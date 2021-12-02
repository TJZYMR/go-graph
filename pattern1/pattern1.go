package pattern1

import (
	"fmt"
	"strconv"

	"y/plotting_main"
	"y/utils"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func W() {
	up, down := utils.Findtrendw()
	fmt.Println("Updward: ", up)
	fmt.Println("downward: ", down)
	pts := [][]int{{11, 100}, {12, 200}, {13, 700}, {14, 1100}, {15, 800}, {16, 500}, {17, 200}, {18, 400}, {19, 700}, {20, 350}, {21, 200}, {22, 310}, {23, 400}, {24, 450}, {25, 500}, {26, 800}, {27, 1000}}
	fmt.Println("Points: ", pts)
	draw := patternW(pts, up, down) //1)return array of inbetwwen points from this funcion
	fmt.Println("Draw: ", draw)
	p := plot.New()

	p.Title.Text = "Graph Problem"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", plotting_main.RandomPoints(pts))

	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	plotter.DefaultLineStyle.Width = vg.Points(2)
	plotutil.AddLinePoints(p, "Pattern W", plotting_main.Linepoints(draw))
	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "W.png"); err != nil {
		panic(err)
	} //plotting finished For M	Pattern

}

func patternW(pts [][]int, up, down map[int]int) (d [][]int) {

	samepoints := make(map[int]int) //same line points on y axis.
	for index, _ := range pts {
		if pts[index][1] == 200 {
			samepoints[index] = pts[index][1]
		} else {
			continue
		}
	}
	fmt.Println("SamePoints: ", samepoints)
	//finding first pattern from where from there downtrend starts.
	//var draw [][]int
	a1 := utils.Findconsquetive(down)
	b_down := utils.Sort(a1)
	for index, i := range b_down {
		b_down[index] = i + 1
	}
	fmt.Println("consequetive down: ", b_down)
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
	//2)find from first point upward trend and then downward trend.
	//first downward trend finding
	start := b_down[0] - 1
	// fmt.Println("start: ", start[])
	//3)checking consequetive downward and upward trend
	peak1 := b_down[0] - 1
	peak2 := b_down[2] - 1 //set it afterwards
	//peak2 := pts[b_down[1]]
	for _, i := range b_down {
		if i >= 3 && i <= 6 { ///find how we can get key from that value.//these are indexes of values.6 is from samepoints
			if i-1 == peak1 {
				fmt.Println(strconv.Itoa(i-1) + "and" + strconv.Itoa(i) + "is in consequetive downward trend between 2 and 3 points")
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 2 and 3 points")

			} else {
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 2 and 3 points")

			}
		} else {
			continue
		}
	}
	//now upward trend
	for _, i := range b_up {
		if i >= 6 && i <= 10 { ///find how we can get key from that value.//
			if i-1 == 6 {
				fmt.Println(strconv.Itoa(i-1) + "and" + strconv.Itoa(i) + "is in consequetive upward trend between 3 and 4 points")
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive upward trend between 3 and 4 points")

			} else {
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive upward trend between 3 and 4 points")

			}
		} else {
			continue
		}
	}
	for _, i := range b_down {
		if i >= 8 && i <= 10 { ///find how we can get key from that value.//these are indexes of values.6 is from samepoints
			if i-1 == peak2 {
				fmt.Println(strconv.Itoa(i-1) + "and" + strconv.Itoa(i) + "is in consequetive downward trend between 4 and 5 points")
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 4 and 5 points")

			} else {
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive downward trend between 4 and 5 points")

			}
		} else {
			continue
		}
	}
	for _, i := range b_up {
		if i >= 10 { ///find how we can get key from that value.//
			if i-1 == 10 {
				fmt.Println(strconv.Itoa(i-1) + "and" + strconv.Itoa(i) + "is in consequetive upward trend between 5 and 6 points")
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive upward trend between 5 and 6 points")

			} else {
				fmt.Println(strconv.Itoa(i) + "and" + strconv.Itoa(i+1) + "is in consequetive upward trend between 5 and 6 points")

			}
		} else {
			continue
		}
	}
	fmt.Println("Hence we can say that W shape is formed")
	var ar [][]int
	for i, j := range samepoints {
		ar = append(ar, []int{i, j})
	}
	var draw [][]int
	//fmt.Println(ar)
	//finding downward trend between 1 and 2 down points
	var e []int
	for j, _ := range up {
		e = append(e, j)

	}
	e1 := utils.Sort(e)
	fmt.Println(e1)
	fmt.Println("Points' Indexes are", start, e1[len(e1)-1])
	fmt.Println("Points' Values are", pts[start], pts[e1[len(e1)-1]])

	for i := range pts {
		if i >= start && e1[len(e1)-1] >= i {
			draw = append(draw, pts[i])
		}
	}
	return draw

}

//conditions will be that(fullfill all the conditions) :1) first find the first downward trend then
//2) find the just after that upward trend from the point in samepoints
//3) find the just after that downward trend with last point from the samepoints.
