package entities

import (
	"encoding/json"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

func (u UserID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *UserID) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := uuid.Parse(s)
	if err != nil {
		return err
	}

	*u = UserID(parsed)

	return nil
}
