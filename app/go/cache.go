package main

import "sync"

type Cache struct {
	IconCache     sync.Map
	IconHashCache sync.Map
	TagCache      sync.Map
	ThemeCache    sync.Map
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

func (c *Cache) GetIconHashCache(userId int64) string {
	if v, ok := c.IconHashCache.Load(userId); ok {
		return v.(string)
	}
	return ""
}

func (c *Cache) SetIconHashCache(userId int64, hash string) {
	c.IconHashCache.Store(userId, hash)
}
