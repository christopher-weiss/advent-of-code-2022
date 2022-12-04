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

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		range1, range2 := buildRanges(line)
		if containsTheOther(range1, range2) {
			sum++
		}
	}

	fmt.Println(sum)
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
