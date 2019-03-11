/*
** - User knows many languages
** - Same language is known by many users
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
	//	db.DropTableIfExists(&UserLanguage{}, &User{}, &Language{})

	//	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Language{}, &UserLanguage{})

	//	Manually set foreign key
	//	db.Model(&UserLanguage{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	//	db.Model(&UserLanguage{}).AddForeignKey("language_id", "languages(id)", "CASCADE", "CASCADE")

	// Create record
	lang1 := []Language{
		{Name: "EN"},
		{Name: "FR"},
		{Name: "DE"},
		{Name: "AR"},
	}
	lang2 := []Language{
		{Name: "EN"},
		{Name: "FR"},
		{Name: "DE"},
	}

	person1 := User{
		Name:      "Robin",
		Languages: lang1,
	}
	person2 := User{
		Name:      "Nami",
		Languages: lang1,
	}
	person3 := User{
		Name:      "Zoro",
		Languages: lang2,
	}
	db.Create(&person1)
	db.Create(&person2)
	db.Create(&person3)

	person4 := []User{
		{Name: "Luffy"},
		{Name: "Sanji"},
	}

	lang3 := Language{
		Name:  "JP",
		Users: person4,
	}
	db.Create(&lang3)
}

type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages";"foreignkey:UserID";"association_foreignkey:ID"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_languages";"foreignkey:LanguageID";"association_foreignkey:ID"`
}

type UserLanguage struct {
	UserID     uint
	LanguageID uint
}
