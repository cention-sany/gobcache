package gobcache

import (
	"reflect"
	"testing"
)

var (
	HostnPort = "localhost:11211"
)

func TestSaveInMemcache(t *testing.T) {
	key := "cention"
	value := "cention contact centre"
	want := "cention contact centre"
	if err := SaveInMemcache(key, value, HostnPort); err != nil {
		t.Error(err)
	}
	var got string
	err := GetFromMemcache(key, &got, HostnPort)
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
	if err := SaveInMemcache(key, want, HostnPort); err != nil {
		t.Error(err)
	}
	var got []string
	err := GetFromMemcache(key, &got, HostnPort)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("[TestSaveInMemcache]->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}

}

type Cookie struct {
	UserId        int
	LastLoginTime int64
	LoggedIn      bool
}

func TestFetchKeys(t *testing.T) {
	HostPort := "localhost:11311"
	key := "Session_MTQ1NzQwMzY5NHxFeGF6dmJjWFlJcTBJV0ZBZm4xV0dCZDZ2UGYydnQzMlV0aFdZUVZkdUswbndjVnpYSkhTdjdkVGVxQlFmRkduYk43VHFZLWlBWFh0VEJGZWNxUHRySDFtfOij83kJOfkNapmefRIVM4TUahF0MagwoxmhTn768DoC"
	//	want := &Cookie{
	//		3, 1457416558, true,
	//	}
	//got := new(Cookie)
	want := "3/1457419037/true"
	get := ""
	sItems, err := GetRawFromMemcache(key, HostPort)
	if err != nil {
		t.Error(err)
	}
	get = string(sItems.Value)
	if want == get {
		t.Errorf("[TestFetchKeys]->Key[%s]:\nWant: %v\n Got: %v", key, want, get)
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
	if err := SaveInMemcache(key, want, HostnPort); err != nil {
		t.Error(err)
	}
	if err := GetFromMemcache(key, &got, HostnPort); err != nil {
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
	if err := SaveInMemcache(key, want, HostnPort); err != nil {
		t.Error(err)
	}

	if err := GetFromMemcache(key, &got, HostnPort); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("[TestAnonymousSaveInMemcache] ->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}
}

func TestDeleteMemcache(t *testing.T) {
	//FlushMemcache()
	DeleteFromMemcache("cention", HostnPort)
}
