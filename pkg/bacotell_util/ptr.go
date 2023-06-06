// Package bacotell_util provides utilities for plugins and BacoTell.
package bacotell_util

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

type nilable[K comparable, V any] interface {
	~*any | ~[]V | ~map[K]V | ~chan V
}

// Ptr returns a pointer to the argument.
func Ptr[T any](v T) *T {
	return &v
}

// Val returns the value pointed to by the argument.
func Val[T any](v *T) T {
	return *v
}

// NilablePtr returns a pointer to the argument if it is not nil else a nil pointer.
func NilablePtr[K comparable, V any, T nilable[K, V]](v T) *T {
	if v == nil {
		return nil
	}

	return &v
}

// NilableVal returns the value pointed to by the argument if it is not nil else nil.
func NilableVal[K comparable, V any, T nilable[K, V]](v *T) T {
	if v == nil {
		return nil
	}

	return *v
}

// PtrConv converts a pointer of type T to a pointer of type R.
// It returns a pointer to the converted value.
// If the input pointer is nil, it returns nil.
func PtrConv[T number, R number](v *T) *R {
	if v == nil {
		return nil
	}

	return Ptr(R(*v))
}
