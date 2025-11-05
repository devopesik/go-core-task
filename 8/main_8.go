package main

import (
	"fmt"
	"sync/atomic"
)

type CustomWaitGroup struct {
	count int32         // Атомарный счётчик
	ch    chan struct{} // Канал для ожидания (действует как семафор с сигналом закрытия)
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		ch: make(chan struct{}),
	}
}

// Add увеличивает счётчик на delta.
func (wg *CustomWaitGroup) Add(delta int) {
	newCount := atomic.AddInt32(&wg.count, int32(delta))
	if newCount < 0 {
		panic("custom wait group: negative count")
	}
}

// Done уменьшает счётчик на 1 и закрывает канал, если счётчик достиг 0.
func (wg *CustomWaitGroup) Done() {
	if atomic.AddInt32(&wg.count, -1) == 0 {
		close(wg.ch)
	}
}

// Wait блокирует до тех пор, пока счётчик не станет 0.
// Если канал уже закрыт, не блокирует.
func (wg *CustomWaitGroup) Wait() {
	if atomic.LoadInt32(&wg.count) == 0 {
		return
	}
	<-wg.ch
}

func main() {
	wg := NewCustomWaitGroup()
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1 done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 done")
	}()

	wg.Wait()
	fmt.Println("All goroutines done")
}
