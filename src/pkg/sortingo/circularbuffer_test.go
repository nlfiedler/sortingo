//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

// numbers is used in testing the circular buffer
var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func TestNewCircularBuffer(t *testing.T) {
	cb := NewCircularBuffer(10)
	if cb.count != 0 {
		t.Error("new buffer has non-zero length")
	}
	if cb.Size() != 0 {
		t.Error("Size() should be zero for new buffer")
	}
	if cb.start != 0 {
		t.Error("new buffer has non-zero start")
	}
	if cb.end != 0 {
		t.Error("new buffer has non-zero end")
	}
	if len(cb.buffer) != 10 {
		t.Error("buffer capacity mismatch")
	}
	if cb.Capacity() != 10 {
		t.Error("Capacity() does not match initial value")
	}
	if cb.Remaining() != 10 {
		t.Error("Remaining() does not match capacity for empty buffer")
	}
	if !cb.Empty() {
		t.Error("Empty() should return true for new buffer")
	}
	if cb.Full() {
		t.Error("Full() should return false for new buffer")
	}
	if cb.Peek() != nil {
		t.Error("Peek() should return nil for new buffer")
	}
}

func TestNewCircularBufferFromSlice(t *testing.T) {
	input := make([]interface{}, 0)
	input = append(input, "foo")
	input = append(input, "bar")
	input = append(input, "baz")
	input = append(input, "quux")
	cb := NewCircularBufferFromSlice(input, false)
	if cb.count != 4 {
		t.Error("new buffer from slice length mismatch")
	}
	if cb.Size() != 4 {
		t.Error("Size() should match initial slice")
	}
	if cb.start != 0 {
		t.Error("new buffer has non-zero start")
	}
	if cb.end != 0 {
		t.Error("new buffer has non-zero end")
	}
	if len(cb.buffer) != 4 {
		t.Error("buffer capacity mismatch")
	}
	if cb.Capacity() != 4 {
		t.Error("Capacity() does not match initial value")
	}
	if cb.Remaining() != 0 {
		t.Error("Remaining() does not match available space")
	}
	if cb.Empty() {
		t.Error("Empty() should return false for non-empty buffer")
	}
	if !cb.Full() {
		t.Error("Full() should return true for full buffer")
	}
	if cb.Peek() == nil {
		t.Error("Peek() should return non-nil for non-empty buffer")
	}
}

func TestCircularBufferAdd(t *testing.T) {
	cb := NewCircularBuffer(10)
	if cb.Size() != 0 {
		t.Error("Size() should be zero for new buffer")
	}
	cb.Add("foo")
	if cb.Size() != 1 {
		t.Error("Size() should be one after calling Add()")
	}
	if cb.Capacity() != 10 {
		t.Error("Capacity() does not match initial value")
	}
	if cb.Remaining() != 9 {
		t.Error("Remaining() does not match available space")
	}
	if cb.Empty() {
		t.Error("Empty() should return false after Add()")
	}
	if cb.Full() {
		t.Error("Full() should return false when Remaining() > 0")
	}
	if cb.Peek() == nil {
		t.Error("Peek() should return non-nil for non-empty buffer")
	}
	elem := cb.Peek()
	if elem.(string) != "foo" {
		t.Error("Peek() returned incorrect element")
	}
	for cb.Add("foo") {
		// add until full
	}
	if cb.Size() != 10 {
		t.Error("Size() should be 10 after calling Add() ten times")
	}
	if cb.Capacity() != 10 {
		t.Error("Capacity() does not match initial value")
	}
	if cb.Remaining() != 0 {
		t.Error("Remaining() does not match available space")
	}
	if cb.Empty() {
		t.Error("Empty() should return false after Add()")
	}
	if !cb.Full() {
		t.Error("Full() should return true when Remaining() == 0")
	}
}

func TestCircularBufferDrain(t *testing.T) {
	cb := NewCircularBuffer(10)
	sink := make([]interface{}, 0, 10)
	count := cb.Drain(sink)
	if count != 0 {
		t.Error("Drain() with empty buffer should return zero")
	}
	// fill the buffer with unique values
	for i := 0; cb.Remaining() > 0; i++ {
		cb.Add(numbers[i])
	}
	sink = make([]interface{}, 0, 5)
	count = cb.Drain(sink)
	if count != 0 {
		t.Error("Drain() with small sink should return zero")
	}
	sink = make([]interface{}, 0, 10)
	count = cb.Drain(sink)
	if count != 10 {
		t.Error("Drain() should return full contents")
	}
	if cb.Remaining() != 10 {
		t.Error("Remaining() does not match available space")
	}
	if cb.Size() != 0 {
		t.Error("Size() should be zero after Drain()")
	}
	if !cb.Empty() {
		t.Error("Empty() should return true after Drain()")
	}
	if cb.Full() {
		t.Error("Full() should return false when Remaining() > 0")
	}
	if cb.Peek() != nil {
		t.Error("Peek() should return nil for empty buffer")
	}
	// examine the results to ensure they are correct
	for i, v := range sink {
		if v.(string) != numbers[i] {
			t.Error("results of Drain() do not match input")
		}
	}
}

