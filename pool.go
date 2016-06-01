package structure

import (
	"errors"
	"unsafe"
)

const (
	cap  = 1 << 20
	size = 1<<5 + 7
)

var ErrPoolEmpty = errors.New("empty pool")

type Cache interface {
	Pointer() unsafe.Pointer
	// sizeof memory
	Sizeof() int
	// reset all elems tobe zero value
	Reset()
}

type Ints [size]int

func (a Ints) Sizeof() int {
	return int(unsafe.Sizeof(a))
}

func (a *Ints) Reset() {
	for i := range a {
		a[i] &= 0
	}
}

func (a *Ints) Pointer() unsafe.Pointer {
	return unsafe.Pointer(a)
}

type ArrayPool struct {
	head  unsafe.Pointer
	pools [cap]int
}

type list struct {
	next unsafe.Pointer
}

func NewNewPool(c Cache) *ArrayPool {
	ap := new(ArrayPool)

	ap.head = unsafe.Pointer(&ap.pools[0])
	size := c.Sizeof()

	for i := 0; ; i += size {
		current := (*list)(unsafe.Pointer(&ap.pools[i]))
		next := i + size

		if next < len(ap.pools) {
			current.next = unsafe.Pointer(&ap.pools[next])
		} else {
			current.next = nil
			break
		}
	}

	return ap
}

func (ap *ArrayPool) Get() (unsafe.Pointer, error) {
	if ap.head == nil {
		return nil, ErrPoolEmpty
	}

	ptr := ap.head

	node := (*list)(ptr)
	ap.head = node.next

	return ptr, nil
}

func (ap *ArrayPool) Put(c Cache) {
	c.Reset()

	head := ap.head
	pointer := c.Pointer()
	ap.head = pointer
	node := (*list)(pointer)
	node.next = head
}
