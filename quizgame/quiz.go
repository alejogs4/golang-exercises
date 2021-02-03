package quizgame

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

// PlayGame is the entry point to quiz game
func PlayGame() {
	timeLimit := flag.Int("Time", 20, "Time in seconds to complete the quiz")
	flag.Parse()

	csvContent, error := getCSVContent()

	if error != nil {
		exit(fmt.Sprintf("Error loading csv file %d", error))
	}
	problems := mapCSVToProblems(csvContent)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	score := 0

	for _, problem := range problems {
		fmt.Printf("The answer to %s is \n", problem.question)

		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Score", score)
			return
		case answer := <-answerChannel:
			if answer == problem.answer {
				score++
			}
		}
	}
	fmt.Println(score)
}

func getCSVContent() ([][]string, error) {
	csvFilename := flag.String("csv", "problems.csv", "Route to csv route with format ('problem', 'answer')")
	flag.Parse()

	file, error := os.Open(*csvFilename)
	defer file.Close()

	if error != nil {
		return nil, error
	}

	csv := csv.NewReader(file)
	csvContent, error := csv.ReadAll()

	return csvContent, error
}

func mapCSVToProblems(csvContent [][]string) []problem {
	problems := make([]problem, len(csvContent))
	for i, line := range csvContent {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
