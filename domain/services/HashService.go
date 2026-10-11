package services

type HashService interface {
	Hashing(value string) (string, error)

	HashChecking(hash, value string) bool
}
