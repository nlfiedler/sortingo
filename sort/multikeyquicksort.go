//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

// As with GCC std::sort, delegate to insertion sort for sizes below 16.
const insertionThreshold = 16

// MultikeyQuickSort is a translation from the original C implementation
// by J. Bentley and R. Sedgewick, from their "Fast algorithms for sorting
// and searching strings" paper published in 1997.
//
// Sorts the slice of strings using a multikey quicksort that chooses
// a pivot point using a "median of three" rule (or pseudo median of
// nine for slices over a certain threshold). For very small slices,
// a simple insertion sort is used.
func MultikeyQuickSort(a []string) {
	MultikeyQuickSortDepth(a, 0)
}

// multikeyQuickSortDepth is like MultikeyQuickSort but it only considers
// the characters in the strings starting from the given offset (depth).
func MultikeyQuickSortDepth(a []string, depth int) {
	n := len(a)
	if n < insertionThreshold {
		insertionSortDepth(a, depth)
		return
	}

	// Find the median of three to determine our pivot value.
	pl := 0
	pm := n / 2
	pn := n - 1
	var r int
	if n > 30 {
		// On larger slices, find a pseudo median of nine elements.
		d := n / 8
		pl = med3(a, 0, d, 2*d, depth)
		pm = med3(a, n/2-d, pm, n/2+d, depth)
		pn = med3(a, n-1-2*d, n-1-d, pn, depth)
	}
	pm = med3(a, pl, pm, pn, depth)

	// Move the pivot to the start of the slice.
	a[0], a[pm] = a[pm], a[0]

	v := int(charAt(a[0], depth))
	var allzeros bool = v == 0
	le, lt := 1, 1
	gt := n - 1
	ge := gt
	for {
		// Move elements smaller than pivot to the left.
		for ; lt <= gt; lt++ {
			r = int(charAt(a[lt], depth)) - v
			if r > 0 {
				break
			} else if r == 0 {
				a[le], a[lt] = a[lt], a[le]
				le++
			} else {
				allzeros = false
			}
		}

		// Move elements larger than pivot to the right.
		for ; lt <= gt; gt-- {
			r = int(charAt(a[gt], depth)) - v
			if r < 0 {
				break
			} else if r == 0 {
				a[gt], a[ge] = a[ge], a[gt]
				ge--
			} else {
				allzeros = false
			}
		}
		if lt > gt {
			break
		}
		a[lt], a[gt] = a[gt], a[lt]
		lt++
		gt--
	}

	pn = n
	r = iMin(le-0, lt-le)
	vecswap(a, 0, lt-r, r)
	r = iMin(ge-gt, pn-ge-1)
	vecswap(a, lt, pn-r, r)
	r = lt - le
	if r > 1 {
		MultikeyQuickSortDepth(a[:r], depth)
	}
	if !allzeros {
		// Only descend if there was at least one string that was
		// of equal or greater length than current depth.
		MultikeyQuickSortDepth(a[r:r+le+n-ge-1], depth+1)
	}
	r = ge - gt
	if r > 1 {
		MultikeyQuickSortDepth(a[n-r:], depth)
	}
}

// Swap the elements between to areas within a slice.
func vecswap(input []string, src, dst, count int) {
	for count > 0 {
		input[src], input[dst] = input[dst], input[src]
		src++
		dst++
		count--
	}
}

// Find the median of three characters, found in the given strings
// at character position 'depth'. One of the three integer values
// (low, med, high) will be returned based on the comparisons.
func med3(a []string, low, med, high, depth int) int {
	va := charAt(a[low], depth)
	vb := charAt(a[med], depth)
	if va == vb {
		return low
	}
	vc := charAt(a[high], depth)
	if vc == va || vc == vb {
		return high
	}
	if va < vb {
		if vb < vc {
			return med
		} else if va < vc {
			return high
		} else {
			return low
		}
	} else {
		if vb > vc {
			return med
		} else if va < vc {
			return low
		} else {
			return high
		}
	}
	panic("unreachable")
}
