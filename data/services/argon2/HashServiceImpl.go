package argon2

import (
	"crypto/rand"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
	"golang.org/x/crypto/argon2"
)

const (
	time    uint32 = 3
	memory  uint32 = 64 * 1024
	threads uint8  = 4
	saltLen uint32 = 16
	keyLen  uint32 = 32
)

type hashService struct{}

func NewArgon2IdHashService() services.HashService {
	return &hashService{}
}

func (i hashService) Hashing(value string) []byte {
	salt := make([]byte, saltLen)

	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	return argon2.IDKey([]byte(value), salt, time, memory, threads, keyLen)
}

func (i hashService) HashChecking(hash []byte, value string) bool {
	panic("")
}
