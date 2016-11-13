package lib

import (
	"math"
)

func QuickSort (arr [] int) []int{
	return quickSort(arr, 0, len(arr) -1)
}

func quickSort (arr [] int, left, right int) []int {
	index := partition(arr, left, right);

	if left < index -1 {
		quickSort(arr, left, index - 1);
	}
	if index < right {
		quickSort(arr, index, right)
	}

	return arr
}

func partition (arr []int, left, right int) int{
	pivot := arr[int(math.Floor(float64((left + right) / 2)))]
	i := left
	j := right

	for i < j {
		for (arr[i] < pivot) {
			i++;
		}

		for (arr[j] > pivot) {
			j--;
		}

		if (i <= j) {
			arr = swap(arr, i, j);
			i++;
			j--;
		}
	}

	return i;
}

func swap (arr[] int, fromIndex, toIndex int) []int{
	fromIndexVal := arr[fromIndex]
	arr[fromIndex] = arr[toIndex]
	arr[toIndex] = fromIndexVal

	return arr
}