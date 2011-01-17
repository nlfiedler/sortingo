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
// largeDataSize is the recommended test size for the faster sorting
// algorithms (e.g. burstsort, funnelsort).
const largeDataSize = 65536

// testSortArguments runs a given sort function with the most
// basic of input sequences in order to test its robustness.
func testSortArguments(t *testing.T, f func ([]string)) {
	// these should silently do nothing
	f(nil)
	f(make([]string, 0))
	f([]string{"a"})
	// test the bare minimum input, two elements
	input := []string{"b", "a"}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("two inputs not sorted")
		t.Fail()
	}
	// now try three elements
	input = []string{"c", "b", "a"}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("three inputs not sorted")
		t.Fail()
	}
	// test with all empty input
	input = []string{"", "", "", "", "", "", "", "", "", ""}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("empty inputs not sorted")
		t.Fail()
	}
	// test with peculiar input
	input = []string{"z", "m", "", "a", "d", "tt", "tt", "tt", "foo", "bar"}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("peculiar inputs not sorted")
		t.Fail()
	}
}

// testSortAnimals runs a given sort function with the simple input
// of a set of animal names.
func testSortAnimals(t *testing.T, f func ([]string)) {
	input := []string{"elephant", "zebra", "giraffe", "monkey", "gazelle"}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("animal input not sorted")
		t.Fail()
	}
}

// testRepeated runs a given sort function over a sequence of
// repeated strings.
func testSortRepeated(t *testing.T, f func ([]string), size int) {
	input := make([]string, size)
	str := strings.Repeat("A", 100)
	for idx, _ := range input {
		input[idx] = str
	}
	f(input)
	if !isRepeated(input, str) {
		t.Log("repeated input not repeating")
		t.Fail()
	}
}

// testSortRepeatedCycle generates a repeating cycle of strings and
// runs the given sort on that data. The size is the number of elements
// to generate for the test.
func testSortRepeatedCycle(t *testing.T, f func ([]string), size int) {
	var strs [100]string
	seed := strings.Repeat("A", 100)
	for i, l := 0, 1; i < len(strs); i, l = i + 1, l + 1 {
		strs[i] = seed[0:l]
	}
	input := make([]string, size)
	for c, i := 0, 0; c < len(input); i, c = i + 1, c + 1 {
		input[c] = strs[i % len(strs)]
	}
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("repeated cycle input not sorted")
		t.Fail()
	}
}

// testSortRandom runs the given sort on a randomly generated data set.
func testSortRandom(t *testing.T, f func ([]string), size int) {
	input := generateData(size, 100)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("random input not sorted")
		t.Fail()
	}
}

// testSortDictWords loads the dictionary words file, shuffles it,
// and runs the given sort function on the result.
func testSortDictWords(t *testing.T, f func ([]string), size int) {
	input := generateUnique(size)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("dictwords input not sorted")
		t.Fail()
	}
}

// testSortSorted runs the given sort function on an input set that
// is already in sorted order.
func testSortSorted(t *testing.T, f func ([]string), size int) {
	input := generateUnique(size)
	sort.SortStrings(input)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("sorted dictwords input not sorted")
		t.Fail()
	}
}

// ReverseStringArray is identical to sort.StringArray except that the
// contents are sorted in reverse order.
type ReverseStringArray []string
func (p ReverseStringArray) Len() int		{ return len(p) }
func (p ReverseStringArray) Less(i, j int) bool { return p[i] > p[j] }
func (p ReverseStringArray) Swap(i, j int)	{ p[i], p[j] = p[j], p[i] }

// testSortReversed runs the given sort function on an input set that
// is in reverse sorted order.
func testSortReversed(t *testing.T, f func ([]string), size int) {
	input := generateUnique(size)
	ri := ReverseStringArray(input)
	sort.Sort(ri)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("reversed dictwords input not sorted")
		t.Fail()
	}
}

// testSortHamletWords runs the given sort function on the set of all
// words appearing in Shakespeare's <<Hamlet>>.
func testSortNonUnique(t *testing.T, f func ([]string), size int) {
	input := generateNonUnique(size)
	f(input);
	if !sort.StringsAreSorted(input) {
		t.Log("non-unique words input not sorted")
		t.Fail()
	}
}

// shuffle randomly shuffles the elements in the string slice.
func shuffle(input []string) {
	indices := rand.Perm(len(input))
	for i := 0; i < len(input); i++ {
		j := indices[i]
		input[i], input[j] = input[j], input[i]
	}
}

// iMin returns the minimum of x and y.
func iMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// isRepeated tests if the array consists only of the s string repeated.
func isRepeated(arr []string, s string) bool {
	for i, a := range arr {
		if a != s {
			fmt.Printf("%s != %s @ %d\n", a, s, i)
			return false
		}
	}
	return true
}

// generateData generates a set of n random strings, each of length l.
// Each string consists of the digits 0 through 9 and upper and lower
// case letters (a..z, A..Z).
func generateData(n, l int) []string {
	list := make([]string, n)
	for ii := 0; ii < n; ii++ {
		bb := bytes.NewBuffer(make([]byte, 0, l))
		for jj := 0; jj < l; jj++ {
			d := rand.Intn(62)
			if (d < 10) {
				bb.WriteRune('0' + d)
			} else if (d < 36) {
				bb.WriteRune('A' + (d - 10))
			} else {
				bb.WriteRune('a' + (d - 36))
			}
		}
		list[ii] = string(bb.Bytes())
	}
	return list
}

// generateUnique generates a set of psuedo words that are unique
// (as in a dictionary).
func generateUnique(n int) []string {
	list := make([]string, 0, n)
	words := make(map[string]bool)
	for ii := 0; ii < n; {
		// Each word is from 1 to 28 characters long.
		l := 1 + rand.Intn(27)
		bb := bytes.NewBuffer(make([]byte, 0, l))
		// Each word consists only of the lowercase letters.
		for jj := 0; jj < l; jj++ {
			d := rand.Intn(26)
			bb.WriteRune('a' + d)
		}
		s := string(bb.Bytes())
		if !words[s] {
			list = append(list, s)
			words[s] = true
			ii++
		}
	}
	return list
}

// generateNonUnique generates a set of psuedo words that may be
// repeated numerous times.
func generateNonUnique(n int) []string {
	list := make([]string, 0, n)
	for cc := 0; cc < n; {
		// Each word is up to 28 characters long.
		l := rand.Intn(28)
		bb := bytes.NewBuffer(make([]byte, 0, l))
		// Each word consists only of the lowercase letters.
		for jj := 0; jj < l; jj++ {
			d := rand.Intn(26)
			bb.WriteRune('a' + d)
		}
		// Repeat the word some number of times.
		c := rand.Intn(100)
		if c > (n - cc) {
			c = n - cc
		}
		s := string(bb.Bytes())
		for jj := 0; jj < c; jj++ {
			list = append(list, s)
		}
		cc += c
	}
	return list
}
