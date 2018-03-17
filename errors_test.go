package onedrive

import (
	"testing"
)

func TestIsError(t *testing.T) {
	testCases := map[string]struct {
		err  *Error
		code string
		exp  bool
	}{
		"nil":       {nil, "", false},
		"Not found": {&Error{Code: "abc"}, "cba", false},
		"Top level": {&Error{Code: "abc"}, "abc", true},
		"2nd level": {&Error{Code: "abc", InnerError: &Error{Code: "bcd"}}, "bcd", true},
		"3rd level": {&Error{Code: "abc", InnerError: &Error{Code: "bcd", InnerError: &Error{Code: "cde"}}}, "cde", true},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got, exp := tc.err.IsError(tc.code), tc.exp; got != exp {
				t.Errorf("expected %v; got %v", exp, got)
			}
		})
	}
}
