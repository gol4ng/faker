package faker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_IntBetween(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		min             int
		max             int
		expectedResults []int
	}{
		{
			name:            "IntBetween(10, 100)",
			min:             1,
			max:             10,
			expectedResults: []int{6, 8, 1, 9, 6, 9, 1, 6, 9, 9, 1, 4},
		},
		{
			name:            "IntBetween(50, 100)",
			min:             50,
			max:             100,
			expectedResults: []int{74, 86, 65, 73, 68, 91, 90, 81, 73, 56, 61, 87},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				i := faker.IntBetween(test.min, test.max)
				//fmt.Printf("%v, ", i)
				assert.Equal(t, expected, i)
			}
		})
		//fmt.Println("")
	}
}

func Test_Int(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []int{5577006791947779410, 8674665223082153551, 6129484611666145821}
	for _, expected := range tests {
		i := faker.Int()
		//fmt.Printf("%v, ", i)
		assert.Equal(t, expected, i)
	}
}

func Test_Intn(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name     string
		value    int
		expected []int
	}{
		{
			name:     "Intn(1)",
			value:    1,
			expected: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Intn(2)",
			value:    2,
			expected: []int{0, 0, 0, 1, 0, 1, 0, 0},
		},
		{
			name:     "Intn(10)",
			value:    10,
			expected: []int{1, 5, 7, 6, 5, 6, 8, 8},
		},
		{
			name:     "Intn(100)",
			value:    100,
			expected: []int{47, 47, 87, 88, 90, 15, 41, 8},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expected {
				i := faker.Intn(test.value)
				//fmt.Printf("%v, ", i)
				assert.Equal(t, expected, i)
			}
		})
		//fmt.Println("")
	}
}

func Test_IntNotZero(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []int{2, 8, 8, 4, 2}
	for _, expected := range tests {
		i := faker.IntNotZero()
		//fmt.Printf("%v, ", i)
		assert.Equal(t, expected, i)
	}
}

func Test_Int32(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []int32{1298498081, 2019727887, 1427131847, 939984059, 911902081}
	for _, expected := range tests {
		i := faker.Int32()
		//fmt.Printf("%v, ", i)
		assert.Equal(t, expected, i)
	}
}

func Test_Intn32(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name     string
		value    int32
		expected []int32
	}{
		{
			name:     "Intn32(1)",
			value:    1,
			expected: []int32{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Intn32(2)",
			value:    2,
			expected: []int32{0, 0, 0, 1, 0, 1, 0, 0},
		},
		{
			name:     "Intn32(10)",
			value:    10,
			expected: []int32{1, 5, 7, 6, 5, 6, 8, 8},
		},
		{
			name:     "Intn32(100)",
			value:    100,
			expected: []int32{47, 47, 87, 88, 90, 15, 41, 8},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expected {
				i := faker.Intn32(test.value)
				//fmt.Printf("%v, ", i)
				assert.Equal(t, expected, i)
			}
		})
		//fmt.Println("")
	}
}

func Test_Uint32(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []uint32{0x9acb0442, 0xf0c5341e, 0xaa209b8e, 0x700e0976, 0x6cb50b02}
	for _, expected := range tests {
		u := faker.Uint32()
		//fmt.Printf("%v, ", u)
		assert.Equal(t, expected, u)
	}
}

func Test_Int64(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []int64{5577006791947779410, 8674665223082153551, 6129484611666145821, 4037200794235010051, 3916589616287113937}
	for _, expected := range tests {
		i := faker.Int64()
		//fmt.Printf("%v, ", i)
		assert.Equal(t, expected, i)
	}
}

func Test_Intn64(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name     string
		value    int64
		expected []int64
	}{
		{
			name:     "Intn64(1)",
			value:    1,
			expected: []int64{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Intn64(2)",
			value:    2,
			expected: []int64{0, 1, 0, 1, 0, 0, 1, 1},
		},
		{
			name:     "Intn64(10)",
			value:    10,
			expected: []int64{8, 1, 0, 1, 3, 6, 1, 7},
		},
		{
			name:     "Intn64(100)",
			value:    100,
			expected: []int64{78, 9, 72, 50, 88, 71, 44, 43},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expected {
				i := faker.Intn64(test.value)
				//fmt.Printf("%v, ", i)
				assert.Equal(t, expected, i)
			}
		})
		//fmt.Println("")
	}
}

func Test_Uint64(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []uint64{0x4d65822107fcfd52, 0x78629a0f5f3f164f, 0xd5104dc76695721d, 0xb80704bb7b4d7c03, 0x365a858149c6e2d1}
	for _, expected := range tests {
		u := faker.Uint64()
		//fmt.Printf("%v, ", u)
		assert.Equal(t, expected, u)
	}
}

func Test_Float32(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []float32{0.6046603, 0.9405091, 0.6645601, 0.4377142, 0.4246375, 0.68682307}
	for _, expected := range tests {
		f := faker.Float32()
		//fmt.Printf("%v, ", f)
		assert.Equal(t, expected, f)
	}
}

