package email

import (
	"airbound/internal/config"
	"airbound/internal/log"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Service struct {
	From    string
	Name    string
	ARN     string
	Charset string
	svc     *ses.SES
}

func NewEmailService(sess *session.Session, cfg *config.Config) EmailService {
	s := &Service{
		From:    fmt.Sprintf(`"%s" <%s>`, cfg.EmailFromName, cfg.EmailFromAddress),
		Name:    "",
		ARN:     cfg.AWSSESARN,
		Charset: "UTF-8",
	}

	s.svc = ses.New(sess)

	return s
}

func (s *Service) send(to string, msg EmailMessage) (*ses.SendEmailOutput, error) {
	input := &ses.SendEmailInput{}

	input.Destination = &ses.Destination{}
	input.Destination.ToAddresses = []*string{aws.String(to)}

	input.Message = &ses.Message{}
	input.Message.Subject = &ses.Content{
		Charset: aws.String(s.Charset),
		Data:    aws.String(msg.subject),
	}

	input.Message.Body = &ses.Body{}
	input.Message.Body.Text = &ses.Content{
		Charset: aws.String(s.Charset),
		Data:    aws.String(msg.body.text),
	}
	input.Message.Body.Html = &ses.Content{
		Charset: aws.String(s.Charset),
		Data:    aws.String(msg.body.html),
	}

	input.Source = aws.String(s.From)
	input.SourceArn = aws.String(s.ARN)

	res, err := s.svc.SendEmail(input)
	if err != nil {
		return nil, err.(awserr.Error)
	}

	return res, nil
}

func (s *Service) bulkSend() {}

func (s *Service) sendWithTemplate() {}

func (s *Service) bulkSendWithTemplate() {}

func (s *Service) createTemplate() {}

func (s *Service) SendAccountVerificationEmail(mail, link string) {

	msg := EmailMessage{
		subject: "Complete Account Registration",
		body: struct {
			text string
			html string
		}{
			html: fmt.Sprintf(`Dear customer,<br/> To complete your account setup, kindly click the button below.<br/><br/><div><a href="%s" style="display:inline-block;box-sizing:border-box;text-decoration:none;font-size:1rem;padding:0.8rem 2.25rem;background:#fbfcfd;color:#444;text-decoration:none;font-weight:bold;border-radius:6px;background-color:#509ee3;color:#fff">Complete Registration<a/></div>`, link),
			text: "Dear customer,<br/> To complete your account setup, kindly click the button below.",
		},
	}

	res, err := s.send(mail, msg)
	if err != nil {
		log.Fatal("[SES]: error sending email - s", err)
	}

	log.Info("[SES]: email sent successfully - id %s", *res.MessageId)
}
