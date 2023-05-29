package bacotell_util

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

func Ptr[T any](v T) *T {
	return &v
}

func PtrConv[T number, R number](v *T) *R {
	if v == nil {
		return nil
	}

	return Ptr(R(*v))
}
