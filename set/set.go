//@Date 2024/12/6 21:30
//@Desc

package set

// Set 是一个通用集合接口，支持添加、删除、检查元素和获取所有元素的功能
// 使用泛型 T 来表示集合中的元素类型，元素类型必须实现 comparable 接口
type Set[T comparable] interface {
	// Add 向集合中添加元素
	Add(key T)
	// Delete 从集合中删除指定的元素
	Delete(key T)
	// Exist 检查元素是否存在于集合中
	Exist(key T) bool
	// Keys 返回集合中所有元素的切片，顺序不固定
	Keys() []T
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

// NewMapSet 创建一个新的 MapSet 实例，并初始化一个空的 map
// size 用于初始化 map 的容量，提供一定的预期大小优化性能
func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

// Add 向集合中添加一个元素
// 如果元素已存在，map 会忽略重复插入
func (s *MapSet[T]) Add(val T) {
	s.m[val] = struct{}{}
}

// Delete 从集合中删除指定的元素
// 如果元素不存在，map 会忽略该操作
func (s *MapSet[T]) Delete(key T) {
	delete(s.m, key)
}

// Exist 检查指定的元素是否存在于集合中
// 返回 true 如果元素存在，否则返回 false
func (s *MapSet[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

// Keys 返回集合中所有元素的切片
// 返回的元素顺序不固定，因为 Go 的 map 不保证遍历顺序
func (s *MapSet[T]) Keys() []T {
	ans := make([]T, 0, len(s.m))
	for key := range s.m {
		ans = append(ans, key)
	}
	return ans
}
