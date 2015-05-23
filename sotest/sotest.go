package main

import (
	"fmt"
	"log"

	"github.com/VonC/goseapi"
)

func main() {
	if err := scrapeQuestions(); err != nil {
		log.Fatal(err)
	}
}

func scrapeQuestions() error {
	var questions []goseapi.Question
	_, err := goseapi.Do("/questions", &questions, goseapi.Params{
		Site:     goseapi.StackOverflow,
		Sort:     goseapi.SortScore,
		Order:    "desc",
		PageSize: 5,
	})
	if err != nil {
		return err
	}

	for _, question := range questions {
		fmt.Printf("%s (ID=%d)\n", question.Title, question.ID)
		answers, err := fetchAnswers(question.ID)
		if err != nil {
			log.Println("Error fetching answers:", err)
		}
		for _, answer := range answers {
			fmt.Printf("  %d %s\n", answer.Score, answer.Body)
		}
	}
	return nil
}

func fetchAnswers(id int) ([]goseapi.Answer, error) {
	var answers []goseapi.Answer
	_, err := goseapi.Do(fmt.Sprintf("/questions/%d/answers", id), &answers, goseapi.Params{
		Site:     goseapi.StackOverflow,
		Sort:     goseapi.SortScore,
		Order:    "desc",
		Filter:   "!-u2CTCBE",
		PageSize: 1,
	})
	return answers, err
}
