package main

// START OMIT

import (
	"strings"
	"testing"
)

var acronymTests = []struct {
	in  string
	out string
}{{
	in:  "hello gophercon, how are you?",
	out: "hghay",
}, {
	in:  "Where will Gophercon be hosted next year?",
	out: "wwgbhny",
}}

func TestAcronym(t *testing.T) {
	for _, test := range acronymTests {
		if Acronym(test.in) != test.out {
			t.Errorf("Expected %s but got %s.", test.out, Acronym(test.in))
		}
	}
}

// END OMIT

func Acronym(s string) string {
	var letters string
	for _, word := range strings.Split(strings.ToLower(s), " ") {
		letters = letters + string(word[0])
	}
	return letters
}

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestTitle", F: TestAcronym})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
