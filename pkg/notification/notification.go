package notification

type NotificationProvider interface {
	Send(title string, message string) error
}

type Notification struct {
	Provider NotificationProvider
}

func (n *Notification) Send(title string, message string) error {
	return n.Provider.Send(title, message)
}
