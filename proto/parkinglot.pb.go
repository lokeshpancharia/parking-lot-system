// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: parkinglot.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message representing a vehicle
type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicensePlate string `protobuf:"bytes,1,opt,name=license_plate,json=licensePlate,proto3" json:"license_plate,omitempty"` // License plate of the vehicle
	VehicleType  string `protobuf:"bytes,2,opt,name=vehicle_type,json=vehicleType,proto3" json:"vehicle_type,omitempty"`    // Type of the vehicle (e.g., car, bike)
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_parkinglot_proto_rawDescGZIP(), []int{0}
}

func (x *Vehicle) GetLicensePlate() string {
	if x != nil {
		return x.LicensePlate
	}
	return ""
}

func (x *Vehicle) GetVehicleType() string {
	if x != nil {
		return x.VehicleType
	}
	return ""
}

// Request message for parking a vehicle
type ParkVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicle *Vehicle `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"` // Vehicle to be parked
}

func (x *ParkVehicleRequest) Reset() {
	*x = ParkVehicleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkVehicleRequest) ProtoMessage() {}

func (x *ParkVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkVehicleRequest.ProtoReflect.Descriptor instead.
func (*ParkVehicleRequest) Descriptor() ([]byte, []int) {
	return file_parkinglot_proto_rawDescGZIP(), []int{1}
}

func (x *ParkVehicleRequest) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

// Response message for parking a vehicle
type ParkVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                  // Indicates if the parking was successful
	TicketId string `protobuf:"bytes,2,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"` // Ticket ID for the parked vehicle
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`                   // Additional message or error description
}

func (x *ParkVehicleResponse) Reset() {
	*x = ParkVehicleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkVehicleResponse) ProtoMessage() {}

func (x *ParkVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkVehicleResponse.ProtoReflect.Descriptor instead.
func (*ParkVehicleResponse) Descriptor() ([]byte, []int) {
	return file_parkinglot_proto_rawDescGZIP(), []int{2}
}

func (x *ParkVehicleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ParkVehicleResponse) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *ParkVehicleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request message for removing a vehicle
type RemoveVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketId string `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"` // Ticket ID of the vehicle to be removed
}

func (x *RemoveVehicleRequest) Reset() {
	*x = RemoveVehicleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveVehicleRequest) ProtoMessage() {}

