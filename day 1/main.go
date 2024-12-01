package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	l1, l2 := parse("input_p2.txt")
	fmt.Println("ans p1:", part_one(l1, l2))
	fmt.Println("ans p2:", part_two(l1, l2))
}

func part_two(l1 []int, l2 []int) int {
	sum := 0
	var count int
	for i := 0; i < len(l1); i++ {
		num := l1[i]
		count = 0
		for j := 0; j < len(l2); j++ {
			if l2[j] == num {
				count++
			}
		}
		sum += num * count
	}

	return sum
}

func part_one(l1 []int, l2 []int) int {
	sum := 0
	sort.Ints(l1)
	sort.Ints(l2)
	var d int
	for i := 0; i < len(l1); i++ {
		d = l1[i] - l2[i]
		if d < 0 {
			d *= -1
		}
		sum += d
	}
	return sum
}

func parse(inputPath string) ([]int, []int) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal("Failed to open")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var l1 []int
	var l2 []int

	for fileScanner.Scan() {
		s := strings.Fields(fileScanner.Text())
		num, _ := strconv.Atoi(s[0])
		l1 = append(l1, num)
		num, _ = strconv.Atoi(s[1])
		l2 = append(l2, num)
	}
	return l1, l2
}
