package sendgrid

import (
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/email"
)

var NotificationCh = make(chan usecase.Message, 10)

func PushNotifyJob(jobs []usecase.Message) {
	for _, job := range jobs {
		NotificationCh <- job
	}
}

func InitNotificationWorker() {
	go func() {
		for {
			select {
			case notification := <-NotificationCh:
				notification.Send()
			default:
			}
		}
	}()
}
