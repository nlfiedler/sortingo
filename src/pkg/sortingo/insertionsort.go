//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// InsertionSort will sort the given slice of strings using the
// basic insertion sort algorithm, with O(n^2) running time.
func InsertionSort(a []string) {
	size := len(a)
        if a == nil || size < 2 {
		return;
        }

        for i := 1; i < size; i++ {
		pivot := a[i];
		j := i;
		for j > 0 && pivot < a[j - 1] {
			a[j] = a[j - 1];
			j--;
		}
		a[j] = pivot;
        }
}
