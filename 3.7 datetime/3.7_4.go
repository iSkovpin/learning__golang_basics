package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var (
		dateStr    string
		dateTime   time.Time
		dateFormat string = "2006-01-02 15:04:05"
		err        error
	)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dateStr = scanner.Text()

	// Layout
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// Input
	// 2020-05-15 08:00:00
	// 2020-05-15 13:00:00
	// Output
	// 2020-05-15 08:00:00
	// 2020-05-16 13:00:00

	dateTime, err = time.Parse(dateFormat, dateStr)
	if err != nil {
		panic(err)
	}

	if dateTime.Hour() >= 13 {
		dateTime = dateTime.AddDate(0, 0, 1)
	}

	fmt.Println(dateTime.Format(dateFormat))
}
