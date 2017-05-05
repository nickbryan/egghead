package main

import (
	"fmt"
	"github.com/nickbryan/egghead/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Egghead REPL!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
