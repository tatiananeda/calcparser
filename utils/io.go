package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RequestInput(q string) string {
	fmt.Println(q)

	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return line
}
