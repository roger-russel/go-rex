package main

import (
	"fmt"

	"github.com/roger-russel/go-rex/gorex"
)

func main() {

	pattern := "^[a-zA-Z0-9_.+-]+@[a-zA-Z-]+\\.[a-z.A-Z0-9.-]+$"
	subject := "fake@email.com.br"

	if gorex.Match(pattern, subject) {
		fmt.Printf("/%s/ matched %s\n", pattern, subject)
	}

	posixPattern := "^[[:alnum:]_.+-]+@[[:alpha:]]+\\.[[:alnum:].-]+$"

	if gorex.Match(posixPattern, subject) {
		fmt.Printf("/%s/ matched %s\n", posixPattern, subject)
	}

}
