package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

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

func shuffleQuestions(data *[][]string) *[][]string {
	// Shuffle the slice in place
	rand.Shuffle(len(*data), func(i, j int) {
		(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
	})
	return data
}

func main() {
	quizData := accumulateData("quiz.csv")
	var wrongAnswers []string
	var limitInt int = len(quizData)
	var score int = 0

	if len(os.Args[1:]) > 0 {
		limit, err := strconv.Atoi(os.Args[1:][0])
		if err != nil {
			fmt.Printf("Invalid limit, limit must be integer\n")
		}
		limitInt = limit
	}
	shuffleQuestions(&quizData)
	for ind := 0; ind < limitInt; ind++ {
		var ans string
		fmt.Printf("Problem #%d: %s = ", (ind + 1), quizData[ind][0])
		fmt.Scan(&ans)
		if ans == quizData[ind][1] {
			score += 1
		} else {
			msg := fmt.Sprintf("Problem #%d: %s -> You answered %s, correct answer is %s", (ind + 1), quizData[ind][0], ans, quizData[ind][1])
			wrongAnswers = append(wrongAnswers, msg)
		}
	}
	fmt.Printf("\n------------------------------\n")
	fmt.Printf("You scored %d out of %d\n", score, limitInt)
	if len(wrongAnswers) > 0 {
		for _, msg := range wrongAnswers {
			fmt.Println(msg)
		}
	}
}
