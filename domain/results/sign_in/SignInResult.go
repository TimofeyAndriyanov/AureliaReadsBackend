package sign_in

type SignInResult interface {
	sealed()
}

type Success struct {
	Data any
}

func (s Success) sealed() {}

type UserNotFound struct{}

func (s UserNotFound) sealed() {}

type WrongPassword struct{}

func (s WrongPassword) sealed() {}

type EmptyFields struct{}

func (s EmptyFields) sealed() {}
