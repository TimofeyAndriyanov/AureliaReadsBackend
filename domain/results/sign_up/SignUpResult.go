package sign_up

import "AureliaReadsBackend/domain/entities"

type SignUpResult interface {
	sealed()
}

type Success struct {
	Data entities.JwtTokens
}

func (s Success) sealed() {}

type UserAlreadyExists struct{}

func (s UserAlreadyExists) sealed() {}

type EmptyFields struct{}

func (s EmptyFields) sealed() {}
