//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package sortingo

import (
	"testing"
)

// TestBurstAppend creates several buckets and ensures they have the
// expected size and contents.
func TestBurstAppend(t *testing.T) {
	var node burstNode
	var buck bucket
	letters := [...]uint8{0, 'a', 'z', 'e', 'g', 'h', 'i', 'm', 'p'}
	words := [...]string{"nullary", "albatross", "zebra", "elephant",
		"giraffe", "hippopotamus", "ibex", "monkey", "platapus"}
	// insert strings
	for i, c := range letters {
		node.append(c, words[i])
	}
	// retrieve strings and verify
	for i, c := range letters {
		elem := node.get(c)
		if _, is_node := elem.(burstNode); is_node {
			t.Error("expected bucket, got node")
		}
		buck = elem.(bucket)
		if len(buck) != 1 {
			t.Error("expected bucket to have one entry")
		}
		if buck[0] != words[i] {
			t.Error("wrong string in bucket")
		}
	}
}

// TestBurstMultiple inserts several strings into a single bucket and
// tests that they can be retrieved in the expected order.
func TestBurstMultiple(t *testing.T) {
	var node burstNode
	var buck bucket
	words := [...]string{"fish", "food", "freeze"}
	for _, w := range words {
		node.append('f', w)
	}
	elem := node.get('f')
	if _, is_node := elem.(burstNode); is_node {
		t.Error("expected bucket, got node")
	}
	buck = elem.(bucket)
	if len(buck) != len(words) {
		t.Error("wrong size for bucket")
	}
	for i, w := range words {
		if buck[i] != w {
			t.Error("wrong string in bucket")
		}
	}
}

// TestBurstNested creates a parent and child node and verifies that
// they have the correct linkage.
func TestBurstNested(t *testing.T) {
	var parent burstNode
	var child burstNode
	// set one entry to be another node
	child.append('l', "oliphaunt")
	parent.set('o', child)
	elem := parent.get('o')
	if _, is_node := elem.(burstNode); !is_node {
		t.Error("expected node from set/get combo")
	}
	node := elem.(burstNode)
	buck := node.get('l').(bucket)
	if buck[0] != "oliphaunt" {
		t.Error("expected string not found")
	}
	// verify sizes of buckets and nodes (i.e. a node entry should be size -1)
	if parent.size('o') != -1 {
		t.Error("expected size -1 for node element")
	}
	if child.size('a') != 0 {
		t.Error("expected unused entry to be size 0")
	}
	if child.size('l') != 1 {
		t.Error("expected 'l' entry to be size 1")
	}
}

func TestBurstExpansion(t *testing.T) {
	// TODO: add many elements to null bucket and other buckets to test slice expansion
	//       (generate a set of random strings, insert them and verify they are all accessible)
}
