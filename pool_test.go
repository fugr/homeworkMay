package structure

import (
	"testing"
	"unsafe"
)

func testCachePool(c Cache, convert func(unsafe.Pointer) Cache, t *testing.T) {
	var (
		first, last unsafe.Pointer
		pool        = NewPool(c)
	)
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
		cache := convert(last)
		pool.Put(cache)
	}
}

func TestInts(t *testing.T) {
	a := Ints{}
	for i := range a {
		a[i] = i
	}
	a.Reset()
	for i := range a {
		if a[i] != 0 {
			t.Error(i, a[i], "!= 0")
		}
	}

	if sizeof := a.Sizeof(); sizeof != len(a)*8 {
		t.Errorf("Sizeof:want %d got %d", sizeof, len(a)*8)
	}

	testCachePool(&a,
		func(ptr unsafe.Pointer) Cache {
			return (*Ints)(ptr)
		},
		t)
}

func TestComposite(t *testing.T) {
	i := 0x1000
	c := Composite{
		a:     0x1F,
		b:     0xEF0000FF,
		s:     "abcdefg",
		p:     &i,
		m:     make(map[string]string, 10),
		slice: make([]int, 10, 20),
		array: [size]int{1, 2, 3, 4, 5},
	}

	c.Reset()

	if c.a != 0 ||
		c.b != 0 ||
		c.s != "" ||
		c.p != nil ||
		c.m != nil ||
		c.slice != nil {
		t.Errorf("Unexpected,got %+v", c)
	}
	for i := range c.array {
		if c.array[i] != 0 {
			t.Error(i, c.array[i], "!= 0")
		}
	}

	if sizeof := c.Sizeof(); sizeof != 384 {
		t.Errorf("Sizeof:want %d got %d", sizeof, 384)
	}

	testCachePool(&c,
		func(ptr unsafe.Pointer) Cache {
			return (*Composite)(ptr)
		},
		t)
}
