package en_GB

import (
	"github.com/gol4ng/faker"
	"io"
)

var DataProviders = GetDataProviders()

func GetDataProviders() *faker.DataProvider {
	return &faker.DataProvider{
		"titleMale":       func(w io.Writer) (int, error) { return w.Write([]byte(TitleMale())) },
		"titleFemale":     func(w io.Writer) (int, error) { return w.Write([]byte(TitleFemale())) },
		"title":           func(w io.Writer) (int, error) { return w.Write([]byte(Title())) },
		"firstNameMale":   func(w io.Writer) (int, error) { return w.Write([]byte(FirstNameMale())) },
		"firstNameFemale": func(w io.Writer) (int, error) { return w.Write([]byte(FirstNameFemale())) },
		"firstName":       func(w io.Writer) (int, error) { return w.Write([]byte(FirstName())) },
		"lastName":        func(w io.Writer) (int, error) { return w.Write([]byte(LastName())) },
		"nameMale":        func(w io.Writer) (int, error) { return w.Write([]byte(NameMale())) },
		"nameFemale":      func(w io.Writer) (int, error) { return w.Write([]byte(NameFemale())) },
		"name":            func(w io.Writer) (int, error) { return w.Write([]byte(Name())) },

		"phoneNumber":  func(w io.Writer) (int, error) { return w.Write([]byte(PhoneNumber())) },
		"mobileNumber": func(w io.Writer) (int, error) { return w.Write([]byte(MobileNumber())) },
	}
}

func CompileValue(value string) string {
	return faker.CompileValue(value, DataProviders)
}
