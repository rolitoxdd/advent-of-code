package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isAsteriskNear(numRange []int, asteriskIndex int) bool {
	if asteriskIndex+1 >= numRange[0] && asteriskIndex <= numRange[1] {
		return true
	} else {
		return false
	}
}

func main() {
	NUMERIC_REGEX := regexp.MustCompile("[0-9]+")
	ASTERISK_REGEX := regexp.MustCompile("[*]")

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
		asterisks := ASTERISK_REGEX.FindAllIndex([]byte(line), -1)
		for _, asterisk := range asterisks {
			actual_num := 1
			near_qtty := 0
			asterisk_index := asterisk[0]
			prev_line_nums := NUMERIC_REGEX.FindAllIndex([]byte(prevLine), -1)
			for _, prev_line_num := range prev_line_nums {
				if isAsteriskNear(prev_line_num, asterisk_index) {
					prev_line_num_str := prevLine[prev_line_num[0]:prev_line_num[1]]
					prev_line_num_int, err := strconv.Atoi(prev_line_num_str)
					if err != nil {
						panic(err)
					}
					actual_num *= prev_line_num_int
					near_qtty++
				}
			}
			line_nums := NUMERIC_REGEX.FindAllIndex([]byte(line), -1)
			for _, line_num := range line_nums {
				if isAsteriskNear(line_num, asterisk_index) {
					line_num_str := line[line_num[0]:line_num[1]]
					line_num_int, err := strconv.Atoi(line_num_str)
					if err != nil {
						panic(err)
					}
					actual_num *= line_num_int
					near_qtty++
				}
			}
			next_line_nums := NUMERIC_REGEX.FindAllIndex([]byte(nextLine), -1)
			for _, next_line_num := range next_line_nums {
				if isAsteriskNear(next_line_num, asterisk_index) {
					next_line_num_str := nextLine[next_line_num[0]:next_line_num[1]]
					next_line_num_int, err := strconv.Atoi(next_line_num_str)
					if err != nil {
						panic(err)
					}
					actual_num *= next_line_num_int
					near_qtty++
				}
			}
			if near_qtty == 2 {
				sum += actual_num

			}

		}

	}
	fmt.Println(sum)

}
