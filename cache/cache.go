package cache

var cache = make(map[string]interface{})

func Set(key string, value interface{}) {
	cache[key] = value
}

func Get(key string) interface{} {
	return cache[key]
}
