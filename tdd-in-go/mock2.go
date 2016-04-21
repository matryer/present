package main

import "github.com/stretchr/testify/mock"

// START OMIT

type mockedThing struct {
	mock.Mock
}

// DoSomething is an ugly way of writing a mock
func (m *mockedThing) DoSomething(n int) (bool, error) {
	returns := m.Called(n)
	if err, hasErr := returns.Get(1).(error); hasErr {
		return false, hasErr
	}
	return returns.Get(0).(bool), nil
}

// END OMIT
