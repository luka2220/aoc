package main

import (
	"fmt"
	"testing"
)

func TestGetSimilarityScore(t *testing.T) {
	fmt.Println("Similarity test...")
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	expected := 31
	result := part2(left, right)

	if result != expected {
		t.Fatalf("Incorrect result, got %d, expected %d", result, expected)
	}
}

