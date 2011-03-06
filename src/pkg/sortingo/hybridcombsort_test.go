//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestHybridCombSort(t *testing.T) {
	testSortArguments(t, HybridCombSort)
	testSortRepeated(t, HybridCombSort, smallDataSize)
	testSortRepeatedCycle(t, HybridCombSort, smallDataSize)
	testSortRandom(t, HybridCombSort, smallDataSize)
	testSortDictWords(t, HybridCombSort, smallDataSize)
	testSortReversed(t, HybridCombSort, smallDataSize)
	testSortNonUnique(t, HybridCombSort, smallDataSize)
}
