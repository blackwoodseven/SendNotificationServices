package requestnotification

import (
	"errors"
	"testing"

	domainmock "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/storage/db/notificationrequest/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestRequest struct {
	requestnotification RequestNotification
}

func TestRequestNotification(t *testing.T) {
	rmock := domainmock.NewRepository(t)
	notificationreq, _ := Init(Input{
		NotificationRequest: rmock,
	})

	var err error = nil
	cases := map[int]TestRequest{
		1: {
			requestnotification: RequestNotification{

				ComponentName: "Mobile-App",
				MediumType:    "Slack",
				SlackID:       "",
				EmailID:       "test-case-1@testcase,com",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
		},
		2: {
			requestnotification: RequestNotification{
				ComponentID:   "2",
				ComponentName: "Mobile-App",
				MediumType:    "Email",
				SlackID:       "Slack-ID",
				EmailID:       "",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
		},
		3: {
			requestnotification: RequestNotification{
				ComponentName: "Mobile-App",
				MediumType:    "Slack",
				SlackID:       "Email",
				EmailID:       "test-case-1@testcase,com",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
		},
	}
	for i, c := range cases {
		switch i {
		case 1:
			rmock.On("ComponentExists", mock.Anything, mock.Anything).Return(false).Once()
			rmock.On("Create", mock.Anything).Return(nil).Once()
		case 2:
			err = errors.New("error")
			rmock.On("ComponentExists", mock.Anything, mock.Anything).Return(false).Once()
			rmock.On("Create", mock.Anything).Return(err).Once()
		case 3:
			rmock.On("ComponentExists", mock.Anything, mock.Anything).Return(false).Once()
			rmock.On("Create", mock.Anything).Return(nil).Once()
		}

		status, slack, mail := notificationreq.RequestNotification(c.requestnotification)

		if err == nil {
			assert.True(t, status)
			assert.False(t, slack)
			assert.False(t, mail)
		} else {
			assert.False(t, status)
			assert.False(t, slack)
			assert.False(t, mail)
			err = nil
		}
	}
}
