//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

// charAt retrieves the character in string s at offset d. If d is
// greater than or equal to the length of the string, return zero.
// This simulates fixed-length strings that are zero-padded.
func charAt(s string, d int) uint8 {
	if d < len(s) {
		return s[d]
	}
	return 0
}

// iMax returns the maximum of x and y.
func iMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// iMin returns the minimum of x and y.
func iMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
