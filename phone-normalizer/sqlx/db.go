package sqlx

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE phone_numbers (
	id SERIAL,
	phone_number VARCHAR(20),
	phone_number_normalized VARCHAR(10)
)
`
var phoneNumbers = []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890",
	"1234567892", "(123)456-7892"}

func SetupDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=gopher dbname=gophers sslmode=disable password=golang")
	if err != nil {
		log.Fatal(err)
	}
	seedDB(db)
	return db
}

func seedDB(db *sqlx.DB) {
	db.MustExec("DROP TABLE IF EXISTS phone_numbers")
	db.MustExec(schema)
	tx := db.MustBegin()
	for _, num := range phoneNumbers {
		tx.MustExec("INSERT INTO phone_numbers (phone_number) VALUES ($1)", num)
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

func GetNumbers(db *sqlx.DB) []string {
	var nums []string
	err := db.Select(&nums, "SELECT phone_number FROM phone_numbers")
	if err != nil {
		fmt.Println(err)
	}
	return nums
}

func UpdateRows(db *sqlx.DB, nums []string) {
	tx := db.MustBegin()
	for _, num := range nums {
		tx.MustExec("UPDATE phone_numbers SET phone_number_normalized = $1", num)
	}
	err := tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
}
