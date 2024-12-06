package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to open input file: %v", err))
	}

	defer file.Close()

	fmt.Printf("Part 1 result: %d", part1(file))
}

func part1(file *os.File) int {
	const pattern = `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	// var list []string
	buffer := bufio.NewReader(file)
	result := 0

	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			break
		}

		match := re.FindAllStringSubmatch(line, -1)
		for i := range match {
			nums := match[i][len(match[i])-2:]
			v1, err := strconv.Atoi(nums[0])
			v2, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(fmt.Sprintf("Error converting string to num: %v\n", err))
			}

			result += v1 * v2
		}
	}

	return result
}

func part2(file *os.File) int {
	return 0
}
