package tokenizer

import (
	"fmt"
	"github.com/tomvodi/bww-format/internal/common/token"
)

type Tokenizer struct {
}

func (t *Tokenizer) Tokenize(data []byte) ([]token.Token, error) {
	return nil, fmt.Errorf("not implemented")
}

func New() *Tokenizer {
	return &Tokenizer{}
}
