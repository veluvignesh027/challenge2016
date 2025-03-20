package server

import "errors"

// Distributor errors
var (
	ErrDistributorNotFound = errors.New("distributor not found")
	ErrInvalidDistributor  = errors.New("invalid distributor")
	ErrCreatingDistributor = errors.New("error creating a new distributor")
	ErrDeletingDistributor = errors.New("error deleting a distributor")
	ErrUpdatingDistributor = errors.New("error updating a distributor")
)

var (
	ErrCityNotFound = errors.New("city not found")
)
var (
	AlertHTML          = "<script>alert({{.}})</script>"
	ErrLoadingCSV      = errors.New("error loading CSV")
	ErrCSVFileNotFound = errors.New("CSV file not found")
)
