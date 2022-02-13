package codepractice

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func ReadArgsAsArray() {
	// the below code works only from main.go as args will be passed to main func from command line
	// Check that our program takes exactly three arguments:
	if len(os.Args) != 4 {
		log.Fatal("Error! Expected 3 arguments only!")
	}

	fmt.Printf("Contents of the os.Args slice = %v\n", os.Args)
	fmt.Printf("FirstName of our program --> os.Args[0] = %[1]s | type: %[1]T\n", os.Args[0])
	fmt.Printf("First cmd-line argument --> os.Args[1] = %[1]s | type: %[1]T\n", os.Args[1])
	fmt.Printf("Second cmd-line argument --> os.Args[2] = %[1]s | type: %[1]T\n", os.Args[2])
	fmt.Printf("Third cmd-line argument --> os.Args[3] = %[1]s | type: %[1]T\n", os.Args[3])
}

// NamedArgs pass the args to main function as following ./main --name Keanu --age 57 --spacing=true
//Note that for Go programs, when specifying a bool flag argument, it is mandatory to use the = character.
func NamedArgs() {
	// Declare string and int flags with default values and help messages:
	name := flag.String("name", "User", "Enter your name")
	age := flag.Int("age", 1, "Enter your age")

	// Another way to declare a flag - bind a command-line flag to an existing variable:
	var spacing bool
	flag.BoolVar(&spacing, "spacing", true, "Insert a new line between each message")

	// After declaring all the flags, enable command-line flag parsing:
	flag.Parse()

	fmt.Printf("Hello, %s! ", *name)
	if spacing {
		fmt.Println()
	}
	fmt.Printf("You are %d years old.", *age)
}

// SubCommands
//Go program that contains two subcommands — reverse and repeat:
//to Run: ./main reverse --name Kanye or ./main repeat --name Wayne --count 3
func SubCommands() {
	reverse := flag.NewFlagSet("reverse", flag.ExitOnError)
	reverseName := reverse.String("name", "User", "Enter the name to reverse")

	repeat := flag.NewFlagSet("repeat", flag.ExitOnError)
	repeatName := repeat.String("name", "User", "Enter the name to be repeated")
	repeatCount := repeat.Int("count", 1, "Enter the number of repetitions")

	switch os.Args[1] {
	case "reverse":
		reverse.Parse(os.Args[2:])
		runes := []rune(*reverseName)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		fmt.Printf("%s", string(runes))
	case "repeat":
		repeat.Parse(os.Args[2:])
		for i := 0; i < *repeatCount; i++ {
			fmt.Printf("%s\n", *repeatName)
		}
	default:
		log.Fatal("Error! Expected 'repeat' or 'reverse' subcommands")
	}

}
