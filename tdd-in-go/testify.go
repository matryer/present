package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	if val, err := MethodToTest(1, 2, 3); assert.NoError(t, err) {
		if assert.NotNil(t, val) {
			assert.Equal(t, val.Name, "Stretchr")
		}
	}

}
