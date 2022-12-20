package main

type Floats interface {
	~float32 | ~float64 | ~int
}

type DFloat[T Floats] []T

func (d DFloat[T]) Count() int {
	return len(d)
}
