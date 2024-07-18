package main

import (
	"math/rand"
	"sync"
	"time"
)



func MissRatio(lenth int) (float64, float64, float64) {
	publicGet, publicPut, fifo, lru, sieve := Init(lenth)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		getRequest(publicGet, lru.Get, sieve.Get, fifo.Get)
		wg.Add(-1)
	}()

	go func() {
		putRequest(publicPut, lru.Put, sieve.Put, fifo.Put)
		wg.Add(-1)
	}()
	wg.Wait()
	return float64(fifo.miss) / float64(lenth), float64(lru.miss) / float64(lenth), float64(sieve.miss) / float64(lenth)
}

func Efficiency(lenth int) (time.Duration, time.Duration, time.Duration) {
	publicGet, publicPut, fifo, lru, sieve := Init(lenth)
	Timer := func(l List) time.Duration {
		start := time.Now()
		var wg sync.WaitGroup
		func() {
			wg.Add(2)
			go func() {
				getRequest(publicGet, l.Get)
				wg.Add(-1)
			}()
			go func() {
				putRequest(publicPut, l.Put)
				wg.Add(-1)
			}()
			wg.Wait()
		}()
		return time.Since(start)
	}
	return Timer(fifo), Timer(lru), Timer(sieve)

}

func Init(lenth int) ([]int, []int, *FIFO, *LRU, *Sieve) {
	rand.Seed(time.Now().UnixNano())
	publicQueue := make([]int, lenth)
	publicGet, publicPut := make([]int, lenth), make([]int, lenth)
	fifo := NewFIFOCache(lenth)
	lru := NewLRUCache(lenth)
	sieve := NewSieveCache(lenth)

	for i := 0; i < lenth; i++ {
		publicQueue[i] = i
		publicGet[i] = rand.Intn(lenth)
		publicPut[i] = rand.Intn(lenth*100)
	}
	rand.Shuffle(lenth, func(i, j int) {
		publicQueue[i], publicQueue[j] = publicQueue[j], publicQueue[i]
	})
	putRequest(publicQueue, lru.Put, sieve.Put, fifo.Put)
	return publicGet, publicPut, fifo, lru, sieve
}

func getRequest(publicGet []int, f ...func(key interface{}) interface{}) {
	for i := 0; i < len(publicGet); i++ {
		for _, fc := range f {
			fc(publicGet[i])
		}
	}
}

func putRequest(publicPut []int, f ...func(key, value interface{})) {
	var data int
	for i := 0; i < len(publicPut); i++ {
		data = publicPut[i]
		for _, fc := range f {
			fc(data, data)
		}
	}
}
