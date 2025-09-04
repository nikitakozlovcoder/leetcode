package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	workerPool2([]job{
		func() {
			fmt.Println(1)
			<-time.Tick(1 * time.Second)
		},
		func() {
			fmt.Println(3)
			<-time.Tick(3 * time.Second)
		},
		func() {
			fmt.Println(4)
			<-time.Tick(4 * time.Second)
		},
		func() {
			fmt.Println(5)
		},
		func() {
			fmt.Println(5)
		},
		func() {
			fmt.Println(5)
		},
		func() {
			fmt.Println(5)
		},
		func() {
			fmt.Println(5)
		},
		func() {
			fmt.Println(5)
		},
		func() {
			fmt.Println(5)
		},
		func() {
			fmt.Println(5)
		},
	}, context.Background())
}

type job = func()

func workerPool(jobs <-chan job) {
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				j()
			}

		}()
	}

	wg.Wait()
}

func workerPool2(jobs []job, ctx context.Context) {
	fmt.Println("workerPool2")
	jobsCh := make(chan job)
	go func() {
		for _, j := range jobs {
			fmt.Println("workerPool2 push")
			select {
			case jobsCh <- j:
			case <-ctx.Done():
				close(jobsCh)
				return
			}
		}

		close(jobsCh)
	}()

	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobsCh {
				fmt.Println("workerPool2 exec")
				j()
			}

		}()
	}

	wg.Wait()
}
