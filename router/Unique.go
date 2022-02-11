package router

import (
	"strings"
	"sync"
)

type UniqueString struct {
	sort []string
	list map[string]bool
	lock *sync.Mutex
}

func NewUniqueString() *UniqueString {
	return &UniqueString{sort: make([]string, 0), list: make(map[string]bool), lock: new(sync.Mutex)}
}

func (this *UniqueString) Append(uni string) *UniqueString {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.sort = append(this.sort, uni)
	this.list[uni] = true
	return this
}

func (this *UniqueString) Exist(key string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
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
