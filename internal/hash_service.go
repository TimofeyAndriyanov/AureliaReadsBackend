package internal

type HashService interface {
	Hashing(val []byte) ([]byte, error)
	HashChecking(hash, val []byte) bool
}
