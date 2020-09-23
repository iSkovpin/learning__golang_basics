package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("./3.6 JSON/group.json")
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}

	var group Group = Group{}
	if err := json.Unmarshal(data, &group); err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}

	var studentsCount, ratingsCount int
	var averageRatingsCount float64
	for _, st := range group.Students {
		studentsCount++
		ratingsCount += len(st.Rating)
	}

	averageRatingsCount = float64(ratingsCount) / float64(studentsCount)
	var avg AverageRatings = AverageRatings{Average: averageRatingsCount}

	str, err := json.MarshalIndent(avg, "", "    ")
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(str)
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type AverageRatings struct {
	Average float64
}
