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
		//if nov then 11 like that putting in the month

		// if x1.Date[3:6] == "Sep" {
		// 	d = 9
		// }
		// if x1.Date[3:6] == "Oct" {
		// 	d = 10
		// }
		// if x1.Date[3:6] == "Nov" {
		// 	d = 11
		// }
		// if x1.Date[3:6] == "Dec" {
		// 	d = 12
		// }

		// c := x1.Date[0:2] //+ x[i].Date[7:11]
		// // a, _ := strconv.Atoi(c) //+ x[i].Date[3:5] + x[i].Date[7:11]
		// d1 := strconv.Itoa(d)
		// // result := fmt.Sprintf("%s%s%s", c, d1, x1.Date[9:11])
		// result := fmt.Sprintf("%s%s", c, d1)
		// b, _ := strconv.ParseFloat(result, 64)

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
