package lfu

import (
	"GoProgrammingTourBook/cache"
	"container/heap"
)

//lfu是一个LFU cache，他并不是并发安全的
type lfu struct {
	//缓存的最大容量，单位是字节
	maxBytes int
	//当一个entry从缓存中移除时调用该函数，默认是nil
	onEvicted func(key string, value interface{})

	//已经使用的字节数，只包括值
	usedBytes int
	queue     *queue
	cache     map[string]*entry
}

func New(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	q := make(queue, 0, 1024)
	return &lfu{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		queue:     &q,
		cache:     make(map[string]*entry),
	}
}

func (l *lfu) Set(key string, value interface{}) {
	if e, ok := l.cache[key]; ok {
		l.usedBytes = l.usedBytes - cache.CalcLen(e.value) + cache.CalcLen(value)
		l.queue.update(e, value, e.weight+1)
		return
	}

	en := &entry{key: key, value: value}
	heap.Push(l.queue, en)
	l.cache[key] = en

	l.usedBytes += en.Len()
	if l.maxBytes > 0 && l.usedBytes > l.maxBytes {
		l.removeElement(heap.Pop(l.queue))
	}
}

func (l *lfu) Get(key string) interface{} {
	if e, ok := l.cache[key]; ok {
		l.queue.update(e, e.value, e.weight+1)
		return e.value
	}
	return nil
}

func (l *lfu) Del(key string) {
	if e, ok := l.cache[key]; ok {
		heap.Remove(l.queue, e.index)
		l.removeElement(e)
	}
}

func (l *lfu) DelOldest() {
	if l.queue.Len() == 0 {
		return
	}
	l.removeElement(heap.Pop(l.queue))
}

func (l *lfu) removeElement(x interface{}) {
	if x == nil {
		return
	}

	en := x.(*entry)

	delete(l.cache, en.key)

	l.usedBytes -= en.Len()

	if l.onEvicted != nil {
		l.onEvicted(en.key, en.value)
	}
}

func (l *lfu) Len() int {
	return l.queue.Len()
}
