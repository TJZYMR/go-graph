package plotting_main

import "gonum.org/v1/plot/plotter"

type XY struct{ X, Y float64 }

func Linepoints(x [][]int) plotter.XYer {
	//pts := make([]plotter.XYer, 3)
	pts := make(plotter.XYs, len(x))

	for i := range x {
		pts[i].X = float64(x[i][0])
		pts[i].Y = float64(x[i][1])
	}
	return pts
}
func RandomPoints(x [][]int) plotter.XYs {
	pts := make(plotter.XYs, len(x))

	for i := range x {
		pts[i].X = float64(x[i][0])
		pts[i].Y = float64(x[i][1])
	}
	return pts
}
