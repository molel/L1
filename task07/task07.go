package main

import "sync"

// используется sync.RWMutex, а не sync.Mutex так как он лучше оптимизирован для конкурентного чтения
type CMap struct {
	sync.RWMutex
	m map[int]string
}

func NewCMap() *CMap {
	return &CMap{
		m: make(map[int]string),
	}
}

func (cm *CMap) Set(key int, value string) {
	cm.Lock()
	defer cm.Unlock()

	cm.m[key] = value
}

func (cm *CMap) Get(key int) string {
	cm.RLock()
	defer cm.RUnlock()

	return cm.m[key]
}
