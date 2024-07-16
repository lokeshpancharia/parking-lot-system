package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/lokeshpancharia/parking-lot-system/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedParkingLotServiceServer
	parkingLotSystem *ParkingLotSystem
}

type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Truck
)

type Vehicle struct {
	LicensePlate string
	VehicleType  VehicleType
}

type ParkingSpot struct {
	SpotID      int
	SpotSize    VehicleType
	IsAvailable bool
	Vehicle     *Vehicle
	mu          sync.Mutex
}

func NewParkingSpot(id int, size VehicleType) *ParkingSpot {
	return &ParkingSpot{
		SpotID:      id,
		SpotSize:    size,
		IsAvailable: true,
	}
}

func (ps *ParkingSpot) AssignVehicle(vehicle *Vehicle) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.Vehicle = vehicle
	ps.IsAvailable = false
}

func (ps *ParkingSpot) RemoveVehicle() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.Vehicle = nil
	ps.IsAvailable = true
}

type Level struct {
	LevelNumber int
	Spots       []*ParkingSpot
	mu          sync.Mutex
}

func NewLevel(number, numSpots int) *Level {
	spots := make([]*ParkingSpot, numSpots)
	for i := 0; i < numSpots; i++ {
		spots[i] = NewParkingSpot(i, Car) // Simplified to only CAR spots for now
	}
	return &Level{
		LevelNumber: number,
		Spots:       spots,
	}
}

func (l *Level) FindAvailableSpot(vehicleType VehicleType) *ParkingSpot {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, spot := range l.Spots {
		if spot.IsAvailable && spot.SpotSize == vehicleType {
			return spot
		}
	}
	return nil
}

type ParkingLot struct {
	Levels         []*Level
	TotalSpots     int
	AvailableSpots int
	mu             sync.Mutex
}

func NewParkingLot(numLevels, spotsPerLevel int) *ParkingLot {
	levels := make([]*Level, numLevels)
	for i := 0; i < numLevels; i++ {
		levels[i] = NewLevel(i, spotsPerLevel)
	}
	return &ParkingLot{
		Levels:         levels,
		TotalSpots:     numLevels * spotsPerLevel,
		AvailableSpots: numLevels * spotsPerLevel,
	}
}

func (pl *ParkingLot) ParkVehicle(vehicle *Vehicle) bool {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	for _, level := range pl.Levels {
		spot := level.FindAvailableSpot(vehicle.VehicleType)
		if spot != nil {
			spot.AssignVehicle(vehicle)
			pl.AvailableSpots--
			return true
		}
	}
	return false
}

func (pl *ParkingLot) RemoveVehicle(vehicle *Vehicle) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	for _, level := range pl.Levels {
		for _, spot := range level.Spots {
			if spot.Vehicle == vehicle {
				spot.RemoveVehicle()
				pl.AvailableSpots++
				return
			}
		}
	}
}

type Ticket struct {
	TicketID    string
	Vehicle     *Vehicle
	EntryTime   time.Time
	ExitTime    time.Time
	ParkingFees float64
}

func NewTicket(ticketID string, vehicle *Vehicle) *Ticket {
	return &Ticket{
		TicketID:  ticketID,
		Vehicle:   vehicle,
		EntryTime: time.Now(),
	}
}

func (t *Ticket) CalculateFees() {
	duration := t.ExitTime.Sub(t.EntryTime).Hours()
	// Simplified fee calculation
	switch t.Vehicle.VehicleType {
	case Motorcycle:
		t.ParkingFees = duration * 1.0 // $1 per hour
	case Car:
		t.ParkingFees = duration * 2.0 // $2 per hour
	case Truck:
		t.ParkingFees = duration * 3.0 // $3 per hour
	}
}

type ParkingLotSystem struct {
	ParkingLot *ParkingLot
	Tickets    map[string]*Ticket
	mu         sync.Mutex
}

func NewParkingLotSystem(numLevels, spotsPerLevel int) *ParkingLotSystem {
	return &ParkingLotSystem{
		ParkingLot: NewParkingLot(numLevels, spotsPerLevel),
		Tickets:    make(map[string]*Ticket),
	}
}

func (pls *ParkingLotSystem) GenerateTicket(vehicle *Vehicle) string {
	ticketID := fmt.Sprintf("%s-%d", vehicle.LicensePlate, time.Now().UnixNano())
	ticket := NewTicket(ticketID, vehicle)
	pls.mu.Lock()
	pls.Tickets[ticketID] = ticket
	pls.mu.Unlock()
	return ticketID
}

func (pls *ParkingLotSystem) ParkVehicle(ticketID string) bool {
	pls.mu.Lock()
	ticket, exists := pls.Tickets[ticketID]
	pls.mu.Unlock()
	if !exists {
		return false
	}
	success := pls.ParkingLot.ParkVehicle(ticket.Vehicle)
	if success {
		ticket.EntryTime = time.Now()
	}
	return success
}

func (pls *ParkingLotSystem) RemoveVehicle(ticketID string) {
	pls.mu.Lock()
	ticket, exists := pls.Tickets[ticketID]
	pls.mu.Unlock()
	if !exists {
		return
	}
	pls.ParkingLot.RemoveVehicle(ticket.Vehicle)
	ticket.ExitTime = time.Now()
	ticket.CalculateFees()
}

func (s *server) ParkVehicle(ctx context.Context, req *pb.ParkVehicleRequest) (*pb.ParkVehicleResponse, error) {
	var vehicleType VehicleType
	switch req.Vehicle.VehicleType {
	case "Motorcycle":
		vehicleType = Motorcycle
	case "Car":
		vehicleType = Car
	case "Truck":
		vehicleType = Truck
	default:
		return &pb.ParkVehicleResponse{
			Success: false,
			Message: "Invalid vehicle type",
		}, nil
	}

	vehicle := &Vehicle{
		LicensePlate: req.Vehicle.LicensePlate,
		VehicleType:  vehicleType,
	}

	ticketID := s.parkingLotSystem.GenerateTicket(vehicle)
	success := s.parkingLotSystem.ParkVehicle(ticketID)
	if success {
		return &pb.ParkVehicleResponse{
			Success:  true,
			TicketId: ticketID,
			Message:  "Vehicle parked successfully",
		}, nil
	}

	return &pb.ParkVehicleResponse{
		Success: false,
		Message: "Failed to park vehicle",
	}, nil
}

func (s *server) RemoveVehicle(ctx context.Context, req *pb.RemoveVehicleRequest) (*pb.RemoveVehicleResponse, error) {
	s.parkingLotSystem.RemoveVehicle(req.TicketId)
	ticket := s.parkingLotSystem.Tickets[req.TicketId]
	return &pb.RemoveVehicleResponse{
		Success:     true,
		Message:     "Vehicle removed successfully",
		ParkingFees: float32(ticket.ParkingFees),
	}, nil
}

func (s *server) GetAvailableSpots(ctx context.Context, req *pb.GetAvailableSpotsRequest) (*pb.GetAvailableSpotsResponse, error) {
	availableSpots := s.parkingLotSystem.ParkingLot.AvailableSpots
	return &pb.GetAvailableSpotsResponse{
		AvailableSpots: int32(availableSpots),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterParkingLotServiceServer(s, &server{parkingLotSystem: NewParkingLotSystem(3, 10)}) // 3 levels, 10 spots each
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
