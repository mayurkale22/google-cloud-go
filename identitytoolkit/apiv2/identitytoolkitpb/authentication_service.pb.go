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
// source: google/cloud/identitytoolkit/v2/authentication_service.proto

package identitytoolkitpb

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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Finalizes sign-in by verifying MFA challenge.
type FinalizeMfaSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Pending credential from first factor sign-in.
	MfaPendingCredential string `protobuf:"bytes,2,opt,name=mfa_pending_credential,json=mfaPendingCredential,proto3" json:"mfa_pending_credential,omitempty"`
	// Proof of completion of the MFA challenge.
	//
	// Types that are assignable to VerificationInfo:
	//
	//	*FinalizeMfaSignInRequest_PhoneVerificationInfo
	VerificationInfo isFinalizeMfaSignInRequest_VerificationInfo `protobuf_oneof:"verification_info"`
	// The ID of the Identity Platform tenant the user is signing in to. If not
	// set, the user will sign in to the default Identity Platform project.
	TenantId string `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *FinalizeMfaSignInRequest) Reset() {
	*x = FinalizeMfaSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeMfaSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeMfaSignInRequest) ProtoMessage() {}

func (x *FinalizeMfaSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeMfaSignInRequest.ProtoReflect.Descriptor instead.
func (*FinalizeMfaSignInRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescGZIP(), []int{0}
}

func (x *FinalizeMfaSignInRequest) GetMfaPendingCredential() string {
	if x != nil {
		return x.MfaPendingCredential
	}
	return ""
}

func (m *FinalizeMfaSignInRequest) GetVerificationInfo() isFinalizeMfaSignInRequest_VerificationInfo {
	if m != nil {
		return m.VerificationInfo
	}
	return nil
}

func (x *FinalizeMfaSignInRequest) GetPhoneVerificationInfo() *FinalizeMfaPhoneRequestInfo {
	if x, ok := x.GetVerificationInfo().(*FinalizeMfaSignInRequest_PhoneVerificationInfo); ok {
		return x.PhoneVerificationInfo
	}
	return nil
}

func (x *FinalizeMfaSignInRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type isFinalizeMfaSignInRequest_VerificationInfo interface {
	isFinalizeMfaSignInRequest_VerificationInfo()
}

type FinalizeMfaSignInRequest_PhoneVerificationInfo struct {
	// Proof of completion of the SMS based MFA challenge.
	PhoneVerificationInfo *FinalizeMfaPhoneRequestInfo `protobuf:"bytes,3,opt,name=phone_verification_info,json=phoneVerificationInfo,proto3,oneof"`
}

func (*FinalizeMfaSignInRequest_PhoneVerificationInfo) isFinalizeMfaSignInRequest_VerificationInfo() {
}

// FinalizeMfaSignIn response.
type FinalizeMfaSignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID token for the authenticated user.
	IdToken string `protobuf:"bytes,1,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	// Refresh token for the authenticated user.
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	// MFA verified sign-in information.
	//
	// Types that are assignable to AuxiliaryAuthInfo:
	//
	//	*FinalizeMfaSignInResponse_PhoneAuthInfo
	AuxiliaryAuthInfo isFinalizeMfaSignInResponse_AuxiliaryAuthInfo `protobuf_oneof:"auxiliary_auth_info"`
}

func (x *FinalizeMfaSignInResponse) Reset() {
	*x = FinalizeMfaSignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeMfaSignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeMfaSignInResponse) ProtoMessage() {}

