package repository

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type UserSessionRepository interface {
	Add() error
	DoesSessionExist(token entities.TokenHash) (bool, error)
	GetSessionsByUserID(id entities.UserID) ([]entities.SessionItem, error)
	Delete(id entities.SessionID) error
}
