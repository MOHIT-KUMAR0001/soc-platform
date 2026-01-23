package schema

import (
	"time"

	"github.com/google/uuid"
)

type ApiSchema struct {
	ID          uuid.UUID `json:"id"`
	WorkerGroup string    `json:"worker_name"`
	ApiKeyHash  string    `json:"api_key_hash"`
	IsDisabled  bool      `json:"is_disabled"`
	Expiry      time.Time `json:"expiry"`
	CreateAt    time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
