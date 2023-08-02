package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	beats := map[string]string{
		"scissors": "rock",
		"rock":     "paper",
		"paper":    "scissors",
	}

	points := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	rival := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	mine := map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	file, err := os.Open("day_2_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	myScore := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rivalMove := rival[line[:1]]
		myMove := mine[line[2:]]

		myScore += points[myMove]

		if myMove == beats[rivalMove] {
			myScore += 6
		} else if myMove == rivalMove {
			myScore += 3
		}
	}

	if scanner.Err() != nil {
		log.Fatalln(scanner.Err())
	}

	fmt.Println("My total score would be:", myScore)
}
