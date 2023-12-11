package domain

import "github.com/google/uuid"

type UserID string

func (id UserID) String() string {
	return string(id)
}

func NewUserId() UserID {
	userId := UserID(uuid.New().String())
	return userId
}

type User struct {
	ID            UserID
	Name          string
	Address       string
	ContactNumber string
}

type ReConstractUserParam struct {
	ID            UserID
	Name          string
	Address       string
	ContactNumber string
}

func ReConstractUser(param ReConstractUserParam) *User {
	return &User{
		ID:            param.ID,
		Name:          param.Name,
		Address:       param.Address,
		ContactNumber: param.ContactNumber,
	}
}