func (x *FinalizeMfaSignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeMfaSignInResponse.ProtoReflect.Descriptor instead.
func (*FinalizeMfaSignInResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescGZIP(), []int{1}
}

func (x *FinalizeMfaSignInResponse) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

func (x *FinalizeMfaSignInResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (m *FinalizeMfaSignInResponse) GetAuxiliaryAuthInfo() isFinalizeMfaSignInResponse_AuxiliaryAuthInfo {
	if m != nil {
		return m.AuxiliaryAuthInfo
	}
	return nil
}

func (x *FinalizeMfaSignInResponse) GetPhoneAuthInfo() *FinalizeMfaPhoneResponseInfo {
	if x, ok := x.GetAuxiliaryAuthInfo().(*FinalizeMfaSignInResponse_PhoneAuthInfo); ok {
		return x.PhoneAuthInfo
	}
	return nil
}

type isFinalizeMfaSignInResponse_AuxiliaryAuthInfo interface {
	isFinalizeMfaSignInResponse_AuxiliaryAuthInfo()
}

type FinalizeMfaSignInResponse_PhoneAuthInfo struct {
	// Extra phone auth info, including android verification proof.
	PhoneAuthInfo *FinalizeMfaPhoneResponseInfo `protobuf:"bytes,3,opt,name=phone_auth_info,json=phoneAuthInfo,proto3,oneof"`
}

func (*FinalizeMfaSignInResponse_PhoneAuthInfo) isFinalizeMfaSignInResponse_AuxiliaryAuthInfo() {}

// Starts multi-factor sign-in by sending the multi-factor auth challenge.
type StartMfaSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Pending credential from first factor sign-in.
	MfaPendingCredential string `protobuf:"bytes,2,opt,name=mfa_pending_credential,json=mfaPendingCredential,proto3" json:"mfa_pending_credential,omitempty"`
	// Required. MFA enrollment id from the user's list of current MFA enrollments.
	MfaEnrollmentId string `protobuf:"bytes,3,opt,name=mfa_enrollment_id,json=mfaEnrollmentId,proto3" json:"mfa_enrollment_id,omitempty"`
	// MFA information by type of 2nd factor.
	//
	// Types that are assignable to SignInInfo:
	//
	//	*StartMfaSignInRequest_PhoneSignInInfo
	SignInInfo isStartMfaSignInRequest_SignInInfo `protobuf_oneof:"sign_in_info"`
	// The ID of the Identity Platform tenant the user is signing in to. If not
	// set, the user will sign in to the default Identity Platform project.
	TenantId string `protobuf:"bytes,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *StartMfaSignInRequest) Reset() {
	*x = StartMfaSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMfaSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMfaSignInRequest) ProtoMessage() {}

func (x *StartMfaSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMfaSignInRequest.ProtoReflect.Descriptor instead.
func (*StartMfaSignInRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescGZIP(), []int{2}
}

func (x *StartMfaSignInRequest) GetMfaPendingCredential() string {
	if x != nil {
		return x.MfaPendingCredential
	}
	return ""
}

func (x *StartMfaSignInRequest) GetMfaEnrollmentId() string {
	if x != nil {
		return x.MfaEnrollmentId
	}
	return ""
}

func (m *StartMfaSignInRequest) GetSignInInfo() isStartMfaSignInRequest_SignInInfo {
	if m != nil {
		return m.SignInInfo
	}
	return nil
}

func (x *StartMfaSignInRequest) GetPhoneSignInInfo() *StartMfaPhoneRequestInfo {
	if x, ok := x.GetSignInInfo().(*StartMfaSignInRequest_PhoneSignInInfo); ok {
		return x.PhoneSignInInfo
	}
	return nil
}

func (x *StartMfaSignInRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type isStartMfaSignInRequest_SignInInfo interface {
	isStartMfaSignInRequest_SignInInfo()
}

type StartMfaSignInRequest_PhoneSignInInfo struct {
	// Verification info to authorize sending an SMS for phone verification.
	PhoneSignInInfo *StartMfaPhoneRequestInfo `protobuf:"bytes,4,opt,name=phone_sign_in_info,json=phoneSignInInfo,proto3,oneof"`
}

func (*StartMfaSignInRequest_PhoneSignInInfo) isStartMfaSignInRequest_SignInInfo() {}

// StartMfaSignIn response.
type StartMfaSignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MultiFactor start sign-in response by 2nd factor type.
	//
	// Types that are assignable to ResponseInfo:
	//
	//	*StartMfaSignInResponse_PhoneResponseInfo
	ResponseInfo isStartMfaSignInResponse_ResponseInfo `protobuf_oneof:"response_info"`
}

func (x *StartMfaSignInResponse) Reset() {
	*x = StartMfaSignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMfaSignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMfaSignInResponse) ProtoMessage() {}

func (x *StartMfaSignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMfaSignInResponse.ProtoReflect.Descriptor instead.
func (*StartMfaSignInResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescGZIP(), []int{3}
}

func (m *StartMfaSignInResponse) GetResponseInfo() isStartMfaSignInResponse_ResponseInfo {
	if m != nil {
		return m.ResponseInfo
	}
	return nil
}

func (x *StartMfaSignInResponse) GetPhoneResponseInfo() *StartMfaPhoneResponseInfo {
	if x, ok := x.GetResponseInfo().(*StartMfaSignInResponse_PhoneResponseInfo); ok {
		return x.PhoneResponseInfo
	}
	return nil
}

type isStartMfaSignInResponse_ResponseInfo interface {
	isStartMfaSignInResponse_ResponseInfo()
}

type StartMfaSignInResponse_PhoneResponseInfo struct {
	// MultiFactor sign-in session information specific to SMS-type second
	// factors. Along with the one-time code retrieved from the sent SMS, the
	// contents of this session information should be passed to
	// FinalizeMfaSignIn to complete the sign in.
	PhoneResponseInfo *StartMfaPhoneResponseInfo `protobuf:"bytes,1,opt,name=phone_response_info,json=phoneResponseInfo,proto3,oneof"`
}

func (*StartMfaSignInResponse_PhoneResponseInfo) isStartMfaSignInResponse_ResponseInfo() {}

var File_google_cloud_identitytoolkit_v2_authentication_service_proto protoreflect.FileDescriptor

var file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f,
	0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x66, 0x61, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x01, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x16, 0x6d, 0x66, 0x61, 0x5f, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x14, 0x6d, 0x66, 0x61, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12,
	0x76, 0x0a, 0x17, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x15, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x19, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x67, 0x0a, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x15, 0x0a, 0x13, 0x61, 0x75, 0x78, 0x69, 0x6c, 0x69, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4d, 0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x39, 0x0a, 0x16, 0x6d, 0x66, 0x61, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x14, 0x6d, 0x66, 0x61, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x11,
	0x6d, 0x66, 0x61, 0x5f, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x6d, 0x66,
	0x61, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x68, 0x0a,
	0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4d, 0x66, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x66,
	0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6c, 0x0a, 0x13, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4d, 0x66, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x11, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0f, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd1,
	0x03, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb6, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x39,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x53, 0x69, 0x67, 0x6e,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a,
	0x22, 0x1f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x6d,
	0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x3a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x12, 0xaa, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x66, 0x61, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c,
	0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x66, 0x61, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4d, 0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a,
	0x22, 0x1c, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x6d,
	0x66, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x3a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x1a, 0x52,
	0xca, 0x41, 0x1e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b,
	0x69, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x42, 0xdf, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x4d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x70, 0x62, 0x3b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x54, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x5c, 0x56, 0x32, 0xea,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescOnce sync.Once
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescData = file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDesc
)

func file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescGZIP() []byte {
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescData)
	})
	return file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDescData
}

var file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_identitytoolkit_v2_authentication_service_proto_goTypes = []any{
	(*FinalizeMfaSignInRequest)(nil),     // 0: google.cloud.identitytoolkit.v2.FinalizeMfaSignInRequest
	(*FinalizeMfaSignInResponse)(nil),    // 1: google.cloud.identitytoolkit.v2.FinalizeMfaSignInResponse
	(*StartMfaSignInRequest)(nil),        // 2: google.cloud.identitytoolkit.v2.StartMfaSignInRequest
	(*StartMfaSignInResponse)(nil),       // 3: google.cloud.identitytoolkit.v2.StartMfaSignInResponse
	(*FinalizeMfaPhoneRequestInfo)(nil),  // 4: google.cloud.identitytoolkit.v2.FinalizeMfaPhoneRequestInfo
	(*FinalizeMfaPhoneResponseInfo)(nil), // 5: google.cloud.identitytoolkit.v2.FinalizeMfaPhoneResponseInfo
	(*StartMfaPhoneRequestInfo)(nil),     // 6: google.cloud.identitytoolkit.v2.StartMfaPhoneRequestInfo
	(*StartMfaPhoneResponseInfo)(nil),    // 7: google.cloud.identitytoolkit.v2.StartMfaPhoneResponseInfo
}
var file_google_cloud_identitytoolkit_v2_authentication_service_proto_depIdxs = []int32{
	4, // 0: google.cloud.identitytoolkit.v2.FinalizeMfaSignInRequest.phone_verification_info:type_name -> google.cloud.identitytoolkit.v2.FinalizeMfaPhoneRequestInfo
	5, // 1: google.cloud.identitytoolkit.v2.FinalizeMfaSignInResponse.phone_auth_info:type_name -> google.cloud.identitytoolkit.v2.FinalizeMfaPhoneResponseInfo
	6, // 2: google.cloud.identitytoolkit.v2.StartMfaSignInRequest.phone_sign_in_info:type_name -> google.cloud.identitytoolkit.v2.StartMfaPhoneRequestInfo
	7, // 3: google.cloud.identitytoolkit.v2.StartMfaSignInResponse.phone_response_info:type_name -> google.cloud.identitytoolkit.v2.StartMfaPhoneResponseInfo
	0, // 4: google.cloud.identitytoolkit.v2.AuthenticationService.FinalizeMfaSignIn:input_type -> google.cloud.identitytoolkit.v2.FinalizeMfaSignInRequest
	2, // 5: google.cloud.identitytoolkit.v2.AuthenticationService.StartMfaSignIn:input_type -> google.cloud.identitytoolkit.v2.StartMfaSignInRequest
	1, // 6: google.cloud.identitytoolkit.v2.AuthenticationService.FinalizeMfaSignIn:output_type -> google.cloud.identitytoolkit.v2.FinalizeMfaSignInResponse
	3, // 7: google.cloud.identitytoolkit.v2.AuthenticationService.StartMfaSignIn:output_type -> google.cloud.identitytoolkit.v2.StartMfaSignInResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_identitytoolkit_v2_authentication_service_proto_init() }
func file_google_cloud_identitytoolkit_v2_authentication_service_proto_init() {
	if File_google_cloud_identitytoolkit_v2_authentication_service_proto != nil {
		return
	}
	file_google_cloud_identitytoolkit_v2_mfa_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FinalizeMfaSignInRequest); i {
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
		file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FinalizeMfaSignInResponse); i {
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
		file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StartMfaSignInRequest); i {
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
		file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StartMfaSignInResponse); i {
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
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[0].OneofWrappers = []any{
		(*FinalizeMfaSignInRequest_PhoneVerificationInfo)(nil),
	}
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[1].OneofWrappers = []any{
		(*FinalizeMfaSignInResponse_PhoneAuthInfo)(nil),
	}
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[2].OneofWrappers = []any{
		(*StartMfaSignInRequest_PhoneSignInInfo)(nil),
	}
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes[3].OneofWrappers = []any{
		(*StartMfaSignInResponse_PhoneResponseInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_identitytoolkit_v2_authentication_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_identitytoolkit_v2_authentication_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_identitytoolkit_v2_authentication_service_proto_msgTypes,
	}.Build()
	File_google_cloud_identitytoolkit_v2_authentication_service_proto = out.File
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_rawDesc = nil
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_goTypes = nil
	file_google_cloud_identitytoolkit_v2_authentication_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	// Verifies the MFA challenge and performs sign-in
	FinalizeMfaSignIn(ctx context.Context, in *FinalizeMfaSignInRequest, opts ...grpc.CallOption) (*FinalizeMfaSignInResponse, error)
	// Sends the MFA challenge
	StartMfaSignIn(ctx context.Context, in *StartMfaSignInRequest, opts ...grpc.CallOption) (*StartMfaSignInResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) FinalizeMfaSignIn(ctx context.Context, in *FinalizeMfaSignInRequest, opts ...grpc.CallOption) (*FinalizeMfaSignInResponse, error) {
	out := new(FinalizeMfaSignInResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.identitytoolkit.v2.AuthenticationService/FinalizeMfaSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) StartMfaSignIn(ctx context.Context, in *StartMfaSignInRequest, opts ...grpc.CallOption) (*StartMfaSignInResponse, error) {
	out := new(StartMfaSignInResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.identitytoolkit.v2.AuthenticationService/StartMfaSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	// Verifies the MFA challenge and performs sign-in
	FinalizeMfaSignIn(context.Context, *FinalizeMfaSignInRequest) (*FinalizeMfaSignInResponse, error)
	// Sends the MFA challenge
	StartMfaSignIn(context.Context, *StartMfaSignInRequest) (*StartMfaSignInResponse, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) FinalizeMfaSignIn(context.Context, *FinalizeMfaSignInRequest) (*FinalizeMfaSignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeMfaSignIn not implemented")
}
func (*UnimplementedAuthenticationServiceServer) StartMfaSignIn(context.Context, *StartMfaSignInRequest) (*StartMfaSignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMfaSignIn not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_FinalizeMfaSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeMfaSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).FinalizeMfaSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.identitytoolkit.v2.AuthenticationService/FinalizeMfaSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).FinalizeMfaSignIn(ctx, req.(*FinalizeMfaSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_StartMfaSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMfaSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).StartMfaSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.identitytoolkit.v2.AuthenticationService/StartMfaSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).StartMfaSignIn(ctx, req.(*StartMfaSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.identitytoolkit.v2.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FinalizeMfaSignIn",
			Handler:    _AuthenticationService_FinalizeMfaSignIn_Handler,
		},
		{
			MethodName: "StartMfaSignIn",
			Handler:    _AuthenticationService_StartMfaSignIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/identitytoolkit/v2/authentication_service.proto",
}
