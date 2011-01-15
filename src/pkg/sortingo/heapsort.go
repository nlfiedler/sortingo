//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// Binary heap sort implementation based on pseudocode from Wikipedia.
// Sort the input array using the heap sort algorithm.
// O(n*logn) running time with constant extra space.
func HeapSort(input []string) {
	size := len(input)
        if input == nil || size < 2 {
		return
        }

        // start is assigned the index in input of the last parent node
        for start := (size - 2) / 2; start >= 0; start-- {
		// sift down the node at index start to the proper place such
		// that all nodes below the start index are in heap order
		root := start
		// While the root has at least one child
		for root * 2 + 1 < size {
			// root*2+1 points to the left child
			child := root * 2 + 1
			// If the child has a sibling and the child's value
			// is less than its sibling's...
			if child + 1 < size && input[child] < input[child + 1] {
				// ... then point to the right child instead
				child++
			}
			// out of max-heap order
			if input[root] < input[child] {
				input[root], input[child] = input[child], input[root]
				// repeat to continue sifting down the child now
				root = child
			} else {
				break
			}
		}
        }
        // after sifting down the root all nodes/elements are in heap order

        for end := size - 1; end > 0; end-- {
		// swap the root (maximum value) of the heap with the last
		// element of the heap
		input[0], input[end] = input[end], input[0]
		// put the heap back in max-heap order
		root := 0
		// While the root has at least one child
		for root * 2 + 1 < end {
			// root*2+1 points to the left child
			child := root * 2 + 1
			// If the child has a sibling and the child's value is
			// less than its sibling's...
			if child + 1 < end && input[child] < input[child + 1] {
				// ... then point to the right child instead
				child++
			}
			// out of max-heap order
			if input[root] < input[child] {
				input[root], input[child] = input[child], input[root]
				// repeat to continue sifting down the child now
				root = child
			} else {
				break
			}
		}
		// end of for loop decreases the size of the heap by one so that
		// the previous max value will stay in its proper placement
        }
}
