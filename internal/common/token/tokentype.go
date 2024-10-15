package token

//go:generate go run github.com/dmarkham/enumer -json -yaml -transform=lower -type=Type

type Type uint8

const (
	Line Type = iota
	Inline
)
