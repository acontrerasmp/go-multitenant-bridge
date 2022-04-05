package shared

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID uuid.UUID `gorm:"primaryKey:type:uuid;" json:"id"`
	gorm.Model
	Name        string
	Description string
}
