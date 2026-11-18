package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Employee struct{}

type Item1 struct{}

func (*Item1) Process() {
	<-time.After(time.Second)
}

type Item2 struct{}

func (*Item2) Process() {
	<-time.After(time.Second * 2)
}

type Item3 struct{}

func (*Item3) Process() {
	<-time.After(time.Second * 3)
}

type Item interface {
	// Process 這是一個耗時操作
	Process()
}

func main() {
	workerNum := 5
	workerStats := make([]int, workerNum)
	itemsCh := make(chan Item)
	var workers sync.WaitGroup
	startedAt := time.Now()

	workers.Add(workerNum)
	for workerIndex := range workerNum {
		go func() {
			defer workers.Done()
			for item := range itemsCh {
				item.Process()
				workerStats[workerIndex]++
			}
		}()
	}

	go func() {
		defer close(itemsCh)

		items := make([]Item, 0, 30)

		for range 10 {
			items = append(items, &Item1{})
			items = append(items, &Item2{})
			items = append(items, &Item3{})
		}

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		rng.Shuffle(len(items), func(i, j int) {
			items[i], items[j] = items[j], items[i]
		})

		for _, item := range items {
			itemsCh <- item
		}
	}()

	fmt.Printf("Start processing at %s\n", startedAt)
	workers.Wait()
	fmt.Printf("Total processing time: %s\n", time.Since(startedAt))
	for workerIndex, count := range workerStats {
		fmt.Printf("Worker %d processed %d items\n", workerIndex+1, count)
	}
}
