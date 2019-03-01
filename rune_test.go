package faker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_RuneBetween(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		min             rune
		max             rune
		expectedResults []rune
	}{
		{
			name:            "RuneBetween('a','g')",
			min:             'a',
			max:             'g',
			expectedResults: []rune{'c', 'b', 'd', 'f', 'f', 'c', 'a'},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				r := faker.RuneBetween(test.min, test.max)
				//fmt.Printf("'%s', ", string(r))
				assert.Equal(t, expected, r)
			}
		})
		//fmt.Println("")
	}
}

func Test_RunesIn(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		length          int
		values          string
		expectedResults []string
	}{
		{
			name:            "RunesIn(5,[]rune(\"aB6\"))",
			length:          5,
			values:          "aB6",
			expectedResults: []string{"6a66B", "aB6Ba", "6Ba6B", "6a666", "6a6Ba", "BBBaa", "Ba66a"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				r := faker.RunesIn(test.length, []rune(test.values))
				//fmt.Printf("\"'%s\", ", string(r))
				assert.Equal(t, []rune(expected), r)
			}
		})
		//fmt.Println("")
	}
}

func Test_RuneIn(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		values          string
		expectedResults []rune
	}{
		{
			name:            "RuneIn([]rune(\"aB6\"))",
			values:          "aB6",
			expectedResults: []rune{'6', 'a', '6', '6', 'B', 'a', 'B'},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				r := faker.RuneIn([]rune(test.values))
				//fmt.Printf("'%s', ", string(r))
				assert.Equal(t, expected, r)
			}
		})
		//fmt.Println("")
	}
}
