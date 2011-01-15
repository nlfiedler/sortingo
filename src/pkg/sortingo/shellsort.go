//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// ShellSort is an implementation of shell sort based on pseudocode
// from Wikipedia. It sorts the input array using a gap sequence
// suggested by Gonnet and Baeza-Yates. Worst case running time is
// O(n^2) though often performs better in practice.
func ShellSort(input []string) {
	size := len(input)
        if input == nil || size < 2 {
		return
        }

        inc := size / 2
        for inc > 0 {
		for ii := inc; ii < size; ii++ {
			temp := input[ii]
			jj := ii
			for jj >= inc && input[jj - inc] > temp {
				input[jj] = input[jj - inc]
				jj -= inc
			}
			input[jj] = temp
		}
		// Another way of dividing by 2.2 to get an integer.
		if inc == 2 {
			inc = 1
		} else {
			inc = inc * 5 / 11
		}
        }
}
