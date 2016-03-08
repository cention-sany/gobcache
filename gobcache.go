package gobcache

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

var mc = &memcache.Client{}
var MemcacheStatus bool

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register([]map[string]interface{}{})
}

func SaveInMemcache(key string, toStore interface{}, HostnPort string) error {
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
	mc = memcache.New(HostnPort)
	if err := mc.Set(item); err != nil && err != memcache.ErrNoServers {
		log.Println("Datastore - `SaveInMemcache` ", err)
		return err
	}
	return nil
}

func GetFromMemcache(key string, data interface{}, HostnPort string) error {
	mc = memcache.New(HostnPort)
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

func SetRawToMemcache(key, toStore, HostnPort string) error {
	item := &memcache.Item{
		Key:   key,
		Value: []byte(toStore),
	}
	mc = memcache.New(HostnPort)
	if err := mc.Set(item); err != nil && err != memcache.ErrNoServers {
		log.Println("RawDataSave - `SetRawToMemcache` ", err)
		return err
	}
	return nil
}

func GetRawFromMemcache(key, HostnPort string) (*memcache.Item, error) {
	mc = memcache.New(HostnPort)
	item, err := mc.Get(key)
	return item, err
}

func FlushMemcache(HostnPort string) {
	mc = memcache.New(HostnPort)
	mc.FlushAll()
}
func DeleteFromMemcache(key string, HostnPort string) {
	mc = memcache.New(HostnPort)
	mc.Delete(key)
}
func DeleteAllFromMemcache(HostnPort string) {
	mc = memcache.New(HostnPort)
	mc.DeleteAll()
}
