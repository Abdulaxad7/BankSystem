package Insert

import (
	"Library/User"
	_ "github.com/lib/pq"
	"log"
)

func (User User.User) CreateAccount(user User.User) {

	User.Open()

	insertUser(user)
}
func insertUser(user User.User) {

	stmt, err := User.db.Prepare(`INSERT INTO UserData(id, name,email,password,Created) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := stmt.Exec(user.Id, user.Name, user.Email, user.Password)
	if err2 != nil {
		log.Fatal(err2)
	}

}
