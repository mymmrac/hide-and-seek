package collection

import "sync"

type SyncMap[K comparable, V any] struct {
	sync.RWMutex
	values map[K]V
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		RWMutex: sync.RWMutex{},
		values:  make(map[K]V),
	}
}

func (m *SyncMap[K, V]) Set(key K, value V) {
	m.Lock()
	m.values[key] = value
	m.Unlock()
}

func (m *SyncMap[K, V]) Remove(key K) {
	m.Lock()
	delete(m.values, key)
	m.Unlock()
}

func (m *SyncMap[K, V]) Get(key K) (V, bool) {
	m.RLock()
	value, ok := m.values[key]
	m.RUnlock()
	return value, ok
}

func (m *SyncMap[K, V]) GetAndRemove(key K) (V, bool) {
	m.Lock()
	value, ok := m.values[key]
	if ok {
		delete(m.values, key)
	}
	m.Unlock()
	return value, ok
}

func (m *SyncMap[K, V]) Has(key K) bool {
	m.RLock()
	_, ok := m.values[key]
	m.RUnlock()
	return ok
}

func (m *SyncMap[K, V]) ForEach(f func(key K, value V) bool) {
	m.RLock()
	for k, v := range m.values {
		if !f(k, v) {
			break
		}
	}
	m.RUnlock()
}

func (m *SyncMap[K, V]) Raw() map[K]V {
	return m.values
}
