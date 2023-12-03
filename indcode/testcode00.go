package indcode

// https://go.dev/play/p/DxsVkkZHEyJ

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type IndexValues struct {
	Date       time.Time
	MonthlyVar float64
}

// Const

const (
	LayoutYYYYMMDD     = "2006/02/01"
	LayoutDDMMYYYY     = "02/01/2006"
	LayoutDDMMYYYYfull = "02/01/2006 15:04:05 -03"
	Layout             = "22 Jan 2006 15:04:05"
)

// FopenCSVfile opens the CSV file and returns ...
func FopenCSVfile(s string) ([]IndexValues, error) {
	f, err := os.Open(s)
	if err != nil {
		log.Println("Occured an Error while opening the CSV file!", err)
	}
	defer f.Close()

	// Create the objected that will be return
	var indice []IndexValues

	// Create a buffered reader to read the file
	fileReader := bufio.NewReader(io.Reader(f))

	for {

		// Read a line from the file
		line, err := fileReader.ReadString('\n')

		// If there is an error, break out of the loop
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// Split the line into fields
		fields := strings.Split(line, ";")

		tmpDate, err := time.Parse(LayoutDDMMYYYY, fields[0])

		tmpMonthlyVar, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			continue
		}

		// Create a Temp Object
		tempIndexes := IndexValues{
			Date:       tmpDate,
			MonthlyVar: tmpMonthlyVar,
		}

		// Store the indexes in the slice
		indice = append(indice, tempIndexes)

	}

	return indice, nil

}
