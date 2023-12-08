package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var NON_NUMERIC_REGEX *regexp.Regexp = regexp.MustCompile("[^0-9]+")

func getCardNumAndValues(line string) (int, map[string]bool, []string) {
	split := strings.Split(line, ": ")
	cardNum, err := strconv.Atoi(string(NON_NUMERIC_REGEX.ReplaceAll([]byte(split[0]), []byte(""))))
	if err != nil {
		panic(err)
	}
	cardValues := split[1]
	split = strings.Split(cardValues, "|")
	winningNumsStr, numsYouHaveStr := split[0], split[1]
	winningNums := map[string]bool{}

	for _, num := range strings.Split(winningNumsStr, " ") {
		key := string(NON_NUMERIC_REGEX.ReplaceAll([]byte(num), []byte("")))
		if key != "" {
			winningNums[key] = true
		}
	}
	numsYouHave := strings.Split(numsYouHaveStr, " ")
	return cardNum, winningNums, numsYouHave
}

func getWinQtty(winningNums map[string]bool, numsYouHave []string) int {
	winQtty := 0
	for _, num := range numsYouHave {
		key := string(NON_NUMERIC_REGEX.ReplaceAll([]byte(num), []byte("")))
		if _, ok := winningNums[key]; ok {
			winQtty++
		}
	}
	return winQtty
}

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	value_multiplier := map[int]int{}
	for scanner.Scan() {
		cardNum, winningNums, numsYouHave := getCardNumAndValues(scanner.Text())
		if _, ok := value_multiplier[cardNum]; !ok {
			value_multiplier[cardNum] = 1
		}
		winQtty := getWinQtty(winningNums, numsYouHave)
		for i := cardNum + 1; i <= cardNum+winQtty; i++ {
			if _, ok := value_multiplier[i]; ok {
				value_multiplier[i] += value_multiplier[cardNum]
			} else {
				value_multiplier[i] = 1 + value_multiplier[cardNum]
			}
		}
		fmt.Println(cardNum, winQtty, value_multiplier[cardNum])
		sum += value_multiplier[cardNum]
	}
	fmt.Println(sum)
}
