// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/deviceconfiguration.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_syncthing_syncthing_lib_protocol "github.com/syncthing/syncthing/lib/protocol"
	protocol "github.com/syncthing/syncthing/lib/protocol"
	_ "github.com/syncthing/syncthing/proto/ext"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DeviceConfiguration struct {
	DeviceID                 github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"deviceID" xml:"id,attr"`
	Name                     string                                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name" xml:"name,attr,omitempty"`
	Addresses                []string                                             `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses" xml:"address,omitempty" default:"dynamic"`
	Compression              protocol.Compression                                 `protobuf:"varint,4,opt,name=compression,proto3,enum=protocol.Compression" json:"compression" xml:"compression,attr"`
	CertName                 string                                               `protobuf:"bytes,5,opt,name=cert_name,json=certName,proto3" json:"certName" xml:"certName,attr,omitempty"`
	Introducer               bool                                                 `protobuf:"varint,6,opt,name=introducer,proto3" json:"introducer" xml:"introducer,attr"`
	SkipIntroductionRemovals bool                                                 `protobuf:"varint,7,opt,name=skip_introduction_removals,json=skipIntroductionRemovals,proto3" json:"skipIntroductionRemovals" xml:"skipIntroductionRemovals,attr"`
	IntroducedBy             github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,8,opt,name=introduced_by,json=introducedBy,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"introducedBy" xml:"introducedBy,attr"`
	Paused                   bool                                                 `protobuf:"varint,9,opt,name=paused,proto3" json:"paused" xml:"paused"`
	AllowedNetworks          []string                                             `protobuf:"bytes,10,rep,name=allowed_networks,json=allowedNetworks,proto3" json:"allowedNetworks" xml:"allowedNetwork,omitempty"`
	AutoAcceptFolders        bool                                                 `protobuf:"varint,11,opt,name=auto_accept_folders,json=autoAcceptFolders,proto3" json:"autoAcceptFolders" xml:"autoAcceptFolders"`
	MaxSendKbps              int                                                  `protobuf:"varint,12,opt,name=max_send_kbps,json=maxSendKbps,proto3,casttype=int" json:"maxSendKbps" xml:"maxSendKbps"`
	MaxRecvKbps              int                                                  `protobuf:"varint,13,opt,name=max_recv_kbps,json=maxRecvKbps,proto3,casttype=int" json:"maxRecvKbps" xml:"maxRecvKbps"`
	IgnoredFolders           []ObservedFolder                                     `protobuf:"bytes,14,rep,name=ignored_folders,json=ignoredFolders,proto3" json:"ignoredFolders" xml:"ignoredFolder"`
	PendingFolders           []ObservedFolder                                     `protobuf:"bytes,15,rep,name=pending_folders,json=pendingFolders,proto3" json:"pendingFolders" xml:"pendingFolder"`
	MaxRequestKiB            int                                                  `protobuf:"varint,16,opt,name=max_request_kib,json=maxRequestKib,proto3,casttype=int" json:"maxRequestKiB" xml:"maxRequestKiB"`
<<<<<<< HEAD
<<<<<<< HEAD
	RemoteGUIPort            int                                                  `protobuf:"varint,17,opt,name=remote_gui_port,json=remoteGuiPort,proto3,casttype=int" json:"remoteGUIPort" xml:"remoteGUIPort"`
=======
	Untrusted                bool                                                 `protobuf:"varint,17,opt,name=untrusted,proto3" json:"untrusted" xml:"untrusted"`
	RemoteGUIPort            int                                                  `protobuf:"varint,18,opt,name=remote_gui_port,json=remoteGuiPort,proto3,casttype=int" json:"remoteGUIPort" xml:"remoteGUIPort"`
>>>>>>> upstream/main
=======
	Untrusted                bool                                                 `protobuf:"varint,17,opt,name=untrusted,proto3" json:"untrusted" xml:"untrusted"`
	RemoteGUIPort            int                                                  `protobuf:"varint,18,opt,name=remote_gui_port,json=remoteGuiPort,proto3,casttype=int" json:"remoteGUIPort" xml:"remoteGUIPort"`
>>>>>>> dcb916d2eb1b0b452a56b496244eb4c1988789fd
}

func (m *DeviceConfiguration) Reset()         { *m = DeviceConfiguration{} }
func (m *DeviceConfiguration) String() string { return proto.CompactTextString(m) }
func (*DeviceConfiguration) ProtoMessage()    {}
func (*DeviceConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_744b782bd13071dd, []int{0}
}
func (m *DeviceConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceConfiguration.Merge(m, src)
}
func (m *DeviceConfiguration) XXX_Size() int {
	return m.ProtoSize()
}
func (m *DeviceConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeviceConfiguration)(nil), "config.DeviceConfiguration")
}

func init() {
	proto.RegisterFile("lib/config/deviceconfiguration.proto", fileDescriptor_744b782bd13071dd)
}

var fileDescriptor_744b782bd13071dd = []byte{
<<<<<<< HEAD
<<<<<<< HEAD
	// 951 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x31, 0x6f, 0xdb, 0x46,
	0x18, 0x15, 0xeb, 0xc4, 0xb6, 0x68, 0xcb, 0xb2, 0x68, 0xc4, 0x61, 0x0c, 0x44, 0x27, 0xa8, 0x1a,
	0x14, 0x34, 0x95, 0x0b, 0xb7, 0x93, 0xd1, 0x0e, 0x65, 0x82, 0xb6, 0x86, 0xd1, 0x24, 0xbd, 0xa2,
	0x8b, 0x17, 0x96, 0xe4, 0x9d, 0x95, 0x83, 0x45, 0x1e, 0x4b, 0x9e, 0x14, 0x09, 0xe8, 0xd0, 0xb1,
	0x43, 0x87, 0x22, 0x6b, 0x97, 0xa2, 0x43, 0x87, 0xfe, 0x12, 0x6f, 0xd6, 0x58, 0x74, 0x38, 0x20,
	0xf6, 0x54, 0x8e, 0x1c, 0x33, 0x15, 0x77, 0x47, 0x51, 0x24, 0x1d, 0x17, 0x01, 0xb2, 0xdd, 0xbd,
	0xf7, 0xee, 0xbd, 0x8f, 0x1f, 0xef, 0x23, 0xf5, 0xde, 0x88, 0xb8, 0xfb, 0x1e, 0x0d, 0x4e, 0xc9,
	0x70, 0x1f, 0xe1, 0x09, 0xf1, 0xb0, 0xda, 0x8c, 0x23, 0x87, 0x11, 0x1a, 0x0c, 0xc2, 0x88, 0x32,
	0x6a, 0xac, 0x2a, 0x70, 0x6f, 0x57, 0xa8, 0x25, 0xe4, 0xd1, 0xd1, 0xbe, 0x8b, 0x43, 0xc5, 0xef,
	0xdd, 0x2b, 0xb8, 0x50, 0x37, 0xc6, 0xd1, 0x04, 0xa3, 0x8c, 0xaa, 0xe3, 0x29, 0x53, 0xcb, 0xee,
	0xbf, 0x4d, 0x7d, 0xe7, 0xb1, 0xcc, 0x78, 0x54, 0xcc, 0x30, 0xfe, 0xd2, 0xf4, 0xba, 0xca, 0xb6,
	0x09, 0x32, 0xb5, 0x8e, 0xd6, 0xdf, 0xb4, 0x7e, 0xd1, 0xce, 0x39, 0xa8, 0xfd, 0xc3, 0xc1, 0x27,
	0x43, 0xc2, 0x9e, 0x8f, 0xdd, 0x81, 0x47, 0xfd, 0xfd, 0x78, 0x16, 0x78, 0xec, 0x39, 0x09, 0x86,
	0x85, 0x55, 0xb1, 0xa2, 0x81, 0x72, 0x3f, 0x7a, 0x7c, 0xc9, 0xc1, 0xfa, 0x62, 0x9d, 0x70, 0xb0,
	0x8e, 0xb2, 0x75, 0xca, 0x41, 0x63, 0xea, 0x8f, 0x0e, 0xbb, 0x04, 0x3d, 0x74, 0x18, 0x8b, 0xba,
	0xc9, 0x45, 0x6f, 0x2d, 0x5b, 0xa7, 0x17, 0xbd, 0x5c, 0xf7, 0xf3, 0xbc, 0xa7, 0xbd, 0x9c, 0xf7,
	0x72, 0x0f, 0xb8, 0x60, 0x90, 0xf1, 0x4c, 0xbf, 0x15, 0x38, 0x3e, 0x36, 0xdf, 0xeb, 0x68, 0xfd,
	0xba, 0xf5, 0x69, 0xc2, 0x81, 0xdc, 0xa7, 0x1c, 0xdc, 0x93, 0xce, 0x62, 0x23, 0xfd, 0x1e, 0x52,
	0x9f, 0x30, 0xec, 0x87, 0x6c, 0x26, 0x52, 0x76, 0xde, 0x80, 0x43, 0x79, 0xd2, 0x98, 0xea, 0x75,
	0x07, 0xa1, 0x08, 0xc7, 0x31, 0x8e, 0xcd, 0x95, 0xce, 0x4a, 0xbf, 0x6e, 0x9d, 0x24, 0x1c, 0x2c,
	0xc1, 0x94, 0x83, 0x07, 0xd2, 0x3b, 0x43, 0x0a, 0xce, 0x1d, 0x84, 0x4f, 0x9d, 0xf1, 0x88, 0x1d,
	0x76, 0xd1, 0x2c, 0x70, 0x7c, 0xe2, 0x89, 0xac, 0xd6, 0x35, 0xdd, 0xeb, 0x8b, 0xde, 0x5a, 0x26,
	0x80, 0x4b, 0x5f, 0x63, 0xa2, 0x6f, 0x78, 0xd4, 0x0f, 0xc5, 0x8e, 0xd0, 0xc0, 0xbc, 0xd5, 0xd1,
	0xfa, 0x5b, 0x07, 0x77, 0x06, 0x79, 0x3b, 0x1f, 0x2d, 0x49, 0xeb, 0xb3, 0x84, 0x83, 0xa2, 0x3a,
	0xe5, 0x60, 0x57, 0x16, 0x55, 0xc0, 0xf2, 0x9e, 0x6e, 0x57, 0x41, 0x58, 0x3c, 0x6a, 0x60, 0xbd,
	0xee, 0xe1, 0x88, 0xd9, 0xb2, 0x91, 0xb7, 0x65, 0x23, 0xbf, 0x12, 0xaf, 0x49, 0x80, 0x4f, 0x54,
	0x33, 0xef, 0x2b, 0xef, 0x0c, 0x78, 0x43, 0x43, 0xef, 0xde, 0xc0, 0xc1, 0xdc, 0xc5, 0x38, 0xd1,
	0x75, 0x12, 0xb0, 0x88, 0xa2, 0xb1, 0x87, 0x23, 0x73, 0xb5, 0xa3, 0xf5, 0xd7, 0xad, 0xc3, 0x84,
	0x83, 0x02, 0x9a, 0x72, 0x70, 0x47, 0x5d, 0x88, 0x1c, 0xca, 0x1f, 0xa2, 0x59, 0xc1, 0x60, 0xe1,
	0x9c, 0xf1, 0x87, 0xa6, 0xef, 0xc5, 0x67, 0x24, 0xb4, 0x17, 0x98, 0xb8, 0xc9, 0x76, 0x84, 0x7d,
	0x3a, 0x71, 0x46, 0xb1, 0xb9, 0x26, 0xc3, 0x50, 0xc2, 0x81, 0x29, 0x54, 0x47, 0x05, 0x11, 0xcc,
	0x34, 0x29, 0x07, 0xef, 0xcb, 0xe8, 0x9b, 0x04, 0x79, 0x21, 0xf7, 0xff, 0x57, 0x01, 0x6f, 0x4c,
	0x30, 0xfe, 0xd4, 0xf4, 0x46, 0x5e, 0x33, 0xb2, 0xdd, 0x99, 0xb9, 0x2e, 0x87, 0xeb, 0xa7, 0x77,
	0x1a, 0xae, 0x84, 0x83, 0xcd, 0xa5, 0xab, 0x35, 0x4b, 0x39, 0xb8, 0x5b, 0xee, 0x21, 0xb2, 0x66,
	0x79, 0xf1, 0xad, 0x6b, 0xa8, 0x18, 0x2e, 0x58, 0x72, 0x30, 0x0e, 0xf4, 0xd5, 0xd0, 0x19, 0xc7,
	0x18, 0x99, 0x75, 0xd9, 0xb8, 0xbd, 0x84, 0x83, 0x0c, 0x49, 0x39, 0xd8, 0x94, 0xee, 0x6a, 0xdb,
	0x85, 0x19, 0x6e, 0xfc, 0xa8, 0x6f, 0x3b, 0xa3, 0x11, 0x7d, 0x81, 0x91, 0x1d, 0x60, 0xf6, 0x82,
	0x46, 0x67, 0xb1, 0xa9, 0xcb, 0xe9, 0xf9, 0x26, 0xe1, 0xa0, 0x99, 0x71, 0x4f, 0x32, 0x2a, 0xe5,
	0xa0, 0xad, 0x66, 0xa8, 0x84, 0x97, 0xef, 0x94, 0x79, 0x13, 0x09, 0xab, 0x76, 0xc6, 0xf7, 0xfa,
	0x8e, 0x33, 0x66, 0xd4, 0x76, 0x3c, 0x0f, 0x87, 0xcc, 0x3e, 0xa5, 0x23, 0x84, 0xa3, 0xd8, 0xdc,
	0x90, 0xe5, 0x7f, 0x94, 0x70, 0xd0, 0x12, 0xf4, 0xe7, 0x92, 0xfd, 0x42, 0x91, 0x79, 0x9f, 0xae,
	0x31, 0x5d, 0x78, 0x5d, 0x6d, 0x3c, 0xd5, 0x1b, 0xbe, 0x33, 0xb5, 0x63, 0x1c, 0x20, 0xfb, 0xcc,
	0x0d, 0x63, 0x73, 0xb3, 0xa3, 0xf5, 0x6f, 0x5b, 0x1f, 0x88, 0x39, 0xf4, 0x9d, 0xe9, 0xb7, 0x38,
	0x40, 0xc7, 0x6e, 0x28, 0x5c, 0x5b, 0xd2, 0xb5, 0x80, 0x75, 0x5f, 0x73, 0xb0, 0x42, 0x02, 0x06,
	0x8b, 0xc2, 0x85, 0x61, 0x84, 0xbd, 0x89, 0x32, 0x6c, 0x94, 0x0c, 0x21, 0xf6, 0x26, 0x55, 0xc3,
	0x05, 0x56, 0x32, 0x5c, 0x80, 0x46, 0xa0, 0x37, 0xc9, 0x30, 0xa0, 0x11, 0x46, 0xf9, 0xf3, 0x6f,
	0x75, 0x56, 0xfa, 0x1b, 0x07, 0xbb, 0x03, 0xf5, 0x2f, 0x18, 0x3c, 0xcd, 0xfe, 0x05, 0xea, 0x99,
	0xac, 0x0f, 0xc5, 0xb5, 0x4b, 0x38, 0xd8, 0xca, 0x8e, 0x2d, 0x1b, 0xb3, 0xa3, 0x2e, 0x50, 0x11,
	0xee, 0xc2, 0x8a, 0x4c, 0xe4, 0x85, 0x38, 0x40, 0x24, 0x18, 0xe6, 0x79, 0xcd, 0xb7, 0xcb, 0xcb,
	0x8e, 0x55, 0xf3, 0x4a, 0x70, 0x17, 0x56, 0x64, 0xc6, 0x6f, 0x9a, 0xde, 0x54, 0x1d, 0xfb, 0x61,
	0x8c, 0x63, 0x66, 0x9f, 0x11, 0xd7, 0xdc, 0x96, 0x3d, 0x8b, 0x2f, 0x39, 0x68, 0x7c, 0x2d, 0x5a,
	0x21, 0x99, 0x63, 0x62, 0x25, 0x1c, 0x34, 0xfc, 0x22, 0x90, 0x87, 0x94, 0xd0, 0x45, 0x23, 0x93,
	0x8b, 0x5e, 0x45, 0x5e, 0x05, 0x5e, 0xce, 0x7b, 0xe5, 0x04, 0x58, 0xe2, 0x5d, 0x59, 0x9d, 0xf8,
	0xde, 0x30, 0x6c, 0x0f, 0xc7, 0xc4, 0x0e, 0x69, 0xc4, 0xcc, 0xd6, 0xb2, 0x3a, 0x28, 0xa9, 0x2f,
	0xbf, 0x3b, 0x7a, 0x46, 0x23, 0x26, 0xaa, 0x8b, 0x8a, 0x40, 0x5e, 0x5d, 0x09, 0x2d, 0x56, 0x57,
	0x96, 0x57, 0x01, 0x51, 0x5d, 0x29, 0x01, 0x2e, 0xf8, 0x31, 0x11, 0x5b, 0xeb, 0xf8, 0xfc, 0x55,
	0xbb, 0x36, 0x7f, 0xd5, 0xae, 0x9d, 0x5f, 0xb6, 0xb5, 0xf9, 0x65, 0x5b, 0xfb, 0xf5, 0xaa, 0x5d,
	0xfb, 0xfd, 0xaa, 0xad, 0xcd, 0xaf, 0xda, 0xb5, 0xbf, 0xaf, 0xda, 0xb5, 0x93, 0x07, 0x6f, 0xf1,
	0xed, 0x51, 0xaf, 0xd5, 0x5d, 0x95, 0xdf, 0xa0, 0x8f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0xff, 0x95, 0x1f, 0xad, 0x08, 0x00, 0x00,
=======
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x31, 0x6f, 0xdb, 0x46,
=======
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x31, 0x6f, 0xdb, 0x46,
>>>>>>> dcb916d2eb1b0b452a56b496244eb4c1988789fd
	0x18, 0x15, 0xeb, 0xc4, 0xb6, 0x68, 0xcb, 0xb2, 0x68, 0xc4, 0x61, 0x0c, 0x44, 0x27, 0xb0, 0x1a,
	0x14, 0x34, 0x95, 0x0b, 0xb7, 0x93, 0xd1, 0x16, 0x28, 0x13, 0xb4, 0x35, 0x8c, 0x26, 0xe9, 0x15,
	0x5d, 0xbc, 0xb0, 0x24, 0xef, 0xac, 0x1c, 0x2c, 0xf2, 0x58, 0xf2, 0xa8, 0x48, 0x40, 0x87, 0x8e,
	0x1d, 0x3a, 0x14, 0x59, 0xbb, 0x14, 0x1d, 0x0a, 0xb4, 0xbf, 0xc4, 0x9b, 0x35, 0x16, 0x1d, 0x0e,
	0x88, 0xbd, 0x71, 0xe4, 0x98, 0xa9, 0xb8, 0x23, 0x45, 0x91, 0x74, 0x5c, 0x04, 0xe8, 0x76, 0xf7,
	0xde, 0xbb, 0xf7, 0xee, 0xfb, 0xf4, 0x9d, 0xa8, 0xf6, 0xc7, 0xc4, 0xd9, 0x77, 0xa9, 0x7f, 0x4a,
	0x46, 0xfb, 0x08, 0x4f, 0x88, 0x8b, 0xb3, 0x4d, 0x1c, 0xda, 0x8c, 0x50, 0x7f, 0x18, 0x84, 0x94,
	0x51, 0x6d, 0x35, 0x03, 0xf7, 0x76, 0x85, 0x5a, 0x42, 0x2e, 0x1d, 0xef, 0x3b, 0x38, 0xc8, 0xf8,
	0xbd, 0x7b, 0x25, 0x17, 0xea, 0x44, 0x38, 0x9c, 0x60, 0x94, 0x53, 0x4d, 0x3c, 0x65, 0xd9, 0xd2,
	0xf8, 0x73, 0x5b, 0xdd, 0x79, 0x2c, 0x33, 0x1e, 0x95, 0x33, 0xb4, 0xbf, 0x14, 0xb5, 0x99, 0x65,
	0x5b, 0x04, 0xe9, 0x4a, 0x4f, 0x19, 0x6c, 0x9a, 0x3f, 0x2b, 0xe7, 0x1c, 0x34, 0xfe, 0xe1, 0xe0,
	0xa3, 0x11, 0x61, 0xcf, 0x63, 0x67, 0xe8, 0x52, 0x6f, 0x3f, 0x9a, 0xf9, 0x2e, 0x7b, 0x4e, 0xfc,
	0x51, 0x69, 0x55, 0xbe, 0xd1, 0x30, 0x73, 0x3f, 0x7a, 0x7c, 0xc9, 0xc1, 0xfa, 0x62, 0x9d, 0x70,
	0xb0, 0x8e, 0xf2, 0x75, 0xca, 0x41, 0x6b, 0xea, 0x8d, 0x0f, 0x0d, 0x82, 0x1e, 0xda, 0x8c, 0x85,
	0x46, 0x72, 0xd1, 0x5f, 0xcb, 0xd7, 0xe9, 0x45, 0xbf, 0xd0, 0xfd, 0x34, 0xef, 0x2b, 0x2f, 0xe7,
	0xfd, 0xc2, 0x03, 0x2e, 0x18, 0xa4, 0x3d, 0x53, 0x6f, 0xf9, 0xb6, 0x87, 0xf5, 0x77, 0x7a, 0xca,
	0xa0, 0x69, 0x7e, 0x9c, 0x70, 0x20, 0xf7, 0x29, 0x07, 0xf7, 0xa4, 0xb3, 0xd8, 0x48, 0xbf, 0x87,
	0xd4, 0x23, 0x0c, 0x7b, 0x01, 0x9b, 0x89, 0x94, 0x9d, 0x37, 0xe0, 0x50, 0x9e, 0xd4, 0xa6, 0x6a,
	0xd3, 0x46, 0x28, 0xc4, 0x51, 0x84, 0x23, 0x7d, 0xa5, 0xb7, 0x32, 0x68, 0x9a, 0x27, 0x09, 0x07,
	0x4b, 0x30, 0xe5, 0xe0, 0x81, 0xf4, 0xce, 0x91, 0x92, 0x73, 0x0f, 0xe1, 0x53, 0x3b, 0x1e, 0xb3,
	0x43, 0x03, 0xcd, 0x7c, 0xdb, 0x23, 0xae, 0xc8, 0xea, 0x5c, 0xd3, 0xbd, 0xbe, 0xe8, 0xaf, 0xe5,
	0x02, 0xb8, 0xf4, 0xd5, 0x26, 0xea, 0x86, 0x4b, 0xbd, 0x40, 0xec, 0x08, 0xf5, 0xf5, 0x5b, 0x3d,
	0x65, 0xb0, 0x75, 0x70, 0x67, 0x58, 0xb4, 0xf3, 0xd1, 0x92, 0x34, 0x3f, 0x49, 0x38, 0x28, 0xab,
	0x53, 0x0e, 0x76, 0xe5, 0xa5, 0x4a, 0x58, 0xd1, 0xd3, 0xed, 0x3a, 0x08, 0xcb, 0x47, 0x35, 0xac,
	0x36, 0x5d, 0x1c, 0x32, 0x4b, 0x36, 0xf2, 0xb6, 0x6c, 0xe4, 0x97, 0xe2, 0x67, 0x12, 0xe0, 0x93,
	0xac, 0x99, 0xf7, 0x33, 0xef, 0x1c, 0x78, 0x43, 0x43, 0xef, 0xde, 0xc0, 0xc1, 0xc2, 0x45, 0x3b,
	0x51, 0x55, 0xe2, 0xb3, 0x90, 0xa2, 0xd8, 0xc5, 0xa1, 0xbe, 0xda, 0x53, 0x06, 0xeb, 0xe6, 0x61,
	0xc2, 0x41, 0x09, 0x4d, 0x39, 0xb8, 0x93, 0x0d, 0x44, 0x01, 0x15, 0x45, 0xb4, 0x6b, 0x18, 0x2c,
	0x9d, 0xd3, 0x7e, 0x57, 0xd4, 0xbd, 0xe8, 0x8c, 0x04, 0xd6, 0x02, 0x13, 0x93, 0x6c, 0x85, 0xd8,
	0xa3, 0x13, 0x7b, 0x1c, 0xe9, 0x6b, 0x32, 0x0c, 0x25, 0x1c, 0xe8, 0x42, 0x75, 0x54, 0x12, 0xc1,
	0x5c, 0x93, 0x72, 0xf0, 0xae, 0x8c, 0xbe, 0x49, 0x50, 0x5c, 0xe4, 0xfe, 0x7f, 0x2a, 0xe0, 0x8d,
	0x09, 0xda, 0x1f, 0x8a, 0xda, 0x2a, 0xee, 0x8c, 0x2c, 0x67, 0xa6, 0xaf, 0xcb, 0xc7, 0xf5, 0xe3,
	0xff, 0x7a, 0x5c, 0x09, 0x07, 0x9b, 0x4b, 0x57, 0x73, 0x96, 0x72, 0x70, 0xb7, 0xda, 0x43, 0x64,
	0xce, 0x8a, 0xcb, 0x77, 0xae, 0xa1, 0xe2, 0x71, 0xc1, 0x8a, 0x83, 0x76, 0xa0, 0xae, 0x06, 0x76,
	0x1c, 0x61, 0xa4, 0x37, 0x65, 0xe3, 0xf6, 0x12, 0x0e, 0x72, 0x24, 0xe5, 0x60, 0x53, 0xba, 0x67,
	0x5b, 0x03, 0xe6, 0xb8, 0xf6, 0x83, 0xba, 0x6d, 0x8f, 0xc7, 0xf4, 0x05, 0x46, 0x96, 0x8f, 0xd9,
	0x0b, 0x1a, 0x9e, 0x45, 0xba, 0x2a, 0x5f, 0xcf, 0xd7, 0x09, 0x07, 0xed, 0x9c, 0x7b, 0x92, 0x53,
	0x29, 0x07, 0xdd, 0xec, 0x0d, 0x55, 0xf0, 0xea, 0x4c, 0xe9, 0x37, 0x91, 0xb0, 0x6e, 0xa7, 0x7d,
	0xa7, 0xee, 0xd8, 0x31, 0xa3, 0x96, 0xed, 0xba, 0x38, 0x60, 0xd6, 0x29, 0x1d, 0x23, 0x1c, 0x46,
	0xfa, 0x86, 0xbc, 0xfe, 0x07, 0x09, 0x07, 0x1d, 0x41, 0x7f, 0x26, 0xd9, 0xcf, 0x33, 0xb2, 0xe8,
	0xd3, 0x35, 0xc6, 0x80, 0xd7, 0xd5, 0xda, 0x53, 0xb5, 0xe5, 0xd9, 0x53, 0x2b, 0xc2, 0x3e, 0xb2,
	0xce, 0x9c, 0x20, 0xd2, 0x37, 0x7b, 0xca, 0xe0, 0xb6, 0xf9, 0x9e, 0x78, 0x87, 0x9e, 0x3d, 0xfd,
	0x06, 0xfb, 0xe8, 0xd8, 0x09, 0x84, 0x6b, 0x47, 0xba, 0x96, 0x30, 0xe3, 0x35, 0x07, 0x2b, 0xc4,
	0x67, 0xb0, 0x2c, 0x5c, 0x18, 0x86, 0xd8, 0x9d, 0x64, 0x86, 0xad, 0x8a, 0x21, 0xc4, 0xee, 0xa4,
	0x6e, 0xb8, 0xc0, 0x2a, 0x86, 0x0b, 0x50, 0xf3, 0xd5, 0x36, 0x19, 0xf9, 0x34, 0xc4, 0xa8, 0xa8,
	0x7f, 0xab, 0xb7, 0x32, 0xd8, 0x38, 0xd8, 0x1d, 0x66, 0xdf, 0x82, 0xe1, 0xd3, 0xfc, 0x5b, 0x90,
	0xd5, 0x64, 0xbe, 0x2f, 0xc6, 0x2e, 0xe1, 0x60, 0x2b, 0x3f, 0xb6, 0x6c, 0xcc, 0x4e, 0x36, 0x40,
	0x65, 0xd8, 0x80, 0x35, 0x99, 0xc8, 0x0b, 0xb0, 0x8f, 0x88, 0x3f, 0x2a, 0xf2, 0xda, 0x6f, 0x97,
	0x97, 0x1f, 0xab, 0xe7, 0x55, 0x60, 0x03, 0xd6, 0x64, 0xda, 0xaf, 0x8a, 0xda, 0xce, 0x3a, 0xf6,
	0x7d, 0x8c, 0x23, 0x66, 0x9d, 0x11, 0x47, 0xdf, 0x96, 0x3d, 0x8b, 0x2e, 0x39, 0x68, 0x7d, 0x25,
	0x5a, 0x21, 0x99, 0x63, 0x62, 0x26, 0x1c, 0xb4, 0xbc, 0x32, 0x50, 0x84, 0x54, 0xd0, 0x45, 0x23,
	0x93, 0x8b, 0x7e, 0x4d, 0x5e, 0x07, 0x5e, 0xce, 0xfb, 0xd5, 0x04, 0x58, 0xe1, 0x1d, 0xed, 0x53,
	0xb5, 0x19, 0xfb, 0x2c, 0x8c, 0x23, 0x86, 0x91, 0xde, 0x91, 0x73, 0xd7, 0x13, 0x9f, 0x8d, 0x02,
	0x4c, 0x39, 0x68, 0xcb, 0x1b, 0x14, 0x88, 0x01, 0x97, 0xac, 0xac, 0x4e, 0xfc, 0x5f, 0x31, 0x6c,
	0x8d, 0x62, 0x62, 0x05, 0x34, 0x64, 0xba, 0xb6, 0xac, 0x0e, 0x4a, 0xea, 0x8b, 0x6f, 0x8f, 0x9e,
	0xd1, 0x90, 0x89, 0xea, 0xc2, 0x32, 0x50, 0x54, 0x57, 0x41, 0xcb, 0xd5, 0x55, 0xe5, 0x75, 0x40,
	0x54, 0x57, 0x49, 0x80, 0x0b, 0x3e, 0x26, 0x62, 0x6b, 0x1e, 0x9f, 0xbf, 0xea, 0x36, 0xe6, 0xaf,
	0xba, 0x8d, 0xf3, 0xcb, 0xae, 0x32, 0xbf, 0xec, 0x2a, 0xbf, 0x5c, 0x75, 0x1b, 0xbf, 0x5d, 0x75,
	0x95, 0xf9, 0x55, 0xb7, 0xf1, 0xf7, 0x55, 0xb7, 0x71, 0xf2, 0xe0, 0x2d, 0xfe, 0xbb, 0xb2, 0xb1,
	0x70, 0x56, 0xe5, 0x7f, 0xd8, 0x87, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x0e, 0x8a, 0x12,
	0xed, 0x08, 0x00, 0x00,
<<<<<<< HEAD
>>>>>>> upstream/main
=======
>>>>>>> dcb916d2eb1b0b452a56b496244eb4c1988789fd
}

func (m *DeviceConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RemoteGUIPort != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.RemoteGUIPort))
		i--
		dAtA[i] = 0x1
		i--
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> dcb916d2eb1b0b452a56b496244eb4c1988789fd
		dAtA[i] = 0x90
	}
	if m.Untrusted {
		i--
		if m.Untrusted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
>>>>>>> upstream/main
		dAtA[i] = 0x88
	}
	if m.MaxRequestKiB != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.MaxRequestKiB))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.PendingFolders) > 0 {
		for iNdEx := len(m.PendingFolders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingFolders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.IgnoredFolders) > 0 {
		for iNdEx := len(m.IgnoredFolders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IgnoredFolders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if m.MaxRecvKbps != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.MaxRecvKbps))
		i--
		dAtA[i] = 0x68
	}
	if m.MaxSendKbps != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.MaxSendKbps))
		i--
		dAtA[i] = 0x60
	}
	if m.AutoAcceptFolders {
		i--
		if m.AutoAcceptFolders {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.AllowedNetworks) > 0 {
		for iNdEx := len(m.AllowedNetworks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedNetworks[iNdEx])
			copy(dAtA[i:], m.AllowedNetworks[iNdEx])
			i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.AllowedNetworks[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.IntroducedBy.ProtoSize()
		i -= size
		if _, err := m.IntroducedBy.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.SkipIntroductionRemovals {
		i--
		if m.SkipIntroductionRemovals {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Introducer {
		i--
		if m.Introducer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.CertName) > 0 {
		i -= len(m.CertName)
		copy(dAtA[i:], m.CertName)
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.CertName)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Compression != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.Compression))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.DeviceID.ProtoSize()
		i -= size
		if _, err := m.DeviceID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDeviceconfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceconfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DeviceID.ProtoSize()
	n += 1 + l + sovDeviceconfiguration(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDeviceconfiguration(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.Compression != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.Compression))
	}
	l = len(m.CertName)
	if l > 0 {
		n += 1 + l + sovDeviceconfiguration(uint64(l))
	}
	if m.Introducer {
		n += 2
	}
	if m.SkipIntroductionRemovals {
		n += 2
	}
	l = m.IntroducedBy.ProtoSize()
	n += 1 + l + sovDeviceconfiguration(uint64(l))
	if m.Paused {
		n += 2
	}
	if len(m.AllowedNetworks) > 0 {
		for _, s := range m.AllowedNetworks {
			l = len(s)
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.AutoAcceptFolders {
		n += 2
	}
	if m.MaxSendKbps != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.MaxSendKbps))
	}
	if m.MaxRecvKbps != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.MaxRecvKbps))
	}
	if len(m.IgnoredFolders) > 0 {
		for _, e := range m.IgnoredFolders {
			l = e.ProtoSize()
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if len(m.PendingFolders) > 0 {
		for _, e := range m.PendingFolders {
			l = e.ProtoSize()
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.MaxRequestKiB != 0 {
		n += 2 + sovDeviceconfiguration(uint64(m.MaxRequestKiB))
	}
<<<<<<< HEAD
<<<<<<< HEAD
=======
	if m.Untrusted {
		n += 3
	}
>>>>>>> upstream/main
=======
	if m.Untrusted {
		n += 3
	}
>>>>>>> dcb916d2eb1b0b452a56b496244eb4c1988789fd
	if m.RemoteGUIPort != 0 {
		n += 2 + sovDeviceconfiguration(uint64(m.RemoteGUIPort))
	}
	return n
}

func sovDeviceconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceconfiguration(x uint64) (n int) {
	return sovDeviceconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceconfiguration
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeviceID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compression", wireType)
			}
			m.Compression = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Compression |= protocol.Compression(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Introducer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Introducer = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkipIntroductionRemovals", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SkipIntroductionRemovals = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntroducedBy", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IntroducedBy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedNetworks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedNetworks = append(m.AllowedNetworks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoAcceptFolders", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AutoAcceptFolders = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSendKbps", wireType)
			}
			m.MaxSendKbps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSendKbps |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRecvKbps", wireType)
			}
			m.MaxRecvKbps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRecvKbps |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoredFolders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IgnoredFolders = append(m.IgnoredFolders, ObservedFolder{})
			if err := m.IgnoredFolders[len(m.IgnoredFolders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingFolders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingFolders = append(m.PendingFolders, ObservedFolder{})
			if err := m.PendingFolders[len(m.PendingFolders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestKiB", wireType)
			}
			m.MaxRequestKiB = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRequestKiB |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> dcb916d2eb1b0b452a56b496244eb4c1988789fd
				return fmt.Errorf("proto: wrong wireType = %d for field Untrusted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Untrusted = bool(v != 0)
		case 18:
			if wireType != 0 {
<<<<<<< HEAD
>>>>>>> upstream/main
=======
>>>>>>> dcb916d2eb1b0b452a56b496244eb4c1988789fd
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteGUIPort", wireType)
			}
			m.RemoteGUIPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteGUIPort |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceconfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDeviceconfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceconfiguration
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDeviceconfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceconfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceconfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceconfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceconfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceconfiguration = fmt.Errorf("proto: unexpected end of group")
)
