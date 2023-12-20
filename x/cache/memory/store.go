package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/go-packagist/go-kratos-components/helper"
	"github.com/go-packagist/go-kratos-components/x/cache"
)

var (
	ErrKeyNotFound = errors.New("memory: key not found")
	ErrKeyExpired  = errors.New("memory: key expired")
)

type data struct {
	value   interface{}
	expired time.Time
}

type options struct {
	size int

	lru         bool
	lruDuration time.Duration
}

type Option func(*options)

func WithSize(size int) Option {
	return func(o *options) {
		if size > 0 {
			o.size = size
		}
	}
}

func WithLRU() Option {
	return func(o *options) {
		o.lru = true
	}
}

func WithLRUDuration(duration time.Duration) Option {
	return func(o *options) {
		o.lruDuration = duration
	}
}

type Store struct {
	opt *options

	data map[string]data
	mu   sync.RWMutex // guard data
}

var _ cache.Store = (*Store)(nil)

func New(opts ...Option) *Store {
	o := &options{
		size:        1024,
		lru:         false,
		lruDuration: time.Second * 10,
	}

	for _, opt := range opts {
		opt(o)
	}

	s := &Store{
		data: make(map[string]data, o.size),
		opt:  o,
	}

	if o.lru {
		go s.lru()
	}

	return s
}

func (s *Store) Has(ctx context.Context, key string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if d, ok := s.data[key]; !ok {
		return false, nil
	} else {
		if d.expired.Before(time.Now()) {
			delete(s.data, key) // TODO: the mu is RLock
			return false, nil
		}

		return true, nil
	}
}

func (s *Store) Get(ctx context.Context, key string, dest interface{}) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if d, ok := s.data[key]; !ok {
		return ErrKeyNotFound
	} else {
		if s.isExpired(d.expired) {
			// Todo: Remove the key (Attention: the mu is RLock)
			return ErrKeyExpired
		}

		return helper.ValueOf(d.value, dest)
	}
}

func (s *Store) Put(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = data{
		value:   value,
		expired: time.Now().Add(ttl),
	}

	return true, nil
}

func (s *Store) Increment(ctx context.Context, key string, value int) (int, error) {
	return 0, nil // TODO
}

func (s *Store) Decrement(ctx context.Context, key string, value int) (int, error) {
	return 0, nil // TODO
}

func (s *Store) Forever(ctx context.Context, key string, value interface{}) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = data{
		value:   value,
		expired: time.Time{},
	}

	return true, nil
}

func (s *Store) Forget(ctx context.Context, key string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; !ok {
		return false, nil
	}

	delete(s.data, key)

	return true, nil
}

func (s *Store) Flush(ctx context.Context) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[string]data, s.opt.size)

	return true, nil
}

func (s *Store) GetPrefix() string {
	return ""
}

func (s *Store) isExpired(expired time.Time) bool {
	return expired.Before(time.Now()) || !expired.IsZero()
}

func (s *Store) lru() {
	for {
		time.Sleep(s.opt.lruDuration)

		s.mu.Lock()

		for k, v := range s.data {
			if s.isExpired(v.expired) {
				delete(s.data, k)
			}
		}

		s.mu.Unlock()
	}
}
