//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// BinaryInsertionSort is an implementation of the binary insertion sort
// algorithm borrowed from timsort, with some minor modifications.
// It requires O(n log n) compares, but O(n^2) data movement (worst case).
func BinaryInsertionSort(arr []string) {
	BinaryInsertionSortDepth(arr, 0)
}

// BinaryInsertionSortDepth is identical to BinaryInsertionSort but takes
// a depth value which indicates the portion of the strings that is to be
// used in sorting (that is, ignoring the characters from 0 to depth).
func BinaryInsertionSortDepth(arr []string, depth int) {
	size := len(arr)
        if arr == nil || size < 2 || depth < 0 {
		return
        }
        for ii := 0; ii < size; ii++ {
		pivot := arr[ii]

		// Set left (and right) to the index where a[start] (pivot) belongs
		left := 0
		right := ii
		// Invariants:
		//   pivot >= all in [lo, left).
		//   pivot <  all in [right, start).
		for left < right {
			mid := (left + right) >> 1
			if compareTail(pivot, arr[mid], depth) < 0 {
				right = mid
			} else {
				left = mid + 1
			}
		}

		// The invariants above still hold, so pivot belongs at left.
		// Note that if there are elements equal to pivot, left points
		// to the first slot after them -- that's why this sort is stable.
		// Slide elements over to make room for the pivot.
		count := ii - left
		// Switch is just an optimization for arraycopy in default case.
		switch count {
                case 2:
			arr[left + 2] = arr[left + 1]
			fallthrough
                case 1:
			arr[left + 1] = arr[left]
                default:
			copy(arr[left + 1:], arr[left:left + count])
		}
		arr[left] = pivot
        }
}

// compareTail compares two strings, starting with the characters at
// offset 'depth' (assumes the leading characters are the same in both
// sequences). Returns a negative integer, zero, or a positive integer as
// the first argument is less than, equal to, or greater than the second.
func compareTail(a, b string, depth int) int {
        idx := depth
	var s, t uint8
        if idx < len(a) { s = a[idx] }
        if idx < len(b) { t = b[idx] }
        for s == t && idx < len(a) {
		idx++
		if s = 0; idx < len(a) { s = a[idx] }
		if t = 0; idx < len(b) { t = b[idx] }
        }
	// Convert unsigned to signed so we can return negatives.
        return int(s) - int(t)
}
