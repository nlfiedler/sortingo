//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestCombSort(t *testing.T) {
	testSortArguments(t, CombSort)
	testSortRepeated(t, CombSort, mediumDataSize)
	testSortRepeatedCycle(t, CombSort, mediumDataSize)
	testSortRandom(t, CombSort, mediumDataSize)
	testSortDictWords(t, CombSort, mediumDataSize)
	testSortReversed(t, CombSort, mediumDataSize)
	testSortNonUnique(t, CombSort, mediumDataSize)
}
