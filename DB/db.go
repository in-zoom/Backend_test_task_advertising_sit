package DB

import (
	"backend_task_advertising_site/data"
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
	db = Сonnect()
	err = createTable()
	return db, err
}

func Сonnect() *sql.DB {
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

func GettingNumberOfRecords(colum string, db *sql.DB) (numbeOfRecords int, err error) {
	query := fmt.Sprintf("SELECT %s FROM ad_table", colum)
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
	tt := t.Format("2006.01.02")
	db, err := open()
	var id int
	ins := "INSERT INTO ad_table (announcement_text, title_ad, price, links, date) VALUES ($1, $2, $3, $4, $5) returning id"
	err = db.QueryRow(ins, description, title, price, pq.Array(arrayOfLinks), tt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ReceiveListAds(attribute, order, offset string) ([]data.Ads, error) {
	db := Сonnect()
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
	db := Сonnect()
	query := fmt.Sprintf("SELECT date, price, title_ad, %s FROM ad_table WHERE id = %s", fields, id)
	rows, err = db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	colNum := len(columns)

	adOne := make([]data.Ads, 0)
	var ad data.Ads
	for rows.Next() {

		cols := make([]interface{}, colNum)
		for i := 0; i < colNum; i++ {
			cols[i] = vehicleCol(columns[i], &ad)
		}

		if err = rows.Scan(cols...); err != nil {
			return nil, err
		}
		adOne = append(adOne, ad)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return adOne, nil
}

func vehicleCol(colname string, ad *data.Ads) interface{} {
	switch colname {
	case "date":
		return &ad.Data
	case "price":
		return &ad.Price
	case "title_ad":
		return &ad.Title
	case "links":
		return pq.Array(&ad.Link)
	case "announcement_text":
		return &ad.Description
	default:
		panic("Неизвестный столбец " + colname)
	}
}
