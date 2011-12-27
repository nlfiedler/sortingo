//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	testSortArguments(t, ShellSort)
	testSortRepeated(t, ShellSort, mediumDataSize)
	testSortRepeatedCycle(t, ShellSort, mediumDataSize)
	testSortRandom(t, ShellSort, mediumDataSize)
	testSortDictWords(t, ShellSort, mediumDataSize)
	testSortReversed(t, ShellSort, mediumDataSize)
	testSortNonUnique(t, ShellSort, mediumDataSize)
}
