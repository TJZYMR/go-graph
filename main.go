package main

import (
	"fmt"
	"y/Generic"
	"y/Semicircle"
	"y/Structs_for_patterns"
)

// func main() {
// 	colorRed := "\033[31m"
// 	colorGreen := "\033[32m"
// 	colorYellow := "\033[33m"
// 	colorBlue := "\033[34m"
// 	//semicircle and inverted ones are working properly.just to make M and w pattern good
// 	//Task:=>Code for Inverted and semicircle is working for four files but not for s1.csv file.so fix that and tell mentor.
// 	Date := Generic.Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/date4.csv")
// 	Semicircle.Plot_Original(Date, "Original.png")
// 	P := &Inverted_semi_circle.Letter_inverted_semicircle{Date}
// 	start, end, count, bol := Inverted_semi_circle.Find(P)
// 	if bol {
// 		fmt.Println(string(colorRed), "Inverted_Semicircle/s found:=>", count)
// 		fmt.Println("Start: ", start, "End: ", end, "Count: ", count)

// 	} else {
// 		fmt.Println(string(colorRed), "Inverted /s not found")
// 	}

// 	P1 := &Semicircle.Letter_semicircle{Date}
// 	start1, end1, count1, bol1 := Semicircle.Find(P1)
// 	if bol1 {
// 		fmt.Println(string(colorGreen), "Semicircle/s found:=>", count1)
// 		fmt.Println("Start: ", start1, "End: ", end1, "Count: ", count1)
// 	} else {
// 		fmt.Println(string(colorGreen), "Semicircle /s not found")
// 	}

// 	P2 := &M.Letterm{Date}
// 	start2, end2, count2, bol2 := M.Find(P2)
// 	if bol2 {
// 		fmt.Println(string(colorYellow), "M /s found:=>", count2)
// 		fmt.Println("Start: ", start2, "End: ", end2, "Count: ", count2)
// 	} else {
// 		fmt.Println(string(colorYellow), "M /s not found")
// 	}

// 	P3 := &W.Letterw{Date}
// 	start3, end3, count3, bol3 := W.Find(P3)
// 	if bol3 {
// 		fmt.Println(string(colorBlue), "W /s found:=>", count3)
// 		fmt.Println("Start: ", start3, "End: ", end3, "Count: ", count3)

// 	} else {
// 		fmt.Println(string(colorBlue), "W /s not found")
// 	}

// Semicircle.Plot_Pattern(Date, count, start, end, "Inverted.png")
// Semicircle.Plot_Pattern(Date, count1, start1, end1, "Semicircle.png")
// Semicircle.Plot_Pattern(Date, count2, start2, end2, "M1.png")
// 	Semicircle.Plot_Pattern(Date, count3, start3, end3, "W.png")

// }
func main() {
	Date := Generic.Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/data1.csv")
	var p Structs_for_patterns.Patterns = &Semicircle.Letter_semicircle{Date}
	start2, end2, count2, bol2 := p.Pattern()
	if bol2 {
		fmt.Println("semi_circle/s found:=>", count2)
		fmt.Println("Start: ", start2, "End: ", end2, "Count: ", count2)
		Semicircle.Plot_Pattern(Date, count2, start2, end2, "semi_circle1.png")
	} else {
		fmt.Println("semi_circle/s not found")
	}
}
