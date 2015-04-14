package gobcache

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
	if want != got {
		t.Errorf("[TestSaveInMemcache]->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}

}

func TestSaveArrayInMemcache(t *testing.T) {
	key := "cention1"
	want := []string{"cention contact centre", "Test", "Test2"}
	if err := SaveInMemcache(key, want); err != nil {
		t.Error(err)
	}
	var got []string
	err := GetFromMemcache(key, &got)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("[TestSaveInMemcache]->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}

}

func TestSaveStructMemcache(t *testing.T) {
	key := "cention_struct"
	var want, got struct {
		Name    string
		Id      int
		Country string
	}
	want.Name = "Mujibur"
	want.Id = 9007
	want.Country = "BD"
	if err := SaveInMemcache(key, want); err != nil {
		t.Error(err)
	}
	if err := GetFromMemcache(key, &got); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("[TestSaveStructMemcache] ->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}
}
func TestAnonymousSaveInMemcache(t *testing.T) {
	key := "anonymous"
	want := []struct {
		Name string
		Data map[string]string
	}{
		{"Mujibur", map[string]string{"as": "cention contact centre", "a1": "Test", "b2": "Test2"}},
		{"Mujibur1", map[string]string{"as": "cention contact centre1", "a1": "Test1", "b2": "Test21"}},
	}
	got := []struct {
		Name string
		Data map[string]string
	}{
		{"Mujibur", map[string]string{"as": "cention contact centre", "a1": "Test", "b2": "Test2"}},
		{"Mujibur1", map[string]string{"as": "cention contact centre1", "a1": "Test1", "b2": "Test21"}},
	}
	if err := SaveInMemcache(key, want); err != nil {
		t.Error(err)
	}

	if err := GetFromMemcache(key, &got); err != nil {
		t.Error(err)
	}
	if reflect.DeepEqual(want, got) {
		t.Errorf("[TestSaveStructMemcache] ->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}
}

func TestDeleteMemcache(t *testing.T) {
	//FlushMemcache()
	DeleteFromMemcache("cention")
}
