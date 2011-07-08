//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// This file contains utility functions for the unit tests.

import (
	"bytes"
	"fmt"
	"rand"
	"sort"
	"strings"
	"testing"
)

// smallDataSize is the recommended test size for the slower sorting
// algorithms (e.g. insertion, gnome, selection).
const smallDataSize = 512
// mediumDataSize is the recommended test size for the sorting algorithms
// that are faster than the typical slow sorts, but not necessarily as
// fast as the memory efficient sorts.
const mediumDataSize = 16384
// largeDataSize is the recommended test size for the faster sorting
// algorithms (e.g. burstsort, funnelsort). This is also the largest
// allowable size for testing.
const largeDataSize = 65536

// repeatedStrings contains a sequence of repeated strings.
var repeatedStrings []string
// repeatedCycleStrings contains a repeating sequence of strings.
var repeatedCycleStrings []string
// randomStrings consists of strings of 100 mixed-case letters and numbers.
var randomStrings []string
// uniqueWords consists of unique pseudo words, similar to a dictionary.
var uniqueWords []string
// nonUniqueWords consists of pseudo words that may repeat numerous times.
var nonUniqueWords []string

// init sets up the test data one time to avoid regenerating repeatedly.
func init() {
	// Generate the repeated strings test data.
	repeatedStrings = make([]string, largeDataSize)
	a100 := strings.Repeat("A", 100)
	for idx := range repeatedStrings {
		repeatedStrings[idx] = a100
	}

	// Generate a repeating cycle of strings.
	strs := make([]string, len(a100))
	for i := range strs {
		strs[i] = a100[0 : i+1]
	}
	repeatedCycleStrings = make([]string, largeDataSize)
	c := 0
	for i := range repeatedCycleStrings {
		repeatedCycleStrings[i] = strs[c]
		if c++; c >= len(strs) {
			c = 0
		}
	}

	// Generate a set of random strings, each of length 100.
	randomStrings = make([]string, largeDataSize)
	for i := range randomStrings {
		bb := bytes.NewBuffer(make([]byte, 0, 100))
		for j := 0; j < 100; j++ {
			d := rand.Intn(62)
			if d < 10 {
				bb.WriteRune('0' + d)
			} else if d < 36 {
				bb.WriteRune('A' + (d - 10))
			} else {
				bb.WriteRune('a' + (d - 36))
			}
		}
		randomStrings[i] = string(bb.Bytes())
	}

	// Generate a set of unique pseudo words.
	uniqueWords = make([]string, largeDataSize)
	wordExists := make(map[string]bool)
	for i := range uniqueWords {
		var s string
		// Loop until a unique random word is generated.
		for {
			// Each word is from 1 to 28 characters long.
			l := 1 + rand.Intn(27)
			bb := bytes.NewBuffer(make([]byte, 0, l))
			// Each word consists only of the lowercase letters.
			for j := 0; j < l; j++ {
				d := rand.Intn(26)
				bb.WriteRune('a' + d)
			}
			s = string(bb.Bytes())
			if !wordExists[s] {
				break
			}
		}
		uniqueWords[i] = s
		wordExists[s] = true
	}

	// Generate a set of pseudo words that may be repeated.
	nonUniqueWords = make([]string, largeDataSize)
	n := len(nonUniqueWords)
	for i := 0; i < n; {
		// Each word is 1 to 28 characters long.
		l := 1 + rand.Intn(27)
		bb := bytes.NewBuffer(make([]byte, 0, l))
		// Each word consists only of the lowercase letters.
		for j := 0; j < l; j++ {
			d := rand.Intn(26)
			bb.WriteRune('a' + d)
		}
		// Repeat the word some number of times.
		c := rand.Intn(100)
		if c > (n - i) {
			c = n - i
		}
		s := string(bb.Bytes())
		for j := 0; j < c; j++ {
			nonUniqueWords[i] = s
			i++
		}
	}
	shuffle(nonUniqueWords)
}

