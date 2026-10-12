package dto

import (
	"encoding/base64"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
)

type SessionTokenDTO struct {
	Token string `json:"token"`
}

func ToDTO(token entities.SessionToken) SessionTokenDTO {
	base63Token := base64.StdEncoding.EncodeToString(token)
	return SessionTokenDTO{Token: base63Token}
}
