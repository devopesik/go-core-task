package main

import (
	"fmt"
	"sync"
)

type StringIntMap struct {
	m  map[string]int
	mu sync.RWMutex
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		m:  make(map[string]int),
		mu: sync.RWMutex{},
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[key] = value
}

func (m *StringIntMap) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.m[key]; ok {
		delete(m.m, key)
	}
}

func (m *StringIntMap) Copy() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	copied := make(map[string]int)
	for k, v := range m.m {
		copied[k] = v
	}
	return copied
}

func (m *StringIntMap) Exists(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.m[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *StringIntMap) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return fmt.Sprintf("%v", m.m)
}

func main() {
	m := NewStringIntMap()
	m.Add("1", 1)
	m.Add("2", 2)
	m.Add("3", 3)
	m.Add("4", 4)
	copied := m.Copy()
	m.Remove("2")
	fmt.Println("copied: ", copied)
	fmt.Println("original: ", m)
	flag := m.Exists("1")
	flag2 := m.Exists("2")
	fmt.Printf("flag %v\n", flag)
	fmt.Printf("flag2 %v\n", flag2)
}
