//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// CombSort will sort the given slice of strings using the
// Comb sort algorithm, namely the Combsort11 variation.
// Its running time is O(n^2) though often does better than
// similar algorithms.
func CombSort(input []string) {
	size := len(input)
        if input == nil || size < 2 {
		return
        }

        gap := size //initialize gap size
        swapped := true

        for gap > 1 || swapped {
		// Update the gap value for the next comb.
		if gap > 1 {
			gap = int(float(gap) / 1.3)
			if gap == 10 || gap == 9 {
				gap = 11
			}
		}

		// a single "comb" over the input list
		swapped = false
		for i := 0; i + gap < size; i++ {
			if input[i] > input[i + gap] {
				input[i], input[i + gap] = input[i + gap], input[i]
				// Signal that the list is not guaranteed sorted.
				swapped = true
			}
		}
        }
}
