package entities

import "crypto/rand"

const tokenSize = 32

type SessionToken []byte

func NewSessionToken() (SessionToken, error) {
	data := make([]byte, tokenSize)
	_, err := rand.Read(data)

	if err != nil {
		return nil, err
	}
	return data, nil
}
