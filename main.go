package main

import (
	"fmt"

	"./process"
)

func main() {
	var table = [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	for _, subTable := range table {
		fmt.Println(process.Xor(subTable))
	}
}
