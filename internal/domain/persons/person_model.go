package persons

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	Id              uuid.UUID `json:"id"`
	FirstName       string    `json:"firstName"`
	EmailAddress    string    `json:"emailAddress"`
	Username        string    `json:"username"`
	Status          string    `json:"status"`
	DatetimeCreated time.Time `json:"datetimeCreated"`
	LastModified    time.Time `json:"lastModified"`
}
