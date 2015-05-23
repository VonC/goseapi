package goseapi

import (
	"encoding/json"
	"strconv"
	"time"
)

// Question represents a question on one of the Stack Exchange sites.
//
// https://api.stackexchange.com/docs/types/question
type Question struct {
	ID                 int       `json:"question_id"`
	AcceptedAnswerID   int       `json:"accepted_answer_id,omitempty"`
	AnswerCount        int       `json:"answer_count"`
	Answers            []Answer  `json:"answers,omitempty"`
	Body               string    `json:"body,omitempty"`
	BountyAmount       int       `json:"bounty_amount,omitempty"`
	BountyCloseDate    Time      `json:"bounty_closes_date,omitempty"`
	CloseVotes         int       `json:"close_vote_count,omitempty"`
	ClosedDate         Time      `json:"closed_date,omitempty"`
	ClosedReason       string    `json:"closed_reason,omitempty"`
	Comments           []Comment `json:"comments,omitempty"`
	CommunityOwnedDate Time      `json:"community_owned_date,omitempty"`
	Created            Time      `json:"creation_date,omitempty"`
	DeleteVotes        int       `json:"delete_vote_count,omitempty"`
	DownVotes          int       `json:"down_vote_count,omitempty"`
	Favorites          int       `json:"favorite_count,omitempty"`
	IsAnswered         bool      `json:"is_answered"`
	LastActivity       Time      `json:"last_activity_date,omitempty"`
	LastEdit           Time      `json:"last_edit_date,omitempty"`
	Link               string    `json:"link,omitempty"`
	LockDate           Time      `json:"locked_date,omitempty"`
	Owner              *User     `json:"owner,omitempty"`
	ProtectedDate      Time      `json:"protected_date,omitempty"`
	ReopenVotes        int       `json:"reopen_vote_count,omitempty"`
	Score              int       `json:"score"`
	Tags               []string  `json:"tags,omitempty"`
	Title              string    `json:"title"`
	UpVotes            int       `json:"up_vote_count,omitempty"`
}

// Answer represents an answer to a question on one of the Stack Exchange sites.
//
// https://api.stackexchange.com/docs/types/answer
type Answer struct {
	ID                 int       `json:"answer_id"`
	Body               string    `json:"body,omitempty"`
	Comments           []Comment `json:"comments,omitempty"`
	CommunityOwnedDate Time      `json:"community_owned_date"`
	Created            Time      `json:"creation_date,omitempty"`
	DownVotes          int       `json:"down_vote_count,omitempty"`
	IsAccepted         bool      `json:"is_accepted"`
	LastActivity       Time      `json:"last_activity_date,omitempty"`
	LastEdit           Time      `json:"last_edit_date,omitempty"`
	Link               string    `json:"link,omitempty"`
	LockDate           Time      `json:"locked_date,omitempty"`
	Owner              *User     `json:"owner,omitempty"`
	QuestionID         int       `json:"question_id"`
	Score              int       `json:"score"`
	Tags               []string  `json:"tags,omitempty"`
	Title              string    `json:"title,omitempty"`
	UpVotes            int       `json:"up_vote_count,omitempty"`
	Views              int       `json:"view_count,omitempty"`
}

// Comment represents a remark on a question or answer.
//
// https://api.stackexchange.com/docs/types/comment
type Comment struct {
	ID      int    `json:"comment_id"`
	Body    string `json:"body,omitempty"`
	Created Time   `json:"creation_date"`
	Edited  bool   `json:"edited"`
	Link    string `json:"link,omitempty"`
	Owner   *User  `json:"owner,omitempty"`
	ReplyTo *User  `json:"reply_to_user,omitempty"`
	Score   int    `json:"score"`

	PostID   int    `json:"post_id"`
	PostType string `json:"post_type,omitempty"`
}

// User represents a Stack Exchange user.
//
// https://api.stackexchange.com/docs/types/user
type User struct {
	ID               int        `json:"user_id"`
	AboutMe          string     `json:"about_me,omitempty"`
	AcceptRate       int        `json:"accept_rate"`
	AccountID        int        `json:"account_id,omitempty"`
	Age              int        `json:"age,omitempty"`
	Answers          int        `json:"answer_count,omitempty"`
	BadgeCounts      BadgeCount `json:"badge_counts,omitempty"`
	Created          Time       `json:"creation_date,omitempty"`
	DisplayName      string     `json:"display_name"`
	DownVotes        int        `json:"down_vote_count,omitempty"`
	IsEmployee       bool       `json:"is_employee"`
	LastAccess       Time       `json:"last_access_date,omitempty"`
	LastModified     Time       `json:"last_modified_date,omitempty"`
	Link             string     `json:"link,omitempty"`
	Location         string     `json:"location,omitempty"`
	ProfileImage     string     `json:"profile_image,omitempty"`
	Questions        int        `json:"question_count,omitempty"`
	TimedPenaltyDate Time       `json:"timed_penalty_date,omitempty"`
	UpVotes          int        `json:"up_vote_count,omitempty"`
	Type             string     `json:"user_type"`
	Views            int        `json:"view_count,omitempty"`
	WebsiteURL       int        `json:"website_url,omitempty"`

	Reputation              int `json:"reputation"`
	ReputationChangeDay     int `json:"reputation_change_day,omitempty"`
	ReputationChangeWeek    int `json:"reputation_change_week,omitempty"`
	ReputationChangeMonth   int `json:"reputation_change_month,omitempty"`
	ReputationChangeQuarter int `json:"reputation_change_quarter,omitempty"`
	ReputationChangeYear    int `json:"reputation_change_year,omitempty"`
}

// BadgeCount is the total number of badges a user has earned.
//
// https://api.stackexchange.com/docs/types/badge-count
type BadgeCount struct {
	Bronze int `json:"bronze"`
	Silver int `json:"silver"`
	Gold   int `json:"gold"`
}

// Total returns the total badges.
func (bc BadgeCount) Total() int {
	return bc.Bronze + bc.Silver + bc.Gold
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
