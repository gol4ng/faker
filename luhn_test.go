package faker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_Luhn(t *testing.T) {
	tests := map[int]int{
		1111: 11114,
		1234: 12344,
		7826487624827: 78264876248278,
	}
	for value, expected := range tests {
		assert.Equal(t, expected, faker.Luhn(value))
	}
}
