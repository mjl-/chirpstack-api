// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/common.proto

package external

import (
	fmt "fmt"
	common "github.com/brocaar/chirpstack-api/go/common"
	gw "github.com/brocaar/chirpstack-api/go/gw"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}

var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}

func (RXWindow) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b2cf469f52d369cc, []int{0}
}

type UplinkFrameLog struct {
	// TX information of the uplink.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX information of the uplink.
	RxInfo []*gw.UplinkRXInfo `protobuf:"bytes,2,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson       string   `protobuf:"bytes,3,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UplinkFrameLog) Reset()         { *m = UplinkFrameLog{} }
func (m *UplinkFrameLog) String() string { return proto.CompactTextString(m) }
func (*UplinkFrameLog) ProtoMessage()    {}
func (*UplinkFrameLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2cf469f52d369cc, []int{0}
}

func (m *UplinkFrameLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkFrameLog.Unmarshal(m, b)
}
func (m *UplinkFrameLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkFrameLog.Marshal(b, m, deterministic)
}
func (m *UplinkFrameLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkFrameLog.Merge(m, src)
}
func (m *UplinkFrameLog) XXX_Size() int {
	return xxx_messageInfo_UplinkFrameLog.Size(m)
}
func (m *UplinkFrameLog) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkFrameLog.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkFrameLog proto.InternalMessageInfo

func (m *UplinkFrameLog) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *UplinkFrameLog) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

func (m *UplinkFrameLog) GetPhyPayloadJson() string {
	if m != nil {
		return m.PhyPayloadJson
	}
	return ""
}

