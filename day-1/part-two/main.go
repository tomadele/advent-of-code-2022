package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	inputFile, err := os.ReadFile("input")
	if err != nil {
		log.Printf("Failed to read input file: %v", err)
		os.Exit(1)
	}

	input := strings.Trim(string(inputFile), "\r\n")

	var (
		allCaloriesByElf   = strings.Split(input, "\n\n")
		totalCaloriesByElf []int
	)
	for _, elfCalories := range allCaloriesByElf {
		allFoodsCaloriesByElf := strings.Split(elfCalories, "\n")

		totalCaloriesForThisElf := 0
		for _, foodCaloriesStr := range allFoodsCaloriesByElf {
			foodCalories, err := strconv.Atoi(foodCaloriesStr)
			if err != nil {
				log.Fatalf("Failed to convert food calories (string) to int: %v", err)
			}
			totalCaloriesForThisElf += foodCalories
		}
		totalCaloriesByElf = append(totalCaloriesByElf, totalCaloriesForThisElf)
	}

	slices.Sort(totalCaloriesByElf)
	fmt.Printf("Answer: %d\n", totalCaloriesByElf[len(totalCaloriesByElf)-1]+totalCaloriesByElf[len(totalCaloriesByElf)-2]+totalCaloriesByElf[len(totalCaloriesByElf)-3])
}
