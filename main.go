package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(counter(os.Args))
}

func counter(args []string) int {
	if len(args[1]) < 1 {
		return 0
	}
	splitString := strings.Split(args[1], " ")
	count := len(splitString)
	return count
}
