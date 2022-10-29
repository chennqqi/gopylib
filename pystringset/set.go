package pystringset

import (
	"sort"
)

type StringSet struct {
	m map[string]bool
}

func New() *StringSet {
	return &StringSet{
		m: make(map[string]bool),
	}
}

func NewStrings(ss []string) *StringSet {
	r := &StringSet{
		m: make(map[string]bool),
	}
	for _, s := range ss {
		r.m[s] = true
	}
	return r
}

func (self *StringSet) Add(v string) {
	self.m[v] = true
}

func (self *StringSet) Remove(v string) {
	//  python: remove() 方法在移除一个不存在的元素时会发生错误
	delete(self.m, v)
}

func (self *StringSet) Discard(v string) {
	_, exists := self.m[v]
	if exists {
		delete(self.m, v)
	}
}

// pop() random pop an element
func (self *StringSet) Pop() string {
	var ret string
	for k, _ := range self.m {
		ret = k
		break
	}
	delete(self.m, ret)
	return ret
}

func (self *StringSet) Clear() {
	self.m = make(map[string]bool)
}

// 交集 & = s & t
// return new set with elements common to s and t
func (self *StringSet) Intersection(other *StringSet) *StringSet {
	set := New()
	for k, _ := range self.m {
		if other.In(k) {
			set.Add(k)
		}
	}
	return set
}

// 并集 | = s | t
// return new set with elements from both s and t
func (self *StringSet) Union(other *StringSet) *StringSet {
	set := New()
	for k, _ := range self.m {
		set.Add(k)
	}
	for k, _ := range other.m {
		set.Add(k)
	}
	return set
}

// 差集 - = s - t
// return new set with elements in s but not in t
func (self *StringSet) Difference(other *StringSet) *StringSet {
	set := New()
	for k, _ := range self.m {
		if !other.In(k) {
			set.Add(k)
		}
	}

	return set
}

// new set with a shallow copy of s
func (self *StringSet) In(v string) bool {
	_, exists := self.m[v]
	return exists
}

// new set with a shallow copy of s
func (self *StringSet) Copy() *StringSet {
	set := New()
	self.Each(func(x string) bool {
		set.Add(x)
		return true
	})
	return set
}

// 对称差 ^ = s ^ t , 对称差相当于两个集合的并集减去交集
func (self *StringSet) SymmetricDifference(other *StringSet) *StringSet {
	u := self.Union(other)
	i := self.Intersection(other)
	return u.IntersectionUpdate(i)
}

func (self *StringSet) Len() int {
	return len(self.m)
}

// 子集 s <= t
func (self *StringSet) IsSubset(other *StringSet) bool {
	for k, _ := range self.m {
		if !other.In(k) {
			return false
		}
	}
	return true
}

// 超集 s >= t
func (self *StringSet) IsSuperset(other *StringSet) bool {
	for k, _ := range other.m {
		if !self.In(k) {
			return false
		}
	}
	return true
}

// 是否有共同元素
func (self *StringSet) IsDisjoint(other *StringSet) bool {
	set := self.Intersection(other)
	return set.Len() <= 0
}

// 	s |= t
func (self *StringSet) Update(other *StringSet) {
	for k, _ := range other.m {
		self.m[k] = true
	}
}

// s &= t
func (self *StringSet) IntersectionUpdate(other *StringSet) *StringSet {
	self.m = self.Intersection(other).m
	return self
}

// s -= t
func (self *StringSet) DiffenceUpdate(other *StringSet) {
	self.m = self.Difference(other).m
}

// s ^= t
func (self *StringSet) SymmetricDifferenceUpdate(other *StringSet) {
	self.m = self.SymmetricDifference(other).m
}

func (self *StringSet) Equals(other *StringSet) bool {
	return self.Len() == other.Len() && self.IsSubset(other)
}

func (self *StringSet) Each(fn func(x string) bool) {
	for k, _ := range self.m {
		if !fn(k) {
			break
		}
	}
}

func (self *StringSet) Values() []string {
	var index int
	out := make([]string, len(self.m))
	self.Each(func(x string) bool {
		out[index] = x
		index++
		return true
	})
	return out
}

func (self *StringSet) SortedValues() []string {
	var index int
	out := make([]string, len(self.m))
	self.Each(func(x string) bool {
		out[index] = x
		index++
		return true
	})
	sort.Strings(out)
	return out
}
