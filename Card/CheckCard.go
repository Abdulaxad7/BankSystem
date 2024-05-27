package Card

import (
	"Library/User"
	"log"
)

func init() {
	Open()
	User.Open()
	OpenTransaction()
}
func UserCards(id int) []map[string]interface{} {
	return selectCards(id)
}

func UserTransaction(id int) []map[string]interface{} {
	return userTran(id)
}

func userTran(id int) []map[string]interface{} {
	const query = `SELECT user_card_transactions.transaction user_card_transactions.transaction_time
	from user_card_data
	INNER JOIN user_card_transactions ON user_card_transactions.id=user_card_data.id
	WHERE user_card_data.id=$1`
	stmt, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	var s [1000]string
	i := 0
	for stmt.Next() {
		err = stmt.Scan(&s[i], &s[i+1])
		i += 2
	}
	var transactions []map[string]interface{}
	insert := make(map[string]interface{})
	for points := 0; points < cap(s); points += 2 {
		if s[points] == "" {
			break
		}
		insert = map[string]interface{}{
			"Id":   s[points],
			"Time": s[points+1],
		}
		transactions = append(transactions, insert)
	}
	return transactions
}

func selectCards(id int) []map[string]interface{} {

	const query = `SELECT user_card_data.*
			FROM userdata
			INNER JOIN user_card_data ON userdata.id = user_card_data.id
			WHERE userdata.id = $1;`

	stmt, err := db.Query(query, id)
	if err != nil {
		recover()
		log.Fatal(err)
	}
	i := 0
	var s [1000]string
	for stmt.Next() {
		err := stmt.Scan(
			&s[i], &s[i+1], &s[i+2], &s[i+3], &s[i+4], &s[i+5], &s[i+6],
		)
		if err != nil {
			log.Fatal(err)
		}
		i += 7

	}
	var results []map[string]interface{}

	insert := make(map[string]interface{})
	for points := 0; points < cap(s); points += 7 {
		if s[points] == "" {
			break
		}
		insert = map[string]interface{}{
			"Id":           s[points],
			"CardNumber":   s[points+1],
			"CardHolder":   s[points+2],
			"CardThruDate": s[points+3],
			"CardPassword": s[points+4],
			"CardBalance":  s[points+5],
			"CardExpenses": s[points+6],
		}
		results = append(results, insert)
	}
	return results
}
