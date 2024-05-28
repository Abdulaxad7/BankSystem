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
		payBills(userId)
	case 4:
		createCard(userId)
	case 5:
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

	var cardNumber string

	fmt.Println("Enter card number: ")
	_, err = fmt.Fscanln(os.Stdin, &cardNumber)
	c_n, _ := strconv.Atoi(cardNumber)
	card := Card.UserTransaction(userId, c_n)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range user {
		fmt.Println(v["CardNumber"], "   ", cardNumber)
		if v["CardNumber"] == cardNumber {

			for _, u := range card {
				fmt.Printf("The transaction was on [%s] and the amount was [%s] \n", u["Time"], u["Amount"])
			}
		}
	}
	time.Sleep(5 * time.Second)
	ShowCardMenu(userId)

}

func payBills(userId int) {
	var receiver int
	var err error
	transaction := Card.Transaction{}
	fmt.Println("Enter transaction card number: ")
	_, err = fmt.Fscanln(os.Stdin, &transaction.TransactionCard)
	fmt.Println("Enter transaction receiver card number: ")
	_, err = fmt.Fscanln(os.Stdin, &receiver)
	if size(receiver) != 16 {
		fmt.Println("\033[1;31mCard Number must contain exactly 16 digits!!!\u001B[1;0m")
		payBills(userId)
	}
	fmt.Println("Enter transaction amount: ")
	_, err = fmt.Fscanln(os.Stdin, &transaction.TransactionAmount)
	if err != nil {
		log.Fatal(err)
	}
	transaction = Card.Transaction{
		Id:                userId,
		TransactionCard:   transaction.TransactionCard,
		TransactionAmount: transaction.TransactionAmount,
	}
	transaction.InsertNewTransaction(transaction, receiver)

	ShowCardMenu(userId)
}
func removeCard(userId int) {
	//var err error
	//fmt.Println("Enter card number to delete: ")

}

func anim() {
	fmt.Print(" \033[1;32m——————————————————————————————————————————————————\n" +
		" |  1)\t\t————————Card Balance————————\t  |\n" +
		" |  2)\t    ————————Recent Transactions————————\t  |\n" +
		" |  3)\t\t  ————————Pay Bills————————\t  |\n" +
		" |  4) \t\t————————Add New Card———————— \t  |\n" +
		" |  5) \t\t ————————Remove Card——————— \t  |\n" +
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
	}
	fmt.Println("\033[1;31mInserting new card...!!!\u001B[1;0m")
	time.Sleep(5 * time.Second)
	card.InsertCard(card)
	ShowCardMenu(userId)
}

func showCards(userId int) {
	user := Card.UserCards(userId)
	if user != nil {

		fmt.Printf("\033[1;36m\t\tHi %s \n", user[0]["CardHolder"])
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
}

func size(d int) int {
	count := 0
	for d != 0 {
		d /= 10
		count++
	}
	return count
}
