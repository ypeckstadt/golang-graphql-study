package satori

import (
	"github.com/gofrs/uuid"
)

// New creates new satori based UUD  service
func New() SatoryUUID {
	return SatoryUUID{
	}
}

type SatoryUUID struct {
}

// Generate generates a new UUID
func (s SatoryUUID) Generate() string {
	var err error
	gen, err := uuid.NewV4()
	return uuid.Must(gen, err).String()
}


