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

// type Letterm struct {
// 	Date_struct []Date
// }

// func Peaks(pts []Date) (a []int) {
// 	var b []int
// 	for i := range pts {
// 		if pts[i].Value2 == pts[0].Value2 {
// 			continue
// 		} else if pts[i].Value2 == pts[len(pts)-1].Value2 {
// 			continue
// 		} else if pts[i-1].Value2 < pts[i].Value2 && pts[i+1].Value2 < pts[i].Value2 && pts[i+1] != pts[len(pts)-1] {
// 			if pts[i-2].Value2 < pts[i].Value2 && pts[i+2].Value2 < pts[i].Value2 && pts[i+2] != pts[len(pts)-1] {
// 				b = append(b, i)
// 				continue
// 			}
// 		}
// 	}
// 	return b
// }
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

// //return all throughs
// //from that point previous point be greater and next point be also greater than itself.
// func Troughs(pts []Date) (a []int) {
// 	var b []int
// 	for i := range pts {
// 		if pts[i].Value3 == pts[0].Value3 {
// 			continue
// 		} else if pts[i].Value3 == pts[len(pts)-1].Value3 {
// 			continue
// 		} else if pts[i-1].Value3 > pts[i].Value3 && pts[i+1].Value3 > pts[i].Value3 && pts[i+1] != pts[len(pts)-1] && pts[i-1] != pts[0] {
// 			if pts[i-2].Value3 > pts[i].Value3 && pts[i+2].Value3 > pts[i].Value3 && pts[i+2] != pts[len(pts)-1] && pts[i-2] != pts[0] {
// 				b = append(b, i)
// 				continue
// 			}
// 		}
// 	}
// 	return b
// }
// func (l *Letterm) Pattern() ([]int, []int, int, bool) {
// 	Peaks := Peaks(l.Date_struct)
// 	// fmt.Println("Peaks =", Peaks)
// 	Troughs := Troughs(l.Date_struct)
// 	// fmt.Println("Troughs =", Troughs)
// 	// l.P = []string{"Upward", "Downward", "Upward", "Downward"}
// 	var All_trends_string []string
// 	var mixed_unsorted []int
// 	for i := range Peaks {
// 		mixed_unsorted = append(mixed_unsorted, Peaks[i])
// 	}
// 	for i := range Troughs {
// 		mixed_unsorted = append(mixed_unsorted, Troughs[i])
// 	}
// 	mixed_sorted := utils.Sort(mixed_unsorted)
// 	var All_trends_values []int
// 	for _, i := range mixed_sorted {
// 		if P_up(l.Date_struct[i : i+2]) {
// 			All_trends_string = append(All_trends_string, "Upward")
// 			All_trends_values = append(All_trends_values, i)

// 		} else {
// 			All_trends_string = append(All_trends_string, "Downward")
// 			All_trends_values = append(All_trends_values, i)
// 		}
// 	}
// 	fmt.Println("All_trends_string =", All_trends_string)
// 	var bol bool //to check whether the code is of the patter that we expect it to be.
// 	// fmt.Println("Mix =", mixed_sorted)
// 	// fmt.Println("Pattern =", All_trends_string)
// 	var start []int
// 	var end []int
// 	var count int
// 	for i := range All_trends_string {

// 		if All_trends_string[i] == "Upward" && i != len(All_trends_string)-1 {
// 			if All_trends_string[i+1] == "Downward" && i+1 != len(All_trends_string)-1 {
// 				if All_trends_string[i+2] == "Upward" && i+2 != len(All_trends_string)-1 {
// 					if All_trends_string[i+3] == "Downward" && i+3 != len(All_trends_string)-1 {
// 						// fmt.Println("Count =", count)
// 						// fmt.Println("Starting Point =", l.Date_struct[All_trends_values[i]])
// 						// fmt.Println("Ending Point =", l.Date_struct[All_trends_values[i+3]+1])
// 						// && IsValidCategoryu(All_trends_values[i+2], a1)
// 						// && IsValidCategoryu(All_trends_values[i+3], a2)
// 						var a1 []int //end point
// 						var x int

// 						for i1 := 0; i1 < 50; i1++ {
// 							x = l.Date_struct[All_trends_values[i]].Value2 + i1
// 							a1 = append(a1, x)
// 						}
// 						for i2 := 0; i2 < 50; i2++ {
// 							x = l.Date_struct[All_trends_values[i]].Value3 - i2
// 							a1 = append(a1, x)
// 						}
// 						var a2 []int //start point
// 						var x1 int

