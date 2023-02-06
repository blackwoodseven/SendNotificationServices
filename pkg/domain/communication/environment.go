package communication

import (
	"fmt"
	"os"
)

const (
	ErrKeyNotFound  = "Key %s not found in the environment"
	ErrInvalidInput = "Too many keys passed"
)

func getvaluefromenvironment(keys ...string) (string, string) {
	os.Setenv("MAIL_FROM", "MAIL_FROM")
	os.Setenv("MAIL_ACCESS_TOKEN", "MAIL_ACCESS_TOKEN")
	os.Setenv("SLACK_BOT_TOKEN", "SLACK_BOT_TOKEN")
	os.Setenv("SLACK_CHANNEL_ID", "SLACK_CHANNEL_ID")
	if len(keys) == 2 {
		var keyvalue = [2]string{}
		for i, key := range keys {
			keyvalue[i] = os.Getenv(key)
			if keyvalue[i] == "" {
				panic(fmt.Sprintf(ErrKeyNotFound, (key)))
			}
		}
		return keyvalue[0], keyvalue[1]
	} else {
		panic(ErrInvalidInput)
	}
}
