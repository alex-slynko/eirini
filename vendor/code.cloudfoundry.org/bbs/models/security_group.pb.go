// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security_group.proto

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PortRange struct {
	Start uint32 `protobuf:"varint,1,opt,name=start" json:"start"`
	End   uint32 `protobuf:"varint,2,opt,name=end" json:"end"`
}

func (m *PortRange) Reset()      { *m = PortRange{} }
func (*PortRange) ProtoMessage() {}
func (*PortRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_security_group_2b6a202766ce71f4, []int{0}
}
func (m *PortRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PortRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PortRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PortRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortRange.Merge(dst, src)
}
func (m *PortRange) XXX_Size() int {
	return m.Size()
}
func (m *PortRange) XXX_DiscardUnknown() {
	xxx_messageInfo_PortRange.DiscardUnknown(m)
}

var xxx_messageInfo_PortRange proto.InternalMessageInfo

func (m *PortRange) GetStart() uint32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *PortRange) GetEnd() uint32 {
	if m != nil {
		return m.End
	}
	return 0
}

type ICMPInfo struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type"`
	Code int32 `protobuf:"varint,2,opt,name=code" json:"code"`
}

func (m *ICMPInfo) Reset()      { *m = ICMPInfo{} }
func (*ICMPInfo) ProtoMessage() {}
func (*ICMPInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_security_group_2b6a202766ce71f4, []int{1}
}
func (m *ICMPInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ICMPInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ICMPInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ICMPInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ICMPInfo.Merge(dst, src)
}
func (m *ICMPInfo) XXX_Size() int {
	return m.Size()
}
func (m *ICMPInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ICMPInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ICMPInfo proto.InternalMessageInfo

func (m *ICMPInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ICMPInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type SecurityGroupRule struct {
	Protocol     string     `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	Destinations []string   `protobuf:"bytes,2,rep,name=destinations" json:"destinations,omitempty"`
	Ports        []uint32   `protobuf:"varint,3,rep,name=ports" json:"ports,omitempty"`
	PortRange    *PortRange `protobuf:"bytes,4,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
	IcmpInfo     *ICMPInfo  `protobuf:"bytes,5,opt,name=icmp_info,json=icmpInfo" json:"icmp_info,omitempty"`
	Log          bool       `protobuf:"varint,6,opt,name=log" json:"log"`
	Annotations  []string   `protobuf:"bytes,7,rep,name=annotations" json:"annotations,omitempty"`
}

func (m *SecurityGroupRule) Reset()      { *m = SecurityGroupRule{} }
func (*SecurityGroupRule) ProtoMessage() {}
func (*SecurityGroupRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_security_group_2b6a202766ce71f4, []int{2}
}
func (m *SecurityGroupRule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SecurityGroupRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SecurityGroupRule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SecurityGroupRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityGroupRule.Merge(dst, src)
}
func (m *SecurityGroupRule) XXX_Size() int {
	return m.Size()
}
func (m *SecurityGroupRule) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityGroupRule.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityGroupRule proto.InternalMessageInfo

func (m *SecurityGroupRule) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *SecurityGroupRule) GetDestinations() []string {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *SecurityGroupRule) GetPorts() []uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *SecurityGroupRule) GetPortRange() *PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *SecurityGroupRule) GetIcmpInfo() *ICMPInfo {
	if m != nil {
		return m.IcmpInfo
	}
	return nil
}

func (m *SecurityGroupRule) GetLog() bool {
	if m != nil {
		return m.Log
	}
	return false
}

