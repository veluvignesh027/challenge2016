package server

import (
	"net/http"

	"github.com/RealImage/challenge2016/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/veluvignesh027/log"
)

// Region Handlers
func (r *Routes) GetRegionBy(c *gin.Context) {
	countryQuery := c.Query("country")
	cityQuery := c.Query("city")
	provinceQuery := c.Query("province")
	log.Debug("GetRegionBy: ", countryQuery, cityQuery, provinceQuery)

	country, err := r.db.GetCities(cityQuery, provinceQuery, countryQuery)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cities": country,
	})
}

func (r *Routes) AddRegion(c *gin.Context) {
	var city models.City
	if err := c.BindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	r.db.AddCityInfo(city)
}

func (r *Routes) UpdateRegion(c *gin.Context) {
	var city models.City
	if err := c.BindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	r.db.UpdateCityInfo(city, city)
}

func (r *Routes) DeleteRegion(c *gin.Context) {
	var city models.City
	if err := c.BindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	if ok := r.db.DeleteCityInfo(city); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "City deleted successfully"})
}
