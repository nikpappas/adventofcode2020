package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var nikos = "nikos"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	return string(dat)
}

func readLines(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var lines []string
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}
		lines = append(lines, line)
		if err != nil {
			fmt.Println(err)
			break
		}

	}
	return lines
}
