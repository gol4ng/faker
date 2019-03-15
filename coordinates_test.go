package faker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_Latitude(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []string{"18.838852", "79.291641", "29.620811", "-11.211449", "-13.565254"}
	for _, expected := range tests {
		lat := faker.Latitude()
		//fmt.Printf("\"%v\", ", lat)
		assert.Equal(t, expected, lat)
	}
}

func Test_Longitude(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []string{"37.677704", "158.583282", "59.241623", "-22.422897", "-27.130508"}
	for _, expected := range tests {
		lon := faker.Longitude()
		//fmt.Printf("\"%v\", ", lon)
		assert.Equal(t, expected, lon)
	}
}

func Test_Coordinates(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []string{"18.838852 158.583282", "29.620811 -22.422897", "-13.565254 67.256302", "-78.185333 -123.653069", "-72.545486 -71.671722"}
	for _, expected := range tests {
		c := faker.Coordinates()
		//fmt.Printf("\"%v\", ", c)
		assert.Equal(t, expected, c)
	}
}
