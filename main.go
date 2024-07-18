package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		fifo_ratio  float64
		lru_ratio   float64
		sieve_ratio float64
	)
	for i := 0; i < 5; i++ {
		a, b, c := MissRatio(10000)
		fifo_ratio += a
		lru_ratio += b
		sieve_ratio += c
	}
	fmt.Println(float64(fifo_ratio) / 5)
	fmt.Println(float64(lru_ratio) / 5)
	fmt.Println(float64(sieve_ratio) / 5)

	var (
		fifo_time  time.Duration
		lru_time   time.Duration
		sieve_time time.Duration
	)
	for i := 0; i < 5; i++ {
		a, b, c := Efficiency(10000)
		fifo_time += a
		lru_time += b
		sieve_time += c
	}

	fmt.Println(fifo_time / 5)
	fmt.Println(lru_time / 5)
	fmt.Println(sieve_time / 5)

}