func (m *SecurityGroupRule) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func init() {
	proto.RegisterType((*PortRange)(nil), "models.PortRange")
	proto.RegisterType((*ICMPInfo)(nil), "models.ICMPInfo")
	proto.RegisterType((*SecurityGroupRule)(nil), "models.SecurityGroupRule")
}
func (this *PortRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PortRange)
	if !ok {
		that2, ok := that.(PortRange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Start != that1.Start {
		return false
	}
	if this.End != that1.End {
		return false
	}
	return true
}
func (this *ICMPInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ICMPInfo)
	if !ok {
		that2, ok := that.(ICMPInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	return true
}
func (this *SecurityGroupRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SecurityGroupRule)
	if !ok {
		that2, ok := that.(SecurityGroupRule)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if len(this.Destinations) != len(that1.Destinations) {
		return false
	}
	for i := range this.Destinations {
		if this.Destinations[i] != that1.Destinations[i] {
			return false
		}
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if this.Ports[i] != that1.Ports[i] {
			return false
		}
	}
	if !this.PortRange.Equal(that1.PortRange) {
		return false
	}
	if !this.IcmpInfo.Equal(that1.IcmpInfo) {
		return false
	}
	if this.Log != that1.Log {
		return false
	}
	if len(this.Annotations) != len(that1.Annotations) {
		return false
	}
	for i := range this.Annotations {
		if this.Annotations[i] != that1.Annotations[i] {
			return false
		}
	}
	return true
}
func (this *PortRange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.PortRange{")
	s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	s = append(s, "End: "+fmt.Sprintf("%#v", this.End)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ICMPInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.ICMPInfo{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SecurityGroupRule) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&models.SecurityGroupRule{")
	s = append(s, "Protocol: "+fmt.Sprintf("%#v", this.Protocol)+",\n")
	if this.Destinations != nil {
		s = append(s, "Destinations: "+fmt.Sprintf("%#v", this.Destinations)+",\n")
	}
	if this.Ports != nil {
		s = append(s, "Ports: "+fmt.Sprintf("%#v", this.Ports)+",\n")
	}
	if this.PortRange != nil {
		s = append(s, "PortRange: "+fmt.Sprintf("%#v", this.PortRange)+",\n")
	}
	if this.IcmpInfo != nil {
		s = append(s, "IcmpInfo: "+fmt.Sprintf("%#v", this.IcmpInfo)+",\n")
	}
	s = append(s, "Log: "+fmt.Sprintf("%#v", this.Log)+",\n")
	if this.Annotations != nil {
		s = append(s, "Annotations: "+fmt.Sprintf("%#v", this.Annotations)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSecurityGroup(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PortRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PortRange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintSecurityGroup(dAtA, i, uint64(m.Start))
	dAtA[i] = 0x10
	i++
	i = encodeVarintSecurityGroup(dAtA, i, uint64(m.End))
	return i, nil
}

func (m *ICMPInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ICMPInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintSecurityGroup(dAtA, i, uint64(m.Type))
	dAtA[i] = 0x10
	i++
	i = encodeVarintSecurityGroup(dAtA, i, uint64(m.Code))
	return i, nil
}

func (m *SecurityGroupRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SecurityGroupRule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintSecurityGroup(dAtA, i, uint64(len(m.Protocol)))
	i += copy(dAtA[i:], m.Protocol)
	if len(m.Destinations) > 0 {
		for _, s := range m.Destinations {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Ports) > 0 {
		for _, num := range m.Ports {
			dAtA[i] = 0x18
			i++
			i = encodeVarintSecurityGroup(dAtA, i, uint64(num))
		}
	}
	if m.PortRange != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSecurityGroup(dAtA, i, uint64(m.PortRange.Size()))
		n1, err := m.PortRange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.IcmpInfo != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSecurityGroup(dAtA, i, uint64(m.IcmpInfo.Size()))
		n2, err := m.IcmpInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	dAtA[i] = 0x30
	i++
	if m.Log {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if len(m.Annotations) > 0 {
		for _, s := range m.Annotations {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintSecurityGroup(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PortRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovSecurityGroup(uint64(m.Start))
	n += 1 + sovSecurityGroup(uint64(m.End))
	return n
}

func (m *ICMPInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovSecurityGroup(uint64(m.Type))
	n += 1 + sovSecurityGroup(uint64(m.Code))
	return n
}

func (m *SecurityGroupRule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Protocol)
	n += 1 + l + sovSecurityGroup(uint64(l))
	if len(m.Destinations) > 0 {
		for _, s := range m.Destinations {
			l = len(s)
			n += 1 + l + sovSecurityGroup(uint64(l))
		}
	}
	if len(m.Ports) > 0 {
		for _, e := range m.Ports {
			n += 1 + sovSecurityGroup(uint64(e))
		}
	}
	if m.PortRange != nil {
		l = m.PortRange.Size()
		n += 1 + l + sovSecurityGroup(uint64(l))
	}
	if m.IcmpInfo != nil {
		l = m.IcmpInfo.Size()
		n += 1 + l + sovSecurityGroup(uint64(l))
	}
	n += 2
	if len(m.Annotations) > 0 {
		for _, s := range m.Annotations {
			l = len(s)
			n += 1 + l + sovSecurityGroup(uint64(l))
		}
	}
	return n
}

func sovSecurityGroup(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSecurityGroup(x uint64) (n int) {
	return sovSecurityGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PortRange) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PortRange{`,
		`Start:` + fmt.Sprintf("%v", this.Start) + `,`,
		`End:` + fmt.Sprintf("%v", this.End) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ICMPInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ICMPInfo{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SecurityGroupRule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SecurityGroupRule{`,
		`Protocol:` + fmt.Sprintf("%v", this.Protocol) + `,`,
		`Destinations:` + fmt.Sprintf("%v", this.Destinations) + `,`,
		`Ports:` + fmt.Sprintf("%v", this.Ports) + `,`,
		`PortRange:` + strings.Replace(fmt.Sprintf("%v", this.PortRange), "PortRange", "PortRange", 1) + `,`,
		`IcmpInfo:` + strings.Replace(fmt.Sprintf("%v", this.IcmpInfo), "ICMPInfo", "ICMPInfo", 1) + `,`,
		`Log:` + fmt.Sprintf("%v", this.Log) + `,`,
		`Annotations:` + fmt.Sprintf("%v", this.Annotations) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSecurityGroup(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PortRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurityGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PortRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PortRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSecurityGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurityGroup
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
func (m *ICMPInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurityGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ICMPInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ICMPInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSecurityGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurityGroup
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
func (m *SecurityGroupRule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSecurityGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SecurityGroupRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SecurityGroupRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destinations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destinations = append(m.Destinations, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSecurityGroup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ports = append(m.Ports, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSecurityGroup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSecurityGroup
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ports) == 0 {
					m.Ports = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSecurityGroup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ports = append(m.Ports, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PortRange == nil {
				m.PortRange = &PortRange{}
			}
			if err := m.PortRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcmpInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IcmpInfo == nil {
				m.IcmpInfo = &ICMPInfo{}
			}
			if err := m.IcmpInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Log = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSecurityGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSecurityGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Annotations = append(m.Annotations, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSecurityGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSecurityGroup
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
func skipSecurityGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSecurityGroup
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
					return 0, ErrIntOverflowSecurityGroup
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
					return 0, ErrIntOverflowSecurityGroup
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSecurityGroup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSecurityGroup
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
				next, err := skipSecurityGroup(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthSecurityGroup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSecurityGroup   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("security_group.proto", fileDescriptor_security_group_2b6a202766ce71f4)
}

var fileDescriptor_security_group_2b6a202766ce71f4 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3d, 0x6e, 0xdb, 0x30,
	0x14, 0xc7, 0x45, 0xcb, 0x4e, 0x2d, 0xa6, 0x01, 0x1a, 0x22, 0x0d, 0x58, 0x0f, 0xb4, 0xa0, 0x49,
	0x43, 0xa3, 0x14, 0x45, 0xd1, 0xa5, 0x40, 0x0b, 0xa8, 0x43, 0x91, 0xa1, 0x40, 0xc0, 0x1e, 0xc0,
	0x50, 0x24, 0x5a, 0x15, 0x20, 0xf1, 0x09, 0x12, 0x35, 0x78, 0xeb, 0x11, 0x7a, 0x8c, 0xde, 0xa1,
	0x17, 0xc8, 0xe8, 0x31, 0x93, 0x51, 0xcb, 0x4b, 0x91, 0x29, 0x47, 0x28, 0x48, 0x59, 0x82, 0xba,
	0xf1, 0xfd, 0x3f, 0x88, 0xf7, 0x7e, 0xf8, 0xa2, 0x16, 0x71, 0x53, 0x65, 0x6a, 0xb3, 0x4a, 0x2b,
	0x68, 0xca, 0xa0, 0xac, 0x40, 0x01, 0x39, 0x29, 0x20, 0x11, 0x79, 0xbd, 0xb8, 0x4a, 0x33, 0xf5,
	0xbd, 0xb9, 0x0b, 0x62, 0x28, 0xae, 0x53, 0x48, 0xe1, 0xda, 0xd8, 0x77, 0xcd, 0xda, 0x4c, 0x66,
	0x30, 0xaf, 0xae, 0xe6, 0x7d, 0xc2, 0xce, 0x2d, 0x54, 0x8a, 0x47, 0x32, 0x15, 0x64, 0x81, 0x67,
	0xb5, 0x8a, 0x2a, 0x45, 0x91, 0x8b, 0xfc, 0xb3, 0x70, 0x7a, 0xbf, 0x5b, 0x5a, 0xbc, 0x93, 0xc8,
	0x25, 0xb6, 0x85, 0x4c, 0xe8, 0x64, 0xe4, 0x68, 0xc1, 0xfb, 0x88, 0xe7, 0x37, 0x9f, 0xbf, 0xde,
	0xde, 0xc8, 0x35, 0x10, 0x8a, 0xa7, 0x6a, 0x53, 0x0a, 0x53, 0x9f, 0x1d, 0x43, 0x46, 0xd1, 0x4e,
	0x0c, 0x89, 0x30, 0xf5, 0xc1, 0xd1, 0x8a, 0xf7, 0x7b, 0x82, 0xcf, 0xbf, 0x1d, 0x0f, 0xfa, 0xa2,
	0xef, 0xe1, 0x4d, 0x2e, 0xc8, 0x7b, 0x3c, 0x37, 0xfb, 0xc5, 0x90, 0x9b, 0xdf, 0x9c, 0x70, 0xa1,
	0x3b, 0x8f, 0xbb, 0x25, 0xe9, 0xf5, 0xd7, 0x50, 0x64, 0x4a, 0x14, 0xa5, 0xda, 0xf0, 0x21, 0x4b,
	0x3c, 0xfc, 0x3c, 0x11, 0xb5, 0xca, 0x64, 0xa4, 0x32, 0x90, 0x35, 0x9d, 0xb8, 0xb6, 0xef, 0xf0,
	0xff, 0x34, 0x72, 0x81, 0x67, 0x25, 0x54, 0xaa, 0xa6, 0xb6, 0x6b, 0xfb, 0x67, 0xbc, 0x1b, 0xc8,
	0x1b, 0x8c, 0xf5, 0x63, 0x55, 0x69, 0x12, 0x74, 0xea, 0x22, 0xff, 0xf4, 0xed, 0x79, 0xd0, 0x41,
	0x0d, 0x06, 0x44, 0xdc, 0x29, 0x07, 0x5a, 0x57, 0xd8, 0xc9, 0xe2, 0xa2, 0x5c, 0x65, 0x72, 0x0d,
	0x74, 0x66, 0x0a, 0x2f, 0xfa, 0x42, 0x8f, 0x84, 0xcf, 0x75, 0xc4, 0xc0, 0xb9, 0xc4, 0x76, 0x0e,
	0x29, 0x3d, 0x71, 0x91, 0x3f, 0xef, 0x01, 0xe6, 0x90, 0x92, 0x0f, 0xf8, 0x34, 0x92, 0x12, 0xd4,
	0x71, 0xe3, 0x67, 0x7a, 0xe3, 0xf0, 0xd5, 0xe3, 0x6e, 0xf9, 0x72, 0x24, 0x8f, 0x8e, 0x1d, 0xa7,
	0xc3, 0x77, 0xdb, 0x3d, 0xb3, 0x1e, 0xf6, 0xcc, 0x7a, 0xda, 0x33, 0xf4, 0xa3, 0x65, 0xe8, 0x57,
	0xcb, 0xd0, 0x7d, 0xcb, 0xd0, 0xb6, 0x65, 0xe8, 0x4f, 0xcb, 0xd0, 0xdf, 0x96, 0x59, 0x4f, 0x2d,
	0x43, 0x3f, 0x0f, 0xcc, 0xda, 0x1e, 0x98, 0xf5, 0x70, 0x60, 0xd6, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x84, 0xdb, 0xa4, 0xfe, 0x42, 0x02, 0x00, 0x00,
}
