package User

import (
	_ "github.com/lib/pq"
	"log"
)

func (User User) CreateAccount(user User) {

	Open()

	insertUser(user)
}
func insertUser(user User) {

	stmt, err := db.Prepare(`INSERT INTO UserData(id, name,email,password,Created) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := stmt.Exec(user.Id, user.Name, user.Email, user.Password)
	if err2 != nil {
		log.Fatal(err2)
	}

}
