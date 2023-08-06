package main

import (
	"bufio"
	"errors"
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

	partOne(bufio.NewScanner(file))
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalln(bufio.NewScanner(file))
	}
	partTwo(bufio.NewScanner(file))
}

func partOne(s *bufio.Scanner) {
	prioritiesSum := 0
	for s.Scan() {
		rucksack := s.Text()
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]

		chars := map[rune]bool{}

		for _, c := range compartment1 {
			chars[c] = true
		}

		var common rune

		for _, c := range compartment2 {
			if chars[c] {
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

	if err := s.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Sum of pritorities:", prioritiesSum)
}

func partTwo(s *bufio.Scanner) {
	badges := []rune{}
	group := make([]string, 0, 3)
	for s.Scan() {
		if len(group) < cap(group) {
			group = append(group, s.Text())
		}

		if len(group) == cap(group) {
			badge, err := processGroup(group)
			if err != nil {
				log.Fatalln(err)
			}
			badges = append(badges, badge)
			group = group[:0]
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalln(err)
	}

	badgesSum := 0
	for _, b := range badges {
		if unicode.IsUpper(b) {
			badgesSum += int(b) - 'A' + 27
		} else {
			badgesSum += int(b) - 'a' + 1
		}
	}

	fmt.Println("Sum of badges:", badgesSum)
}

func processGroup(g []string) (rune, error) {
	map1 := map[rune]bool{}

	for _, c := range g[0] {
		map1[c] = true
	}

	common := map[rune]bool{}

	for _, c := range g[1] {
		if map1[c] {
			common[c] = true
		}
	}

	for _, c := range g[2] {
		if common[c] {
			return c, nil
		}
	}

	return ' ', errors.New("no common rune in group")
}
