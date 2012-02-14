//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sort

// alphabetSize is the number of characters supported for the trie used
// in sorting (strings are treated as arrays of uint8 values).
const alphabetSize = 256

//
// To support multi-byte character sets, would need a larger alphabet
// size and an efficient, sparse array for indexing the runes. Instead
// of accessing bytes in strings using array access, would need to cast
// the strings to []rune and access the runes, using those as indices
// into the trie nodes.
//
// Would also need to switch out multikey quick sort for something that
// works with runes, or possibly a comparator based sort.
//

// nullterm represents the null terminator character.
const nullterm = '\000'

// threshold is the maximum number of elements in any given bucket; for
// null bucket set, this is the size of each of the chained buckets.
const threshold = 8192

// thresholdMinusOne is used to store reference to next bucket in last
// cell of bucket.
const thresholdMinusOne = threshold - 1

// bucketStartSize is the initial size for new buckets.
const bucketStartSize = 16

// bucketGrowthFactor is the bucket growth factor.
const bucketGrowthFactor = 8

// bucket is a simplistic vector of strings which is grown according
// to the bucketStartSize and bucketGrowthFactor values.
type bucket []interface{}

// realloc resizes the bucket to the desired length.
func (b *bucket) realloc(length, capacity int) (n bucket) {
	if capacity < bucketStartSize {
		capacity = bucketStartSize
	}
	if capacity < length {
		capacity = length
	}
	n = make(bucket, length, capacity)
	copy(n, *b)
	*b = n
	return
}

// burstNode is the primary structure of the burst trie. Each entry in
// the trie node can be either nil, a bucket of strings, or a pointer to
// another trie node. A bucket is a slice of strings. The null bucket
// is a collection of those strings that are completely consumed by the
// trie structure. The last entry in a null bucket is a pointer to
// another null bucket, effectively daisy chaining the buckets.
type burstNode struct {
	// the last null bucket in the chain, starting from the first
	// pointer in elements[0]
	nulltail bucket
	// index into last null bucket (we manage the size ourselves to
	// avoid relocating the underlying array)
	nulltailidx int
	// counts is the number of items in each bucket, or -1 if
	// reference to trie node; for the null bucket it is the total
	// size of all null buckets in the chain.
	counts [alphabetSize]int
	// pointers to buckets or trie node
	elements [alphabetSize]interface{}
}

// add adds the given string into the appropriate bucket, using the
// character index into the trie. Presumably the character is from the
// string, but not necessarily so. The character may be the null
// character, in which case the string is added to the null bucket.
// Buckets are expanded as needed to accomodate the new string.
func (n *burstNode) add(c uint8, s string) {
	// are buckets already created?
	if n.counts[c] < 1 {
		// need to create bucket
		if c == nullterm {
			// allocate memory for the null bucket,
			// which is always sized at the maximum
			n.nulltail = make(bucket, threshold)
			n.nulltail[0] = s
			n.nulltailidx = 1
			n.elements[c] = n.nulltail
			n.counts[c]++
		} else {
			// allocate memory for the bucket
			b := make(bucket, 1, bucketStartSize)
			b[0] = s
			n.elements[c] = b
			n.counts[c]++
		}
	} else {
		// bucket already created
		if c == nullterm {
			// check if the bucket has reached the threshold
			if n.counts[c]%thresholdMinusOne == 0 {
				// grow the null bucket by daisy chaining a new slice
				tmp := make(bucket, threshold)
				n.nulltail[n.nulltailidx] = tmp
				n.nulltailidx = 0
				// point to the first cell in the new slice
				n.nulltail = tmp
			}
			// insert string in bucket and increment the item counter
			n.nulltail[n.nulltailidx] = s
			n.nulltailidx++
			n.counts[c]++
		} else {
			// insert string in bucket and increment the item counter
			b := n.elements[c].(bucket)
			b = append(b, s)
			n.counts[c]++
			// when bucket fills, increase its size up to the threshold
			l := len(b)
			if n.counts[c] == l && n.counts[c] < threshold {
				b = b.realloc(l, l*bucketGrowthFactor)
			}
			n.elements[c] = b
		}
	}
}

