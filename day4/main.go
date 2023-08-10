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
	return (a.start <= b.start && a.end >= b.end) || (b.start <= a.start && b.end >= a.end)
}

func overlaps(a, b interval) bool {
	return a.start <= b.end && a.end >= b.start
}

func main() {
	file, err := os.Open("day_4_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	containsCount := 0
	overlapsCount := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		pair := strings.Split(s.Text(), ",")
		firstInterval := convertToInterval(pair[0])
		secondInterval := convertToInterval(pair[1])
		if contains(firstInterval, secondInterval) {
			containsCount++
		}
		if overlaps(firstInterval, secondInterval) {
			overlapsCount++
		}
	}
	fmt.Println("Part one:", containsCount, "pairs fully contain one another")
	fmt.Println("Part two:", overlapsCount, "pairs overlap")
}

func convertToInterval(intervalStr string) interval {
	intervalLimits := strings.Split(intervalStr, "-")
	if len(intervalLimits) != 2 {
		log.Fatalf("malformed pair: %v\n", intervalStr)
	}

	start, err := strconv.Atoi(intervalLimits[0])
	if err != nil {
		log.Fatalf("malformed start value: %v", err)
	}

	end, err := strconv.Atoi(intervalLimits[1])
	if err != nil {
		log.Fatalf("malformed end value: %v", err)
	}

	return interval{start: start, end: end}
}
