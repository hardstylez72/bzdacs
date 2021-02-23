package util

import "time"

type Times struct {
	CreatedAt time.Time    `json:"createdAt" db:"created_at" swaggertype:"string"`
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at" swaggertype:"string"`
	DeletedAt JsonNullTime `json:"deletedAt" db:"deleted_at" swaggertype:"string" extensions:"x-nullable"`
}
