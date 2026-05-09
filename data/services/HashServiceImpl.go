package services

import (
	"AureliaReadsBackend/domain/services"
	"crypto/sha512"
	"encoding/base64"
)

type hashServiceImpl struct{}

func NewHashService() services.HashService {
	return &hashServiceImpl{}
}

func (i hashServiceImpl) Hashing(value string) string {
	binHash := sha512.Sum512([]byte(value))
	stringHash := base64.StdEncoding.EncodeToString(binHash[:])

	return stringHash
}

func (i hashServiceImpl) HashChecking(hash string, value string) bool {
	return i.Hashing(value) == hash
}
