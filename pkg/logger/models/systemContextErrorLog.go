package models

type SystemContextErrorLog struct {
	Time               string      `json:"asctime"`
	SeverityStatusName string      `json:"levelname"`
	SeverityStatusNo   int         `json:"levelno"`
	ErrorMsg           string      `json:"message"`
	Error              string      `json:"error"`
	Stack              interface{} `json:"stack,omitempty"`
}
