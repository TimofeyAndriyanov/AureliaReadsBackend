package sign_in

import "AureliaReadsBackend/domain/entities"

type SignInResult interface {
	sealed()
}

type Success struct {
	Data entities.JwtTokens
}

func (s Success) sealed() {}

type UserNotFound struct{}

func (s UserNotFound) sealed() {}

type WrongPassword struct{}

func (s WrongPassword) sealed() {}

type EmptyFields struct{}

func (s EmptyFields) sealed() {}
