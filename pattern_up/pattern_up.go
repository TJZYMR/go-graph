package pattern_up

//{2, 703}
func P_up(pts [][]int) (a bool) {
	//t := [][]int{{1, 697}, {2, 703}, {3, 692}, {4, 710}, {5, 736}, {6, 745}, {7, 736}, {8, 756}, {9, 785}, {10, 811}, {11, 795}, {12, 785}, {13, 787}, {14, 801}, {15, 838}}

	var b bool = false
	for i := range pts {
		if len(pts) != 1 {
			if pts[i][1] < pts[i+1][1] && pts[i+1][0] != pts[len(pts)-1][0] {
				continue
			} else if pts[i][1] < pts[i+1][1] && pts[i+1][0] == pts[len(pts)-1][0] {
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

func P_down(pts [][]int) (a bool) {
	var b bool = false
	for i := range pts {
		if pts[i][1] > pts[i+1][1] && pts[i+1][0] != pts[len(pts)-1][0] {
			continue
		} else if pts[i][1] > pts[i+1][1] && pts[i+1][0] == pts[len(pts)-1][0] {
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
func Peaks(pts [][]int) (a []int) {
	var b []int
	for i := range pts {
		if pts[i][1] == pts[0][1] {
			continue
		} else if pts[i][1] == pts[len(pts)-1][1] {
			continue
		} else if pts[i-1][1] < pts[i][1] && pts[i+1][1] < pts[i][1] {
			b = append(b, i)
			continue
		}
	}
	return b
}

//return all throughs
//from that point previous point be greater and next point be also greater than itself.
func Troughs(pts [][]int) (a []int) {
	var b []int
	for i := range pts {
		if pts[i][1] == pts[0][1] {
			continue
		} else if pts[i][1] == pts[len(pts)-1][1] {
			continue
		} else if pts[i-1][1] > pts[i][1] && pts[i+1][1] > pts[i][1] {
			b = append(b, i)
			continue
		}
	}
	return b
}

//0 {11,100}
