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
	testSortRepeated(t, IntroSort, mediumDataSize)
	testSortRepeatedCycle(t, IntroSort, mediumDataSize)
	testSortRandom(t, IntroSort, mediumDataSize)
	testSortDictWords(t, IntroSort, mediumDataSize)
	testSortReversed(t, IntroSort, mediumDataSize)
	testSortNonUnique(t, IntroSort, mediumDataSize)
}
