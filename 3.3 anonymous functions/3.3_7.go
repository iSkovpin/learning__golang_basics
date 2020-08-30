package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint
	fmt.Scan(&a)

	fn := func(a uint) uint {
		var (
			res       string
			resUint   uint
			resUint64 uint64
			err       error
		)

		aStr := strconv.FormatUint(uint64(a), 10)

		for _, char := range aStr {
			num, err := strconv.ParseInt(string(char), 10, 8)
			if err != nil {
				panic(err)
			}

			if num%2 != 0 || num == 0 {
				continue
			}

			res += string(char)
		}

		if res == "" {
			return 100
		}

		resUint64, err = strconv.ParseUint(res, 10, 64)
		if err != nil {
			panic(err)
		}

		resUint = uint(resUint64)

		return resUint
	}

	fmt.Println(fn(a))
}
