package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

const limit = 32768

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
	intArray := make([]int, 0, limit)

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

// SumFrequency ...
func SumFrequency(intArray []int) int {
	frequency := 0
	for _, intValue := range intArray {
		frequency += intValue
	}
	return frequency
}

func contains(intArray []int, targetValue int) bool {
	for _, intValue := range intArray {
		if intValue == targetValue {
			return true
		}
	}
	return false
}

// RepeatedFrequency ...
func RepeatedFrequency(intArray []int) *int {
	frequency := 0
	pastFrequencies := make([]int, 0, limit)
	pastFrequencies = append(pastFrequencies, frequency)
	for i := 1; i <= limit; i++ {
		for _, intValue := range intArray {
			frequency += intValue
			if contains(pastFrequencies, frequency) == true {
				return &frequency
			}
			pastFrequencies = append(pastFrequencies, frequency)
		}
	}
	return nil
}

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	intArray := fileToIntArray(filename)
	fmt.Println("Sum frequency is: " + strconv.Itoa(SumFrequency(intArray)))

	repeatedFrequency := RepeatedFrequency(intArray)
	if repeatedFrequency == nil {
		fmt.Println("Repeated frequency is: NOT FOUND")
	} else {
		fmt.Println("Repeated frequency is: " + strconv.Itoa(*repeatedFrequency))
	}
}
