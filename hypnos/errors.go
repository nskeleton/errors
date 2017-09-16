package hypnos

import (
	"github.com/nskeleton/errors"
)

// IHypnosError ...
type IHypnosError interface { // Stutter is for demo purposes, you should proably call this NError or something
	errors.NError
	isHypnosError() bool
}

// HypnosError ...
type HypnosError struct {
	errors.nError
}

func (e *HypnosError) isHypnosError() bool {
	return true
}

// IsHypnosError ...
func IsHypnosError(err error) bool {
	e, ok := err.(HypnosError)
	return ok && e.isHypnosError()
}
