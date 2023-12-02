package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/rolitoxdd/advent-of-code/2"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const MAX_BLUE = 14
	const MAX_RED = 12
	const MAX_GREEN = 13
	validSum := 0
	for scanner.Scan() {
		gameId, line := utils.ParseLine(scanner.Text())
		isValid := true
		split := strings.Split(line, "; ")
		for _, v := range split {
			// split the line in the , character
			split2 := strings.Split(v, ", ")
			for _, v2 := range split2 {
				split3 := strings.Split(v2, " ")
				ballsColor := split3[1]
				ballsNumber, err := strconv.Atoi(split3[0])
				if err != nil {
					panic(err)
				}
				if ballsColor == "blue" && ballsNumber > MAX_BLUE {
					isValid = false
				} else if ballsColor == "red" && ballsNumber > MAX_RED {
					isValid = false
				} else if ballsColor == "green" && ballsNumber > MAX_GREEN {
					isValid = false
				}
			}
		}
		if isValid {
			validSum += gameId
		}
	}
	fmt.Println(validSum)
}
