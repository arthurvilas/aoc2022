package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var beats = map[string]string{
	"scissors": "rock",
	"rock":     "paper",
	"paper":    "scissors",
}

var movePoints = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var rivalMap = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

func main() {

	file, err := os.Open("day_2_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	partOne(bufio.NewScanner(file))

	// Reset file pointer
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	partTwo(bufio.NewScanner(file))
}

func partOne(s *bufio.Scanner) {
	myMap := map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	myScore := 0
	for s.Scan() {
		line := s.Text()
		rivalMove := rivalMap[line[:1]]
		myMove := myMap[line[2:]]
		myScore += movePoints[myMove]
		if myMove == beats[rivalMove] {
			myScore += 6
		} else if myMove == rivalMove {
			myScore += 3
		}
	}

	if s.Err() != nil {
		log.Fatalln(s.Err())
	}

	fmt.Println("My total score would be:", myScore)
}

func partTwo(s *bufio.Scanner) {
	myScore := 0
	for s.Scan() {
		line := s.Text()
		rivalMove := rivalMap[line[:1]]
		outcome := line[2:]
		var myMove string
		if outcome == "Z" {
			myMove = beats[rivalMove]
			myScore += 6
		} else if outcome == "Y" {
			myMove = rivalMove
			myScore += 3
		} else {
			myMove = beats[beats[rivalMove]]
			myScore += 0
		}

		myScore += movePoints[myMove]
	}

	if s.Err() != nil {
		log.Fatalln(s.Err())
	}

	fmt.Println("My total score would be:", myScore)
}
