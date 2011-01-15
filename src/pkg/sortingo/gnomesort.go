//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// $Id$
//

package sortingo

// GnomeSort is an implementation of the Gnome sort algorithm, based on
// pseudocode on Wikipedia. The running time is O(n^2).
func GnomeSort(input []string) {
	size := len(input)
        if input == nil || size < 2 {
		return
        }
        i := 1
        j := 2
        for i < size {
		if input[i - 1] <= input[i] {
			i = j
			j++
		} else {
			input[i - 1], input[i] = input[i], input[i - 1]
			i--
			if i == 0 {
				i = j
				j++
			}
		}
        }
}
