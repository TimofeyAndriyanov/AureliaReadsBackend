package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

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

type Argon2Configuration struct {
	HashRaw    []byte
	Salt       []byte
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLength  uint32
}

type hashService struct{}

func NewArgon2IdHashService() services.HashService {
	return &hashService{}
}

func genCryptoSalt() ([]byte, error) {
	salt := make([]byte, saltLen)

	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

func (i hashService) Hashing(value string) (string, error) {
	salt, saltErr := genCryptoSalt()

	if saltErr != nil {
		return "", saltErr
	}

	hash := argon2.IDKey(
		[]byte(value),
		salt,
		time,
		memory,
		threads,
		keyLen,
	)

	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		time,
		threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encodedHash, nil
}

func parseHash(encodedHash string) (*Argon2Configuration, error) {
	components := strings.Split(encodedHash, "$")
	if len(components) != 6 {
		return nil, errors.New("invalid hash format structure")
	}

	// Validate algorithm identifier
	if !strings.HasPrefix(components[1], "argon2id") {
		return nil, errors.New("unsupported algorithm variant")
	}

	// Extract version information
	var version int
	fmt.Sscanf(components[2], "v=%d", &version)

	// Parse configuration parameters
	config := &Argon2Configuration{}
	fmt.Sscanf(components[3], "m=%d,t=%d,p=%d",
		&config.MemoryCost, &config.TimeCost, &config.Threads)

	// Decode salt component
	salt, err := base64.RawStdEncoding.DecodeString(components[4])
	if err != nil {
		return nil, fmt.Errorf("salt decoding failed: %w", err)
	}
	config.Salt = salt

	// Decode hash component
	hash, err := base64.RawStdEncoding.DecodeString(components[5])
	if err != nil {
		return nil, fmt.Errorf("hash decoding failed: %w", err)
	}
	config.HashRaw = hash
	config.KeyLength = uint32(len(hash))

	return config, nil
}

func (i hashService) HashChecking(hash, value string) bool {
	config, configErr := parseHash(hash)

	if configErr != nil {
		return false
	}

	computedHash := argon2.IDKey(
		[]byte(value),
		config.Salt,
		config.TimeCost,
		config.MemoryCost,
		config.Threads,
		config.KeyLength,
	)

	return subtle.ConstantTimeCompare(config.HashRaw, computedHash) == 1
}
