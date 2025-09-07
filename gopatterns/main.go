package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()

	wg := sync.WaitGroup{}
	go func() {
		panic("aaaa!")
	}()

	wg.Wait()
	c := make(chan struct {
		int
		error 
	})
	c <- struct {
		int
		error
	}{1, nil}

	/*c1 := producer(true)
	c2 := producer(false)
	res := merge(c1, c2)
	for v := range res {
		fmt.Println(v)
	}*/

	/*err := wpool(context.Background(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100)
	if err != nil {
		log.Fatal(err)
	}*/
}

func merge(chans ...<-chan int) <-chan int {
	res := make(chan int)
	wg := sync.WaitGroup{}
	for _, c := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				res <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func producer(negative bool) <-chan int {
	res := make(chan int)
	go func() {
		for i := range 1000 {
			if negative {
				res <- -i
			} else {
				res <- i
			}

			time.Sleep(10000)
		}
		close(res)
	}()

	return res
}

func wpool(ctx context.Context, work []int, n int) error {
	c := make(chan int)
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	wg.Add(1)
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()
	var wpoolErr error
	go func() {
		defer wg.Done()
		for _, w := range work {
			select {
			case c <- w:
			case <-ctx.Done():
				return
			}

		}
	}()

	for range n {
		go func() {
			for v := range c {
				err := doWork(ctx, v)
				if err != nil {
					m.Lock()
					wpoolErr = errors.Join(wpoolErr, err)
					cancelCtx()
					m.Unlock()
				}
			}
		}()
	}

	wg.Wait()
	close(c)
	return wpoolErr
}

func doWork(ctx context.Context, work int) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	if work == 7 {
		return fmt.Errorf("some error")
	}

	fmt.Println(work)
	return nil
}
