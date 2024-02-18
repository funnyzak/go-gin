package singleton

import (
	nn "go-gin/pkg/notification"
	"go-gin/pkg/utils"
)

var Notifications map[string][]interface{}

func LoadNotifications() {
	Notifications = make(map[string][]interface{})

	for _, _notification := range Conf.Notifications {
		switch _notification.Type {
		case "apprise":
			for _, instance := range _notification.Instances {
				Notifications["apprise"] = append(Notifications["apprise"], nn.Apprise{Payload: nn.ApprisePayload{AppriseUrl: instance["url"]}})
			}
		case "dingtalk":
			for _, instance := range _notification.Instances {
				Notifications["dingtalk"] = append(Notifications["dingtalk"], nn.DingTalk{Payload: nn.DingTalkPayload{Webhook: instance["webhook"]}})
			}
		case "ifttt":
			for _, instance := range _notification.Instances {
				Notifications["ifttt"] = append(Notifications["ifttt"], nn.IFTTT{Payload: nn.IFTTTPayload{Event: instance["event"], Key: instance["key"]}})
			}
		case "telegram":
			for _, instance := range _notification.Instances {
				Notifications["telegram"] = append(Notifications["telegram"], nn.Telegram{Payload: nn.TelegramPayload{BotToken: instance["token"], ChatId: instance["chat_id"]}})
			}
		case "wecom":
			for _, instance := range _notification.Instances {
				Notifications["wecom"] = append(Notifications["wecom"], nn.WeCom{Payload: nn.WeComPayload{Key: instance["key"]}})
			}
		case "smtp":
			for _, instance := range _notification.Instances {
				Notifications["smtp"] = append(Notifications["smtp"], nn.SMTP{Payload: nn.SMTPPayload{Host: instance["host"], Port: utils.ParseInt(instance["port"], 587), Username: instance["username"], Password: instance["password"], From: instance["from"], To: instance["to"]}})
			}
		default:
			Log.Error().Msgf("Unknown notification type: %s", _notification.Type)
		}
	}
}
