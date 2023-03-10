package communication

import (
	"crypto/tls"

	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/logger"

	"gopkg.in/gomail.v2"
)

const (
	MailFromKey        = "MAIL_FROM"
	MailAccessTokenKey = "MAIL_ACCESS_TOKEN"
	MailTo             = "To"
	MailFrom           = "From"
	MailSubject        = "Subject"
	MailContentType    = "text/html"
	MailSMTP           = "smtp.gmail.com"
	MailSMPTPort       = 587
)

func SendEmail(communicationmodel CommunicationModel) bool {

	mailfrom, mailcredentials := getvaluefromenvironment(MailFromKey, MailAccessTokenKey)
	email := gomail.NewMessage()
	email.SetHeader(MailTo, communicationmodel.EmailID)
	email.SetHeader(MailFrom, mailfrom)
	email.SetHeader(MailSubject, communicationmodel.Subject)
	email.SetBody(MailContentType, communicationmodel.Message)

	n := gomail.NewDialer(MailSMTP, MailSMPTPort, mailfrom, mailcredentials)
	n.TLSConfig = &tls.Config{InsecureSkipVerify: false}
	if err := n.DialAndSend(email); err != nil {
		logger.Log("", logger.SeverityCritical, "Error while delivering mail", err, nil)
		return false
	}
	return true
}
