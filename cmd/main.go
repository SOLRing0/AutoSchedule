package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 2 {
		fmt.Println("ERROR: expected input arguments are: [people.csv] [schedule.csv]")
		return
	}
}