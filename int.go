package faker

// Return an int between [a,b]
func IntBetween(a int, b int) int {
	min := a
	max := b
	if a > b {
		min = b
		max = a
	}
	return int(Intn64(int64(max-min))) + min
}

// Return an int between [0-9]
func Int() int {
	return FakerRand.Int()
}

//http://www.pcg-random.org/posts/bounded-rands.html
//https://github.com/golang/go/issues/21835
// Return an int between [0-n)
func Intn(n int) int {
	return FakerRand.Intn(n)
}

// Return an int between [1-9]
func IntNotZero() int {
	return int(Intn32(8) + 1)
}

// Return an int32
func Int32() int32 {
	return FakerRand.Int31()
}

// Return an int32 between [0-n)
func Intn32(n int32) int32 {
	return FakerRand.Int31n(n)
}

// Return an uint32
func Uint32() uint32 {
	return FakerRand.Uint32()
}

// Return an int64
func Int64() int64 {
	return FakerRand.Int63()
}

// Return an int64 between [0-n)
func Intn64(n int64) int64 {
	return FakerRand.Int63n(n)
}

// Return an uint64
func Uint64() uint64 {
	return FakerRand.Uint64()
}

func Float32() float32 {
	return FakerRand.Float32()
}

func Float32Between(a float32, b float32) float32 {
	min := a
	max := b
	if a > b {
		min = b
		max = a
	}
	return min + Float32()*(max-min)
}

func Floatn32(n float32) float32 {
	return Float32Between(0, n)
}

func Float64() float64 {
	return FakerRand.Float64()
}

func Float64Between(a float64, b float64) float64 {
	min := a
	max := b
	if a > b {
		min = b
		max = a
	}
	return min + Float64()*(max-min)
}

func Floatn64(n float64) float64 {
	return Float64Between(0, n)
}
