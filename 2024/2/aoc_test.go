package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	fmt.Println("Part 1 Test")
	input1 := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	expected1 := 2
	result1 := part1(input1)

	if result1 != expected1 {
		t.Fatalf("Incorrect result, got = %d, expected = %d\n", result1, expected1)
	}

	input2 := [][]int {
		{1, 5, 6, 7, 8, 0},
		{11, 8, 5, 4, 3, 2},
		{1, 5, 6, 7, 8, 0},
		{1, 5, 6, 7, 8, 0},
		{0, 1, 4, 7, 10, 11},
		{50, 47, 44, 41, 38, 35},
	}

	expected2 := 3
	result2 := part1(input2)

	if result2 != expected2 {
		t.Fatalf("Incorrect result, got = %d, expected = %d\n", result2, expected2)
	}	
}

func TestPart2(t *testing.T) {
	fmt.Println("Part 2 test")

	input1 := [][]int {
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	expected1 := 4
	result1 := part2(input1)

	if result1 != expected1 {
		t.Fatalf("Incorrect result, got = %d, expected = %d\n", result1, expected1)
	}
}
