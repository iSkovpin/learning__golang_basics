package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Timiryzevo", "Nizhnerypinsk"}, // города с населением 10-99 тыс. человек
		100:  []string{"Batat", "Shizjat"},            // города с населением 100-999 тыс. человек
		1000: []string{"New York", "Old York"},        // города с населением 1000 тыс. человек и более
	}

	cityPopulation100 := map[string]int{
		"Batat":      127,
		"Shizjat":    542,
		"Timiryzevo": 101,
	}

	contain := func(a []string, x string) bool {
		for _, n := range a {
			if x == n {
				return true
			}
		}
		return false
	}

	for city := range cityPopulation100 {
		if !contain(groupCity[100], city) {
			delete(cityPopulation100, city)
		}
	}

	fmt.Println(cityPopulation100)
}
