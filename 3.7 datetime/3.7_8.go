package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const now = 1589570165

/**
Sample input:
12 мин. 13 сек.

Sample output:
Fri May 15 19:28:18 UTC 2020
*/
func main() {
	var (
		durationStr string
		baseDate    time.Time
		duration    time.Duration
		loc         *time.Location
		err         error
	)

	baseDate = time.Unix(now, 0)
	loc, err = time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	baseDate = baseDate.In(loc)

	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	durationStr = scanner.Text()
	durationStr = strings.ReplaceAll(durationStr, " ", "")
	durationStr = strings.ReplaceAll(durationStr, "мин.", "m")
	durationStr = strings.ReplaceAll(durationStr, "сек.", "s")

	duration, err = time.ParseDuration(durationStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(baseDate.Add(duration).Format(time.UnixDate))
}
