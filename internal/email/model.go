package email

import "github.com/aws/aws-sdk-go/service/ses"

type EmailService interface {
	send(to string, msg EmailMessage) (*ses.SendEmailOutput, error)
	bulkSend()
	sendWithTemplate()
	bulkSendWithTemplate()
	createTemplate()
	SendAccountVerificationEmail(mail, link string)
}

type EmailMessage struct {
	subject string
	body    struct {
		text string
		html string
	}
}
