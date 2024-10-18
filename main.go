package main

import (
	"fmt"
	"log"

	u "github.com/tatiananeda/calcparser/utils"
)

func main() {
	inp := u.RequestInput("Enter the expression. \nOperators supported: * / + - \nNumbers and operators must be separated by whitespace\nExample: 6.7 + 3 * 2")
	res, err := u.ProcessInput(inp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %f", res)
}
