package sha512

import (
	"crypto/sha512"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type hashService struct{}

func NewSHA512HashService() services.HashService {
	return &hashService{}
}

func (i hashService) Hashing(value string) []byte {
	hash := sha512.Sum512([]byte(value))

	return hash[:]
}

func (i hashService) HashChecking(hash []byte, value string) bool {
	panic("")
	//return i.Hashing(value) == hash
}
