package posts

import "github.com/google/uuid"

type CreatePostDTO struct {
	PersonId uuid.UUID `json:"personId"`
	Content  string    `json:"content" validate:"required,min=5,max=255"`
}
