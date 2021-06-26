package main

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

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer")
	timeLimit := flag.Int("limit", 30, "The timeLimit for the quiz in secconds")

	problems := readCSVFile(*csvFilename)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var score int

	for i, problem := range problems {
		fmt.Printf("%v. What is %s?\n", i+1, trim(problem.question))
		answerChannel := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			answerChannel <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Your score is %v out of %v.", score, len(problems))
			return
		case <-answerChannel:
			if <-answerChannel == trim(problem.answer) {
				score += 1
			}
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
	return problems
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
