//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package main

import (
	"bytes"
	"fmt"
	"rand"
	"sortingo"
	"strings"
	"time"
)

// Largest data size used in the micro benchmarks.
const largeDataSize = 400

// sorterNames are the name of the sort algorithms in the desired run order.
var sorterNames = []string{"Binsert", "Comb", "Gnome", "Heap", "Insert", "MkQ", "Select", "Shell"}
// sorters maps sort algorithm names to implementing functions.
var sorters = make(map[string]func([]string))
// sortSizes are the different sizes of data used in testing, in the desired run order.
var sortSizes = []int{10, 20, 50, 100, 400}
// sortCounts is the number of times to test a sort for a given size.
var sortCounts = make(map[int]int)
// dataSetNames are the names of the data sets, in the desired run order.
var dataSetNames = []string{"Repeat", "RepeatCycle", "Random", "PseudoWords", "SmallAlphabet", "Genome"}
// dataSets contains the various data sets that are used in testing.
var dataSets = make(map[string][]string)

// init sets up the benchmark data structures.
func init() {
	sorters["Binsert"] = sortingo.BinaryInsertionSort
	sorters["Comb"] = sortingo.CombSort
	sorters["Gnome"] = sortingo.GnomeSort
	sorters["Heap"] = sortingo.HeapSort
	sorters["Insert"] = sortingo.InsertionSort
	sorters["MkQ"] = sortingo.MultikeyQuickSort
	sorters["Select"] = sortingo.SelectionSort
	sorters["Shell"] = sortingo.ShellSort
	sortCounts[10] = 500000
	sortCounts[20] = 250000
	sortCounts[50] = 100000
	sortCounts[100] = 25000
	sortCounts[400] = 5000

	// Generate the repeated strings test data.
	repeatedStrings := make([]string, largeDataSize)
	a100 := strings.Repeat("A", 100)
	for idx, _ := range repeatedStrings {
		repeatedStrings[idx] = a100
	}
	dataSets["Repeat"] = repeatedStrings

	// Generate a repeating cycle of strings.
	strs := make([]string, len(a100))
	for i, _ := range strs {
		strs[i] = a100[0:i + 1]
	}
	repeatedCycleStrings := make([]string, largeDataSize)
	c := 0
	for i, _ := range repeatedCycleStrings {
		repeatedCycleStrings[i] = strs[c]
		if c++; c >= len(strs) {
			c = 0
		}
	}
	dataSets["RepeatCycle"] = repeatedCycleStrings

	// Generate a set of random strings, each of length 100.
	randomStrings := make([]string, largeDataSize)
	for i, _ := range randomStrings {
		bb := bytes.NewBuffer(make([]byte, 0, 100))
		for j := 0; j < 100; j++ {
			d := rand.Intn(95)
			bb.WriteRune(' ' + d)
		}
		randomStrings[i] = string(bb.Bytes())
	}
	dataSets["Random"] = randomStrings

	// Generate a set of unique pseudo words.
	uniqueWords := make([]string, largeDataSize)
	wordExists := make(map[string]bool)
	for i, _ := range uniqueWords {
		var s string
		// Loop until a unique random word is generated.
		for {
			// Each word is from 1 to 28 characters long.
			l := 1 + rand.Intn(27)
			bb := bytes.NewBuffer(make([]byte, 0, l))
			// Each word consists only of the lowercase letters.
			for j := 0; j < l; j++ {
				d := rand.Intn(26)
				bb.WriteRune('a' + d)
			}
			s = string(bb.Bytes())
			if !wordExists[s] {
				break
			}
		}
		uniqueWords[i] = s
		wordExists[s] = true
	}
	dataSets["PseudoWords"] = uniqueWords

	// Generate a set of random strings, each of length 100,
	// consisting of a small alphabet of characters.
	smallAlphaStrings := make([]string, largeDataSize)
	for i, _ := range smallAlphaStrings {
		l := 1 + rand.Intn(100)
		bb := bytes.NewBuffer(make([]byte, 0, l))
		for j := 0; j < l; j++ {
			d := rand.Intn(9)
			bb.WriteRune('a' + d)
		}
		smallAlphaStrings[i] = string(bb.Bytes())
	}
	dataSets["SmallAlphabet"] = smallAlphaStrings

	// Generate a set of random "genome" strings, each of length 9,
	// consisting of the letters a, c, g, t.
	genomeStrings := make([]string, largeDataSize)
	for i, _ := range genomeStrings {
		bb := bytes.NewBuffer(make([]byte, 0, 9))
		for j := 0; j < 9; j++ {
			d := rand.Intn(4)
			switch d {
			case 0:
				bb.WriteRune('a')
			case 1:
				bb.WriteRune('c')
			case 2:
				bb.WriteRune('g')
			case 3:
				bb.WriteRune('t')
			}
		}
		genomeStrings[i] = string(bb.Bytes())
	}
	dataSets["Genome"] = genomeStrings
}

// main runs the micro benchmarks on the slower sorting algorithms.
func main() {
        // For each type of data set...
	// and each data set size...
	// and each sort implementation...
	// run the sort many times and calculate an average run time.
        for _, dataSetName := range dataSetNames {
		fmt.Printf("%s...\n", dataSetName)
		dataSet := dataSets[dataSetName]
		for _, size := range sortSizes {
			input := make([]string, size)
			runCount := sortCounts[size]
			fmt.Printf("\t%d...\n", size)
			for _, sorterName := range sorterNames {
				sorter := sorters[sorterName]
				fmt.Printf("\t\t%-10s:\t", sorterName)
				start := time.Nanoseconds()
				for run := 0; run < runCount; run++ {
					copy(input, dataSet)
					sorter(input)
				}
				runtime := time.Nanoseconds() - start
				fmt.Printf("%d ms\n", runtime / 1e6)
			}
		}
        }
}
