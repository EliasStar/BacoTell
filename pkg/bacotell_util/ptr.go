package bacotell_util

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

type nilable[K comparable, V any] interface {
	~*any | ~[]V | ~map[K]V | ~chan V
}

func Ptr[T any](v T) *T {
	return &v
}

func Val[T any](v *T) T {
	return *v
}

func NilablePtr[K comparable, V any, T nilable[K, V]](v T) *T {
	if v == nil {
		return nil
	}

	return &v
}

func NilableVal[K comparable, V any, T nilable[K, V]](v *T) T {
	if v == nil {
		return nil
	}

	return *v
}

func PtrConv[T number, R number](v *T) *R {
	if v == nil {
		return nil
	}

	return Ptr(R(*v))
}
