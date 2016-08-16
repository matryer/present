package main

import "errors"

// START OMIT

func BrilliantFunction() (*Thing, error) {
	something, err := GetSomething()
	if err != nil {
		return nil, err
	}
	defer something.Close()
	if !something.OK() {
		return nil, errors.New("something went wrong")
	}
	another, err := something.Else()
	if err != nil {
		return nil, &customErr{err: err, location: "BrilliantFunction"}
	}
	another.Lock()
	defer another.Unlock()
	err = another.Update(1)
	if err != nil {
		return nil, err
	}
	return another.Thing(), nil
}

// END OMIT

type Thing struct{}
type something struct{}

func (s *something) Close()   {}
func (s *something) OK() bool { return false }
func (s *something) Else() (*something, error) {
	return s, nil
}
func (s *something) Lock()              {}
func (s *something) Unlock()            {}
func (s *something) Thing() *Thing      { return nil }
func (s *something) Update(i int) error { return nil }

func GetSomething() (*something, error) {
	return nil, nil
}

type customErr struct {
	err      error
	location string
}

func (c customErr) Error() string {
	return "err"
}
