package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}

	defer f.Close()

	buffer := bufio.NewReader(f)

	var ll []int
	var lr []int

	for {
		l, err := buffer.ReadString('\n')
		if err != nil {
			break
		}

		split := strings.Split(l, " ")

		var v int

		v, err = strconv.Atoi(split[0])
		if err != nil {
			panic(fmt.Sprintf("Error converting string to int: %v", err))
		}
		ll = append(ll, v)

		v, err = strconv.Atoi(split[3][:len(split[3])-1])
		if err != nil {
			panic(fmt.Sprintf("Error converting string to int: %v", err))
		}
		lr = append(lr, v)
	}

	slices.Sort(ll)
	slices.Sort(lr)

	result1 := part1(ll, lr)
	fmt.Printf("Result1: %d\n", result1)

	result2 := part2(ll, lr)
	fmt.Printf("Result2: %d", result2)
}

func part1(l, r []int) int {
	v := 0

	for i := range l {
		t := l[i] - r[i]
		if t < 0 {
			t *= -1
		}

		v += t
	}

	return v
}

func part2(l, r []int) int {
	hashRight := make(map[int]int)
	score := 0

	for _, num := range r {
		hashRight[num]++
	}

	for _, num := range l {
		score += hashRight[num] * num 
	}

	return score
}

