package blink

import "sync"

type Blink struct {
	values map[string]interface{}
	mu     sync.RWMutex
}

var blink *Blink

func Instance() *Blink {
	if blink == nil {
		blink = New()
	}

	return blink
}

func New() *Blink {
	return &Blink{
		values: make(map[string]interface{}),
	}
}

func (b *Blink) Put(key string, value interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.values[key] = value
}

func (b *Blink) Get(key string) (interface{}, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	value, ok := b.values[key]

	return value, ok
}

func (b *Blink) Has(key string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	_, ok := b.values[key]

	return ok
}

func (b *Blink) Delete(key string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.values, key)
}

func (b *Blink) Forget(key string) {
	b.Delete(key)
}

func (b *Blink) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.values = make(map[string]interface{})
}

func (b *Blink) Flush() {
	b.Clear()
}

func (b *Blink) Values() map[string]interface{} {
	return b.values
}

func (b *Blink) All() map[string]interface{} {
	return b.values
}

func (b *Blink) Size() int {
	return len(b.values)
}

func (b *Blink) Keys() []string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	keys := make([]string, 0, len(b.values))

	for key := range b.values {
		keys = append(keys, key)
	}

	return keys
}

func (b *Blink) Pull(key string) (interface{}, bool) {
	b.mu.RLock()
	value, ok := b.values[key]
	b.mu.RUnlock()

	if ok {
		b.Delete(key)
	}

	return value, ok
}

// func (b *Blink) Increment(key string, values ...int) int {
// 	value := 1
// 	if len(values) > 0 {
// 		value = values[0]
// 	}
//
// 	oldValue, ok := b.Get(key)
// 	if !ok {
// 		b.Put(key, value)
// 		return
// 	}
//
// 	b.Put(key, 1)
//
// }
