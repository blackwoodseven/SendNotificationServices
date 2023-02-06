package requestnotification

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
)

var (
	ErrInvalidInputNotificationRequest = errors.New("invalid input. NotificationRequest domain missing")
)

type repository struct {
	Input
}

type Repository interface {
	RequestNotification(c *gin.Context)
}

type Input struct {
	NotificationRequest requestnotification.Domain
}

func validateInput(input Input) error {
	if input.NotificationRequest == nil {
		return ErrInvalidInputNotificationRequest
	}

	return nil
}

func Init(input Input) (Repository, error) {
	if err := validateInput(input); err != nil {
		return nil, err
	}

	return &repository{
		input,
	}, nil
}
