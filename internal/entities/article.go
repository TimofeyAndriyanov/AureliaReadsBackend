package entities

import (
	"time"
	"uuid"
)

type Article struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
}
