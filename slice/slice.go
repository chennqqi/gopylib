package slice

type Slice[T comparable] []T

func (s Slice[T]) Contains(v T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return true
		}
	}
	return false
}
