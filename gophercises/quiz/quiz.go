package main

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

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer")

	problems := readCSVFile(*csvFilename)

	var userAnswer string
	var score int

	for i, problem := range problems {
		fmt.Printf("%v. What is %s?\n", i+1, trim(problem.question))
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == trim(problem.answer) {
			score += 1
		}
	}

	fmt.Printf("Your score is %v out of %v", score, len(problems))
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func trim(str string) string {
	return strings.TrimSpace(str)
}

func readCSVFile(path string) []problem {
	file, err := os.Open(path)
	if err != nil {
		exit(fmt.Sprintf("Failed to read file: %s", path))
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse file: %s", path))
	}

	problems := parseLines(lines)
	file.Close()
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problems
}
