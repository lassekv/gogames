package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Question string
	Answer   string
}

type Result struct {
	right int
	wrong int
}

var fileName = flag.String("problems", "2problems.csv", "specify the file to load problems from")
var shuffle = flag.Bool("shuffle", false, "should the questions be shuffled?")
var duration = flag.Int("duration", 10, "duration on the quiz")

func parseAnswers(in <-chan string, result *Result, whenDone chan<- string) {
	for res := range in {
		if res == "YES" {
			result.right += 1
		} else if res == "NO" {
			result.wrong += 1
		} else {
			log.Fatal("Invalid answer transmitted")
		}
	}
	whenDone <- "No more questions"
}

func main() {
	parseCLIParameters()
	quiz := readQuizFromFile(*fileName)
	shuffleIfNeeded(quiz)

	answers := make(chan string)
	doneEarly := make(chan string)
	go runQuiz(quiz, answers)

	result := Result{}
	go parseAnswers(answers, &result, doneEarly)
	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	select {
	case <-timer.C:
		fmt.Println("Time is up")
	case <-doneEarly:
		if !timer.Stop() {
			<-timer.C
		}
	}
	fmt.Printf("\nYou answered %d question(s) correct and %d wrong.", result.right, result.wrong)
}

func runQuiz(quiz []Question, answerOut chan<- string) {
	for _, q := range quiz {
		fmt.Printf("What is %v?\n", q.Question)
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer == q.Answer {
			answerOut <- "YES"
		} else {
			answerOut <- "NO"
		}
	}
	close(answerOut)
}

func shuffleIfNeeded(quiz []Question) {
	if *shuffle {
		rand.Shuffle(len(quiz), func(i, j int) { quiz[i], quiz[j] = quiz[j], quiz[i] })
	}
}

func parseCLIParameters() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	ok := flag.Parsed()
	if !ok {
		fmt.Println("There was a problem parsing the command line arguments!")
		os.Exit(1)
	}
}

func readQuizFromFile(filename string) []Question {
	var results = make([]Question, 0)
	csvFile, err := os.Open(filename)

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		b := Question{
			Question: record[0],
			Answer:   record[1],
		}
		results = append(results, b)
	}
	return results
}
