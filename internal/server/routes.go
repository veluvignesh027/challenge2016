package server

import (
	"net/http"
	"os"

	"github.com/RealImage/challenge2016/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/veluvignesh027/log"
)

type Routes struct {
	db database.DataStore
}

func NewRoutes() *Routes {
	return &Routes{db: *database.NewDataStore()}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	routes := NewRoutes()
	distributor := r.Group("/distributors/v1/")
	{
		distributor.GET("/get/:name", routes.GetDistributorByName)
		distributor.GET("/get", routes.GetAllDistributors)
		distributor.POST("/create", routes.CreateDistributor)
		distributor.PUT("/update/:name", routes.UpdateDistributor)
		distributor.DELETE("/delete/:name", routes.DeleteDistributor)
		permisson := distributor.Group("/distributor/permission")
		{
			permisson.GET("/check", routes.CheckDistributionFor)
			permisson.POST("/allow", routes.AllowDistributionTo)
			permisson.POST("/apply/contract", routes.ApplyContractToDistributor)
			permisson.POST("/exclude", routes.ExcludeDistributionTo)
			permisson.GET("/permission/:name", routes.GetPermissionsOfDistributor)
		}
	}

	region := r.Group("/region")
	{
		region.GET("/get", routes.GetRegionBy)
		region.POST("/create", routes.AddRegion)
		region.PUT("/update/:name", routes.UpdateRegion)
		region.DELETE("/delete/:name", routes.DeleteRegion)
	}

	debug := r.Group("/debug/api")
	{
		debug.GET("/loadcsv", routes.LoadCSV)
	}

	log.Info("Registered all the routes for the application.")
	return r
}

// Debug API Handlers
func (r *Routes) LoadCSV(c *gin.Context) {
	err := database.LoadCSVFile()
	if err != nil {
		if err == os.ErrNotExist {
			c.HTML(http.StatusNotFound, AlertHTML, ErrCSVFileNotFound)
			return
		}
		c.HTML(http.StatusInternalServerError, AlertHTML, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CSV loaded successfully"})
}
