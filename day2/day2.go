package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pointMap map[rune]int
var winningMove map[rune]rune

func main() {
	winningMove = map[rune]rune{
		'A': 'Y',
		'B': 'Z',
		'C': 'X',
	}
	pointMap = map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': 1,
		'Y': 2,
		'Z': 3,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalScore int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalScore += calculateRoundScore(rune(line[0]), rune(line[2]))
	}

	fmt.Println(totalScore)
}

func calculateRoundScore(opponentMove rune, myMove rune) int {
	score := pointMap[myMove]
	if pointMap[opponentMove] == pointMap[myMove] {
		score += 3
	}
	if winningMove[opponentMove] == myMove {
		score += 6
	}

	return score
}
