package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
const (
	crtWidth int = 40
	crtHeight int = 6
)
var x, cycle int
var signalMap map[int]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x = 1
	cycle = 1
	signalMap = make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		signalMap[cycle] = x
		if line == "noop" {
			cycle++
		} else {
			command := strings.Split(line, " ")
			val, _ := strconv.Atoi(command[1])
			signalMap[cycle + 1] = x
			x += val
			signalMap[cycle + 2] = x
			cycle += 2
		}
	}

	solution1 := signalMap[20] * 20 + signalMap[60] * 60 + signalMap[100] * 100 + signalMap[140] * 140 + signalMap[180] * 180 + signalMap[220] * 220
	fmt.Println(solution1)

	drawCrt()
}

func drawCrt() {
	for i := 1; i <= crtHeight * crtWidth; i++ {
		char := ""
		pos := (i%crtWidth)-1
		if signalMap[i] == pos || signalMap[i] == pos-1 || signalMap[i] == pos+1 {
			char = "#"
		} else {
			char = "."
		}
		if i % crtWidth == 0 {
			fmt.Printf("%s\n", char)
		} else {
			fmt.Printf("%s", char)
		}
	}
}
