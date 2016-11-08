package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheekybits/is"
)

var tests = []struct {
	Method       string
	Path         string
	Body         io.Reader
	BodyContains string
	Status       int
}{{
	Method:       "GET",
	Path:         "/things",
	BodyContains: "Hello Gophers",
	Status:       http.StatusOK,
}, {
	Method:       "POST",
	Path:         "/things",
	Body:         strings.NewReader(`{"name":"Gophers"}`),
	BodyContains: "Hello Gophers",
	Status:       http.StatusCreated,
}}

// START OMIT
func TestAll(t *testing.T) {
	server := httptest.NewServer(&myhandler{}) // HL
	defer server.Close()                       // HL
	for _, test := range tests {
		t.Run(test.Method + " " + test.Path, func(t *testing.T) {
			is := is.New(t)
			r, err := http.NewRequest(test.Method, server.URL+test.Path, test.Body) // HL
			is.NoErr(err)
			// call handler
			response, err := http.DefaultClient.Do(r) // HL
			is.NoErr(err)
			actualBody, err := ioutil.ReadAll(response.Body)
			is.NoErr(err)
			// make assertions
			is.True(strings.Contains(string(actualBody), test.BodyContains))        // HL
			is.Equal(test.Status, response.StatusCode) // HL
		})
	}
}

// END OMIT
type myhandler struct{}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestAll", F: TestAll})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
