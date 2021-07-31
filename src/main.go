package main

import (
	"battle/strategy"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	inputFile := "input.txt"
	if len(args) > 0 {
		inputFile = args[0]
	}
	// read input file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error while reading file", err)
	}
	result := strategy.BuildStrategy(string(data))
	fmt.Println(result)
}