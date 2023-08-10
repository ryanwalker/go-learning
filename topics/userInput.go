package topics

import (
	"fmt"
)

func ReadInput() {
	fmt.Print("What is your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is:", name)
}
