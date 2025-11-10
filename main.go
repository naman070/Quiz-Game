// Package main implements a simple timed quiz game that reads
// questions and answers from a CSV file.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// accumulateData reads a CSV file and returns its contents
// as a slice of string slices.
func accumulateData(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error occurred while opening the file %s : %v\n", filename, err)
	}
	defer file.Close()

	var data [][]string
	r := csv.NewReader(file)
	for {
		record, readErr := r.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			fmt.Printf("Error occurred while reading : %v\n", readErr)
		}
		data = append(data, record)
	}
	return data
}

// shuffleQuestions randomizes the order of the questions in-place.
func shuffleQuestions(data *[][]string) *[][]string {
	// Shuffle the slice in place
	rand.Shuffle(len(*data), func(i, j int) {
		(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
	})
	return data
}

func quizEnd(wrongAnswers *[]string, score int, limitInt int) {
	fmt.Printf("\n------------------------------\n")
	fmt.Printf("You scored %d out of %d\n", score, limitInt)
	if len(*wrongAnswers) > 0 {
		for _, msg := range *wrongAnswers {
			fmt.Println(msg)
		}
	}
}

func main() {
	csvFilename := flag.String("csv", "quiz.csv", "A CSV File in the format of 'questions,answer'")
	questionsLimit := flag.Int("n", 20, "Max Number of questions in Quiz")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
	flag.Parse()
	quizData := accumulateData(*csvFilename)
	var wrongAnswers []string
	var limitInt int = min(*questionsLimit, len(quizData))
	var score int = 0

	shuffleQuestions(&quizData)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for ind := 0; ind < limitInt; ind++ {
		fmt.Printf("Problem #%d: %s = ", (ind + 1), quizData[ind][0])
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scan(&ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\n\nTimer Expired.")
			quizEnd(&wrongAnswers, score, limitInt)
			return
		case ans := <-answerCh:
			if ans == quizData[ind][1] {
				score += 1
			} else {
				msg := fmt.Sprintf("Problem #%d: %s -> You answered %s, correct answer is %s", (ind + 1), quizData[ind][0], ans, quizData[ind][1])
				wrongAnswers = append(wrongAnswers, msg)
			}
		}
	}
	quizEnd(&wrongAnswers, score, limitInt)
}
