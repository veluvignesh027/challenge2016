package server

import (
	"net/http"

	"github.com/RealImage/challenge2016/internal/models"
	"github.com/gin-gonic/gin"
)

// Permission Handlers
func (r *Routes) CheckDistributionFor(c *gin.Context) {
	name := c.Query("name")
	entityType := c.Query("type") // "city", "province", or "country"
	entityName := c.Query("entity")

	distributor, exists := r.db.Distributors[name]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distributor not found"})
		return
	}

	var allowed bool
	switch entityType {
	case "city":
		_, allowed = distributor.IncludedCites[entityName]
	case "province":
		_, allowed = distributor.IncludedProvinces[entityName]
	case "country":
		_, allowed = distributor.IncludedCountries[entityName]
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"allowed": allowed})
}
func (r *Routes) AllowDistributionTo(c *gin.Context) {
	var request struct {
		Name       string `json:"name"`
		EntityType string `json:"type"` // "city", "province", or "country"
		EntityName string `json:"entity"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	distributor, exists := r.db.Distributors[request.Name]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distributor not found"})
		return
	}

	switch request.EntityType {
	case "city":
		if distributor.IncludedCites == nil {
			distributor.IncludedCites = make(map[string]*models.City)
		}
		distributor.IncludedCites[request.EntityName] = &models.City{CityName: request.EntityName}
	case "province":
		if distributor.IncludedProvinces == nil {
			distributor.IncludedProvinces = make(map[string]*models.Province)
		}
		distributor.IncludedProvinces[request.EntityName] = &models.Province{ProvinceName: request.EntityName}
	case "country":
		if distributor.IncludedCountries == nil {
			distributor.IncludedCountries = make(map[string]*models.Country)
		}
		distributor.IncludedCountries[request.EntityName] = &models.Country{CountryName: request.EntityName}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Distribution allowed successfully"})
}
func (r *Routes) ApplyContractToDistributor(c *gin.Context) {

}
func (r *Routes) ExcludeDistributionTo(c *gin.Context) {
	var request struct {
		Name       string `json:"name"`
		EntityType string `json:"type"` // "city", "province", or "country"
		EntityName string `json:"entity"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	distributor, exists := r.db.Distributors[request.Name]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distributor not found"})
		return
	}

	switch request.EntityType {
	case "city":
		delete(distributor.IncludedCites, request.EntityName)
	case "province":
		delete(distributor.IncludedProvinces, request.EntityName)
	case "country":
		delete(distributor.IncludedCountries, request.EntityName)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Distribution excluded successfully"})
}
func (r *Routes) GetPermissionsOfDistributor(c *gin.Context) {
	name := c.Query("name")
	distributor, exists := r.db.Distributors[name]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distributor not found"})
		return
	}

	permissions := gin.H{
		"allowed_cities":    distributor.IncludedCites,
		"allowed_provinces": distributor.IncludedProvinces,
		"allowed_countries": distributor.IncludedCountries,
	}

	c.JSON(http.StatusOK, permissions)
}
