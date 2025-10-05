package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chNum := 10
	chStart := make(chan int, chNum)
	chResult := make(chan int, chNum)
	go createRandSlice(chNum, chStart)
	go squareNum(chStart, chResult)
	var resultSlice []int
	for num := range chResult {
		resultSlice = append(resultSlice, num)
	}

	fmt.Println("Конечный слайс: ", resultSlice)
}

func createRandSlice(chNum int, ch chan int) {
	var slice []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < chNum; i++ {
		n := r.Intn(101)
		slice = append(slice, n)
	}
	for _, num := range slice {
		ch <- num
	}
	close(ch)
}

func squareNum(ch chan int, chResult chan int) {
	for num := range ch {
		chResult <- num * num
	}
	close(chResult)
}
