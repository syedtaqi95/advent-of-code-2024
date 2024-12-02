package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

// Embed input.txt within the binary as a string

//go:embed input.txt
var input_text string

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
	// Split the input text into lines
	// Split each line to get the integers
	var ans int
	for _, line := range strings.Split(input_text, "\n") {
		if line == "" {
			continue
		}

		var num1, num2 int
		if _, err := fmt.Sscanf(line, "%d %d", &num1, &num2); err == nil {
			diff := num1 - num2
			if diff < 0 {
				diff = -diff
			}

			ans += diff
		}
	}

	return ans
}

func part2(input_text string) int {
	return 2
}
