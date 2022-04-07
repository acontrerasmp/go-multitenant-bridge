package public

import (
	"gorm.io/gorm"
)

type Tenant struct {
	// ID uuid.UUID `gorm:"primaryKey;type:uuid;unique" json:"id"`
	ID int `gorm:"primarykey"`
	gorm.Model
	Name   string `gorm:"unique;not null" json:"name"`
	Domain string `gorm:"unique;not null" json:"domain"`
}
