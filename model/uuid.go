package model

import (
	"fmt"

	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

// UUID aliases & extends the support of the uuid type used for model objects
type UUID [16]byte

// NewUUID creates a new UUID
func NewUUID() (*UUID, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	uuidOut := new(UUID)
	*uuidOut = [16]byte(*uuid)
	return uuidOut, nil
}

func (u *UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// MarshalJSON overrides the JSON representation of the object
func (u *UUID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", u.String())), nil
}

// UnmarshalJSON overrides the JSON representation of the object
func (u *UUID) UnmarshalJSON(byt []byte) error {
	str := string(byt)
	str = strings.Trim(str, "\"")
	fmt.Println("candidate:", str)
	uuid, err := uuid.ParseHex(str)
	if err != nil {
		return err
	}
	*u = [16]byte(*uuid)
	return nil
}
