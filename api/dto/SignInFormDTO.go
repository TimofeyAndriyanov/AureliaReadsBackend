package dto

import "AureliaReadsBackend/domain/entities"

type SignInFormDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d SignInFormDTO) ToDomain() entities.SignInForm {
	return entities.SignInForm{
		Username: d.Username,
		Password: d.Password,
	}
}
