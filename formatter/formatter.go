package formatter

import "fmt"

type Formatter struct {
}

func (f *Formatter) Format(data []byte) ([]byte, error) {
	// Split into tokens (line or inline)
	// Cleanup and fix if needed
	// Put tokens together
	return nil, fmt.Errorf("not implemented")
}

func New() *Formatter {
	return &Formatter{}
}
