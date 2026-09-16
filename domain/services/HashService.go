package services

type HashService interface {
	Hashing(value string) []byte

	HashChecking(hash []byte, value string) bool
}
