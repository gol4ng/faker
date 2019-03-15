package faker

import (
	"errors"
	"io"
)

var GlobalFunc = &DataProvider{
	"randomInt":        func(w io.Writer) (int, error) { return w.Write([]byte(string(Int()))) },
	"randomIntNotZero": func(w io.Writer) (int, error) { return w.Write([]byte(string(IntNotZero()))) },
	"randomInt32":      func(w io.Writer) (int, error) { return w.Write([]byte(string(Int32()))) },
	"randomUint32":     func(w io.Writer) (int, error) { return w.Write([]byte(string(Uint32()))) },
	"randomInt64":      func(w io.Writer) (int, error) { return w.Write([]byte(string(Int64()))) },
	"randomUint64":     func(w io.Writer) (int, error) { return w.Write([]byte(string(Uint64()))) },

	"randomNumeric": func(w io.Writer) (int, error) { return w.Write([]byte(Numeric())) },
	"randomLetter":  func(w io.Writer) (int, error) { return w.Write([]byte(Letter())) },
	"randomAscii":   func(w io.Writer) (int, error) { return w.Write([]byte(Ascii())) },

	"e164PhoneNumber": func(w io.Writer) (int, error) { return w.Write([]byte(E164PhoneNumber())) },
	"imei":            func(w io.Writer) (int, error) { return w.Write([]byte(Imei())) },

	"latitude":         func(w io.Writer) (int, error) { return w.Write([]byte(Latitude())) },
	"longitude":        func(w io.Writer) (int, error) { return w.Write([]byte(Longitude())) },
	"coordinates":      func(w io.Writer) (int, error) { return w.Write([]byte(Coordinates())) },
}

type DataProvider map[string]func(w io.Writer) (int, error)

func (d *DataProvider) Write(value string, w io.Writer) (int, error) {
	if dataProvider, ok := (*d)[value]; ok {
		return dataProvider(w)
	}
	return 0, errors.New("DataProvider not found")
}

func (d *DataProvider) CompileValue(value string) string {
	return CompileValue(value, d)
}

type DataProviderRegistry struct {
	providers     map[string]*DataProvider
	providersKeys []string
}

func (r *DataProviderRegistry) Get(name string) *DataProvider {
	if provider, ok := r.providers[name]; ok {
		return provider
	}
	panic("DataProvider not found")
}

func (r *DataProviderRegistry) Add(name string, provider *DataProvider) *DataProviderRegistry {
	r.providers[name] = provider
	r.providersKeys = append(r.providersKeys, name)
	return r
}

func (r *DataProviderRegistry) Rand() *DataProvider {
	return r.providers[StringIn(r.providersKeys)]
}

func NewDataProviderRegistry() *DataProviderRegistry {
	return &DataProviderRegistry{providers: map[string]*DataProvider{}, providersKeys: []string{}}
}
