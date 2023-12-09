package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getLinesUntilNewline(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	return lines
}

func createMap(lines []string) func(int) int {
	var destinationRanges [][]int
	var sourceRanges [][]int

	for _, line := range lines {
		values := strings.Split(line, " ")
		destinationStartRange, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		sourceStartRange, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(values[2])
		if err != nil {
			panic(err)
		}
		sourceEndRange := sourceStartRange + rangeLength
		destinationEndRange := destinationStartRange + rangeLength
		sourceRanges = append(sourceRanges, []int{sourceStartRange, sourceEndRange})
		destinationRanges = append(destinationRanges, []int{destinationStartRange, destinationEndRange})
	}

	return func(source int) int {
		for i, range_ := range sourceRanges {
			if range_[0] <= source && source < range_[1] {
				destinationRange := destinationRanges[i]
				return (source - range_[0]) + destinationRange[0]
			}
		}
		return source
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var seeds []int
	var seedToSoil func(int) int
	var soilToFertilizer func(int) int
	var fertilizerToWater func(int) int
	var waterToLight func(int) int
	var lightToTemperature func(int) int
	var temperatureToHumidity func(int) int
	var humidityToLocation func(int) int
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		fmt.Println(lineNum)
		line := scanner.Text()
		// if line starts with "seeds: " then remove the "seeds: " part and split in " "
		if line == "" {
			continue
		}
		if line[:7] == "seeds: " {
			line = line[7:]
			values := strings.Split(line, " ")
			for _, value := range values {
				seed, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, seed)
			}
		} else if line == "seed-to-soil map:" {
			seedToSoil = createMap(getLinesUntilNewline(scanner))
		} else if line == "soil-to-fertilizer map:" {
			soilToFertilizer = createMap(getLinesUntilNewline(scanner))
		} else if line == "fertilizer-to-water map:" {
			fertilizerToWater = createMap(getLinesUntilNewline(scanner))
		} else if line == "water-to-light map:" {
			waterToLight = createMap(getLinesUntilNewline(scanner))
		} else if line == "light-to-temperature map:" {
			lightToTemperature = createMap(getLinesUntilNewline(scanner))
		} else if line == "temperature-to-humidity map:" {
			temperatureToHumidity = createMap(getLinesUntilNewline(scanner))
		} else if line == "humidity-to-location map:" {
			humidityToLocation = createMap(getLinesUntilNewline(scanner))
		}
	}
	min := math.MaxInt
	for _, seed := range seeds {
		soil := seedToSoil(seed)
		fertilizer := soilToFertilizer(soil)
		water := fertilizerToWater(fertilizer)
		light := waterToLight(water)
		temperature := lightToTemperature(light)
		humidity := temperatureToHumidity(temperature)
		location := humidityToLocation(humidity)

		if location < min {
			min = location
		}
	}
	fmt.Println(min)
}
