package indcode

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadCSVtoSlices(s string) [][]string {
	f, err := os.Open(s)
	if err != nil {
		log.Println("Error while opening the CSV file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error while parsing CSV data.", err)
	}

	return data
}

// FStrToSlices takes a slice of strings and returns two slices of different types
func FStrToSlices(s string) ([]time.Time, []float64) {
	data := ReadCSVtoSlices(s)

	sliceDate := make([]time.Time, 0)
	sliceVar := make([]float64, 0)

	for _, i := range data {
		for _, lin := range i {

			line := strings.Split(lin, ";")

			tmpDataCol, err := time.Parse(Layout012006, string(line[0]))
			if err != nil {
				continue
			}

			sliceDate = append(sliceDate, tmpDataCol)

			tmpVarCol, err := strconv.ParseFloat(string(line[1]), 64)
			if err != nil {
				continue
			}

			sliceVar = append(sliceVar, tmpVarCol)

		}

	}

	return sliceDate, sliceVar
}
