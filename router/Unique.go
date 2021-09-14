package router

import "strings"

type UniqueString struct {
	sort []string
	list map[string]bool
}

func NewUniqueString() *UniqueString {
	return &UniqueString{make([]string, 0), make(map[string]bool)}
}

func (this *UniqueString) Append(uni string) *UniqueString {
	this.sort = append(this.sort, uni)
	this.list[uni] = true
	return this
}

func (this *UniqueString) Exist(key string) bool {
	_, ok := this.list[key]
	return ok
}

func (this *UniqueString) String() string {
	return strings.Join(this.sort, "/")
}

func (this *UniqueString) ForEach(callable StringCallable) {
	for _, index := range this.sort {
		callable(index, this.list[index])
	}
}

func (this *UniqueString) Sort() []string {
	return this.sort
}

func (this *UniqueString) List() map[string]bool {
	return this.list
}
