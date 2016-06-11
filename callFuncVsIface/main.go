package main

type digiter interface {
	getPointer() *int
	getValue() int
}

type D struct {
	n int
}

func (d D) getPointer() *int {
	n := d.n
	return &n
}

func (d D) getValue() int {
	n := d.n
	return n
}

func getValue(d D) int {

	n := d.n
	return n
}

func getPointer(d D) *int {
	n := d.n
	return &n
}

func main() {

}
