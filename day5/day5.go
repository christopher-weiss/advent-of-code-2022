package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var crates []string

func main() {
	crates = make([]string, 10)
	crates[0] = "."
	crates[1] = "RNPG"
	crates[2] = "TJBLCSVH"
	crates[3] = "TDBMNL"
	crates[4] = "RVPSB"
	crates[5] = "GCQSWMVH"
	crates[6] = "WQSCDBJ"
	crates[7] = "FQL"
	crates[8] = "WMHTDLFV"
	crates[9] = "LPBVMJF"

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "move") {
			instruction := strings.Split(line, " ")
			origin, _ := strconv.Atoi(instruction[3])
			dest, _ := strconv.Atoi(instruction[5])
			amount, _ := strconv.Atoi(instruction[1])

			moveMany(origin, dest, amount)
		}
	}
	fmt.Printf("1: %s\n", crates[1])
	fmt.Printf("2: %s\n", crates[2])
	fmt.Printf("3: %s\n", crates[3])
	fmt.Printf("4: %s\n", crates[4])
	fmt.Printf("5: %s\n", crates[5])
	fmt.Printf("6: %s\n", crates[6])
	fmt.Printf("7: %s\n", crates[7])
	fmt.Printf("8: %s\n", crates[8])
	fmt.Printf("9: %s\n", crates[9])
}

func moveOne(origin int, dest int) {
	crates[dest] += string(crates[origin][len(crates[origin])-1])
	crates[origin] = crates[origin][0 : len(crates[origin])-1]
}

func moveMany(origin int, dest int, amount int) {
	for i := 0; i < amount; i++ {
		moveOne(origin, dest)
	}
}
