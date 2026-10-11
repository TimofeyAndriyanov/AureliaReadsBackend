package articles

import (
	"context"
	"database/sql"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	articleRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"
)

type repository struct {
	db      *sql.DB
	context context.Context
}

func NewArticlesRepository(db *sql.DB) articleRepo.Repository {
	return &repository{
		db:      db,
		context: context.Background(),
	}
}

func (i repository) AddArticle(uid entities.UserID, article entities.NewArticle) error {
	query := `
		INSERT INTO articles(
			title,
			content,
			user_id
		) VALUES ($1, $2, $3)
	`

	if err := i.db.QueryRowContext(
		i.context,
		query,
		article.Title,
		article.Content,
		uid,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (i repository) GetArticleById(id entities.ArticleID) (*entities.Article, error) {
	var dest *entities.Article

	query := "SELECT * FROM articles WHERE id=$1"

	if err := i.db.QueryRowContext(
		i.context,
		query,
		id,
	).Scan(&dest); err != nil {
		return nil, err
	}

	return dest, nil
}

func (i repository) DeleteArticleById(id entities.ArticleID, uid entities.UserID) error {
	query := "DELETE FROM articles WHERE id=$1 AND user_id=$2"

	if err := i.db.QueryRowContext(
		i.context,
		query,
		id,
		uid,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (i repository) GetArticlesByUid(uid entities.UserID) ([]entities.Article, error) {
	var dest []entities.Article

	query := "SELECT * FROM articles WHERE user_id = $1"

	if err := i.db.QueryRowContext(i.context, query, uid).Scan(&dest); err != nil {
		return nil, err
	}

	return dest, nil
}

func (i repository) AllArticles() ([]entities.Article, error) {
	var dest []entities.Article

	query := "SELECT * FROM articles"

	if err := i.db.QueryRowContext(i.context, query).Scan(&dest); err != nil {
		return nil, err
	}

	return dest, nil
}
