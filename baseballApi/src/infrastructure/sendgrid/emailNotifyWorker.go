package sendgrid

import (
	"fmt"
	"time"

	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/email"
)

var NotificationCh = make(chan usecase.Message, 10)

func PushNotifyJob(job usecase.Message) {
	NotificationCh <- job
}

func InitNotificationWorker() {
	for {
		select {
		case notification := <-NotificationCh:
			notification.Send()
		default:
			// 他のロジックや待機処理
			fmt.Println("Notification server running...")
			time.Sleep(1 * time.Second)
		}
	}
}
