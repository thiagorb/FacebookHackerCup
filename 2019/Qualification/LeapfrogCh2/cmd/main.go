package main

import (
	"brute"
	"bufio"
	"fmt"
	"optimized"
	"os"
	"strconv"
)

func main() {

	canReachRightmostLilypad := optimized.CanReachRightmostLilypad

	if len(os.Args) > 1 && os.Args[1] == "--brute" {
		canReachRightmostLilypad = brute.CanReachRightmostLilypad
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numberOfDays, err := strconv.ParseUint(scanner.Text(), 10, 9)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for day := uint64(1); day <= numberOfDays; day++ {
			if !scanner.Scan() {
				fmt.Printf("Failed to read case %d\n", day)
				os.Exit(1)
			}

			var resultString string
			if canReachRightmostLilypad(scanner.Text()) {
				resultString = "Y"
			} else {
				resultString = "N"
			}

			fmt.Printf("Case #%d: %s\n", day, resultString)
		}
	}
}
