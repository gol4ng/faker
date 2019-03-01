package provider

import (
	"github.com/gol4ng/faker"
	"github.com/gol4ng/faker/provider/en_GB"
	"github.com/gol4ng/faker/provider/fr_FR"
)

var Registry = faker.NewDataProviderRegistry().
	Add("en_GB", en_GB.GetDataProviders()).
	Add("fr_FR", fr_FR.GetDataProviders())

func CompileValue(value string) string {
	return Rand().CompileValue(value)
}

func Get(name string) *faker.DataProvider {
	return Registry.Get(name)
}

func Rand() *faker.DataProvider {
	return Registry.Rand()
}
