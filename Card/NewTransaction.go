package Card

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func init() {
	OpenTransaction()
}
func (transaction Transaction) InsertNewTransaction(card Transaction, receiver int) {
	if currentTransaction(card.TransactionAmount, card.TransactionCard, card.Id) {
		insertTransaction(transaction)
		fmt.Printf("The transaction was successful with amount [%d] to the receiver [%d]\n", card.TransactionAmount, receiver)
		time.Sleep(5 * time.Second)
	}
}
func insertTransaction(transaction Transaction) {
	stmt, err := db.Prepare(`INSERT INTO User_Card_Transactions(id,card_number,transaction,transaction_time) VALUES($1,$2,$3,CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(transaction.Id, transaction.TransactionCard, transaction.TransactionAmount)
	if err != nil {
		log.Fatal(err)
	}

}

func currentTransaction(amount int, cardNumber int, userId int) bool {
	cardnumber := strconv.Itoa(cardNumber)
	var err error
	user := UserCards(userId)
	const query = `

UPDATE  User_Card_Data
	SET card_balance = card_balance - $1 
	WHERE  card_balance>=$1 and  User_Card_Data.card_number=$2 `
	stmt, err := db.Prepare(query)
	var check bool
	for _, v := range user {

		if v["CardNumber"] == cardnumber {

			check = true
		}
	}

	if check {
		_, err = stmt.Exec(amount, cardNumber)
	} else {
		fmt.Printf("There is no card with number [%d] in your wallet", cardNumber)
		time.Sleep(5 * time.Second)
		return false
	}
	if err != nil {
		log.Fatal(err)
	}
	return true
}
func Remove(cardNumber int) {
	var err error
	const query = `delete from user_card_data where card_number=$1`
	stmt, err := db.Prepare(query)
	_, err = stmt.Exec(cardNumber)
	if err != nil {
		log.Fatal(err)
	}
	removeTransaction(cardNumber)
}
func removeTransaction(cardNumber int) {
	const query = `delete from user_card_transactions where card_number=$1`
	stmt, err := db.Prepare(query)
	_, err = stmt.Exec(cardNumber)
	if err != nil {
		log.Fatal(err)
	}
}
