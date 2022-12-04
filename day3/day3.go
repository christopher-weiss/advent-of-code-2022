package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sumPart1, sumPart2 int

	groups := make([][]string, 500)
	groups[0] = make([]string, 3)

	var groupIndex, groupSubIndex int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()

		// Part 1
		comp1, comp2 := separateIntoCompartments(rucksack)
		for _, r := range comp1[:] {
			if strings.ContainsRune(comp2, r) {
				sumPart1 += priority(r)
				break
			}
		}

		// Part 2
		if groupSubIndex%3 == 0 {
			groupIndex++
			groups[groupIndex] = make([]string, 3)
			groupSubIndex = 0
		}
		groups[groupIndex][groupSubIndex] = rucksack
		groupSubIndex++
	}

	for _, group := range groups {
		if group != nil {
			sumPart2 += priority(findBadge(group))
		}
	}

	fmt.Printf("Sum part 1: %d\n", sumPart1)
	fmt.Printf("Sum part 2: %d\n", sumPart2)

}

func separateIntoCompartments(rucksack string) (string, string) {
	size := len(rucksack) / 2

	return rucksack[:size], rucksack[size:]
}

func priority(item rune) int {
	priorities := ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(priorities, string(item))
}

func findBadge(rucksacks []string) rune {
	candidates := make([]rune, 0)
	for _, r1 := range rucksacks[0] {
		if strings.ContainsRune(rucksacks[1], r1) {
			candidates = append(candidates, r1)
		}
	}
	for _, r2 := range candidates {
		if strings.ContainsRune(rucksacks[2], r2) {
			return r2
		}
	}
	return '.'
}
