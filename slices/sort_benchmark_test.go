// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
)

// These benchmarks compare sorting a large slice of int with sort.Ints vs.
// slices.Sort
func makeRandomInts(n int) []int {
	ints := make([]int, n)
	fillRandomInts(ints)
	return ints
}

func fillRandomInts(ints []int) {
	rand.Seed(42)
	n := len(ints)
	for i := 0; i < len(ints); i++ {
		ints[i] = rand.Intn(n)
	}
}

func makeSortedInts(n int) []int {
	ints := make([]int, n)
	fillSortedInts(ints)
	return ints
}

func fillSortedInts(ints []int) {
	for i := 0; i < len(ints); i++ {
		ints[i] = i
	}
}

func makeReversedInts(n int) []int {
	ints := make([]int, n)
	fillReversedInts(ints)
	return ints
}

func fillReversedInts(ints []int) {
	n := len(ints)
	for i := 0; i < n; i++ {
		ints[i] = n - i
	}
}

func makeMixedInts(n int) []int {
	ints := make([]int, n)
	m := n / 3
	fillSortedInts(ints[:m])
	fillRandomInts(ints[m : n-m])
	fillReversedInts(ints[n-m:])
	return ints
}

const N = 100_000

func BenchmarkSortInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeRandomInts(N)
		b.StartTimer()
		sort.Ints(ints)
	}
}

func BenchmarkSlicesSortInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeRandomInts(N)
		b.StartTimer()
		Sort(ints)
	}
}

func BenchmarkSlicesSortInts_Sorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeSortedInts(N)
		b.StartTimer()
		Sort(ints)
	}
}

func BenchmarkSlicesSortInts_Reversed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeReversedInts(N)
		b.StartTimer()
		Sort(ints)
	}
}

func BenchmarkSlicesSortInts_Mixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ints := makeMixedInts(N)
		b.StartTimer()
		Sort(ints)
	}
}

// The following is a benchmark for sorting strings.

// makeRandomStrings generates n random strings with alphabetic runes of
// varying lenghts.
func makeRandomStrings(n int) []string {
	rand.Seed(42)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		var sb strings.Builder
		slen := 2 + rand.Intn(50)
		for j := 0; j < slen; j++ {
			sb.WriteRune(letters[rand.Intn(len(letters))])
		}
		ss[i] = sb.String()
	}
	return ss
}

func BenchmarkSortStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStrings(N)
		b.StartTimer()
		sort.Strings(ss)
	}
}

func BenchmarkSlicesSortStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStrings(N)
		b.StartTimer()
		Sort(ss)
	}
}

// These benchmarks compare sorting a slice of structs with sort.Sort vs.
// slices.SortFunc.
type myStruct struct {
	a, b, c, d string
	n          int
}

type myStructs []*myStruct

func (s myStructs) Len() int           { return len(s) }
func (s myStructs) Less(i, j int) bool { return s[i].n < s[j].n }
func (s myStructs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func makeRandomStructs(n int) myStructs {
	rand.Seed(42)
	structs := make([]*myStruct, n)
	for i := 0; i < n; i++ {
		structs[i] = &myStruct{n: rand.Intn(n)}
	}
	return structs
}

func BenchmarkSortStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStructs(N)
		b.StartTimer()
		sort.Sort(ss)
	}
}

func BenchmarkSortStableStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStructs(N)
		b.StartTimer()
		sort.Stable(ss)
	}
}

func BenchmarkSortFuncStructs(b *testing.B) {
	less := func(a, b *myStruct) bool { return a.n < b.n }
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStructs(N)
		b.StartTimer()
		SortFunc(ss, less)
	}
}

func BenchmarkSortStableFuncStructs(b *testing.B) {
	less := func(a, b *myStruct) bool { return a.n < b.n }
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStructs(N)
		b.StartTimer()
		SortStableFunc(ss, less)
	}
}
