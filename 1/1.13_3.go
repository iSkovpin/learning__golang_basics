package main

import "fmt"

func main() {
	var sec, hours, mins int
	fmt.Scan(&sec)

	hours = sec / 3600
	mins = sec % 3600 / 60

	fmt.Printf("It is %d hours %d minutes.", hours, mins)
}
