package expirymap

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	m := New[int, bool](time.Minute, time.Minute)
	if m == nil {
		t.Fatalf("expected ExpiryMap, got nil")
	}
}

func TestGetUnset(t *testing.T) {
	key := 1

	m := New[int, string](time.Minute, time.Minute)

	v, found := m.Get(key)
	if !reflect.DeepEqual(v, "") {
		t.Fatalf("expected %v, got %v", "", v)
	}

	if found {
		t.Fatalf("expected %v, got %v", false, found)
	}
}

func TestSetGet(t *testing.T) {
	key := 1
	value := "alice"

	m := New[int, string](time.Minute, time.Minute)
	m.Set(key, value)

	if v, _ := m.Get(key); !reflect.DeepEqual(v, value) {
		t.Fatalf("expected %v, got %v", value, v)
	}
}

func TestSetDeleteGet(t *testing.T) {
	key := 2
	value := "bob"

	m := New[int, string](time.Minute, time.Minute)
	m.Set(key, value)
	m.Delete(key)

	if v, _ := m.Get(key); !reflect.DeepEqual(v, "") {
		t.Fatalf("expected %v, got %v", "", v)
	}
}

func TestDeleteUnset(t *testing.T) {
	key := 3

	m := New[int, string](time.Minute, time.Minute)
	m.Delete(key)
}

func TestLen(t *testing.T) {
	key := 4
	value := "charlie"

	m := New[int, string](time.Minute, time.Minute)

	if l := m.Len(); l != 0 {
		t.Fatalf("expected %v, got %v", 0, l)
	}

	m.Set(key, value)

	if l := m.Len(); l != 1 {
		t.Fatalf("expected %v, got %v", 1, l)
	}
}

func TestIterate(t *testing.T) {
	key := 5
	value := "dave"

	m := New[int, string](time.Minute, time.Minute)
	m.Set(key, value)

	var count int
	var k int
	var v string
	for k, v = range m.Iterate() {
		count++
	}

	if count != 1 {
		t.Fatalf("expected %v, got %v", 1, count)
	}

	if k != key {
		t.Fatalf("expected %v, got %v", key, k)
	}

	if v != value {
		t.Fatalf("expected %v, got %v", value, v)
	}
}

func TestClear(t *testing.T) {
	key := 6
	value := "eve"

	m := New[int, string](time.Minute, time.Minute)
	m.Set(key, value)
	m.Clear()

	if l := m.Len(); l != 0 {
		t.Fatalf("expected %v, got %v", 0, l)
	}
}

func TestStop(t *testing.T) {
	m := New[int, string](time.Minute, time.Minute)
	m.Stop()
}

func TestGargabeClean(t *testing.T) {
	key := 7
	value := "frank"

	m := New[int, string](time.Nanosecond, time.Millisecond)
	m.Set(key, value)

	if v, _ := m.Get(key); !reflect.DeepEqual(v, value) {
		t.Fatalf("expected %v, got %v", value, v)
	}

	time.Sleep(time.Millisecond * 2)

	if v, _ := m.Get(key); v != "" {
		t.Fatalf("expected nil, got %v", v)
	}
}
