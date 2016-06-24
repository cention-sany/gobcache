package gobcache

import (
	"bytes"
	"encoding/gob"
	"log"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
)

type Cache struct {
	mc *memcache.Client
}

func NewCache(hostAndPort string) *Cache {
	return &Cache{mc: memcache.New(hostAndPort)}
}

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([]map[string]interface{}{})
}

func (c *Cache) SaveInMemcache(key string, toStore interface{}) error {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	if err := enc.Encode(toStore); err != nil {
		log.Println("Error on `SaveInMemcache` while encoding", err)
		return err
	}
	item := &memcache.Item{
		Key:   key,
		Value: data.Bytes(),
	}

	if err := c.mc.Set(item); err != nil && err != memcache.ErrNoServers {
		log.Println("Error on `SaveInMemcache`: ", err)
		return err
	}
	return nil
}

func (c *Cache) GetFromMemcache(key string, data interface{}) error {
	item, err := c.mc.Get(key)
	if err != nil && err != memcache.ErrCacheMiss {
		log.Println("Error on `GetFromMemcache`: ", err)
		return err
	} else if err == memcache.ErrCacheMiss {
		return err
	}
	dec := gob.NewDecoder(bytes.NewBuffer(item.Value))
	if err := dec.Decode(data); err != nil {
		log.Println("FetchData Decode - `GetFromMemcache` failed: ", err)
		return err
	}

	return nil
}

func (c *Cache) SetRawToMemcache(key, toStore string) error {
	item := &memcache.Item{
		Key:   key,
		Value: []byte(toStore),
	}
	if err := c.mc.Set(item); err != nil && err != memcache.ErrNoServers {
		log.Println("Error `SetRawToMemcache`: ", err)
		return err
	}
	return nil
}

func (c *Cache) GetRawFromMemcache(key string) (*memcache.Item, error) {
	item, err := c.mc.Get(key)
	return item, err
}

func (c *Cache) FlushMemcache() {
	c.mc.FlushAll()
}

func (c *Cache) DeleteFromMemcache(key string) {
	n := strings.LastIndex(key, "/")
	if !(n < 0) {
		c.mc.Delete(key[:n]) // delete ferite generate key too
	}
	c.mc.Delete(key)
}

func (c *Cache) DeleteAllFromMemcache() {
	c.mc.DeleteAll()
}
