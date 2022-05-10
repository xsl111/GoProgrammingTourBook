package fifo

import (
	"GoProgrammingTourBook/cache"
	"container/list"
)

//fifo是一个FIFO cache,它并不是并发安全的
type fifo struct {
	//缓存的最大容量,单位是字节
	maxBytes int
	//当一个entry从缓存中移除时,调用该函数
	onEvicted func(key string, value interface{})

	//已经使用的字节数,只包括值,不包括key
	usedBytes int
	ll        *list.List
	cache     map[string]*list.Element
}

type entry struct {
	key   string
	value interface{}
}

func (e *entry) Len() int {
	return cache.CalcLen(e.value)
}

// New 创建一个新的 Cache，如果 maxBytes 是 0，表示没有容量限制
func New(maxBytes int, onEvicted func(key string, value interface{})) cache.Cache {
	return &fifo{
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
	}
}

// Set 往 Cache 尾部增加一个元素 (如果已经存在，则移至尾部，并修改值)
func (f *fifo) Set(key string, value interface{}) {
	if e, ok := f.cache[key]; ok {
		f.ll.MoveToBack(e)
		en := e.Value.(*entry)
		f.usedBytes = f.usedBytes - cache.CalcLen(en.value) + cache.CalcLen(value)
		en.value = value
		return
	}

	en := &entry{key, value}
	e := f.ll.PushBack(en)
	f.cache[key] = e

	f.usedBytes += en.Len()
	if f.maxBytes > 0 && f.usedBytes > f.maxBytes {
		f.DelOldest()
	}
}

// Get 从 cache 中获取 key 对应的值，nil 表示 key 不存在
func (f *fifo) Get(key string) interface{} {
	if e, ok := f.cache[key]; ok {
		return e.Value.(*entry).value
	}
	return nil
}

// Del 从 cache 中删除 key 对应的记录
func (f *fifo) Del(key string) {
	if e, ok := f.cache[key]; ok {
		f.removeElement(e)
	}
}

// DelOldest 从 cache 中删除最旧的记录
func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Front())
}

func (f *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}

	f.ll.Remove(e)
	en := e.Value.(*entry)
	f.usedBytes -= en.Len()
	delete(f.cache, en.key)

	if f.onEvicted != nil {
		f.onEvicted(en.key, en.value)
	}
}

func (f *fifo) Len() int {
	return f.ll.Len()
}
