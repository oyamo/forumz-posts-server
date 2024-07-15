package connections

import (
	"github.com/google/uuid"
)

type CreateConnectionDTO struct {
	UserId       uuid.UUID `json:"userId"`
	ConnectionTo uuid.UUID `json:"connectionTo" validate:"required"`
}
