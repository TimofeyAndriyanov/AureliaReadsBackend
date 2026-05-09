package dto

import "AureliaReadsBackend/domain/entities"

type SignUpFormDTO struct {
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (d SignUpFormDTO) ToDomain() entities.SignUpForm {
	return entities.SignUpForm{
		FirstName: d.FirstName,
		Username:  d.Username,
		Password:  d.Password,
	}
}
