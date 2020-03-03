package DB

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var rows *sql.Rows

func Open() (*sql.DB, error) {
	db = Connect()
	err = createTable()
	return db, err
}

func Connect() *sql.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func createTable() (err error) {
    ins := "CREATE TABLE IF NOT EXISTS ad_table (id SERIAL, date DATE, price MONEY, announcement_text VARCHAR, title_ad VARCHAR, links TEXT[])"
	_, err = db.Exec(ins)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func AddNewAd(description, title string, price float64, arrayOfLinks []string) (err error) {
	ins := "INSERT INTO ad_table (announcement_text, title_ad, price, links) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(ins, description, title, price, pq.Array(arrayOfLinks))
	if err != nil {
		return err
	}
	return nil
}
