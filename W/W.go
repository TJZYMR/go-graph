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
						count = count + 1
						// fmt.Println("Count =", count)
						// fmt.Println("Starting Point =", l.Date_struct[All_trends_values[i]])
						// fmt.Println("Ending Point =", l.Date_struct[All_trends_values[i+3]+2])

						start = append(start, All_trends_values[i])
						end = append(end, All_trends_values[i+3]+2)
						bol = true

					}
				}
			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}

type Patterns interface {
	Pattern() ([]int, []int, int, bool)
}

func Find(P Patterns) ([]int, []int, int, bool) {
	start, end, count, bol := P.Pattern()
	return start, end, count, bol
}
