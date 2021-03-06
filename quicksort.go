package main

import (
	"fmt"
)
// Go Playground 
// https://play.golang.org/p/HhFzhdf8Kc

func partition(arr []int, start int, end int) int {
	pivot := arr[end] // Last element is the pivot
	pIndex := start   // Let the pindex be the first
	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex = pIndex + 1
		}
	}
	arr[pIndex], arr[end] = arr[end], arr[pIndex]
	return pIndex
}

func qsort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	pIndex := partition(arr, start, end)
	fmt.Println("array ", arr, "index ", pIndex)
	qsort(arr, start, pIndex-1)
	qsort(arr, pIndex+1, end)

}


//array  [3 1 2 4 7 6 5] index  3
//array  [1 2 3 4 7 6 5] index  1
//array  [1 2 3 4 5 6 7] index  4
//array  [1 2 3 4 5 6 7] index  6
//Sorted Array:  [1 2 3 4 5 6 7]

func main() {
	arr := []int{3, 5, 1, 2, 7, 6, 4}
	qsort(arr, 0, len(arr)-1)
	fmt.Println("Sorted Array: ", arr)
}
