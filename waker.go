// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flux

// Waker coalesces notifications.
type Waker struct {
	c chan struct{}
}

// MakeWaker for a single outstanding notification.
func MakeWaker() Waker {
	return Waker{make(chan struct{}, 1)}
}

// Chan is unblocked when it's time to wake up.  It is closed at Finish.
func (w Waker) Chan() <-chan struct{} {
	return w.c
}

// Poke the observer to wake it up, unless a wake-up is already pending.
func (w Waker) Poke() {
	select {
	case w.c <- struct{}{}:
	default:
	}
}

// Finish wakes up the observer for the last time.  Chan will be closed.
func (w Waker) Finish() {
	close(w.c)
}
