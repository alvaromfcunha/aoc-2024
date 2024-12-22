package main

import (
	"bufio"
	"fmt"
	"os"
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

	if len(*col1) != len(*col2) {
		panic("length of col1 differs from col2")
	}

	freq := make([]int, len(*col1))
	for i, n1 := range *col1 {
		f := 0
		for _, n2 := range *col2 {
			if n1 == n2 {
				f++
			}
		}
		freq[i] = f
	}

	scoreList := make([]int, len(freq))
	for i, f := range freq {
		scoreList[i] = (*col1)[i] * f
	}

	score := 0
	for _, s := range scoreList {
		score += s
	}

	fmt.Println(score)
}
