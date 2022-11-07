// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flux

import (
	"sync/atomic"
)

// Uint64 value in flux.
type Uint64 struct {
	value uint64 // Atomic.
	Waker
}

// MakeUint64 is for initializing a field.  Multiple copies of the struct value
// must not be used.
func MakeUint64(initial uint64) Uint64 {
	return Uint64{
		value: initial,
		Waker: MakeWaker(),
	}
}

// NewUint64 value in flux.
func NewUint64(initial uint64) *Uint64 {
	x := MakeUint64(initial)
	return &x
}

// Value takes a snapshot.  The value may have wrapped around (multiple times)
// since previous call.
func (x *Uint64) Value() uint64 {
	return atomic.LoadUint64(&x.value)
}

// Increment the value and possibly notify the observer.  The value may wrap
// around.
func (x *Uint64) Increment(i uint64) {
	atomic.AddUint64(&x.value, i)
	x.Poke()
}
