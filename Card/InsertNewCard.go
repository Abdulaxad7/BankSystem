package Card

import (
	"log"
)

func (Card Card) InsertCard(card Card) {

	Open()
	insertCard(card)
}

func insertCard(card Card) {
	stmt, err := db.Prepare(`INSERT INTO User_Card_Data(id, card_number,card_holder,card_thru_date,password,card_balance,created) VALUES($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := stmt.Exec(card.Id, card.CardNumber, card.CardHolder, card.CardThruDate, card.CardPassword, card.CardBalance)
	if err2 != nil {
		log.Fatal(err2)
	}

}
