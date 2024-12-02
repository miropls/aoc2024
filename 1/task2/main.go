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
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		first, second, _ := strings.Cut(scanner.Text(), " ")

		firstAsNum, _ := strconv.Atoi(strings.TrimSpace(first))
		secondAsNum, _ := strconv.Atoi(strings.TrimSpace(second))

		left = append(left, firstAsNum)
		right = append(right, secondAsNum)
	}

	for i, lNum := range left {
		timesAppearingOnRight := 0

		for _, rNum := range right {
			if lNum == rNum {
				timesAppearingOnRight += 1
			}
		}

		left[i] = lNum * timesAppearingOnRight
	}

	similarityScore := 0

	for _, num := range left {
		similarityScore += num
	}

	fmt.Println(similarityScore)
}
