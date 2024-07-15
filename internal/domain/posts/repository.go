package posts

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, post *Post) error
	Get(ctx context.Context, id uuid.UUID) (*Post, error)
	Delete(ctx context.Context, id uuid.UUID, initiatorId uuid.UUID) error
}
