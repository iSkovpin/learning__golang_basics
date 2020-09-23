package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/**
Sample input:
13.03.2018 14:00:15,12.03.2018 14:00:15

Sample output:
24h0m0s
*/
func main() {
	var (
		datesStr   []string
		dateFormat string = "02.01.2006 15:04:05"
		dateOne    time.Time
		dateTwo    time.Time
		err        error
	)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	datesStr = strings.Split(scanner.Text(), ",")

	dateOne, err = time.Parse(dateFormat, datesStr[0])
	if err != nil {
		panic(err)
	}

	dateTwo, err = time.Parse(dateFormat, datesStr[1])
	if err != nil {
		panic(err)
	}

	diff := dateOne.Sub(dateTwo)
	if diff < 0 {
		diff *= -1
	}

	fmt.Println(diff)
}
