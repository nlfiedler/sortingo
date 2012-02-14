//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testSortArguments(t, SelectionSort)
	testSortRepeated(t, SelectionSort, smallDataSize)
	testSortRepeatedCycle(t, SelectionSort, smallDataSize)
	testSortRandom(t, SelectionSort, smallDataSize)
	testSortDictWords(t, SelectionSort, smallDataSize)
	testSortReversed(t, SelectionSort, smallDataSize)
	testSortNonUnique(t, SelectionSort, smallDataSize)
}
