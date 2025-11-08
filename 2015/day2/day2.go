package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []string {

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	strData := strings.TrimSpace(string(data))
	strSlice := strings.Split(strData, "\n")
	return strSlice
}
func returnSmallestPart(data []int) int {
	lw := data[0] * data[1]
	wh := data[1] * data[2]
	hl := data[2] * data[0]

	if lw <= wh && lw <= hl {
		return lw
	} else if wh <= lw && wh <= hl {
		return wh
	} else {
		return hl
	}

}
func partOne(data []string) int {
	var result int = 0
	for _, elem := range data {
		elemSlice := strings.Split(elem, "x")
		if len(elemSlice) < 3 {
			continue
		}
		intSlice := make([]int, len(elemSlice))
		for i, s := range elemSlice {
			n, e := strconv.Atoi(s)
			if e != nil {
				fmt.Println(len(elemSlice))
				fmt.Println(elemSlice)

				panic(e)
			}
			intSlice[i] = n
		}
		smallestPart := returnSmallestPart(intSlice)
		result += 2*intSlice[0]*intSlice[1] + 2*intSlice[1]*intSlice[2] + 2*intSlice[2]*intSlice[0] + smallestPart

	}

	return result
}
func returnSumOfSides(data []int) int {
	l := data[0]
	w := data[1]
	h := data[2]
	lw := l + l + w + w
	wh := w + w + h + h
	hl := h + h + l + l
	if lw <= wh && lw <= hl {
		return lw
	} else if wh <= lw && wh <= hl {
		return wh
	} else {
		return hl
	}
}
func partTwo(data []string) int {
	var result int = 0
	for _, elem := range data {
		elemSlice := strings.Split(elem, "x")
		if len(elemSlice) < 3 {
			continue
		}
		intSlice := make([]int, len(elemSlice))
		for i, s := range elemSlice {
			n, e := strconv.Atoi(s)
			if e != nil {
				fmt.Println(len(elemSlice))
				fmt.Println(elemSlice)

				panic(e)
			}
			intSlice[i] = n
		}
		sumOfSides := returnSumOfSides(intSlice)
		result += (intSlice[0] * intSlice[1] * intSlice[2]) + sumOfSides
	}
	return result
}

func main() {
	fmt.Println(partOne(readFile("./day2.txt")))
	fmt.Println(partTwo(readFile("./day2.txt")))

}
