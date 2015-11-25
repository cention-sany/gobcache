package gobcache

import (
	"log"
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
type Cookie struct {
	UserId        int
	LastLoginTime int64
	LoggedIn      bool
}

func TestFetchKeys(t *testing.T){
	key := "SSS_MTQ0NzkyMzQ5NXxaUGVLOWJ3YzdXSDZqblo4dWJIOF9wSVN3eXZZZzFiSTZvaXFXYXhmWWQ4T2pwQ2poaVhBVW5hdDlKS0RSTE5Velc5cV9fMjY4QVlCZXhjSThkeFU4ZTg0fB7AnGKfhELQGLOjw0vQhPmJvPY5Vpj5H0yzLWtqvQtb"
	want := &Cookie{
		2,1447925985,true,
	}	
	got := new(Cookie)
	err := GetFromMemcache(key, &got)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("[TestFetchKeys]->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
	}

}

type Cookie struct {
	UserId        int
	LastLoginTime int64
	LoggedIn      bool
}

func TestFetchKeys(t *testing.T) {
	key := "Session_MTQ0Nzk5OTYyNnwyUlJRRm93eW9BQm1Bc3M2OE9jWDlhOHMtRGpCd1hSWExDNldhdElybGZZX1V4ZWxPMGlkaG5WenNlZXJOMDFKZURYMGNSbmdpT2tURFFGQUFJQl8yWHBXfILxGWO0ac0DQuQwt8MXSmXaSjKmTEczJ9cafN4gKzs4"
	want := &Cookie{
		2, 1447925985, true,
	}
	got := new(Cookie)
	err := GetFromMemcache(key, &got)
	if err != nil {
		t.Error(err)
	}
	log.Println("Got:", got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("[TestFetchKeys]->Key[%s]:\nWant: %v\n Got: %v", key, want, got)
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
