package Card

import (
	"database/sql"
	"log"
)

func OpenTransaction() {

	const connStr = "postgres://postgres:secret@localhost:5435/bank?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	query := `create table if not exists User_Card_Transactions(
		id  int not null,
		card_number varchar(100) not null,
		transaction varchar(100) not null,
    	transaction_time timestamp default now()
		)`
	if err != nil {
		log.Fatal(err)
	}
	defer db.Exec(query)
}
