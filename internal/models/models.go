package models

type Distributor struct {
	Name              string
	IncludedCountries map[string]*Country
	ExcludedCountries map[string]*Country
	IncludedProvinces map[string]*Province
	ExcludedProvinces map[string]*Province
	IncludedCites     map[string]*City
	ExcludedCities    map[string]*City
}

type City struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

type Province struct {
	ProvinceCode string
	ProvinceName string
	Cities       []City
}

type Country struct {
	CountryCode string
	CountryName string
	Cities      []City
	Provinces   []Province
}

func NewDistributor(name string) *Distributor {
	return &Distributor{
		Name:              name,
		IncludedCountries: make(map[string]*Country),
		ExcludedCountries: make(map[string]*Country),
		IncludedProvinces: make(map[string]*Province),
		ExcludedProvinces: make(map[string]*Province),
		IncludedCites:     make(map[string]*City),
		ExcludedCities:    make(map[string]*City),
	}
}
