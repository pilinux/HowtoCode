/*
** - User has many emails
** - UserID is the foreign key [it holds unique value from users(id)]
 */

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	// Auto migration block needs to be executed only when it is required

	//	foreignkey is defined in &Email{}, hence it needs to be deleted first
	//	db.DropTableIfExists(&Email{}, &User{})

	//	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Email{})

	//	Manually set foreign key
	//	db.Model(&Email{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// Create record
	email1 := []Email{
		{Email: "Joe1@gmail.com"},
		{Email: "Joe2@gmail.com"},
		{Email: "Joe3@gmail.com"},
		{Email: "Joe4@gmail.com"},
	}
	user1 := User{Name: "Joe1", Emails: email1}
	db.Create(&user1)
}

type User struct {
	gorm.Model
	Name   string
	Emails []Email `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint
}
