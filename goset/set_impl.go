package goset

type None struct{}

var none None

type Key interface {
	~int | ~string | ~float32 | ~int64 | ~float64 | ~int32
}

type Set[K Key] struct {
	item map[K]None
}

// NewSet new一个Set对象
func NewSet[K Key]() *Set[K] {
	return &Set[K]{
		item: make(map[K]None),
	}
}

// Add 新增key，会去重
func (s *Set[K]) Add(key K) {
	s.item[key] = none
}

// Size 计算size
func (s *Set[K]) Size() int {
	return len(s.item)
}

func (s *Set[K]) Remove(key K) {
	if _, ok := s.item[key]; !ok {
		return
	}
	delete(s.item, key)
}

// Values  获取全部key
func (s *Set[K]) Values() []K {
	var values []K
	for value := range s.item {
		values = append(values, value)
	}
	return values
}

func (s *Set[K]) Find(key K) bool {
	if _, ok := s.item[key]; ok {
		return true
	}
	return false
}
