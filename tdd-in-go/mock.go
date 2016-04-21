package main

// START OMIT

// make a mocked thing
thing := new(mockedThing)
thing.On("DoSomething", 123).Return(true, nil)

// call to code being tested
LetsDoSomething(thing)

// make sure things happened as expected
mock.AssertExpectations(t)

// END OMIT