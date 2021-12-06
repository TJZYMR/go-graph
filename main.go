package main

import (
	"image/color"
	"y/W"
	"y/plotting_main"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	t := [][]int{{1, 697}, {2, 703}, {3, 692}, {4, 710}, {5, 736}, {6, 745}, {7, 736}, {8, 756}, {9, 785}, {10, 811}, {11, 795}, {12, 785}, {13, 787}, {14, 801}, {15, 838}, {16, 826}, {17, 801}, {18, 778}, {19, 778}, {20, 786}, {21, 781}, {22, 764}, {23, 773}, {24, 789}, {25, 773}, {26, 775}, {27, 792}, {28, 779}, {29, 780}, {30, 798}, {31, 807}, {32, 806}, {33, 802}, {34, 794}, {35, 779}, {36, 752}, {37, 739}, {38, 725}, {39, 728}, {40, 738}, {41, 731}, {42, 743}, {43, 745}, {44, 772}, {45, 765}, {46, 756}, {47, 755}, {48, 741}, {49, 738}, {50, 751}, {51, 770}, {52, 786}, {53, 783}, {54, 783}, {55, 773}, {56, 768}, {57, 774}, {58, 781}, {59, 796}, {60, 798}, {61, 781}}
	//t := [][]int{{12, 400}, {13, 567}, {14, 456}, {15, 654}, {16, 549}, {17, 711}, {18, 500}, {19, 444}, {20, 765}}
	a, b, count, _ := W.W(t)
	// fmt.Println(t[a : b+1])
	plot1(t, count, a, b)

}
func plot1(t [][]int, count int, a []int, b []int) {
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

	if err := p.Save(16*vg.Inch, 16*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
