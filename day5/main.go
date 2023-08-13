package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type stack []string

func (s *stack) push(data string) {
	*s = append(*s, data)
}

func (s *stack) pop() (string, error) {
	if len(*s) == 0 {
		return "", errors.New("empty stack")
	}
	element := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return element, nil
}

func (s *stack) peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

func main() {
	file, err := os.Open("day_5_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	stacks := getInitialStacks(s)
	for s.Scan() {
		command := strings.Fields(s.Text())
		if len(command) < 6 {
			log.Fatalln("malformed command")
		}

		times, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatalln("command conversion error")
		}
		from, err := strconv.Atoi(command[3])
		if err != nil {
			log.Fatalln("command conversion error")
		}
		to, err := strconv.Atoi(command[5])
		if err != nil {
			log.Fatalln("command conversion error")
		}

		for i := 0; i < times; i++ {
			popped, err := stacks[from-1].pop()
			if err != nil {
				log.Fatalln(err)
			}
			stacks[to-1].push(popped)
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack.peek())
	}
	fmt.Println()
}

func getInitialStacks(s *bufio.Scanner) []stack {
	stacks := []stack{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		crates := []string{}
		for i := 0; i < len(line); i += 3 {
			i++
			crates = append(crates, line[i:i+1])
		}
		for len(stacks) < len(crates) {
			stacks = append(stacks, stack{})
		}
		for i, crate := range crates {
			if unicode.IsUpper(rune(crate[0])) {
				stacks[i] = append(stack{crate}, stacks[i]...)
			}
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalln(err)
	}

	return stacks
}
