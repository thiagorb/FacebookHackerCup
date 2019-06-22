package main

import (
	"bufio"
	"fmt"
	"mrx"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	/*
		var f io.Reader
		var err error
		if f, err = os.Open("../problem.txt"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		scanner = bufio.NewScanner(bufio.NewReader(f))
		//*/

	if !scanner.Scan() {
		fmt.Println("Failed to read number of tests")
		os.Exit(1)
	}

	numberOfTests, err := strconv.ParseUint(scanner.Text(), 10, 9)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for test := uint64(1); test <= numberOfTests; test++ {
		if !scanner.Scan() {
			fmt.Printf("Failed to read case %d\n", test)
		}

		expression := mrx.Parse(scanner.Text())

		minimumModifications := mrx.CalculateMinimumModifications(expression)
		fmt.Printf("Case #%d: %d\n", test, minimumModifications)
	}
}
