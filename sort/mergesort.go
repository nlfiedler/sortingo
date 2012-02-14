//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

// MergeSort will sort the given slice of strings using the
// basic merge sort algorithm, with O(n log n) running time.
func MergeSort(a []string) {
	size := len(a)
	if a == nil || size < 2 {
		return
	}

	// for small sets, delegate to insertion sort
	if size < 7 {
		InsertionSort(a)
		return
	}

	// recursively sort the left and right sides
	middle := size / 2
	left := a[:middle]
	right := a[middle:]
	MergeSort(left)
	MergeSort(right)

	// merge the sorted halves into the result
	result := make([]string, 0, size)
	li := 0
	ls := len(left)
	ri := 0
	rs := len(right)
	for li < ls && ri < rs {
		if left[li] <= right[ri] {
			result = append(result, left[li])
			li++
		} else {
			result = append(result, right[ri])
			ri++
		}
	}
	if li < ls {
		result = append(result, left[li:]...)
	} else if ri < rs {
		result = append(result, right[ri:]...)
	}

	// copy into the original array
	copy(a, result)
}
