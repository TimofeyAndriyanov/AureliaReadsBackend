package services

type HashService interface {
	Hashing(value string) string

	HashChecking(hash string, value string) bool
}
