package faker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_E164PhoneNumber(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"+7570858058803",
		"+6682815682011",
		"+5371538352118",
		"+2371076634137",
		"+1446684377313",
		"+6442041705786",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.E164PhoneNumber())
	}
}

func Test_Imei(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"570858058803769",
		"828156820116374",
		"153835211863711",
		"76634137044660",
		"843773133442049",
		"170578655785524",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Imei())
	}
}
