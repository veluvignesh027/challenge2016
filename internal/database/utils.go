package database

import (
	"github.com/RealImage/challenge2016/internal/models"
	"github.com/veluvignesh027/log"
)

func (db *DataStore) AddCityInfo(info models.City) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	db.CitiesStore[info.CityName] = append(db.CitiesStore[info.CityName], info)
	db.ProvincesStore[info.ProvinceName] = append(db.ProvincesStore[info.ProvinceName], info)
	db.CountriesStore[info.CountryName] = append(db.CountriesStore[info.CountryName], info)
}

func (db *DataStore) GetCities(cityName, provinceName, countryName string) ([]models.City, error) {
	db.Mutex.RLock()
	defer db.Mutex.RUnlock()

	var cities []models.City

	if countryName != "" {
		cities = db.CountriesStore[countryName]
		log.Debug("Country:", countryName)
	} else if provinceName != "" {
		cities = db.ProvincesStore[provinceName]
		log.Debug("Province:", provinceName)
	} else if cityName != "" {
		cities = db.CitiesStore[cityName]
		log.Debug("City:", cityName)
	}

	if len(cities) == 0 {
		return []models.City{}, nil
	}

	var res []models.City
	for _, city := range cities {
		if (provinceName == "" || city.ProvinceName == provinceName) &&
			(countryName == "" || city.CountryName == countryName) &&
			(cityName == "" || city.CityName == cityName) {
			res = append(res, city)
		}
	}

	return res, nil
}

func (db *DataStore) DeleteCityInfo(info models.City) bool {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	flag := false
	// Remove the city from CitiesStore
	if cities, exists := db.CitiesStore[info.CityName]; exists {
		for i, city := range cities {
			if city == info {
				// Remove the city from the slice
				db.CitiesStore[info.CityName] = append(cities[:i], cities[i+1:]...)
				flag = true
				break
			}
		}
	}
	return flag == true
}

func (db *DataStore) UpdateCityInfo(oldInfo models.City, newInfo models.City) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	// Remove the old city information from CitiesStore
	if cities, exists := db.CitiesStore[oldInfo.CityName]; exists {
		for i, city := range cities {
			if city == oldInfo {
				// Remove the old city from the slice
				db.CitiesStore[oldInfo.CityName] = append(cities[:i], cities[i+1:]...)
				break
			}
		}
	}

	// Remove the old city information from ProvincesStore
	if provinces, exists := db.ProvincesStore[oldInfo.ProvinceName]; exists {
		for i, city := range provinces {
			if city == oldInfo {
				// Remove the old city from the slice
				db.ProvincesStore[oldInfo.ProvinceName] = append(provinces[:i], provinces[i+1:]...)
				break
			}
		}
	}

	// Remove the old city information from CountriesStore
	if countries, exists := db.CountriesStore[oldInfo.CountryName]; exists {
		for i, city := range countries {
			if city == oldInfo {
				// Remove the old city from the slice
				db.CountriesStore[oldInfo.CountryName] = append(countries[:i], countries[i+1:]...)
				break
			}
		}
	}

	// Add the new city information to CitiesStore
	db.CitiesStore[newInfo.CityName] = append(db.CitiesStore[newInfo.CityName], newInfo)

	// Add the new city information to ProvincesStore
	db.ProvincesStore[newInfo.ProvinceName] = append(db.ProvincesStore[newInfo.ProvinceName], newInfo)

	// Add the new city information to CountriesStore
	db.CountriesStore[newInfo.CountryName] = append(db.CountriesStore[newInfo.CountryName], newInfo)
}
