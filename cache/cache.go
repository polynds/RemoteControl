package cache

var cache = make(map[string]interface{})

func Set(key string, value interface{}) {
	cache[key] = value
}

func Get(key string, defaultVal interface{}) interface{} {
	if value, ok := cache[key]; ok {
		return value
	}

	return defaultVal
}
