package requestnotification

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/auth"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/communication"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/utility"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/storage/db/notificationrequest"
)

var EncncryptionKey = "ENCRYPTION_KEY"

func (d *domain) RequestNotification(requestnotification RequestNotification) (bool, bool, bool) {

	models := notificationrequest.NotificationRequest{
		ComponentID:   requestnotification.ComponentID,
		ComponentName: requestnotification.ComponentName,
	}

	communicationmodel := communication.CommunicationModel{
		MediumType: requestnotification.MediumType,
		SlackID:    requestnotification.SlackID,
		EmailID:    requestnotification.EmailID,
		Subject:    requestnotification.Subject,
		Message:    requestnotification.Message,
	}

	models.AuthToken = utility.Encrypt(hex.EncodeToString([]byte(utility.GetRamdomKey())), requestnotification.AuthToken)
	err := d.NotificationRequest.Create(&models)
	fmt.Println("err: ", err)
	if err != nil {
		return false, false, false
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
		if strings.ToUpper(communicationmodel.MediumType) == "SLACK" {
			emaildelivered = true
			slackmessagedelivered = communication.SendSlackMessage(communicationmodel)
		}
		if strings.ToUpper(communicationmodel.MediumType) == "MAIL" {
			slackmessagedelivered = true
			emaildelivered = communication.SendEmail(communicationmodel)
		}
		if strings.ToUpper(communicationmodel.MediumType) == "ALL" {
			slackmessagedelivered = communication.SendSlackMessage(communicationmodel)
			emaildelivered = communication.SendEmail(communicationmodel)
		}
	}

	return slackmessagedelivered, emaildelivered
}

func (d *domain) ValidateToken(authToken string) bool {

	if len(authToken) == 0 {
		return false
	}
	claims, err := auth.ValidateGSuiteToken(authToken)
	if err != nil {
		return false
	}
	if claims != nil && claims.Hd != "blackwoodseven.com" && claims.Hd != "kantar.com" {
		return false
	}

	return true
}
