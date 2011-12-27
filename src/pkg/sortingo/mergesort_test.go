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
	testSortRepeated(t, MergeSort, mediumDataSize)
	testSortRepeatedCycle(t, MergeSort, mediumDataSize)
	testSortRandom(t, MergeSort, mediumDataSize)
	testSortDictWords(t, MergeSort, mediumDataSize)
	testSortReversed(t, MergeSort, mediumDataSize)
	testSortNonUnique(t, MergeSort, mediumDataSize)
}
