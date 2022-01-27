package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"y/utils"
)

type Date struct {
	Index  int
	Date   string
	Value1 int
	Value2 int
	Value3 int
	Value4 int
}

func main() {
	var genfile Genericfiles1 = Date{}
	Date := genfile.Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/data.csv")
	// Semicircle.Plot_Original(Date, "Original.png")
	fmt.Println(Date)
	var p Patterns = &Date_struct{Date}
	start, end, count, bol := p.Pattern()
	fmt.Println(start, end, count, bol)
	if bol {
		fmt.Println("inverted_semicircle/s found:=>", count)
		fmt.Println("Start: ", start, "End: ", end, "Count: ", count)
		fmt.Println("Start: ", Date[start[0]], Date[start[1]], "End: ", Date[end[0]], Date[end[1]], "Count: ", count)
		// Semicircle.Plot_Pattern(Date, count, start, end, "Pattern.png")
	} else {
		fmt.Println("inverted_semicircle /s not found")
	}
}

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
		b1, _ := strconv.Atoi(row[2])
		b2, _ := strconv.Atoi(row[3])
		b3, _ := strconv.Atoi(row[4])
		date = append(date, Date{Index: Index, Date: row[0], Value1: b, Value2: b1, Value3: b2, Value4: b3})

	}
	return date
}

type Date_struct struct {
	Date_struct []Date
}

func (l *Date_struct) Pattern() ([]int, []int, int, bool) {
	peaks := peaks(l.Date_struct)
	troughs := troughs(l.Date_struct)
	fmt.Println(peaks)
	fmt.Println(troughs)
	var All_trends_string []string
	var mixed_unsorted []int
	for i := range peaks {
		mixed_unsorted = append(mixed_unsorted, peaks[i])
	}
	for i := range troughs {
		mixed_unsorted = append(mixed_unsorted, troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted) //mixed values
	var All_trends_index_values []int
	for _, i := range mixed_sorted {
		if P_up(l.Date_struct[i : i+2]) {
			All_trends_string = append(All_trends_string, "Upward")
			All_trends_index_values = append(All_trends_index_values, i)

		} else {
			All_trends_string = append(All_trends_string, "Downward")
			All_trends_index_values = append(All_trends_index_values, i)
		}
	}

	var start []int
	var end []int
	var count int
	var bol bool //just to check if the pattern is semicircle or not,this is boolean that will tell us that.

	for i := range All_trends_string {
		if All_trends_string[i] == "Downward" && i != len(All_trends_string)-1 {
			if All_trends_string[i+1] == "Upward" && i+1 != len(All_trends_string)-1 { //Task:=>check here for first and last point to be in the align manner.
				var a1 []int //end point
				var x int

				for i1 := 0; i1 < 30; i1++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value3 + i1
					a1 = append(a1, x)
				}
				for i2 := 0; i2 < 30; i2++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value3 - i2
					a1 = append(a1, x)
				}
				var a2 []int //start point
				var x1 int

				for i1 := 0; i1 < 30; i1++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value2 + i1
					a2 = append(a2, x1)
				}
				for i2 := 0; i2 < 30; i2++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value2 - i2
					a2 = append(a2, x1)
				}
				// if Is_Valid11(l.Date_struct[All_trends_index_values[i]].Value, a1) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i])
				// 	count++
				// 	bol = true //do tuing here
				// } else
				if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value3, a1) {
					start = append(start, All_trends_index_values[i]+1)
					end = append(end, All_trends_index_values[i+2]+1) //last U
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+2].Value3, a1) {
					start = append(start, All_trends_index_values[i]+2)
					end = append(end, All_trends_index_values[i+2]+1)
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]].Value2, a2) {
					start = append(start, All_trends_index_values[i])
					end = append(end, All_trends_index_values[i]+3)
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value2, a2) {
					start = append(start, All_trends_index_values[i])
					end = append(end, All_trends_index_values[i]+5) //some connection
					count++
					bol = true
				}
			}
		}

	}

	// fmt.Println("Count =", count)
	return start, end, count, bol
}

