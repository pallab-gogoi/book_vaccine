// initialize connection b/w database and this code
package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func setUp() {
	host := "localhost"
	port := "5432"
	dbName := "vaccine"
	userName := "postgres"
	passWord := "postgres"
	arg := "host=" + host + " port=" + port + " user=" + userName + " dbname=" + dbName + " sslmode=disable password=" + passWord
	db, err := gorm.Open("postgres", arg)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(model.details{})
	DB = db
}
