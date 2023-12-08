package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	NON_NUMERIC_REGEX := regexp.MustCompile("[^0-9]+")
	for scanner.Scan() {
		actual_num := 0
		input := scanner.Text()
		cardValues := strings.Split(input, ": ")[1]
		split := strings.Split(cardValues, "|")

		winningNumsStr, numsYouHaveStr := split[0], split[1]
		winningNums := map[string]bool{}
		for _, num := range strings.Split(winningNumsStr, " ") {
			key := string(NON_NUMERIC_REGEX.ReplaceAll([]byte(num), []byte("")))
			if key != "" {
				winningNums[key] = true
			}
		}

		for _, num := range strings.Split(numsYouHaveStr, " ") {
			key := string(NON_NUMERIC_REGEX.ReplaceAll([]byte(num), []byte("")))
			if _, ok := winningNums[string(key)]; ok {
				if actual_num == 0 {
					actual_num = 1
				} else {
					actual_num = actual_num * 2
				}
			}
		}
		sum += actual_num
	}
	fmt.Println(sum)
}
