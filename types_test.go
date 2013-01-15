package stackexchange

import (
	"testing"
	"time"
)

func TestUnmarshalTime(t *testing.T) {
	tests := []struct {
		In  string
		Out time.Time
	}{
		{"0", time.Unix(0, 0)},
		{"-42", time.Unix(-42, 0)},
		{"42", time.Unix(42, 0)},
	}
	for _, test := range tests {
		mytime := new(Time)
		if in, out, err := test.In, test.Out, mytime.UnmarshalJSON([]byte(test.In)); err != nil {
			t.Errorf("new(Time).UnmarshalJSON(%q) error: %v", in, err)
		} else if mytime := time.Time(*mytime); mytime != out {
			t.Errorf("new(Time).UnmarshalJSON(%q) = %v; want %v", in, mytime, out)
		}
	}
}
