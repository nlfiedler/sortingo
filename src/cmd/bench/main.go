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
	"sortingo"
	"strings"
	"time"
)

// runCount is the number of times each sort is measured to compute an average.
const runCount = 5

// sorterNames are the name of the sort algorithms in the desired run order.
var sorterNames = []string{"Merge"}
// sorters maps sort algorithm names to implementing functions.
var sorters = make(map[string]func([]string))
// sortSizes are the different sizes of data used in testing, in the desired run order.
var sortSizes = []int{330000, 1000000, 3000000}
// dataSetNames are the names of the data sets, in the desired run order.
var dataSetNames = []string{"Repeat", "RepeatCycle", "Random", "PseudoWords", "SmallAlphabet", "Genome"}
// dataGenerators maps data set names to data generator functions.
var dataGenerators = make(map[string]func(size int) []string)

// init sets up the benchmark data structures.
func init() {
	sorters["Merge"] = sortingo.MergeSort

	dataGenerators["Repeat"] = generateRepeated
	dataGenerators["RepeatCycle"] = generateRepeatedCycle
	dataGenerators["Random"] = generateRandom
	dataGenerators["PseudoWords"] = generateUniqueWords
	dataGenerators["SmallAlphabet"] = generateSmallAlpha
	dataGenerators["Genome"] = generateGenome
}

// generateRepeated generates the repeated strings test data.
func generateRepeated(size int) []string {
	repeatedStrings := make([]string, size)
	a100 := strings.Repeat("A", 100)
	for idx := range repeatedStrings {
		repeatedStrings[idx] = a100
	}
	return repeatedStrings
}

// generateRepeatedCycle generates a repeating cycle of strings.
func generateRepeatedCycle(size int) []string {
	a100 := strings.Repeat("A", 100)
	strs := make([]string, len(a100))
	for i := range strs {
		strs[i] = a100[0 : i+1]
	}
	repeatedCycleStrings := make([]string, size)
	c := 0
	for i := range repeatedCycleStrings {
		repeatedCycleStrings[i] = strs[c]
		if c++; c >= len(strs) {
			c = 0
		}
	}
	return repeatedCycleStrings
}

// generateRandom generates a set of random strings, each of length 100.
func generateRandom(size int) []string {
	randomStrings := make([]string, size)
	for i := range randomStrings {
		bb := bytes.NewBuffer(make([]byte, 0, 100))
		for j := 0; j < 100; j++ {
			d := rand.Int31n(95)
			bb.WriteRune(' ' + d)
		}
		randomStrings[i] = string(bb.Bytes())
	}
	return randomStrings
}

// generateUniqueWords generates a set of unique pseudo words.
func generateUniqueWords(size int) []string {
	uniqueWords := make([]string, size)
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
				d := rand.Int31n(26)
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
	return uniqueWords
}

// generateSmallAlpha generates a set of random strings, each of length
// 100, consisting of a small alphabet of characters.
func generateSmallAlpha(size int) []string {
	smallAlphaStrings := make([]string, size)
	for i := range smallAlphaStrings {
		l := 1 + rand.Intn(100)
		bb := bytes.NewBuffer(make([]byte, 0, l))
		for j := 0; j < l; j++ {
			d := rand.Int31n(9)
			bb.WriteRune('a' + d)
		}
		smallAlphaStrings[i] = string(bb.Bytes())
	}
	return smallAlphaStrings
}

// Generate a set of random "genome" strings, each of length 9,
// consisting of the letters a, c, g, t.
func generateGenome(size int) []string {
	genomeStrings := make([]string, size)
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
	return genomeStrings
}

// usage displays command line usage information.
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

// main runs the benchmarks on the "faster" sorting algorithms using
// large data sets.
func main() {
	var help = flag.Bool("help", false, "show usage information")
	var list = flag.Bool("list", false, "list supported data sets and algorithms")
	var data = flag.String("data", "", "regex to select data sets to sort")
	var algo = flag.String("sort", "", "regex to select sorting algorithms")
	// TODO: add a 'size' flag to select the data sizes (small, medium, large)
	// TODO: add a 'file' flag to take a file to be sorted instead of generated data
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
		for _, size := range sortSizes {
			fmt.Printf("\t%d...\n", size)
			dataSet := dataGenerators[dataSetName](size)
			input := inputSets[size]
			for _, sorterName := range sorterNames {
				fmt.Printf("\t\t%-10s:\t", sorterName)
				sorter := sorters[sorterName]
				times := new([runCount]int64)
				for run := 0; run < runCount; run++ {
					copy(input, dataSet)
					t1 := time.Now()
					sorter(input)
					t2 := time.Now()
					times[run] = (t2.Sub(t1).Nanoseconds()) / 1000000
				}

				// Find the average of the run times. The run times
				// should never be more than a couple of minutes,
				// so these calculations will never overflow.
				var total int64 = 0
				var highest int64 = 0
				var lowest int64 = 2147483648
				for run := 0; run < runCount; run++ {
					total += times[run]
					if times[run] > highest {
						highest = times[run]
					}
					if times[run] < lowest {
						lowest = times[run]
					}
				}
				average := total / runCount
				fmt.Printf("%4d %4d %4d (low/avg/high) ms\n", lowest, average, highest)
			}
		}
	}
}
