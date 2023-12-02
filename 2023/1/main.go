package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// this replace with letter at the end of the number is because, we want to reuse the letter for other numbers as well
// like "oneight": "one" and "eight" both have "e" (at the end and at the beginning) so we can use the same letter for both
var replacements = map[string]string{
	"one":   "1e",
	"two":   "2o",
	"three": "3e",
	"four":  "4r",
	"five":  "5e",
	"six":   "6x",
	"seven": "7n",
	"eight": "8t",
	"nine":  "9e",
}
var backwardReplacements = map[string]string{
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	REGEX := `one|two|three|four|five|six|seven|eight|nine`
	scanner := bufio.NewScanner(os.Stdin)
	var counter int
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("(" + REGEX + ")")
		line = string(re.ReplaceAllFunc([]byte(line), func(match []byte) []byte {
			if _, ok := replacements[string(match)]; ok {
				return []byte(replacements[string(match)])
			}
			return match
		}))
		// reverse the string
		line = reverse(line)
		re = regexp.MustCompile("(" + reverse(REGEX) + ")")
		line = string(re.ReplaceAllFunc([]byte(line), func(match []byte) []byte {
			if _, ok := backwardReplacements[string(match)]; ok {
				return []byte(backwardReplacements[string(match)])
			}
			return match
		}))
		// remove all non numeric characters
		re = regexp.MustCompile(`[^0-9]`)
		line = string(re.ReplaceAllFunc([]byte(line), func(match []byte) []byte {
			return []byte("")
		}))
		// get the first and last character
		first := line[0]
		last := line[len(line)-1]
		num, err := strconv.Atoi(string(last) + string(first)) // swap the first and last character (cause it was reversed)
		if err != nil {
			panic(err)
		}
		counter += num
	}
	fmt.Println(counter)

}
