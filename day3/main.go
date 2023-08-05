package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("day_3_input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	prioritiesSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]

		chars1 := map[rune]bool{}
		chars2 := map[rune]bool{}

		for _, c := range compartment1 {
			chars1[c] = true
		}

		for _, c := range compartment2 {
			chars2[c] = true
		}

		var common rune

		for c := range chars1 {
			if chars2[c] {
				common = c
				break
			}
		}

		if unicode.IsUpper(common) {
			prioritiesSum += int(common) - 'A' + 27
		} else {
			prioritiesSum += int(common) - 'a' + 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Sum of pritorities:", prioritiesSum)
}
