package main

import (
	"bitbucket.org/zombiezen/stackexchange"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := scrapeQuestions(); err != nil {
		log.Fatal(err)
	}
}

func scrapeQuestions() error {
	resp, err := http.Get(stackexchange.Root + "/questions?order=desc&sort=votes&site=stackoverflow&pagesize=5")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct{
		Questions []stackexchange.Question `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	for _, question := range result.Questions {
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
	resp, err := http.Get(fmt.Sprintf("%s/questions/%d/answers?order=desc&sort=votes&site=stackoverflow&filter=!-u2CTCBE&pagesize=1", stackexchange.Root, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct{
		Answers []stackexchange.Answer `json:"items"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.Answers, err
}
