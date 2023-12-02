package utils

import (
	"regexp"
	"strconv"
)

func ParseLine(s string) (gameId int, ballsInfo string) {
	re := regexp.MustCompile(`(Game \d+: )`)
	ballsInfo = re.ReplaceAllString(s, "")
	game := re.FindString(s)

	// delete all non numeric characters
	re = regexp.MustCompile(`[^0-9]`)
	gameId, err := strconv.Atoi(re.ReplaceAllString(game, ""))
	if err != nil {
		panic(err)
	}
	return gameId, ballsInfo
}
