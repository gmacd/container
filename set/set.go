package set

type Set struct {
	m map[interface{}]struct{}
}

func NewSetOfValues(items ...interface{}) *Set {
	set := &Set{make(map[interface{}]struct{})}
	set.AddMany(items...)
	return set
}

func (set *Set) Add(item interface{}) {
	set.m[item] = struct{}{}
}

func (set *Set) Remove(item interface{}) {
	delete(set.m, item)
}

func (set *Set) AddMany(items ...interface{}) {
	for _, item := range items {
		set.m[item] = struct{}{}
	}
}

func (set *Set) RemoveMany(items ...interface{}) {
	for _, item := range items {
		delete(set.m, item)
	}
}

func (set *Set) Contains(item interface{}) bool {
	_, ok := set.m[item]
	return ok
}

func (set *Set) Len() int {
	return len(set.m)
}

func (set *Set) Clear() {
	set.m = make(map[interface{}]struct{})
}

func (set *Set) Foreach(action func(interface{})) {
	for item, _ := range set.m {
		action(item)
	}
}
