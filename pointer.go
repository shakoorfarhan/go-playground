package main

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	zeroval(i)
	println("zeroval:", i)

	zeroptr(&i)
	println("zeroptr:", i)
}
