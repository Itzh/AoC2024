package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := parse("input.txt")
	fmt.Println("ans:", solve(list))
}

func solve(input [][]int) int {
	correct_lines := 0
	for _, row := range input {
		test := test_input(row)
		if test == 1 {
			correct_lines++
			continue
		}
		for i := 0; i < len(row); i++ {
			cpy := make([]int, len(row))
			copy(cpy, row)
			cpy = append(cpy[:i], cpy[i+1:]...)
			test := test_input(cpy)
			if test == 1 {
				correct_lines++
				break
			}
		}
	}
	return correct_lines
}

func test_input(input []int) int {
	increasing := input[1] > input[0]

	correct := 0
	for j := 0; j < len(input)-1; j++ {
		num := input[j+1] - input[j]
		num2 := num * num
		if num2 > 0 && num2 <= 9 && (increasing && num > 0 || !increasing && num < 0) {
			correct++
		}
	}
	if correct == len(input)-1 {
		return 1
	}
	return 0
}

func parse(inputPath string) [][]int {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("Failed to open")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var list [][]int

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		var arr []int
		for _, s := range line {
			num, _ := strconv.Atoi(s)
			arr = append(arr, num)
		}
		list = append(list, arr)

	}
	return list
}
