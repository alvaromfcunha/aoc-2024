package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	col1 := new([]int)
	col2 := new([]int)
	for s.Scan() {
		numbers := strings.Split(s.Text(), "   ")
		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic("a line contain a non numeric first number")
		}
		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic("a line contain a non numeric second number")
		}

		*col1 = append(*col1, n1)
		*col2 = append(*col2, n2)
	}

	sort.Slice(*col1, func(i, j int) bool {
		return (*col1)[i] < (*col1)[j]
	})

	sort.Slice(*col2, func(i, j int) bool {
		return (*col2)[i] < (*col2)[j]
	})

	if len(*col1) != len(*col2) {
		panic("length of col1 differs from col2")
	}

	colDiff := make([]int, len(*col1))
	for i := 0; i < len(*col1); i++ {
		colDiff[i] = AbsInt((*col1)[i] - (*col2)[i])
	}

	dist := 0
	for i := 0; i < len(colDiff); i++ {
		dist += colDiff[i]
	}

	fmt.Println(dist)
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
