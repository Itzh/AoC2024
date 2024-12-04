package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputs := parse("input.txt")
	fmt.Println(solve_p1(inputs))
	fmt.Println(solve_p2(inputs))
}

func solve_p2(input []string) int {
	res := 0
	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input)-1; x++ {
			if input[y][x] == 'A' {
				if input[y-1][x-1] == 'M' && input[y-1][x+1] == 'S' && input[y+1][x-1] == 'M' && input[y+1][x+1] == 'S' {
					res++
				}
				if input[y-1][x-1] == 'S' && input[y-1][x+1] == 'M' && input[y+1][x-1] == 'S' && input[y+1][x+1] == 'M' {
					res++
				}
				if input[y-1][x-1] == 'M' && input[y-1][x+1] == 'M' && input[y+1][x-1] == 'S' && input[y+1][x+1] == 'S' {
					res++
				}
				if input[y-1][x-1] == 'S' && input[y-1][x+1] == 'S' && input[y+1][x-1] == 'M' && input[y+1][x+1] == 'M' {
					res++
				}
			}
		}
	}
	return res
}

func solve_p1(input []string) int {
	res := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == 'X' {
				res += findMas_p1(y, x, input)
			}
		}
	}
	return res
}

func findMas_p1(y, x int, input []string) int {
	max := len(input)
	res := 0
	// upright
	if y-3 >= 0 && x+3 < max && input[y-1][x+1] == 'M' && input[y-2][x+2] == 'A' && input[y-3][x+3] == 'S' {
		res++
	}
	//upleft
	if y-3 >= 0 && x-3 >= 0 && input[y-1][x-1] == 'M' && input[y-2][x-2] == 'A' && input[y-3][x-3] == 'S' {
		res++
	}
	//downright
	if y+3 < max && x+3 < max && input[y+1][x+1] == 'M' && input[y+2][x+2] == 'A' && input[y+3][x+3] == 'S' {
		res++
	}
	//downleft
	if y+3 < max && x-3 >= 0 && input[y+1][x-1] == 'M' && input[y+2][x-2] == 'A' && input[y+3][x-3] == 'S' {
		res++
	}
	//horisontal
	if x+3 < max && input[y][x+1] == 'M' && input[y][x+2] == 'A' && input[y][x+3] == 'S' {
		res++
	}
	//horisontal-rev
	if x-3 >= 0 && input[y][x-1] == 'M' && input[y][x-2] == 'A' && input[y][x-3] == 'S' {
		res++
	}
	//vertical
	if y+3 < max && input[y+1][x] == 'M' && input[y+2][x] == 'A' && input[y+3][x] == 'S' {
		res++
	}
	//vertical-rev
	if y-3 >= 0 && input[y-1][x] == 'M' && input[y-2][x] == 'A' && input[y-3][x] == 'S' {
		res++
	}

	return res
}

func parse(inputPath string) []string {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("Failed to open")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var input []string
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	return input
}
