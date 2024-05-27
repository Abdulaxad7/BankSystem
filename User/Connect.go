package User

import (
	"database/sql"
	"log"
)

func Connect(user User) bool {
	Open()

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	if selecting(user, db).Next() {
		return true
	} else {
		return false
	}
	return false
}
func selecting(user User, db *sql.DB) *sql.Rows {
	const query = `select *from UserData where (id,password)=($1,$2)`
	row, err := db.Query(query, user.Id, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return row

}
