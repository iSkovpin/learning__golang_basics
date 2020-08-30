package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(adding(a, b))
}

func adding(a, b string) int64 {
	a = filterNonDigit(a)
	b = filterNonDigit(b)

	aInt64, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}

	bInt64, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		panic(err)
	}

	return aInt64 + bInt64
}

func filterNonDigit(str string) string {
	var strClean string
	for _, r := range []rune(str) {
		if unicode.IsDigit(r) {
			strClean = strClean + string(r)
		}
	}

	return strClean
}
