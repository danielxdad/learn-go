package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutimes() {
	// go say("world")
	// say("hello")

	// d := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// d := rand.Perm(5)
	d := []int{4, 2, 3, 1, 0}
	c := make(chan []int)
	go mergesort(d, c)
	fmt.Printf("%v\n%v\n", d, <-c)
}

func mergesort(data []int, c chan []int) {
	if len(data) < 2 {
		c <- data
		return
	}

	lc := make(chan []int)
	rc := make(chan []int)

	go mergesort(data[:len(data)/2], lc)
	go mergesort(data[len(data)/2:], rc)

	left := <-lc
	right := <-rc

	result := make([]int, 0, len(data))

	if left[len(left)-1] < right[0] {
		result = append(result, left...)
		result = append(result, right...)
	} else if left[0] > right[len(right)-1] {
		result = append(result, right...)
		result = append(result, left...)
	} else {
		n := min(len(left), len(right))

		for i := range n {
			x, y := left[i], right[i]
			result = append(result, min(x, y), max(x, y))
		}

		if len(left) > n {
			result = append(result, left[n:]...)
		}

		if len(right) > n {
			result = append(result, right[n:]...)
		}
	}

	c <- result
}
