package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var forestSize = 99
var forest = make([][]int, forestSize)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var yIndex int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		forest[yIndex] = make([]int, len(line))
		for xIndex, height := range line {
			val, _ := strconv.Atoi(string(height))
			forest[yIndex][xIndex] = val
		}
		yIndex++
	}

	solution1 := 0
	solution2 := 0
	for i, hor := range forest {
		for j := range hor {
			solution1 += visible(i, j)
			score := calculateScenicScore(i,j)
			if score > solution2 {
				solution2 = score
			}
		}
	}
	fmt.Println(solution1)
	fmt.Println(solution2)
}

func visible(i, j int) int {
	// edges
	if i == 0 || j == 0 || i == forestSize-1 || j == forestSize-1 {
		return 1
	}

	if visibleVertical(i, j) || visibleHorizontal(i, j) {
		return 1
	}

	return 0
}

func visibleHorizontal(xIndex, yIndex int) bool {
	highestTree := 0
	var visibleLeft, visibleRight bool
	for i, tree := range forest[yIndex] {
		if i == xIndex {
			visibleLeft = highestTree < forest[yIndex][xIndex]
			highestTree = 0
		}
		if i != xIndex && tree > highestTree {
			highestTree = tree
		}
		if i == forestSize-1 {
			visibleRight = highestTree < forest[yIndex][xIndex]
		}
	}
	return visibleLeft || visibleRight
}

func visibleVertical(xIndex, yIndex int) bool {
	highestTree := 0
	var visibleUp, visibleDown bool
	for i := 0; i < forestSize; i++ {
		if i == yIndex {
			visibleUp = highestTree < forest[yIndex][xIndex]
			highestTree = 0
		}
		tree := forest[i][xIndex]
		if i != yIndex && tree > highestTree {
			highestTree = tree
		}
		if i == forestSize-1 {
			visibleDown = highestTree < forest[yIndex][xIndex]
		}
	}
	return visibleUp || visibleDown
}

func calculateScenicScore(xIndex, yIndex int) int {
	height := forest[yIndex][xIndex]

	var scUp, scDown, scLeft, scRight int

	// left
	if xIndex >= 1 {
		for i := xIndex - 1; i >= 0; i-- {
			scLeft = xIndex - i
			if forest[yIndex][i] >= height {
				break
			}
		}
	} else {
		if xIndex > 0 {
			scLeft = 1
		}
	}

	// right
	if xIndex < forestSize-2 {
		for i := xIndex + 1; i <= forestSize -1 ; i++ {
			scRight = i - xIndex
			if forest[yIndex][i] >= height {
				break
			}
		}
	} else {
		if xIndex < forestSize-1 {
			scRight = 1
		}
	}

	// up
	if yIndex >= 1 {
		for i := yIndex - 1; i >= 0; i-- {
			scUp = yIndex - i
			if forest[i][xIndex] >= height {
				break
			}
		}
	} else {
		if yIndex > 0 {
			scUp = 1
		}
	}

	// down
	if yIndex <= forestSize-2 {
		for i := yIndex + 1; i <= forestSize -1; i++ {
			scDown = i - yIndex
			if forest[i][xIndex] >= height {
				break
			}
		}
	} else {
		if yIndex < forestSize -1 {
			scDown = 1
		}
	}

	return scUp * scDown * scLeft * scRight
}