package entities

type Article struct {
	Id        ArticleID `sql:"id"`
	UserID    UserID    `sql:"user_id"`
	Title     string    `sql:"title"`
	Content   string    `sql:"content"`
	CreatedAt string    `sql:"created_at"`
}
