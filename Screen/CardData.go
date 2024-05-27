package Screen

import (
	"Library/Card"
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func ShowCardMenu(userId int) {
	var choice int
	showCards(userId)
	anim()
	fmt.Println("Enter your choice: ")
	_, err := fmt.Fscanln(os.Stdin, &choice)
	if err != nil {
		log.Fatal(err)
	}
	switch choice {
	case 1:
		cardBalance(userId)

	case 2:
		recentTransactions(userId)
	case 3:
		cardFullInfo(userId)
	case 4:
		payBills(userId)
	case 5:
		createCard(userId)
	case 6:
		removeCard(userId)
	}

}

func cardBalance(userId int) {

	user := Card.UserCards(userId)
	var err error
	var cardNumber string
	var balance interface{}
	fmt.Println("Enter card number: ")
	_, err = fmt.Fscanln(os.Stdin, &cardNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range user {
		if v["CardNumber"] == cardNumber {
			balance = v["CardBalance"]
		}
	}
	if balance != nil {
		fmt.Printf("The card balance is: %s\n", balance)
		time.Sleep(3 * time.Second)
	} else {
		fmt.Printf("\033[1;31mThe card with number [%q] did not found!!!\n", cardNumber)
		time.Sleep(3 * time.Second)
	}

	ShowCardMenu(userId)
}
func recentTransactions(userId int) {
	var err error
	user := Card.UserCards(userId)
	card := Card.UserTransaction(userId)

	var cardNumber string

	fmt.Println("Enter card number: ")
	_, err = fmt.Fscanln(os.Stdin, &cardNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range user {
		if v["CardNumber"] == cardNumber {
			for _, u := range card {
				fmt.Printf("The transaction was on [%q] and the amount was [%q] \n", u["transaction_time"], u["transaction"])
			}
		}
	}
	ShowCardMenu(userId)

}

func cardFullInfo(userId int) {

}
func payBills(userId int) {

}
func removeCard(userId int) {
	//var err error
	//fmt.Println("Enter card number to delete: ")

}

func anim() {
	fmt.Print(" \033[1;32m——————————————————————————————————————————————————\n" +
		" |  1)\t\t————————Card Balance————————\t  |\n" +
		" |  2)\t    ————————Recent Transactions————————\t  |\n" +
		" |  3)\t       ————————Card Full Info————————\t  |\n" +
		" |  4)\t\t  ————————Pay Bills————————\t  |\n" +
		" |  5) \t\t————————Add New Card———————— \t  |\n" +
		" |  6) \t\t ————————Remove Card——————— \t  |\n" +
		" \033[1;32m——————————————————————————————————————————————————\033[1;0m\n",
	)
}

func createCard(userId int) {
	reader := bufio.NewScanner(os.Stdin)
	var err error
	var card Card.Card
	fmt.Println("Enter Card Number: ")
	//_, err = fmt.Fscanln(os.Stdin, &card.CardNumber)
	if reader.Scan() {
		card.CardNumber, err = strconv.Atoi(reader.Text())
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Enter Card Holder Name: ")

	if reader.Scan() {
		card.CardHolder = reader.Text()
	}
	fmt.Println("Enter Card Thru Date: ")
	_, err = fmt.Fscanln(os.Stdin, &card.CardThruDate)

	fmt.Println("Generate Card Password: ")
	_, err = fmt.Fscanln(os.Stdin, &card.CardPassword)

	if err != nil {
		log.Fatal(err)
	}
	card = Card.Card{
		Id:           userId,
		CardNumber:   card.CardNumber,
		CardHolder:   card.CardHolder,
		CardThruDate: card.CardThruDate,
		CardPassword: card.CardPassword,
		CardBalance:  rand.Intn(10000),
		CardExpenses: 0,
	}
	card.InsertCard(card)
}

func showCards(userId int) {
	user := Card.UserCards(userId)
	fmt.Printf("\033[1;36m\t\tHi %s \n", user[1]["CardHolder"])
	fmt.Print("\033[1;36m ~~~~~~~~~~~~~~~~~~~~~Your cards~~~~~~~~~~~~~~~~~~\n\n")
	for _, card := range user {
		fmt.Printf(" \033[1;32m——————————————————————————————————————————————————\n")
		fmt.Printf(" |    \t\t\t\t\t\t  |\n")
		fmt.Printf(" |    \t \033[1;42m      \033[1;0;32m\t\t\t\t\t  |\n")
		fmt.Printf(" |    \t \033[1;42m      \033[1;0;32m\t\t\t\t\t  |\n")

		fmt.Printf(" |    \t\t\t%q\t  |\n", card["CardHolder"])
		fmt.Printf(" |    \t\t\t\t%q  \t  |\n", card["CardThruDate"])
		fmt.Printf(" |       %q	\t\t  |\n", card["CardNumber"])
		fmt.Printf(" |    \t\t\t\t\t\t  |\n")
		fmt.Printf(" \033[1;32m——————————————————————————————————————————————————\n")

	}
}
