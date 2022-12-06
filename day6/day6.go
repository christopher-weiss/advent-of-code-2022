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

	var stream string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stream = scanner.Text()
	}

	buffer := ""
	var marker int
	for _, r := range stream[:] {
		if len(buffer) < 4 {
			buffer += string(r)
		} else {
			if isMarker(buffer) {
				fmt.Println(marker)
				break
			}
			buffer = buffer[1:]
			buffer += string(r)
		}
		marker++
	}
}

func isMarker(buffer string) bool {
	for _, r := range buffer[:] {
		if strings.Count(buffer, string(r)) > 1 {
			return false
		}
	}
	return true
}
