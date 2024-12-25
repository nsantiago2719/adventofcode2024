// main package
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func unmarshalData(csvFile *os.File) (locationIDLists, error) {
	csvReader := csv.NewReader(csvFile)
	rawData, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	var data locationIDLists

	for _, rawDatum := range rawData {
		firstLocationIDs, err := strconv.Atoi(rawDatum[0])
		if err != nil {
			fmt.Println(err)
		}
		secondLocationIDs, err := strconv.Atoi(rawDatum[1])
		if err != nil {
			fmt.Println(err)
		}
		data.firstLocationIDs = append(data.firstLocationIDs, firstLocationIDs)
		data.secondLocationIDs = append(data.secondLocationIDs, secondLocationIDs)
	}

	return data, nil
}
