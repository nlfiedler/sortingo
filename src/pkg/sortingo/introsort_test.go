//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestIntroSort(t *testing.T) {
	testSortArguments(t, IntroSort)
	testSortRepeated(t, IntroSort, smallDataSize)
	testSortRepeatedCycle(t, IntroSort, smallDataSize)
	testSortRandom(t, IntroSort, smallDataSize)
	testSortDictWords(t, IntroSort, smallDataSize)
	testSortReversed(t, IntroSort, smallDataSize)
	testSortNonUnique(t, IntroSort, smallDataSize)
}
