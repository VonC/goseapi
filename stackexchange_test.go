package stackexchange

import (
	"log"
	"testing"
)

func TestFillPlaceholders(t *testing.T) {
	tests := []struct {
		Path   string
		Args   []string
		String string
	}{
		{"", nil, ""},
		{"", []string{}, ""},
		{"", []string{"foo"}, ""},
		{"/questions", nil, "/questions"},
		{"/questions", []string{}, "/questions"},
		{"/questions", []string{"foo"}, "/questions"},
		{"/questions/{id}", nil, "/questions/{id}"},
		{"/questions/{id}", []string{}, "/questions/{id}"},
		{"/questions/{id}", []string{"42"}, "/questions/42"},
		{"/questions/{id}", []string{"42", "bacon"}, "/questions/42"},
		{"/questions/{id}/comments", nil, "/questions/{id}/comments"},
		{"/questions/{id}/comments", []string{}, "/questions/{id}/comments"},
		{"/questions/{id}/comments", []string{"42"}, "/questions/42/comments"},
		{"/questions/{id}/comments", []string{"42", "bacon"}, "/questions/42/comments"},
		{"/questions/{id/comments", []string{"42", "bacon"}, "/questions/{id/comments"},
		{"/tags/{tag}/top-askers/{period}", []string{"bacon", "42"}, "/tags/bacon/top-askers/42"},
	}
	for _, test := range tests {
		if out := fillPlaceholders(test.Path, test.Args); out != test.String {
			t.Errorf("fillPlaceholders(%q, %q) = %q; want %q", test.Path, test.Args, out, test.String)
		}
	}
}

func TestJoinIDs(t *testing.T) {
	tests := []struct {
		IDs    []int
		String string
	}{
		{nil, ""},
		{[]int{}, ""},
		{[]int{1}, "1"},
		{[]int{40}, "40"},
		{[]int{1, 40}, "1;40"},
		{[]int{1, 40, 72}, "1;40;72"},
		{[]int{1, 5, 2, 3}, "1;5;2;3"},
	}
	for _, test := range tests {
		if out := JoinIDs(test.IDs); out != test.String {
			t.Errorf("JoinIDs(%v) = %q; want %q", test.IDs, out, test.String)
		}
	}
}

func TestExampleDo(t *testing.T) {
	var questions []Question
	wrapper, err := Do(PathQuestions, &questions, &Params{
		Site: StackOverflow,
		Args: []string{"11227809"},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(wrapper)
	log.Println(questions[0].Title)
}
