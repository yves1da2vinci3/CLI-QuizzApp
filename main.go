package main

import (
	"fmt"
	"time"
)

type Quiz struct {
	Question        string
	PossibleAnswers []map[string]string
	CorrectAnswer   string
}

func main() {
	quizList := []Quiz{
		{
			Question: "What is the capital of France?",
			PossibleAnswers: []map[string]string{
				{"A": "London"},
				{"B": "Paris"},
				{"C": "Berlin"},
			},
			CorrectAnswer: "B",
		},
		{
			Question: "Who painted the famous artwork 'The Starry Night'?",
			PossibleAnswers: []map[string]string{
				{"A": "Vincent van Gogh"},
				{"B": "Pablo Picasso"},
				{"C": "Leonardo da Vinci"},
			},
			CorrectAnswer: "A",
		},
		{
			Question: "Which of the following is not a programming language?",
			PossibleAnswers: []map[string]string{
				{"A": "Python"},
				{"B": "Java"},
				{"C": "HTML"},
			},
			CorrectAnswer: "C",
		},
	}

	fmt.Println("Welcome to the quiz!")
	fmt.Println("You have 10 seconds to answer each question.\n")

	score := 0
	for i, quiz := range quizList {
		fmt.Printf("Question #%d: %s\n", i+1, quiz.Question)

		// Print possible answers
		for _, answer := range quiz.PossibleAnswers {
			for letter, text := range answer {
				fmt.Printf("%s) %s\n", letter, text)
			}
		}

		// Start timer
		timer := time.NewTimer(10 * time.Second)

		// Listen for answer input
		ch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			ch <- answer
		}()

		// Wait for answer or timeout
		select {
		case answer := <-ch:
			if answer == quiz.CorrectAnswer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Printf("Incorrect. The correct answer is %s.\n", quiz.CorrectAnswer)
			}
		case <-timer.C:
			fmt.Printf("Time's up! The correct answer is %s.\n", quiz.CorrectAnswer)
		}

		fmt.Println()
	}

	fmt.Printf("Quiz complete! You scored %d out of %d.\n", score, len(quizList))
}
