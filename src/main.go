package main

import (
	"fmt"
	"interpreter/src/repl"
	"os"
	"os/user"
)

//from here we are welcoming the REPL user

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Funny programming langugage!", user.Username)
	fmt.Printf("Feel free to type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
