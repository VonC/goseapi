package main

import (
	"bitbucket.org/zombiezen/stackexchange"
	"fmt"
	"log"
)

func main() {
	if err := scrapeQuestions(); err != nil {
		log.Fatal(err)
	}
}

func scrapeQuestions() error {
	var questions []stackexchange.Question
	_, err := stackexchange.Do("/questions", &questions, stackexchange.Params{
		Site:     stackexchange.StackOverflow,
		Sort:     stackexchange.SortScore,
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

func fetchAnswers(id int) ([]stackexchange.Answer, error) {
	var answers []stackexchange.Answer
	_, err := stackexchange.Do(fmt.Sprintf("/questions/%d/answers", id), &answers, stackexchange.Params{
		Site:     stackexchange.StackOverflow,
		Sort:     stackexchange.SortScore,
		Order:    "desc",
		Filter:   "!-u2CTCBE",
		PageSize: 1,
	})
	return answers, err
}
