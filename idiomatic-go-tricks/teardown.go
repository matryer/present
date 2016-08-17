package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// START OMIT

func setup(t *testing.T) (*os.File, func(), error) {
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
			t.Error("setup: Close:", err)
		}
		// delete the test file
		err = os.RemoveAll(f.Name())
		if err != nil {
			t.Error("setup: RemoveAll:", err)
		}
	}
	return f, teardown, nil
}

// END OMIT

// TWO OMIT

func TestSomething(t *testing.T) {
	f, teardown, err := setup(t)
	defer teardown()
	if err != nil {
		t.Error("setup:", err)
	}
	// do something with f
}

// TWO END OMIT
