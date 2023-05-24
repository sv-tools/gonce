package gonce

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	o := Once[int64]{}
	res, err := o.Do(func() (result int64, err error) {
		return -1, errors.New("foo")
	})
	if err == nil {
		t.Fatalf("expected an error, but got nil")
	}
	if res != -1 {
		t.Fatalf("expected res = -1, but got %d", res)
	}
	res, err = o.Do(func() (result int64, err error) {
		return rand.Int63(), nil
	})
	if err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
	if res == -1 {
		t.Fatalf("expected non-negative value, but got %d", res)
	}
	res2, err := o.Do(func() (result int64, err error) {
		return rand.Int63(), nil
	})
	if err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
	if res != res2 {
		t.Fatalf("expected same values, but got res: %d; res2: %d", res, res2)
	}
}

var benchmarkSyncOnce int

func BenchmarkSyncOnce(b *testing.B) {
	o := sync.Once{}
	b.RunParallel(func(pb *testing.PB) {
		var res int
		for pb.Next() {
			o.Do(func() {
				res = 42
			})
		}
		benchmarkSyncOnce = res
	})
}

var benchmarkOnce int

func BenchmarkOnce(b *testing.B) {
	o := Once[int]{}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var res int
		for pb.Next() {
			res, _ = o.Do(func() (int, error) {
				return 42, nil
			})
		}
		benchmarkOnce = res
	})
}

func ExampleOnce() {
	o := Once[int64]{}
	res1, err := o.Do(func() (result int64, err error) {
		return rand.Int63(), nil
	})
	if err != nil {
		panic(err)
	}
	res2, err := o.Do(func() (result int64, err error) {
		return rand.Int63(), nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("res1: (%T); res2: (%T); res1 == res2: %v", res1, res2, res1 == res2)
	// Output: res1: (int64); res2: (int64); res1 == res2: true
}
