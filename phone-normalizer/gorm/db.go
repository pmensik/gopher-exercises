package gorm

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var phoneNumbers = []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890",
	"1234567892", "(123)456-7892"}

type PhoneNumbers struct {
	gorm.Model
	PhoneNumber          string
	PhoneNumbeNormalized string
}

func SetupDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=gopher dbname=gophers sslmode=disable password=golang")
	if err != nil {
		log.Fatal(err)
	}
	seedData(db)
	return db
}

func seedData(db *gorm.DB) {
	db.DropTable(&PhoneNumbers{})
	db.AutoMigrate(&PhoneNumbers{})
	for _, num := range phoneNumbers {
		db.Create(&PhoneNumbers{PhoneNumber: num})
	}
}

func GetNumbers(db *gorm.DB) []PhoneNumbers {
	var nums []PhoneNumbers
	db.Find(&nums)
	return nums
}

func UpdateRows(db *gorm.DB, nums []string) {
	for _, num := range nums {
		db.Exec("UPDATE phone_numbers SET phone_number_normalized = ?", num)
	}
}
