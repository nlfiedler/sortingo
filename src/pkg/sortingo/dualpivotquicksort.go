//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// DualPivotQuickSort will sort the given slice of strings using the
// two pivot value quicksort variation by Vladimir Yaroslavskiy.
func DualPivotQuickSort(a []string) {
	dualPivotQuicksort(a, 0, len(a)-1)
}

// dualPivotQuicksort is a variation of quicksort that uses two pivot
// values rather than one, and was created by Vladimir Yaroslavskiy.
// This Go implementation is a translation of the original Java code.
func dualPivotQuicksort(a []string, left int, right int) {
	len := right - left

	// insertion sort on tiny array
	if len < 17 {
		for i := left + 1; i <= right; i++ {
			for j := i; j > left && a[j] < a[j-1]; j-- {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
		return
	}

	// median indexes
	sixth := len / 6
	m1 := left + sixth
	m2 := m1 + sixth
	m3 := m2 + sixth
	m4 := m3 + sixth
	m5 := m4 + sixth

	// 5-element sorting network
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

	// pivots: [ < pivot1 | pivot1 <= && <= pivot2 | > pivot2 ]
	pivot1 := a[m2]
	pivot2 := a[m4]

	diffPivots := pivot1 != pivot2

	a[m2] = a[left]
	a[m4] = a[right]

	// center part pointers
	less := left + 1
	great := right - 1

	// sorting
	if diffPivots {
		for k := less; k <= great; k++ {
			x := a[k]

			if x < pivot1 {
				a[k] = a[less]
				a[less] = x
				less++
			} else if x > pivot2 {
				for a[great] > pivot2 && k < great {
					great--
				}
				a[k] = a[great]
				a[great] = x
				great--
				x = a[k]

				if x < pivot1 {
					a[k] = a[less]
					a[less] = x
					less++
				}
			}
		}
	} else {
		for k := less; k <= great; k++ {
			x := a[k]

			if x == pivot1 {
				continue
			}
			if x < pivot1 {
				a[k] = a[less]
				a[less] = x
				less++
			} else {
				for a[great] > pivot2 && k < great {
					great--
				}
				a[k] = a[great]
				a[great] = x
				great--
				x = a[k]

				if x < pivot1 {
					a[k] = a[less]
					a[less] = x
					less++
				}
			}
		}
	}

	// swap
	a[left] = a[less-1]
	a[less-1] = pivot1
	a[right] = a[great+1]
	a[great+1] = pivot2

	// left and right parts
	dualPivotQuicksort(a, left, less-2)
	dualPivotQuicksort(a, great+2, right)

	// equal elements
	if great-less > len-13 && diffPivots {
		for k := less; k <= great; k++ {
			x := a[k]

			if x == pivot1 {
				a[k] = a[less]
				a[less] = x
				less++
			} else if x == pivot2 {
				a[k] = a[great]
				a[great] = x
				great--
				x = a[k]

				if x == pivot1 {
					a[k] = a[less]
					a[less] = x
					less++
				}
			}
		}
	}

	// center part
	if diffPivots {
		dualPivotQuicksort(a, less, great)
	}
}
