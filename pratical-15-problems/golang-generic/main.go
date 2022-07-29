package main

import (
	"fmt"
)

type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func newGenericFunc[age Number](myAge age) {
	fmt.Println(myAge)
	val := int(myAge) + 1
	fmt.Println(val)
}
func BubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				fmt.Println(input)
				swapped = true
			}
		}
	}
	return input
}
func BubbleSortRecursion[N Number](input []N, size int) []N {
	if size == 1 {
		return input
	}
	var i = 0
	for i < size-1 {
		if input[i] > input[i+1] {
			input[i], input[i+1] = input[i+1], input[i]
		}
		i++
	}
	BubbleSortRecursion(input, size-1)
	return input
}
func InsertionSort[N Number](input []N) []N {
	for i := 1; i < len(input); i++ {
		key := input[i]
		j := i - 1
		for j >= 0 && input[j] > key {
			input[j+1] = input[j]
			fmt.Println(input)
			j = j - 1
		}
		input[j+1] = key
	}
	return input
}
func SelectionSort[N Number](input []N) []N {
	for i := 0; i < len(input)-1; i++ {
		min := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		if min != i {
			input[i], input[min] = input[min], input[i]
		}
		fmt.Println(input)
	}
	return input
}
func MergeSort[N Number](input []N) []N {
	if len(input) <= 1 {
		return input
	}
	mid := len(input) / 2
	left := MergeSort(input[:mid])
	right := MergeSort(input[mid:])
	return Merge(left, right)
}
func Merge[N Number](left, right []N) []N {
	var result []N
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func main() {
	fmt.Println("Go Generic Tutorial")
	// var testAge int64 = 28
	// var testAge2 float32 = 28.7

	// newGenericFunc(testAge)
	// newGenericFunc(testAge2)
	list := []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// listfloat := []float64{4.4, 5.5, 6.6, 7.7, 1.1, 2.2, 3.3, 8.8, 9.9, 10.10}
	// for i := 0; i < 10000; i++ {
	// 	listfloat = append(listfloat, math.Floor((rand.Float64()*1000)+1)/100)
	// }
	fmt.Println(BubbleSort(list))
	// fmt.Println(BubbleSort(listfloat))
	// fmt.Println(BubbleSortRecursion(listfloat, len(listfloat)))
	// fmt.Println(list)
	// fmt.Println(InsertionSort(list))
	// fmt.Println(SelectionSort(list))
	// fmt.Println(MergeSort(list))
}
