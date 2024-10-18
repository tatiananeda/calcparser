package main

import (
	u "calcparser/utils"
)

func main() {
	inp := u.RequestInput("Enter the expression. \nOperators supported: * / + - \nNumbers and operators must be separated by whitespace\nExample: 6.7 + 3 * 2")
	u.ProcessInput(inp)
}
