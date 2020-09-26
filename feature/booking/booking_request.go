package booking

import (
	"Accounting/config"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetBookingRequest to handle request GET to select Booking
func GetBookingRequest(c *gin.Context) {
	var (
		rep    = NeBookingRepository(context.Background(), config.Getdatabase())
		filter = map[string]string{
			"id":       c.Param("id"),
			"type":     c.Query("type"),
			"category": c.Query("category"),
			"dateFrom": c.Query("dateFrom"),
			"dateTo":   c.Query("dateTo"),
		}
	)

	req, err := rep.GetBookingFromDatabase(filter)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if req == nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "No Data"})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, req)
}

func PostBookingRequest(c *gin.Context) {
	var (
		req BookingModel
		rep = NeBookingRepository(context.Background(), config.Getdatabase())
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := rep.PostBookingToDatabase(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Insert Bookkeping Data"})
}

func PutBookingRequest(c *gin.Context) {
	var (
		req    BookingModel
		rep    = NeBookingRepository(context.Background(), config.Getdatabase())
		filter = map[string]string{
			"id": c.Param("id"),
		}
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := rep.PutBookingOnDatabase(&req, filter); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Update Bookkeping Data for Index " + filter["id"]})
}

func DeleteBookingRequest(c *gin.Context) {
	var (
		rep    = NeBookingRepository(context.Background(), config.Getdatabase())
		filter = map[string]string{
			"id": c.Param("id"),
		}
	)

	if err := rep.DeleteBookingOnDatabase(filter); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Success Delete Bookkeping Data for Index " + filter["id"]})
}
