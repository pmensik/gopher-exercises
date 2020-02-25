package main

import (
	"regexp"

	"github.com/pmensik/gopher-exercises/phone-normalizer/gorm"
	"github.com/pmensik/gopher-exercises/phone-normalizer/sqlx"
)

var reg = regexp.MustCompile(`[^0-9]`)

func main() {
	// runSqlx()
	runGorm()
}

func runGorm() {
	db := gorm.SetupDB()
	nums := gorm.GetNumbers(db)
	var numsSlice []string
	for _, num := range nums {
		numsSlice = append(numsSlice, num.PhoneNumber)
	}
	defer db.Close()
}

func runSqlx() {
	db := sqlx.SetupDB()
	defer db.Close()
	nums := sqlx.GetNumbers(db)
	sqlx.UpdateRows(db, Normalize(nums))
}

func Normalize(nums []string) []string {
	var normalized []string
	for _, num := range nums {
		normalized = append(normalized, reg.ReplaceAllString(num, ""))
	}
	return normalized
}