func peaks(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value2 == pts[0].Value2 {
			continue
		} else if pts[i].Value2 == pts[len(pts)-1].Value2 {
			continue
		} else if pts[i-1].Value2 < pts[i].Value2 && pts[i+1].Value2 < pts[i].Value2 && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {

			b = append(b, i)
			continue

		} else {
			continue
		}
	}
	return b
}
func troughs(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value3 == pts[0].Value3 {
			continue
		} else if pts[i].Value3 == pts[len(pts)-1].Value3 {
			continue
		} else if pts[i-1].Value3 > pts[i].Value3 && pts[i+1].Value3 > pts[i].Value3 && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {
			var a1 []int
			var x int
			for i1 := 0; i1 < 30; i1++ {
				x = pts[i].Value3 + i1
				a1 = append(a1, x)
			}
			for i2 := 0; i2 < 30; i2++ {
				x = pts[i].Value3 - i2
				a1 = append(a1, x)
			}
			if Is_Valid30(pts[i-1].Value3, a1) && pts[i-1].Value3 < pts[i-2].Value3 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue
			}
			if Is_Valid30(pts[i+1].Value3, a1) && pts[i+1].Value3 < pts[i+2].Value3 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue

			}

		} else {
			continue
		}
	}
	return b

}
func Is_Valid1(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23]:
		// n[24],
		// n[25],
		// n[26],
		// n[27],
		// n[28],
		// n[29],
		// n[30],
		// n[31],
		// n[32],
		// n[33],
		// n[34]:
		// n[35],
		// n[36],
		// n[37],
		// n[38],
		// n[39]:
		// n[40],
		// n[41],
		// n[42],
		// n[43],
		// n[44],
		// n[45],
		// n[46],
		// n[47],
		// n[48]:

		return true
	}
	return false
}

func Is_Valid11(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23],
		n[24],
		n[25],
		n[26],
		n[27],
		n[28],
		n[29],
		n[30],
		n[31],
		n[32],
		n[33],
		n[34],
		n[35],
		n[36],
		n[37],
		n[38],
		n[39],
		n[40],
		n[41],
		n[42],
		n[43],
		n[44],
		n[45],
		n[46],
		n[47],
		n[48],
		n[49],
		n[50],
		n[51],
		n[52],
		n[53],
		n[54],
		n[55],
		n[56],
		n[57],

		n[58],
		n[59]:
		return true
	}
	return false
}
func Is_Valid30(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23],
		n[24],
		n[25],
		n[26],
		n[27],
		n[28],
		n[29]:
		return true
	}
	return false
}

type Patterns interface {
	Pattern() ([]int, []int, int, bool)
}
type Genericfiles1 interface {
	Genericcsv(filepath string) []Date
}

func P_up(pts []Date) (a bool) { //pts [][]int

	var b bool = false
	for i := range pts {
		if len(pts) != 1 {
			if pts[i].Value2 < pts[i+1].Value2 && pts[i+1].Date != pts[len(pts)-1].Date {
				continue
			} else if pts[i].Value2 < pts[i+1].Value2 && pts[i+1].Date == pts[len(pts)-1].Date {
				//fmt.Println("Down trend breaks here", pts[i+1])
				b = true
				break
			} else {
				b = false
				break
			}
		} else {
			b = true
			break
		}

	}
	return b
}

func P_down(pts []Date) (a bool) {
	var b bool = false
	for i := range pts {
		if pts[i].Value3 > pts[i+1].Value3 && pts[i+1].Date != pts[len(pts)-1].Date {
			continue
		} else if pts[i].Value3 > pts[i+1].Value3 && pts[i+1].Date == pts[len(pts)-1].Date {
			//fmt.Println("Down trend breaks here", pts[i+1])
			b = true
			break
		} else {
			b = false
			break
		}
	}
	return b
}
