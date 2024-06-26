// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: services/payment/grpc/payment.proto

package grpc

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

type RecordTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Type          string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Paymentmethod string `protobuf:"bytes,4,opt,name=paymentmethod,proto3" json:"paymentmethod,omitempty"`
	CheckinTime   string `protobuf:"bytes,5,opt,name=checkin_time,json=checkinTime,proto3" json:"checkin_time,omitempty"`
	CheckoutTime  string `protobuf:"bytes,6,opt,name=checkout_time,json=checkoutTime,proto3" json:"checkout_time,omitempty"`
}

func (x *RecordTransactionRequest) Reset() {
	*x = RecordTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_payment_grpc_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordTransactionRequest) ProtoMessage() {}

func (x *RecordTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_payment_grpc_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordTransactionRequest.ProtoReflect.Descriptor instead.
func (*RecordTransactionRequest) Descriptor() ([]byte, []int) {
	return file_services_payment_grpc_payment_proto_rawDescGZIP(), []int{0}
}

func (x *RecordTransactionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RecordTransactionRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RecordTransactionRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RecordTransactionRequest) GetPaymentmethod() string {
	if x != nil {
		return x.Paymentmethod
	}
	return ""
}

func (x *RecordTransactionRequest) GetCheckinTime() string {
	if x != nil {
		return x.CheckinTime
	}
	return ""
}

func (x *RecordTransactionRequest) GetCheckoutTime() string {
	if x != nil {
		return x.CheckoutTime
	}
	return ""
}

type RecordTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RecordTransactionResponse) Reset() {
	*x = RecordTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_payment_grpc_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordTransactionResponse) ProtoMessage() {}

func (x *RecordTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_payment_grpc_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordTransactionResponse.ProtoReflect.Descriptor instead.
func (*RecordTransactionResponse) Descriptor() ([]byte, []int) {
	return file_services_payment_grpc_payment_proto_rawDescGZIP(), []int{1}
}

func (x *RecordTransactionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RecordTransactionResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ProcessPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Type          string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Paymentmethod string `protobuf:"bytes,4,opt,name=paymentmethod,proto3" json:"paymentmethod,omitempty"`
	CheckinTime   string `protobuf:"bytes,5,opt,name=checkin_time,json=checkinTime,proto3" json:"checkin_time,omitempty"`
	CheckoutTime  string `protobuf:"bytes,6,opt,name=checkout_time,json=checkoutTime,proto3" json:"checkout_time,omitempty"`
}

func (x *ProcessPaymentRequest) Reset() {
	*x = ProcessPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_payment_grpc_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentRequest) ProtoMessage() {}

func (x *ProcessPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_payment_grpc_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentRequest.ProtoReflect.Descriptor instead.
func (*ProcessPaymentRequest) Descriptor() ([]byte, []int) {
	return file_services_payment_grpc_payment_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessPaymentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProcessPaymentRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ProcessPaymentRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ProcessPaymentRequest) GetPaymentmethod() string {
	if x != nil {
		return x.Paymentmethod
	}
	return ""
}

func (x *ProcessPaymentRequest) GetCheckinTime() string {
	if x != nil {
		return x.CheckinTime
	}
	return ""
}

func (x *ProcessPaymentRequest) GetCheckoutTime() string {
	if x != nil {
		return x.CheckoutTime
	}
	return ""
}

type ProcessPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ProcessPaymentResponse) Reset() {
	*x = ProcessPaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_payment_grpc_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentResponse) ProtoMessage() {}

func (x *ProcessPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_payment_grpc_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentResponse.ProtoReflect.Descriptor instead.
func (*ProcessPaymentResponse) Descriptor() ([]byte, []int) {
	return file_services_payment_grpc_payment_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessPaymentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ProcessPaymentResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_services_payment_grpc_payment_proto protoreflect.FileDescriptor

var file_services_payment_grpc_payment_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6e, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x19, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32,
	0xdb, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x68, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25,
	0x2e, 0x65, 0x6e, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x6e, 0x74, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a,
	0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_payment_grpc_payment_proto_rawDescOnce sync.Once
	file_services_payment_grpc_payment_proto_rawDescData = file_services_payment_grpc_payment_proto_rawDesc
)

func file_services_payment_grpc_payment_proto_rawDescGZIP() []byte {
	file_services_payment_grpc_payment_proto_rawDescOnce.Do(func() {
		file_services_payment_grpc_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_payment_grpc_payment_proto_rawDescData)
	})
	return file_services_payment_grpc_payment_proto_rawDescData
}

var file_services_payment_grpc_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_payment_grpc_payment_proto_goTypes = []interface{}{
	(*RecordTransactionRequest)(nil),  // 0: entpaymentance.RecordTransactionRequest
	(*RecordTransactionResponse)(nil), // 1: entpaymentance.RecordTransactionResponse
	(*ProcessPaymentRequest)(nil),     // 2: entpaymentance.ProcessPaymentRequest
	(*ProcessPaymentResponse)(nil),    // 3: entpaymentance.ProcessPaymentResponse
}
var file_services_payment_grpc_payment_proto_depIdxs = []int32{
	0, // 0: entpaymentance.PaymentService.RecordTransaction:input_type -> entpaymentance.RecordTransactionRequest
	2, // 1: entpaymentance.PaymentService.ProcessPayment:input_type -> entpaymentance.ProcessPaymentRequest
	1, // 2: entpaymentance.PaymentService.RecordTransaction:output_type -> entpaymentance.RecordTransactionResponse
	3, // 3: entpaymentance.PaymentService.ProcessPayment:output_type -> entpaymentance.ProcessPaymentResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_payment_grpc_payment_proto_init() }
func file_services_payment_grpc_payment_proto_init() {
	if File_services_payment_grpc_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_payment_grpc_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordTransactionRequest); i {
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
		file_services_payment_grpc_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordTransactionResponse); i {
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
		file_services_payment_grpc_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessPaymentRequest); i {
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
		file_services_payment_grpc_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessPaymentResponse); i {
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
			RawDescriptor: file_services_payment_grpc_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_payment_grpc_payment_proto_goTypes,
		DependencyIndexes: file_services_payment_grpc_payment_proto_depIdxs,
		MessageInfos:      file_services_payment_grpc_payment_proto_msgTypes,
	}.Build()
	File_services_payment_grpc_payment_proto = out.File
	file_services_payment_grpc_payment_proto_rawDesc = nil
	file_services_payment_grpc_payment_proto_goTypes = nil
	file_services_payment_grpc_payment_proto_depIdxs = nil
}
