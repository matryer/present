package main

// START OMIT

import (
	"testing"
)

func TestMeaningOfLife(t *testing.T) {
	result := add(36, 6)
	if result != 42 {
		t.Errorf("%d is not the meaning of life", result)
	}
}

// END OMIT

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestMeaningOfLife", F: TestMeaningOfLife})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
