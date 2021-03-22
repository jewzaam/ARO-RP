package uuid

import (
	uuid "github.com/gofrs/uuid"
)

// This is a facade to aid in conversion from github.com/satori/go.uuid to github.com/gofrs/uuid
// Changing due to https://www.whitesourcesoftware.com/vulnerability-database/WS-2018-0594

// helper so we can in-line uuid generation with gofrs, drops the error that would be returned
func NewV4() uuid.UUID {
	u, _ := uuid.NewV4()
	return u
}

// simple passthrough
func Must(u uuid.UUID, err error) uuid.UUID {
	return uuid.Must(u, err)
}

// simple passthroughs
func FromString(input string) (uuid.UUID, error) {
	return uuid.FromString(input)
}
