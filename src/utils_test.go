//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// This file contains utility functions for the unit tests.

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
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
	input, err := loadFileData("data/dictwords", false, size)
	if err != nil {
		t.Log(err.String())
		t.Fail()
	}
	shuffle(input)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("dictwords input not sorted")
		t.Fail()
	}
}

// testSortSorted runs the given sort function on an input set that
// is already in sorted order.
func testSortSorted(t *testing.T, f func ([]string), size int) {
	input, err := loadFileData("data/dictwords", false, size)
	if err != nil {
		t.Log(err.String())
		t.Fail()
	}
	sort.SortStrings(input)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("dictwords input not sorted")
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
	input, err := loadFileData("data/dictwords", false, size)
	if err != nil {
		t.Log(err.String())
		t.Fail()
	}
	ri := ReverseStringArray(input)
	sort.Sort(ri)
	f(input)
	if !sort.StringsAreSorted(input) {
		t.Log("dictwords input not sorted")
		t.Fail()
	}
}

// testSortHamletWords runs the given sort function on the set of all
// words appearing in Shakespeare's <<Hamlet>>.
func testSortHamletWords(t *testing.T, f func ([]string), size int) {
	input, err := loadFileData("data/hamletwords", false, size)
	if err != nil {
		t.Log(err.String())
		t.Fail()
	}
	shuffle(input)
	f(input);
	if !sort.StringsAreSorted(input) {
		t.Log("hamlet words input not sorted")
		t.Fail()
	}
}

// testSortDictCalls runs the given sort function on a set of library calls.
func testSortDictCalls(t *testing.T, f func ([]string), size int) {
	input, err := loadFileData("data/dictcalls.gz", true, size)
	if err != nil {
		t.Log(err.String())
		t.Fail()
	}
	f(input);
	if !sort.StringsAreSorted(input) {
		t.Log("dictionary calls input not sorted")
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

// loadFileData loads the named file, splitting on the newline character
// and returning the result as a slice of strings. The count is the number
// of lines to read from the file.
func loadFileData(file string, compressed bool, count int) (list []string, err os.Error) {
	var f io.ReadCloser
	f, err = os.Open(file, os.O_RDONLY, 0400)
	if err != nil {
		return
	}
	defer f.Close()
	if (compressed) {
		f, err = gzip.NewReader(f)
		if err != nil {
			return
		}
		defer f.Close()
	}
	br := bufio.NewReader(f)
	list = make([]string, 0, iMin(count, 8192))
	for {
		line, err := br.ReadString('\n')
		if err == nil {
			line = strings.TrimRight(line, "\n")
		} else if err != os.EOF {
			return
		}
		list = append(list, line)
		if err == os.EOF {
			break
		}
		if count--; count == 0 {
			break
		}
	}
	return
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
	list := make([]string, 0, n)
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
		list = append(list, string(bb.Bytes()))
	}
	return list
}
