package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	NUMERIC_REGEX := regexp.MustCompile("[0-9]+")
	NOT_NUM_OR_DOT_REGEX := regexp.MustCompile("[^0-9.]")

	scanner := bufio.NewScanner(os.Stdin)
	// read the input until eof and save in a list
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var sum int

	// iterate over the input, getting the actual line, the previous line and the next line
	for i := 0; i < len(input); i++ {
		var line string
		var prevLine string
		var nextLine string
		if i == 0 {
			line = input[i]
			prevLine = ""
			nextLine = input[i+1]
		} else if i == len(input)-1 {
			line = input[i]
			prevLine = input[i-1]
			nextLine = ""
		} else {
			line = input[i]
			prevLine = input[i-1]
			nextLine = input[i+1]
		}
		// find the indexes of numeric characters in the line
		indexes := NUMERIC_REGEX.FindAllIndex([]byte(line), -1)

		for _, index := range indexes {
			i, j := index[0], index[1]
			// if in prevLine or nextLine range is a character different from a number or '.', print the number
			startRange := i - 1
			endRange := j + 1
			if i-1 < 0 {
				startRange = 0
			}
			if j+1 > len(line)-1 {
				endRange = len(line)
			}
			prevChar := string(line[startRange])
			nextChar := string(line[endRange-1])

			var prevLineSlice string
			var nextLineSlice string
			if prevLine != "" {
				prevLineSlice = prevLine[startRange:endRange]
			}
			if nextLine != "" {
				nextLineSlice = nextLine[startRange:endRange]
			}
			if NOT_NUM_OR_DOT_REGEX.MatchString(prevLineSlice) || NOT_NUM_OR_DOT_REGEX.MatchString(nextLineSlice) || NOT_NUM_OR_DOT_REGEX.MatchString(prevChar) || NOT_NUM_OR_DOT_REGEX.MatchString(nextChar) {
				num, err := strconv.Atoi(line[i:j])
				if err != nil {
					panic(err)
				}
				sum += num
			}
		}
	}
	fmt.Println(sum)

}
