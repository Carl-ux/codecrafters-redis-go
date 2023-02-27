package main

type store struct {
	data map[string]string
}

func newStore() *store {
	return &store{
		data: make(map[string]string),
	}
}

func (kv *store) GET(key string) string {
	return kv.data[key]
}

func (kv *store) SET(key string, value string) {
	kv.data[key] = value
}
