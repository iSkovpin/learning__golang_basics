package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		s, a, b         string
		aF64, bF64, res float64
		numsStr         []string
		err             error
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()

	mapping := func(r rune) rune {
		if r == ' ' {
			return -1
		}
		if r == ',' {
			return '.'
		}
		return r
	}

	numsStr = strings.Split(s, ";")
	a, b = strings.Map(mapping, numsStr[0]), strings.Map(mapping, numsStr[1])

	aF64, err = strconv.ParseFloat(a, 64)
	if err != nil {
		panic(err)
	}

	bF64, err = strconv.ParseFloat(b, 64)
	if err != nil {
		panic(err)
	}

	res = aF64 / bF64

	fmt.Printf("%.4f", res)
}
