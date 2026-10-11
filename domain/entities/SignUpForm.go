package entities

type SignUpForm struct {
	FirstName  string
	Username   string
	Password   string
	DeviceName string
}

func (f *SignUpForm) IsEmpty() bool {
	if len(f.FirstName) == 0 && len(f.Username) == 0 && len(f.Password) == 0 && len(f.DeviceName) == 0 {
		return true
	}
	return false
}

func (f *SignUpForm) CopyPass(password string) SignUpForm {
	return SignUpForm{
		FirstName:  f.FirstName,
		Username:   f.Username,
		Password:   password,
		DeviceName: f.DeviceName,
	}
}
