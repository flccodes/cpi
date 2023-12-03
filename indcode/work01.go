package indcode

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type IndexesSeries struct {
	Date      time.Time
	Variation float64
}

const (
	Layout20060127     = "2006/01/27"
	Layout27012006     = "27/01/2006"
	Layout27012006full = "27/01/2006 15:04:03 time.UTC"
	Layout012006       = "01/2006"
)

// FreadCSVtoStructs reads a CSV file into a struct
func FreadCSVtoStruct(s string) []IndexesSeries {
	// Opening the CSV file
	f, err := os.Open(s)
	if err != nil {
		log.Printf("Error while opening the CSV file: %s %v", s, err)
	}

	defer f.Close()

	// Reading the CSV file's content
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("Error while parsing CSV data: %v", err)
	}

	// Variable of the type struct which the function is going to return
	var indexData []IndexesSeries

	// Iteranting over the CSV file's content and put it into a struct
	for _, i := range data {
		line := strings.Split(i[0], ";")

		tmpDataCol, err := time.Parse(Layout012006, line[0])
		if err != nil {
			continue
		}

		tmpVarCol, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			continue
		}

		indexData = append(indexData, IndexesSeries{
			Date:      tmpDataCol,
			Variation: tmpVarCol,
		})

	}

	return indexData
}

// Retrieving data from struct

// FgetMonthlyVariationsByDate returns the monthly variation by date
func FgetMonthlyVariationsByDate(data []IndexesSeries, date time.Time) (float64, error) {
	for _, d := range data {
		if d.Date.Equal(date) {
			return d.Variation, nil
		}
	}
	return 0, fmt.Errorf("Data not found for date: %v", date)
}

// FgetMonthlyVariationsInRange returns the variation in an interval of dates
func FgetMonthlyVariationsInRange(data []IndexesSeries, startDate, endDate time.Time) []float64 {
	var variations []float64

	for _, d := range data {
		if d.Date.After(startDate) && d.Date.Before(endDate) {
			variations = append(variations, d.Variation)
		}
	}

	return variations
}

// FgetMonthlyVariation returns the date and the variation at a specific date
func FgetMonthlyVariation(data []IndexesSeries, date time.Time) (IndexesSeries, error) {
	for _, d := range data {
		if d.Date.Equal(date) {
			return IndexesSeries{
				Date:      d.Date,
				Variation: d.Variation,
			}, nil
		}
	}
	return IndexesSeries{}, fmt.Errorf("Data not found for %v", date)
}

// FprintFormattedOutPut prints the formatted output using text/tabwriter
func FprintFormattedOutPut(data []IndexesSeries) {
	// Setting the TabWriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	// Printing the Header
	// fmt.Fprintln(w, "Date\tVariation")
	fmt.Println("Date\tVariation")
	/*
		maxDateLen := 0
		for _, d := range data {
			dateLen := len(d.Date.Format(Layout012006))
			if dateLen > maxDateLen {
				maxDateLen = dateLen
			}
		}
		fmt.Printf("%-*s %s\n", maxDateLen, "Date", "Variation")
	*/
	// Printing the formatted data
	for _, d := range data {
		fmt.Printf("%s\t%0.2f\t\n", d.Date.Format(Layout012006), d.Variation)
	}
	// Flushing the TabWriter
	w.Flush()

}

// FFilterAndPrintWhitinInterval returns the filtered data whitin an interval
func FFilterAndPrintWhitinInterval(data []IndexesSeries, startDate, endDate time.Time) ([]IndexesSeries, error) {
	var filteredData []IndexesSeries
	for _, d := range data {
		if d.Date.After(startDate) && d.Date.Before(endDate) {
			filteredData = append(filteredData, d)
		}
	}
	return filteredData, nil
}

// FprintFormattedOutPut prints the formatted output without tabwriter
func FprintFormattedOutPut02(data []IndexesSeries) {
	// Finding the maximum length of Date
	maxDateLen := 0
	for _, d := range data {
		dateLen := len(d.Date.Format(Layout012006))
		if dateLen > maxDateLen {
			maxDateLen = dateLen
		}
	}

	// Printing the Header
	fmt.Printf("%-*s %s\n", maxDateLen, "Date", "Variation")

	// Printing the formatted data
	for _, d := range data {
		fmt.Printf("%-*s %.2f\n", maxDateLen, d.Date.Format(Layout012006), d.Variation)
	}
}

// FcumulatedVar returns the cumulated variations at an interval
func FcumulatedVar(data []IndexesSeries) ([]float64, error) {
	var cumVars []float64
	var calcVar float64
	calcVar = 1.
	for _, v := range data {
		calcVar *= ((v.Variation / 100) + 1)
		cumVars = append(cumVars, calcVar)
	}

	return cumVars, nil
}
