// main package
package main

import (
	"fmt"
	"math"
)

// Day1 of advent of code
func day1_1(data locationIDLists) {
	sortedFirstLocationIDs := mergeSort(data.firstLocationIDs)
	sortedSecondLocationIDs := mergeSort(data.secondLocationIDs)

	totalDistance := computeTotalDistance(sortedFirstLocationIDs, sortedSecondLocationIDs)

	fmt.Printf("Flag 1: %v\n", totalDistance)
}

func computeTotalDistance(firstLocationIDs, secondLocationIDs []int) int {
	var totalDistance float64
	for i := range firstLocationIDs {
		totalDistance += math.Abs(float64(firstLocationIDs[i]) - float64(secondLocationIDs[i]))
	}

	return int(totalDistance)
}

func day1_2(data locationIDLists) {
	var similarityScore int
	for _, number := range data.firstLocationIDs {
		numberOfOccurances := getOccurances(data.secondLocationIDs, number)
		similarityScore += number * numberOfOccurances
	}
	fmt.Printf("Flag 2: %v\n", similarityScore)
}

// return how many times the number is found in the numbersList
func getOccurances(numbersList []int, number int) int {
	count := 0
	for _, n := range numbersList {
		if number == n {
			count++
		}
	}
	return count
}

func merge(left, right []int) []int {
	data := []int{}

	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			data = append(data, left[leftIndex])
			leftIndex++
		} else {
			data = append(data, right[rightIndex])
			rightIndex++
		}
	}

	for ; leftIndex < len(left); leftIndex++ {
		data = append(data, left[leftIndex])
	}

	for ; rightIndex < len(right); rightIndex++ {
		data = append(data, right[rightIndex])
	}

	return data
}

func mergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	leftData := mergeSort(data[:len(data)/2])
	rightData := mergeSort(data[len(data)/2:])
	return merge(leftData, rightData)
}
