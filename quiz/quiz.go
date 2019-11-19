package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Problem -
type Problem struct {
	question string
	answer   string
}

func parseProblems(text string) []Problem {
	lineCount := strings.Count(text, "\n")
	problems := make([]Problem, lineCount)

	for i := 0; i < lineCount; i++ {
		newLine := strings.Index(text, "\n")
		line := text[:newLine]

		comma := strings.LastIndex(line, ",")
		problem := Problem{
			question: line[:comma],
			answer:   line[comma+1:],
		}

		text = text[newLine+1:]
		problems[i] = problem
	}
	return problems
}

func randomized(p []Problem) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(p); i++ {
		rand := rand.Intn(i + 1)
		if i != rand {
			p[rand], p[i] = p[i], p[rand]
		}
	}
}
func solveQuiz(timeLimit int, problems []Problem) {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("problem #%d: %s = ", i+1, p.question)
		scoreChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answer = strings.TrimSpace(answer)
			scoreChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nFinal score: %d / %d\n", correct, len(problems))
			return
		case answer := <-scoreChan:
			if answer == p.answer {
				correct++
			}
		}
	}

}
func main() {
	filePath := flag.String("problem", "problems.csv", "Path to CSV problems file")
	timeLimit := flag.Int("time", 5, "Time limit in seconds.")
	flag.Parse()

	problem, err := ioutil.ReadFile(*filePath)
	if err != nil {
		os.Exit(3)
	}

	problems := parseProblems(string(problem))
	randomized(problems)
	solveQuiz(*timeLimit, problems)

}
