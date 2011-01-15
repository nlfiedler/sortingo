//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// $Id$
//

package sortingo

// InsertionSort will sort the given slice of strings using the
// basic insertion sort algorithm, with O(n^2) running time.
func InsertionSort(a []string) {
	InsertionSortRange(a, 0, len(a) - 1)
}

// InsertionSortRange will sort the specific range of elements in
// the slice of strings. The low and high values are inclusive.
func InsertionSortRange(a []string, low int, high int) {
        if a == nil || len(a) < 2 || low < 0 || high <= low {
		return;
        }

        for i := low + 1; i <= high; i++ {
		pivot := a[i];
		j := i;
		for j > low && pivot < a[j - 1] {
			a[j] = a[j - 1];
			j--;
		}
		a[j] = pivot;
        }
}
