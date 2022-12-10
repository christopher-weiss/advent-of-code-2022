package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x,y int
}

var head, tail position

var visited map[string]bool

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	visited = make(map[string]bool)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		movement := strings.Split(line, " ")

		dir := movement[0]
		dist,_ := strconv.Atoi(movement[1])

		for i:=0; i < dist; i++ {
			moveHead(dir)
			moveTail()
			visited[fmt.Sprintf("%d,%d",tail.x, tail.y)] = true
		}
	}

	fmt.Println(len(visited))
}

func moveHead(dir string) {
	switch dir {
	case "R": head.x += 1
	case "U": head.y += 1
	case "L": head.x -= 1
	case "D": head.y -= 1
	default:
		panic("direction not known")
	}
}

func moveTail() {
	distToHeadX := math.Abs(float64(head.x - tail.x))
	distToHeadY := math.Abs(float64(head.y - tail.y))

	if distToHeadX > 1 || distToHeadY > 1 {
		if distToHeadX + distToHeadY >= 3 {
			if (head.x > tail.x) && (head.y > tail.y) {
				tail.x += 1
				tail.y += 1
			}
			if (head.x > tail.x) && (head.y < tail.y) {
				tail.x += 1
				tail.y -= 1
			}
			if (head.x < tail.x) && (head.y > tail.y) {
				tail.x -= 1
				tail.y += 1
			}
			if (head.x < tail.x) && (head.y < tail.y) {
				tail.x -= 1
				tail.y -= 1
			}
		} else {
			if distToHeadX > 1 {
				if tail.x < head.x {
					tail.x += head.x - tail.x - 1
				} else {
					tail.x -= tail.x - head.x - 1
				}
			}
			if distToHeadY > 1 {
				if tail.y < head.y {
					tail.y += head.y - tail.y - 1
				} else {
					tail.y -= tail.y - head.y - 1
				}
			}
		}
	}
}
