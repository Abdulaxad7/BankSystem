package User

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db *sql.DB

func Open() {

	const connStr = "postgres://postgres:secret@localhost:5435/bank?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	query := `create table if not exists UserData(
		id serial primary key,
		name varchar(100) not null,
    	email varchar(25) not null,
		password varchar(100) not null,
		created timestamp default now()
		)`
	if err != nil {
		log.Fatal(err)
	}
	defer db.Exec(query)
}
