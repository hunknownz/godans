package ndslice

import "sync"

type Ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type sliceInt[T Ints] []T

type NDSliceInt[T Ints] struct {
	data    sliceInt[T]
	nanOnce sync.Once
	nans    *BitSet
}
