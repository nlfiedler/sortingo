//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	testSortArguments(t, HeapSort)
	testSortRepeated(t, HeapSort, mediumDataSize)
	testSortRepeatedCycle(t, HeapSort, mediumDataSize)
	testSortRandom(t, HeapSort, mediumDataSize)
	testSortDictWords(t, HeapSort, mediumDataSize)
	testSortReversed(t, HeapSort, mediumDataSize)
	testSortNonUnique(t, HeapSort, mediumDataSize)
}
