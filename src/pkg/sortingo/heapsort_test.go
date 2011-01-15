//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	testSortArguments(t, HeapSort)
	testSortAnimals(t, HeapSort)
	testSortRepeated(t, HeapSort, smallDataSize)
	testSortRepeatedCycle(t, HeapSort, smallDataSize)
	testSortRandom(t, HeapSort, smallDataSize)
	testSortDictWords(t, HeapSort, smallDataSize)
	testSortReversed(t, HeapSort, smallDataSize)
	testSortNonUnique(t, HeapSort, smallDataSize)
}