func (x *RemoveVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveVehicleRequest.ProtoReflect.Descriptor instead.
func (*RemoveVehicleRequest) Descriptor() ([]byte, []int) {
	return file_parkinglot_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveVehicleRequest) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

// Response message for removing a vehicle
type RemoveVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                             // Indicates if the removal was successful
	Message     string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`                              // Additional message or error description
	ParkingFees float32 `protobuf:"fixed32,3,opt,name=parking_fees,json=parkingFees,proto3" json:"parking_fees,omitempty"` // Parking fees for the vehicle
}

func (x *RemoveVehicleResponse) Reset() {
	*x = RemoveVehicleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveVehicleResponse) ProtoMessage() {}

func (x *RemoveVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveVehicleResponse.ProtoReflect.Descriptor instead.
func (*RemoveVehicleResponse) Descriptor() ([]byte, []int) {
	return file_parkinglot_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveVehicleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveVehicleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RemoveVehicleResponse) GetParkingFees() float32 {
	if x != nil {
		return x.ParkingFees
	}
	return 0
}

// Request message for getting available spots
type GetAvailableSpotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAvailableSpotsRequest) Reset() {
	*x = GetAvailableSpotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableSpotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableSpotsRequest) ProtoMessage() {}

func (x *GetAvailableSpotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableSpotsRequest.ProtoReflect.Descriptor instead.
func (*GetAvailableSpotsRequest) Descriptor() ([]byte, []int) {
	return file_parkinglot_proto_rawDescGZIP(), []int{5}
}

// Response message for getting available spots
type GetAvailableSpotsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableSpots int32 `protobuf:"varint,1,opt,name=available_spots,json=availableSpots,proto3" json:"available_spots,omitempty"` // Number of available parking spots
}

func (x *GetAvailableSpotsResponse) Reset() {
	*x = GetAvailableSpotsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parkinglot_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableSpotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableSpotsResponse) ProtoMessage() {}

func (x *GetAvailableSpotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parkinglot_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableSpotsResponse.ProtoReflect.Descriptor instead.
func (*GetAvailableSpotsResponse) Descriptor() ([]byte, []int) {
	return file_parkinglot_proto_rawDescGZIP(), []int{6}
}

func (x *GetAvailableSpotsResponse) GetAvailableSpots() int32 {
	if x != nil {
		return x.AvailableSpots
	}
	return 0
}

var File_parkinglot_proto protoreflect.FileDescriptor

var file_parkinglot_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x6c, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x74, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x51, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x12, 0x50, 0x61, 0x72,
	0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x66, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x6b, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33,
	0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x46,
	0x65, 0x65, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x70, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x44, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x70, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x6f, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x70, 0x6f, 0x74, 0x73, 0x32, 0xcb, 0x02, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x4c, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x70, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x74,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x70,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x70, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x70, 0x6f, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c,
	0x6f, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x6f, 0x6b, 0x65, 0x73, 0x68, 0x70, 0x61, 0x6e, 0x63, 0x68, 0x61, 0x72, 0x69,
	0x61, 0x2f, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x6c, 0x6f, 0x74, 0x2d, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_parkinglot_proto_rawDescOnce sync.Once
	file_parkinglot_proto_rawDescData = file_parkinglot_proto_rawDesc
)

func file_parkinglot_proto_rawDescGZIP() []byte {
	file_parkinglot_proto_rawDescOnce.Do(func() {
		file_parkinglot_proto_rawDescData = protoimpl.X.CompressGZIP(file_parkinglot_proto_rawDescData)
	})
	return file_parkinglot_proto_rawDescData
}

var file_parkinglot_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_parkinglot_proto_goTypes = []any{
	(*Vehicle)(nil),                   // 0: parking_lot_system.Vehicle
	(*ParkVehicleRequest)(nil),        // 1: parking_lot_system.ParkVehicleRequest
	(*ParkVehicleResponse)(nil),       // 2: parking_lot_system.ParkVehicleResponse
	(*RemoveVehicleRequest)(nil),      // 3: parking_lot_system.RemoveVehicleRequest
	(*RemoveVehicleResponse)(nil),     // 4: parking_lot_system.RemoveVehicleResponse
	(*GetAvailableSpotsRequest)(nil),  // 5: parking_lot_system.GetAvailableSpotsRequest
	(*GetAvailableSpotsResponse)(nil), // 6: parking_lot_system.GetAvailableSpotsResponse
}
var file_parkinglot_proto_depIdxs = []int32{
	0, // 0: parking_lot_system.ParkVehicleRequest.vehicle:type_name -> parking_lot_system.Vehicle
	1, // 1: parking_lot_system.ParkingLotService.ParkVehicle:input_type -> parking_lot_system.ParkVehicleRequest
	3, // 2: parking_lot_system.ParkingLotService.RemoveVehicle:input_type -> parking_lot_system.RemoveVehicleRequest
	5, // 3: parking_lot_system.ParkingLotService.GetAvailableSpots:input_type -> parking_lot_system.GetAvailableSpotsRequest
	2, // 4: parking_lot_system.ParkingLotService.ParkVehicle:output_type -> parking_lot_system.ParkVehicleResponse
	4, // 5: parking_lot_system.ParkingLotService.RemoveVehicle:output_type -> parking_lot_system.RemoveVehicleResponse
	6, // 6: parking_lot_system.ParkingLotService.GetAvailableSpots:output_type -> parking_lot_system.GetAvailableSpotsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_parkinglot_proto_init() }
func file_parkinglot_proto_init() {
	if File_parkinglot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parkinglot_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Vehicle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parkinglot_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ParkVehicleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parkinglot_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ParkVehicleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parkinglot_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveVehicleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parkinglot_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveVehicleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parkinglot_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvailableSpotsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parkinglot_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvailableSpotsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_parkinglot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parkinglot_proto_goTypes,
		DependencyIndexes: file_parkinglot_proto_depIdxs,
		MessageInfos:      file_parkinglot_proto_msgTypes,
	}.Build()
	File_parkinglot_proto = out.File
	file_parkinglot_proto_rawDesc = nil
	file_parkinglot_proto_goTypes = nil
	file_parkinglot_proto_depIdxs = nil
}
