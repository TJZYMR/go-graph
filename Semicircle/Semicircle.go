package Semicircle

import (
	"fmt"
	"y/Generic"
	"y/pattern_up"
	"y/utils"
)

type Lettersemicricle struct {
	Date_struct []Generic.Date
	Patternn    []string
}

//1.changing the method to find peaks and troughs.
//2.after matching he pattern upward and then downward
//3.then matching starting and ending point's value are same or euivalent and if yes then declaring the points as semicircle.
func (l *Lettersemicricle) Pattern() ([]int, []int, int, bool) {
	Peaks := Peaks(l.Date_struct)
	Troughs := Troughs(l.Date_struct)
	fmt.Println("Peaks", Peaks)
	fmt.Println("Troughs", Troughs)

	var All_trends_string []string
	var mixed_unsorted []int
	for i := range Peaks {
		mixed_unsorted = append(mixed_unsorted, Peaks[i])
	}
	for i := range Troughs {
		mixed_unsorted = append(mixed_unsorted, Troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted) //mixed values
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
	var bol bool //just to check if the pattern is semicircle or not,this is boolean that will tell us that.
	fmt.Println("Mix =", mixed_sorted)
	for _, j := range mixed_sorted {
		fmt.Println("Mix =", l.Date_struct[j])
	}
	fmt.Println("main Pattern =", l.Patternn)
	fmt.Println("Pattern =", All_trends_string)
	var start []int
	var end []int
	var count int
	for i := range All_trends_string {
		if All_trends_string[i] == "Upward" && i != len(All_trends_string)-1 {

			if All_trends_string[i+1] == "Downward" && i+1 != len(All_trends_string)-1 {

				count = count + 1
				// fmt.Println("Count =", count)
				// fmt.Println("Starting Point =", l.T[All_trends_values[i]])
				// fmt.Println("Ending Point =", l.T[All_trends_values[i+1]+1])

				start = append(start, All_trends_values[i])
				end = append(end, All_trends_values[i+1+1])
				bol = true

			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}
func Peaks(pts []Generic.Date) (a []int) {
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
						if pts[i+5].Value < pts[i].Value && pts[i-5].Value < pts[i].Value && pts[i-5] != pts[0] && pts[i+5] != pts[len(pts)-1] {
							if pts[i+6].Value < pts[i].Value && pts[i-6].Value < pts[i].Value && pts[i-6] != pts[0] && pts[i+6] != pts[len(pts)-1] {
								if pts[i+7].Value < pts[i].Value && pts[i-7].Value < pts[i].Value && pts[i-7] != pts[0] && pts[i+7] != pts[len(pts)-1] {
									if pts[i+8].Value < pts[i].Value && pts[i-8].Value < pts[i].Value && pts[i-8] != pts[0] && pts[i+8] != pts[len(pts)-1] {
										if pts[i+9].Value < pts[i].Value && pts[i-9].Value < pts[i].Value && pts[i-9] != pts[0] && pts[i+9] != pts[len(pts)-1] {
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
		} else {
			continue
		}
	}
	return b
}

func Troughs(pts []Generic.Date) (a []int) {
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
						if pts[i+5].Value > pts[i].Value && pts[i-5].Value > pts[i].Value && pts[i-5] != pts[0] && pts[i+5] != pts[len(pts)-1] {
							if pts[i+6].Value > pts[i].Value && pts[i-6].Value > pts[i].Value && pts[i-6] != pts[0] && pts[i+6] != pts[len(pts)-1] {
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
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return b

}

func Adjustpoints(start, end []int, count int, date []Generic.Date) ([]int, []int, int) { //
	//1->make these variables as many times as count variable c.

	nearby_points := make([][]int, count)
	//var nearby_points2 []int
	t := make([][]int, count)  //this variable for the index of the first pattern
	t1 := make([][]int, count) //this variable for the index of the second pattern

	// var t1 [][]int

	//1
	//2->make automatic isvalidcategory function and loop thorough the points until the last point.
	for i := 0; i < count; i++ {
		for _, All_trends_string := range date[start[i] : end[i]+1] {
			if IsValidCategory2(All_trends_string.Value) {
				nearby_points[i] = append(nearby_points[i], All_trends_string.Index)
			}
		}

		for _, All_trends_values := range date[start[i] : end[i]+1] {
			if IsValidCategory1(All_trends_values.Value) {
				nearby_points[i] = append(nearby_points[i], All_trends_values.Index)
			}

		}
	}
	fmt.Println("nearby_points1:", nearby_points[0])
	fmt.Println("nearby_points2:", nearby_points[1])
	//2
	//3-->adjust the points for evry start-end pair of points.
	for i := 0; i < count; i++ {
		for _, All_trends_string := range date[end[i]-2 : end[i]+1] {
			if IsValidCategory2(All_trends_string.Value) {
				fmt.Println("Foundpoint nearby the ending point 1:", All_trends_string.Index)
				t[i] = append(t[i], All_trends_string.Index)
			}
		}

		for _, All_trends_values := range date[end[i]-4 : end[i]+1] {
			if IsValidCategory1(All_trends_values.Value) {
				fmt.Println("Foundpoint nearby the ending point 2:", All_trends_values.Index)
				t1[i] = append(t1[i], All_trends_values.Index)
			}
		}
	}
	fmt.Println("nearby_points1:", nearby_points)
	fmt.Println("t:", t)
	fmt.Println("t1:", t1)

	//3
	//4->according to the number of couunts work this out ,the showing of variable and all condition.
	for i := 0; i < count; i++ {
		if len(t[i]) == 0 || len(t[i+1]) == 0 {
			fmt.Println("No nearby points found for:")
			fmt.Println("elemenating the first semicircle")
			start = start[1:]
			end = end[1:]
			break
			// fmt.Println(len(a), len(b))
		} else {
			fmt.Println("nearby points found for:")
			fmt.Println(len(t))
			end = end[1:]
			end = append(end, t1[1][1])
			break
		}
	}
	for i := 0; i < count; i++ {
		if len(t1[i]) == 0 || len(t1[i+1]) == 0 {
			fmt.Println(" no nearby points found for pair of points:=>")
			start = start[1:]
			end = end[1:]
			break
		} else {
			fmt.Println("nearby points found for:")
			fmt.Println(len(t1))
			end = end[1:]
			end = append(end, t1[1][1])
			break
		}
	}

	fmt.Println(len(end))

	count = len(start)
	fmt.Println(start, end, count)
	//4
	return start, end, count
}

//This is kind of like in operator in python language.
func IsValidCategory1(category int) bool {
	switch category {
	case
		795,
		796,
		797,
		798,
		799,
		800,
		801,
		802,
		803,
		804,
		805,
		806:
		return true
	}
	return false
}
func IsValidCategory2(category int) bool {
	switch category {
	case
		774,
		775,
		776,
		777,
		778,
		779,
		780,
		773,
		772,
		771,
		770,
		769,
		768,
		767:
		return true
	}
	return false
}

// var a1 []Generic.Date
// var b1 []Generic.Date

// for _, i := range Peaks {
// 	a1 = append(a1, l.Date_struct[i])
// }
// for _, i := range Troughs {
// 	b1 = append(b1, l.Date_struct[i])
// }
