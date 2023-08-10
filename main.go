package main

import (
	"bufio"
	"firstGo/topics"
	"fmt"
	"io"
	"math"
	"os" // How to import a package without using it and still compile, not sure why i'd want this yet though.
	"strconv"

	"github.com/mactsouk/go/simpleGitHub"
)

var GlobalInt int = 123
var GlobalInferredInt = 456

func main() {
	// printingWithFmt()
	// stdout()
	// readInput()
	// commandLine()
	// mathfunc()
	// switc()
	// forloop()
	// topics.SlicesBasics()
	// topics.ReadInput()
	topics.Cla()
}

func forloop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n-----------")

	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
}

func switc() {
	args := os.Args
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}

	input := args[1]

	switch input {
	case "0":
		fmt.Println("ZERO")
	case "1":
		fmt.Println("ONE")
		fallthrough
	default:
		fmt.Println("Default branch, Value:", input)
	}

	switch {
	case input == "1":
		var one, _ = strconv.Atoi(input)
		fmt.Println(one, "hey hey hey")
	case input == "2":
		fmt.Println("dos")
	default:
		fmt.Println("default")
	}

}

func mathfunc() {
	k := math.Abs(float64(GlobalInt))
	fmt.Println(k)
}

func printingWithFmt() {
	fmt.Println("hell world")
	fmt.Println(simpleGitHub.AddTwo(1, 4))

	v1 := "123"
	v2 := "hi ther \n"
	number := 111

	fmt.Printf("%s %s %b\n", v1, v2, number)

}

func stdout() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString+"\n")
}

func readInput() {
	var stdin *os.File = os.Stdin
	defer stdin.Close()

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

}

func commandLine() {
	if len(os.Args) == 1 {
		fmt.Println("Give 1 or more floats")
		os.Exit(1)
	}

	arguments := os.Args
	min, what := strconv.ParseFloat(arguments[1], 64)
	fmt.Println(min, what)
}

/*

go build hello.go
./hello.go

or

go run hello.go

*/
