package Screen

import (
	"fmt"
	"log"
	"os"
)

func Start() {
	var o int
	ShowUserMenu()
	fmt.Println("Enter choice: ")
	_, err := fmt.Scanln(&o)
	if err != nil {
		log.Fatalln(err)

	}
	switch o {
	case 1:
		Login()
	case 2:
		Create()
	case 0:
		os.Exit(0)
	default:
		Start()
	}

}
