// Copyright 2009-2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sortingo

// A simple circular buffer of fixed length which has an empty and full
// state. When full, the buffer will not accept any new entries.

// CircularBuffer is an array of elements that is fixed in size and thus
// will stop receiving new elements once full. Removing elements makes
// room for new ones. Buffers may share a single underlying slice if
// given lower and upper bounds that do not overlap.
type CircularBuffer struct {
	buffer []interface{} // contains the buffer elements
	start  int           // position of first element
	end    int           // position of last element
	count  int           // number of elements
}

// NewCircularBuffer constructs a CircularBuffer with the given capacity.
func NewCircularBuffer(capacity int) *CircularBuffer {
	cb := new(CircularBuffer)
	cb.buffer = make([]interface{}, capacity)
	return cb
}

// NewCircularBufferFromSlice constructs a CircularBuffer from the given
// slice. All entries in the slice are assumed to be valid data such
// that the buffer count will be equal to the length of the given slice.
// If dupe is true, will create a new slice and copy the contents of
// initial to that slice.
func NewCircularBufferFromSlice(initial []interface{}, dupe bool) *CircularBuffer {
	cb := new(CircularBuffer)
	if dupe {
		cb.buffer = make([]interface{}, 0, len(initial))
		copy(cb.buffer, initial)
	} else {
		cb.buffer = initial
	}
	cb.count = len(initial)
	return cb
}

// Add adds the given value to the buffer, returning true if
// successful, or false if the buffer is full.
func (cb *CircularBuffer) Add(e interface{}) bool {
	if cb.count == cap(cb.buffer) {
		return false
	}
	cb.count++
	cb.buffer[cb.end] = e
	cb.end++
	if cb.end == len(cb.buffer) {
		cb.end = 0
	}
	return true
}

// Capacity returns the total number of elements this buffer can hold.
func (cb *CircularBuffer) Capacity() int {
	return cap(cb.buffer)
}

// Drain moves the contents of the circular buffer into the given output
// slice in an efficient manner, leaving this buffer empty. Returns the
// number of elements copied to the slice.
func (cb *CircularBuffer) Drain(sink []interface{}) int {
	if cb.count == 0 {
		return 0 // nothing to copy
	}
	if cap(sink)-len(sink) < cb.count {
		return 0 // insufficient space
	}
	if cb.end <= cb.start {
		// elements wrap around, must make two calls to copy()
		copy(sink, cb.buffer[cb.start:])
		copy(sink, cb.buffer[:cb.end+1])
	} else {
		// elements are in one contiguous region
		copy(sink, cb.buffer[cb.start:cb.end+1])
	}
	cb.start = 0
	cb.end = 0
	count := cb.count
	cb.count = 0
	return count
}

// Empty returns true if the buffer is empty, false otherwise.
func (cb *CircularBuffer) Empty() bool {
	return cb.count == 0
}

// Full returns true if the buffer is full, false otherwise.
func (cb *CircularBuffer) Full() bool {
	return cb.count == cap(cb.buffer)
}

// Move removes the given number of elements from the circular buffer
// and adds them to the sink buffer in an efficient manner. This is
// equivalent to repeatedly removing elements from this buffer and
// adding them to the sink. Returns the number of elements moved, which
// may be less than requested if the origin has fewer items, or if the
// sink has insufficient space.
func (cb *CircularBuffer) Move(sink *CircularBuffer, count int) int {
	if cb.count < count {
		count = cb.count
	}
	if sink.Remaining() < count {
		count = sink.Remaining()
	}
	tocopy := count
	capacity := cap(cb.buffer)
	sapacity := cap(sink.buffer)
	for tocopy > 0 {
		// compute how much can be copied from source
		var available int
		if cb.start < cb.end {
			available = cb.end - cb.start
		} else {
			// wraps around, start with upper portion
			available = capacity - cb.start
		}
		// compute how much can be copied to sink
		var willfit int
		if sink.start <= sink.end {
			willfit = sapacity - sink.end
		} else {
			// wraps around
			willfit = sink.start - sink.end
		}
		willcopy := iMin(available, willfit)
		if willcopy <= 0 {
			break
		}
		copy(sink.buffer[sink.end:], cb.buffer[cb.start:cb.start+willcopy])
		sink.end += willcopy
		if sink.end >= sapacity {
			sink.end = 0
		}
		cb.start += willcopy
		if cb.start >= capacity {
			cb.start = 0
		}
		tocopy -= willcopy
	}
	if tocopy > 0 {
		panic("failed to move circular buffer contents")
	}
	sink.count += count
	cb.count -= count
	return count
}

// Peek returns the first element in the buffer, without removing it.
// If the buffer is empty, nil is returned.
func (cb *CircularBuffer) Peek() interface{} {
	if cb.count == 0 {
		return nil
	}
	return cb.buffer[cb.start]
}

// Remove removes the first element in the buffer, reducing the
// number of elements in the buffer by one. If the buffer is
// empty, nil is returned.
func (cb *CircularBuffer) Remove() interface{} {
	if cb.count == 0 {
		return nil
	}
	cb.count--
	e := cb.buffer[cb.start]
	cb.start++
	if cb.start == cap(cb.buffer) {
		cb.start = 0
	}
	return e
}

// Remaining returns the number of empty spaces within this buffer.
func (cb *CircularBuffer) Remaining() int {
	return cap(cb.buffer) - cb.count
}

// Size returns the number of elements in the circular buffer.
func (cb *CircularBuffer) Size() int {
	return cb.count
}
