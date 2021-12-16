package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"y/pattern_up"
)

type Lettersemicricle struct {
	T []pattern_up.Date
	P []string
}

//1.changing the method to find peaks and troughs.
//2.after matching he pattern upward and then downward
//3.then matching starting and ending point's value are same or euivalent and if yes then declaring the points as semicircle.
func (l *Lettersemicricle) Pattern() {
	a := Peaks(l.T)
	fmt.Println("Peaks =", a)
	b := Troughs(l.T)
	fmt.Println("Troughs =", b)
	// l.P = []string{"Upward", "Downward", "Upward", "Downward"}
	var a1 [][]int
	var a2 []int

	for _, j := range a {
		a1 = append(a1, []int{l.T[j].Value, l.T[j].Value})
		a2 = append(a2, l.T[j].Value)
	}
	a3 := MinMax(a2)
	fmt.Println("Max value of peaks =", a3)
	// fmt.Println("Count =", count)

}
func MinMax(array []int) int {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max
}

// type Date struct {
// 	Index int
// 	Date  string
// 	Value int
// }

func P_up(pts []pattern_up.Date) (a bool) { //pts [][]int

	var b bool = false
	for i := range pts {
		if len(pts) != 1 {
			if pts[i].Value < pts[i+1].Value && pts[i+1].Date != pts[len(pts)-1].Date {
				continue
			} else if pts[i].Value < pts[i+1].Value && pts[i+1].Date == pts[len(pts)-1].Date {
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

func P_down(pts []pattern_up.Date) (a bool) {
	var b bool = false
	for i := range pts {
		if pts[i].Value > pts[i+1].Value && pts[i+1].Date != pts[len(pts)-1].Date {
			continue
		} else if pts[i].Value > pts[i+1].Value && pts[i+1].Date == pts[len(pts)-1].Date {
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

//return all peaks
//from that point previous point be lesser and next point be also lesser than itself.
//last four and next four points should be lesser than the peak point.
func Peaks(pts []pattern_up.Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value < pts[i].Value && pts[i+1].Value < pts[i].Value && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {
			if pts[i-2].Value < pts[i].Value && pts[i+2].Value < pts[i].Value && pts[i-2] != pts[0] && pts[i+2] != pts[len(pts)-1] {
				if pts[i-3].Value < pts[i].Value && pts[i+3].Value < pts[i].Value && pts[i-3] != pts[0] && pts[i+3] != pts[len(pts)-1] {
					if pts[i+4].Value < pts[i].Value && pts[i-4].Value < pts[i].Value && pts[i-4] != pts[0] && pts[i+4] != pts[len(pts)-1] {
						b = append(b, i)
						continue

					} else {
						continue
					}
				} else {
					continue
				}
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return b
}

//return all throughs
//from that point previous point be greater and next point be also greater than itself.
func Troughs(pts []pattern_up.Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value > pts[i].Value && pts[i+1].Value > pts[i].Value && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {
			if pts[i-2].Value > pts[i].Value && pts[i+2].Value > pts[i].Value && pts[i-2] != pts[0] && pts[i+2] != pts[len(pts)-1] {
				if pts[i-3].Value > pts[i].Value && pts[i+3].Value > pts[i].Value && pts[i-3] != pts[0] && pts[i+3] != pts[len(pts)-1] {
					if pts[i+4].Value > pts[i].Value && pts[i-4].Value > pts[i].Value && pts[i-4] != pts[0] && pts[i+4] != pts[len(pts)-1] {
						b = append(b, i)
						continue
					} else {
						continue
					}
				} else {
					continue
				}
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return b

}

//0 {11,100}

//finding semi circle by coding by all the peaks and troughs and p_up p_down functions using.
func main() {
	date := Getdata()
	// p1 := []string{"Upward", "Downward"}
	// w := &Lettersemicricle{date, p1}

	a := Peaks(date)
	fmt.Println("Peaks =", a)
	b := Troughs(date)
	fmt.Println("Troughs =", b)
	// l.P = []string{"Upward", "Downward", "Upward", "Downward"}
	var a1 [][]int
	var a2 []int

	for _, j := range a {
		a1 = append(a1, []int{date[j].Index, date[j].Value})
		a2 = append(a2, date[j].Value)
	}
	a3 := MinMax(a2)
	fmt.Println("Max value of peaks =", a3)
	for _, j := range a {
		if a3 == date[j].Value {
			fmt.Println("Max value's index  =", j)
		}
	}
}
func Getdata() []pattern_up.Date {
	f, err := os.Open("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/s1.csv")
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
	return date
}
