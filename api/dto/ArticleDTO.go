package dto

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type ArticleDTO struct {
	Id        entities.ArticleID `json:"id"`
	UserID    entities.UserID    `json:"user_id"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	CreatedAt string             `json:"created_at"`
}
