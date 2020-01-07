package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

var defaultFileName string = "problems.csv"

type Problem struct {
	question string
	result   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a CSV file with questions and answers (default problems.csv)")
	// limit := flag.Int("limit", 30, "time limit for quiz in seconds (default 30)")
	flag.Parse()

	problems := ParseCsvFile(*csvFile)
	correct := 0
	scanner := bufio.NewScanner(os.Stdin)
	for i, p := range problems {
		fmt.Printf("Problem # %d: %s = \n", i+1, p.question)
		scanner.Scan()
		answer := scanner.Text()
		if answer == p.result {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func parseRowValues(row []string) Problem {
	p := Problem{
		question: row[0],
		result:   row[1],
	}
	return p
}

func ParseCsvFile(filename string) []Problem {
	f, err := os.Open(defaultFileName)
	if err != nil {
		fmt.Println("Error parsing the file:", filename)
		os.Exit(1)
	}
	defer f.Close()
	csvr := csv.NewReader(f)

	var problems []Problem
	for {
		row, err := csvr.Read()
		if err != nil && err != io.EOF {
			fmt.Printf("Error while reading line %s in file %s", row, filename)
		}
		if len(row) == 0 {
			break
		}
		problems = append(problems, parseRowValues(row))
	}
	return problems
}
