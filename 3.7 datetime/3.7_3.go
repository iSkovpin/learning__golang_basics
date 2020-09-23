package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var (
		dateStr  string
		dateTime time.Time
		err      error
	)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dateStr = scanner.Text()

	// Layout
	// Mon Jan 2 15:04:05 -0700 MST 2006
	// Input
	// 1986-04-16T05:20:00+06:00
	// Output
	// Wed Apr 16 05:20:00 +0600 1986

	dateTime, err = time.Parse("2006-01-02T15:04:05-07:00", dateStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(dateTime.Format(time.UnixDate))
}
