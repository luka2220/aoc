package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}

	var input [][]int
	buffer := bufio.NewReader(f)

	for {
		report, err := buffer.ReadSlice('\n')
		if err != nil {
			break
		}

		clean := bytes.Split(report, []byte(" "))
		var reportInt []int

		for i, v := range clean {
			if i == len(clean)-1 {
				num, err := strconv.Atoi(string(v)[:len(v)-1])
				if err != nil {
					panic(fmt.Sprintf("Error converting string to int: %v", err))
				}
				reportInt = append(reportInt, num)
				continue
			}

			num, err := strconv.Atoi(string(v))
			if err != nil {
				panic(fmt.Sprintf("Error converting string to int: %v", err))
			}

			reportInt = append(reportInt, num)
		}

		input = append(input, reportInt)
	}

	p1ans := part1(input)
	fmt.Println(p1ans)
	p2ans := part2(input)
	fmt.Println(p2ans)
}

func part1(data [][]int) int {
	safe := 0

	for i := range data {
		result := true
		if data[i][0] < data[i][1] {
			// Increasing logic
			for j, v := range data[i] {
				if j == len(data[i]) - 1 {
					break
				}

				level := data[i][j+1] - v
				if level > 3 || level <= 0 {
					result = false
					break
				}
			}

			if result {
				safe += 1
			}
		} else if data[i][0] > data[i][1] {
			// Decreasing logic
			for j, v := range data[i] {
				if j == len(data[i]) - 1 {
					break
				}

				level := v - data[i][j+1]
				if level > 3 || level <= 0 {
					result = false
					break
				}
			}

			if result {
				safe += 1
			}
		}
	}

	return safe
}

func part2(data [][]int) int {
	safe := 0
	for _, list := range data {
		if part1([][]int{list}) == 1 {
			safe++
		} else {
			for i := range list {
				if part1([][]int{remove(list, i)}) == 1 {
					safe++
					break	
				}
			}
		}
	}

	return safe
}

func remove(slice []int, i int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:i]...)
	return append(newSlice, slice[i+1:]...)
}