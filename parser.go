package main

import "fmt"

func ParseLine(line *[]byte) {
	fmt.Printf("Got line: %s\n", string(*line))
}
