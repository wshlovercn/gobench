package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

var lockMap = NewLockMap()
var syncMap = &sync.Map{}

func init()  {
	for i := 0; i < 10000; i ++ {
		k := "key" + string(i)
		v := "value" + string(i)
		lockMap.Put(k, v)
		syncMap.Store(k, v)
	}
}

func TestLockMap(t *testing.T) {
	m := NewLockMap()
	m.Put("key1", "value1")
	v, ok := m.Get("key1")
	fmt.Println(v, ok)
}

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	m.Store("key1", "value1")
	v, ok := m.Load("key1")
	fmt.Println(v, ok)
}

func Benchmark_LockMap_Read(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(256)
	for k := 0; k < 256; k++ {
		go func() {
			for i := 0; i < b.N; i++ {
				key := "key" + string(i)
				_, _ = lockMap.Get(key)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func Benchmark_LockMap_Write10(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(256)
	for k := 0; k < 256; k++ {
		go func() {
			for i := 0; i < b.N; i++ {
				key := "key" + string(i)
				if i % 10 == 0 {
					lockMap.Put(key, key)
				} else {
					_, _ = lockMap.Get(key)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}


func Benchmark_SyncMap_Read(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(256)
	for k := 0; k < 256; k++ {
		go func() {
			for i := 0; i < b.N; i++ {
				key := "key" + string(i)
				_, _ = syncMap.Load(key)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func Benchmark_SyncMap_Write10(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(256)
	for k := 0; k < 256; k++ {
		go func() {
			for i := 0; i < b.N; i++ {
				key := "key" + string(i)
				if i % 10 == 0 {
					syncMap.Store(key, key)
				} else {
					_, _ = syncMap.Load(key)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func Benchmark_CPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := "key" + string(i)
		if i % 10 == 0 {
			syncMap.Store(key, key)
		} else {
			_, _ = syncMap.Load(key)
		}
	}
	fmt.Println("**********    Benchmark_CPU   ******************")
}
