package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// START OMIT

func Setup(t *testing.T) (io.Reader, func(), error) {
	teardown := func() {}
	// make a test file
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		return nil, teardown, err
	}
	teardown = func() {
		// close f
		err := f.Close()
		if err != nil {
			t.Error("Setup: Close:", err)
		}
		// delete the test file
		err = os.RemoveAll(f.Name())
		if err != nil {
			t.Error("Setup: RemoveAll:", err)
		}
	}
	return f, teardown, nil
}

// END OMIT

// TWO OMIT

func TestSomething(t *testing.T) {
	r, teardown, err := Setup(t)
	defer teardown()
	if err != nil {
		t.Error("Setup:", err)
	}
	// do something with r
}

// TWO END OMIT
