package fr_FR

import (
	"io"

	"github.com/gol4ng/faker"
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
		"prefix":          func(w io.Writer) (int, error) { return w.Write([]byte(Prefix())) },
		"nameMale":        func(w io.Writer) (int, error) { return w.Write([]byte(NameMale())) },
		"nameFemale":      func(w io.Writer) (int, error) { return w.Write([]byte(NameFemale())) },
		"name":            func(w io.Writer) (int, error) { return w.Write([]byte(Name())) },

		"phoneNumber":                func(w io.Writer) (int, error) { return w.Write([]byte(PhoneNumber())) },
		"phoneNumber07":              func(w io.Writer) (int, error) { return w.Write([]byte(PhoneNumber07())) },
		"phoneNumber07WithSeparator": func(w io.Writer) (int, error) { return w.Write([]byte(PhoneNumber07WithSeparator())) },
		//"phoneNumber08":              func(w io.Writer) (int, error) { return w.Write([]byte(PhoneNumber08())) },
		//"phoneNumber08WithSeparator": func(w io.Writer) (int, error) { return w.Write([]byte(PhoneNumber08WithSeparator())) },
		"mobileNumber":  func(w io.Writer) (int, error) { return w.Write([]byte(MobileNumber())) },
		"serviceNumber": func(w io.Writer) (int, error) { return w.Write([]byte(ServiceNumber())) },

		"secondaryAddress": func(w io.Writer) (int, error) { return w.Write([]byte(SecondaryAddress())) },
		"cityPrefix":       func(w io.Writer) (int, error) { return w.Write([]byte(CityPrefix())) },
		"streetPrefix":     func(w io.Writer) (int, error) { return w.Write([]byte(StreetPrefix())) },
		"buildingNumber":   func(w io.Writer) (int, error) { return w.Write([]byte(BuildingNumber())) },
		"city":             func(w io.Writer) (int, error) { return w.Write([]byte(City())) },
		"streetName":       func(w io.Writer) (int, error) { return w.Write([]byte(StreetName())) },
		"streetAddress":    func(w io.Writer) (int, error) { return w.Write([]byte(StreetAddress())) },
		"postcode":         func(w io.Writer) (int, error) { return w.Write([]byte(Postcode())) },
		"address":          func(w io.Writer) (int, error) { return w.Write([]byte(Address())) },
		"country":          func(w io.Writer) (int, error) { return w.Write([]byte(Country())) },
		"region":           func(w io.Writer) (int, error) { return w.Write([]byte(Region())) },
		"department":       func(w io.Writer) (int, error) { return w.Write([]byte(Department())) },
		"departmentName":   func(w io.Writer) (int, error) { return w.Write([]byte(DepartmentName())) },
		"departmentNumber": func(w io.Writer) (int, error) { return w.Write([]byte(DepartmentNumber())) },
	}
}

func CompileValue(value string) string {
	return faker.CompileValue(value, DataProviders)
}
