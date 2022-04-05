package db

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb() {
	var err error
	Prod, _ = strconv.ParseBool(os.Getenv("PROD"))
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = 5432 // Default port
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)
	postgresconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=CST7CDT", host, user, password, dbname, port)
	Db, err = gorm.Open(postgres.Open(postgresconn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("connection not opened to database")
	}
	fmt.Println("connection opened to database")
	// Db.AutoMigrate(&public.Tenant{})
	// Db.AutoMigrate(&models.Account{}, &models.Team{}, &models.Project{}, &models.Organization{})
	fmt.Println("migrations completed")

}
