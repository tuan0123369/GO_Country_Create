package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const ( //thay doi cho nay de them bang vao database
	host         = "localhost"
	port         = 5432
	user         = "admin"
	password     = "1234"
	databaseName = "postgres"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func connect(host, user, password, databaseName string, port int) *sql.DB {
	msql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)
	db, err := sql.Open(databaseName, msql)
	checkError(err)
	return db
}

func createCountryTable(db *sql.DB) {
	_, e := db.Exec("CREATE TABLE Country (Country_code	varchar(10) PRIMARY KEY, Country_Name varchar(50))")
	checkError(e)
}

func addTable(db *sql.DB, value1 string, value2 string) {
	_, e := db.Exec("INSERT INTO country(Country_code, Country_Name) values($1,$2)", value1, value2)
	checkError(e)
}

func readCsv(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	record, _ := r.Read()
	return record
}

func splitString(str string, key string) (string, string) {
	s := strings.Split(str, key)
	return s[0], s[1]
}

func main() {
	record := readCsv("country.csv")
	db := connect(host, user, password, databaseName, port)
	createCountryTable(db)
	for value := range record {
		s1, s2 := splitString(record[value], "|")
		addTable(db, s1, s2)
	}
	fmt.Println("Create Table Success")
	defer db.Close()
}
