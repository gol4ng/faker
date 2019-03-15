package faker_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_StringBool(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"true", "true", "true", "true", "true", "false"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.StringBool())
	}
}

func Test_Number(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		n               int
		expectedResults []string
	}{
		{
			name:            "Number(1)",
			n:               1,
			expectedResults: []string{"5", "7", "0", "8", "5", "8"},
		},
		{
			name:            "Number(10)",
			n:               10,
			expectedResults: []string{"75880376828", "45682011637", "55383521186", "57107663413", "20446684377", "61334420417"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				n := faker.Number(test.n)
				//fmt.Printf("\"%s\", ", n)
				assert.Equal(t, expected, n)
			}
		})
		//fmt.Println("")
	}
}

func Test_Numeric(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"5", "7", "0", "8", "5", "8", "0"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Numeric())
	}
}

func Test_Numerics(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		n               int
		expectedResults []string
	}{
		{
			name:            "Numerics(1)",
			n:               1,
			expectedResults: []string{"5", "7", "0", "8", "5", "8"},
		},
		{
			name:            "Numerics(10)",
			n:               10,
			expectedResults: []string{"0588037682", "8156820116", "3715383521", "1863710766", "3413704466", "8437731334"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				n := faker.Numerics(test.n)
				//fmt.Printf("\"%s\", ", n)
				assert.Equal(t, expected, n)
			}
		})
		//fmt.Println("")
	}
}

func Test_NumericNotZero(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"3", "8", "6", "4", "2", "1", "7"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.NumericNotZero())
	}
}

func Test_NumericsNotZero(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		n               int
		expectedResults []string
	}{
		{
			name:            "NumericsNotZero(1)",
			n:               1,
			expectedResults: []string{"3", "8", "6", "4", "2", "1"},
		},
		{
			name:            "NumericsNotZero(10)",
			n:               10,
			expectedResults: []string{"7512587542", "1474258636", "1354148444", "4455455854", "8776226124", "6237761316"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				n := faker.NumericsNotZero(test.n)
				//fmt.Printf("\"%s\", ", n)
				assert.Equal(t, expected, n)
			}
		})
		fmt.Println("")
	}
}

func Test_Letter(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"k", "b", "v", "b", "m", "u", "i"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Letter())
	}
}

func Test_Letters(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		n               int
		expectedResults []string
	}{
		{
			name:            "Letters(1)",
			n:               1,
			expectedResults: []string{"X", "V", "l", "B", "z", "g"},
		},
		{
			name:            "Letters(10)",
			n:               10,
			expectedResults: []string{"baiCMRAjWw", "hTHctcuAxh", "xKQFDaFpLS", "jFbcXoEFfR", "sWxPLDnJOb", "CsNVlgTeMa"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				l := faker.Letters(test.n)
				//fmt.Printf("\"%s\", ", l)
				assert.Equal(t, expected, l)
			}
		})
		fmt.Println("")
	}
}

func Test_LettersIn(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		n               int
		values          string
		expectedResults []string
	}{
		{
			name:            "LettersIn(1, \"abhslds\")",
			n:               1,
			values:          "abhslds",
			expectedResults: []string{"s", "h", "b", "a", "s", "l"},
		},
		{
			name:            "LettersIn(10, \"XhhuTSkuxK\")",
			n:               10,
			values:          "XhhuTSkuxK",
			expectedResults: []string{"SXkXThhKxT", "hSukSkxxuu", "uxXShxuhKk", "uhSkuXTuuu", "xTKuuhKKXS", "xxuShXSkkx"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				l := faker.LettersIn(test.n, test.values)
				//fmt.Printf("\"%s\", ", l)
				assert.Equal(t, expected, l)
			}
		})
		//fmt.Println("")
	}
}

func Test_Alphanumeric(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"B", "p", "L", "n", "f", "g", "D"}
	for _, expected := range tests {
		a := faker.Alphanumeric()
		//fmt.Printf("\"%v\", ", a)
		assert.Equal(t, expected, a)
	}
}

func Test_Alphanumerics(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		n               int
		expectedResults []string
	}{
		{
			name:            "Alphanumerics(1)",
			n:               1,
			expectedResults: []string{"B", "p", "L", "n", "f", "g"},
		},
		{
			name:            "Alphanumerics(10)",
			n:               10,
			expectedResults: []string{"Dsc2WD8F2q", "NfHK5a84jj", "JkwzDkh9h2", "fhfUVuS9jZ", "8uVbhV3vC5", "AWX39IVUWS"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				l := faker.Alphanumerics(test.n)
				//fmt.Printf("\"%s\", ", l)
				assert.Equal(t, expected, l)
			}
		})
		//fmt.Println("")
	}
}

func Test_Ascii(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"h", "g", "-", "G", "&", "&", "3"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Ascii())
	}
}

func Test_StringIn(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		values          []string
		expectedResults []string
	}{
		{
			name:            "StringIn([]string{\"foo\", \"bar\", \"baz\"})",
			values:          []string{"foo", "bar", "baz"},
			expectedResults: []string{"baz", "foo", "baz"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				s := faker.StringIn(test.values)
				//fmt.Printf("\"%s\", ", s)
				assert.Equal(t, expected, s)
			}
		})
		//fmt.Println("")
	}
}
