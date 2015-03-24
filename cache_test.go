package cache

import (
	"reflect"
	"testing"
)

func TestSaveInMemcache(t *testing.T) {
	key := "cention"
	value := "cention contact centre"
	want := "cention contact centre"
	if err := SaveInMemcache(key, value); err != nil {
		t.Error(err)
	}
	var got string
	err := GetFromMemcache(key, &got)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if want != got {
		t.Errorf("[TestSaveInMemcache]->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}

}

func TestSaveArrayInMemcache(t *testing.T) {
	key := "cention"
	want := []string{"cention contact centre", "Test", "Test2"}
	if err := SaveInMemcache(key, want); err != nil {
		t.Error(err)
	}
	var got []string
	err := GetFromMemcache(key, &got)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("[TestSaveInMemcache]->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}

}
