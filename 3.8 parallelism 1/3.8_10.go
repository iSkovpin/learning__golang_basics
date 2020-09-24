package main

import (
	"fmt"
)

func main() {
	inputStream := make(chan string, 16)
	outputStream := make(chan string)

	inputStream <- "A"
	inputStream <- "B"
	inputStream <- "B"
	inputStream <- "B"
	inputStream <- "C"
	inputStream <- "c"
	inputStream <- "A"
	close(inputStream)

	go removeDuplicates(inputStream, outputStream)
	printFromChannel(outputStream)
}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	lastStr := <-inputStream
	outputStream <- lastStr

	for str := range inputStream {
		if str == lastStr {
			continue
		}

		lastStr = str
		outputStream <- str
	}

	close(outputStream)
}

func printFromChannel(ch chan string) {
	for str := range ch {
		fmt.Print(str)
	}
}
