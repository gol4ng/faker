package faker_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_DataProvider_CompileValue(t *testing.T) {
	faker.FakerRand.Seed(1)
	dataProvider := &faker.DataProvider{
		"foo": func(w io.Writer) (int, error) { return w.Write([]byte("foo_value")) },
		"bar": func(w io.Writer) (int, error) { return w.Write([]byte("bar_value")) },
		"baz": func(w io.Writer) (int, error) { return w.Write([]byte("baz_value")) },
	}

	assert.Equal(t, "foo_value bar_value baz_value", dataProvider.CompileValue("{{foo}} {{bar}} {{baz}}"))
}

func Test_DataProviderRegistry_Get_WillPanic(t *testing.T) {
	faker.FakerRand.Seed(1)
	dataProviderRegistry := faker.NewDataProviderRegistry()

	assert.PanicsWithValue(
		t,
		"DataProvider not found",
		func() { dataProviderRegistry.Get("A") },
	)
}

func Test_DataProviderRegistry_Get(t *testing.T) {
	faker.FakerRand.Seed(1)
	dataProviderA := &faker.DataProvider{"foo": func(w io.Writer) (int, error) { return w.Write([]byte("foo_value")) }}
	dataProviderB := &faker.DataProvider{"bar": func(w io.Writer) (int, error) { return w.Write([]byte("bar_value")) }}
	dataProviderRegistry := faker.NewDataProviderRegistry().Add("A", dataProviderA).Add("B", dataProviderB)

	assert.Equal(t, dataProviderA, dataProviderRegistry.Get("A"))
	assert.Equal(t, dataProviderB, dataProviderRegistry.Get("B"))
}

func Test_DataProviderRegistry_Rand(t *testing.T) {
	faker.FakerRand.Seed(1)
	dataProviderA := &faker.DataProvider{"foo": func(w io.Writer) (int, error) { return w.Write([]byte("foo_value")) }}
	dataProviderB := &faker.DataProvider{"bar": func(w io.Writer) (int, error) { return w.Write([]byte("bar_value")) }}
	dataProviderRegistry := faker.NewDataProviderRegistry().Add("A", dataProviderA).Add("B", dataProviderB)

	expectedDataProviders := []*faker.DataProvider{
		dataProviderB,
		dataProviderB,
		dataProviderB,
		dataProviderB,
		dataProviderB,
		dataProviderA,
	}

	for _, expected := range expectedDataProviders {
		assert.Equal(t, expected, dataProviderRegistry.Rand())
	}
}
