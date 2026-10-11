package dto

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type SignInFormDTO struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	DeviceName string `json:"device_name"`
}

func (d SignInFormDTO) ToDomain() entities.SignInForm {
	return entities.SignInForm{
		Username:   d.Username,
		Password:   d.Password,
		DeviceName: d.DeviceName,
	}
}
