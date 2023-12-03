package indcode

/*
Some very short examples:

https://go.dev/play/p/wnQ6GrTMwut

https://go.dev/play/p/bFVHEvol0AS

*/

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// FreadCSVdata opens and put the CSV file's content into a struct using a bufio.NewReader
func FreadCSVdata(s string) ([]IndexesSeries, error) {
	f, err := os.Open(s)
	if err != nil {
		log.Println("Error while reading the data")
	}
	defer f.Close()

	var indiceOutPut []IndexesSeries

	fileReader := bufio.NewReader(io.Reader(f))
	for {
		line, err := fileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		fields := strings.Split(line, ";")
		tmpDate, err := time.Parse(LayoutDDMMYYYY, fields[0])
		if err != nil {
			continue
		}
		tmpVar, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			continue
		}
		indices := IndexesSeries{
			Date:      tmpDate,
			Variation: tmpVar,
		}
		indiceOutPut = append(indiceOutPut, indices)
	}

	return indiceOutPut, nil

}
