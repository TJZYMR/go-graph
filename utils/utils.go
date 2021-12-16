package utils

import (
	"fmt"
	"image/color"
	"strconv"
	"y/pattern_up"
	"y/plotting_main"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Findconsquetive(m map[int]int) (keys []int) {
	for k, _ := range m {
		if _, ok := m[k+1]; ok {
			keys = append(keys, k)
		}
	}
	return
}
func Sort(m []int) (keys []int) {
	for i := range m {
		for j := i + 1; j < len(m); j++ {
			if m[i] > m[j] {
				m[i], m[j] = m[j], m[i]
			}
		}
	}
	return m
}
func Findtrendw() (up1, down1 map[int]int) {
	up := make(map[int]int)
	down := make(map[int]int)
	points := [][]int{{11, 100}, {12, 200}, {13, 700}, {14, 1100}, {15, 800}, {16, 500}, {17, 200}, {18, 400}, {19, 700}, {20, 350}, {21, 210}, {22, 310}, {23, 400}, {24, 450}, {25, 500}, {26, 800}, {27, 1000}}
	fmt.Println(points) //x,y pairs
	for i := range points {
		if points[i][1] < points[i+1][1] {
			if points[i+1][1] != points[len(points)-1][1] {
				fmt.Println("Updward Trend from: " + strconv.Itoa(points[i][1]) + " to " + strconv.Itoa(points[i+1][1]))
				//fmt.Println(points[i])
				up[len(points[:i])] = points[i][1]

			} else if points[i+1][1] == points[len(points)-1][1] {
				fmt.Println("Updward Trend from: " + strconv.Itoa(points[i][1]) + " to " + strconv.Itoa(points[i+1][1]))
				//fmt.Println(points[i])
				up[len(points[:i+1])] = points[i+1][1]
				//neglecting the last point if otherwise want reapete the above code.
				break
			}
		} else {
			fmt.Println("Downward Trend from: " + strconv.Itoa(points[i][1]) + " to " + strconv.Itoa(points[i+1][1]))
			down[len(points[:i])] = points[i][1]

		}
		//for w pattern woulbe be too points up ,two down and one in between.

	}
	return up, down
}
func Findtrendm() (up1, down1 map[int]int) {
	up := make(map[int]int)
	down := make(map[int]int)
	points := [][]int{{11, 100}, {12, 200}, {13, 700}, {14, 1100}, {15, 800}, {16, 500}, {17, 200}, {18, 400}, {19, 700}, {20, 350}, {21, 210}, {22, 310}, {23, 400}, {24, 450}, {25, 500}, {26, 800}, {27, 1000}}
	fmt.Println(points) //x,y pairs
	for i := range points {
		if points[i][1] < points[i+1][1] {
			if points[i+1][1] != points[len(points)-1][1] {
				fmt.Println("Updward Trend from: " + strconv.Itoa(points[i][1]) + " to " + strconv.Itoa(points[i+1][1]))
				//fmt.Println(points[i])
				up[len(points[:i])] = points[i][1]

			} else {
				break //neglecting the last point if otherwise want reapete the above code.
			}
		} else {
			fmt.Println("Downward Trend from: " + strconv.Itoa(points[i][1]) + " to " + strconv.Itoa(points[i+1][1]))
			down[len(points[:i])] = points[i][1]

		}
		//for w pattern woulbe be too points up ,two down and one in between.

	}
	return up, down
}
func Plot1(t []pattern_up.Date, count int, a []int, b []int) {
	//rand.Seed(int64(0))

	p := plot.New()

	p.Title.Text = "Graph Problem"
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
		plotter.DefaultLineStyle.Width = vg.Points(2.5)

		plotutil.AddLinePoints(p, "Pattern M", plotting_main.Linepoints(t[a[i]:b[i]]))

	}

	// if pattern() == "M" {
	// 	plotutil.AddLinePoints(p, "Pattern M", linepoints())
	//3)this is for printing the pattern on to the screen after pattern matching.
	// }

	if err := p.Save(18*vg.Inch, 18*vg.Inch, "points1.png"); err != nil {
		panic(err)
	}
}
