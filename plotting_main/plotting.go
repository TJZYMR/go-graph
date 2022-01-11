package plotting_main

import (
	"y/pattern_up"

	"gonum.org/v1/plot/plotter"
)

type XY struct{ X, Y float64 }

func Linepoints(x []pattern_up.Date) plotter.XYer {
	//pts := make([]plotter.XYer, 3)
	pts := make(plotter.XYs, len(x))
	// var d int
	for i, x1 := range x {

		pts[i].X = float64(x1.Index)
		pts[i].Y = float64(x[i].Value)
		// fmt.Println(pts[i].X, pts[i].Y)

	}
	return pts
}
func RandomPoints(x []pattern_up.Date) plotter.XYs {
	pts := make(plotter.XYs, len(x))

	for i := range x {
		// b, _ := strconv.Atoi(x[i].Date[0:2] + "." + x[i].Date[3:5] + "." + x[i].Date[7:11])

		pts[i].X = float64(x[i].Index)
		pts[i].Y = float64(x[i].Value)
	}
	return pts
}
