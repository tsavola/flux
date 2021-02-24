// Copyright (c) 2021 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

Package threshold provides concurrently accessible scalar value types which can
be updated incrementally.  A single observer goroutine can receive change
notifications via a channel.  There can be one or more mutator goroutines.

Mutator goroutines can call the Increment, Poke, Finish and Value methods.  If
there is only one mutator, it may also call the ValueNonatomic method.  The
consumer goroutine can call the Chan and Value methods.

There isn't a one-to-one correspondence between updates and notifications:
multiple Increment calls may be coalesced to a single successful channel
reception.  It's also possible that the same value is observed after receiving
a notification.  The value may wrap around; it's the application's
responsibility to take that into account.

*/
package threshold
