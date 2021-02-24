// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package threshold

import (
	"sync/atomic"
)

// Uint64 threshold.
type Uint64 struct {
	Waker
	value uint64
}

// MakeUint64 is for initializing a field.  Multiple copies of the struct
// value must not be used.
func MakeUint64(initial uint64) Uint64 {
	return Uint64{
		Waker: MakeWaker(),
		value: initial,
	}
}

// NewUint64 threshold value.
func NewUint64(value uint64) *Uint64 {
	t := MakeUint64(value)
	return &t
}

// Value is safe to call concurrently.  The value may have wrapped around since
// previous call.
func (t *Uint64) Value() uint64 {
	return atomic.LoadUint64(&t.value)
}

// ValueNonatomic is a racy version of Value.
func (t *Uint64) ValueNonatomic() uint64 {
	return t.value
}

// Increment the value and possibly notify the observer.  The value may wrap
// around.
func (t *Uint64) Increment(i uint64) {
	atomic.AddUint64(&t.value, i)
	t.Poke()
}
