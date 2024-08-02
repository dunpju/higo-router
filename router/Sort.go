package router

type Sort[T int | string] struct {
	sort []T
}

func newSort[T int | string]() *Sort[T] {
	return &Sort[T]{sort: make([]T, 0)}
}

func (this *Sort[T]) Exist(value T) bool {
	for _, v := range this.sort {
		if v == value {
			return true
		}
	}
	return false
}

func (this *Sort[T]) Append(value T) {
	if !this.Exist(value) {
		this.sort = append(this.sort, value)
	}
}

func (this *Sort[T]) Range(fn func(index int, value T) bool) {
	for i, v := range this.sort {
		if !fn(i, v) {
			break
		}
	}
}

func (this *Sort[T]) Len() int {
	return len(this.sort)
}
