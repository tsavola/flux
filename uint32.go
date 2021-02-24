// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package threshold

import (
	"sync/atomic"
)

// Uint32 threshold.
type Uint32 struct {
	Waker
	value uint32
}

// MakeUint32 is for initializing a field.  Multiple copies of the struct
// value must not be used.
func MakeUint32(initial uint32) Uint32 {
	return Uint32{
		Waker: MakeWaker(),
		value: initial,
	}
}

// NewUint32 threshold value.
func NewUint32(value uint32) *Uint32 {
	t := MakeUint32(value)
	return &t
}

// Value is safe to call concurrently.  The value may have wrapped around since
// previous call.
func (t *Uint32) Value() uint32 {
	return atomic.LoadUint32(&t.value)
}

// ValueNonatomic is a racy version of Value.
func (t *Uint32) ValueNonatomic() uint32 {
	return t.value
}

// Increment the value and possibly notify the observer.  The value may wrap
// around.
func (t *Uint32) Increment(i uint32) {
	atomic.AddUint32(&t.value, i)
	t.Poke()
}
