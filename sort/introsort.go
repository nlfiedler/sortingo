//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

// Implementation of the introspective sort algorithm, developed by
// David Musser; implementation copied from the paper on introsort
// by Ralph Unden, with some modifications.

import (
	"math"
)

// IntroSort sorts the array of strings using an introspective sort
// algorithm, so expect O(log(n)) running time.
func IntroSort(a []string) {
	size := len(a)
	if a == nil || size < 2 {
		return
	}
	floor := int(math.Floor(math.Log2(float64(size))))
	introsortLoop(0, size, 2*floor, a)
	insertionsort(0, size, a)
}

// introsortLoop is a modified quicksort that delegates to heapsort when
// the depth limit has been reached. Does nothing if the range is below
// the threshold.
func introsortLoop(low, high, depth_limit int, a []string) {
	for high-low > 16 {
		if depth_limit == 0 {
			// perform a basic heap sort
			n := high - low
			for i := n / 2; i >= 1; i-- {
				d := a[low+i-1]
				j := i
				for j <= n/2 {
					child := 2 * j
					if child < n && a[low+child-1] < a[low+child] {
						child++
					}
					if d >= a[low+child-1] {
						break
					}
					a[low+j-1] = a[low+child-1]
					j = child
				}
				a[low+j-1] = d
			}
			for i := n; i > 1; i-- {
				a[low], a[low+i-1] = a[low+i-1], a[low]
				d := a[low+i-1]
				j := 1
				m := i - 1
				for j <= m/2 {
					child := 2 * j
					if child < m && a[low+child-1] < a[low+child] {
						child++
					}
					if d >= a[low+child-1] {
						break
					}
					a[low+j-1] = a[low+child-1]
					j = child
				}
				a[low+j-1] = d
			}
			return
		}
		depth_limit--
		p := introsortPartition(low, high, introsortMedian(low, low+((high-low)/2)+1, high-1, a), a)
		introsortLoop(p, high, depth_limit, a)
		high = p
	}
}

// Partitions the elements in the given range such that elements
// less than the pivot appear before those greater than the pivot.
func introsortPartition(low, high int, x string, a []string) int {
	i := low
	j := high
	for {
		for a[i] < x {
			i++
		}
		j--
		for x < a[j] {
			j--
		}
		if i >= j {
			return i
		}
		a[i], a[j] = a[j], a[i]
		i++
	}
	panic("unreachable")
}

// introsortMedian finds the median of three elements in the given range.
func introsortMedian(low, mid, high int, a []string) string {
	if a[mid] < a[low] {
		if a[high] < a[mid] {
			return a[mid]
		} else {
			if a[high] < a[low] {
				return a[high]
			} else {
				return a[low]
			}
		}
	} else {
		if a[high] < a[mid] {
			if a[high] < a[low] {
				return a[low]
			} else {
				return a[high]
			}
		} else {
			return a[mid]
		}
	}
	panic("unreachable")
}

// insertionsort performs a simple insertion sort that operates on the
// given range.
func insertionsort(low, high int, a []string) {
	for i := low; i < high; i++ {
		j := i
		t := a[i]
		for j != low && t < a[j-1] {
			a[j] = a[j-1]
			j--
		}
		a[j] = t
	}
}
