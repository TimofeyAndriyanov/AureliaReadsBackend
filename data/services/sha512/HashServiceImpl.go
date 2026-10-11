package sha512

import (
	"crypto/sha512"
	"encoding/base64"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type hashService struct{}

func NewSHA512HashService() services.HashService {
	return &hashService{}
}

func (i hashService) Hashing(value string) (string, error) {
	hash := sha512.Sum512([]byte(value))

	return base64.RawStdEncoding.EncodeToString(hash[:]), nil
}

func (i hashService) HashChecking(hash, value string) bool {
	hashing, err := i.Hashing(value)

	if err != nil {
		return false
	}

	return hashing == hash
}
