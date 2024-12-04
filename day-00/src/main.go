package main

import (
	"fmt"
	"strconv"
	"os"
	"main/mymath"
)

func parseArgs(args []string) ([]int, error) {
	var numbers []int
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", arg)
		}
		if num < -100000 || num > 100000 {
			return nil, fmt.Errorf("number %d out of bounds, must be between -100000 and 100000", num)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide numbers as arguments.")
		return
	}

	args := os.Args[1:]

	numbers, err := parseArgs(args)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)
		switch v := result.(type) {
		case int:
			fmt.Printf("%s: %d\n", statName[i], v)
		case float64:
			fmt.Printf("%s: %.2f\n", statName[i], v)
		default:
			fmt.Printf("%s: idk how its calling\n", statName[i]) //probably unused
		}
    }
}