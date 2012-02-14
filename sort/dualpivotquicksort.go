//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

// DualPivotQuickSort will sort the given slice of strings using the
// two pivot value quicksort variation by Vladimir Yaroslavskiy.
func DualPivotQuickSort(a []string) {
	size := len(a)
	if a == nil || size < 2 {
		return
	}
	dualPivotQuicksort(a, 0, size-1)
}

// dualPivotQuicksort is a variation of quicksort that uses two pivot
// values rather than one, and was created by Vladimir Yaroslavskiy.
// This Go implementation is a translation of the original Java code,
// with some simplification of the code.
func dualPivotQuicksort(a []string, left int, right int) {
	len := right - left

	// perform insertion sort on small ranges
	if len < 17 {
		for i := left + 1; i <= right; i++ {
			for j := i; j > left && a[j] < a[j-1]; j-- {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
		return
	}

	// compute indices of medians
	sixth := len / 6
	m1 := left + sixth
	m2 := m1 + sixth
	m3 := m2 + sixth
	m4 := m3 + sixth
	m5 := m4 + sixth

	// order the medians in preparation for partitioning
	if a[m1] > a[m2] {
		a[m1], a[m2] = a[m2], a[m1]
	}
	if a[m4] > a[m5] {
		a[m4], a[m5] = a[m5], a[m4]
	}
	if a[m1] > a[m3] {
		a[m1], a[m3] = a[m3], a[m1]
	}
	if a[m2] > a[m3] {
		a[m2], a[m3] = a[m3], a[m2]
	}
	if a[m1] > a[m4] {
		a[m1], a[m4] = a[m4], a[m1]
	}
	if a[m3] > a[m4] {
		a[m3], a[m4] = a[m4], a[m3]
	}
	if a[m2] > a[m5] {
		a[m2], a[m5] = a[m5], a[m2]
	}
	if a[m2] > a[m3] {
		a[m2], a[m3] = a[m3], a[m2]
	}
	if a[m4] > a[m5] {
		a[m4], a[m5] = a[m5], a[m4]
	}

	// select the pivots such that [ < pivot1 | pivot1 <= && <= pivot2 | > pivot2 ]
	pivot1 := a[m2]
	pivot2 := a[m4]

	diffPivots := pivot1 != pivot2

	// move the pivots out of the away
	a[m2] = a[left]
	a[m4] = a[right]
	less := left + 1
	great := right - 1

	// partition the elements
	if diffPivots {
		for k := less; k <= great; k++ {
			x := a[k]
			if x > pivot2 {
				for a[great] > pivot2 && k < great {
					great--
				}
				a[k] = a[great]
				a[great] = x
				great--
				x = a[k]
			}
			if x < pivot1 {
				a[k] = a[less]
				a[less] = x
				less++
			}
		}
	} else {
		for k := less; k <= great; k++ {
			x := a[k]
			if x == pivot1 {
				continue
			}
			if x > pivot1 {
				for a[great] > pivot2 && k < great {
					great--
				}
				a[k] = a[great]
				a[great] = x
				great--
				x = a[k]
			}
			if x < pivot1 {
				a[k] = a[less]
				a[less] = x
				less++
			}
		}
	}

	// swap the pivots back into position
	a[left] = a[less-1]
	a[less-1] = pivot1
	a[right] = a[great+1]
	a[great+1] = pivot2

	// recursively sort the left and right partitions
	dualPivotQuicksort(a, left, less-2)
	dualPivotQuicksort(a, great+2, right)

	// order the equal elements in the middle
	if great-less > len-13 && diffPivots {
		for k := less; k <= great; k++ {
			x := a[k]
			if x == pivot2 {
				a[k] = a[great]
				a[great] = x
				great--
				x = a[k]
			}
			if x == pivot1 {
				a[k] = a[less]
				a[less] = x
				less++
			}
		}
	}

	// recursively sort the middle partition
	if diffPivots {
		dualPivotQuicksort(a, less, great)
	}
}
