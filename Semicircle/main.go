package main

import (
	"fmt"
	"image/color"
	"y/Generic"
	"y/pattern_up"
	"y/plotting_main"
	"y/utils"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Letter_semicircle struct {
	Date_struct []pattern_up.Date
}

func (l *Letter_semicircle) PatternU() ([]int, []int, int, bool) {
	peaks := peaks(l.Date_struct)
	troughs := troughs(l.Date_struct)
	fmt.Println("Peaks:", peaks)
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
		if pattern_up.P_up(l.Date_struct[i : i+2]) {
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
		if All_trends_string[i] == "Upward" && i != len(All_trends_string)-1 {
			if All_trends_string[i+1] == "Downward" && i+1 != len(All_trends_string)-1 { //Task:=>check here for first and last point to be in the align manner.
				var a1 []int //end point
				var x int

				for i1 := 0; i1 < 40; i1++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value + i1
					a1 = append(a1, x)
				}
				for i2 := 0; i2 < 40; i2++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value - i2
					a1 = append(a1, x)
				}
				var a2 []int //start point
				var x1 int

				for i1 := 0; i1 < 40; i1++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value + i1
					a2 = append(a2, x1)
				}
				for i2 := 0; i2 < 40; i2++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value - i2
					a2 = append(a2, x1)
				}
				// if Is_Valid11(l.Date_struct[All_trends_index_values[i]].Value, a1) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i])
				// 	count++
				// 	bol = true //do tuing here
				// } else
				if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value, a1) {
					start = append(start, All_trends_index_values[i]-1)
					end = append(end, All_trends_index_values[i+1]+2) //last U
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+2].Value, a1) {
					start = append(start, All_trends_index_values[i]+2)
					end = append(end, All_trends_index_values[i+1]+2) //last U
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+3].Value, a1) {
					start = append(start, All_trends_index_values[i]+3)
					end = append(end, All_trends_index_values[i+1]+2) //last U
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i+1]].Value, a2) {
					start = append(start, All_trends_index_values[i])
					end = append(end, All_trends_index_values[i+1]+3) //third u
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value, a2) {
					start = append(start, All_trends_index_values[i]+1)
					end = append(end, All_trends_index_values[i+2]+1) //first,second and fourth U.
					count++
					bol = true
				}
				// } else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+2].Value, a1) {
				// 	start = append(start, All_trends_index_values[i]+2)
				// 	end = append(end, All_trends_index_values[i+1]+2) //last U
				// 	count++
				// 	bol = true
				// } else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+3].Value, a1) {
				// 	start = append(start, All_trends_index_values[i]+3)
				// 	end = append(end, All_trends_index_values[i+1]+2) //last U
				// 	count++
				// 	bol = true
				// }
				// } else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+2].Value, a1) {
				// 	start = append(start, All_trends_index_values[i]+2)
				// 	end = append(end, All_trends_index_values[i+2]+2)
				// 	count++
				// 	bol = true
				// } else if Is_Valid11(l.Date_struct[All_trends_index_values[i]].Value, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i]+4)
				// 	count++
				// 	bol = true
				// } else if Is_Valid11(l.Date_struct[All_trends_index_values[i+1]].Value, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	fmt.Println("here!!")
				// 	end = append(end, All_trends_index_values[i+1]+2) //some connection
				// 	count++
				// 	bol = true
				// } else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i+1]+7) //last U
				// 	count++
				// 	bol = true
				// }
				// } else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	fmt.Println("here1!!")
				// 	end = append(end, All_trends_index_values[i]+8) //some connection
				// 	count++
				// 	bol = true
				// }
			}
		}

	}

	// fmt.Println("Count =", count)
	return start, end, count, bol
}

func peaks(pts []pattern_up.Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value < pts[i].Value && pts[i+1].Value < pts[i].Value && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {

			var a1 []int
			var x int
			for i1 := 0; i1 < 12; i1++ {
				x = pts[i].Value + i1
				a1 = append(a1, x)
			}
			for i2 := 0; i2 < 12; i2++ {
				x = pts[i].Value - i2
				a1 = append(a1, x)
			}
			if Is_Valid1(pts[i-1].Value, a1) && pts[i-1].Value > pts[i-2].Value { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue
			}
			if Is_Valid1(pts[i+1].Value, a1) && pts[i+1].Value > pts[i+2].Value { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
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
		n[79]:
		return true
	}
	return false
}

func troughs(pts []pattern_up.Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value > pts[i].Value && pts[i+1].Value > pts[i].Value && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {

			b = append(b, i)
			continue

		}

	}
	return b

}
func main() {
	Date := Generic.Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/data1.csv")
	Plot_Original(Date, "Main_Without_Patterns.png")
	P := &Letter_semicircle{Date_struct: Date}
	start, end, count, bol := Find(P)
	fmt.Println(start, end, count, bol)
	if bol {
		Plot_Pattern(Date, count, start, end, "Main_With_Patterns_U.png")
		fmt.Println(Date[28], Date[35])
	} else {
		fmt.Println("No Pattern Found")
	}
}

func Plot_Pattern(t []pattern_up.Date, count int, a []int, b []int, Name string) {
	//rand.Seed(int64(0))

	p := plot.New()

	p.Title.Text = Name
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", plotting_main.RandomPoints(t))

	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	//p.BackgroundColor.RGBA()
	plotter.DefaultLineStyle.Width = vg.Points(2)

	for i := 0; i < count; i++ {
		plotter.DefaultLineStyle.Color = color.RGBA{B: 255, A: 255}
		plotter.DefaultLineStyle.Width = vg.Points(3.5)

		plotutil.AddLinePoints(p, "Pattern M", plotting_main.Linepoints(t[a[i]:b[i]]))

	}

	if err := p.Save(18*vg.Inch, 18*vg.Inch, Name); err != nil {
		panic(err)
	}
}
func Plot_Original(Date []pattern_up.Date, Name string) {
	p := plot.New()

	p.Title.Text = "Graph Problem"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", plotting_main.RandomPoints(Date))

	if err != nil {
		panic(err)
	}

	if err := p.Save(18*vg.Inch, 18*vg.Inch, Name); err != nil {
		panic(err)
	}
}

type Patterns interface {
	PatternU() ([]int, []int, int, bool)
}

func Find(P Patterns) ([]int, []int, int, bool) {
	start, end, count, bol := P.PatternU()
	return start, end, count, bol
} //end of Find and with this adter that dfgfddfdfdf
