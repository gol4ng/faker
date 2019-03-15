package faker

import (
	"bytes"
	"io"
	"math/rand"

	"github.com/valyala/fasttemplate"
)

const starTag = "{{"
const endTag = "}}"

var FakerRand = rand.New(NewLockedSource(rand.NewSource(1)))

// TODO REFACTO the compile value process
func CompileValue(value string, dataProviders *DataProvider) string {
	return fasttemplate.ExecuteFuncString(
		value,
		starTag,
		endTag,
		func(w io.Writer, tag string) (int, error) {
			t := []byte(tag)
			find := false
			find = numerify(t)
			find = lexify(t)
			find = asciify(t)
			if find {
				return w.Write(t)
			}
			if dataProvider, ok := (*GlobalFunc)[tag]; ok {
				return dataProvider(w)
			}
			if dataProvider, ok := (*dataProviders)[tag]; ok {
				return dataProvider(w)
			}
			return w.Write(t)
		},
	)
}

func CompileRandValue(values []string, dataProviders *DataProvider) string {
	return CompileValue(StringIn(values), dataProviders)
}

func Compute(b []byte) string {
	a := append([]byte{}, b...)
	numerify(a)
	lexify(a)
	asciify(a)

	return string(a)
}

func Numerify(b []byte) string {
	a := append([]byte{}, b...)
	numerify(a)
	return string(a)
}
func Lexify(b []byte) string {
	a := append([]byte{}, b...)
	lexify(a)
	return string(a)
}
func Asciify(b []byte) string {
	a := append([]byte{}, b...)
	asciify(a)
	return string(a)
}

func numerify(b []byte) bool {
	find := false
	for {
		if i := bytes.Index(b, []byte("#")); i != -1 {
			find = true
			copy(b[i:], Numeric())
			continue
		}
		break
	}
	for {
		if i := bytes.Index(b, []byte("%")); i != -1 {
			find = true
			copy(b[i:], NumericNotZero())
			continue
		}
		break
	}
	return find
}

func lexify(b []byte) bool {
	find := false
	for {
		if i := bytes.Index(b, []byte("?")); i != -1 {
			find = true
			copy(b[i:], Letter())
			continue
		}
		break
	}
	return find
}

func asciify(b []byte) bool {
	find := false
	for {
		if i := bytes.Index(b, []byte("*")); i != -1 {
			find = true
			copy(b[i:], Ascii())
			continue
		}
		break
	}
	return find
}
