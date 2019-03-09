package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	// Update record
	// Method 1
	// UPDATE `addressBook` SET last_name='Google' WHERE user_id=3;
	db.Model(&Person{}).Where("user_id = ?", 3).Update("last_name", "Google")
	// Method 2
	// UPDATE `addressBook` SET last_name='Yahoo' WHERE user_id=3;
	db.Model(&Person{}).Where(Person{UserID: 3}).Update(Person{LastName: "Yahoo"})
	// Method 3
	// UPDATE `addressBook` SET last_name='Microsof' WHERE user_id=3;
	person := []Person{}
	db.Model(&person).Where("user_id = ?", 3).Update("last_name", "Microsoft")
}

type Person struct {
	UserID    int    `gorm:"primary_key";"AUTO_INCREMENT"`
	FirstName string `gorm:"size:15"`
	LastName  string `gorm:"size:15"`
	Address   string `gorm:"type:varchar(100)"`
	Phone     string `gorm:"size:20"`
}

// Set custom table name
func (Person) TableName() string {
	return "addressBook"
}
