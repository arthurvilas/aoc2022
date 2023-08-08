package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	start, end int
}

func contains(a, b interval) bool {
	return a.start <= b.start && a.end >= b.end
}

func main() {
	file, err := os.Open("day_4_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	partOne(bufio.NewScanner(file))
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalln(bufio.NewScanner(file))
	}
}

func partOne(s *bufio.Scanner) {
	containsCount := 0
	for s.Scan() {
		pair := strings.Split(s.Text(), ",")
		firstInterval := convertToInterval(pair[0])
		secondInterval := convertToInterval(pair[1])
		if contains(firstInterval, secondInterval) || contains(secondInterval, firstInterval) {
			containsCount++
		}
	}
	fmt.Println(containsCount, "pairs fully contain one another")
}

func convertToInterval(intervalStr string) interval {
	intervalLimits := strings.Split(intervalStr, "-")
	res := interval{}
	for i, limit := range intervalLimits {
		if num, err := strconv.Atoi(limit); err != nil {
			log.Fatalf("malformed pair at index %d: %v", i, err)
		} else {
			if i == 0 {
				res.start = num
			} else if i == 1 {
				res.end = num
			}
		}
	}
	return res
}
