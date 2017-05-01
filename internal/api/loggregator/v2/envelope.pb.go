// Code generated by protoc-gen-go.
// source: envelope.proto
// DO NOT EDIT!

package loggregator_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Log_Type int32

const (
	Log_OUT Log_Type = 0
	Log_ERR Log_Type = 1
)

var Log_Type_name = map[int32]string{
	0: "OUT",
	1: "ERR",
}
var Log_Type_value = map[string]int32{
	"OUT": 0,
	"ERR": 1,
}

func (x Log_Type) String() string {
	return proto.EnumName(Log_Type_name, int32(x))
}
func (Log_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type Envelope struct {
	Timestamp  int64             `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	SourceId   string            `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	InstanceId string            `protobuf:"bytes,8,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	Tags       map[string]*Value `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Message:
	//	*Envelope_Log
	//	*Envelope_Counter
	//	*Envelope_Gauge
	//	*Envelope_Timer
	Message isEnvelope_Message `protobuf_oneof:"message"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isEnvelope_Message interface {
	isEnvelope_Message()
}

type Envelope_Log struct {
	Log *Log `protobuf:"bytes,4,opt,name=log,oneof"`
}
type Envelope_Counter struct {
	Counter *Counter `protobuf:"bytes,5,opt,name=counter,oneof"`
}
type Envelope_Gauge struct {
	Gauge *Gauge `protobuf:"bytes,6,opt,name=gauge,oneof"`
}
type Envelope_Timer struct {
	Timer *Timer `protobuf:"bytes,7,opt,name=timer,oneof"`
}

func (*Envelope_Log) isEnvelope_Message()     {}
func (*Envelope_Counter) isEnvelope_Message() {}
func (*Envelope_Gauge) isEnvelope_Message()   {}
func (*Envelope_Timer) isEnvelope_Message()   {}

func (m *Envelope) GetMessage() isEnvelope_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Envelope) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Envelope) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Envelope) GetTags() map[string]*Value {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Envelope) GetLog() *Log {
	if x, ok := m.GetMessage().(*Envelope_Log); ok {
		return x.Log
	}
	return nil
}

func (m *Envelope) GetCounter() *Counter {
	if x, ok := m.GetMessage().(*Envelope_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *Envelope) GetGauge() *Gauge {
	if x, ok := m.GetMessage().(*Envelope_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (m *Envelope) GetTimer() *Timer {
	if x, ok := m.GetMessage().(*Envelope_Timer); ok {
		return x.Timer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Envelope) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Envelope_OneofMarshaler, _Envelope_OneofUnmarshaler, _Envelope_OneofSizer, []interface{}{
		(*Envelope_Log)(nil),
		(*Envelope_Counter)(nil),
		(*Envelope_Gauge)(nil),
		(*Envelope_Timer)(nil),
	}
}

func _Envelope_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case *Envelope_Counter:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Counter); err != nil {
			return err
		}
	case *Envelope_Gauge:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gauge); err != nil {
			return err
		}
	case *Envelope_Timer:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Timer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Envelope.Message has unexpected type %T", x)
	}
	return nil
}

func _Envelope_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Envelope)
	switch tag {
	case 4: // message.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Log)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Log{msg}
		return true, err
	case 5: // message.counter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Counter)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Counter{msg}
		return true, err
	case 6: // message.gauge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Gauge)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Gauge{msg}
		return true, err
	case 7: // message.timer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Timer)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Timer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Envelope_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		s := proto.Size(x.Log)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Counter:
		s := proto.Size(x.Counter)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Gauge:
		s := proto.Size(x.Gauge)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Timer:
		s := proto.Size(x.Timer)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Value struct {
	// Types that are valid to be assigned to Data:
	//	*Value_Text
	//	*Value_Integer
	//	*Value_Decimal
	Data isValue_Data `protobuf_oneof:"data"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isValue_Data interface {
	isValue_Data()
}

type Value_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,oneof"`
}
type Value_Integer struct {
	Integer int64 `protobuf:"varint,2,opt,name=integer,oneof"`
}
type Value_Decimal struct {
	Decimal float64 `protobuf:"fixed64,3,opt,name=decimal,oneof"`
}

func (*Value_Text) isValue_Data()    {}
func (*Value_Integer) isValue_Data() {}
func (*Value_Decimal) isValue_Data() {}

func (m *Value) GetData() isValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Value) GetText() string {
	if x, ok := m.GetData().(*Value_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Value) GetInteger() int64 {
	if x, ok := m.GetData().(*Value_Integer); ok {
		return x.Integer
	}
	return 0
}

func (m *Value) GetDecimal() float64 {
	if x, ok := m.GetData().(*Value_Decimal); ok {
		return x.Decimal
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_Text)(nil),
		(*Value_Integer)(nil),
		(*Value_Decimal)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *Value_Integer:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Integer))
	case *Value_Decimal:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Decimal))
	case nil:
	default:
		return fmt.Errorf("Value.Data has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // data.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &Value_Text{x}
		return true, err
	case 2: // data.integer
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Value_Integer{int64(x)}
		return true, err
	case 3: // data.decimal
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Data = &Value_Decimal{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *Value_Integer:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Integer))
	case *Value_Decimal:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Log struct {
	Payload []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type    Log_Type `protobuf:"varint,2,opt,name=type,enum=loggregator.v2.Log_Type" json:"type,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Log) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Log) GetType() Log_Type {
	if m != nil {
		return m.Type
	}
	return Log_OUT
}

type Counter struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Counter_Delta
	//	*Counter_Total
	Value isCounter_Value `protobuf_oneof:"value"`
}

func (m *Counter) Reset()                    { *m = Counter{} }
func (m *Counter) String() string            { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()               {}
func (*Counter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type isCounter_Value interface {
	isCounter_Value()
}

type Counter_Delta struct {
	Delta uint64 `protobuf:"varint,2,opt,name=delta,oneof"`
}
type Counter_Total struct {
	Total uint64 `protobuf:"varint,3,opt,name=total,oneof"`
}

func (*Counter_Delta) isCounter_Value() {}
func (*Counter_Total) isCounter_Value() {}

func (m *Counter) GetValue() isCounter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Counter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Counter) GetDelta() uint64 {
	if x, ok := m.GetValue().(*Counter_Delta); ok {
		return x.Delta
	}
	return 0
}

func (m *Counter) GetTotal() uint64 {
	if x, ok := m.GetValue().(*Counter_Total); ok {
		return x.Total
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Counter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Counter_OneofMarshaler, _Counter_OneofUnmarshaler, _Counter_OneofSizer, []interface{}{
		(*Counter_Delta)(nil),
		(*Counter_Total)(nil),
	}
}

func _Counter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Counter)
	// value
	switch x := m.Value.(type) {
	case *Counter_Delta:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Delta))
	case *Counter_Total:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Total))
	case nil:
	default:
		return fmt.Errorf("Counter.Value has unexpected type %T", x)
	}
	return nil
}

func _Counter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Counter)
	switch tag {
	case 2: // value.delta
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Counter_Delta{x}
		return true, err
	case 3: // value.total
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Counter_Total{x}
		return true, err
	default:
		return false, nil
	}
}

func _Counter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Counter)
	// value
	switch x := m.Value.(type) {
	case *Counter_Delta:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Delta))
	case *Counter_Total:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Total))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Gauge struct {
	Metrics map[string]*GaugeValue `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Gauge) Reset()                    { *m = Gauge{} }
