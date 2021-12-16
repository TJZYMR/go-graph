package pattern_up

//{2, 703}
type Date struct {
	Index int
	Date  string
	Value int
}

func P_up(pts []Date) (a bool) { //pts [][]int

	var b bool = false
	for i := range pts {
		if len(pts) != 1 {
			if pts[i].Value < pts[i+1].Value && pts[i+1].Date != pts[len(pts)-1].Date {
				continue
			} else if pts[i].Value < pts[i+1].Value && pts[i+1].Date == pts[len(pts)-1].Date {
				//fmt.Println("Down trend breaks here", pts[i+1])
				b = true
				break
			} else {
				b = false
				break
			}
		} else {
			b = true
			break
		}

	}
	return b
}

func P_down(pts []Date) (a bool) {
	var b bool = false
	for i := range pts {
		if pts[i].Value > pts[i+1].Value && pts[i+1].Date != pts[len(pts)-1].Date {
			continue
		} else if pts[i].Value > pts[i+1].Value && pts[i+1].Date == pts[len(pts)-1].Date {
			//fmt.Println("Down trend breaks here", pts[i+1])
			b = true
			break
		} else {
			b = false
			break
		}
	}
	return b
}

//return all peaks
//from that point previous point be lesser and next point be also lesser than itself.
func Peaks(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value < pts[i].Value && pts[i+1].Value < pts[i].Value {
			b = append(b, i)
			continue
		}
	}
	return b
}

//return all throughs
//from that point previous point be greater and next point be also greater than itself.
func Troughs(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value == pts[0].Value {
			continue
		} else if pts[i].Value == pts[len(pts)-1].Value {
			continue
		} else if pts[i-1].Value > pts[i].Value && pts[i+1].Value > pts[i].Value {
			b = append(b, i)
			continue
		}
	}
	return b
}

//0 {11,100}
