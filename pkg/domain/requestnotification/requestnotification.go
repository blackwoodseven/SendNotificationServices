package requestnotification

import (
	"encoding/hex"
	"fmt"

	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/communication"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/utility"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/storage/db/notificationrequest"
)

const (
	Slack string = "SLACK"
	Mail  string = "MAIL"
	All   string = "All"
)

func (d *domain) RequestNotification(requestnotification RequestNotification) (bool, bool, bool) {

	models := getNotificationRequest(requestnotification)
	communicationmodel := getCommunicationModel(requestnotification)

	if !d.NotificationRequest.ComponentExists(models.ComponentID, models.ComponentName) {
		models.AuthToken = utility.Encrypt(hex.EncodeToString([]byte(utility.GetRamdomKey())), requestnotification.AuthToken)
		err := d.NotificationRequest.Create(&models)
		fmt.Println("err: ", err)
		if err != nil {
			return false, false, false
		}
	}

	slackmessagedelivered, emaildelivered := false, false

	go func() {
		slackmessagedelivered, emaildelivered = d.communication()
	}()
	d.incoming <- communicationmodel

	return true, slackmessagedelivered, emaildelivered

}

func (d *domain) communication() (bool, bool) {

	slackmessagedelivered, emaildelivered := false, false

	for communicationmodel := range d.incoming {

		switch communicationmodel.MediumType {

		case Slack:
			emaildelivered = true
			slackmessagedelivered = communication.SendSlackMessage(communicationmodel)
		case Mail:
			slackmessagedelivered = true
			emaildelivered = communication.SendEmail(communicationmodel)
		case All:
			slackmessagedelivered = communication.SendSlackMessage(communicationmodel)
			emaildelivered = communication.SendEmail(communicationmodel)
		}
	}

	return slackmessagedelivered, emaildelivered
}

func getCommunicationModel(requestnotification RequestNotification) communication.CommunicationModel {

	return communication.CommunicationModel{
		MediumType: requestnotification.MediumType,
		SlackID:    requestnotification.SlackID,
		EmailID:    requestnotification.EmailID,
		Subject:    requestnotification.Subject,
		Message:    requestnotification.Message,
	}
}

func getNotificationRequest(requestnotification RequestNotification) notificationrequest.NotificationRequest {

	return notificationrequest.NotificationRequest{
		ComponentID:   requestnotification.ComponentID,
		ComponentName: requestnotification.ComponentName,
	}
}
