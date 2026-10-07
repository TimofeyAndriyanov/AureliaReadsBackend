package user

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type SessionRepository interface {
	Add(uid entities.UserID, deviceName string) (entities.SessionToken, error)
	DoesSessionExist(token entities.TokenHash) (bool, error)
	GetSessionsByUserID(id entities.UserID) ([]entities.SessionItem, error)
	Delete(token entities.TokenHash) error
}
