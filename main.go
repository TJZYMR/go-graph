package main

import (
	"y/Generic"
	"y/Semicircle"
	"y/Structs_for_patterns"
	"y/utils"
)

//main function for mathcing semicircle.
func main() {
	date := Generic.Genericcsv("Generic/s1.csv")
	pattern := []string{"Upward", "Downward"}
	Semicircle_data_struct := &Semicircle.Lettersemicricle{date, pattern}
	Start1, End1, Count1, _ := Structs_for_patterns.Find(Semicircle_data_struct) //start,end,count,bool
	Start, End, Count := Semicircle.Adjustpoints(Start1, End1, Count1, date)     //Adjusting the output to meet the semicircle atleast requirements.
	utils.Plot1(date, Count, Start, End)

}

//1.sending the data so that it will be in searlized form for graph plotting and the pattern matches then

//2.we will send the data to the graph plotting function
// randomPoints returns some random x, y points.
// func randomPoints() plotter.XYs {
// 	pts := make(plotter.XYs, 1)
// 	pts[0].X = "11 - nov - 2021"
// 	pts[0].Y = 1

// 	return pts
// }
// bol := P_up(date)
// date1 := []Date{{1, "11-nov-2021", 200}, {2, "12-nov-2021", 200}, {3, "13-nov-2021", 300}, {4, "14-nov-2021", 400}, {5, "15-nov-2021", 500}, {6, "16-nov-2021", 600}, {7, "17-nov-2021", 700}}
// p := []string{"Upward", "Downward", "Upward", "Downward"}
// m := &M.Letterm{date, p}
// // a, b, c, _ := m.Pattern()
// a, b, c, _ := Structs_for_patterns.Find(m)
// fmt.Println(a, b, c)
// utils.Plot1(date, c, a, b)
// Get value from cell by given worksheet name and axis.
// func main() {
// 	// t := [][]int{{1, 697}, {2, 703}, {3, 692}, {4, 710}, {5, 736}, {6, 745}, {7, 736}, {8, 756}, {9, 785}, {10, 811}, {11, 795}, {12, 785}, {13, 787}, {14, 801}, {15, 838}, {16, 826}, {17, 801}, {18, 778}, {19, 778}, {20, 786}, {21, 781}, {22, 764}, {23, 773}, {24, 789}, {25, 773}, {26, 775}, {27, 792}, {28, 779}, {29, 780}, {30, 798}, {31, 807}, {32, 806}, {33, 802}, {34, 794}, {35, 779}, {36, 752}, {37, 739}, {38, 725}, {39, 728}, {40, 738}, {41, 731}, {42, 743}, {43, 745}, {44, 772}, {45, 765}, {46, 756}, {47, 755}, {48, 741}, {49, 738}, {50, 751}, {51, 770}, {52, 786}, {53, 783}, {54, 783}, {55, 773}, {56, 768}, {57, 774}, {58, 781}, {59, 796}, {60, 798}, {61, 781}}
// 	// //var m Patterns
// 	// p := []string{"Upward", "Downward", "Upward", "Downward"}
// 	// m := &M.Letterm{t, p}
// 	// // a, b, c, _ := m.Pattern()
// 	// a, b, c, _ := Structs_for_patterns.Find(m)
// 	// utils.Plot1(t, c, a, b)
// 	//w
// 	// p1 := []string{"Downward", "Upward", "Downward", "Upward"}
// 	// w := &m_struct.Letterw{t, p1}
// 	// // a1, b1, c1, _ := w.Pattern()
// 	// a1, b1, c1, _ := m_struct.Find(w)
// 	// utils.Plot1(t, c1, a1, b1)

// 	//1.finding file type
// 	//now we have file name and we wamt to find the file type
// 	// var filetype string
// 	// var filepath string

// 	path := "/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/s1.csv"
// 	var data []Generic.Date
// 	fileExtension := filepath.Ext(path)
// 	fmt.Println(fileExtension)
// 	if fileExtension == ".xlsx" {
// 		data = Generic.Genericexcel(path, "date-values")
// 		fmt.Println("data received:=>", "File Type:=>Excel   ", data)
// 	} else if fileExtension == ".csv" {
// 		data = Generic.Genericcsv(path)
// 		fmt.Println("data received:=>", "File Type:=>csv   ", data)

// 	} else if fileExtension == ".json" {
// 		data = Generic.Genericjson(path)
// 		fmt.Println("data received:=>", "File Type:=>json   ", data)

// 	} else if fileExtension == ".txt" {
// 		data = Generic.Genericexcel(path, "date-values")
// 		fmt.Println("data received:=>", "File Type:=>txt   ", data)

// 	} else if fileExtension == ".xml" {
// 		data = Generic.Genericexcel(path, "date-values")
// 		fmt.Println("data received:=>", "File Type:=>xml   ", data)

// 	} else if fileExtension == ".html" {
// 		data = Generic.Genericexcel(path, "date-values")
// 		fmt.Println("data received:=>", "File Type:=>html   ", data)

// 	} else if fileExtension == ".xls" {
// 		data = Generic.Genericexcel(path, "date-values")
// 		fmt.Println("data received:=>", "File Type:=>xls   ", data)

// 	} else { //here database is the file type
// 		data := Generic.Genericdatabase("127.0.0.1:3306", "tatva", "tatva", "zymr@123")
// 		fmt.Println("data received:=>", data)

// 	}
// 	//first set t to the accepted data supported by recieving function.
// 	//second draw on graph with the help of the recieved data which is in string form.
// 	p := []string{"Upward", "Downward", "Upward", "Downward"}
// 	m := &M.Letterm{t, p}
// 	// a, b, c, _ := m.Pattern()
// 	a, b, c, _ := Structs_for_patterns.Find(m)
// 	utils.Plot1(t, c, a, b)
// }
// func main() {

// 	f, err := os.Open("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/date-values.csv")
// 	if err != nil {
// 		fmt.Println(err)

// 	}
// 	var date []pattern_up.Date
// 	if err != nil {
// 		fmt.Println(err)

// 	}
// 	csvReader := csv.NewReader(f)
// 	records, _ := csvReader.ReadAll()
// 	records = records[1:]
// 	var Index int
// 	for _, row := range records {
// 		Index = Index + 1
// 		b, _ := strconv.Atoi(row[1])
// 		date = append(date, pattern_up.Date{Index, row[0], b})
// 	}

// 	p1 := []string{"Downward", "Upward", "Downward", "Upward"}
// 	w := &W.Letterw{date, p1}
// 	a2, b2, c2, _ := Structs_for_patterns.Find(w)
// 	fmt.Println(a2, b2, c2)
// 	utils.Plot1(date, c2, a2, b2)
// }
