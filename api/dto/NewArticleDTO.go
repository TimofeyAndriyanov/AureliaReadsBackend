package dto

import "AureliaReadsBackend/domain/entities"

type NewArticleDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (d NewArticleDTO) ToDomain() entities.NewArticle {
	return entities.NewArticle{
		Title:   d.Title,
		Content: d.Content,
	}
}
