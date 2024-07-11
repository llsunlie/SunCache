package lru

import (
	"SunCache/cache/core"
	"container/list"
)

type Cache struct {
	maxBytes int64
	useBytes int64
	linkList *list.List
	elemtMap map[string]*list.Element
}

type entry struct {
	key   string
	value core.Value
}

func NewCache(maxBytes int64) (cache *Cache) {
	cache = &Cache{
		maxBytes: maxBytes,
		useBytes: 0,
		linkList: list.New(),
		elemtMap: make(map[string]*list.Element),
	}
	return
}

func (c *Cache) get(key string) (value core.Value, ok bool) {
	if elemt, ok := c.elemtMap[key]; ok {
		c.linkList.MoveToFront(elemt)
		entry := elemt.Value.(*entry)
		return entry.value, true
	}
	return
}

func (c *Cache) removeOldest() {
	if elemt := c.linkList.Back(); elemt != nil {
		c.linkList.Remove(elemt)
		entry := elemt.Value.(*entry)
		delete(c.elemtMap, entry.key)
		c.useBytes -= int64(len(entry.key) + entry.value.Len())
	}
}

func (c *Cache) add(key string, value core.Value) {
	if elemt, ok := c.elemtMap[key]; ok {
		c.linkList.MoveToFront(elemt)
		entry := elemt.Value.(*entry)
		c.useBytes += int64(value.Len() - entry.value.Len())
		entry.value = value
	} else {
		elemt := c.linkList.PushFront(&entry{key, value})
		c.elemtMap[key] = elemt
		c.useBytes += int64(len(key) + value.Len())
	}
	for c.useBytes > c.maxBytes {
		c.removeOldest()
	}
}
