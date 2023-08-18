package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type queue []rune

func (q *queue) enqueue(r rune) {
	*q = append(*q, r)
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func (q *queue) length() int {
	return len(*q)
}

func (q *queue) dequeue() rune {
	if q.empty() {
		log.Fatalln("dequeue of empty queue")
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *queue) uniqueItems() bool {
	set := map[rune]bool{}
	for _, char := range *q {
		if set[char] {
			return false
		}
		set[char] = true
	}
	return true
}

func main() {
	file, err := os.Open("day_6_input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	q := make(queue, 0, 4)
	position := 0

	reader := bufio.NewReader(file)
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		position++

		if q.length() < 4 {
			q.enqueue(char)
		} else {
			q.dequeue()
			q.enqueue(char)
			if q.uniqueItems() {
				fmt.Println("First marker after character", position)
				return
			}
		}
	}

	fmt.Println("No marker found")
}
