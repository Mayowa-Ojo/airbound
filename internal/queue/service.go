package queue

import (
	"airbound/internal/config"
	"airbound/internal/event"
	"airbound/internal/log"
	"airbound/repository/flights"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

type Service struct {
	reservationsQueue string
	delay             int64
	visibility        int64
	messagePerRequest int64
	backpressureDelay int
	svc               *sqs.SQS // service client
	emitter           *event.Emitter
}

func NewQueueService(sqss *session.Session, cfg *config.Config, emitter *event.Emitter) QueueService {
	s := &Service{
		reservationsQueue: cfg.AWSSQSReservationsQueueName,
		delay:             int64(cfg.AWSSQSDelay),
		visibility:        int64(cfg.AWSSQSVisibility * 60 * 60),
		messagePerRequest: int64(1),
		emitter:           emitter,
	}

	s.svc = sqs.New(sqss)
	s.backpressureDelay = 700

	go s.startSQSListner(context.Background())

	return s
}

func (s *Service) createQueue(ctx context.Context) {}

func (s *Service) listQueues(ctx context.Context) (*sqs.ListQueuesOutput, error) {
	res, err := s.svc.ListQueuesWithContext(ctx, &sqs.ListQueuesInput{
		MaxResults: aws.Int64(10),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) deleteQueue(ctx context.Context) {}

func (s *Service) sendMessage(ctx context.Context, msg *Message, queueURL string) (*sqs.SendMessageOutput, error) {
	input := &sqs.SendMessageInput{
		DelaySeconds:           aws.Int64(s.delay),
		MessageAttributes:      msg.attributes,
		MessageBody:            aws.String(msg.body),
		MessageDeduplicationId: aws.String(msg.deduplicationID),
		MessageGroupId:         aws.String(msg.groupID),
		QueueUrl:               aws.String(queueURL),
	}

	res, err := s.svc.SendMessageWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) receiveMessage(ctx context.Context, queueURL string) (*sqs.ReceiveMessageOutput, error) {
	receiveRequestAttemptID := uuid.New()

	input := &sqs.ReceiveMessageInput{
		QueueUrl:                &queueURL,
		MaxNumberOfMessages:     aws.Int64(s.messagePerRequest),
		ReceiveRequestAttemptId: aws.String(receiveRequestAttemptID.String()),
		VisibilityTimeout:       aws.Int64(s.visibility),
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			aws.String(sqs.MessageSystemAttributeNameApproximateFirstReceiveTimestamp),
			aws.String(sqs.MessageSystemAttributeNameMessageDeduplicationId),
		},
		MessageAttributeNames: []*string{aws.String(sqs.QueueAttributeNameAll)},
	}

	res, err := s.svc.ReceiveMessageWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	go func(msgs []*sqs.Message) {
		logger := log.WithField(string(log.LogFieldFunctionName), "<QueueService>deleteMessage")

		for _, msg := range msgs {
			if _, err := s.deleteMessage(ctx, queueURL, *msg.ReceiptHandle); err != nil {
				logger.Error("error occurred while attempting to delete message - %s", err)
			}
		}
	}(res.Messages)

	return res, nil
}

func (s *Service) deleteMessage(ctx context.Context, queueURL, messageHandle string) (*sqs.DeleteMessageOutput, error) {
	input := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(messageHandle),
	}

	res, err := s.svc.DeleteMessageWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) getQueueURL(ctx context.Context, queueName string) (*string, error) {
	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}

	res, err := s.svc.GetQueueUrlWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	return res.QueueUrl, nil
}

func (s *Service) pollMessages(ctx context.Context, ch chan<- *ChannelMessage) {
	defer close(ch)

	for {
		res, err := s.ReceiveFlightReservationMessage(ctx)
		if err != nil {
			continue
		}

		for _, msg := range res.Messages {
			ch <- &ChannelMessage{queue: FlightReservationsQueue, Message: msg}
		}

		s.backpressure()
	}
}

func (s *Service) startSQSListner(ctx context.Context) {
	msgChan := make(chan *ChannelMessage, 2)

	go s.pollMessages(ctx, msgChan)

	for msg := range msgChan {
		switch msg.queue {
		case FlightReservationsQueue:
			s.emitter.Dispatch(ctx, event.CONFIRM_FLIGHT_RESERVATION, msg)
		}
	}
}

func (s *Service) backpressure() {
	time.Sleep(time.Millisecond * time.Duration(s.backpressureDelay))
}

func (s *Service) SendFlightReservationMessage(ctx context.Context, fr *flights.FlightReservation, customerID uuid.UUID) error {
	logger := log.WithField(string(log.LogFieldFunctionName), "<QueueService>SendFlightReservationMessage")

	queueURL, err := s.getQueueURL(ctx, s.reservationsQueue)
	if err != nil {
		logger.Error("error occurred while fetching queue url - %s", err)
		return err
	}

	body := fmt.Sprintf("Flight reservation for customer - %s", "")

	msg := &Message{
		deduplicationID: fr.ID.String(),
		groupID:         fr.ID.String(),
		attributes: map[string]*sqs.MessageAttributeValue{
			"ReservationStatus": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(fr.ReservationStatus)),
			},
			"ReservationNumber": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(fr.ReservationNumber)),
			},
			"FlightInstanceID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(fr.FlightInstance.ID.String())),
			},
			"CustomerID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(customerID.String())),
			},
		},
		body: body,
	}

	if _, err := s.sendMessage(ctx, msg, *queueURL); err != nil {
		logger.Error("error occurred while sending flight reservation message - %s", err)
		return err
	}

	return nil
}

func (s *Service) ReceiveFlightReservationMessage(ctx context.Context) (*sqs.ReceiveMessageOutput, error) {
	logger := log.WithField(string(log.LogFieldFunctionName), "<QueueService>ReceiveFlightReservationMessage")

	queueURL, err := s.getQueueURL(ctx, s.reservationsQueue)
	if err != nil {
		logger.Error("error occurred while fetching queue url - %s", err)
		return nil, err
	}

	res, err := s.receiveMessage(ctx, *queueURL)
	if err != nil {
		logger.Error("error occured while attempting to receive message - %s", err)
		return nil, err
	}

	return res, nil
}
