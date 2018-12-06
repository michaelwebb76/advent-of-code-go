package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToIntArray(filename string) []int {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	intArray := make([]int, 0, 32768)

	for err == nil {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err == nil {
			lineString := string(line)

			if intValue, err := strconv.Atoi(lineString); err == nil {
				intArray = append(intArray, intValue)
			} else {
				panic(err)
			}
		}
	}

	return intArray
}

func sumFrequency(intArray []int) int {
	frequency := 0
	for _, intValue := range intArray {
		frequency += intValue
	}
	return frequency
}

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	intArray := fileToIntArray(filename)
	fmt.Println("Sum frequency is: " + strconv.Itoa(sumFrequency(intArray)))
}
