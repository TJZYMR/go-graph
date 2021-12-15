package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type Date struct {
	Day   string
	Month string
	Year  string
	Value string
}

//1.doing for excel file
func Genericexcel(filepath, sheetname string) []Date {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)

	}
	// Get value from cell by given worksheet name and axis.
	var date []Date
	rows, err := f.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)

	}
	rows = rows[1:]
	for _, row := range rows {
		date = append(date, Date{Day: row[0][:2], Month: row[0][3:6], Year: row[0][7:11], Value: row[1]})

	}
	return date
}

//2.doing for csv file
func Genericcsv(filepath, sheetname string) []Date {
	f, err := os.Open("date-values.csv")
	if err != nil {
		fmt.Println(err)

	}
	// Get value from cell by given worksheet name and axis.
	var date []Date
	if err != nil {
		fmt.Println(err)

	}
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	records = records[1:]
	for _, row := range records {
		date = append(date, Date{Day: row[0][:2], Month: row[0][3:6], Year: row[0][7:11], Value: row[1]})

	}
	return date
}

//3.doing for database file
//4.doing for json file
