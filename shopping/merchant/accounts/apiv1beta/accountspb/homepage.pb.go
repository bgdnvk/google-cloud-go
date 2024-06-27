// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/shopping/merchant/accounts/v1beta/homepage.proto

package accountspb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A store's homepage.
type Homepage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the store's homepage.
	// Format: `accounts/{account}/homepage`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The URI (typically a URL) of the store's homepage.
	Uri *string `protobuf:"bytes,2,opt,name=uri,proto3,oneof" json:"uri,omitempty"`
	// Output only. Whether the homepage is claimed. See
	// https://support.google.com/merchants/answer/176793.
	Claimed bool `protobuf:"varint,3,opt,name=claimed,proto3" json:"claimed,omitempty"`
}

func (x *Homepage) Reset() {
	*x = Homepage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Homepage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Homepage) ProtoMessage() {}

func (x *Homepage) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Homepage.ProtoReflect.Descriptor instead.
func (*Homepage) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescGZIP(), []int{0}
}

func (x *Homepage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Homepage) GetUri() string {
	if x != nil && x.Uri != nil {
		return *x.Uri
	}
	return ""
}

func (x *Homepage) GetClaimed() bool {
	if x != nil {
		return x.Claimed
	}
	return false
}

// Request message for the `GetHomepage` method.
type GetHomepageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the homepage to retrieve.
	// Format: `accounts/{account}/homepage`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetHomepageRequest) Reset() {
	*x = GetHomepageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomepageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomepageRequest) ProtoMessage() {}

