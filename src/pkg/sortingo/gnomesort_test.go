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

func TestGnomeSort(t *testing.T) {
	testSortArguments(t, GnomeSort)
	testSortAnimals(t, GnomeSort)
	testSortRepeated(t, GnomeSort, smallDataSize)
	testSortRepeatedCycle(t, GnomeSort, smallDataSize)
	testSortRandom(t, GnomeSort, smallDataSize)
	testSortDictWords(t, GnomeSort, smallDataSize)
	testSortReversed(t, GnomeSort, smallDataSize)
	testSortNonUnique(t, GnomeSort, smallDataSize)
}
