package sendgrid

import (
	"fmt"

	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/email"
)

func NewUserRegistrationMessage(name string) usecase.Message {
	return &UserRegistrationMessage{
		name: name,
	}
}

type UserRegistrationMessage struct {
	name string
}

func (urm *UserRegistrationMessage) Send() {
	message := urm.Format()

	// 実際にメール送信する部分
	fmt.Println("Sending email:", message)
}

func (urm *UserRegistrationMessage) Format() string {
	return fmt.Sprintf("Hello %s, thank you for registering!", urm.name)
}
