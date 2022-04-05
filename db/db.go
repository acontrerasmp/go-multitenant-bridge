package db

import (
	"database/sql"

	"gorm.io/gorm"
)

var Db *gorm.DB

var SqlConnection *sql.DB

var Prod bool
