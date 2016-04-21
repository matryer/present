package main

import "testing"

// START OMIT
func sum(nums ...int) int {
	var r int
	for _, n := range nums {
		r += n
	}
	return r
}
func add(a, b int) int {
	return sum(a, b)
}
func minus(a, b int) int {
	return sum(a, 0-b)
}

// END OMIT
func TestAll(t *testing.T) {
	if sum(1, 2, 3) != 6 {
		t.Error("sum failed")
	}
	if add(1, 4) != 5 {
		t.Error("add failed")
	}
	if minus(10, 2) != 8 {
		t.Error("minus failed")
	}
}

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestAll", F: TestAll})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
