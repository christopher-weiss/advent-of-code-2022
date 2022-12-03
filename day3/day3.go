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

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		comp1, comp2 := seperateIntoCompartments(rucksack)

		for _, r := range comp1[:] {
			if strings.ContainsRune(comp2, r) {
				sum += priority(r)
				break
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}

func seperateIntoCompartments(rucksack string) (string, string) {
	size := int(len(rucksack) / 2)

	return rucksack[0:size], rucksack[size:]
}

func priority(item rune) int {
	priorities := ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(priorities, string(item))
}
