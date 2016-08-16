package main

import "errors"

// START OMIT

func UnbrilliantFunction() (*Thing, error) {
	something, err := GetSomething()
	if err != nil {
		return nil, err
	}
	defer something.Close()
	if something.OK() {
		another, err := something.Else()
		if err != nil {
			return nil, &customErr{err: err, location: "BrilliantFunction"}
		}
		another.Lock()
		err = another.Update(1)
		if err == nil {
			another.Unlock()
			return another.Thing(), nil
		}
		another.Unlock()
		return nil, err
	} else {
		return nil, errors.New("something went wrong")
	}
}

// END OMIT
