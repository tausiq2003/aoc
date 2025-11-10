package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	var num int = 1
	for {
		input := "bgvyzdsv" + strconv.Itoa(num)
		data := []byte(input)
		hash := md5.Sum([]byte(data))
		result := hex.EncodeToString(hash[:])
		// if result[:5] == "00000" { part one
		if result[:5] == "000000" { // this will take lots of time i need to read golang more
			break
		}
		num++

	}
	fmt.Println(num)

}
