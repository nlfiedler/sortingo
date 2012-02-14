//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

import (
	"testing"
)

func TestBinaryInsertionSort(t *testing.T) {
	testSortArguments(t, BinaryInsertionSort)
	testSortRepeated(t, BinaryInsertionSort, mediumDataSize)
	testSortRepeatedCycle(t, BinaryInsertionSort, mediumDataSize)
	testSortRandom(t, BinaryInsertionSort, mediumDataSize)
	testSortDictWords(t, BinaryInsertionSort, mediumDataSize)
	testSortReversed(t, BinaryInsertionSort, mediumDataSize)
	testSortNonUnique(t, BinaryInsertionSort, mediumDataSize)
}
