package main

import "time"

type ValueWithExpiry struct {
	value     string
	expiresAt time.Time
}

type store struct {
	data map[string]ValueWithExpiry
}

func (v ValueWithExpiry) IsExpired() bool {
	if v.expiresAt.IsZero() {
		return false
	}
	return v.expiresAt.Before(time.Now())
}

func newStore() *store {
	return &store{
		data: make(map[string]ValueWithExpiry),
	}
}

func (kv *store) GET(key string) (string, bool) {
	valueWithExpiry, ok := kv.data[key]
	if !ok {
		return "", false
	}

	if valueWithExpiry.IsExpired() {
		delete(kv.data, key) //delete expired data
		return "", false
	}
	return valueWithExpiry.value, true
}

func (kv *store) SET(key string, value string) {
	kv.data[key] = ValueWithExpiry{value: value}
}

func (kv *store) SetWithExpiry(key string, value string, expiry time.Duration) {
	kv.data[key] = ValueWithExpiry{
		value:     value,
		expiresAt: time.Now().Add(expiry),
	}
}
