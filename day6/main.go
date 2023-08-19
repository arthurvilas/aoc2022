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
	findMarker(file, 4)
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalln(err)
	}
	findMarker(file, 14)
}

func findMarker(file *os.File, size int) {
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

		if q.length() < size {
			q.enqueue(char)
		} else {
			q.dequeue()
			q.enqueue(char)
			if q.uniqueItems() {
				fmt.Printf("First marker of len %v after character %v\n", size, position)
				return
			}
		}
	}

	fmt.Println("No marker found")
}
