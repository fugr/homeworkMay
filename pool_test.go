package structure

import (
	"testing"
	"unsafe"
)

func TestArrayPool(t *testing.T) {
	a := &Ints{}

	pool := NewNewPool(a)

	var first, last unsafe.Pointer

	for {
		pointer, err := pool.Get()
		if err != nil {
			t.Log(err)
			break
		}
		t.Logf("%p", pointer)

		first = unsafe.Pointer(pointer)
		if last != nil && first != last {
			t.Errorf("Unexpected,address from pool isnot the last one,%p!=%p", first, last)
		}

		pointer1, err := pool.Get()
		if err != nil {
			t.Log(err)
			break
		}
		t.Logf("%p", pointer1)
		last = unsafe.Pointer(pointer1)

		if first == last {
			t.Errorf("Unexpected,got same address from pool,%p==%p", first, last)
		}

		// put back last address to pool
		cache := (*Ints)(last)
		pool.Put(cache)
	}
}
