package main

import (
	"fmt"
	"os"
)

type coordinates struct {
	x int
	y int
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func partOne(data string) int {
	set := make(map[coordinates]struct{})
	currentPosition := coordinates{0, 0}
	set[currentPosition] = struct{}{}

	for _, s := range data {
		switch {
		case s == '^':
			currentPosition.y += 1
		case s == '>':
			currentPosition.x += 1
		case s == 'v':
			currentPosition.y -= 1
		case s == '<':
			currentPosition.x -= 1

		}

		set[currentPosition] = struct{}{}
	}

	return len(set)

}
func partTwo(data string) int {
	set := make(map[coordinates]struct{})
	santaPosition := coordinates{0, 0}
	roboPosition := coordinates{0, 0}
	set[santaPosition] = struct{}{}
	currentPosition := &santaPosition

	for i, s := range data {
		if i%2 == 0 {
			currentPosition = &santaPosition
		} else {
			currentPosition = &roboPosition
		}
		switch {
		case s == '^':
			currentPosition.y += 1
		case s == '>':
			currentPosition.x += 1
		case s == 'v':
			currentPosition.y -= 1
		case s == '<':
			currentPosition.x -= 1

		}
		set[*currentPosition] = struct{}{}

	}
	return len(set)

}

func main() {
	fmt.Println(partOne(readFile("./day3.txt")))
	fmt.Println(partTwo(readFile("./day3.txt")))

}
