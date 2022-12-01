package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var currCal, maxCal int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currCal > maxCal {
				maxCal = currCal
			}
			currCal = 0
		} else {
			num, _ := strconv.Atoi(line)
			currCal += num
		}
	}

	fmt.Println(maxCal)
}
