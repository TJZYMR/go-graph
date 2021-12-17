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

//1.changing the method to find peaks and troughs.
//2.after matching he pattern upward and then downward
//3.then matching starting and ending point's value are same or euivalent and if yes then declaring the points as semicircle.
func (l *Lettersemicricle) Pattern() ([]int, []int, int, bool) {
	a := Peaks(l.T)
	b := Troughs(l.T)
	fmt.Println("Peaks", a)
	fmt.Println("Troughs", b)
	var a1 []pattern_up.Date
	var b1 []pattern_up.Date

	for _, i := range a {
		a1 = append(a1, l.T[i])
	}
	for _, i := range b {
		b1 = append(b1, l.T[i])
	}

	var i1 []string
	var w1 []int
	for i := range a {
		w1 = append(w1, a[i])
	}
	for i := range b {
		w1 = append(w1, b[i])
	}
	w2 := utils.Sort(w1) //mixed values
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
	for _, j := range w2 {
		fmt.Println("Mix =", l.T[j])
	}
	fmt.Println("main Pattern =", l.P)
	fmt.Println("Pattern =", i1)
	var start []int
	var end []int
	var count int
	for i := range i1 {
		if i1[i] == "Upward" && i != len(i1)-1 {

			if i1[i+1] == "Downward" && i+1 != len(i1)-1 {

				count = count + 1
				// fmt.Println("Count =", count)
				// fmt.Println("Starting Point =", l.T[i2[i]])
				// fmt.Println("Ending Point =", l.T[i2[i+1]+1])

				start = append(start, i2[i])
				end = append(end, i2[i+1+1])
				bol = true

			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}
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
func main() {
	date := Getdata()
	p1 := []string{"Upward", "Downward"}
	w := &Lettersemicricle{date, p1}
	a, b, c, _ := Structs_for_patterns.Find(w) //start,end,count,bool
	// fmt.Println(b[0])

	a1, b1, c1 := Adjustpoints(a, b, c, date)
	utils.Plot1(date, c1, a1, b1)

}
func Adjustpoints(a, b []int, c int, date []pattern_up.Date) ([]int, []int, int) { //
	//1->make these variables as many times as count variable c.

	nearby_points := make([][]int, c)
	//var nearby_points2 []int
	t := make([][]int, c)
	t1 := make([][]int, c)

	// var t1 [][]int

	//1
	//2->make automatic isvalidcategory function and loop thorough the points until the last point.
	for i := 0; i < c; i++ {
		for _, i1 := range date[a[i] : b[i]+1] {
			if IsValidCategory2(i1.Value) {
				nearby_points[i] = append(nearby_points[i], i1.Index)
			}
		}

		for _, i2 := range date[a[i] : b[i]+1] {
			if IsValidCategory1(i2.Value) {
				nearby_points[i] = append(nearby_points[i], i2.Index)
			}

		}
	}
	fmt.Println("nearby_points1:", nearby_points[0])
	fmt.Println("nearby_points2:", nearby_points[1])
	//2
	//3-->adjust the points for evry start-end pair of points.
	for i := 0; i < c; i++ {
		for _, i1 := range date[b[i]-2 : b[i]+1] {
			if IsValidCategory2(i1.Value) {
				fmt.Println("Foundpoint nearby the ending point 1:", i1.Index)
				t[i] = append(t[i], i1.Index)
			}
		}

		for _, i2 := range date[b[i]-4 : b[i]+1] {
			if IsValidCategory1(i2.Value) {
				fmt.Println("Foundpoint nearby the ending point 2:", i2.Index)
				t1[i] = append(t1[i], i2.Index)
			}
		}
	}
	fmt.Println("nearby_points1:", nearby_points)
	fmt.Println("t:", t)
	fmt.Println("t1:", t1)

	//3
	//4->according to the number of couunts work this out ,the showing of variable and all condition.
	for i := 0; i < c; i++ {
		if len(t[i]) == 0 || len(t[i+1]) == 0 {
			fmt.Println("No nearby points found for:")
			fmt.Println("elemenating the first semicircle")
			a = a[1:]
			b = b[1:]
			break
			// fmt.Println(len(a), len(b))
		} else {
			fmt.Println("nearby points found for:")
			fmt.Println(len(t))
			b = b[1:]
			b = append(b, t1[1][1])
			break
		}
	}
	for i := 0; i < c; i++ {
		if len(t1[i]) == 0 || len(t1[i+1]) == 0 {
			fmt.Println(" no nearby points found for pair of points:=>")
			a = a[1:]
			b = b[1:]
			break
		} else {
			fmt.Println("nearby points found for:")
			fmt.Println(len(t1))
			b = b[1:]
			b = append(b, t1[1][1])
			break
		}
	}

	fmt.Println(len(b))

	c = len(a)
	fmt.Println(a, b, c)
	//4
	return a, b, c
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
