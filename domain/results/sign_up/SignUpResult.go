package sign_up

type SignUpResult interface {
	sealed()
}

type Success struct {
	Data any
}

func (s Success) sealed() {}

type UserAlreadyExists struct{}

func (s UserAlreadyExists) sealed() {}

type EmptyFields struct{}

func (s EmptyFields) sealed() {}
