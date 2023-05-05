package main

import "math"

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}


func max[T Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func min[T Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func fequal(a, b float64) bool {
	return math.Abs(a - b) < 1e-9
}
