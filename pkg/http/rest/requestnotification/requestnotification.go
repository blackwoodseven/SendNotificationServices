package requestnotification

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
)

var (
	ErrUnAuthorizedToken = errors.New("Unauthorized Token")
	ErrBadRequest        = errors.New("Bad request, invalid paramter(s)")
	ErrSlackOrEmail      = errors.New("Couldn't deliver the message to slack or email")
)

const (
	Slack           string = "SLACK"
	Mail            string = "MAIL"
	All             string = "ALL"
	Authorization   string = "Authorization"
	ResponseMessage string = "Notification delivered to slack and mail"
)

type RequestResponse struct {
	Message string `json:"message"`
}

func (r *repository) RequestNotification(c *gin.Context) {
	var req requestnotification.RequestNotification
	var authtoken = c.Request.Header[Authorization][0]

	if len(authtoken) == 0 {
		handleError(c, http.StatusUnauthorized, ErrUnAuthorizedToken)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	if !ValidateInput(req) {
		handleError(c, http.StatusBadRequest, ErrBadRequest)
		return
	}

	requestNotification, slackmessagedelivered, emaildelivered := r.NotificationRequest.RequestNotification(req)

	if !requestNotification {
		handleError(c, http.StatusBadRequest, ErrBadRequest)
		return
	}
	if !slackmessagedelivered || !emaildelivered {
		handleError(c, http.StatusInternalServerError, ErrSlackOrEmail)
		return
	}

	c.JSON(http.StatusOK, RequestResponse{Message: ResponseMessage})
}

func ValidateInput(req requestnotification.RequestNotification) bool {

	if req.ComponentID == "" {
		return false
	}

	if strings.ToUpper(req.MediumType) == Slack && req.SlackID == "" {
		return false
	}

	if strings.ToUpper(req.MediumType) == Mail && req.EmailID == "" {
		return false
	}

	switch strings.ToUpper(req.MediumType) {
	case Mail:
		return true
	case Slack:
		return true
	case All:
		return true
	default:
		return false
	}
}
