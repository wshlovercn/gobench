package concurrent

import "sync"

type LockMap struct {
	lock sync.RWMutex
	datas map[string]interface{}
}

func NewLockMap() *LockMap {
	return &LockMap {
		datas:make(map[string]interface{}),
	}
}

func (lm *LockMap) Put(key string, value interface{}) {
	lm.lock.Lock()
	lm.datas[key] = value
	lm.lock.Unlock()
}

func (lm *LockMap) Get(key string) (interface{}, bool) {
	lm.lock.RLock()
	v, ok := lm.datas[key]
	lm.lock.RUnlock()
	return v, ok
}

func (lm *LockMap) Remove(key string)  {
	lm.lock.Lock()
	delete(lm.datas, key)
	lm.lock.Unlock()
}