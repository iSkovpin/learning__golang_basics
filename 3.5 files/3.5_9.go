package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var sum int
	var scanner = bufio.NewScanner(os.Stdin)

	for scan := scanner.Scan(); scan != false; scan = scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		sum += num
	}

	var writer = bufio.NewWriter(os.Stdout)
	_, err := writer.WriteString(strconv.Itoa(sum))
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
