//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// SelectionSort is a basic selection sort implementation based on
// pseudocode found on Wikipedia. It sorts the input array in
// O(n^2) running time.
func SelectionSort(input []string) {
	size := len(input)
        if input == nil || size < 2 {
		return
        }

        for ii := 0; ii < size; ii++ {
		min := ii
		for jj := ii + 1; jj < size; jj++ {
			if input[jj] < input[min] {
				min = jj
			}
		}
		if ii != min {
			input[ii], input[min] = input[min], input[ii]
		}
	}
}
