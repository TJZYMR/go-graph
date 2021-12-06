package M

import (
	"fmt"
	"y/pattern_up"
	"y/utils"
)

func M(t [][]int) ([]int, []int, int, bool) {
	a := pattern_up.Peaks(t)
	fmt.Println("Peaks =", a)
	b := pattern_up.Troughs(t)
	fmt.Println("Troughs =", b)
	p := []string{"Upward", "Downward", "Upward", "Downward"}
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
		if pattern_up.P_up(t[i : i+2]) {
			i1 = append(i1, "Upward")
			i2 = append(i2, i)

		} else {
			i1 = append(i1, "Downward")
			i2 = append(i2, i)
		}
	}
	var bol bool
	fmt.Println("Mix =", w2)
	fmt.Println("main Pattern =", p)
	fmt.Println("Pattern =", i1)
	var start []int
	var end []int
	var count int
	for i := range i1 {
		if i1[i] == "Upward" && i != len(i1)-1 {
			if i1[i+1] == "Downward" && i+1 != len(i1)-1 {
				if i1[i+2] == "Upward" && i+2 != len(i1)-1 {
					if i1[i+3] == "Downward" && i+3 != len(i1)-1 {
						count = count + 1
						fmt.Println("Count =", count)
						fmt.Println("Starting Point =", t[i2[i]])
						fmt.Println("Ending Point =", t[i2[i+3]+1])

						start = append(start, i2[i])
						end = append(end, i2[i+3+1])
						bol = true

					}
				}
			}

		}

	}
	fmt.Println("Count =", count)
	return start, end, count, bol
}

//matching subarray
// if pattern_up.P_up(t[b[0]:a[1]]) {
// 	i = append(i, "Upward")
// } else {
// 	i = append(i, "Downward")

// }

// if pattern_up.P_up(t[a[1]:b[1]]) {
// 	i = append(i, "Upward")

// } else {
// 	i = append(i, "Downward")

// }

// if pattern_up.P_up(t[b[1]:a[2]]) {
// 	i = append(i, "Upward")

// } else {
// 	i = append(i, "Downward")

// }

// if pattern_up.P_up(t[a[2]:a[2]]) {
// 	i = append(i, "Upward")

// } else {
// 	i = append(i, "Downward")

// }
//now write dynamic logic

// if pattern_up.P_up(t[b[0]:a[1]]) {
// 	fmt.Println("Upward")
// 	i = append(i, "Upward")
// } else {
// 	fmt.Println("Downward")
// 	i = append(i, "Downward")

// }
// if i[0] == "Upward" {
// 	if pattern_up.P_up(t[a[1]:b[1]]) {
// 		fmt.Println("Upward")
// 		i = append(i, "Upward")

// 	} else {
// 		fmt.Println("Downward")
// 		i = append(i, "Downward")

// 	}
// }
// if i[1] == "Downward" {
// 	if pattern_up.P_up(t[b[1]:a[2]]) {
// 		fmt.Println("Upward")
// 		i = append(i, "Upward")

// 	} else {
// 		fmt.Println("Downward")
// 		i = append(i, "Downward")

// 	}
// }
// if i[2] == "Upward" {
// 	if pattern_up.P_up(t[a[2]:a[2]]) {
// 		fmt.Println("Upward")
// 		i = append(i, "Upward")

// 	} else {
// 		fmt.Println("Downward")
// 		i = append(i, "Downward")

// 	}
// } //now write dynamic logic

// for index1, i := range a {
// 	for index2, j := range b {
// 		// if i < j {
// 		// 	if pattern_up.P_up(t[a[index1]:b[index2]]) {
// 		// 		i1 = append(i1, "Upward")
// 		// 		break
// 		// 	} else {
// 		// 		i1 = append(i1, "Downward")
// 		// 		break
// 		// 	}
// 		// } else if
// 		if j < i {
// 			if pattern_up.P_up(t[b[index2]:a[index1]]) {
// 				i1 = append(i1, "Upward")
// 				break
// 			} else {
// 				i1 = append(i1, "Downward")
// 				break
// 			}
// 		}
// 	}

// }
