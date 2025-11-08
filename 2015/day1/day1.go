package main

import (
	//	"bufio"
	"fmt"
	//	"io"
	"os"
)

func readFile(path string) string {

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func partOne(data string) int {
	// if ( then +1 else -1
	var cnt int = 0
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "(" {
			cnt += 1
		} else if string(data[i]) == ")" {
			cnt -= 1
		}

	}
	return cnt

}
func partTwo(data string) int {
	var cnt int = 0
	for i := 0; i < len(data); i++ {
		if cnt == -1 {
			return i
		}
		if string(data[i]) == "(" {
			cnt += 1

		} else if string(data[i]) == ")" {
			cnt -= 1
		}
	}
	return 0
}

func main() {
	fmt.Println(partOne(readFile("./day1.txt")))
	fmt.Println(partTwo(readFile("./day1.txt")))

}
