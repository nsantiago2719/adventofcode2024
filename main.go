// main package
package main

import (
	"fmt"
	"os"
)

type locationIDLists struct {
	firstLocationIDs  []int
	secondLocationIDs []int
}

func main() {
	csvFile, err := os.Open("./data/day-1.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	// read csv and unmarshal data to the struct
	data, err := unmarshalData(csvFile)
	if err != nil {
		fmt.Println(err)
	}

	day1_1(data)
	day1_2(data)
}
