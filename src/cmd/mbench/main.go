//
// Copyright 2011 Nathan Fiedler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"sortingo"
	"strings"
	"testing"
)

// Largest data size used in the micro benchmarks.
const largeDataSize = 400

// sorterNames are the name of the sort algorithms in the desired run order.
var sorterNames = []string{"Binsert", "Comb", "2PivotQ", "Gnome", "Heap", "HybridComb", "Insert", "MkQ", "Quick", "Select", "Shell"}
// sorters maps sort algorithm names to implementing functions.
var sorters = make(map[string]func([]string))
// sortSizes are the different sizes of data used in testing, in the desired run order.
var sortSizes = []int{10, 20, 50, 100, 400}
// dataSetNames are the names of the data sets, in the desired run order.
var dataSetNames = []string{"Repeat", "RepeatCycle", "Random", "PseudoWords", "SmallAlphabet", "Genome"}
// dataSets contains the various data sets that are used in testing.
var dataSets = make(map[string][]string)

// init sets up the benchmark data structures.
func init() {
	sorters["Binsert"] = sortingo.BinaryInsertionSort
	sorters["Comb"] = sortingo.CombSort
	sorters["2PivotQ"] = sortingo.DualPivotQuickSort
	sorters["Gnome"] = sortingo.GnomeSort
	sorters["Heap"] = sortingo.HeapSort
	sorters["HybridComb"] = sortingo.HybridCombSort
	sorters["Insert"] = sortingo.InsertionSort
	sorters["MkQ"] = sortingo.MultikeyQuickSort
	sorters["Quick"] = sort.Strings
	sorters["Select"] = sortingo.SelectionSort
	sorters["Shell"] = sortingo.ShellSort

	// Generate the repeated strings test data.
	repeatedStrings := make([]string, largeDataSize)
	a100 := strings.Repeat("A", 100)
	for idx := range repeatedStrings {
		repeatedStrings[idx] = a100
	}
	dataSets["Repeat"] = repeatedStrings

	// Generate a repeating cycle of strings.
	strs := make([]string, len(a100))
	for i := range strs {
		strs[i] = a100[0 : i+1]
	}
	repeatedCycleStrings := make([]string, largeDataSize)
	c := 0
	for i := range repeatedCycleStrings {
		repeatedCycleStrings[i] = strs[c]
		if c++; c >= len(strs) {
			c = 0
		}
	}
	dataSets["RepeatCycle"] = repeatedCycleStrings

	// Generate a set of random strings, each of length 100.
	randomStrings := make([]string, largeDataSize)
	for i := range randomStrings {
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
	for i := range uniqueWords {
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
	for i := range smallAlphaStrings {
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
	for i := range genomeStrings {
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

func usage() {
	fmt.Println("Usage: mbench [options]")
	fmt.Println("\t--data <regex>")
	fmt.Println("\t\tSelect the data set whose name matches the regular expression.")
	fmt.Println("\t\tFor example, '--data random' would use only the random data set.")
	fmt.Println("\t--help")
	fmt.Println("\t\tDisplay this usage information.")
	fmt.Println("\t--list")
	fmt.Println("\t\tDisplay a list of the supported data sets and sorting algorithms.")
	fmt.Println("\t--sort <regex>")
	fmt.Println("\t\tSelect the sort algorithms whose name matches the regular")
	fmt.Println("\t\texpression. For example, '--sort (comb|insert)' would run")
	fmt.Println("\t\tboth versions of the insertion and comb sort algorithms.")
}

// main runs the micro benchmarks on the "slower" sorting algorithms
// using small data sets (under 500 elements).
func main() {
	var help = flag.Bool("help", false, "show usage information")
	var list = flag.Bool("list", false, "list supported data sets and algorithms")
	var data = flag.String("data", "", "regex to select data sets to sort")
	var algo = flag.String("sort", "", "regex to select sorting algorithms")
	flag.Parse()

	if *help {
		usage()
		os.Exit(0)
	}

	if *list {
		fmt.Println("Data sets")
		for _, dataSetName := range dataSetNames {
			fmt.Printf("\t%s\n", dataSetName)
		}
		fmt.Println("Sorting algorithms")
		for _, sorterName := range sorterNames {
			fmt.Printf("\t%s\n", sorterName)
		}
		os.Exit(0)
	}

	if *data != "" {
		re, err := regexp.Compile(strings.ToLower(*data))
		if err != nil {
			fmt.Printf("%s in '%s' of --data flag\n", err, *data)
			os.Exit(1)
		}
		newData := make([]string, 0, len(dataSetNames))
		for _, dataSetName := range dataSetNames {
			if idx := re.FindStringIndex(strings.ToLower(dataSetName)); idx != nil {
				newData = append(newData, dataSetName)
			}
		}
		dataSetNames = newData
	}

	if *algo != "" {
		re, err := regexp.Compile(strings.ToLower(*algo))
		if err != nil {
			fmt.Printf("%s in '%s' of --sort flag\n", err, *algo)
			os.Exit(1)
		}
		newSort := make([]string, 0, len(sorterNames))
		for _, sorterName := range sorterNames {
			if idx := re.FindStringIndex(strings.ToLower(sorterName)); idx != nil {
				newSort = append(newSort, sorterName)
			}
		}
		sorterNames = newSort
	}

	// Avoid recreating the input arrays over and over again.
	inputSets := make(map[int][]string)
	for _, size := range sortSizes {
		inputSets[size] = make([]string, size)
	}

	// For each type of data set...
	// and each data set size...
	// and each sort implementation...
	// run the sort via the testing package benchmark facility.
	for _, dataSetName := range dataSetNames {
		fmt.Printf("%s...\n", dataSetName)
		dataSet := dataSets[dataSetName]
		for _, size := range sortSizes {
			input := inputSets[size]
			fmt.Printf("\t%d...\n", size)
			for _, sorterName := range sorterNames {
				sorter := sorters[sorterName]
				fmt.Printf("\t\t%-10s:\t", sorterName)
				harness := func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						copy(input, dataSet)
						b.StartTimer()
						sorter(input)
					}
				}
				result := testing.Benchmark(harness)
				fmt.Println(result)
			}
		}
	}
}
