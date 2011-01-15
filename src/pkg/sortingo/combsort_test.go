//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// $Id$
//

package sortingo

import (
	"testing"
)

func TestCombSort(t *testing.T) {
	testSortArguments(t, CombSort)
	testSortAnimals(t, CombSort)
	testSortRepeated(t, CombSort, smallDataSize)
	testSortRepeatedCycle(t, CombSort, smallDataSize)
	testSortRandom(t, CombSort, smallDataSize)
	testSortDictWords(t, CombSort, smallDataSize)
	testSortReversed(t, CombSort, smallDataSize)
	testSortNonUnique(t, CombSort, smallDataSize)
}
