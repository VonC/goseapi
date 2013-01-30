// Package stackexchange provides access to the Stack Exchange 2.0 API.
//
// http://api.stackexchange.com/
package stackexchange

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Root is the Stack Exchange API endpoint.
const Root = "https://api.stackexchange.com/" + Version

// Version is the API version identifier.
const Version = "2.1"

// Well-known Stack Exchange sites
const (
	StackOverflow = "stackoverflow"
)

// Sort orders
const (
	SortActivity     = "activity"
	SortCreationDate = "creation"
	SortHot          = "hot"
	SortWeek         = "week"
	SortMonth        = "month"
	SortScore        = "votes"
)

type Params struct {
	Site string

	Sort     string
	Order    string
	Page     int
	PageSize int

	Filter string
}

// DefaultClient uses the default HTTP client and API root.
var DefaultClient *Client = nil

// Do performs an API request using the default client.
func Do(path string, v interface{}, params Params) (*Wrapper, error) {
	return DefaultClient.Do(path, v, params)
}

// A Client can make API requests.
type Client struct {
	Client *http.Client
	Root   string

	// Pass these fields if you have an OAuth 2.0 application registered with stackapps.com.
	AccessToken string
	Key         string
}

// Do performs an API request.
func (c *Client) Do(path string, v interface{}, params Params) (*Wrapper, error) {
	// Get arguments
	client := http.DefaultClient
	if c != nil && c.Client != nil {
		client = c.Client
	}
	root := Root
	if c != nil && c.Root != "" {
		root = c.Root
	}

	// Build URL parameters
	vals := url.Values{
		"site": {params.Site},
	}
	if params.Sort != "" {
		vals.Set("sort", params.Sort)
	}
	if params.Order != "" {
		vals.Set("order", params.Order)
	}
	if params.Page != 0 {
		vals.Set("page", strconv.Itoa(params.Page))
	}
	if params.PageSize != 0 {
		vals.Set("pagesize", strconv.Itoa(params.PageSize))
	}
	if params.Filter != "" {
		vals.Set("filter", params.Filter)
	}
	if c.AccessToken != "" {
		vals.Set("access_token", c.AccessToken)
	}
	if c.Key != "" {
		vals.Set("key", c.Key)
	}

	// Send request
	resp, err := client.Get(root + path + "?" + vals.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response
	var result struct {
		Items items `json:"items"`

		ErrorID      int    `json:"error_id"`
		ErrorName    string `json:"error_name"`
		ErrorMessage string `json:"error_message"`

		Page     int  `json:"page"`
		PageSize int  `json:"page_size"`
		HasMore  bool `json:"has_more"`

		Backoff        int `json:"backoff"`
		QuotaMax       int `json:"quota_max"`
		QuotaRemaining int `json:"quota_remaining"`

		Total int    `json:"total"`
		Type  string `json:"type"`
	}
	result.Items = items{v}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &Wrapper{
		Error: Error{
			ID:      result.ErrorID,
			Name:    result.ErrorName,
			Message: result.ErrorMessage,
		},
		Page:           result.Page,
		PageSize:       result.PageSize,
		HasMore:        result.HasMore,
		Backoff:        result.Backoff,
		QuotaMax:       result.QuotaMax,
		QuotaRemaining: result.QuotaRemaining,
		Total:          result.Total,
		Type:           result.Type,
	}, err
}

type items struct {
	val interface{}
}

func (i items) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, i.val)
}
