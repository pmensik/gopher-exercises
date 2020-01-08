package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type Problem struct {
	question string
	result   string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a CSV file with questions and answers")
	limit := flag.Int("limit", 30, "time limit for quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle questions?")
	flag.Parse()

	problems := ParseCsvFile(*csvFile, *shuffle)
	correct := 0
	qChannel := make(chan bool, 1)

	go startQuestionnaire(problems, &correct, qChannel)
	select {
	case <-qChannel:
		fmt.Println("Succesfully finished test in a deadline:")
	case <-time.After(time.Duration(*limit) * time.Second):
		fmt.Println("Deadline occured")
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func startQuestionnaire(problems []Problem, correct *int, ch chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for i, p := range problems {
		fmt.Printf("Problem # %d: %s = \n", i+1, p.question)
		scanner.Scan()
		answer := scanner.Text()
		if answer == p.result {
			*correct++
		}
	}
	ch <- true
}

func parseRowValues(row []string) Problem {
	p := Problem{
		question: row[0],
		result:   row[1],
	}
	return p
}

func ParseCsvFile(filename string, shuffle bool) []Problem {
	f, err := os.Open(filename)
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
	if shuffle {
		return Shuffle(problems)
	} else {
		return problems
	}
}

func Shuffle(problems []Problem) []Problem {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Problem, len(problems))
	perm := r.Perm(len(problems))
	for i, randIndex := range perm {
		ret[i] = problems[randIndex]
	}
	return ret
}
