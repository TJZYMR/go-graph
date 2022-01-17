package main

import (
	"fmt"
	"y/Generic"
	"y/Inverted_semi_circle"
	"y/M"
	"y/Semicircle"
	"y/W"
)

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"

	Date := Generic.Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/Generic/data.csv")
	P := &Inverted_semi_circle.Letter_inverted_semicircle{Date}
	start, end, count, bol := Inverted_semi_circle.Find(P)
	if bol {
		fmt.Println(string(colorRed), "Semicircle/s found:=>", count)
		fmt.Println("Start: ", start, "End: ", end, "Count: ", count)
	}

	P1 := &Semicircle.Letter_semicircle{Date}
	start1, end1, count1, bol1 := Semicircle.Find(P1)
	if bol1 {
		fmt.Println(string(colorGreen), "Inverted_Semicircle/s found:=>", count1)
		fmt.Println("Start: ", start1, "End: ", end1, "Count: ", count1)
	}

	P2 := &M.Letterm{Date}
	start2, end2, count2, bol2 := M.Find(P2)
	if bol2 {
		fmt.Println(string(colorYellow), "M /s found:=>", count2)
		fmt.Println("Start: ", start2, "End: ", end2, "Count: ", count2)
	}

	P3 := &W.Letterw{Date}
	start3, end3, count3, bol3 := W.Find(P3)
	if bol3 {
		fmt.Println(string(colorBlue), "W /s found:=>", count3)
		fmt.Println("Start: ", start3, "End: ", end3, "Count: ", count3)
	}
}
