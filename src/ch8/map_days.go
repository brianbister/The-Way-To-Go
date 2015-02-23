package main

import "fmt"

func main() {
	map_days := map[string]int{
		"Sunday": 1,
		"Monday": 2,
		"Tuesday": 3,
		"Wednesday": 4,
		"Thursday": 5,
		"Friday": 6,
		"Saturday": 7,
	}
	for key, value := range map_days {
		fmt.Printf("Name: %d, Day: %s\n", value, key)
	}
	if _, ok := map_days["Tuesday"]; ok {
		fmt.Println("Tuesday is present.")
	}
	if _, ok := map_days["Hollyday"]; ok {
		fmt.Println("Hollyday is not present.")
	}
}