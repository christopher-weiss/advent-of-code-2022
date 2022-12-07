package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
			//fmt.Printf("%s : %d\n", currDir, dirMap[currDir])

			//fmt.Println(scanner.Text())
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
			if _, exists := dirMap[output[1]]; !exists {
				//dirMap[output[1]] = 0
			}
		default:
			size, _ := strconv.Atoi(output[0])
			addFileSizeToFolders(size)
		}
	}

	solution := 0
	for _, v := range dirMap {
		//fmt.Printf("%s -> %d\n", k, v)
		if v <= 100000 {
			solution += v
		}
	}
	fmt.Println(solution)
}

func addFileSizeToFolders(size int) {
	// add filesize to current folder
	dirMap[currDir] += size
	// add filesize to all previous folders
	dirs := strings.Split(currDir, "/")
	curr := "/"
	for _, d := range dirs {
		if curr == currDir {
			break
		}

		dirMap[curr] += size

		if curr == "/" {
			curr += d
		} else {
			curr += "/" + d
		}
	}
}
