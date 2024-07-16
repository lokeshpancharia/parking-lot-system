package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/lokeshpancharia/parking-lot-system/proto"
	
)

var client pb.ParkingLotServiceClient

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = pb.NewParkingLotServiceClient(conn)

	// Set up Gin router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showLoginPage)
	router.POST("/login", handleLogin)
	router.GET("/dashboard", showDashboard)
	router.POST("/park", parkVehicle)
	router.POST("/remove", removeVehicle)
	router.GET("/available-spots", getAvailableSpots)

	router.Run(":8080")
}

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func handleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// For simplicity, accept any username/password
	if username != "" && password != "" {
		c.Redirect(http.StatusFound, "/dashboard")
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}

func showDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func parkVehicle(c *gin.Context) {
	licensePlate := c.PostForm("license_plate")
	vehicleType := c.PostForm("vehicle_type")

	vehicle := &pb.Vehicle{
		LicensePlate: licensePlate,
		VehicleType:  vehicleType,
	}

	req := &pb.ParkVehicleRequest{Vehicle: vehicle}
	res, err := client.ParkVehicle(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res.Message, "ticket_id": res.TicketId})
}

func removeVehicle(c *gin.Context) {
	ticketID := c.PostForm("ticket_id")

	req := &pb.RemoveVehicleRequest{TicketId: ticketID}
	res, err := client.RemoveVehicle(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res.Message, "parking_fees": res.ParkingFees})
}

func getAvailableSpots(c *gin.Context) {
	req := &pb.GetAvailableSpotsRequest{}
	res, err := client.GetAvailableSpots(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available_spots": res.AvailableSpots})
}
