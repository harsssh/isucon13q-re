package main

import "sync"

type Cache struct {
	IconCache sync.Map
}

var cache Cache

func (c *Cache) GetIconCache(userId int64) []byte {
	if v, ok := c.IconCache.Load(userId); ok {
		return v.([]byte)
	}
	return nil
}

func (c *Cache) SetIconCache(userId int64, icon []byte) {
	c.IconCache.Store(userId, icon)
}
