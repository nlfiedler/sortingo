//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestDualPivotQuickSort(t *testing.T) {
	testSortArguments(t, DualPivotQuickSort)
	testSortRepeated(t, DualPivotQuickSort, mediumDataSize)
	testSortRepeatedCycle(t, DualPivotQuickSort, mediumDataSize)
	testSortRandom(t, DualPivotQuickSort, mediumDataSize)
	testSortDictWords(t, DualPivotQuickSort, mediumDataSize)
	testSortReversed(t, DualPivotQuickSort, mediumDataSize)
	testSortNonUnique(t, DualPivotQuickSort, mediumDataSize)
}
