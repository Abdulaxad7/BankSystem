package Card

import (
	_ "github.com/lib/pq"

	"database/sql"
	"log"
)

type Card struct {
	Id           int    `json:"id"`
	CardNumber   int    `json:"card_number"`
	CardHolder   string `json:"card_holder"`
	CardThruDate string `json:"card_thru_date"`
	CardPassword string `json:"card_password"`
	CardBalance  int    `json:"card_balance"`
}

var db *sql.DB

func Open() {

	const connStr = "postgres://postgres:secret@localhost:5435/bank?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)

	query := `create table if not exists User_Card_Data(
		Id  int,
		Card_Number text primary key not null,
    	Card_Holder varchar(25) not null,
		Card_Thru_Date varchar(100) not null,
    	Password varchar(100) not null,
    	Card_Balance int,
		Created timestamp default now()
		)`
	if err != nil {
		log.Fatal(err)
	}
	defer db.Exec(query)
}
