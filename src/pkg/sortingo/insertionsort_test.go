//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testSortArguments(t, InsertionSort)
	testSortRepeated(t, InsertionSort, smallDataSize)
	testSortRepeatedCycle(t, InsertionSort, smallDataSize)
	testSortRandom(t, InsertionSort, smallDataSize)
	testSortDictWords(t, InsertionSort, smallDataSize)
	testSortReversed(t, InsertionSort, smallDataSize)
	testSortNonUnique(t, InsertionSort, smallDataSize)
}
