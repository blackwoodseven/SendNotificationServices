package models

import (
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/logger/interfaces"
)

type SystemInfoLog struct {
	Time               string              `json:"asctime"`
	GroupName          string              `json:"group,omitempty"`
	Message            string              `json:"message,omitempty"`
	SeverityStatusName string              `json:"levelname"`
	SeverityStatusNo   int                 `json:"levelno"`
	AdditionalInfo     interfaces.LogValue `json:"additionalInfo"`
}
