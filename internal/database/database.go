package database

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/RealImage/challenge2016/internal/models"
)

var (
	ErrCityNotFound = fmt.Errorf("city not found")
)

var (
	datastore *DataStore
	once      sync.Once
)

type DataStore struct {
	Mutex          sync.RWMutex
	CitiesStore    map[string][]models.City
	ProvincesStore map[string][]models.City
	CountriesStore map[string][]models.City
	Distributors   map[string]*models.Distributor
}

func NewDataStore() *DataStore {
	once.Do(func() {
		datastore = &DataStore{
			Mutex:          sync.RWMutex{},
			CitiesStore:    make(map[string][]models.City),
			ProvincesStore: make(map[string][]models.City),
			CountriesStore: make(map[string][]models.City),
			Distributors:   make(map[string]*models.Distributor),
		}
	})
	return datastore
}

func LoadCSVFile() error {
	csvfile, err := os.OpenFile(os.Getenv("CSVFILE"), os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var cities []models.City
	for i, record := range records {
		if i == 0 {
			continue
		}

		city := models.City{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}
		cities = append(cities, city)
		NewDataStore().AddCityInfo(city)
	}
	return nil
}
