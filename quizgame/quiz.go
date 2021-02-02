package quizgame

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
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

// PlayGame is the entry point to quiz game
func PlayGame() {
	csvContent, error := getCSVContent()

	if error != nil {
		exit(fmt.Sprintf("Error loading csv file %d", error))
	}
	problems := mapCSVToProblems(csvContent)

	score := 0
	for _, problem := range problems {
		fmt.Printf("The answer to %s is \n", problem.question)
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == problem.answer {
			score++
		}
	}
	fmt.Println(score)
}
