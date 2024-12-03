package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	text := parse("input.txt")
	fmt.Println("p1: ", solve_p1(text))
	fmt.Println("p2: ", solve_p2(text))
}

func solve_p1(input string) int {
	res := 0
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	numRegex, _ := regexp.Compile(`[0-9]+`)
	mulOps := r.FindAllString(input, -1)
	for _, op := range mulOps {
		nums := numRegex.FindAllString(op, 2)
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		res += n1 * n2
	}
	return res
}

func solve_p2(input string) int {
	res := 0
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	numRegex, _ := regexp.Compile(`[0-9]+`)
	enableRegex, _ := regexp.Compile(`do\(\)`)
	disableRegex, _ := regexp.Compile(`don't\(\)`)
	mulOps := r.FindAllString(input, -1)
	mulEnabled := true
	for _, op := range mulOps {
		if enableRegex.FindString(op) != "" {
			mulEnabled = true
			continue
		} else if disableRegex.FindString(op) != "" {
			mulEnabled = false
			continue
		} else if mulEnabled {
			nums := numRegex.FindAllString(op, 2)
			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])
			res += n1 * n2
		}
	}
	return res
}

func parse(inputPath string) string {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("Failed to open")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Scan()

	return fileScanner.Text()
}