// testSortArguments runs a given sort function with the most
// basic of input sequences in order to test its robustness.
func testSortArguments(t *testing.T, f func([]string)) {
	// these should silently do nothing
	f(nil)
	f(make([]string, 0))
	f([]string{"a"})
	// test the bare minimum input, two elements
	input := []string{"b", "a"}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("two inputs not sorted")
	}
	// now try three elements
	input = []string{"c", "b", "a"}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("three inputs not sorted")
	}
	// test with all empty input
	input = []string{"", "", "", "", "", "", "", "", "", ""}
	f(input)
	// test with peculiar input
	input = []string{"z", "m", "", "a", "d", "tt", "tt", "tt", "foo", "bar"}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("peculiar inputs not sorted")
	}
}

// testRepeated runs a given sort function over a sequence of
// repeated strings.
func testSortRepeated(t *testing.T, f func([]string), size int) {
	checkTestSize(t, size)
	input := make([]string, size)
	copy(input, repeatedStrings)
	f(input)
	if !isRepeated(input) {
		t.Error("repeated input not repeating")
	}
}

// testSortRepeatedCycle generates a repeating cycle of strings and
// runs the given sort on that data. The size is the number of elements
// to generate for the test.
func testSortRepeatedCycle(t *testing.T, f func([]string), size int) {
	checkTestSize(t, size)
	input := make([]string, size)
	copy(input, repeatedCycleStrings)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("repeated cycle input not sorted")
	}
}

// testSortRandom runs the given sort on a randomly generated data set
// consisting of strings of 100 letters and upper and lower case letters.
func testSortRandom(t *testing.T, f func([]string), size int) {
	checkTestSize(t, size)
	input := make([]string, size)
	copy(input, randomStrings)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("random input not sorted")
	}
}

// testSortDictWords generates a set of random pseudo-words and runs the
// given sort function on that set.
func testSortDictWords(t *testing.T, f func([]string), size int) {
	checkTestSize(t, size)
	input := make([]string, size)
	copy(input, uniqueWords)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("dictwords input not sorted")
	}
}

// testSortSorted runs the given sort function on an input set that
// is already in sorted order.
func testSortSorted(t *testing.T, f func([]string), size int) {
	checkTestSize(t, size)
	input := make([]string, size)
	copy(input, uniqueWords)
	sort.Strings(input)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("sorted dictwords input not sorted")
	}
}

// ReverseStringArray is identical to sort.StringArray except that the
// contents are sorted in reverse order.
type ReverseStringArray []string

func (p ReverseStringArray) Len() int           { return len(p) }
func (p ReverseStringArray) Less(i, j int) bool { return p[i] > p[j] }
func (p ReverseStringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// testSortReversed runs the given sort function on an input set that
// is in reverse sorted order.
func testSortReversed(t *testing.T, f func([]string), size int) {
	checkTestSize(t, size)
	input := make([]string, size)
	copy(input, uniqueWords)
	ri := ReverseStringArray(input)
	sort.Sort(ri)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("reversed dictwords input not sorted")
	}
}

// testSortNonUnique runs the given sort function on a set of words
// that are not necessarily unique (many will repeat a random number
// of times).
func testSortNonUnique(t *testing.T, f func([]string), size int) {
	checkTestSize(t, size)
	input := make([]string, size)
	copy(input, nonUniqueWords)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Error("non-unique words input not sorted")
	}
}

// checkTestSize compares the given size argument to the maximum
// allowable value, logging an error and failing the test if the
// value is too large.
func checkTestSize(t *testing.T, size int) {
	if size > largeDataSize {
		t.Error("size is larger than", largeDataSize)
	}
}

// shuffle randomly shuffles the elements in the string slice.
func shuffle(input []string) {
	n := len(input)
	indices := rand.Perm(n)
	for i := 0; i < n; i++ {
		j := indices[i]
		input[i], input[j] = input[j], input[i]
	}
}

// isRepeated tests if the array consists only of repeated strings.
func isRepeated(arr []string) bool {
	s := arr[0]
	for i, a := range arr {
		if a != s {
			fmt.Printf("%s != %s @ %d\n", a, s, i)
			return false
		}
	}
	return true
}
