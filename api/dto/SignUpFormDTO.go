package dto

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type SignUpFormDTO struct {
	FirstName  string `json:"first_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	DeviceName string `json:"device_name"`
}

func (d SignUpFormDTO) ToDomain() entities.SignUpForm {
	return entities.SignUpForm{
		FirstName:  d.FirstName,
		Username:   d.Username,
		Password:   d.Password,
		DeviceName: d.DeviceName,
	}
}
