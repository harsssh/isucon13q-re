package main

import "sync"

func getOrInsertMap(m *sync.Map, key any, f func() any) any {
	if v, ok := m.Load(key); ok {
		return v
	}
	v := f()
	m.Store(key, v)
	return v
}
