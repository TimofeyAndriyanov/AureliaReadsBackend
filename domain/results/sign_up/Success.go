package sign_up

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type Success struct {
	Data entities.SessionToken
}

func (s Success) sealed() {}
