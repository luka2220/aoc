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

	//ll := make([]int, buffer.Size()/2)
	// lr := make([]int, buffer.Size()/2)
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

	result1 := getListDiff(ll, lr)
	fmt.Printf("Result1: %d\n", result1)

	result2 := getSimilarityScore(ll, lr)
	fmt.Printf("Result2: %d", result2)
}

// Gets the difference of ascending numbers in the left list compared to right
func getListDiff(l, r []int) int {
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

// Gets the similarity difference of values from the left list in the right list
func getSimilarityScore(l, r []int) int {
	hashRight := make(map[int]int)
	score := 0

	for i := range r {
		_, ok := hashRight[r[i]]
		if !ok {
			hashRight[r[i]] = 1
			continue
		}

		hashRight[r[i]] += 1
	}

	for i := range l {
		value, ok := hashRight[l[i]]
		if !ok {
			continue
		}

		score += l[i] * value
	}

	return score
}

