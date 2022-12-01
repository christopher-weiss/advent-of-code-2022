package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalCals := make([]int, 0)
	var currCal int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			totalCals = append(totalCals, currCal)
			currCal = 0
		} else {
			num, _ := strconv.Atoi(line)
			currCal += num
		}
	}

	sort.Ints(totalCals)

	first := totalCals[len(totalCals)-1]
	second := totalCals[len(totalCals)-2]
	third := totalCals[len(totalCals)-3]
	top3Total := first + second + third

	fmt.Printf("Most calories: %d\nTop 3 total calories: %d\n", first, top3Total)
}
