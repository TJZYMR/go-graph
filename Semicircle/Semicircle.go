package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"y/Structs_for_patterns"
	"y/pattern_up"
	"y/utils"
)

type Lettersemicricle struct {
	T []pattern_up.Date
	P []string
}

//changing the method to find peaks and troughs.
//after matching he pattern upward and then downward
//then matching starting and ending point's value are same or euivalent and if yes then declaring the points as semicircle.
func (l *Lettersemicricle) Pattern() ([]int, []int, int, bool) {
	a := pattern_up.Peaks(l.T)
	fmt.Println("Peaks =", a)
	b := pattern_up.Troughs(l.T)
	fmt.Println("Troughs =", b)
	// l.P = []string{"Upward", "Downward", "Upward", "Downward"}
	var i1 []string
	var w1 []int
	for i := range a {
		w1 = append(w1, a[i])
	}
	for i := range b {
		w1 = append(w1, b[i])
	}
	w2 := utils.Sort(w1)
	var i2 []int
	for _, i := range w2 {
		if pattern_up.P_up(l.T[i : i+2]) {
			i1 = append(i1, "Upward")
			i2 = append(i2, i)

		} else {
			i1 = append(i1, "Downward")
			i2 = append(i2, i)
		}
	}
	var bol bool
	fmt.Println("Mix =", w2)
	fmt.Println("main Pattern =", l.P)
	fmt.Println("Pattern =", i1)
	var start []int
	var end []int
	var count int
	for i := range i1 {
		if i1[i] == "Upward" && i != len(i1)-1 {
			if i1[i+1] == "Downward" && i+1 != len(i1)-1 {

				count = count + 1
				fmt.Println("Count =", count)
				fmt.Println("Starting Point =", l.T[i2[i]])
				fmt.Println("Ending Point =", l.T[i2[i+1]+1])

				start = append(start, i2[i])
				end = append(end, i2[i+1+1])
				bol = true

			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}

//finding semi circle by coding by all the peaks and troughs and p_up p_down functions using.
func main() {

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
	p1 := []string{"Upward", "Downward"}
	w := &Lettersemicricle{date, p1}
	a2, b2, c2, _ := Structs_for_patterns.Find(w)
	fmt.Println(a2, b2, c2)
	utils.Plot1(date, c2, a2, b2)
}
