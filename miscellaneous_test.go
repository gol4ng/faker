package faker_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gol4ng/faker"
)

func Test_LanguageCode(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"sm", "zu", "hz", "af", "mn", "ki", "oc"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.LanguageCode())
	}
}

func Test_CountryCode(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"PS", "WS", "CU", "HM", "GR", "BN", "AO"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.CountryCode())
	}
}

func Test_CurrencyCode(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"XPF", "LAK", "DKK", "FJD", "DJF", "KES", "ILS"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.CurrencyCode())
	}
}

func Test_ISOAlpha3(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"PRT", "WLF", "COM", "GUY", "GNQ", "BLM", "ARE"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.ISOAlpha3())
	}
}

func Test_Local(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"es_HN", "ii_CN", "ha_NG", "zu_ZA", "ku_IR", "ku_TR", "gez_ET"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Local())
	}
}

func Test_LocalDash(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"es-HN", "ii-CN", "ha-NG", "zu-ZA", "ku-IR", "ku-TR", "gez-ET"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.LocalDash())
	}
}

func Test_Emoji(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"😣", "😤", "😤", "😁", "😑", "😭", "😃"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Emoji())
	}
}

func Test_Md5(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"85f4da88848c5f0edef581d374c38fd4", "9e43e9e992088aa20c05bff425c971ec", "2ab01702ea8ff904c9d410e0086a41aa", "aac64d47a5359edc512ffbd7574d090e"}
	for _, expected := range tests {
		h := faker.Md5()
		//fmt.Printf("\"%v\", ", h)
		assert.Equal(t, expected, h)
	}
}

func Test_Sha1(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"6b45b07116bcb511f33408828a959b4ace5dde11", "534e5f731a1ba17556a8e388b8fff9d1123ef0e8", "d788cf8077d7141868ec3e2db796842052facd76"}
	for _, expected := range tests {
		h := faker.Sha1()
		//fmt.Printf("\"%v\", ", h)
		assert.Equal(t, expected, h)
	}
}

func Test_Sha256(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"f4b757126210cc299faa1c7aba3f684f52fccb9a8c5be06baec3517cd7b80868", "d9c914cb166d332e1863d72a8879244e1489e59e390d50e8c4753110dac9c6d6"}
	for _, expected := range tests {
		h := faker.Sha256()
		//fmt.Printf("\"%v\", ", h)
		assert.Equal(t, expected, h)
	}
}

func Test_Sha512(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"7b9a55dde97003cf5dbc1e56375c685b17cc4776ce8ff9cad7068d025d41e72dcbf17af23f9c714b0072046a6687a2917176638386d0eb54d730f0438986ce15",
		"386a4f2df84d9a7c32edae1476ab8931745059319a7532fa74b68947365dd805c93c015396a6ee5a2dcaf6cc84029523f3bf633b1f4a980b063a75e6996d5d02",
	}
	for _, expected := range tests {
		h := faker.Sha512()
		//fmt.Printf("\"%v\", ", h)
		assert.Equal(t, expected, h)
	}
}

func Test_Boolean(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []bool{true, false, false, false}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Bool())
	}
}

func Test_BooleanRatio(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []bool{true, false, false, false}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.BoolRatio(10))
	}
}
