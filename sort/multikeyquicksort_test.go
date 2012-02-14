//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

import (
	"testing"
)

func TestMultikeyQuickSort(t *testing.T) {
	testSortArguments(t, MultikeyQuickSort)
	testSortRepeated(t, MultikeyQuickSort, largeDataSize)
	testSortRepeatedCycle(t, MultikeyQuickSort, largeDataSize)
	testSortRandom(t, MultikeyQuickSort, largeDataSize)
	testSortDictWords(t, MultikeyQuickSort, largeDataSize)
	testSortReversed(t, MultikeyQuickSort, largeDataSize)
	testSortNonUnique(t, MultikeyQuickSort, largeDataSize)
}
