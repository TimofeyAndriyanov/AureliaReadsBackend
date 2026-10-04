package entities

type UserCredential struct {
	Id           UserID `sql:"id"`
	HashPassword string `sql:"hash_password"`
}
