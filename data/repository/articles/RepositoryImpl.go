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

//
// CREATE TABLE users(
//     id UUID PRIMARY KEY DEFAULT random_uuid(),
//     username TEXT UNIQUE NOT NULL,
//     first_name TEXT NOT NULL,
//     password_hash BYTEA NOT NULL,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
// );
//
// CREATE TABLE user_sessions (
//     id UUID PRIMARY KEY DEFAULT random_uuid(),
//     user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
//     token_hash BYTEA NOT NULL,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
// );
//
// CREATE TABLE articles(
//     id UUID PRIMARY KEY DEFAULT random_uuid(),
//     user_id UUID REFERENCES users(id) ON DELETE CASCADE,
//     title TEXT NOT NULL,
//     content TEXT NOT NULL,
//     is_hide BOOLEAN NOT NULL DEFAULT false,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
// )
//

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
