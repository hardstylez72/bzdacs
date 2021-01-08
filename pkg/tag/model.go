package tag

import (
	"github.com/hardstylez72/bblog/ad/pkg/util"
	"time"
)

type Tag struct {
	Id        int               `json:"id" db:"id"`
	Name      string            `json:"name" db:"name"`
	CreatedAt time.Time         `json:"createdAt" db:"created_at"`
	UpdatedAt util.JsonNullTime `json:"updatedAt" db:"updated_at"`
	DeletedAt util.JsonNullTime `json:"deletedAt" db:"deleted_at"`
}
