package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day_1_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	calsByElf := []int{}
	currElfCals := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totalCals := 0
			for _, cal := range currElfCals {
				totalCals += cal
			}
			calsByElf = append(calsByElf, totalCals)
			currElfCals = nil
			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}
		currElfCals = append(currElfCals, cals)

		if scanner.Err() != nil {
			log.Fatalln(scanner.Err())
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calsByElf)))

	fmt.Println("Most calories:", calsByElf[0])

	if len(calsByElf) < 3 {
		log.Fatalln("Less than three Elves provided for part two")
	}

	topCals := calsByElf[0] + calsByElf[1] + calsByElf[2]
	fmt.Println("Sum of top three Elves:", topCals)
}
