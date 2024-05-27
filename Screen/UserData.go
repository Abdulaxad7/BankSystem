package Screen

import (
	"Library/User"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func ShowUserMenu() {
	fmt.Print("\033[1;36m ~~~~~~~~~~~~~~~~~~~~~Bank System~~~~~~~~~~~~~~~~~~\n" +
		" \033[1;32m——————————————————————————————————————————————————\n" +
		" |  1)\t\t————————Log In————————\t\t  |\n" +
		" |  2)\t    ————————Create Account————————\t  |\n" +
		" |  0) \t\t ————————Exit———————— \t\t  |\n" +
		" \033[1;32m——————————————————————————————————————————————————\033[1;0m\n",
	)

}
func Create() {
	var err error
	reader := bufio.NewScanner(os.Stdin)
	var user User.User
	fmt.Println("\u001B[1;36m ~~~~~~~~~~~~~~~~~~~~Create Account~~~~~~~~~~~~~~~~")
	fmt.Println("\t\t\u001B[1;32mEnter your id: ")
	_, err = fmt.Fscanln(os.Stdin, &user.Id)
	if err != nil {
		log.Println(err)
		ShowUserMenu()
	}
	if user.Id <= 0 {
		fmt.Print("\033[1;31mUser id has to be positive number!!!\n\n\n\n\t \033[1;0m")
		Create()
	} else if user.Id < 999 {
		fmt.Print("\033[1;31mUser id has to contain 4 digit at least!!!\n\n\n\n\t \033[1;0m")
		Create()
	}

	fmt.Println("\t\t\u001B[1;32mEnter your name: ")
	if reader.Scan() {
		user.Name = reader.Text()
	}

	if bytes.Contains([]byte(user.Name), []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}) {
		fmt.Print("\033[1;31mUser name cannot contain integers!!!\n\n\n\n\t \033[1;0m")
		Create()
	}

	fmt.Println("\t\t\u001B[1;32mEnter your email: ")
	_, err = fmt.Fscanln(os.Stdin, &user.Email)
	if err != nil {
		log.Println(err)
		ShowUserMenu()
	}
	if !(strings.Contains(user.Email, "@") || strings.ContainsAny(user.Email, "gmail") || strings.Contains(user.Email, ".")) {
		fmt.Print("\\033[1;31mEmail has to be as [example@gmail.com] !!!\n\n\n\n\t \033[1;0m")
		Create()
	}
	fmt.Println("\t\t\u001B[1;32mEnter your password: ")
	_, err = fmt.Fscanln(os.Stdin, &user.Password)
	if err != nil {
		log.Println(err)
		ShowUserMenu()
	}
	if !(bytes.Contains([]byte(user.Password), []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}) || len(user.Password) > 5) {
		fmt.Print("\\033[1;31mYour password is too easily guessable!!!\n\n\n\n\t \033[1;0m")
		Create()
	}
	user.CreateAccount(user)
}

func Login() {
	var user User.User
	var err error
	fmt.Println("\u001B[1;36m ~~~~~~~~~~~~~~~~~~~~Log In~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("\t\t\u001B[1;32mEnter your id: ")
	_, err = fmt.Fscanln(os.Stdin, &user.Id)
	if err != nil {
		log.Println(err)
		ShowUserMenu()
	}
	fmt.Println("\t\t\u001B[1;32mEnter your password: ")
	_, err = fmt.Fscanln(os.Stdin, &user.Password)
	if err != nil {
		log.Println(err)
		ShowUserMenu()
	}
	if User.Connect(user) {
		ShowCardMenu(user.Id)
	} else {
		Start()
	}

}
