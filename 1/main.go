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
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left, right, difference []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		first, second, _ := strings.Cut(scanner.Text(), " ")

		firstAsNum, _ := strconv.Atoi(strings.TrimSpace(first))
		secondAsNum, _ := strconv.Atoi(strings.TrimSpace(second))

		left = append(left, firstAsNum)
		right = append(right, secondAsNum)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for i := range left {
		if left[i] > right[i] {
			difference = append(difference, left[i]-right[i])
		} else {
			difference = append(difference, right[i]-left[i])
		}
	}

	total := 0

	for _, num := range difference {
		total = total + num
	}

	fmt.Println(total)
}
