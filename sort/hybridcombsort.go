//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

// HybridCombSort is an implementation of comb sort that delegates to
// insertion sort when the gap value has dropped below a certain
// threshold. This variation was proposed by David B. Ring of Palo Alto
// and demonstrated to be 10 to 15 percent faster than traditional comb
// sort. This particular implementation uses the Combsort11 variation
// for determining the gap values.
func HybridCombSort(a []string) {
	size := len(a)
	if a == nil || size < 2 {
		return
	}

	gap := size
	for gap > 8 {
		gap = (10 * gap) / 13
		if gap == 10 || gap == 9 {
			gap = 11
		}
		for i := 0; i+gap < size; i++ {
			j := i + gap
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	// At this point the input is nearly sorted, a case for which
	// insertion sort performs very well.
	InsertionSort(a)
}
