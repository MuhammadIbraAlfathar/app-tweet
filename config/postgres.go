package config

import (
	"github.com/MuhammadIbraAlfathar/app-tweet/internal/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgres() (*gorm.DB, error) {

	host := Env.PostgresHost
	user := Env.PostgresUser
	password := Env.PostgresPassword
	dbName := Env.PostgresDbName
	port := Env.PostgresPort

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error to connect database")
	}

	execEnumType(db)

	err = migratePostgresTable(db)
	if err != nil {
		log.Println("Migrate database failed:", err)
	}

	log.Println("connected")

	return db, nil

}

func migratePostgresTable(db *gorm.DB) error {

	models := []interface{}{
		&schema.User{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}

func execEnumType(db *gorm.DB) {

	var exists bool
	db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role')").Scan(&exists)

	if !exists {
		err := db.Exec("CREATE TYPE users_role AS ENUM ('admin', 'user', 'guest')").Error
		if err != nil {
			log.Println("ERROR CREATE ENUM USERS_ROLE :	", err)
		}
	} else {
		log.Println("Enum type 'user_role' already exists")
	}

	db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender')").Scan(&exists)
	if !exists {
		err := db.Exec("CREATE TYPE gender AS ENUM ('male', 'female')").Error
		if err != nil {
			log.Println("ERROR CREATE ENUM GENDER :	", err)
		}
	} else {
		log.Println("Enum type 'gender' already exists")
	}

	db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tags')").Scan(&exists)
	if !exists {
		err := db.Exec("CREATE TYPE tags AS ENUM ('Technology', 'Programming', 'Lifestyle', 'Business')").Error
		if err != nil {
			log.Println("ERROR CREATE ENUM TAGS :	", err)
		}
	} else {
		log.Println("Enum type 'tags' already exists")
	}

}
