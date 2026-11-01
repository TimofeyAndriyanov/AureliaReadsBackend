package verify_session_token

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type Success struct {
	Data entities.UserID
}

func (s Success) sealed() {}
