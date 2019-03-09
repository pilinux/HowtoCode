package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func main() {
	dbType := "mysql"
	dbUser := "userName"
	dbName := "databaseName"
	dbPass := "userPass"
	dbInfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbName)

	db, err := gorm.Open(dbType, dbInfo)
	defer db.Close()

	if err != nil {
		fmt.Println("DB connection error: ", err)
	}
	if err == nil {
		fmt.Println("Connection to database is successful!")
	}

	// Auto migration
	/*
		- Automatically migrate schema
		- Only create tables with missing columns and missing indexes
		- Will not change/delete any existing columns and their types
		- db.Set() --> add table suffix during auto migration
	*/

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Person{})
}

type Person struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
