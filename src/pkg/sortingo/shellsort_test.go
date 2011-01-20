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
	testSortRepeated(t, ShellSort, smallDataSize)
	testSortRepeatedCycle(t, ShellSort, smallDataSize)
	testSortRandom(t, ShellSort, smallDataSize)
	testSortDictWords(t, ShellSort, smallDataSize)
	testSortReversed(t, ShellSort, smallDataSize)
	testSortNonUnique(t, ShellSort, smallDataSize)
}
