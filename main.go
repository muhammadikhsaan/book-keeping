package main

import (
	"Accounting/config"
	"Accounting/feature/authorization"
	"Accounting/feature/booking"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

var (
	endpointV1          = "/v1"
	endpointAccount     = "/account"
	endpointAccountByID = "/account/:id"
	endpointBooking     = "/book"
	endpointBookingByID = "/book/:id"
)

func main() {
	// auth := r.Group(endpointV1)
	// {
	// 	auth.POST(endpointAccount)
	// 	auth.GET(endpointAccount)
	// 	auth.GET(endpointAccountByID)
	// 	auth.PUT(endpointAccountByID)
	// 	auth.DELETE(endpointAccountByID)
	// }

	feature := r.Group(endpointV1, authorization.ValidationRequest)
	{
		feature.GET(endpointBooking, booking.GetBookingRequest)
		feature.GET(endpointBookingByID, booking.GetBookingRequest)
		feature.POST(endpointBooking, booking.PostBookingRequest)
		feature.PUT(endpointBookingByID, booking.PutBookingRequest)
		feature.DELETE(endpointBookingByID, booking.DeleteBookingRequest)
	}

	log.Fatal(r.Run(config.APPLICATIONPORT))
}
