//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestBinaryInsertionSort(t *testing.T) {
	testSortArguments(t, BinaryInsertionSort)
	testSortAnimals(t, BinaryInsertionSort)
	testSortRepeated(t, BinaryInsertionSort, smallDataSize)
	testSortRepeatedCycle(t, BinaryInsertionSort, smallDataSize)
 	testSortRandom(t, BinaryInsertionSort, smallDataSize)
  	testSortDictWords(t, BinaryInsertionSort, largeDataSize)
	testSortReversed(t, BinaryInsertionSort, smallDataSize)
	testSortNonUnique(t, BinaryInsertionSort, largeDataSize)
}