type DownlinkFrameLog struct {
	// TX information of the downlink.
	TxInfo *gw.DownlinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson       string   `protobuf:"bytes,2,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkFrameLog) Reset()         { *m = DownlinkFrameLog{} }
func (m *DownlinkFrameLog) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrameLog) ProtoMessage()    {}
func (*DownlinkFrameLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2cf469f52d369cc, []int{1}
}

func (m *DownlinkFrameLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrameLog.Unmarshal(m, b)
}
func (m *DownlinkFrameLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrameLog.Marshal(b, m, deterministic)
}
func (m *DownlinkFrameLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrameLog.Merge(m, src)
}
func (m *DownlinkFrameLog) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrameLog.Size(m)
}
func (m *DownlinkFrameLog) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrameLog.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrameLog proto.InternalMessageInfo

func (m *DownlinkFrameLog) GetTxInfo() *gw.DownlinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *DownlinkFrameLog) GetPhyPayloadJson() string {
	if m != nil {
		return m.PhyPayloadJson
	}
	return ""
}

// This is a copy of gw.UplinkRXInfo with the only change that the
// gateway_id is of type string so that we can return it as HEX encoded
// instead of base64.
type UplinkRXInfo struct {
	// Gateway ID.
	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayID,proto3" json:"gateway_id,omitempty"`
	// Uplink ID (UUID).
	UplinkId string `protobuf:"bytes,16,opt,name=uplink_id,json=uplinkID,proto3" json:"uplink_id,omitempty"`
	// RX time (only set when the gateway has a GPS module).
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// RX time since GPS epoch (only set when the gateway has a GPS module).
	TimeSinceGpsEpoch *duration.Duration `protobuf:"bytes,3,opt,name=time_since_gps_epoch,json=timeSinceGpsEpoch,proto3" json:"time_since_gps_epoch,omitempty"`
	// RSSI.
	Rssi int32 `protobuf:"varint,5,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// LoRa SNR.
	LoraSnr float64 `protobuf:"fixed64,6,opt,name=lora_snr,json=loraSnr,proto3" json:"lora_snr,omitempty"`
	// Channel.
	Channel uint32 `protobuf:"varint,7,opt,name=channel,proto3" json:"channel,omitempty"`
	// RF Chain.
	RfChain uint32 `protobuf:"varint,8,opt,name=rf_chain,json=rfChain,proto3" json:"rf_chain,omitempty"`
	// Board.
	Board uint32 `protobuf:"varint,9,opt,name=board,proto3" json:"board,omitempty"`
	// Antenna.
	Antenna uint32 `protobuf:"varint,10,opt,name=antenna,proto3" json:"antenna,omitempty"`
	// Location.
	Location *common.Location `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	// Fine-timestamp type.
	FineTimestampType gw.FineTimestampType `protobuf:"varint,12,opt,name=fine_timestamp_type,json=fineTimestampType,proto3,enum=gw.FineTimestampType" json:"fine_timestamp_type,omitempty"`
	// Fine-timestamp data.
	//
	// Types that are valid to be assigned to FineTimestamp:
	//	*UplinkRXInfo_EncryptedFineTimestamp
	//	*UplinkRXInfo_PlainFineTimestamp
	FineTimestamp isUplinkRXInfo_FineTimestamp `protobuf_oneof:"fine_timestamp"`
	// Gateway specific context.
	Context              []byte   `protobuf:"bytes,15,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UplinkRXInfo) Reset()         { *m = UplinkRXInfo{} }
func (m *UplinkRXInfo) String() string { return proto.CompactTextString(m) }
func (*UplinkRXInfo) ProtoMessage()    {}
func (*UplinkRXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2cf469f52d369cc, []int{2}
}

func (m *UplinkRXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkRXInfo.Unmarshal(m, b)
}
func (m *UplinkRXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkRXInfo.Marshal(b, m, deterministic)
}
func (m *UplinkRXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkRXInfo.Merge(m, src)
}
func (m *UplinkRXInfo) XXX_Size() int {
	return xxx_messageInfo_UplinkRXInfo.Size(m)
}
func (m *UplinkRXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkRXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkRXInfo proto.InternalMessageInfo

func (m *UplinkRXInfo) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *UplinkRXInfo) GetUplinkId() string {
	if m != nil {
		return m.UplinkId
	}
	return ""
}

func (m *UplinkRXInfo) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *UplinkRXInfo) GetTimeSinceGpsEpoch() *duration.Duration {
	if m != nil {
		return m.TimeSinceGpsEpoch
	}
	return nil
}

func (m *UplinkRXInfo) GetRssi() int32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *UplinkRXInfo) GetLoraSnr() float64 {
	if m != nil {
		return m.LoraSnr
	}
	return 0
}

func (m *UplinkRXInfo) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *UplinkRXInfo) GetRfChain() uint32 {
	if m != nil {
		return m.RfChain
	}
	return 0
}

func (m *UplinkRXInfo) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *UplinkRXInfo) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

func (m *UplinkRXInfo) GetLocation() *common.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *UplinkRXInfo) GetFineTimestampType() gw.FineTimestampType {
	if m != nil {
		return m.FineTimestampType
	}
	return gw.FineTimestampType_NONE
}

type isUplinkRXInfo_FineTimestamp interface {
	isUplinkRXInfo_FineTimestamp()
}

type UplinkRXInfo_EncryptedFineTimestamp struct {
	EncryptedFineTimestamp *EncryptedFineTimestamp `protobuf:"bytes,13,opt,name=encrypted_fine_timestamp,json=encryptedFineTimestamp,proto3,oneof"`
}

type UplinkRXInfo_PlainFineTimestamp struct {
	PlainFineTimestamp *gw.PlainFineTimestamp `protobuf:"bytes,14,opt,name=plain_fine_timestamp,json=plainFineTimestamp,proto3,oneof"`
}

func (*UplinkRXInfo_EncryptedFineTimestamp) isUplinkRXInfo_FineTimestamp() {}

func (*UplinkRXInfo_PlainFineTimestamp) isUplinkRXInfo_FineTimestamp() {}

func (m *UplinkRXInfo) GetFineTimestamp() isUplinkRXInfo_FineTimestamp {
	if m != nil {
		return m.FineTimestamp
	}
	return nil
}

func (m *UplinkRXInfo) GetEncryptedFineTimestamp() *EncryptedFineTimestamp {
	if x, ok := m.GetFineTimestamp().(*UplinkRXInfo_EncryptedFineTimestamp); ok {
		return x.EncryptedFineTimestamp
	}
	return nil
}

func (m *UplinkRXInfo) GetPlainFineTimestamp() *gw.PlainFineTimestamp {
	if x, ok := m.GetFineTimestamp().(*UplinkRXInfo_PlainFineTimestamp); ok {
		return x.PlainFineTimestamp
	}
	return nil
}

func (m *UplinkRXInfo) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UplinkRXInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UplinkRXInfo_EncryptedFineTimestamp)(nil),
		(*UplinkRXInfo_PlainFineTimestamp)(nil),
	}
}

// this s a copy of gw.EncryptedFineTimestamp which the only change that
// the fpga_id is of type string so that it can be returned in HEX format
// instead of base64.
type EncryptedFineTimestamp struct {
	// AES key index used for encrypting the fine timestamp.
	AesKeyIndex uint32 `protobuf:"varint,1,opt,name=aes_key_index,json=aesKeyIndex,proto3" json:"aes_key_index,omitempty"`
	// Encrypted 'main' fine-timestamp (ns precision part of the timestamp).
	EncryptedNs []byte `protobuf:"bytes,2,opt,name=encrypted_ns,json=encryptedNS,proto3" json:"encrypted_ns,omitempty"`
	// FPGA ID.
	FpgaId               string   `protobuf:"bytes,3,opt,name=fpga_id,json=fpgaID,proto3" json:"fpga_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedFineTimestamp) Reset()         { *m = EncryptedFineTimestamp{} }
func (m *EncryptedFineTimestamp) String() string { return proto.CompactTextString(m) }
func (*EncryptedFineTimestamp) ProtoMessage()    {}
func (*EncryptedFineTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2cf469f52d369cc, []int{3}
}

func (m *EncryptedFineTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedFineTimestamp.Unmarshal(m, b)
}
func (m *EncryptedFineTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedFineTimestamp.Marshal(b, m, deterministic)
}
func (m *EncryptedFineTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedFineTimestamp.Merge(m, src)
}
func (m *EncryptedFineTimestamp) XXX_Size() int {
	return xxx_messageInfo_EncryptedFineTimestamp.Size(m)
}
func (m *EncryptedFineTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedFineTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedFineTimestamp proto.InternalMessageInfo

func (m *EncryptedFineTimestamp) GetAesKeyIndex() uint32 {
	if m != nil {
		return m.AesKeyIndex
	}
	return 0
}

func (m *EncryptedFineTimestamp) GetEncryptedNs() []byte {
	if m != nil {
		return m.EncryptedNs
	}
	return nil
}

func (m *EncryptedFineTimestamp) GetFpgaId() string {
	if m != nil {
		return m.FpgaId
	}
	return ""
}

// Same comment as above applies to this message.
type DownlinkTXInfo struct {
	// Gateway ID.
	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayID,proto3" json:"gateway_id,omitempty"`
	// Downlink ID (UUID).
	DownlinkId string `protobuf:"bytes,17,opt,name=downlink_id,json=downlinkID,proto3" json:"downlink_id,omitempty"`
	// TX frequency (in Hz).
	Frequency uint32 `protobuf:"varint,5,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// TX power (in dBm).
	Power int32 `protobuf:"varint,6,opt,name=power,proto3" json:"power,omitempty"`
	// Modulation.
	Modulation common.Modulation `protobuf:"varint,7,opt,name=modulation,proto3,enum=common.Modulation" json:"modulation,omitempty"`
	// Types that are valid to be assigned to ModulationInfo:
	//	*DownlinkTXInfo_LoraModulationInfo
	//	*DownlinkTXInfo_FskModulationInfo
	ModulationInfo isDownlinkTXInfo_ModulationInfo `protobuf_oneof:"modulation_info"`
	// The board identifier for emitting the frame.
	Board uint32 `protobuf:"varint,10,opt,name=board,proto3" json:"board,omitempty"`
	// The antenna identifier for emitting the frame.
	Antenna uint32 `protobuf:"varint,11,opt,name=antenna,proto3" json:"antenna,omitempty"`
	// Timing defines the downlink timing to use.
	Timing gw.DownlinkTiming `protobuf:"varint,12,opt,name=timing,proto3,enum=gw.DownlinkTiming" json:"timing,omitempty"`
	// Types that are valid to be assigned to TimingInfo:
	//	*DownlinkTXInfo_ImmediatelyTimingInfo
	//	*DownlinkTXInfo_DelayTimingInfo
	//	*DownlinkTXInfo_GpsEpochTimingInfo
	TimingInfo isDownlinkTXInfo_TimingInfo `protobuf_oneof:"timing_info"`
	// Gateway specific context.
	// In case of a Class-A downlink, this contains a copy of the uplink context.
	Context              []byte   `protobuf:"bytes,16,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownlinkTXInfo) Reset()         { *m = DownlinkTXInfo{} }
func (m *DownlinkTXInfo) String() string { return proto.CompactTextString(m) }
func (*DownlinkTXInfo) ProtoMessage()    {}
func (*DownlinkTXInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2cf469f52d369cc, []int{4}
}

func (m *DownlinkTXInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkTXInfo.Unmarshal(m, b)
}
func (m *DownlinkTXInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkTXInfo.Marshal(b, m, deterministic)
}
func (m *DownlinkTXInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkTXInfo.Merge(m, src)
}
func (m *DownlinkTXInfo) XXX_Size() int {
	return xxx_messageInfo_DownlinkTXInfo.Size(m)
}
func (m *DownlinkTXInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkTXInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkTXInfo proto.InternalMessageInfo

func (m *DownlinkTXInfo) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *DownlinkTXInfo) GetDownlinkId() string {
	if m != nil {
		return m.DownlinkId
	}
	return ""
}

func (m *DownlinkTXInfo) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DownlinkTXInfo) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *DownlinkTXInfo) GetModulation() common.Modulation {
	if m != nil {
		return m.Modulation
	}
	return common.Modulation_LORA
}

type isDownlinkTXInfo_ModulationInfo interface {
	isDownlinkTXInfo_ModulationInfo()
}

type DownlinkTXInfo_LoraModulationInfo struct {
	LoraModulationInfo *gw.LoRaModulationInfo `protobuf:"bytes,8,opt,name=lora_modulation_info,json=loraModulationInfo,proto3,oneof"`
}

type DownlinkTXInfo_FskModulationInfo struct {
	FskModulationInfo *gw.FSKModulationInfo `protobuf:"bytes,9,opt,name=fsk_modulation_info,json=fskModulationInfo,proto3,oneof"`
}

func (*DownlinkTXInfo_LoraModulationInfo) isDownlinkTXInfo_ModulationInfo() {}

func (*DownlinkTXInfo_FskModulationInfo) isDownlinkTXInfo_ModulationInfo() {}

func (m *DownlinkTXInfo) GetModulationInfo() isDownlinkTXInfo_ModulationInfo {
	if m != nil {
		return m.ModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetLoraModulationInfo() *gw.LoRaModulationInfo {
	if x, ok := m.GetModulationInfo().(*DownlinkTXInfo_LoraModulationInfo); ok {
		return x.LoraModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetFskModulationInfo() *gw.FSKModulationInfo {
	if x, ok := m.GetModulationInfo().(*DownlinkTXInfo_FskModulationInfo); ok {
		return x.FskModulationInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *DownlinkTXInfo) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

func (m *DownlinkTXInfo) GetTiming() gw.DownlinkTiming {
	if m != nil {
		return m.Timing
	}
	return gw.DownlinkTiming_IMMEDIATELY
}

type isDownlinkTXInfo_TimingInfo interface {
	isDownlinkTXInfo_TimingInfo()
}

type DownlinkTXInfo_ImmediatelyTimingInfo struct {
	ImmediatelyTimingInfo *gw.ImmediatelyTimingInfo `protobuf:"bytes,13,opt,name=immediately_timing_info,json=immediatelyTimingInfo,proto3,oneof"`
}

type DownlinkTXInfo_DelayTimingInfo struct {
	DelayTimingInfo *gw.DelayTimingInfo `protobuf:"bytes,14,opt,name=delay_timing_info,json=delayTimingInfo,proto3,oneof"`
}

type DownlinkTXInfo_GpsEpochTimingInfo struct {
	GpsEpochTimingInfo *gw.GPSEpochTimingInfo `protobuf:"bytes,15,opt,name=gps_epoch_timing_info,json=gpsEpochTimingInfo,proto3,oneof"`
}

func (*DownlinkTXInfo_ImmediatelyTimingInfo) isDownlinkTXInfo_TimingInfo() {}

func (*DownlinkTXInfo_DelayTimingInfo) isDownlinkTXInfo_TimingInfo() {}

func (*DownlinkTXInfo_GpsEpochTimingInfo) isDownlinkTXInfo_TimingInfo() {}

func (m *DownlinkTXInfo) GetTimingInfo() isDownlinkTXInfo_TimingInfo {
	if m != nil {
		return m.TimingInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetImmediatelyTimingInfo() *gw.ImmediatelyTimingInfo {
	if x, ok := m.GetTimingInfo().(*DownlinkTXInfo_ImmediatelyTimingInfo); ok {
		return x.ImmediatelyTimingInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetDelayTimingInfo() *gw.DelayTimingInfo {
	if x, ok := m.GetTimingInfo().(*DownlinkTXInfo_DelayTimingInfo); ok {
		return x.DelayTimingInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetGpsEpochTimingInfo() *gw.GPSEpochTimingInfo {
	if x, ok := m.GetTimingInfo().(*DownlinkTXInfo_GpsEpochTimingInfo); ok {
		return x.GpsEpochTimingInfo
	}
	return nil
}

func (m *DownlinkTXInfo) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DownlinkTXInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DownlinkTXInfo_LoraModulationInfo)(nil),
		(*DownlinkTXInfo_FskModulationInfo)(nil),
		(*DownlinkTXInfo_ImmediatelyTimingInfo)(nil),
		(*DownlinkTXInfo_DelayTimingInfo)(nil),
		(*DownlinkTXInfo_GpsEpochTimingInfo)(nil),
	}
}

func init() {
	proto.RegisterEnum("external.RXWindow", RXWindow_name, RXWindow_value)
	proto.RegisterType((*UplinkFrameLog)(nil), "external.UplinkFrameLog")
	proto.RegisterType((*DownlinkFrameLog)(nil), "external.DownlinkFrameLog")
	proto.RegisterType((*UplinkRXInfo)(nil), "external.UplinkRXInfo")
	proto.RegisterType((*EncryptedFineTimestamp)(nil), "external.EncryptedFineTimestamp")
	proto.RegisterType((*DownlinkTXInfo)(nil), "external.DownlinkTXInfo")
}

func init() { proto.RegisterFile("as/external/common.proto", fileDescriptor_b2cf469f52d369cc) }

var fileDescriptor_b2cf469f52d369cc = []byte{
	// 916 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xe1, 0x72, 0xda, 0x46,
	0x10, 0x8e, 0xe2, 0x60, 0x60, 0x65, 0x30, 0x9c, 0x63, 0x47, 0x71, 0xd3, 0x86, 0xf2, 0x8b, 0xa6,
	0xad, 0x34, 0xa1, 0x4f, 0x50, 0x17, 0xc7, 0x25, 0x76, 0x53, 0xcf, 0xe1, 0x4e, 0x3d, 0x9d, 0xce,
	0x68, 0x0e, 0xe9, 0x24, 0xae, 0x88, 0x3b, 0x55, 0x12, 0x03, 0x7a, 0x8b, 0xbe, 0x44, 0x5f, 0xa6,
	0x4f, 0xd5, 0xb9, 0x3b, 0x09, 0x10, 0xa6, 0xd3, 0x5f, 0xd2, 0x7e, 0xfb, 0xed, 0x77, 0x8b, 0xf6,
	0xdb, 0x03, 0x2c, 0x92, 0x3a, 0x74, 0x9d, 0xd1, 0x84, 0x93, 0xc8, 0xf1, 0xc4, 0x62, 0x21, 0xb8,
	0x1d, 0x27, 0x22, 0x13, 0xa8, 0x51, 0xc2, 0x97, 0x6f, 0x43, 0x21, 0xc2, 0x88, 0x3a, 0x0a, 0x9f,
	0x2e, 0x03, 0x27, 0x63, 0x0b, 0x9a, 0x66, 0x64, 0x11, 0x6b, 0xea, 0xe5, 0x17, 0xfb, 0x04, 0x7f,
	0x99, 0x90, 0x8c, 0x95, 0x52, 0x97, 0x67, 0x5a, 0xb8, 0xa2, 0x7f, 0x69, 0x86, 0x2b, 0x27, 0x5c,
	0xe9, 0xa0, 0xff, 0x97, 0x01, 0xed, 0x5f, 0xe2, 0x88, 0xf1, 0xf9, 0x87, 0x84, 0x2c, 0xe8, 0x9d,
	0x08, 0xd1, 0x57, 0x50, 0xcf, 0xd6, 0x2e, 0xe3, 0x81, 0xb0, 0x8c, 0x9e, 0x31, 0x30, 0x87, 0x1d,
	0x3b, 0x5c, 0xd9, 0x9a, 0xf4, 0xf0, 0x38, 0xe6, 0x81, 0xc0, 0xc7, 0xd9, 0x5a, 0x3e, 0x25, 0x35,
	0x29, 0xa8, 0xcf, 0x7b, 0x47, 0x55, 0x2a, 0x2e, 0xa8, 0x89, 0xa6, 0x0e, 0xa0, 0x13, 0xcf, 0x72,
	0x37, 0x26, 0x79, 0x24, 0x88, 0xef, 0xfe, 0x91, 0x0a, 0x6e, 0x1d, 0xf5, 0x8c, 0x41, 0x13, 0xb7,
	0xe3, 0x59, 0x7e, 0xaf, 0xe1, 0x8f, 0x93, 0x9f, 0x3f, 0xf5, 0x19, 0x74, 0x46, 0x62, 0xc5, 0x2b,
	0x3d, 0x7d, 0xbd, 0xdf, 0x13, 0x92, 0x07, 0x95, 0xb4, 0xbd, 0xae, 0x0e, 0x1d, 0xf5, 0xfc, 0xe0,
	0x51, 0x7f, 0xd7, 0xe0, 0x64, 0xb7, 0x5b, 0xf4, 0x39, 0x40, 0x48, 0x32, 0xba, 0x22, 0xb9, 0xcb,
	0x7c, 0x75, 0x54, 0x13, 0x37, 0x0b, 0x64, 0x3c, 0x42, 0x9f, 0x41, 0x73, 0xa9, 0xe8, 0x32, 0xdb,
	0x51, 0xd9, 0x86, 0x06, 0xc6, 0x23, 0x64, 0xc3, 0x0b, 0x39, 0x1f, 0x75, 0x94, 0x39, 0xbc, 0xb4,
	0xf5, 0x6c, 0xec, 0x72, 0x36, 0xf6, 0x43, 0x39, 0x3c, 0xac, 0x78, 0xe8, 0x23, 0xbc, 0x94, 0x4f,
	0x37, 0x65, 0xdc, 0xa3, 0x6e, 0x18, 0xa7, 0x2e, 0x8d, 0x85, 0x37, 0x53, 0x5f, 0xc5, 0x1c, 0xbe,
	0x7e, 0x52, 0x3f, 0x2a, 0x66, 0x8b, 0xbb, 0xb2, 0x6c, 0x22, 0xab, 0x6e, 0xe2, 0xf4, 0x5a, 0xd6,
	0x20, 0x04, 0x2f, 0x92, 0x34, 0x65, 0x56, 0xad, 0x67, 0x0c, 0x6a, 0x58, 0xbd, 0xa3, 0xd7, 0xd0,
	0x88, 0x44, 0x42, 0xdc, 0x94, 0x27, 0xd6, 0x71, 0xcf, 0x18, 0x18, 0xb8, 0x2e, 0xe3, 0x09, 0x4f,
	0x90, 0x05, 0x75, 0x6f, 0x46, 0x38, 0xa7, 0x91, 0x55, 0xef, 0x19, 0x83, 0x16, 0x2e, 0x43, 0x59,
	0x94, 0x04, 0xae, 0x37, 0x23, 0x8c, 0x5b, 0x0d, 0x9d, 0x4a, 0x82, 0x1f, 0x64, 0x88, 0x5e, 0x42,
	0x6d, 0x2a, 0x48, 0xe2, 0x5b, 0x4d, 0x85, 0xeb, 0x40, 0x4a, 0x11, 0x9e, 0x51, 0xce, 0x89, 0x05,
	0x9a, 0x5f, 0x84, 0xe8, 0x1b, 0x79, 0xbe, 0xa7, 0x5a, 0xb6, 0xcc, 0xc2, 0x48, 0x85, 0x11, 0xef,
	0x0a, 0x1c, 0x6f, 0x18, 0xe8, 0x1a, 0xce, 0x02, 0xc6, 0xa9, 0xbb, 0xb1, 0xb8, 0x9b, 0xe5, 0x31,
	0xb5, 0x4e, 0x7a, 0xc6, 0xa0, 0x3d, 0x3c, 0x97, 0xd3, 0xfe, 0xc0, 0x38, 0xdd, 0x7c, 0xc3, 0x87,
	0x3c, 0xa6, 0xb8, 0x1b, 0xec, 0x43, 0xe8, 0x77, 0xb0, 0x28, 0xf7, 0x92, 0x3c, 0xce, 0xa8, 0xef,
	0x56, 0x05, 0xad, 0x96, 0x6a, 0xa2, 0x67, 0x97, 0xfb, 0x65, 0x5f, 0x97, 0xcc, 0x8a, 0xf4, 0x8f,
	0xcf, 0xf0, 0x05, 0x3d, 0x98, 0x91, 0x23, 0x8b, 0x23, 0xc2, 0xf8, 0xbe, 0x72, 0x5b, 0x29, 0x5f,
	0xc8, 0x2e, 0xef, 0x65, 0x7e, 0x5f, 0x0f, 0xc5, 0x4f, 0x50, 0x35, 0x03, 0xc1, 0x33, 0xba, 0xce,
	0xac, 0xd3, 0x9e, 0x31, 0x38, 0xc1, 0x65, 0x78, 0xd5, 0x81, 0x76, 0x55, 0xbf, 0xbf, 0x86, 0x8b,
	0xc3, 0xbd, 0xa2, 0x3e, 0xb4, 0x08, 0x4d, 0xdd, 0x39, 0xcd, 0x5d, 0xc6, 0x7d, 0xba, 0x56, 0x9e,
	0x6d, 0x61, 0x93, 0xd0, 0xf4, 0x96, 0xe6, 0x63, 0x09, 0xa1, 0x2f, 0xe1, 0x64, 0xfb, 0x4d, 0x78,
	0xaa, 0x0c, 0x7a, 0x82, 0xcd, 0x0d, 0xf6, 0x69, 0x82, 0x5e, 0x41, 0x3d, 0x88, 0x43, 0x22, 0x6d,
	0xad, 0x97, 0xf2, 0x58, 0x86, 0xe3, 0x51, 0xff, 0x9f, 0x1a, 0xb4, 0xab, 0x6b, 0xf6, 0x7f, 0x3b,
	0xf2, 0x16, 0x4c, 0xbf, 0x28, 0x90, 0xf9, 0xae, 0xca, 0x43, 0x09, 0x8d, 0x47, 0xe8, 0x0d, 0x34,
	0x83, 0x84, 0xfe, 0xb9, 0xa4, 0xdc, 0xcb, 0x95, 0x61, 0x5b, 0x78, 0x0b, 0x48, 0x97, 0xc5, 0x62,
	0x45, 0xb5, 0x65, 0x6b, 0x58, 0x07, 0x68, 0x08, 0xb0, 0x10, 0xfe, 0x32, 0xd2, 0x6e, 0xaa, 0x2b,
	0x53, 0xa0, 0xd2, 0x4d, 0x3f, 0x6d, 0x32, 0x78, 0x87, 0x25, 0x87, 0xa5, 0xfc, 0xbf, 0x85, 0xf4,
	0x05, 0xd2, 0xd8, 0x0e, 0xeb, 0x4e, 0x60, 0xb2, 0xad, 0x96, 0xbf, 0x4e, 0x0e, 0x4b, 0x56, 0x55,
	0x51, 0x74, 0x03, 0x67, 0x41, 0x3a, 0x7f, 0x22, 0xd5, 0x54, 0x52, 0xda, 0x9d, 0x93, 0xdb, 0x27,
	0x4a, 0xdd, 0x20, 0x9d, 0xef, 0x09, 0x6d, 0x96, 0x08, 0xfe, 0x63, 0x89, 0xcc, 0xea, 0x12, 0xbd,
	0x83, 0xe3, 0x8c, 0x2d, 0x18, 0x0f, 0x8b, 0x4d, 0xa8, 0xde, 0x7b, 0x2a, 0x83, 0x0b, 0x06, 0x9a,
	0xc0, 0x2b, 0xb6, 0x58, 0x50, 0x9f, 0x91, 0x8c, 0x46, 0xb9, 0xab, 0x51, 0xdd, 0x68, 0xab, 0xbc,
	0x53, 0x56, 0xf6, 0x78, 0x4b, 0xd1, 0xf5, 0xaa, 0x59, 0x03, 0x9f, 0xb3, 0x43, 0x09, 0xf4, 0x3d,
	0x74, 0x7d, 0x1a, 0x91, 0xaa, 0x9c, 0xf6, 0xfb, 0x99, 0xea, 0x45, 0x26, 0x2b, 0x42, 0xa7, 0x7e,
	0x15, 0x42, 0xb7, 0x70, 0xbe, 0xb9, 0xdd, 0x2a, 0x32, 0xa7, 0xdb, 0x49, 0xdc, 0xdc, 0x4f, 0xd4,
	0x4d, 0x56, 0x51, 0x42, 0x61, 0x71, 0xbf, 0xed, 0x88, 0xed, 0xac, 0x4d, 0xa7, 0xba, 0x36, 0x5d,
	0x38, 0xdd, 0x9b, 0xcf, 0x55, 0x0b, 0xcc, 0x9d, 0xf3, 0xde, 0xbd, 0x81, 0x06, 0x7e, 0xfc, 0x95,
	0x71, 0x5f, 0xac, 0x50, 0x1d, 0x8e, 0xf0, 0xe3, 0xfb, 0xce, 0x33, 0xfd, 0x32, 0xec, 0x18, 0x57,
	0xef, 0x7f, 0x73, 0x42, 0x96, 0xcd, 0x96, 0x53, 0xe9, 0x2b, 0x67, 0x9a, 0x08, 0x8f, 0x24, 0x8e,
	0x37, 0x63, 0x49, 0x9c, 0x66, 0xc4, 0x9b, 0x7f, 0x4b, 0x62, 0xe6, 0x84, 0xc2, 0xd9, 0xf9, 0xdf,
	0x9e, 0x1e, 0xab, 0xcb, 0xf9, 0xbb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x68, 0x8a, 0xd8,
	0xcd, 0x07, 0x00, 0x00,
}
