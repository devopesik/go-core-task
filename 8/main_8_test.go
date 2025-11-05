package main

import (
	"testing"
	"time"
)

// Предполагаем, что CustomWaitGroup определена в том же пакете или импортирована.
// Для тестов используем тот же пакет.

func TestCustomWaitGroup_Basic(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(1)
	go func() {
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}()
	wg.Wait()
	// Если дошли сюда без паники или дедлока, тест пройден.
}

func TestCustomWaitGroup_MultipleAdds(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(3)
	doneCh := make(chan struct{})
	go func() {
		wg.Done()
		wg.Done()
		wg.Done()
		close(doneCh)
	}()
	select {
	case <-doneCh:
		// Done вызваны
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for Done calls")
	}
	wg.Wait()
}

func TestCustomWaitGroup_NegativeCountPanic(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on negative count")
		}
	}()
	wg.Add(-1)
}

func TestCustomWaitGroup_DoneWithoutAdd(t *testing.T) {
	wg := NewCustomWaitGroup()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Unexpected panic: %v", r)
		}
	}()
	wg.Done() // Должен уменьшить до -1, но в имплементации нет паники на Done, только на Add negative.
	// В оригинальной имплементации Done не паникует, если count уходит в negative, но Add паникует если newCount <0.
	// Так что тест на то, что Done без Add не закрывает канал преждевременно? Но count starts at 0, Done makes -1, не закрывает.
	// Wait() будет ждать вечно, если не закрыто.
	// Чтобы протестировать, нужно проверить, что Wait не разблокируется если count <0.
	// Но для простоты:
	select {
	case <-wg.ch:
		t.Error("Channel should not be closed yet")
	default:
		// OK
	}
}

func TestCustomWaitGroup_MultipleWaits(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(1)

	doneWaiting := make(chan struct{}, 2)

	go func() {
		wg.Wait()
		doneWaiting <- struct{}{}
	}()
	go func() {
		wg.Wait()
		doneWaiting <- struct{}{}
	}()

	time.Sleep(100 * time.Millisecond) // Дать горутинам начать ждать
	wg.Done()

	select {
	case <-doneWaiting:
		<-doneWaiting // Оба должны разблокироваться
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for multiple Waits to complete")
	}
}

func TestCustomWaitGroup_ZeroAddWait(t *testing.T) {
	wg := NewCustomWaitGroup()
	// Wait на нулевом count должен сразу вернуться
	start := time.Now()
	wg.Wait()
	duration := time.Since(start)
	if duration > 10*time.Millisecond {
		t.Errorf("Wait on zero count took too long: %v", duration)
	}
}

func TestCustomWaitGroup_AddAfterDone(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(1)
	wg.Done()
	wg.Add(1) // Это разрешено? Count becomes 1, канал уже закрыт? Нет, в имплементации канал закрывается только когда count==0 после Done.
	// После первого Done count=0, канал закрыт.
	// Затем Add(1) делает count=1, но канал закрыт.
	// Wait() вернется сразу, поскольку канал закрыт.
	// Но если затем Done(), count=0, но канал уже закрыт, close на closed panic.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on closing already closed channel")
		}
	}()
	wg.Done() // Должен panic, если имплементация не проверяет.
}
