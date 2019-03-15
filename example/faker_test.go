package main

import (
	"io"
	"testing"

	"github.com/gol4ng/faker"
	"github.com/gol4ng/faker/provider"
	"github.com/gol4ng/faker/provider/fr_FR"
)

var value = "{{titleMale}}"

var benchmarkProvider = &faker.DataProvider{
	"titleMale": func(w io.Writer) (int, error) { return w.Write([]byte("benchmark")) },
}

func Benchmark_Faker_CompileValue(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		faker.CompileValue(value, benchmarkProvider)
	}
}

func Benchmark_DataProvider_CompileValue(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		benchmarkProvider.CompileValue(value)
	}
}

func Benchmark_Provider_CompileValue(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		provider.CompileValue(value)
	}
}

func Benchmark_fr_FR_Provider_CompileValue(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		fr_FR.CompileValue(value)
	}
}
func Benchmark_RandString(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		faker.Letters(10)
	}
}
