package main

import "fmt"

func main() {
	fmt.Println(season(0))
	fmt.Println(season(1))
	fmt.Println(season(12))
	fmt.Println(season(100))
}

func season(month int) string {
	switch {
	case month <= 0:
		return "None"
	case month > 0 && month <= 3:
		return "Winter"
	case month > 3 && month <= 6:
		return "Spring"
	case month > 6 && month <= 9:
		return "Summer"
	case month > 9 && month <= 12:
		return "Fall"
	case month > 12:
		return "None"
	}
	return "None"
}