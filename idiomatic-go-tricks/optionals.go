package main

import (
	"encoding/json"
	"errors"
	"io"
)

// START OMIT

type Valid interface {
	OK() error
}

func (p Person) OK() error {
	if p.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func Decode(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return err
	}
	obj, ok := v.(Valid)
	if !ok {
		return nil // no OK method
	}
	err = obj.OK()
	if err != nil {
		return err
	}
	return nil
}

// END OMIT

type Person struct {
	Name string
}
