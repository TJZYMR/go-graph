package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	//rand.Seed(int64(0))

	p := plot.New()

	p.Title.Text = "Graph Problem"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"First", randomPoints())

	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	//p.BackgroundColor.RGBA()
	//plotter.DefaultLineStyle.Width = vg.Points(2)
	// if pattern() == "M" {
	// 	plotutil.AddLinePoints(p, "Pattern M", linepoints())
	//3)this is for printing the pattern on to the screen after pattern matching.
	// }

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

//type XYs []XY

func linepoints() plotter.XYer {
	//pts := make([]plotter.XYer, 3)
	pts := make(plotter.XYs, 5)

	pts[0].X = 12
	pts[0].Y = 200
	pts[1].X = 14
	pts[1].Y = 1100
	pts[2].X = 17
	pts[2].Y = 200
	pts[3].X = 19
	pts[3].Y = 700
	pts[4].X = 21
	pts[4].Y = 210

	return pts
}

// randomPoints returns some random x, y points.
func randomPoints() plotter.XYs {
	pts := make(plotter.XYs, 17)
	pts[0].X = 11
	pts[1].X = 12
	pts[2].X = 13
	pts[3].X = 14
	pts[4].X = 15
	pts[5].X = 16
	pts[6].X = 17
	pts[7].X = 18
	pts[8].X = 19
	pts[9].X = 20
	pts[10].X = 21
	pts[11].X = 22
	pts[12].X = 23
	pts[13].X = 24
	pts[14].X = 25
	pts[15].X = 26
	pts[16].X = 27

	pts[0].Y = 697
	pts[1].Y = 703
	pts[2].Y = 692
	pts[3].Y = 710
	pts[4].Y = 736
	pts[5].Y = 745
	pts[6].Y = 736
	pts[7].Y = 756
	pts[8].Y = 785
	pts[9].Y = 811
	pts[10].Y = 795
	pts[11].Y = 785
	pts[12].Y = 787
	pts[13].Y = 801
	pts[14].Y = 838
	pts[15].Y = 826
	pts[16].Y = 801

	return pts
}

//err = plotutil.AddLinePoints(p,
//1)//try here drawing line between points with different colour(pattern 1 then put it in struct)
// l, _ := plotter.NewLine(linepoints())
// l.Color = plotutil.Color(2)
// l.LineStyle.Width = vg.Points(8)
// plotutil.AddLines(p, l)
// Save the plot to a PNG file.
