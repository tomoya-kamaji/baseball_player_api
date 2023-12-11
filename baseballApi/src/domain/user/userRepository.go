package domain

import "context"

type UserRepository interface {
	Save(c context.Context, u *User) error
	GetByID(c context.Context, id UserID) (*User, error)
}
