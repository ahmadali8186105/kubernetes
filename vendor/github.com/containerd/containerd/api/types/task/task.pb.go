// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/types/task/task.proto

package task

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	StatusUnknown Status = 0
	StatusCreated Status = 1
	StatusRunning Status = 2
	StatusStopped Status = 3
	StatusPaused  Status = 4
	StatusPausing Status = 5
)

var Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "CREATED",
	2: "RUNNING",
	3: "STOPPED",
	4: "PAUSED",
	5: "PAUSING",
}

var Status_value = map[string]int32{
	"UNKNOWN": 0,
	"CREATED": 1,
	"RUNNING": 2,
	"STOPPED": 3,
	"PAUSED":  4,
	"PAUSING": 5,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_391ef18c8ab0dc16, []int{0}
}

type Process struct {
	ContainerID          string    `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ID                   string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Pid                  uint32    `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Status               Status    `protobuf:"varint,4,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
	Stdin                string    `protobuf:"bytes,5,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout               string    `protobuf:"bytes,6,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               string    `protobuf:"bytes,7,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Terminal             bool      `protobuf:"varint,8,opt,name=terminal,proto3" json:"terminal,omitempty"`
	ExitStatus           uint32    `protobuf:"varint,9,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt             time.Time `protobuf:"bytes,10,opt,name=exited_at,json=exitedAt,proto3,stdtime" json:"exited_at"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Process) Reset()      { *m = Process{} }
func (*Process) ProtoMessage() {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_391ef18c8ab0dc16, []int{0}
}
func (m *Process) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Process.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return m.Size()
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

type ProcessInfo struct {
	// PID is the process ID.
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	// Info contains additional process information.
	//
	// Info varies by platform.
	Info                 *types.Any `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProcessInfo) Reset()      { *m = ProcessInfo{} }
func (*ProcessInfo) ProtoMessage() {}
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_391ef18c8ab0dc16, []int{1}
}
func (m *ProcessInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessInfo.Merge(m, src)
}
func (m *ProcessInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProcessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessInfo proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("containerd.v1.types.Status", Status_name, Status_value)
	proto.RegisterType((*Process)(nil), "containerd.v1.types.Process")
	proto.RegisterType((*ProcessInfo)(nil), "containerd.v1.types.ProcessInfo")
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/types/task/task.proto", fileDescriptor_391ef18c8ab0dc16)
}

var fileDescriptor_391ef18c8ab0dc16 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7d, 0x6e, 0xeb, 0xa6, 0xe7, 0xb6, 0x18, 0x13, 0x55, 0xc6, 0x20, 0xdb, 0xea, 0x64,
	0x31, 0xd8, 0x22, 0xdd, 0xd8, 0xf2, 0x4f, 0xc8, 0x42, 0x72, 0x23, 0x27, 0x11, 0x6c, 0x91, 0x13,
	0x5f, 0xcc, 0xa9, 0xcd, 0x9d, 0x65, 0x9f, 0x81, 0x6c, 0x8c, 0xa8, 0x13, 0x5f, 0xa0, 0x13, 0x7c,
	0x0a, 0x3e, 0x41, 0x46, 0x26, 0xc4, 0x14, 0xa8, 0x3f, 0x09, 0x3a, 0xdb, 0x49, 0x23, 0x60, 0x39,
	0xbd, 0xef, 0xf3, 0x7b, 0xee, 0xbd, 0xf7, 0x1e, 0xf8, 0x22, 0xc6, 0xec, 0x6d, 0x3e, 0x75, 0x66,
	0x74, 0xe1, 0xce, 0x28, 0x61, 0x21, 0x26, 0x28, 0x8d, 0x76, 0xcb, 0x30, 0xc1, 0x2e, 0x5b, 0x26,
	0x28, 0x73, 0x59, 0x98, 0x5d, 0x95, 0x87, 0x93, 0xa4, 0x94, 0x51, 0xf5, 0xd1, 0xbd, 0xcb, 0x79,
	0xf7, 0xdc, 0x29, 0x4d, 0x7a, 0x33, 0xa6, 0x31, 0x2d, 0xb9, 0xcb, 0xab, 0xca, 0xaa, 0x9b, 0x31,
	0xa5, 0xf1, 0x35, 0x72, 0xcb, 0x6e, 0x9a, 0xcf, 0x5d, 0x86, 0x17, 0x28, 0x63, 0xe1, 0x22, 0xa9,
	0x0d, 0x8f, 0xff, 0x36, 0x84, 0x64, 0x59, 0xa1, 0xf3, 0x42, 0x84, 0x87, 0x83, 0x94, 0xce, 0x50,
	0x96, 0xa9, 0x2d, 0x78, 0xbc, 0x7d, 0x74, 0x82, 0x23, 0x0d, 0x58, 0xc0, 0x3e, 0xea, 0x3c, 0x28,
	0xd6, 0xa6, 0xdc, 0xdd, 0xe8, 0x5e, 0x2f, 0x90, 0xb7, 0x26, 0x2f, 0x52, 0xcf, 0xa0, 0x88, 0x23,
	0x4d, 0x2c, 0x9d, 0x52, 0xb1, 0x36, 0x45, 0xaf, 0x17, 0x88, 0x38, 0x52, 0x15, 0xb8, 0x97, 0xe0,
	0x48, 0xdb, 0xb3, 0x80, 0x7d, 0x12, 0xf0, 0x52, 0xbd, 0x80, 0x52, 0xc6, 0x42, 0x96, 0x67, 0xda,
	0xbe, 0x05, 0xec, 0xd3, 0xd6, 0x13, 0xe7, 0x3f, 0x3f, 0x74, 0x86, 0xa5, 0x25, 0xa8, 0xad, 0x6a,
	0x13, 0x1e, 0x64, 0x2c, 0xc2, 0x44, 0x3b, 0xe0, 0x2f, 0x04, 0x55, 0xa3, 0x9e, 0xf1, 0x51, 0x11,
	0xcd, 0x99, 0x26, 0x95, 0x72, 0xdd, 0xd5, 0x3a, 0x4a, 0x53, 0xed, 0x70, 0xab, 0xa3, 0x34, 0x55,
	0x75, 0xd8, 0x60, 0x28, 0x5d, 0x60, 0x12, 0x5e, 0x6b, 0x0d, 0x0b, 0xd8, 0x8d, 0x60, 0xdb, 0xab,
	0x26, 0x94, 0xd1, 0x07, 0xcc, 0x26, 0xf5, 0x6e, 0x47, 0xe5, 0xc2, 0x90, 0x4b, 0xd5, 0x2a, 0x6a,
	0x1b, 0x1e, 0xf1, 0x0e, 0x45, 0x93, 0x90, 0x69, 0xd0, 0x02, 0xb6, 0xdc, 0xd2, 0x9d, 0x2a, 0x50,
	0x67, 0x13, 0xa8, 0x33, 0xda, 0x24, 0xde, 0x69, 0xac, 0xd6, 0xa6, 0xf0, 0xf9, 0x97, 0x09, 0x82,
	0x46, 0x75, 0xad, 0xcd, 0xce, 0x3d, 0x28, 0xd7, 0x19, 0x7b, 0x64, 0x4e, 0x37, 0xd9, 0x80, 0xfb,
	0x6c, 0x6c, 0xb8, 0x8f, 0xc9, 0x9c, 0x96, 0x39, 0xca, 0xad, 0xe6, 0x3f, 0xe3, 0xdb, 0x64, 0x19,
	0x94, 0x8e, 0x67, 0x3f, 0x00, 0x94, 0xea, 0xc5, 0x0c, 0x78, 0x38, 0xf6, 0x5f, 0xf9, 0x97, 0xaf,
	0x7d, 0x45, 0xd0, 0x1f, 0xde, 0xdc, 0x5a, 0x27, 0x15, 0x18, 0x93, 0x2b, 0x42, 0xdf, 0x13, 0xce,
	0xbb, 0x41, 0xbf, 0x3d, 0xea, 0xf7, 0x14, 0xb0, 0xcb, 0xbb, 0x29, 0x0a, 0x19, 0x8a, 0x38, 0x0f,
	0xc6, 0xbe, 0xef, 0xf9, 0x2f, 0x15, 0x71, 0x97, 0x07, 0x39, 0x21, 0x98, 0xc4, 0x9c, 0x0f, 0x47,
	0x97, 0x83, 0x41, 0xbf, 0xa7, 0xec, 0xed, 0xf2, 0x21, 0xa3, 0x49, 0x82, 0x22, 0xf5, 0x29, 0x94,
	0x06, 0xed, 0xf1, 0xb0, 0xdf, 0x53, 0xf6, 0x75, 0xe5, 0xe6, 0xd6, 0x3a, 0xae, 0xf0, 0x20, 0xcc,
	0xb3, 0x6a, 0x3a, 0xa7, 0x7c, 0xfa, 0xc1, 0xee, 0x6d, 0x8e, 0x31, 0x89, 0xf5, 0xd3, 0x4f, 0x5f,
	0x0c, 0xe1, 0xdb, 0x57, 0xa3, 0xfe, 0x4d, 0x47, 0x5b, 0xdd, 0x19, 0xc2, 0xcf, 0x3b, 0x43, 0xf8,
	0x58, 0x18, 0x60, 0x55, 0x18, 0xe0, 0x7b, 0x61, 0x80, 0xdf, 0x85, 0x01, 0xde, 0x08, 0x53, 0xa9,
	0x0c, 0xe2, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x32, 0xd2, 0x86, 0x50, 0x03, 0x00,
	0x00,
}

func (m *Process) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Process) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Process) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExitedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExitedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTask(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	if m.ExitStatus != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.ExitStatus))
		i--
		dAtA[i] = 0x48
	}
	if m.Terminal {
		i--
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.Stderr) > 0 {
		i -= len(m.Stderr)
		copy(dAtA[i:], m.Stderr)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stderr)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Stdout) > 0 {
		i -= len(m.Stdout)
		copy(dAtA[i:], m.Stdout)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdout)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Stdin) > 0 {
		i -= len(m.Stdin)
		copy(dAtA[i:], m.Stdin)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdin)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Pid != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContainerID) > 0 {
		i -= len(m.ContainerID)
		copy(dAtA[i:], m.ContainerID)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ContainerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProcessInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTask(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Pid != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	offset -= sovTask(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Process) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if m.Status != 0 {
		n += 1 + sovTask(uint64(m.Status))
	}
	l = len(m.Stdin)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stdout)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stderr)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Terminal {
		n += 2
	}
	if m.ExitStatus != 0 {
		n += 1 + sovTask(uint64(m.ExitStatus))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExitedAt)
	n += 1 + l + sovTask(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProcessInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTask(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Process) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Process{`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Stdin:` + fmt.Sprintf("%v", this.Stdin) + `,`,
		`Stdout:` + fmt.Sprintf("%v", this.Stdout) + `,`,
		`Stderr:` + fmt.Sprintf("%v", this.Stderr) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`ExitStatus:` + fmt.Sprintf("%v", this.ExitStatus) + `,`,
		`ExitedAt:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ExitedAt), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessInfo{`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Info:` + strings.Replace(fmt.Sprintf("%v", this.Info), "Any", "types.Any", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTask(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Process) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
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
			return fmt.Errorf("proto: Process: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Process: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stderr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stderr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
			m.Terminal = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitStatus", wireType)
			}
			m.ExitStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitStatus |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExitedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTask
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
func (m *ProcessInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
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
			return fmt.Errorf("proto: ProcessInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
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
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &types.Any{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTask
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
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
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
					return 0, ErrIntOverflowTask
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
					return 0, ErrIntOverflowTask
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
				return 0, ErrInvalidLengthTask
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTask
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTask
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTask        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTask = fmt.Errorf("proto: unexpected end of group")
)
