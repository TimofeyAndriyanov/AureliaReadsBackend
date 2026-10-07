package entities

type SignInForm struct {
	Username   string
	Password   string
	DeviceName string
}

func (f SignInForm) IsEmpty() bool {
	if len(f.Username) == 0 && len(f.Password) == 0 && len(f.DeviceName) == 0 {
		return true
	}
	return false
}
