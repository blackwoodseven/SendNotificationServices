package requestnotification

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
	domain "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
	domainmock "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification/mocks"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"net/url"
)

type TestRequest struct {
	response     RequestResponse
	Error        string
	RequestInput domain.RequestNotification
	authToken    string
}

func TestRequestNotification(t *testing.T) {
	mock := domainmock.NewDomain(t)
	requestnotification, _ := Init(Input{
		NotificationRequest: mock,
	})
	status, slack, mail := false, false, false
	cases := map[int]TestRequest{
		1: {
			RequestInput: domain.RequestNotification{
				ComponentName: "Mobile-App",
				MediumType:    "Slack",
				SlackID:       "",
				EmailID:       "test-case-1@testcase,com",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
			Error:     "internal server error",
			authToken: "ZnJlZDpmcmVk",
		},
		2: {
			RequestInput: domain.RequestNotification{
				ComponentName: "Mobile-App",
				MediumType:    "Email",
				SlackID:       "Slack-ID",
				EmailID:       "",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
			Error: "internal server error",
		},
		3: {
			RequestInput: domain.RequestNotification{
				ComponentName: "Mobile-App",
				MediumType:    "Slack",
				SlackID:       "Email",
				EmailID:       "test-case-1@testcase,com",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
			Error:     "",
			authToken: "ZnJlZDpmcmVk",
		},
	}

	for i, c := range cases {
		h := httptest.NewRecorder()
		httpcontext := GetTestGinContext(h)
		MockJsonPostWithAuthHeader(httpcontext, c.RequestInput, c.authToken)
		switch i {
		case 1:
			status, slack, mail = false, false, false
			mock.On("RequestNotification", c).Return(false, false, false)
		case 2:
			status, slack = true, true
			mail = false
			mock.On("RequestNotification", c).Return(true, true, false)
		case 3:
			status, slack, mail = true, true, true
			mock.On("RequestNotification", c).Return(true, true, true)
		}
		requestnotification.RequestNotification(httpcontext)
		if status && slack && mail {
			assert.EqualValues(t, http.StatusOK, h.Code)
		} else if status && (slack || mail) {
			assert.EqualValues(t, http.StatusInternalServerError, h.Code)
		} else {
			assert.EqualValues(t, http.StatusBadRequest, h.Code)
		}
	}
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MockJsonPostWithAuthHeader(c *gin.Context, content interface{}, authToken string) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("Authorization", authToken)

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}
