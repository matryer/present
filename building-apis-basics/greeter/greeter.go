package greeter

import (
	"errors"
	"fmt"
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
