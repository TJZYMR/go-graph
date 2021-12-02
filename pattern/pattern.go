package pattern

import (
	"fmt"
	"strconv"
	"y/utils"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// type pattern interface {
// 	//draw(dc *gg.Context)
// 	//w()//w pattern
// 	//m()//m pattern
// }

func M() {
	up, down := findtrend()
	fmt.Println(up)
	fmt.Println(down)
	//pattern finding for M starts........
	//1)find three same y points
	pts := [][]int{{11, 100}, {12, 200}, {13, 700}, {14, 1100}, {15, 800}, {16, 500}, {17, 200}, {18, 400}, {19, 700}, {20, 350}, {21, 200}, {22, 310}, {23, 400}, {24, 450}, {25, 500}, {26, 800}, {27, 1000}}
	draw := patternM(pts, up, down) //1)return array of inbetwwen points from this funcion
	fmt.Println(draw)
	//2)plotting from returened points from above functions
	p := plot.New()

	p.Title.Text = "Graph Problem"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", randomPoints(pts))

	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	plotter.DefaultLineStyle.Width = vg.Points(2)
	plotutil.AddLinePoints(p, "Pattern M", linepoints(draw))
	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "M.png"); err != nil {
		panic(err)
	} //plotting finished For M	Pattern

	//pattern finding for W starts........
}

// 	}

//2.1)find upward trend opto some point, mark the point from where there are no two consequetive keys.
//task:checking for consecutive keys.type XYs []XY
type XY struct{ X, Y float64 }

func linepoints(x [][]int) plotter.XYer {
	//pts := make([]plotter.XYer, 3)
	pts := make(plotter.XYs, len(x))

	for i := range x {
		pts[i].X = float64(x[i][0])
		pts[i].Y = float64(x[i][1])
	}
	return pts
}
func randomPoints(x [][]int) plotter.XYs {
	pts := make(plotter.XYs, len(x))

	for i := range x {
		pts[i].X = float64(x[i][0])
		pts[i].Y = float64(x[i][1])
	}
	return pts
}
func findtrend() (up1, down1 map[int]int) {
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
func patternM(pts [][]int, up, down map[int]int) (p [][]int) {

	samepoints := make(map[int]int)
	for index, _ := range pts {
		if pts[index][1] == 200 {
			samepoints[index] = pts[index][1]
		} else {
			continue
		}
	}
	fmt.Println(samepoints)
	//2)find from first point upward trend and then downward trend.
	var draw [][]int
	a := utils.Findconsquetive(up)
	b_up := utils.Sort(a)

	for index, i := range b_up {
		b_up[index] = i + 1
	}
	fmt.Println(b_up)
	for _, v := range b_up {
		fmt.Print(pts[v][1])
		fmt.Print(" ,")
	}

	fmt.Println("\n")
	a1 := utils.Findconsquetive(down)
	b_down := utils.Sort(a1)
	for index, i := range b_down {
		b_down[index] = i + 1
	}
	fmt.Println(b_down)

	for _, v := range b_down {
		fmt.Print(pts[v][1])
		fmt.Print(" ,")

	} //for peak,if from b_up,b_down:peak point will not be on the both array,that will be peak;back to the last samepoints.write code afterwards
	peak1 := b_down[0] - 1
	peak2 := b_down[2] - 1
	fmt.Println("\n")
	//b_up is the upward trend and b_down is the downward trend.
	//finding upward trend between 1 and 2 down points
	for _, i := range b_up {
		if i <= 6 { ///find how we can get key from that value.//
			fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive upward trend between 1 and 2 points")
		} else {
			break
		}
	}
	//finding downward trend between 1 and 2 down points

	for _, i := range b_down {
		if i <= 6 { ///find how we can get key from that value.//
			if i-1 == peak1 {
				fmt.Println(strconv.Itoa(i-1) + "and" + strconv.Itoa(i) + "is in consequetive downward trend between 1 and 2 points")
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 1 and 2 points")

			} else {
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 1 and 2 points")

			}
		} else {
			continue
		}
	}
	fmt.Println(peak1, peak2, "both are peakes of M shape")
	//finding upward trend between 2 and 3 down points
	for _, i := range b_up {
		if i >= 6 && i <= 10 { ///find how we can get key from that value.//
			if i-1 == 6 {
				fmt.Println(strconv.Itoa(i-1) + "and" + strconv.Itoa(i) + "is in consequetive upward trend between 2 and 3 points")
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive upward trend between 2 and 3 points")

			} else {
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive upward trend between 2 and 3 points")

			}
		} else {
			continue
		}
	}
	for _, i := range b_down {
		if i >= 6 && i <= 10 { ///find how we can get key from that value.//
			if i-1 == peak2 {
				fmt.Println(strconv.Itoa(i-1) + "and" + strconv.Itoa(i) + "is in consequetive downward trend between 2 and 3 points")
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 2 and 3 points")

			} else {
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 2 and 3 points")

			}
		} else {
			continue
		}
	}
	fmt.Println("Hence we can say that M shape is formed")
	var start int
	var end int
	var ar [][]int
	for i, j := range samepoints {
		ar = append(ar, []int{i, j})
	}
	//fmt.Println(ar)
	//finding downward trend between 1 and 2 down points
	start = ar[0][0]
	end = ar[len(ar)-1][0]
	fmt.Println("Points' Indexes are", start, end)
	fmt.Println("Points' Values are", pts[start], pts[end])
	fmt.Println("Peaks are", peak1, ":", pts[peak1], peak2, ":", pts[peak2])

	for i := range pts {
		if i >= start && end >= i {
			draw = append(draw, pts[i])

		}

	}
	return draw

}
