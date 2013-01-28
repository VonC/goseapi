package stackexchange

import (
	"encoding/json"
	"strconv"
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
	Score            int    `json:"score"`
	Title            string `json:"title"`
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
	Title      string `json:"title"`
}

// Error is an error returned from the Stack Exchange API wrapper.
//
// See: https://api.stackexchange.com/docs/wrapper
type Error struct {
	ID      int    `json:"error_id"`
	Name    string `json:"error_name"`
	Message string `json:"error_message"`
}

func (err *Error) Error() string {
	return err.Message + " (" + strconv.Itoa(err.ID) + " " + err.Name + ")"
}

// Wrapper records the common fields in the JSON wrapper.  It is intended to be
// embedded, since the "items" key is specifically not included.
//
// See: https://api.stackexchange.com/docs/wrapper
//
// BUG(light): encoding/json does not support embedded structs in Go 1.0, so
// embedding doesn't work as expected yet.
type Wrapper struct {
	Error

	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasMore  bool `json:"has_more"`

	Backoff        int `json:"backoff"`
	QuotaMax       int `json:"quota_max"`
	QuotaRemaining int `json:"quota_remaining"`

	Total int    `json:"total"`
	Type  string `json:"type"`
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
