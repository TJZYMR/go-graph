package Generic

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

// type Date struct {
// 	Index int
// 	Date  string
// 	Value int
// }
type Date struct {
	Index int
	Date  string
	Value int
}

//1.doing for excel file
type Genericfiles interface {
	Genericexcel(filepath, sheetname string) []Date
	Genericcsv(filepath string) []Date
	Genericdatabase(port, username, dbname, password string) []Date
	Genericjson(filename string) []Date
}

func (d Date) GenericProto(filepath, sheetname string) []Date {
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
	for i, row := range rows {
		b, _ := strconv.Atoi(row[1])
		date = append(date, Date{Index: i, Date: row[0], Value: b})

	}
	return date
}
func (d Date) Genericexcel(filepath, sheetname string) []Date {
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
	for i, row := range rows {
		b, _ := strconv.Atoi(row[1])
		date = append(date, Date{Index: i, Date: row[0], Value: b})

	}
	return date
}

//sockets can also be introduced.
//2.doing for csv file
func (d Date) Genericcsv(filepath string) []Date {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)

	}
	// Get value from cell by given worksheet name and axis.
	var date []Date
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
		date = append(date, Date{Index: Index, Date: row[0], Value: b})

	}
	return date
}

//3.doing for database file
func (d Date) Genericdatabase(port, username, dbname, password string) []Date {
	db, _ := sql.Open("mysql", username+":"+password+"@tcp("+port+")/"+dbname)
	// db, _ := sql.Open("mysql", "tatva:zymr@123@tcp(127.0.0.1:3306)/tatva")
	defer db.Close()
	ab, _ := db.Query("select * from tatva.date_values")
	var date []Date
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
		date = append(date, Date{Index: i, Date: t.Tar, Value: t.Value})
	}
	return date
}

//4.doing for json file
func (d Date) Genericjson(filename string) []Date {
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
	var date []Date
	for i, row := range t {
		date = append(date, Date{Index: i, Date: row.Tar, Value: row.Value})

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
