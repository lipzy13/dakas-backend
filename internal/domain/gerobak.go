package domain

import (
	"github.com/google/uuid"
)

type Gerobak struct {
	ID   uuid.UUID `gorm:"column:gerobak_id;type:uuid;primary_key;" json:"id"`
	Name string    `gorm:"column:nama_gerobak;" json:"nama_gerobak"`
}
