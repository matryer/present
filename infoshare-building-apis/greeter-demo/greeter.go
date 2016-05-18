package main

import (
	"errors"
	"fmt"
	"log"
)

type Greeter struct {
	Format string
}

func NewGreeter(format string) (*Greeter, error) {
	if len(format) == 0 {
		return nil, errors.New("format required")
	}
	g := &Greeter{
		Format: format,
	}
	return g, nil
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf(g.Format, name)
}

// START OMIT

func main() {

	greeter, err := NewGreeter("Hello %s")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeter.Greet("Mat"))
	fmt.Println(greeter.Greet("David"))
	fmt.Println(greeter.Greet("Laurie"))

}

// END OMIT
