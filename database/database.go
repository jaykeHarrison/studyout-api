package database

import (
	"log"
	"os"

	"github.com/jaykeHarrison/studyout-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//GORM DB instance
type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	/*
		Connects to local psql database

		you will need to create a database with name studyout
		# CREATE DATABASE studyout

		check the port that psql uses, this is default 5432
		# SELECT * FROM pg_settings WHERE name = 'port'
		if this is not 5432, change the port for dsn
	*/
	dsn := "host=postgres://whukezqqivadin:f1810aa7bc247490000ba184427d7ffcb25e372e18cce181546d542d22b0e9c5@ec2-3-248-121-12.eu-west-1.compute.amazonaws.com:5432/d27jvfvqp86ufd dbname=d27jvfvqp86ufd port=5432 user=whukezqqivadin password=f1810aa7bc247490000ba184427d7ffcb25e372e18cce181546d542d22b0e9c5"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//if there is an error connecting to db, this will throw the error
	if err != nil {
		//log error message
		log.Fatal("Failed to connect to the database! \n", err.Error())

		//ends program with cleanup
		os.Exit(2)
	}

	//if no error
	log.Println("Connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	//Add migrations - these are tables using structs
	db.AutoMigrate(&models.Users{}, &models.UserPreference{}, &models.Location{}, &models.LocationFeature{}, &models.Review{}, &models.Bookmark{})
	//instance of DbInstance struct; db is a gorm.db
	Database = DbInstance{Db: db}
}
