package main

import "testing"

func BenchmarkValueCall(b *testing.B) {
	d := D{
		n: 100,
	}

	for i := 0; i < b.N; i++ {
		_ = getValue(d)
	}
}

func BenchmarkValueIface(b *testing.B) {
	d := D{
		n: 100,
	}

	var valuer digiter = d

	for i := 0; i < b.N; i++ {
		_ = valuer.getValue()
	}
}

func BenchmarkPointerCall(b *testing.B) {
	d := D{
		n: 100,
	}

	for i := 0; i < b.N; i++ {
		_ = getPointer(d)
	}
}

func BenchmarkPointerIface(b *testing.B) {
	d := D{
		n: 100,
	}

	var pointer digiter = d

	for i := 0; i < b.N; i++ {
		_ = pointer.getPointer()
	}
}
