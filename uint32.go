// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flux

import (
	"sync/atomic"
)

// Uint32 value in flux.
type Uint32 struct {
	value uint32 // Atomic.
	Waker
}

// MakeUint32 is for initializing a field.  Multiple copies of the struct value
// must not be used.
func MakeUint32(initial uint32) Uint32 {
	return Uint32{
		value: initial,
		Waker: MakeWaker(),
	}
}

// NewUint32 value in flux.
func NewUint32(initial uint32) *Uint32 {
	x := MakeUint32(initial)
	return &x
}

// Value takes a snapshot.  The value may have wrapped around (multiple times)
// since previous call.
func (x *Uint32) Value() uint32 {
	return atomic.LoadUint32(&x.value)
}

// Increment the value and possibly notify the observer.  The value may wrap
// around.
func (x *Uint32) Increment(i uint32) {
	atomic.AddUint32(&x.value, i)
	x.Poke()
}
