package billing

import (
	"time"
)

type Invoice struct {
	ID          string    `json:"id"`
	Client      string    `json:"client"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
