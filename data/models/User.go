package models

import "AureliaReadsBackend/domain/entities"

type User struct {
	Id        entities.UserID
	FirstName string
	Username  string
	Password  string
}

func (u *User) ToUserCredential() *entities.UserCredential {
	return &entities.UserCredential{
		Id:           u.Id,
		HashPassword: u.Password,
	}
}
