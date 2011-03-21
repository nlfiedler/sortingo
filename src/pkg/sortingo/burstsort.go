//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

// alphabetSize is the number of character supported for the trie used
// in sorting (strings are treated as arrays of uint8 values).
const alphabetSize = 256
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
func (b *bucket) realloc(length, capacity int) (n []interface{}) {
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
// trie structure. The last entry in a null bucket may be a pointer to
// another null bucket, effectively daisy chaining the buckets.
type burstNode struct {
	// the last null bucket in the chain, starting from the
	// first pointer in elements[0]
	nulltail bucket
        // counts is the number of items in each bucket,
	// or -1 if reference to trie node; for the null
	// bucket it is the total size of all null buckets
	// in the chain.
        counts [alphabetSize]int
        // pointers to buckets or trie node
        elements [alphabetSize]interface{}
}

// append adds the given string into the appropriate bucket, using the
// character index into the trie. Presumably the character is from the
// string, but not necessarily so. The character may be the null
// character, in which case the string is added to the null bucket.
// Buckets are expanded as needed to accomodate the new string.
func (n *burstNode) append(c uint8, s string) {
        // are buckets already created?
        if n.counts[c] < 1 {
                // need to create bucket
                if c == nullterm {
			// allocate memory for the null bucket,
			// which is always sized at the maximum
			n.nulltail = make(bucket, 1, threshold)
			n.nulltail[0] = s
			n.elements[c] = n.nulltail
			n.counts[c]++
                } else {
			// allocate memory for the bucket
			cs := make(bucket, 1, bucketStartSize)
			cs[0] = s
			n.elements[c] = cs
			n.counts[c]++
                }
        } else {
                // bucket already created
                if (c == nullterm) {
			// check if the bucket has reached the threshold
			if (n.counts[c] % thresholdMinusOne == 0) {
				// grow the null bucket by daisy chaining a new slice
				tmp := make(bucket, 0, threshold)
				n.nulltail = append(n.nulltail, tmp)
				// point to the first cell in the new slice
				n.nulltail = tmp
			}
			// insert string in bucket and increment the item counter
			n.nulltail = append(n.nulltail, s)
			n.counts[c]++
                } else {
			cs := n.elements[c].(bucket)
			// when bucket fills, increase its size up to the threshold
			l := len(cs)
			if (n.counts[c] < threshold && n.counts[c] == l) {
				cs = cs.realloc(l, l * bucketGrowthFactor)
			}
			// insert string in bucket and increment the item counter
			n.elements[c] = append(cs, s)
			n.counts[c]++
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
	if _, is_node := o.(burstNode); is_node {
                // flag to indicate pointer to trie node and not bucket
                n.counts[c] = -1
        }
}

// size gets the number of strings stored for the given character.
func (n *burstNode) size(c uint8) int {
        return n.counts[c]
}

//
// TODO: implement insert, traverse, and sort
//

//
// TODO: implement concurrent version using go routines
//
