package main

/*

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	numberOfDays, err := strconv.ParseUint(os.Args[1], 10, 9)

	if err != nil {
		fmt.Println("Unable to parse number of days:", err)
		os.Exit(1)
	}

	numberOfLilypads, err := strconv.ParseUint(os.Args[2], 10, 13)

	if err != nil {
		fmt.Println("Unable to parse number of lilypads")
		os.Exit(1)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println(numberOfDays)

	for i := uint64(0); i < numberOfDays; i++ {
		fmt.Print("A")
		for j := uint64(0); j < numberOfLilypads; j++ {
			if rand.Intn(2) > 0 {
				fmt.Print("B")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
*/
