package posts

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id              uuid.UUID `json:"id"`
	PersonId        uuid.UUID `json:"personId"`
	Content         string    `json:"content"`
	DatetimeCreated time.Time `json:"datetimeCreated"`
	LastModified    time.Time `json:"lastModified"`
}
