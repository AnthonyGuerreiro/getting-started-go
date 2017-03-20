package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"./fibonacci"
	"./prime"
	"./coin"
	"./sort"
	"sort"
)

func main() {
	customBSort()
}

func customBSort() {

	var size = readInt()
	var elements = readNInts(size)

	var start = time.Now()
	custom_sort.BSort(elements)
	var elapsed = time.Since(start)

	log.Println("Elapsed", elapsed)
	log.Println("Sorted", elements)
}


func customQSort() {

	var size = readInt()
	var elements = readNInts(size)

	var start = time.Now()
	custom_sort.Qsort(elements)
	var elapsed = time.Since(start)

	log.Println("Elapsed", elapsed)
	log.Println("Sorted", elements)
}

func flipCoinNTimes() {

	var flips = readInt()

	var start = time.Now()
	var heads = coin.FlipNTimes(flips)
	var elapsed = time.Since(start)

	log.Println("Elapsed", elapsed)
	log.Println("Flipped coin", flips, "times. Obtained", heads, "heads and", flips-heads, "tails")
}

func nextPrime() {
	var number = readInt64()

	var start = time.Now()
	var result = prime.NextPrime(number)
	var elapsed = time.Since(start)

	log.Println("Elapsed", elapsed)
	log.Println("Next prime greater than or equal to", number, ":", result)
}

func factors() {
	var number = readInt64()

	var start = time.Now()
	var result = prime.Factors(number)
	var elapsed = time.Since(start)

	log.Println("Elapsed", elapsed)
	log.Println("Factors of", number, ":", result)
}

func rec_fyb() {

	var index = readInt()

	var start = time.Now()
	var result = fibonacci.Rec(index)
	var elapsed = time.Since(start)

	log.Println("Elapsed", elapsed)
	log.Println("Number at index", index, "in fibonacci seq:", result)
}

func dyn_fyb() {

	var index = readInt()

	var start = time.Now()
	var result = fibonacci.Dyn(index)
	var elapsed = time.Since(start)

	log.Println("Elapsed", elapsed)
	log.Println("Number at index", index, "in fibonacci seq:", result)
}

func readInt() int {
	var i int
	fmt.Println("Insert int")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

func readInt32() int32 {
	var i int32
	fmt.Println("Insert int")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

func readInt64() int64 {
	var i int64
	fmt.Println("Insert int64")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}

func readNInts(n int) []int {
	var elements []int;
	var i int
	for i = 0; i < n; i++ {
		elements = append(elements, readInt())
	}
	return elements
}
