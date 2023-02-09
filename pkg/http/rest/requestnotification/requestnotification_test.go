package requestnotification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	domain "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
	domainmock "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification/mocks"
	"github.com/stretchr/testify/assert"
)

type TestRequest struct {
	response        RequestResponse
	Error           string
	RequestInput    domain.RequestNotification
	authToken       string
	isValidToken    bool
	inputValidation bool
}

func TestRequestNotification(t *testing.T) {
	dmock := domainmock.NewDomain(t)
	requestnotification, _ := Init(Input{
		NotificationRequest: dmock,
	})
	status, slack, mail := false, false, false
	cases := map[int]TestRequest{
		1: {
			RequestInput: domain.RequestNotification{
				ComponentID:   "12",
				ComponentName: "Mobile-App",
				MediumType:    "Slack",
				SlackID:       "Email",
				EmailID:       "test-case-1@testcase,com",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
			Error:           "internal server error",
			authToken:       "",
			isValidToken:    false,
			inputValidation: true,
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
			Error:           "internal server error",
			authToken:       "ZnJlZDpmcmVk",
			isValidToken:    true,
			inputValidation: false,
		},
		3: {
			RequestInput: domain.RequestNotification{
				ComponentID:   "12",
				ComponentName: "Mobile-App",
				MediumType:    "Slack",
				SlackID:       "Email",
				EmailID:       "test-case-1@testcase,com",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
			Error:           "",
			authToken:       "ZnJlZDpmcmVk",
			isValidToken:    true,
			inputValidation: true,
		},
		4: {
			RequestInput: domain.RequestNotification{
				ComponentID:   "12",
				ComponentName: "Mobile-App",
				MediumType:    "Slack",
				SlackID:       "Email",
				EmailID:       "test-case-1@testcase,com",
				Subject:       "test-case-subject",
				Message:       "test-case-message-body",
			},
			Error:           "",
			authToken:       "ZnJlZDpmcmVk",
			isValidToken:    true,
			inputValidation: true,
		},
	}

	for i, c := range cases {
		h := httptest.NewRecorder()
		httpcontext := GetTestGinContext(h)
		MockJsonPostWithAuthHeader(httpcontext, c.RequestInput, c.authToken)
		switch i {
		case 1:
			status, slack, mail = false, false, false
		case 2:
			status, slack = true, true
			mail = false
		case 3:
			status, slack, mail = true, true, true
			dmock.On("RequestNotification", c.RequestInput).Return(status, slack, mail).Once()

		case 4:
			status, slack, mail = true, true, false
			dmock.On("RequestNotification", c.RequestInput).Return(status, slack, mail).Once()
		}
		requestnotification.RequestNotification(httpcontext)
		if status && slack && mail && c.isValidToken && c.inputValidation {
			assert.EqualValues(t, http.StatusOK, h.Code)
			fmt.Println(http.StatusOK, h.Code)
		} else if !c.isValidToken {
			assert.EqualValues(t, http.StatusUnauthorized, h.Code)
			fmt.Println(http.StatusUnauthorized, h.Code)
		} else if status && (slack || mail) && c.inputValidation {
			assert.EqualValues(t, http.StatusInternalServerError, h.Code)
			fmt.Println(http.StatusInternalServerError, h.Code)
		} else {
			assert.EqualValues(t, http.StatusBadRequest, h.Code)
			fmt.Println(http.StatusBadRequest, h.Code)
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
