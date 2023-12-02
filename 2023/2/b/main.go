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
	sum := 0
	for scanner.Scan() {
		_, line := utils.ParseLine(scanner.Text())
		split := strings.Split(line, "; ")
		gameMaxBlue := 0
		gameMaxRed := 0
		gameMaxGreen := 0
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
				if ballsColor == "blue" && ballsNumber > gameMaxBlue {
					gameMaxBlue = ballsNumber
				} else if ballsColor == "red" && ballsNumber > gameMaxRed {
					gameMaxRed = ballsNumber
				} else if ballsColor == "green" && ballsNumber > gameMaxGreen {
					gameMaxGreen = ballsNumber
				}
			}
		}
		sum += gameMaxBlue * gameMaxRed * gameMaxGreen
	}
	fmt.Println(sum)
}
