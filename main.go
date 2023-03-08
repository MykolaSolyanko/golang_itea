package main

import (
	"fmt"
	"os"
	"sort"
)

//
// The Goal
// Casablanca’s hippodrome is organizing a new type of horse racing: duals.
// During a dual, only two horses will participate in the race.
// In order for the race to be interesting, it is necessary to try to select two horses with similar strength.
//
// Write a program which, using a given number of strengths, identifies the two closest strengths
// and shows their difference with an integer (≥ 0).
//
// Input
//Line 1: Number N of horses
//
//The N following lines: the strength Pi of each horse. Pi is an integer.
//
// Output
// The difference D between the two closest strengths. D is an integer greater than or equal to 0.

func main() {
	var strengths []int = getInputData()
	fmt.Println(strengths)
	preProcessStrengths(&strengths)
	fmt.Println(strengths)
	minimumDiff := findMinimumDifference(strengths)
	fmt.Println("Minimum difference: ", minimumDiff)
}

func getInputData() []int {
	var horses int
	fmt.Scanf("%d", &horses)
	if horses < 2 {
		fmt.Printf("The minimum horses value is 2.")
		os.Exit(1)
	}

	strengths := make([]int, horses)

	for i := 0; i < horses; i++ {
		fmt.Scanf("%d", &strengths[i])
	}
	return strengths
}

func preProcessStrengths(strengths *[]int) {
	sort.Ints(*strengths)
}

func findMinimumDifference(strengths []int) int {
	minDiff := strengths[1] - strengths[0]
	for i := 2; i < len(strengths); i++ {
		if dif := strengths[i] - strengths[i-1]; dif < minDiff {
			minDiff = dif
		}
	}
	return minDiff
}
