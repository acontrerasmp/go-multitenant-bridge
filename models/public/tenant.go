package public

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;unique" json:"id"`
	gorm.Model
	Name   string `gorm:"unique;not null" json:"name"`
	Domain string `gorm:"unique;not null" json:"domain"`
}
