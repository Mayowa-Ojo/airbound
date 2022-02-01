package queue

import (
	"airbound/repository/flights"
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

// model for a FIFO queue that handles flight reservations
type QueueService interface {
	createQueue(ctx context.Context)
	listQueues(ctx context.Context) (*sqs.ListQueuesOutput, error)
	deleteQueue(ctx context.Context)
	sendMessage(ctx context.Context, msg *Message, queueURL string) (*sqs.SendMessageOutput, error)
	receiveMessage(ctx context.Context, queueURL string) (*sqs.ReceiveMessageOutput, error)
	deleteMessage(ctx context.Context, queueURL, messageHandle string) (*sqs.DeleteMessageOutput, error)
	getQueueURL(ctx context.Context, queueName string) (*string, error)

	SendFlightReservationMessage(ctx context.Context, fr *flights.FlightReservation, customerID uuid.UUID) error
	ReceiveFlightReservationMessage(ctx context.Context) (*sqs.ReceiveMessageOutput, error)
}

type Message struct {
	attributes      map[string]*sqs.MessageAttributeValue
	body            string
	deduplicationID string
	groupID         string
}

type MessageQueue string

const (
	FlightReservationsQueue MessageQueue = "FLIGHT_RESERVATIONS_QUEUE"
)

type ChannelMessage struct {
	queue MessageQueue
	*sqs.Message
}