// get retrieves either a trie node or a bucket for the given character.
func (n *burstNode) get(c uint8) interface{} {
	return n.elements[c]
}

// set sets the trie node or bucket for a character.
func (n *burstNode) set(c uint8, o interface{}) {
	n.elements[c] = o
	if _, is_node := o.(*burstNode); is_node {
		// flag to indicate pointer to trie node and not bucket
		n.counts[c] = -1
	}
}

// size gets the number of strings stored for the given character.
func (n *burstNode) size(c uint8) int {
	return n.counts[c]
}

// burstInsert adds a set of strings into the burst trie structure, in
// preparation for in-order traversal (hence sorting).
func burstInsert(root *burstNode, strings []string) {
	for _, word := range strings {
		// start at root each time
		curr := root
		// locate trie node in which to insert string
		p := 0
		c := charAt(word, p)
		for curr.size(c) < 0 {
			curr = curr.get(c).(*burstNode)
			p++
			c = charAt(word, p)
		}

		curr.add(c, word)

		// This section is incredibly slow, and this is made
		// worse when the input consists of long, repeated
		// strings.

		// is bucket size above the threshold?
		for curr.size(c) >= threshold && c != nullterm {
			// advance depth of character
			p++
			// allocate memory for new trie node
			newt := new(burstNode)
			// burst...
			var cc uint8 = nullterm
			ptrs := curr.get(c).(bucket)
			size := curr.size(c)
			for j := 0; j < size; j++ {
				// access the next depth character
				str := ptrs[j].(string)
				cc = charAt(str, p)
				newt.add(cc, str)
			}
			// old pointer points to the new trie node
			curr.set(c, newt)
			// used to burst recursive, so point curr to new
			curr = newt
			// point to character used in previous string
			c = cc
		}
	}
}

// burstTraverse traverses the trie structure, ordering the strings in
// the array to conform to their lexicographically sorted order as
// determined by the trie structure.
func burstTraverse(node *burstNode, strings []string, pos, depth int) int {
	for c := 0; c < alphabetSize; c++ {
		idx := uint8(c)
		count := node.size(idx)
		if count < 0 {
			pos = burstTraverse(node.get(idx).(*burstNode), strings, pos, depth+1)
		} else if count > 0 {
			off := pos
			if c == 0 {
				// Visit all of the null buckets, which are daisy-chained
				// together with the last reference in each bucket pointing
				// to the next bucket in the chain.
				num_buckets := (count / thresholdMinusOne) + 1
				nullbucket := node.get(idx).(bucket)
				for k := 1; k <= num_buckets; k++ {
					var num_elements_in_bucket int
					if k == num_buckets {
						num_elements_in_bucket = count % thresholdMinusOne
					} else {
						num_elements_in_bucket = thresholdMinusOne
					}
					// copy the string tails to the sorted array
					j := 0
					for j < num_elements_in_bucket {
						strings[off] = nullbucket[j].(string)
						off++
						j++
					}
					if nullbucket[j] != nil {
						nullbucket = nullbucket[j].(bucket)
					}
				}
			} else {
				// copy to final destination
				bucket := node.get(idx).(bucket)
				dst := strings[off : off+count]
				for i, v := range bucket {
					// convert types while copying
					dst[i] = v.(string)
				}
				// sort the tail string bucket
				if count > 1 {
					MultikeyQuickSortDepth(dst, depth+1)
				}
			}
			pos += count
		}
	}
	return pos
}

// BurstSort sorts the given set of strings using the original
// (P-)burstsort algorithm.
func BurstSort(strings []string) {
	if strings != nil && len(strings) > 1 {
		root := new(burstNode)
		burstInsert(root, strings)
		burstTraverse(root, strings, 0, 0)
	}
}

//
// TODO: implement concurrent version using go routines
//