func (x *GetHomepageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomepageRequest.ProtoReflect.Descriptor instead.
func (*GetHomepageRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescGZIP(), []int{1}
}

func (x *GetHomepageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for the `UpdateHomepage` method.
type UpdateHomepageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The new version of the homepage.
	Homepage *Homepage `protobuf:"bytes,1,opt,name=homepage,proto3" json:"homepage,omitempty"`
	// Required. List of fields being updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateHomepageRequest) Reset() {
	*x = UpdateHomepageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHomepageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHomepageRequest) ProtoMessage() {}

func (x *UpdateHomepageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHomepageRequest.ProtoReflect.Descriptor instead.
func (*UpdateHomepageRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHomepageRequest) GetHomepage() *Homepage {
	if x != nil {
		return x.Homepage
	}
	return nil
}

func (x *UpdateHomepageRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request message for the `ClaimHomepage` method.
type ClaimHomepageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the homepage to claim.
	// Format: `accounts/{account}/homepage`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ClaimHomepageRequest) Reset() {
	*x = ClaimHomepageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimHomepageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimHomepageRequest) ProtoMessage() {}

func (x *ClaimHomepageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimHomepageRequest.ProtoReflect.Descriptor instead.
func (*ClaimHomepageRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescGZIP(), []int{3}
}

func (x *ClaimHomepageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for the `UnclaimHomepage` method.
type UnclaimHomepageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the homepage to unclaim.
	// Format: `accounts/{account}/homepage`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UnclaimHomepageRequest) Reset() {
	*x = UnclaimHomepageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnclaimHomepageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnclaimHomepageRequest) ProtoMessage() {}

func (x *UnclaimHomepageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnclaimHomepageRequest.ProtoReflect.Descriptor instead.
func (*UnclaimHomepageRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescGZIP(), []int{4}
}

func (x *UnclaimHomepageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_shopping_merchant_accounts_v1beta_homepage_proto protoreflect.FileDescriptor

var file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x08, 0x48, 0x6f, 0x6d,
	0x65, 0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x69, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x07, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x3a, 0x5a, 0xea, 0x41, 0x57, 0x0a, 0x23,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x6d, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65,
	0x2a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x73, 0x32, 0x08, 0x68, 0x6f, 0x6d,
	0x65, 0x70, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x69, 0x22, 0x55, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53,
	0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x6f, 0x6d, 0x65,
	0x70, 0x61, 0x67, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x57, 0x0a, 0x14, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x48, 0x6f,
	0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x25, 0x0a, 0x23, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x59,
	0x0a, 0x16, 0x55, 0x6e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x0a, 0x23,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x6d, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x8d, 0x07, 0x0a, 0x0f, 0x48, 0x6f,
	0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbb, 0x01,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65,
	0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x3a, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x7d, 0x12, 0xe4, 0x01, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x12, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x22, 0x5d, 0xda, 0x41, 0x14, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65,
	0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x40, 0x3a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x32, 0x34, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b,
	0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67,
	0x65, 0x7d, 0x12, 0xc1, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x48, 0x6f, 0x6d, 0x65,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36,
	0x3a, 0x01, 0x2a, 0x22, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x7d,
	0x3a, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0xc7, 0x01, 0x0a, 0x0f, 0x55, 0x6e, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x12, 0x40, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x55, 0x6e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x48, 0x6f, 0x6d,
	0x65, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x3a, 0x01, 0x2a, 0x22, 0x33, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x68,
	0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x7d, 0x3a, 0x75, 0x6e, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x1a, 0x47, 0xca, 0x41, 0x1a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x8f, 0x01, 0x0a, 0x2c, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0d, 0x48, 0x6f, 0x6d, 0x65,
	0x70, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x70, 0x62,
	0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescOnce sync.Once
	file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescData = file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDesc
)

func file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescGZIP() []byte {
	file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescOnce.Do(func() {
		file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescData)
	})
	return file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDescData
}

var file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_shopping_merchant_accounts_v1beta_homepage_proto_goTypes = []any{
	(*Homepage)(nil),               // 0: google.shopping.merchant.accounts.v1beta.Homepage
	(*GetHomepageRequest)(nil),     // 1: google.shopping.merchant.accounts.v1beta.GetHomepageRequest
	(*UpdateHomepageRequest)(nil),  // 2: google.shopping.merchant.accounts.v1beta.UpdateHomepageRequest
	(*ClaimHomepageRequest)(nil),   // 3: google.shopping.merchant.accounts.v1beta.ClaimHomepageRequest
	(*UnclaimHomepageRequest)(nil), // 4: google.shopping.merchant.accounts.v1beta.UnclaimHomepageRequest
	(*fieldmaskpb.FieldMask)(nil),  // 5: google.protobuf.FieldMask
}
var file_google_shopping_merchant_accounts_v1beta_homepage_proto_depIdxs = []int32{
	0, // 0: google.shopping.merchant.accounts.v1beta.UpdateHomepageRequest.homepage:type_name -> google.shopping.merchant.accounts.v1beta.Homepage
	5, // 1: google.shopping.merchant.accounts.v1beta.UpdateHomepageRequest.update_mask:type_name -> google.protobuf.FieldMask
	1, // 2: google.shopping.merchant.accounts.v1beta.HomepageService.GetHomepage:input_type -> google.shopping.merchant.accounts.v1beta.GetHomepageRequest
	2, // 3: google.shopping.merchant.accounts.v1beta.HomepageService.UpdateHomepage:input_type -> google.shopping.merchant.accounts.v1beta.UpdateHomepageRequest
	3, // 4: google.shopping.merchant.accounts.v1beta.HomepageService.ClaimHomepage:input_type -> google.shopping.merchant.accounts.v1beta.ClaimHomepageRequest
	4, // 5: google.shopping.merchant.accounts.v1beta.HomepageService.UnclaimHomepage:input_type -> google.shopping.merchant.accounts.v1beta.UnclaimHomepageRequest
	0, // 6: google.shopping.merchant.accounts.v1beta.HomepageService.GetHomepage:output_type -> google.shopping.merchant.accounts.v1beta.Homepage
	0, // 7: google.shopping.merchant.accounts.v1beta.HomepageService.UpdateHomepage:output_type -> google.shopping.merchant.accounts.v1beta.Homepage
	0, // 8: google.shopping.merchant.accounts.v1beta.HomepageService.ClaimHomepage:output_type -> google.shopping.merchant.accounts.v1beta.Homepage
	0, // 9: google.shopping.merchant.accounts.v1beta.HomepageService.UnclaimHomepage:output_type -> google.shopping.merchant.accounts.v1beta.Homepage
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_shopping_merchant_accounts_v1beta_homepage_proto_init() }
func file_google_shopping_merchant_accounts_v1beta_homepage_proto_init() {
	if File_google_shopping_merchant_accounts_v1beta_homepage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Homepage); i {
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
		file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetHomepageRequest); i {
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
		file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateHomepageRequest); i {
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
		file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ClaimHomepageRequest); i {
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
		file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UnclaimHomepageRequest); i {
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
	file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_shopping_merchant_accounts_v1beta_homepage_proto_goTypes,
		DependencyIndexes: file_google_shopping_merchant_accounts_v1beta_homepage_proto_depIdxs,
		MessageInfos:      file_google_shopping_merchant_accounts_v1beta_homepage_proto_msgTypes,
	}.Build()
	File_google_shopping_merchant_accounts_v1beta_homepage_proto = out.File
	file_google_shopping_merchant_accounts_v1beta_homepage_proto_rawDesc = nil
	file_google_shopping_merchant_accounts_v1beta_homepage_proto_goTypes = nil
	file_google_shopping_merchant_accounts_v1beta_homepage_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HomepageServiceClient is the client API for HomepageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HomepageServiceClient interface {
	// Retrieves a store's homepage.
	GetHomepage(ctx context.Context, in *GetHomepageRequest, opts ...grpc.CallOption) (*Homepage, error)
	// Updates a store's homepage. Executing this method requires admin access.
	UpdateHomepage(ctx context.Context, in *UpdateHomepageRequest, opts ...grpc.CallOption) (*Homepage, error)
	// Claims a store's homepage. Executing this method requires admin access.
	//
	// If the homepage is already claimed, this will recheck the
	// verification (unless the merchant is exempted from claiming, which also
	// exempts from verification) and return a successful response. If ownership
	// can no longer be verified, it will return an error, but it won't clear the
	// claim. In case of failure, a canonical error message will be returned:
	//   - PERMISSION_DENIED: user doesn't have the necessary permissions on this
	//     MC account;
	//   - FAILED_PRECONDITION:
	//   - The account is not a Merchant Center account;
	//   - MC account doesn't have a homepage;
	//   - claiming failed (in this case the error message will contain more
	//     details).
	ClaimHomepage(ctx context.Context, in *ClaimHomepageRequest, opts ...grpc.CallOption) (*Homepage, error)
	// Unclaims a store's homepage. Executing this method requires admin access.
	UnclaimHomepage(ctx context.Context, in *UnclaimHomepageRequest, opts ...grpc.CallOption) (*Homepage, error)
}

type homepageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomepageServiceClient(cc grpc.ClientConnInterface) HomepageServiceClient {
	return &homepageServiceClient{cc}
}

func (c *homepageServiceClient) GetHomepage(ctx context.Context, in *GetHomepageRequest, opts ...grpc.CallOption) (*Homepage, error) {
	out := new(Homepage)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.HomepageService/GetHomepage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homepageServiceClient) UpdateHomepage(ctx context.Context, in *UpdateHomepageRequest, opts ...grpc.CallOption) (*Homepage, error) {
	out := new(Homepage)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.HomepageService/UpdateHomepage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homepageServiceClient) ClaimHomepage(ctx context.Context, in *ClaimHomepageRequest, opts ...grpc.CallOption) (*Homepage, error) {
	out := new(Homepage)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.HomepageService/ClaimHomepage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homepageServiceClient) UnclaimHomepage(ctx context.Context, in *UnclaimHomepageRequest, opts ...grpc.CallOption) (*Homepage, error) {
	out := new(Homepage)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.accounts.v1beta.HomepageService/UnclaimHomepage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomepageServiceServer is the server API for HomepageService service.
type HomepageServiceServer interface {
	// Retrieves a store's homepage.
	GetHomepage(context.Context, *GetHomepageRequest) (*Homepage, error)
	// Updates a store's homepage. Executing this method requires admin access.
	UpdateHomepage(context.Context, *UpdateHomepageRequest) (*Homepage, error)
	// Claims a store's homepage. Executing this method requires admin access.
	//
	// If the homepage is already claimed, this will recheck the
	// verification (unless the merchant is exempted from claiming, which also
	// exempts from verification) and return a successful response. If ownership
	// can no longer be verified, it will return an error, but it won't clear the
	// claim. In case of failure, a canonical error message will be returned:
	//   - PERMISSION_DENIED: user doesn't have the necessary permissions on this
	//     MC account;
	//   - FAILED_PRECONDITION:
	//   - The account is not a Merchant Center account;
	//   - MC account doesn't have a homepage;
	//   - claiming failed (in this case the error message will contain more
	//     details).
	ClaimHomepage(context.Context, *ClaimHomepageRequest) (*Homepage, error)
	// Unclaims a store's homepage. Executing this method requires admin access.
	UnclaimHomepage(context.Context, *UnclaimHomepageRequest) (*Homepage, error)
}

// UnimplementedHomepageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHomepageServiceServer struct {
}

func (*UnimplementedHomepageServiceServer) GetHomepage(context.Context, *GetHomepageRequest) (*Homepage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomepage not implemented")
}
func (*UnimplementedHomepageServiceServer) UpdateHomepage(context.Context, *UpdateHomepageRequest) (*Homepage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHomepage not implemented")
}
func (*UnimplementedHomepageServiceServer) ClaimHomepage(context.Context, *ClaimHomepageRequest) (*Homepage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimHomepage not implemented")
}
func (*UnimplementedHomepageServiceServer) UnclaimHomepage(context.Context, *UnclaimHomepageRequest) (*Homepage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnclaimHomepage not implemented")
}

func RegisterHomepageServiceServer(s *grpc.Server, srv HomepageServiceServer) {
	s.RegisterService(&_HomepageService_serviceDesc, srv)
}

func _HomepageService_GetHomepage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomepageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomepageServiceServer).GetHomepage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.HomepageService/GetHomepage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomepageServiceServer).GetHomepage(ctx, req.(*GetHomepageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomepageService_UpdateHomepage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomepageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomepageServiceServer).UpdateHomepage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.HomepageService/UpdateHomepage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomepageServiceServer).UpdateHomepage(ctx, req.(*UpdateHomepageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomepageService_ClaimHomepage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimHomepageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomepageServiceServer).ClaimHomepage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.HomepageService/ClaimHomepage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomepageServiceServer).ClaimHomepage(ctx, req.(*ClaimHomepageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomepageService_UnclaimHomepage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnclaimHomepageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomepageServiceServer).UnclaimHomepage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.accounts.v1beta.HomepageService/UnclaimHomepage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomepageServiceServer).UnclaimHomepage(ctx, req.(*UnclaimHomepageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HomepageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.accounts.v1beta.HomepageService",
	HandlerType: (*HomepageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHomepage",
			Handler:    _HomepageService_GetHomepage_Handler,
		},
		{
			MethodName: "UpdateHomepage",
			Handler:    _HomepageService_UpdateHomepage_Handler,
		},
		{
			MethodName: "ClaimHomepage",
			Handler:    _HomepageService_ClaimHomepage_Handler,
		},
		{
			MethodName: "UnclaimHomepage",
			Handler:    _HomepageService_UnclaimHomepage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/accounts/v1beta/homepage.proto",
}
