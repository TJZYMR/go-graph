package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"y/pattern_up"
)

//finding semi circle by coding by all the peaks and troughs and p_up p_down functions using.
func main() {

	f, err := os.Open("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/date-values.csv")
	if err != nil {
		fmt.Println(err)

	}
	var date []pattern_up.Date
	if err != nil {
		fmt.Println(err)

	}
	csvReader := csv.NewReader(f)
	records, _ := csvReader.ReadAll()
	records = records[1:]
	var Index int
	for _, row := range records {
		Index = Index + 1
		b, _ := strconv.Atoi(row[1])
		date = append(date, pattern_up.Date{Index, row[0], b})
	}
}
