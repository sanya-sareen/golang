package database

import (
	"fmt"
	 "os"
	 "log"
	"gorm.io/driver/postgres"
	"gorm.io/gorum"
)

type Dbinstance struct{
	Db *gorum.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host"=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeOne=Asia/Shanghai",
		 os.Getenv("DB_USER"),
		 os.Getenv("DB_PASSWORD"),
		 os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gormConfig{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil{
		log.Fatal("failed to connect",err)
		os.Exit(2)
	}

	log.Println("connected")
	db.logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{
		Db:db,
	}
}