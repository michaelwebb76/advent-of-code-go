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

func main() {
	flag.Parse()

	filename := flag.Arg(0)

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)

	frequency := 0
	for err == nil {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err == nil {
			lineString := string(line)

			sign := lineString[0:1]
			number := lineString[1:len(lineString)]

			if newFrequency, err := strconv.Atoi(number); err == nil {
				if sign == "+" {
					frequency += newFrequency
				} else {
					frequency -= newFrequency
				}
			} else {
				panic(err)
			}
		}
	}
	fmt.Println(frequency)

	return
}
