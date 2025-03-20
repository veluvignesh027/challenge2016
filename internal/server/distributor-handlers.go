package server

import (
	"net/http"

	"github.com/RealImage/challenge2016/internal/models"
	"github.com/gin-gonic/gin"
)

// Distributor Handlers
func (r *Routes) GetDistributorByName(c *gin.Context) {
	name := c.Query("name")
	distributor, exists := r.db.Distributors[name]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distributor not found"})
		return
	}
	c.JSON(http.StatusOK, distributor)
}

func (r *Routes) GetAllDistributors(c *gin.Context) {
	distributors := make([]*models.Distributor, 0, len(r.db.Distributors))
	for _, distributor := range r.db.Distributors {
		distributors = append(distributors, distributor)
	}
	c.JSON(http.StatusOK, distributors)
}

func (r *Routes) CreateDistributor(c *gin.Context) {
	var distributor *models.Distributor
	if err := c.BindJSON(&distributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Check if the distributor already exists
	if _, exists := r.db.Distributors[distributor.Name]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Distributor already exists"})
		return
	}

	new := models.NewDistributor(distributor.Name)
	if distributor.ExcludedCities != nil {
		new.ExcludedCities = distributor.ExcludedCities
	}
	if distributor.IncludedCites != nil {
		new.IncludedCites = distributor.IncludedCites
	}
	if distributor.ExcludedCountries != nil {
		new.ExcludedCountries = distributor.ExcludedCountries
	}
	if distributor.IncludedCountries != nil {
		new.IncludedCountries = distributor.IncludedCountries
	}
	if distributor.ExcludedProvinces != nil {
		new.ExcludedProvinces = distributor.ExcludedProvinces
	}
	if distributor.IncludedProvinces != nil {
		new.IncludedProvinces = distributor.IncludedProvinces
	}

	// Add the new distributor to the database
	r.db.Distributors[distributor.Name] = new
	c.JSON(http.StatusCreated, gin.H{"message": "Distributor created successfully"})
}
func (r *Routes) UpdateDistributor(c *gin.Context) {
	var distributor models.Distributor
	if err := c.BindJSON(&distributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Check if the distributor exists
	if _, exists := r.db.Distributors[distributor.Name]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distributor not found"})
		return
	}

	// Update the distributor in the database
	r.db.Distributors[distributor.Name] = &distributor
	c.JSON(http.StatusOK, gin.H{"message": "Distributor updated successfully"})
}
func (r *Routes) DeleteDistributor(c *gin.Context) {
	name := c.Query("name")
	if _, exists := r.db.Distributors[name]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distributor not found"})
		return
	}

	// Delete the distributor from the database
	delete(r.db.Distributors, name)
	c.JSON(http.StatusOK, gin.H{"message": "Distributor deleted successfully"})
}