func TestCircularBufferDrainWrap(t *testing.T) {
	cb := NewCircularBuffer(10)
	// add some elements and then remove them so that the next batch
	// we add will end up wrapping around within the buffer
	for i := 0; i < 5; i++ {
		cb.Add(numbers[i])
	}
	for cb.Size() > 0 {
		cb.Remove()
	}
	// now fill the buffer with unique values
	for i := 0; cb.Remaining() > 0; i++ {
		cb.Add(numbers[i])
	}
	sink := make([]interface{}, 0, 10)
	count := cb.Drain(sink)
	if count != 10 {
		t.Error("Drain() should return full contents")
	}
	if cb.Remaining() != 10 {
		t.Error("Remaining() does not match available space")
	}
	if cb.Size() != 0 {
		t.Error("Size() should be zero after Drain()")
	}
	if !cb.Empty() {
		t.Error("Empty() should return true after Drain()")
	}
	if cb.Full() {
		t.Error("Full() should return false when Remaining() > 0")
	}
	if cb.Peek() != nil {
		t.Error("Peek() should return nil for empty buffer")
	}
	// examine the results to ensure they are correct
	for i, v := range sink {
		if v.(string) != numbers[i] {
			t.Error("results of Drain() do not match input")
		}
	}
}

func TestCircularBufferRemove(t *testing.T) {
	cb := NewCircularBuffer(10)
	// add some elements and then remove them so that the next batch
	// we add will end up wrapping around within the buffer
	for i := 0; i < 5; i++ {
		cb.Add(numbers[i])
	}
	for i := 0; cb.Size() > 0; i++ {
		e := cb.Remove()
		if e.(string) != numbers[i] {
			t.Error("value from Remove() does not match input")
		}
	}
	// now fill the buffer with unique values
	for i := 0; cb.Remaining() > 0; i++ {
		cb.Add(numbers[i])
	}
	if cb.Size() != 10 {
		t.Error("adding values to test Remove() failed")
	}
	for i := 0; cb.Size() > 0; i++ {
		e := cb.Remove()
		if e.(string) != numbers[i] {
			t.Error("value from Remove() does not match input")
		}
	}
	if cb.Remaining() != 10 {
		t.Error("Remaining() does not match available space")
	}
	if !cb.Empty() {
		t.Error("Empty() should return true after Drain()")
	}
	if cb.Full() {
		t.Error("Full() should return false when Remaining() > 0")
	}
	if cb.Peek() != nil {
		t.Error("Peek() should return nil for empty buffer")
	}
}

func TestCircularBufferMoveSmallSource(t *testing.T) {
	source := NewCircularBuffer(10)
	for i := 0; i < 5; i++ {
		source.Add(numbers[i])
	}
	sink := NewCircularBuffer(10)
	count := source.Move(sink, 10)
	if count != 5 {
		t.Error("Move() returned unexpected count")
	}
	for i := 0; sink.Size() > 0; i++ {
		e := sink.Remove()
		if e.(string) != numbers[i] {
			t.Error("value from Remove() does not match input")
		}
	}
	if !sink.Empty() {
		t.Error("sink should be empty now")
	}
}

func TestCircularBufferMoveSmallSink(t *testing.T) {
	source := NewCircularBuffer(10)
	for i := 0; i < 10; i++ {
		source.Add(numbers[i])
	}
	sink := NewCircularBuffer(5)
	count := source.Move(sink, 10)
	if count != 5 {
		t.Error("Move() returned unexpected count")
	}
	for i := 0; sink.Size() > 0; i++ {
		e := sink.Remove()
		if e.(string) != numbers[i] {
			t.Error("value from Remove() does not match input")
		}
	}
	if !sink.Empty() {
		t.Error("sink should be empty now")
	}
}

func TestCircularBufferMove(t *testing.T) {
	source := NewCircularBuffer(10)
	for i := 0; i < 10; i++ {
		source.Add(numbers[i])
	}
	sink := NewCircularBuffer(10)
	count := source.Move(sink, 10)
	if count != 10 {
		t.Error("Move() returned unexpected count")
	}
	for i := 0; sink.Size() > 0; i++ {
		e := sink.Remove()
		if e.(string) != numbers[i] {
			t.Error("value from Remove() does not match input")
		}
	}
	if !sink.Empty() {
		t.Error("sink should be empty now")
	}
}

func TestCircularBufferMoveSourceWrap(t *testing.T) {
	source := NewCircularBuffer(10)
	// fill the source buffer, but force it to wrap around
	for i := 0; i < 5; i++ {
		source.Add(numbers[i])
	}
	for !source.Empty() {
		source.Remove()
	}
	for i := 0; i < 10; i++ {
		source.Add(numbers[i])
	}
	sink := NewCircularBuffer(10)
	count := source.Move(sink, 10)
	if count != 10 {
		t.Error("Move() returned unexpected count")
	}
	for i := 0; sink.Size() > 0; i++ {
		e := sink.Remove()
		if e.(string) != numbers[i] {
			t.Error("value from Remove() does not match input")
		}
	}
	if !sink.Empty() {
		t.Error("sink should be empty now")
	}
}

func TestCircularBufferMoveSinkWrap(t *testing.T) {
	source := NewCircularBuffer(10)
	for i := 0; i < 10; i++ {
		source.Add(numbers[i])
	}
	sink := NewCircularBuffer(10)
	// cause the sink to wrap around after Move() is called
	for i := 0; i < 5; i++ {
		sink.Add(numbers[i])
	}
	for !sink.Empty() {
		sink.Remove()
	}
	count := source.Move(sink, 10)
	if count != 10 {
		t.Error("Move() returned unexpected count")
	}
	for i := 0; sink.Size() > 0; i++ {
		e := sink.Remove()
		if e.(string) != numbers[i] {
			t.Error("value from Remove() does not match input")
		}
	}
	if !sink.Empty() {
		t.Error("sink should be empty now")
	}
}
