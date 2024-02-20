package singleton

import (
	nfn "go-gin/pkg/notification"
	"go-gin/pkg/utils/parse"
)

var Notifications map[string][]interface{}

func LoadNotifications() {
	Notifications = make(map[string][]interface{})

	for _, _notification := range Conf.Notifications {
		switch _notification.Type {
		case "apprise":
			for _, instance := range _notification.Instances {
				Notifications["apprise"] = append(Notifications["apprise"], nfn.Apprise{Payload: nfn.ApprisePayload{AppriseUrl: instance["url"]}})
			}
		case "dingtalk":
			for _, instance := range _notification.Instances {
				Notifications["dingtalk"] = append(Notifications["dingtalk"], nfn.DingTalk{Payload: nfn.DingTalkPayload{Webhook: instance["webhook"]}})
			}
		case "ifttt":
			for _, instance := range _notification.Instances {
				Notifications["ifttt"] = append(Notifications["ifttt"], nfn.IFTTT{Payload: nfn.IFTTTPayload{Event: instance["event"], Key: instance["key"]}})
			}
		case "telegram":
			for _, instance := range _notification.Instances {
				Notifications["telegram"] = append(Notifications["telegram"], nfn.Telegram{Payload: nfn.TelegramPayload{BotToken: instance["token"], ChatId: instance["chat_id"]}})
			}
		case "wecom":
			for _, instance := range _notification.Instances {
				Notifications["wecom"] = append(Notifications["wecom"], nfn.WeCom{Payload: nfn.WeComPayload{Key: instance["key"]}})
			}
		case "smtp":
			for _, instance := range _notification.Instances {
				Notifications["smtp"] = append(Notifications["smtp"], nfn.SMTP{Payload: nfn.SMTPPayload{Host: instance["host"], Port: parse.ParseInt(instance["port"], 587), Username: instance["username"], Password: instance["password"], From: instance["from"], To: instance["to"]}})
			}
		default:
			Log.Error().Msgf("Unknown notification type: %s", _notification.Type)
		}
	}
}

// SendNotification sends a notification to all instances
func SendNotification(title string, message string) {
	for _, _notification := range Notifications {
		for _, instance := range _notification {
			switch _instance := instance.(type) {
			case nfn.Apprise:
				_instance.Send(title, message)
			case nfn.DingTalk:
				_instance.Send(title, message)
			case nfn.IFTTT:
				_instance.Send(title, message)
			case nfn.Telegram:
				_instance.Send(title, message)
			case nfn.WeCom:
				_instance.Send(title, message)
			case nfn.SMTP:
				_instance.Send(title, message)
			default:
				Log.Error().Msgf("Unknown notification instance: %v", _instance)
			}
		}
	}
}

// SendNotificationByType sends a notification by type， e.g. "apprise", "dingtalk", "ifttt", "telegram", "wecom", "smtp"
func SendNotificationByType(notificationType string, title string, message string) {
	for _, instance := range Notifications[notificationType] {
		switch _instance := instance.(type) {
		case nfn.Apprise:
			_instance.Send(title, message)
		case nfn.DingTalk:
			_instance.Send(title, message)
		case nfn.IFTTT:
			_instance.Send(title, message)
		case nfn.Telegram:
			_instance.Send(title, message)
		case nfn.WeCom:
			_instance.Send(title, message)
		case nfn.SMTP:
			_instance.Send(title, message)
		default:
			Log.Error().Msgf("Unknown notification instance: %v", _instance)
		}
	}
}
