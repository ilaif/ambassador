// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto

package envoy_config_filter_network_mongo_proxy_v2

import (
	fmt "fmt"
	v2 "github.com/datawire/ambassador/go/apis/envoy/config/filter/fault/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MongoProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mongo_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated :ref:`runtime <config_network_filters_mongo_proxy_runtime>`.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Inject a fixed delay before proxying a Mongo operation. Delays are
	// applied to the following MongoDB operations: Query, Insert, GetMore,
	// and KillCursors. Once an active delay is in progress, all incoming
	// data up until the timer event fires will be a part of the delay.
	Delay *v2.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
	// Flag to specify whether :ref:`dynamic metadata
	// <config_network_filters_mongo_proxy_dynamic_metadata>` should be emitted. Defaults to false.
	EmitDynamicMetadata  bool     `protobuf:"varint,4,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MongoProxy) Reset()         { *m = MongoProxy{} }
func (m *MongoProxy) String() string { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()    {}
func (*MongoProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d590dd12f767c61, []int{0}
}
func (m *MongoProxy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MongoProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MongoProxy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MongoProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoProxy.Merge(m, src)
}
func (m *MongoProxy) XXX_Size() int {
	return m.Size()
}
func (m *MongoProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoProxy.DiscardUnknown(m)
}

var xxx_messageInfo_MongoProxy proto.InternalMessageInfo

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *MongoProxy) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *MongoProxy) GetEmitDynamicMetadata() bool {
	if m != nil {
		return m.EmitDynamicMetadata
	}
	return false
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.config.filter.network.mongo_proxy.v2.MongoProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto", fileDescriptor_4d590dd12f767c61)
}

var fileDescriptor_4d590dd12f767c61 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x1c, 0xc5, 0xc9, 0xdd, 0x29, 0x36, 0x37, 0x08, 0x15, 0xb1, 0x1c, 0x58, 0x8a, 0x53, 0xb9, 0x21,
	0x81, 0xba, 0x38, 0x88, 0x43, 0x39, 0x9c, 0x3c, 0x28, 0xdd, 0x9c, 0x4a, 0x6c, 0xd3, 0x12, 0x6c,
	0xfb, 0x2f, 0x69, 0xac, 0xd7, 0xaf, 0xe6, 0xe4, 0x28, 0xb8, 0xf8, 0x11, 0xa4, 0x9b, 0xdf, 0x42,
	0xd2, 0x9c, 0x78, 0xc3, 0x0d, 0xb7, 0xfd, 0x93, 0xdf, 0x7b, 0xef, 0x9f, 0x3c, 0x7c, 0xcb, 0xeb,
	0x0e, 0x7a, 0x9a, 0x42, 0x9d, 0x8b, 0x82, 0xe6, 0xa2, 0x54, 0x5c, 0xd2, 0x9a, 0xab, 0x57, 0x90,
	0xcf, 0xb4, 0x82, 0xba, 0x80, 0xa4, 0x91, 0xb0, 0xe9, 0x69, 0x17, 0xec, 0x1e, 0x49, 0x23, 0x41,
	0x81, 0xbd, 0x1c, 0xdd, 0xc4, 0xb8, 0x89, 0x71, 0x93, 0xad, 0x9b, 0xec, 0xca, 0xbb, 0x60, 0xe1,
	0xef, 0xdb, 0x94, 0xb3, 0x97, 0x52, 0xe9, 0xec, 0x71, 0x30, 0xa9, 0x8b, 0x8b, 0x8e, 0x95, 0x22,
	0x63, 0x8a, 0xd3, 0xbf, 0xc1, 0x80, 0xab, 0x4f, 0x84, 0xf1, 0x5a, 0xa7, 0x46, 0x3a, 0xd4, 0x5e,
	0xe2, 0x79, 0xab, 0x98, 0x4a, 0x1a, 0xc9, 0x73, 0xb1, 0x71, 0x90, 0x87, 0x7c, 0x2b, 0xb4, 0xde,
	0x7e, 0xde, 0xa7, 0x33, 0x39, 0xf1, 0x50, 0x8c, 0x35, 0x8d, 0x46, 0x68, 0x5f, 0x62, 0xcc, 0xd2,
	0x94, 0xb7, 0x6d, 0x52, 0x42, 0xe1, 0x4c, 0xb4, 0x34, 0xb6, 0xcc, 0xcd, 0x03, 0x14, 0xf6, 0x1d,
	0x3e, 0xca, 0x78, 0xc9, 0x7a, 0x67, 0xea, 0x21, 0x7f, 0x1e, 0xf8, 0x64, 0xdf, 0xc7, 0xcc, 0x1b,
	0xbb, 0x80, 0xdc, 0xeb, 0x61, 0xa5, 0xf5, 0xb1, 0xb1, 0xd9, 0x01, 0x3e, 0xe7, 0x95, 0x50, 0x49,
	0xd6, 0xd7, 0xac, 0x12, 0x69, 0x52, 0x71, 0xc5, 0x32, 0xa6, 0x98, 0x33, 0xf3, 0x90, 0x7f, 0x12,
	0x9f, 0x69, 0xb8, 0x32, 0x6c, 0xbd, 0x45, 0xe1, 0xe3, 0xc7, 0xe0, 0xa2, 0xaf, 0xc1, 0x45, 0xdf,
	0x83, 0x8b, 0xf0, 0x8d, 0x00, 0xb3, 0xd4, 0x14, 0x76, 0x78, 0xb1, 0xe1, 0xe9, 0x7f, 0x25, 0x91,
	0xae, 0x29, 0x42, 0x4f, 0xc7, 0x63, 0x5f, 0xd7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0x9c,
	0xc4, 0x94, 0xde, 0x01, 0x00, 0x00,
}

func (m *MongoProxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MongoProxy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MongoProxy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EmitDynamicMetadata {
		i--
		if m.EmitDynamicMetadata {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Delay != nil {
		{
			size, err := m.Delay.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMongoProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AccessLog) > 0 {
		i -= len(m.AccessLog)
		copy(dAtA[i:], m.AccessLog)
		i = encodeVarintMongoProxy(dAtA, i, uint64(len(m.AccessLog)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintMongoProxy(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMongoProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovMongoProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MongoProxy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovMongoProxy(uint64(l))
	}
	l = len(m.AccessLog)
	if l > 0 {
		n += 1 + l + sovMongoProxy(uint64(l))
	}
	if m.Delay != nil {
		l = m.Delay.Size()
		n += 1 + l + sovMongoProxy(uint64(l))
	}
	if m.EmitDynamicMetadata {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMongoProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMongoProxy(x uint64) (n int) {
	return sovMongoProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MongoProxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMongoProxy
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
			return fmt.Errorf("proto: MongoProxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MongoProxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
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
				return ErrInvalidLengthMongoProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessLog", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
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
				return ErrInvalidLengthMongoProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessLog = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
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
				return ErrInvalidLengthMongoProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Delay == nil {
				m.Delay = &v2.FaultDelay{}
			}
			if err := m.Delay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmitDynamicMetadata", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMongoProxy
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
			m.EmitDynamicMetadata = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMongoProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMongoProxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMongoProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMongoProxy
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
					return 0, ErrIntOverflowMongoProxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMongoProxy
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
				return 0, ErrInvalidLengthMongoProxy
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMongoProxy
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMongoProxy
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMongoProxy(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMongoProxy
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMongoProxy = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMongoProxy   = fmt.Errorf("proto: integer overflow")
)