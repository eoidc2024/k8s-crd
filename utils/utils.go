package utils

import (
	"golang.org/x/exp/constraints"
)

func NewPointerOfType[T constraints.Ordered](value T) *T {
	return &value
}

// NewBoolPointer returns pointer on a new bool value instance
func NewBoolPointer(b bool) *bool {
	return &b
}

// DefaultBoolIfUnset sets default value d of a boolean if unset
func DefaultBoolIfUnset(valPtr **bool, d bool) {
	if *valPtr == nil {
		*valPtr = &d
	}
}

func DefaultIfUnset[T constraints.Ordered](valPtr **T, d T) {
	if *valPtr == nil {
		*valPtr = &d
	}
}
