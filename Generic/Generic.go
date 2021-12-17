package Generic

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"y/pattern_up"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

// type Date struct {
// 	Index int
// 	Date  string
// 	Value int
// }

//1.doing for excel file
func Genericexcel(filepath, sheetname string) []pattern_up.Date {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)

	}
	// Get value from cell by given worksheet name and axis.
	var date []pattern_up.Date
	rows, err := f.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)

	}
	rows = rows[1:]
	for i, row := range rows {
		b, _ := strconv.Atoi(row[1])
		date = append(date, pattern_up.Date{Index: i, Date: row[0], Value: b})

	}
	return date
}

//2.doing for csv file
func Genericcsv(filepath string) []pattern_up.Date {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)

	}
	// Get value from cell by given worksheet name and axis.
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
		date = append(date, pattern_up.Date{Index: Index, Date: row[0], Value: b})

	}
	return date
}

//3.doing for database file
func Genericdatabase(port, username, dbname, password string) []pattern_up.Date {
	db, _ := sql.Open("mysql", username+":"+password+"@tcp("+port+")/"+dbname)
	// db, _ := sql.Open("mysql", "tatva:zymr@123@tcp(127.0.0.1:3306)/tatva")
	defer db.Close()
	ab, _ := db.Query("select * from tatva.date_values")
	var date []pattern_up.Date
	type T struct {
		Index int
		Tar   string
		Value int
	}
	var i int = 0
	for ab.Next() {
		i++
		var t T
		ab.Scan(&t.Tar, &t.Value)
		date = append(date, pattern_up.Date{Index: i, Date: t.Tar, Value: t.Value})
	}
	return date
}

//4.doing for json file
func Genericjson(filename string) []pattern_up.Date {
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	type T struct {
		Index int    `json:"Index"`
		Tar   string `json:"Date"`
		Value int    `json:"Value"`
	}
	var t []T
	json.Unmarshal(byteValue, &t)
	var date []pattern_up.Date
	for i, row := range t {
		date = append(date, pattern_up.Date{Index: i, Date: row.Tar, Value: row.Value})

	}
	return date
}

//main file
// func main() {
// 	d := Genericdatabase("127.0.0.1:3306", "tatva", "tatva", "zymr@123")
// 	fmt.Println("here it runs", "\n", d)
// }
// func main() {
// 	d := Genericjson("date-values.json")
// 	fmt.Println("here it runs", "\n", d)

// }
