package DB

import (
	"Backend_task_advertising_site/data"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var rows *sql.Rows

func open() (*sql.DB, error) {
	db = connect()
	err = createTable()
	return db, err
}

func connect() *sql.DB {
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

func createTable() error {
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

func GettingNumberOfRecords() (numbeOfRecords int, err error) {
	db := connect()
	query := "SELECT count(*) FROM ad_table"
	rows, err = db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var numberOfRecords int
	for rows.Next() {
		if err = rows.Scan(&numberOfRecords); err != nil {
			return 0, err
		}
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}
	return numberOfRecords, nil
}

func AddNewAd(description, title string, price float64, arrayOfLinks []string) (int, error) {
	t := time.Now()
	tt := t.Format("02.01.2006")
	open()
	var id int
	ins := "INSERT INTO ad_table (announcement_text, title_ad, price, links, date) VALUES ($1, $2, $3, $4, $5) returning id"
	err = db.QueryRow(ins, description, title, price, pq.Array(arrayOfLinks), tt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ReceiveListAds(attribute, order, offset string) ([]data.Ads, error) {
	db := connect()
	query := "SELECT id, date, title_ad, links[1:1], price FROM ad_table" + " " + attribute + " " + order + " " + "limit 10" + " " + offset
	rows, err = db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]data.Ads, 0)
	var ads data.Ads
	for rows.Next() {
		if err = rows.Scan(&ads.Id, &ads.Data, &ads.Title, pq.Array(&ads.Link), &ads.Price); err != nil {
			fmt.Println(err)
			return nil, err
		}
		list = append(list, ads)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func GetOneAd(id, fields string) ([]data.Ads, error) {
	db := connect()
	var query string

	if len(fields) == 0 {
		query = fmt.Sprintf("SELECT price, title_ad, links[1:1] FROM ad_table WHERE id = %s", id)
	} else {
		query = fmt.Sprintf("SELECT announcement_text, price, title_ad, links[1:3] FROM ad_table WHERE id = %s", id)
	}

	rows, err = db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	adOne := make([]data.Ads, 0)
	var ad data.Ads
	for rows.Next() {
		if len(fields) == 0 {
			if err = rows.Scan(&ad.Price, &ad.Title, pq.Array(&ad.Link)); err != nil {
				return nil, err
			}
		} else {
			if err = rows.Scan(&ad.Description, &ad.Price, &ad.Title, pq.Array(&ad.Link)); err != nil {
				return nil, err
			}
		}
		adOne = append(adOne, ad)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return adOne, nil
}
