package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pointMap map[rune]int
var winningMove, losingMove map[rune]rune

func main() {
	winningMove = map[rune]rune{
		'A': 'Y',
		'B': 'Z',
		'C': 'X',
	}
	losingMove = map[rune]rune{
		'A': 'Z',
		'B': 'X',
		'C': 'Y',
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

	var totalScoreStrategy1, totalScoreStrategy2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalScoreStrategy1 += calculateRoundScore(rune(line[0]), rune(line[2]))
		myMove := chooseMoveForStrategy2(rune(line[0]), rune(line[2]))
		totalScoreStrategy2 += calculateRoundScore(rune(line[0]), myMove)
	}

	fmt.Printf("Strategy 1: %d\n", totalScoreStrategy1)
	fmt.Printf("Strategy 2: %d\n", totalScoreStrategy2)
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

func chooseMoveForStrategy2(opponentMove rune, myGoal rune) rune {
	if myGoal == 'X' {
		return losingMove[opponentMove]
	}
	if myGoal == 'Y' {
		return opponentMove
	}
	if myGoal == 'Z' {
		return winningMove[opponentMove]
	}
	return ' '
}
