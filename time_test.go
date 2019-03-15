package faker_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_TimeBetween(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		min             time.Time
		max             time.Time
		expectedResults []int64
	}{
		{
			name:            "TimeBetween(time.Unix(10, 0), time.Unix(15, 0))",
			min:             time.Unix(10, 0),
			max:             time.Unix(15, 0),
			expectedResults: []int64{11947779410, 13082153551, 11666145821, 14235010051, 11287113937, 14549167320, 12632969758, 12331776148},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				i := faker.TimeBetween(test.min, test.max)
				//fmt.Printf("%v, ", i.UnixNano())
				assert.Equal(t, expected, i.UnixNano())
				assert.True(t, i.After(test.min))
				assert.True(t, i.Before(test.max))
			}
		})
		//fmt.Println("")
	}
}

func Test_TimePast(t *testing.T) {
	for i := 0; i <= 10; i++ {
		assert.True(t, faker.TimePast().Before(time.Now()))
	}
}

func Test_TimeFuture(t *testing.T) {
	for i := 0; i <= 10; i++ {
		assert.True(t, faker.TimeFuture().After(time.Now()))
	}
}
