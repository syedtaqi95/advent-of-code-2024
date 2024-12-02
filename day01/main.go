package main

import (
	_ "embed"
	"flag"
	"fmt"
)

// Embed input.txt within the binary as a string
//go:embed input.txt
var input_text string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Parse command line inputs
	var part int
	flag.IntVar(&part, "part", 1, "part '1' or '2'")
	flag.Parse()
	fmt.Println("Running part", part)

	// Run the selected part and print the answer
	var ans int
	if part == 1 {
		ans = part1(input_text)
	} else {
		ans = part2(input_text)
	}
	fmt.Println("Output:", ans)
}

func part1(input_text string) int {
	return 1
}

func part2(input_text string) int {
	return 2
}
