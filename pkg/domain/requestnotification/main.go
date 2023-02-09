package requestnotification

import (
	"errors"

	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/communication"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/storage/db/notificationrequest"
)

type RequestNotification struct {
	ComponentID   string `json:"componentid"`
	AuthToken     string `json:"idtoken"`
	ComponentName string `json:"componentname"`
	MediumType    string `json:"mediumtype"`
	SlackID       string `json:"slackid"`
	EmailID       string `json:"emailid"`
	Subject       string `json:"subject"`
	Message       string `json:"message"`
}

type domain struct {
	Input
	incoming chan communication.CommunicationModel
}

type Input struct {
	NotificationRequest notificationrequest.Repository
}

type Domain interface {
	RequestNotification(requestnotification RequestNotification) (bool, bool, bool)
	Scheduler()
}

var (
	ErrNotFound                        = errors.New("Not found")
	ErrInvalidInputRequestNotification = errors.New("Invalid input. Request Notification repository missing")
)

func validateInput(input Input) error {
	if input.NotificationRequest == nil {
		return ErrInvalidInputRequestNotification
	}

	return nil
}

func Init(input Input) (Domain, error) {
	if err := validateInput(input); err != nil {
		return nil, err
	}

	d := &domain{
		input,
		make(chan communication.CommunicationModel),
	}

	return d, nil
}
