//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []string{"elephant", "zebra", "giraffe", "monkey", "gazelle"}
	InsertionSort(input)
	if !sort.StringsAreSorted(input) {
		t.Log("small input not sorted")
		t.Fail()
	}
}
