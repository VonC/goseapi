package stackexchange

import (
	"encoding/json"
	"time"
)

// Question represents a question on one of the Stack Exchange sites.
//
// https://api.stackexchange.com/docs/types/question
type Question struct {
	ID               int    `json:"question_id"`
	AcceptedAnswerID int    `json:"accepted_answer_id"`
	AnswerCount      int    `json:"answer_count"`
	ClosedDate       Time   `json:"closed_date"`
	ClosedReason     string `json:"closed_reason"`
	Created          Time   `json:"creation_date"`
	IsAnswered       bool   `json:"is_answered"`
	Link             string `json:"link"`
}

// Answer represents an answer to a question on one of the Stack Exchange sites.
//
// https://api.stackexchange.com/docs/types/answer
type Answer struct {
	ID         int    `json:"answer_id"`
	Body       string `json:"body"`
	Created    Time   `json:"creation_date"`
	IsAccepted bool   `json:"is_accepted"`
	Link       string `json:"link"`
	QuestionID int    `json:"question_id"`
	Score      int    `json:"score"`
}

// Time converts Stack Exchange API time values to/from JSON into Go time values.
// Stack Exchange stores time values over the wire as int64 Unix epoch time.
//
// See: https://api.stackexchange.com/docs/dates
type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix())
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var tt int64
	if err := json.Unmarshal(data, &tt); err != nil {
		return err
	}
	*t = Time(time.Unix(tt, 0))
	return nil
}