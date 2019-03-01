package faker_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_CompileValue(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		value           string
		dataProviders   *faker.DataProvider
		expectedResults []string
	}{
		{
			name:  "foo bar baz",
			value: "{{foo}}",
			dataProviders: &faker.DataProvider{
				"foo": func(w io.Writer) (int, error) { return w.Write([]byte("foo_value")) },
				"bar": func(w io.Writer) (int, error) { return w.Write([]byte("bar_value")) },
				"baz": func(w io.Writer) (int, error) { return w.Write([]byte("baz_value")) },
			},
			expectedResults: []string{"foo_value"},
		},
		{
			name:  "numerify lexify asciify",
			value: "{{## %% ?? **}} {{foo}}",
			dataProviders: &faker.DataProvider{
				"foo": func(w io.Writer) (int, error) { return w.Write([]byte("foo_value")) },
			},
			expectedResults: []string{
				"57 64 mu 3n foo_value",
				"88 58 yl Pb foo_value",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				assert.Equal(t, expected, faker.CompileValue(test.value, test.dataProviders))
			}
		})
	}
}

func Test_CompileRandValue(t *testing.T) {
	faker.FakerRand.Seed(1)

	tests := []struct {
		name            string
		values          []string
		dataProviders   *faker.DataProvider
		expectedResults []string
	}{
		{
			name:   "foo bar baz",
			values: []string{"{{foo}}", "{{bar}}", "{{baz}}"},
			dataProviders: &faker.DataProvider{
				"foo": func(w io.Writer) (int, error) { return w.Write([]byte("foo_value")) },
				"bar": func(w io.Writer) (int, error) { return w.Write([]byte("bar_value")) },
				"baz": func(w io.Writer) (int, error) { return w.Write([]byte("baz_value")) },
			},
			expectedResults: []string{"baz_value", "foo_value", "baz_value"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, expected := range test.expectedResults {
				assert.Equal(t, expected, faker.CompileRandValue(test.values, test.dataProviders))
			}
		})
	}
}

func Test_Compute(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []struct {
		value    []byte
		expected []string
	}{
		{
			value:    []byte("## %% ?? **"),
			expected: []string{"57 64 mu 3n", "88 58 yl Pb", "81 74 xg f4", "16 13 nv f/"},
		},
		{
			value:    []byte("#"),
			expected: []string{"3", "5", "2", "1"},
		},
		{
			value:    []byte("#####"),
			expected: []string{"18637", "10766", "34137", "04466"},
		},
		{
			value:    []byte("foo ### baz"),
			expected: []string{"foo 843 baz", "foo 773 baz", "foo 133 baz", "foo 442 baz"},
		},
		{
			value:    []byte("?"),
			expected: []string{"t", "a", "h", "x"},
		},
		{
			value:    []byte("?????"),
			expected: []string{"mcfhs", "btstq", "nvlfs", "fqdgr"},
		},
		{
			value:    []byte("foo ??? baz"),
			expected: []string{"foo gsh baz", "foo oef baz", "foo avk baz", "foo xqf baz"},
		},
		{
			value:    []byte("*"),
			expected: []string{"'", "@", "j", "X"},
		},
		{
			value:    []byte("*****"),
			expected: []string{"-\\R=D", "Cx7dI", "GhDo_", "efZ4%"},
		},
		{
			value:    []byte("foo *** baz"),
			expected: []string{"foo L=6 baz", "foo ?dh baz", "foo JH^ baz", "foo &dU baz"},
		},
		{
			value:    []byte("foo"),
			expected: []string{"foo", "foo"},
		},
	}
	for _, test := range tests {
		t.Run(string(test.value), func(t *testing.T) {
			for _, expected := range test.expected {
				assert.Equal(t, expected, faker.Compute(test.value))
			}
		})
	}
}

func Test_Numerify(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []struct {
		value    []byte
		expected []string
	}{
		{
			value:    []byte("## %% ?? **"),
			expected: []string{"57 64 ?? **", "58 75 ?? **", "88 58 ?? **", "76 42 ?? **"},
		},
		{
			value:    []byte("#"),
			expected: []string{"8", "1", "5", "6"},
		},
		{
			value:    []byte("#####"),
			expected: []string{"82011", "63715", "38352", "11863"},
		},
		{
			value:    []byte("foo ### baz"),
			expected: []string{"foo 710 baz", "foo 766 baz", "foo 341 baz", "foo 370 baz"},
		},
		{
			value:    []byte("%"),
			expected: []string{"6", "1", "2", "4"},
		},
		{
			value:    []byte("%%%%%"),
			expected: []string{"62377", "61316", "68672", "41782"},
		},
		{
			value:    []byte("foo %%% baz"),
			expected: []string{"foo 464 baz", "foo 372 baz", "foo 568 baz", "foo 188 baz"},
		},
	}
	for _, test := range tests {
		t.Run(string(test.value), func(t *testing.T) {
			for _, expected := range test.expected {
				assert.Equal(t, expected, faker.Numerify(test.value))
			}
		})
	}
}

func Test_Lexify(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []struct {
		value    []byte
		expected []string
	}{
		{
			value:    []byte("## %% ?? **"),
			expected: []string{"## %% kb **", "## %% vb **", "## %% mu **", "## %% ix **"},
		},
		{
			value:    []byte("?"),
			expected: []string{"q", "y", "j", "m"},
		},
		{
			value:    []byte("?????"),
			expected: []string{"ylpxs", "qpgxg", "lmdjw", "anvts"},
		},
		{
			value:    []byte("foo ??? baz"),
			expected: []string{"foo xjd baz", "foo oih baz", "foo hfp baz", "foo quv baz"},
		},
	}
	for _, test := range tests {
		t.Run(string(test.value), func(t *testing.T) {
			for _, expected := range test.expected {
				assert.Equal(t, expected, faker.Lexify(test.value))
			}
		})
	}
}

func Test_Asciify(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []struct {
		value    []byte
		expected []string
	}{
		{
			value:    []byte("## %% ?? **"),
			expected: []string{"## %% ?? hg", "## %% ?? -G", "## %% ?? &&", "## %% ?? 3n"},
		},
		{
			value:    []byte("*"),
			expected: []string{"h", "#", "N", "!"},
		},
		{
			value:    []byte("*****"),
			expected: []string{"|HPb)", "%qZG5", "f4FoQ", "|j;f/"},
		},
		{
			value:    []byte("foo *** baz"),
			expected: []string{"foo E,b baz", "foo .@Y baz", "foo H.[ baz", "foo HRZ baz"},
		},
	}
	for _, test := range tests {
		t.Run(string(test.value), func(t *testing.T) {
			for _, expected := range test.expected {
				assert.Equal(t, expected, faker.Asciify(test.value))
			}
		})
	}
}