func Test_Float32Between(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		min             float32
		max             float32
		expectedResults []float32
	}{
		{
			name:            "Float32Between(0, 1)",
			min:             0,
			max:             1,
			expectedResults: []float32{0.6046603, 0.9405091, 0.6645601, 0.4377142, 0.4246375, 0.68682307, 0.06563702, 0.15651925, 0.09696952, 0.30091187, 0.51521266, 0.81363994},
		},
		{
			name:            "Float32Between(10, 100)",
			min:             1,
			max:             10,
			expectedResults: []float32{2.9283748, 4.425915, 3.8625236, 5.2200084, 3.5473073, 3.6379166, 7.111762, 2.9669776, 2.828682, 4.247843, 6.1360598, 8.762423},
		},
		{
			name:            "Float32Between(50, 100)",
			min:             50,
			max:             100,
			expectedResults: []float32{64.655716, 64.854126, 87.62865, 60.329132, 93.266754, 84.83596, 76.19101, 51.415154, 57.916412, 80.36267, 98.762085, 53.972683},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				f := faker.Float32Between(test.min, test.max)
				//fmt.Printf("%v, ", f)
				assert.Equal(t, expected, f)
				assert.True(t, f < test.max)
				assert.True(t, f >= test.min)
			}
		})
	}
}

func Test_Floatn32(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		value           float32
		expectedResults []float32
	}{
		{
			name:            "Floatn32(0, 1)",
			value:           1,
			expectedResults: []float32{0.6046603, 0.9405091, 0.6645601, 0.4377142, 0.4246375, 0.68682307, 0.06563702, 0.15651925, 0.09696952, 0.30091187, 0.51521266, 0.81363994},
		},
		{
			name:            "Floatn32(10, 100)",
			value:           10,
			expectedResults: []float32{2.1426387, 3.806572, 3.1805816, 4.688898, 2.8303413, 2.9310184, 6.790847, 2.1855304, 2.0318687, 3.608714, 5.7067327, 8.624914},
		},
		{
			name:            "Floatn32(50, 100)",
			value:           100,
			expectedResults: []float32{29.311424, 29.708258, 75.2573, 20.658266, 86.5335, 69.67192, 52.382027, 2.8303082, 15.832828, 60.72534, 97.52416, 7.9453626},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				f := faker.Floatn32(test.value)
				//fmt.Printf("%v, ", f)
				assert.Equal(t, expected, f)
				assert.True(t, f < test.value)
				assert.True(t, f >= 0)
			}
		})
	}
}

func Test_Float64(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []float64{0.6046602879796196, 0.9405090880450124, 0.6645600532184904, 0.4377141871869802, 0.4246374970712657, 0.6868230728671094}
	for _, expected := range tests {
		f := faker.Float64()
		//fmt.Printf("%v, ", f)
		assert.Equal(t, expected, f)
	}
}

func Test_Float64Between(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		min             float64
		max             float64
		expectedResults []float64
	}{
		{
			name:            "Float64Between(0, 1)",
			min:             0,
			max:             1,
			expectedResults: []float64{0.6046602879796196, 0.9405090880450124, 0.6645600532184904, 0.4377141871869802, 0.4246374970712657, 0.6868230728671094, 0.06563701921747622, 0.15651925473279124, 0.09696951891448456, 0.30091186058528707, 0.5152126285020654, 0.8136399609900968},
		},
		{
			name:            "Float64Between(10, 100)",
			min:             1,
			max:             10,
			expectedResults: []float64{2.9283748532413743, 4.425914703697174, 3.8625235689729687, 5.220008604121809, 3.5473073606240066, 3.637916716031342, 7.111762083281946, 2.966977473334879, 2.828681889825906, 4.247842751712154, 6.136059484639203, 8.762422937030976},
		},
		{
			name:            "Float64Between(50, 100)",
			min:             50,
			max:             100,
			expectedResults: []float64{64.6557122276929, 64.85412817781457, 87.62865177758059, 60.32913309568493, 93.26675065007805, 84.83595828733174, 76.19101530250003, 51.4151541662945, 57.91641388725638, 80.36267197727577, 98.76208094302892, 53.9726811686936},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				f := faker.Float64Between(test.min, test.max)
				//fmt.Printf("%v, ", f)
				assert.Equal(t, expected, f)
				assert.True(t, f < test.max)
				assert.True(t, f >= test.min)
			}
		})
	}
}

func Test_Floatn64(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		value           float64
		expectedResults []float64
	}{
		{
			name:            "Floatn64(0, 1)",
			value:           1,
			expectedResults: []float64{0.6046602879796196, 0.9405090880450124, 0.6645600532184904, 0.4377141871869802, 0.4246374970712657, 0.6868230728671094, 0.06563701921747622, 0.15651925473279124, 0.09696951891448456, 0.30091186058528707, 0.5152126285020654, 0.8136399609900968},
		},
		{
			name:            "Floatn64(10, 100)",
			value:           10,
			expectedResults: []float64{2.1426387258237494, 3.80657189299686, 3.1805817433032986, 4.688898449024232, 2.830341511804452, 2.9310185733681577, 6.790846759202163, 2.1855305259276427, 2.0318687664732287, 3.60871416856906, 5.706732760710226, 8.624914374478864},
		},
		{
			name:            "Floatn64(50, 100)",
			value:           100,
			expectedResults: []float64{29.311424455385804, 29.708256355629153, 75.25730355516119, 20.658266191369858, 86.5335013001561, 69.67191657466347, 52.38203060500008, 2.8303083325889995, 15.832827774512765, 60.72534395455153, 97.52416188605784, 7.9453623373871975},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				f := faker.Floatn64(test.value)
				//fmt.Printf("%v, ", f)
				assert.Equal(t, expected, f)
				assert.True(t, f < test.value)
				assert.True(t, f >= 0)
			}
		})
	}
}
