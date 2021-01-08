package user

import (
	"github.com/hardstylez72/bzdacs/pkg/util"
	"time"
)

type User struct {
	Id         int    `json:"id" db:"id"`
	ExternalId string `json:"externalId" db:"external_id"`

	CreatedAt time.Time         `json:"createdAt" db:"created_at"`
	UpdatedAt util.JsonNullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt util.JsonNullTime `json:"deletedAt" db:"deleted_at"`
}
