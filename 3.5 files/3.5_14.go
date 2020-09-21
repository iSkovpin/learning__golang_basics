package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f, err = os.Open("./3.5 files/numSequence.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var reader = bufio.NewReader(f)
	var delim, zero byte = ';', '0'

	var i int = 1
	for str, err := reader.ReadString(delim); err == nil; str, err = reader.ReadString(delim) {
		if len(str) == 2 && str[0] == zero {
			break
		}
		i++
	}

	fmt.Println(i)
}
