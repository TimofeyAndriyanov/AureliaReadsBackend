package dto

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type SessionTokenDTO struct {
	Token string `json:"token"`
}

func ToDTO(token entities.SessionToken) SessionTokenDTO {
	return SessionTokenDTO{Token: string(token)}
}
