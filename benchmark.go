package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/djonathanb/ppgcomp-algorithms-ds/sort"
)

func createArray(size int) []int {
	v := make([]int, size)
	for i := range v {
		v[i] = rand.Intn(size * 100)
	}
	return v
}

func cpArray(v []int) []int {
	c := make([]int, len(v))
	for i, n := range v {
		c[i] = n
	}
	return c
}

func verify(v []int) error {
	d := v[0]
	for i := 1; i < len(v); i++ {
		if d > v[i] {
			return errors.New("unsorted")
		}
		d = v[i]
	}
	return nil
}

func test(original []int, fn func([]int), name string) {
	cp := cpArray(original)
	start := time.Now()
	fn(cp)
	fmt.Printf("%s took %v\n", name, time.Since(start))

	err := verify(cp)
	if err != nil {
		fmt.Printf("%v\n", cp)
		fmt.Println(err)
	}
}

func main() {
	original := createArray(50000)

	test(original, sort.BubbleSort, "bubble")
	test(original, sort.SelectionSort, "selection")
	test(original, sort.InsertionSort, "insertion")
	test(original, func(v []int) { sort.HeapSort(v) }, "heap")
	test(original, sort.BucketSort, "bucket")
	test(original, sort.ShellSort, "shell")
	test(original, sort.MergeSort, "merge")
	test(original, sort.QuickSort, "quick")
}
