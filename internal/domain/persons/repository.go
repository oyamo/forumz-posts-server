package persons

import (
	"context"
	"github.com/google/uuid"
)

type PersonRepository interface {
	Upsert(context.Context, *Person) error
	Find(context.Context, uuid.UUID) (*Person, error)
	Exists(context.Context, uuid.UUID) (bool, error)
}
