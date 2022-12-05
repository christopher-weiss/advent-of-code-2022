package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type _range struct {
	start int
	end   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sumContains, sumOverlaps int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		range1, range2 := buildRanges(line)
		if containsTheOther(range1, range2) {
			sumContains++
		}
		if overlaps(range1, range2) {
			sumOverlaps++
		}
	}

	fmt.Printf("Contains: %d\n", sumContains)
	fmt.Printf("Overlaps: %d\n", sumOverlaps)
}

func buildRanges(line string) (_range, _range) {
	ranges := strings.Split(line, ",")
	range1 := strings.Split(ranges[0], "-")
	range2 := strings.Split(ranges[1], "-")
	range1Start, _ := strconv.Atoi(range1[0])
	range1End, _ := strconv.Atoi(range1[1])
	range2Start, _ := strconv.Atoi(range2[0])
	range2End, _ := strconv.Atoi(range2[1])
	return _range{range1Start, range1End}, _range{range2Start, range2End}
}

func containsTheOther(range1 _range, range2 _range) bool {
	return (range1.start >= range2.start && range1.end <= range2.end) ||
		(range2.start >= range1.start && range2.end <= range1.end)
}

func overlaps(range1 _range, range2 _range) bool {
	range1Nums := toRange(range1)
	range2Nums := toRange(range2)

	for _, i := range range1Nums {
		for _, j := range range2Nums {
			if i == j {
				return true
			}
		}
	}

	return false
}

func toRange(r _range) []int {
	nums := make([]int, 0)
	for i := r.start; i <= r.end; i++ {
		nums = append(nums, i)
	}
	return nums
}
