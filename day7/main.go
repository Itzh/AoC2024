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
	input := parse("input.txt")
	solution := solve_loop(input)

	fmt.Println(solution)

}

func solve_loop(input [][]int64) int64 {
	var sum int64
	for i := 0; i < len(input); i++ {
		if solve(0, 1, input[i]) {
			sum += input[i][0]
		}
	}
	return sum
}

func solve(curr_sum int64, i int, input []int64) bool {
	if curr_sum > input[0] {
		return false
	}
	if i == len(input) {
		return curr_sum == input[0]
	}
	return (solve(curr_sum+input[i], i+1, input) ||
		solve(curr_sum*input[i], i+1, input) ||
		solve(concat(curr_sum, input[i]), i+1, input))
}

func concat(a, b int64) int64 {
	r, _ := strconv.ParseInt(fmt.Sprintf("%d%d", a, b), 10, 64)
	return r
}

func parse(inputPath string) [][]int64 {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("Failed to open")
	}
	defer file.Close()
	var res [][]int64
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		var arr []int64
		for _, s := range line {
			s = strings.Replace(s, ":", "", 1)
			num, _ := strconv.ParseInt(s, 10, 64)
			arr = append(arr, num)
		}
		res = append(res, arr)
	}

	return res
}
