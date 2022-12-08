package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var dirMap map[string]int
var currDir = "/"

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dirMap = make(map[string]int, 0)
	dirMap["/"] = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output := strings.Split(scanner.Text(), " ")
		switch output[0] {
		case "$":
			if output[1] == "cd" {
				if output[2] == ".." {
					dirs := strings.Split(currDir, "/")
					currDir = strings.Join(dirs[0:len(dirs)-1], "/")
					if currDir == "" {
						currDir = "/"
					}
				} else if output[2] != "/" {
					if currDir != "/" {
						currDir += "/" + output[2]
					} else {
						currDir += output[2]
					}
				}
			} else if output[1] == "ls" {

				continue
			} else {
				panic("unknown command")
			}
		case "dir":
		default:
			size, _ := strconv.Atoi(output[0])
			addFileSizeToFolders(size)
		}
	}

	solution1 := 0

	used := dirMap["/"]
	total := 70_000_000
	unused := total - used
	needed := 30_000_000
	candidates := make([]int, 0)

	for _, v := range dirMap {
		if v <= 100000 {
			solution1 += v
		}
		if unused+v >= needed {
			candidates = append(candidates, v)
		}
	}

	sort.Ints(candidates)

	fmt.Println(solution1)
	fmt.Println(candidates[0])
}

func addFileSizeToFolders(size int) {
	dirMap[currDir] += size

	if currDir == "/" {
		return
	}
	// add filesize to all previous folders
	dirs := strings.Split(currDir, "/")
	curr := "/"
	for _, d := range dirs {
		if d == "" {
			dirMap["/"] += size
			continue
		}
		if curr == "/" {
			curr += d
		} else {
			curr += "/" + d
		}
		if curr == currDir {
			break
		}

		dirMap[curr] += size

	}
}
