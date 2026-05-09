package entities

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ArticleID uuid.UUID

func (u ArticleID) String() string {
	return uuid.UUID(u).String()
}

func (u ArticleID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *ArticleID) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := uuid.Parse(s)
	if err != nil {
		return err
	}

	*u = ArticleID(parsed)

	return nil
}
