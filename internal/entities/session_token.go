package entities

import "crypto/rand"

const tokenSize = 32

type SessionToken [tokenSize]byte

func NewSessionToken() (SessionToken, error) {
	var token SessionToken

	if _, err := rand.Read(token[:]); err != nil {
		return SessionToken{}, err
	}

	return token, nil
}
