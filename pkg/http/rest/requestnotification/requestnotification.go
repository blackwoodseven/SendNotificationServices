package requestnotification

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
)

var (
	ErrAuthTokenEmpty = errors.New("Authorization token can't be empty")
	ErrBadRequest     = errors.New("Bad request, invalid paramter(s)")
	ErrSlackOrEmail   = errors.New("Couldn't deliver the message to slack or email")
)

type RequestResponse struct {
	Message string `json:"message"`
}

func (r *repository) RequestNotification(c *gin.Context) {
	var req requestnotification.RequestNotification
	var authtoken = c.Request.Header["Authorization"][0]

	if len(authtoken) == 0 {
		handleError(c, http.StatusInternalServerError, ErrAuthTokenEmpty)
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

	c.JSON(http.StatusOK, RequestResponse{Message: "Notification delivered to slack and mail"})
}

func ValidateInput(req requestnotification.RequestNotification) bool {

	if req.ComponentID == "" {
		return false
	}

	if strings.ToUpper(req.MediumType) == "SLACK" && req.SlackID == "" {
		return false
	}

	if strings.ToUpper(req.MediumType) == "EMAIL" && req.EmailID == "" {
		return false
	}
	return true
}