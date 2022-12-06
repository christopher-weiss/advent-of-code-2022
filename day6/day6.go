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

	fmt.Printf("Start of Packet:%d\n", getMarker(stream, 4))
	fmt.Printf("Start of Message:%d\n", getMarker(stream, 14))
}

func getMarker(stream string, size int) int {
	var buffer string
	var marker int
	for _, r := range stream[:] {
		if len(buffer) < size {
			buffer += string(r)
		} else {
			if isMarker(buffer) {
				return marker
			}
			buffer = buffer[1:]
			buffer += string(r)
		}
		marker++
	}
	return -1
}

func isMarker(buffer string) bool {
	for _, r := range buffer[:] {
		if strings.Count(buffer, string(r)) > 1 {
			return false
		}
	}
	return true
}