func (m *Gauge) String() string            { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()               {}
func (*Gauge) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Gauge) GetMetrics() map[string]*GaugeValue {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type GaugeValue struct {
	Unit  string  `protobuf:"bytes,1,opt,name=unit" json:"unit,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *GaugeValue) Reset()                    { *m = GaugeValue{} }
func (m *GaugeValue) String() string            { return proto.CompactTextString(m) }
func (*GaugeValue) ProtoMessage()               {}
func (*GaugeValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GaugeValue) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *GaugeValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Timer struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Start int64  `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	Stop  int64  `protobuf:"varint,3,opt,name=stop" json:"stop,omitempty"`
}

func (m *Timer) Reset()                    { *m = Timer{} }
func (m *Timer) String() string            { return proto.CompactTextString(m) }
func (*Timer) ProtoMessage()               {}
func (*Timer) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *Timer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Timer) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Timer) GetStop() int64 {
	if m != nil {
		return m.Stop
	}
	return 0
}

func init() {
	proto.RegisterType((*Envelope)(nil), "loggregator.v2.Envelope")
	proto.RegisterType((*Value)(nil), "loggregator.v2.Value")
	proto.RegisterType((*Log)(nil), "loggregator.v2.Log")
	proto.RegisterType((*Counter)(nil), "loggregator.v2.Counter")
	proto.RegisterType((*Gauge)(nil), "loggregator.v2.Gauge")
	proto.RegisterType((*GaugeValue)(nil), "loggregator.v2.GaugeValue")
	proto.RegisterType((*Timer)(nil), "loggregator.v2.Timer")
	proto.RegisterEnum("loggregator.v2.Log_Type", Log_Type_name, Log_Type_value)
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x94, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x9b, 0xdb, 0xa4, 0x69, 0xa6, 0x47, 0x29, 0x6b, 0xd5, 0xa5, 0x0a, 0x96, 0xbc, 0x58,
	0x50, 0x83, 0xf4, 0xe0, 0x10, 0xf1, 0xa9, 0x52, 0xec, 0xc1, 0xa9, 0xb0, 0xd4, 0xbe, 0x89, 0xac,
	0xcd, 0xb2, 0x04, 0xd3, 0x6c, 0x48, 0xb6, 0xc5, 0x7e, 0x18, 0x3f, 0x86, 0xdf, 0x4f, 0x76, 0x36,
	0xf1, 0xee, 0x4a, 0xde, 0x66, 0xe6, 0xff, 0x9b, 0x69, 0xe6, 0x3f, 0x4b, 0x61, 0x24, 0x8b, 0xa3,
	0xcc, 0x75, 0x29, 0x93, 0xb2, 0xd2, 0x46, 0xd3, 0x51, 0xae, 0x95, 0xaa, 0xa4, 0x12, 0x46, 0x57,
	0xc9, 0x71, 0x11, 0xff, 0x25, 0x30, 0x58, 0x35, 0x08, 0x7d, 0x0e, 0x91, 0xc9, 0xf6, 0xb2, 0x36,
	0x62, 0x5f, 0x32, 0x6f, 0xe6, 0xcd, 0x09, 0xbf, 0x2b, 0xd0, 0x67, 0x10, 0xd5, 0xfa, 0x50, 0xed,
	0xe4, 0x8f, 0x2c, 0x65, 0x17, 0x33, 0x6f, 0x1e, 0xf1, 0x81, 0x2b, 0xdc, 0xa4, 0xf4, 0x05, 0x0c,
	0xb3, 0xa2, 0x36, 0xa2, 0x70, 0xf2, 0x00, 0x65, 0x68, 0x4b, 0x37, 0x29, 0xbd, 0x06, 0xdf, 0x08,
	0x55, 0x33, 0x32, 0x23, 0xf3, 0xe1, 0x22, 0x4e, 0x1e, 0x7e, 0x47, 0xd2, 0x7e, 0x43, 0xb2, 0x11,
	0xaa, 0x5e, 0x15, 0xa6, 0x3a, 0x71, 0xe4, 0xe9, 0x4b, 0x20, 0xb9, 0x56, 0xcc, 0x9f, 0x79, 0xf3,
	0xe1, 0xe2, 0xd1, 0x79, 0xdb, 0xad, 0x56, 0xeb, 0x1e, 0xb7, 0x04, 0xbd, 0x82, 0x70, 0xa7, 0x0f,
	0x85, 0x91, 0x15, 0x0b, 0x10, 0x7e, 0x7a, 0x0e, 0x7f, 0x74, 0xf2, 0xba, 0xc7, 0x5b, 0x92, 0xbe,
	0x81, 0x40, 0x89, 0x83, 0x92, 0xac, 0x8f, 0x2d, 0x8f, 0xcf, 0x5b, 0x3e, 0x59, 0x71, 0xdd, 0xe3,
	0x8e, 0xb2, 0xb8, 0xf5, 0xa3, 0x62, 0x61, 0x37, 0xbe, 0xb1, 0xa2, 0xc5, 0x91, 0x9a, 0x7e, 0x81,
	0xe8, 0xff, 0x3a, 0x74, 0x0c, 0xe4, 0x97, 0x3c, 0xa1, 0xad, 0x11, 0xb7, 0x21, 0x7d, 0x05, 0xc1,
	0x51, 0xe4, 0x07, 0x89, 0x66, 0x76, 0x4c, 0xdb, 0x5a, 0x91, 0x3b, 0xe6, 0xfd, 0xc5, 0x3b, 0x6f,
	0x19, 0x41, 0xb8, 0x97, 0x75, 0x2d, 0x94, 0x8c, 0xbf, 0x43, 0x80, 0x32, 0x9d, 0x80, 0x6f, 0xe4,
	0x6f, 0xe3, 0xe6, 0xae, 0x7b, 0x1c, 0x33, 0x3a, 0x85, 0x30, 0x2b, 0x8c, 0x54, 0xb2, 0xc2, 0xe1,
	0xc4, 0xee, 0xdc, 0x14, 0xac, 0x96, 0xca, 0x5d, 0xb6, 0x17, 0x39, 0x23, 0x33, 0x6f, 0xee, 0x59,
	0xad, 0x29, 0x2c, 0xfb, 0xe0, 0xa7, 0xc2, 0x88, 0x58, 0x01, 0xb9, 0xd5, 0x8a, 0x32, 0x08, 0x4b,
	0x71, 0xca, 0xb5, 0x48, 0x71, 0xfe, 0x25, 0x6f, 0x53, 0xfa, 0x1a, 0x7c, 0x73, 0x2a, 0xdd, 0xa7,
	0x8f, 0x16, 0xac, 0xe3, 0x2e, 0xc9, 0xe6, 0x54, 0x4a, 0x8e, 0x54, 0xcc, 0xc0, 0xb7, 0x19, 0x0d,
	0x81, 0x7c, 0xfd, 0xb6, 0x19, 0xf7, 0x6c, 0xb0, 0xe2, 0x7c, 0xec, 0xc5, 0x5b, 0x08, 0x9b, 0xb3,
	0x50, 0x0a, 0x7e, 0x21, 0xf6, 0xb2, 0x71, 0x08, 0x63, 0xfa, 0x04, 0x82, 0x54, 0xe6, 0x46, 0xe0,
	0xef, 0xf8, 0xd6, 0x59, 0x4c, 0x6d, 0xdd, 0x68, 0xd3, 0x6c, 0x80, 0x75, 0x4c, 0x97, 0x61, 0x63,
	0x69, 0xfc, 0xc7, 0x83, 0x00, 0x8f, 0x47, 0x3f, 0x58, 0xd3, 0x4c, 0x95, 0xed, 0x6a, 0xe6, 0x75,
	0xbf, 0x3d, 0xe4, 0x92, 0xcf, 0x0e, 0x72, 0x6f, 0xaf, 0x6d, 0x99, 0x6e, 0xe1, 0xf2, 0xbe, 0xd0,
	0x71, 0xc5, 0xb7, 0x0f, 0xaf, 0x38, 0xed, 0x9c, 0x7e, 0x7e, 0xca, 0xf8, 0x1a, 0xe0, 0x4e, 0xb0,
	0xab, 0x1f, 0x8a, 0xcc, 0xb4, 0xab, 0xdb, 0x98, 0x4e, 0xee, 0xcf, 0xf5, 0x9a, 0xde, 0x78, 0x05,
	0x01, 0x3e, 0xb2, 0x4e, 0xb7, 0x26, 0x10, 0xd4, 0x46, 0x54, 0xc6, 0xdd, 0x9c, 0xbb, 0xc4, 0x92,
	0xb5, 0xd1, 0x25, 0x5a, 0x45, 0x38, 0xc6, 0x3f, 0xfb, 0xf8, 0x6f, 0x70, 0xf5, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0xa3, 0x8d, 0xd7, 0x1f, 0x04, 0x00, 0x00,
}
