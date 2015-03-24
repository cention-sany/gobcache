package cache

import (
	"bytes"
	"encoding/gob"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var mc = memcache.New("localhost:11211")

func SaveInMemcache(key string, toStore interface{}) error {
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

	if err := mc.Set(item); err != nil && err != memcache.ErrNoServers {
		log.Println("Datastore - `SaveInMemcache` ", err)
		return err
	}
	return nil
}

func GetFromMemcache(key string, data interface{}) error {
	item, err := mc.Get(key)
	if err != nil && err != memcache.ErrCacheMiss { //Error if nil and key not exists
		log.Println("FetchData - `GetFromMemcache` error: ", err)
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
