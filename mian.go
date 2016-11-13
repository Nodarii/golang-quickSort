package main

import (
	"fmt"
	"./lib"
	"math/rand"
	"time"
)


func main(){
	randomArr := getRandomArray(100, 1000);
	fmt.Println("Random array is", randomArr);
	fmt.Println("Sorted array is", lib.QuickSort(randomArr))
}

func getRandomArray(size, maxVal int) []int {
	arr := make([]int, size)

	for i := 0; size > i; i++ {
		arr[i] = randInt(0, maxVal)
	}

	return arr;
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())

	return min + rand.Intn(max-min)
}
