package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const limit = 32768

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToStringArray(filename string) []string {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	stringArray := make([]string, 0, limit)

	for err == nil {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err == nil {
			lineString := string(line)
			stringArray = append(stringArray, lineString)
		}
	}

	return stringArray
}

func containsTheSameLetterXTimes(stringValue string, count int) bool {
	for _, r := range stringValue {
		if strings.Count(stringValue, string(r)) == count {
			return true
		}
	}
	return false
}

// CalculateChecksum ...
func CalculateChecksum(stringArray []string) int {
	sumOfTwos := 0
	sumOfThrees := 0
	for _, stringValue := range stringArray {
		if containsTheSameLetterXTimes(stringValue, 2) {
			sumOfTwos++
		}
		if containsTheSameLetterXTimes(stringValue, 3) {
			sumOfThrees++
		}
	}
	return sumOfTwos * sumOfThrees
}

func choose(stringArray []string, test func(string) bool) (result []string) {
	for _, string := range stringArray {
		if test(string) {
			result = append(result, string)
		}
	}
	return
}

// CharsDifferent ...
func CharsDifferent(firstString string, secondString string) int {
	firstStringLen := len(firstString)
	secondStringLen := len(secondString)
	result := int(math.Abs(float64(firstStringLen - secondStringLen)))

	longestString := ""
	shortestString := ""
	if firstStringLen > secondStringLen {
		longestString = firstString
		shortestString = secondString
	} else {
		longestString = secondString
		shortestString = firstString
	}

	for pos, testRune := range shortestString {
		if testRune != rune(longestString[pos]) {
			result++
		}
	}
	return result
}

// StringsWithOneCharDifference ...
func StringsWithOneCharDifference(stringArray []string) (result []string) {
	for _, stringValue := range stringArray {
		chooseFunc := func(testString string) bool { return CharsDifferent(testString, stringValue) == 1 }
		result = append(result, choose(stringArray, chooseFunc)...)
	}
	return
}

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	stringArray := fileToStringArray(filename)
	fmt.Println("Checksum is: " + strconv.Itoa(CalculateChecksum(stringArray)))
	fmt.Println("Boxes with one char difference are: " + strings.Join(StringsWithOneCharDifference(stringArray), ","))
}
