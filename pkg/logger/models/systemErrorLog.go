package models

import "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/logger/interfaces"

type SystemErrorLog struct {
	Time               string              `json:"asctime"`
	GroupName          string              `json:"group,omitempty"`
	SeverityStatusName string              `json:"levelname"`
	SeverityStatusNo   int                 `json:"levelno"`
	ErrorMsg           string              `json:"message"`
	Error              string              `json:"error"`
	AdditionalInfo     interfaces.LogValue `json:"additionalInfo"`
	Stack              interface{}         `json:"stack"`
}