// 						for i1 := 0; i1 < 50; i1++ {
// 							x1 = l.Date_struct[All_trends_values[i+1]].Value2 + i1
// 							a2 = append(a2, x1)
// 						}
// 						for i2 := 0; i2 < 50; i2++ {
// 							x1 = l.Date_struct[All_trends_values[i+1]].Value3 - i2
// 							a2 = append(a2, x1)
// 						}
// 						if IsValidCategoryu(l.Date_struct[All_trends_values[i+2]].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]].Value3, a2) {
// 							start = append(start, All_trends_values[i])
// 							end = append(end, All_trends_values[i+3]+4) //last U
// 							count++
// 							bol = true
// 						}
// 						// start = append(start, All_trends_values[i])
// 						// end = append(end, All_trends_values[i+3+1])
// 						// bol = true

// 					}
// 				}
// 			}

// 		}

// 	}
// 	// fmt.Println("Count =", count)
// 	return start, end, count, bol
// }
// func IsValidCategoryu(category int, n []int) bool {
// 	switch category {
// 	case
// 		n[0],
// 		n[1],
// 		n[2],
// 		n[3],
// 		n[4],
// 		n[5],
// 		n[6],
// 		n[7],
// 		n[8],
// 		n[9],
// 		n[10],
// 		n[11],
// 		n[12],
// 		n[13],
// 		n[14],
// 		n[15],
// 		n[16],
// 		n[17],
// 		n[18],
// 		n[19],
// 		n[20],
// 		n[21],
// 		n[22],
// 		n[23],
// 		n[24],
// 		n[25],
// 		n[26],
// 		n[27],
// 		n[28],
// 		n[29],
// 		n[30],
// 		n[31],
// 		n[32],
// 		n[33],
// 		n[34],
// 		n[35],
// 		n[36],
// 		n[37],
// 		n[38],
// 		n[39],
// 		n[40],
// 		n[41],
// 		n[42],
// 		n[43],
// 		n[44],
// 		n[45],
// 		n[46],
// 		n[47],
// 		n[48],
// 		n[49],
// 		n[50],
// 		n[51],
// 		n[52],
// 		n[53],
// 		n[54],
// 		n[55],
// 		n[56],
// 		n[57],
// 		n[58],
// 		n[59],
// 		n[60],
// 		n[61],
// 		n[62],
// 		n[63],
// 		n[64],
// 		n[65],
// 		n[66],
// 		n[67],
// 		n[68],
// 		n[69],
// 		n[70],
// 		n[71],
// 		n[72],
// 		n[73],
// 		n[74],
// 		n[75],
// 		n[76],
// 		n[77],
// 		n[78],
// 		n[79],
// 		n[80],
// 		n[81],
// 		n[82],
// 		n[83],
// 		n[84],
// 		n[85],
// 		n[86],
// 		n[87],
// 		n[88],
// 		n[89],
// 		n[90],
// 		n[91],
// 		n[92],
// 		n[93],
// 		n[94],
// 		n[95],
// 		n[96],
// 		n[97],
// 		n[98],
// 		n[99]:
// 		return true
// 	}
// 	return false
// }

// type Patterns interface {
// 	Pattern() ([]int, []int, int, bool)
// }

// func Find(P Patterns) ([]int, []int, int, bool) {
// 	start, end, count, bol := P.Pattern()
// 	return start, end, count, bol
// }
func main() {
	Date := Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/data.csv")
	// Semicircle.Plot_Original(Date, "Original.png")
	fmt.Println(Date)
	P := &Letter_inverted_semicircle{Date}
	start, end, count, bol := Find(P)
	fmt.Println(start, end, count, bol)
	if bol {
		fmt.Println("M/s found:=>", count)
		fmt.Println("Start: ", start, "End: ", end, "Count: ", count)

	} else {
		fmt.Println("M /s not found")
	}
}
func Genericcsv(filepath string) []Date {
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

type Letter_inverted_semicircle struct {
	Date_struct []Date
}

func (l *Letter_inverted_semicircle) Pattern() ([]int, []int, int, bool) {
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
			for i1 := 0; i1 < 12; i1++ {
				x = pts[i].Value3 + i1
				a1 = append(a1, x)
			}
			for i2 := 0; i2 < 12; i2++ {
				x = pts[i].Value3 - i2
				a1 = append(a1, x)
			}
			if Is_Valid1(pts[i-1].Value3, a1) && pts[i-1].Value3 < pts[i-2].Value3 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue
			}
			if Is_Valid1(pts[i+1].Value3, a1) && pts[i+1].Value3 < pts[i+2].Value3 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue

			}

		} else {
			continue
		}
	}
	return b

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

type Patterns interface {
	Pattern() ([]int, []int, int, bool)
}

func Find(P Patterns) ([]int, []int, int, bool) {
	start, end, count, bol := P.Pattern()
	return start, end, count, bol
} //end of Find and with this adter that dfgfddfdfdf
