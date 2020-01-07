package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var defaultFileName string = "problems.csv"

type Question struct {
	addend1 int
	addend2 int
	sign    string
	result  int
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a CSV file with questions and answers (default problems.csv)")
	// limit := flag.Int("limit", 30, "time limit for quiz in seconds (default 30)")
	flag.Parse()

	questions := ParseCsvFile(*csvFile)
	correctAnswers := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i, q := range questions {
		fmt.Printf("Problem # %d: %d %s %d = \n", i+1, q.addend1, q.sign, q.addend2)
		scanner.Scan()
		answer := scanner.Text()
		if correctAnswer(q, answer) {
			correctAnswers = correctAnswers + 1
		}
	}
	fmt.Printf("You scored %d out of %d", correctAnswers, len(questions))
}

func correctAnswer(q Question, answer string) bool {
	answerNum, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("Error parsing result value:", err)
	}
	return q.addend1+q.addend2 == answerNum
}

func parseRowValues(row []string) Question {
	result, err := strconv.Atoi(row[1])
	if err != nil {
		fmt.Println("Error parsing result value:", row[1])
	}
	addend1, err := strconv.Atoi(strings.Split(row[0], "+")[0])
	if err != nil {
		fmt.Println("Error parsing addend1 value:", row[0])
	}
	addend2, err := strconv.Atoi(strings.Split(row[0], "+")[1])
	if err != nil {
		fmt.Println("Error parsing addend2 value:", row[0])
	}
	q := Question{
		addend1: addend1,
		addend2: addend2,
		sign:    "+",
		result:  result,
	}
	return q
}

func ParseCsvFile(filename string) []Question {
	f, err := os.Open(defaultFileName)
	if err != nil {
		fmt.Println("Error parsing the file:", filename)
		os.Exit(1)
	}
	defer f.Close()
	csvr := csv.NewReader(f)

	var questions []Question
	for {
		row, err := csvr.Read()
		if err != nil && err != io.EOF {
			fmt.Printf("Error while reading line %s in file %s", row, filename)
		}
		if len(row) == 0 {
			break
		}
		questions = append(questions, parseRowValues(row))
	}
	return questions
}
