package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"
)

// Embed input.txt within the binary as a string

//go:embed input.txt
var inputText string

func main() {
	// Parse command line inputs
	part := flag.Int("part", 1, "part '1' or '2'")
	flag.Parse()
	fmt.Println("Running part", *part)

	// Run the selected part and print the answer
	var ans int
	switch *part {
	case 1:
		ans = part1(inputText)
	case 2:
		ans = part2(inputText)
	default:
		fmt.Println("invalid part")
		return
	}
	fmt.Println("Output:", ans)
}

func part1(inputText string) int {
	// Parse the input text to get the integers
	var firstNums []int
	var secondNums []int
	for _, line := range strings.Split(inputText, "\n") {
		if line == "" {
			continue
		}

		var num1, num2 int
		if _, err := fmt.Sscanf(line, "%d %d", &num1, &num2); err == nil {
			firstNums = append(firstNums, num1)
			secondNums = append(secondNums, num2)
		}
	}

	// Sort both in ascending order
	sort.Ints(firstNums)
	sort.Ints(secondNums)

	// Add the diffs between the two arrays
	var ans int
	for i, first := range firstNums {
		second := secondNums[i]
		diff := first - second
		if diff < 0 {
			diff = -diff
		}

		ans += diff
	}

	return ans
}

func part2(input_text string) int {
	// Parse the input text to get the integers
	var firstNums []int
	var secondNums []int
	for _, line := range strings.Split(inputText, "\n") {
		if line == "" {
			continue
		}

		var num1, num2 int
		if _, err := fmt.Sscanf(line, "%d %d", &num1, &num2); err == nil {
			firstNums = append(firstNums, num1)
			secondNums = append(secondNums, num2)
		}
	}

	// Count the occurrences of each number in secondNums
	counts := make(map[int]int)
	for _, num := range secondNums {
		counts[num] += 1
	}

	var ans int
	for _, first := range firstNums {
		similarity := first * counts[first]
		ans += similarity
	}

	return ans
}
