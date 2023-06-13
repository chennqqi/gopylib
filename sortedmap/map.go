package sortedmap

// not thread safe
type SortedMap[Key comparable, Value any] struct {
	dict map[Key]Value
	keys []Key
}

func New[Key comparable, Value any]() *SortedMap[Key, Value] {
	return &SortedMap[Key, Value]{
		dict: make(map[Key]Value),
		keys: make([]Key, 0),
	}
}

func (self *SortedMap[Key, Value]) Get(key Key) (Value, bool) {
	v, exists := self.dict[key]
	return v, exists
}

func (self *SortedMap[Key, Value]) Set(key Key, value Value) {
	_, exists := self.dict[key]
	if exists {
		self.dict[key] = value
		return
	} else {
		self.dict[key] = value
		self.keys = append(self.keys, key)
	}
}

func (self *SortedMap[Key, Value]) Exists(key Key) bool {
	_, exists := self.dict[key]
	return exists
}

func (self *SortedMap[Key, Value]) Keys() []Key {
	r := make([]Key, len(self.keys))
	r = append(r, self.keys...)
	return r
}

func (self *SortedMap[Key, Value]) Values() []Value {
	r := make([]Value, len(self.dict))
	for _, k := range self.keys {
		r = append(r, self.dict[k])
	}
	return r
}
