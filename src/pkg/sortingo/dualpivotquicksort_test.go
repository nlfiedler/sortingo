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
	testSortRepeated(t, DualPivotQuickSort, smallDataSize)
	testSortRepeatedCycle(t, DualPivotQuickSort, smallDataSize)
	testSortRandom(t, DualPivotQuickSort, smallDataSize)
	testSortDictWords(t, DualPivotQuickSort, smallDataSize)
	testSortReversed(t, DualPivotQuickSort, smallDataSize)
	testSortNonUnique(t, DualPivotQuickSort, smallDataSize)
}
