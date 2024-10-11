package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/LautaroCasti/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hola %s, al lenguaje bostero.\n", user.Username)
	fmt.Printf("Escriba sus comandos a continuacion\n")
	repl.Start(os.Stdin, os.Stdout)
}
