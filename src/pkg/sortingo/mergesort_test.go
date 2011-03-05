//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	testSortArguments(t, MergeSort)
	testSortRepeated(t, MergeSort, smallDataSize)
	testSortRepeatedCycle(t, MergeSort, smallDataSize)
	testSortRandom(t, MergeSort, smallDataSize)
	testSortDictWords(t, MergeSort, smallDataSize)
	testSortReversed(t, MergeSort, smallDataSize)
	testSortNonUnique(t, MergeSort, smallDataSize)
}
