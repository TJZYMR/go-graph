package W

import (
	"fmt"
	"y/pattern_up"
	"y/utils"
)

type Letterw struct {
	T [][]int
	P []string
}

func (l *Letterw) Pattern() ([]int, []int, int, bool) {
	a := pattern_up.Peaks(l.T)
	fmt.Println("Peaks =", a)
	b := pattern_up.Troughs(l.T)
	fmt.Println("Troughs =", b)
	// P := []string{"Downward", "Upward", "Downward", "Upward"}
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
		if i1[i] == "Downward" && i != len(i1)-1 {
			if i1[i+1] == "Upward" && i+1 != len(i1)-1 {
				if i1[i+2] == "Downward" && i+2 != len(i1)-1 {
					if i1[i+3] == "Upward" && i+3 != len(i1)-1 {
						count = count + 1
						fmt.Println("Count =", count)
						fmt.Println("Starting Point =", l.T[i2[i]])
						fmt.Println("Ending Point =", l.T[i2[i+3]+2])

						start = append(start, i2[i])
						end = append(end, i2[i+3]+2)
						bol = true

					}
				}
			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}