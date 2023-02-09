package communication

import (
	"errors"

	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/logger"
	"github.com/slack-go/slack"
)

const (
	SlackBotTokenEnvKey = "SLACK_BOT_TOKEN"
	SlackClientEnvKey   = "SLACK_CHANNEL_ID"
	SlackElementType    = "mrkdwn"
	SlackErrorMessage   = "Failed to send message to slack channel. Error: %s"
)

var (
	ErrSlackBotToken = errors.New("Slack bot token not found")
	ErrSlackClientID = errors.New("Slack client id not found")
)

func SendSlackMessage(communicationmodel CommunicationModel) bool {

	slackbottoken, slackchannelid := getvaluefromenvironment(SlackBotTokenEnvKey, SlackClientEnvKey)
	msgText := slack.NewTextBlockObject("mrkdwn", communicationmodel.Message, false, false)
	msgSection := slack.NewSectionBlock(msgText, nil, nil)
	msg := slack.MsgOptionBlocks(
		msgSection,
	)

	slackapi := slack.New(slackbottoken)

	_, _, _, err := slackapi.SendMessage(slackchannelid, msg)
	if err != nil {
		logger.Log("", logger.SeverityCritical, "Error while sending message to slack", err, nil)
		return false
	}
	return true
}
