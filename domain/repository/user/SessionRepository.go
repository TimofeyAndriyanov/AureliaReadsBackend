package user

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type SessionRepository interface {
	Add(uid entities.UserID, deviceName string, tokenHash entities.TokenHash) error
	DoesSessionExist(token entities.TokenHash) (*entities.UserID, error)
	GetSessionsByUserID(id entities.UserID) ([]entities.SessionItem, error)
	Delete(token entities.TokenHash) error
}
