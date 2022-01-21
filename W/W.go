package W

import (
	"y/pattern_up"
	"y/utils"
)

type Letterw struct {
	Date_struct []pattern_up.Date
}

func (l *Letterw) Pattern() ([]int, []int, int, bool) {
	Peaks := pattern_up.Peaks(l.Date_struct)
	// fmt.Println("Peaks =", Peaks)
	Troughs := pattern_up.Troughs(l.Date_struct)
	// fmt.Println("Troughs =", Troughs)
	// P := []string{"Downward", "Upward", "Downward", "Upward"}
	var All_trends_string []string
	var mixed_unsorted []int
	for i := range Peaks {
		mixed_unsorted = append(mixed_unsorted, Peaks[i])
	}
	for i := range Troughs {
		mixed_unsorted = append(mixed_unsorted, Troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted)
	var All_trends_values []int
	for _, i := range mixed_sorted {
		if pattern_up.P_up(l.Date_struct[i : i+2]) {
			All_trends_string = append(All_trends_string, "Upward")
			All_trends_values = append(All_trends_values, i)

		} else {
			All_trends_string = append(All_trends_string, "Downward")
			All_trends_values = append(All_trends_values, i)
		}
	}
	var bol bool //to check whether the code is of the patter that we expect it to be.
	// fmt.Println("Mix =", mixed_sorted)
	// fmt.Println("Pattern =", All_trends_string)
	var start []int
	var end []int
	var count int
	for i := range All_trends_string {
		if All_trends_string[i] == "Downward" && i != len(All_trends_string)-1 {
			if All_trends_string[i+1] == "Upward" && i+1 != len(All_trends_string)-1 {
				if All_trends_string[i+2] == "Downward" && i+2 != len(All_trends_string)-1 {
					if All_trends_string[i+3] == "Upward" && i+3 != len(All_trends_string)-1 {
						// fmt.Println("Count =", count)
						// fmt.Println("Starting Point =", l.Date_struct[All_trends_values[i]])
						// fmt.Println("Ending Point =", l.Date_struct[All_trends_values[i+3]+2])

						var a1 []int //end point
						var x int

						for i1 := 0; i1 < 50; i1++ {
							x = l.Date_struct[All_trends_values[i]].Value + i1
							a1 = append(a1, x)
						}
						for i2 := 0; i2 < 50; i2++ {
							x = l.Date_struct[All_trends_values[i]].Value - i2
							a1 = append(a1, x)
						}
						var a2 []int //start point
						var x1 int

						for i1 := 0; i1 < 50; i1++ {
							x1 = l.Date_struct[All_trends_values[i+1]].Value + i1
							a2 = append(a2, x1)
						}
						for i2 := 0; i2 < 50; i2++ {
							x1 = l.Date_struct[All_trends_values[i+1]].Value - i2
							a2 = append(a2, x1)
						}
						if IsValidCategoryu(l.Date_struct[All_trends_values[i+2]].Value, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]].Value, a2) {
							start = append(start, All_trends_values[i])
							end = append(end, All_trends_values[i+3]+4) //last U
							count++
							bol = true
						}

					}
				}
			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}
func IsValidCategoryu(category int, n []int) bool {
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
		n[59],
		n[60],
		n[61],
		n[62],
		n[63],
		n[64],
		n[65],
		n[66],
		n[67],
		n[68],
		n[69],
		n[70],
		n[71],
		n[72],
		n[73],
		n[74],
		n[75],
		n[76],
		n[77],
		n[78],
		n[79],
		n[80],
		n[81],
		n[82],
		n[83],
		n[84],
		n[85],
		n[86],
		n[87],
		n[88],
		n[89],
		n[90],
		n[91],
		n[92],
		n[93],
		n[94],
		n[95],
		n[96],
		n[97],
		n[98],
		n[99]:
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
}
