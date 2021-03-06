// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type All struct {
	Global               *Global  `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	Proxy                *Proxy   `protobuf:"bytes,2,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Install              *Install `protobuf:"bytes,3,opt,name=install,proto3" json:"install,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{0}
}

func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (m *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(m, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetGlobal() *Global {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *All) GetProxy() *Proxy {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *All) GetInstall() *Install {
	if m != nil {
		return m.Install
	}
	return nil
}

type Global struct {
	LinkerdNamespace string `protobuf:"bytes,1,opt,name=linkerd_namespace,json=linkerdNamespace,proto3" json:"linkerd_namespace,omitempty"`
	CniEnabled       bool   `protobuf:"varint,2,opt,name=cni_enabled,json=cniEnabled,proto3" json:"cni_enabled,omitempty"`
	// Control plane version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// If present, configures identity.
	IdentityContext        *IdentityContext   `protobuf:"bytes,4,opt,name=identity_context,json=identityContext,proto3" json:"identity_context,omitempty"`
	AutoInjectContext      *AutoInjectContext `protobuf:"bytes,6,opt,name=auto_inject_context,json=autoInjectContext,proto3" json:"auto_inject_context,omitempty"` // Deprecated: Do not use.
	OmitWebhookSideEffects bool               `protobuf:"varint,7,opt,name=omitWebhookSideEffects,proto3" json:"omitWebhookSideEffects,omitempty"`
	// Override default `cluster.local`
	ClusterDomain        string   `protobuf:"bytes,8,opt,name=cluster_domain,json=clusterDomain,proto3" json:"cluster_domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{1}
}

func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (m *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(m, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetLinkerdNamespace() string {
	if m != nil {
		return m.LinkerdNamespace
	}
	return ""
}

func (m *Global) GetCniEnabled() bool {
	if m != nil {
		return m.CniEnabled
	}
	return false
}

func (m *Global) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Global) GetIdentityContext() *IdentityContext {
	if m != nil {
		return m.IdentityContext
	}
	return nil
}

// Deprecated: Do not use.
func (m *Global) GetAutoInjectContext() *AutoInjectContext {
	if m != nil {
		return m.AutoInjectContext
	}
	return nil
}

func (m *Global) GetOmitWebhookSideEffects() bool {
	if m != nil {
		return m.OmitWebhookSideEffects
	}
	return false
}

func (m *Global) GetClusterDomain() string {
	if m != nil {
		return m.ClusterDomain
	}
	return ""
}

type Proxy struct {
	ProxyImage              *Image                `protobuf:"bytes,1,opt,name=proxy_image,json=proxyImage,proto3" json:"proxy_image,omitempty"`
	ProxyInitImage          *Image                `protobuf:"bytes,2,opt,name=proxy_init_image,json=proxyInitImage,proto3" json:"proxy_init_image,omitempty"`
	ControlPort             *Port                 `protobuf:"bytes,3,opt,name=control_port,json=controlPort,proto3" json:"control_port,omitempty"`
	IgnoreInboundPorts      []*PortRange          `protobuf:"bytes,4,rep,name=ignore_inbound_ports,json=ignoreInboundPorts,proto3" json:"ignore_inbound_ports,omitempty"`
	IgnoreOutboundPorts     []*PortRange          `protobuf:"bytes,5,rep,name=ignore_outbound_ports,json=ignoreOutboundPorts,proto3" json:"ignore_outbound_ports,omitempty"`
	InboundPort             *Port                 `protobuf:"bytes,6,opt,name=inbound_port,json=inboundPort,proto3" json:"inbound_port,omitempty"`
	AdminPort               *Port                 `protobuf:"bytes,7,opt,name=admin_port,json=adminPort,proto3" json:"admin_port,omitempty"`
	OutboundPort            *Port                 `protobuf:"bytes,8,opt,name=outbound_port,json=outboundPort,proto3" json:"outbound_port,omitempty"`
	Resource                *ResourceRequirements `protobuf:"bytes,9,opt,name=resource,proto3" json:"resource,omitempty"`
	ProxyUid                int64                 `protobuf:"varint,10,opt,name=proxy_uid,json=proxyUid,proto3" json:"proxy_uid,omitempty"`
	LogLevel                *LogLevel             `protobuf:"bytes,11,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	DisableExternalProfiles bool                  `protobuf:"varint,12,opt,name=disable_external_profiles,json=disableExternalProfiles,proto3" json:"disable_external_profiles,omitempty"`
	ProxyVersion            string                `protobuf:"bytes,13,opt,name=proxy_version,json=proxyVersion,proto3" json:"proxy_version,omitempty"`
	ProxyInitImageVersion   string                `protobuf:"bytes,14,opt,name=proxy_init_image_version,json=proxyInitImageVersion,proto3" json:"proxy_init_image_version,omitempty"`
	DebugImage              *Image                `protobuf:"bytes,15,opt,name=debug_image,json=debugImage,proto3" json:"debug_image,omitempty"`
	DebugImageVersion       string                `protobuf:"bytes,16,opt,name=debug_image_version,json=debugImageVersion,proto3" json:"debug_image_version,omitempty"`
	DestinationGetNetworks  string                `protobuf:"bytes,17,opt,name=destination_get_networks,json=destinationGetNetworks,proto3" json:"destination_get_networks,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{2}
}

func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxy.Unmarshal(m, b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
}
func (m *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(m, src)
}
func (m *Proxy) XXX_Size() int {
	return xxx_messageInfo_Proxy.Size(m)
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

func (m *Proxy) GetProxyImage() *Image {
	if m != nil {
		return m.ProxyImage
	}
	return nil
}

func (m *Proxy) GetProxyInitImage() *Image {
	if m != nil {
		return m.ProxyInitImage
	}
	return nil
}

func (m *Proxy) GetControlPort() *Port {
	if m != nil {
		return m.ControlPort
	}
	return nil
}

func (m *Proxy) GetIgnoreInboundPorts() []*PortRange {
	if m != nil {
		return m.IgnoreInboundPorts
	}
	return nil
}

func (m *Proxy) GetIgnoreOutboundPorts() []*PortRange {
	if m != nil {
		return m.IgnoreOutboundPorts
	}
	return nil
}

func (m *Proxy) GetInboundPort() *Port {
	if m != nil {
		return m.InboundPort
	}
	return nil
}

func (m *Proxy) GetAdminPort() *Port {
	if m != nil {
		return m.AdminPort
	}
	return nil
}

func (m *Proxy) GetOutboundPort() *Port {
	if m != nil {
		return m.OutboundPort
	}
	return nil
}

func (m *Proxy) GetResource() *ResourceRequirements {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Proxy) GetProxyUid() int64 {
	if m != nil {
		return m.ProxyUid
	}
	return 0
}

func (m *Proxy) GetLogLevel() *LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func (m *Proxy) GetDisableExternalProfiles() bool {
	if m != nil {
		return m.DisableExternalProfiles
	}
	return false
}

func (m *Proxy) GetProxyVersion() string {
	if m != nil {
		return m.ProxyVersion
	}
	return ""
}

func (m *Proxy) GetProxyInitImageVersion() string {
	if m != nil {
		return m.ProxyInitImageVersion
	}
	return ""
}

func (m *Proxy) GetDebugImage() *Image {
	if m != nil {
		return m.DebugImage
	}
	return nil
}

func (m *Proxy) GetDebugImageVersion() string {
	if m != nil {
		return m.DebugImageVersion
	}
	return ""
}

func (m *Proxy) GetDestinationGetNetworks() string {
	if m != nil {
		return m.DestinationGetNetworks
	}
	return ""
}

type Image struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	PullPolicy           string   `protobuf:"bytes,2,opt,name=pull_policy,json=pullPolicy,proto3" json:"pull_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Image) GetPullPolicy() string {
	if m != nil {
		return m.PullPolicy
	}
	return ""
}

type Port struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{4}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type PortRange struct {
	PortRange            string   `protobuf:"bytes,1,opt,name=port_range,json=portRange,proto3" json:"port_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortRange) Reset()         { *m = PortRange{} }
func (m *PortRange) String() string { return proto.CompactTextString(m) }
func (*PortRange) ProtoMessage()    {}
func (*PortRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{5}
}

func (m *PortRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortRange.Unmarshal(m, b)
}
func (m *PortRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortRange.Marshal(b, m, deterministic)
}
func (m *PortRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortRange.Merge(m, src)
}
func (m *PortRange) XXX_Size() int {
	return xxx_messageInfo_PortRange.Size(m)
}
func (m *PortRange) XXX_DiscardUnknown() {
	xxx_messageInfo_PortRange.DiscardUnknown(m)
}

var xxx_messageInfo_PortRange proto.InternalMessageInfo

func (m *PortRange) GetPortRange() string {
	if m != nil {
		return m.PortRange
	}
	return ""
}

type ResourceRequirements struct {
	RequestCpu           string   `protobuf:"bytes,1,opt,name=request_cpu,json=requestCpu,proto3" json:"request_cpu,omitempty"`
	RequestMemory        string   `protobuf:"bytes,2,opt,name=request_memory,json=requestMemory,proto3" json:"request_memory,omitempty"`
	LimitCpu             string   `protobuf:"bytes,3,opt,name=limit_cpu,json=limitCpu,proto3" json:"limit_cpu,omitempty"`
	LimitMemory          string   `protobuf:"bytes,4,opt,name=limit_memory,json=limitMemory,proto3" json:"limit_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{6}
}

func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (m *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(m, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetRequestCpu() string {
	if m != nil {
		return m.RequestCpu
	}
	return ""
}

func (m *ResourceRequirements) GetRequestMemory() string {
	if m != nil {
		return m.RequestMemory
	}
	return ""
}

func (m *ResourceRequirements) GetLimitCpu() string {
	if m != nil {
		return m.LimitCpu
	}
	return ""
}

func (m *ResourceRequirements) GetLimitMemory() string {
	if m != nil {
		return m.LimitMemory
	}
	return ""
}

// Deprecated: Do not use.
type AutoInjectContext struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoInjectContext) Reset()         { *m = AutoInjectContext{} }
func (m *AutoInjectContext) String() string { return proto.CompactTextString(m) }
func (*AutoInjectContext) ProtoMessage()    {}
func (*AutoInjectContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{7}
}

func (m *AutoInjectContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoInjectContext.Unmarshal(m, b)
}
func (m *AutoInjectContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoInjectContext.Marshal(b, m, deterministic)
}
func (m *AutoInjectContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoInjectContext.Merge(m, src)
}
func (m *AutoInjectContext) XXX_Size() int {
	return xxx_messageInfo_AutoInjectContext.Size(m)
}
func (m *AutoInjectContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoInjectContext.DiscardUnknown(m)
}

var xxx_messageInfo_AutoInjectContext proto.InternalMessageInfo

type IdentityContext struct {
	TrustDomain          string             `protobuf:"bytes,1,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	TrustAnchorsPem      string             `protobuf:"bytes,2,opt,name=trust_anchors_pem,json=trustAnchorsPem,proto3" json:"trust_anchors_pem,omitempty"`
	IssuanceLifetime     *duration.Duration `protobuf:"bytes,3,opt,name=issuance_lifetime,json=issuanceLifetime,proto3" json:"issuance_lifetime,omitempty"`
	ClockSkewAllowance   *duration.Duration `protobuf:"bytes,4,opt,name=clock_skew_allowance,json=clockSkewAllowance,proto3" json:"clock_skew_allowance,omitempty"`
	Scheme               string             `protobuf:"bytes,5,opt,name=scheme,proto3" json:"scheme,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IdentityContext) Reset()         { *m = IdentityContext{} }
func (m *IdentityContext) String() string { return proto.CompactTextString(m) }
func (*IdentityContext) ProtoMessage()    {}
func (*IdentityContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{8}
}

func (m *IdentityContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityContext.Unmarshal(m, b)
}
func (m *IdentityContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityContext.Marshal(b, m, deterministic)
}
func (m *IdentityContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityContext.Merge(m, src)
}
func (m *IdentityContext) XXX_Size() int {
	return xxx_messageInfo_IdentityContext.Size(m)
}
func (m *IdentityContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityContext.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityContext proto.InternalMessageInfo

func (m *IdentityContext) GetTrustDomain() string {
	if m != nil {
		return m.TrustDomain
	}
	return ""
}

func (m *IdentityContext) GetTrustAnchorsPem() string {
	if m != nil {
		return m.TrustAnchorsPem
	}
	return ""
}

func (m *IdentityContext) GetIssuanceLifetime() *duration.Duration {
	if m != nil {
		return m.IssuanceLifetime
	}
	return nil
}

func (m *IdentityContext) GetClockSkewAllowance() *duration.Duration {
	if m != nil {
		return m.ClockSkewAllowance
	}
	return nil
}

func (m *IdentityContext) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

type LogLevel struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{9}
}

func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (m *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(m, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

func (m *LogLevel) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

// Stores information about the last installation/upgrade.
//
// Useful for driving upgrades.
type Install struct {
	// The CLI version that drove the last install or upgrade.
	CliVersion string `protobuf:"bytes,2,opt,name=cli_version,json=cliVersion,proto3" json:"cli_version,omitempty"`
	// The CLI arguments to the install (or upgrade) command, indicating the
	// installer's intent.
	Flags                []*Install_Flag `protobuf:"bytes,3,rep,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{10}
}

func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (m *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(m, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

func (m *Install) GetCliVersion() string {
	if m != nil {
		return m.CliVersion
	}
	return ""
}

func (m *Install) GetFlags() []*Install_Flag {
	if m != nil {
		return m.Flags
	}
	return nil
}

type Install_Flag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install_Flag) Reset()         { *m = Install_Flag{} }
func (m *Install_Flag) String() string { return proto.CompactTextString(m) }
func (*Install_Flag) ProtoMessage()    {}
func (*Install_Flag) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc332a44e926b360, []int{10, 0}
}

func (m *Install_Flag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install_Flag.Unmarshal(m, b)
}
func (m *Install_Flag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install_Flag.Marshal(b, m, deterministic)
}
func (m *Install_Flag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install_Flag.Merge(m, src)
}
func (m *Install_Flag) XXX_Size() int {
	return xxx_messageInfo_Install_Flag.Size(m)
}
func (m *Install_Flag) XXX_DiscardUnknown() {
	xxx_messageInfo_Install_Flag.DiscardUnknown(m)
}

var xxx_messageInfo_Install_Flag proto.InternalMessageInfo

func (m *Install_Flag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Install_Flag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*All)(nil), "linkerd2.config.All")
	proto.RegisterType((*Global)(nil), "linkerd2.config.Global")
	proto.RegisterType((*Proxy)(nil), "linkerd2.config.Proxy")
	proto.RegisterType((*Image)(nil), "linkerd2.config.Image")
	proto.RegisterType((*Port)(nil), "linkerd2.config.Port")
	proto.RegisterType((*PortRange)(nil), "linkerd2.config.PortRange")
	proto.RegisterType((*ResourceRequirements)(nil), "linkerd2.config.ResourceRequirements")
	proto.RegisterType((*AutoInjectContext)(nil), "linkerd2.config.AutoInjectContext")
	proto.RegisterType((*IdentityContext)(nil), "linkerd2.config.IdentityContext")
	proto.RegisterType((*LogLevel)(nil), "linkerd2.config.LogLevel")
	proto.RegisterType((*Install)(nil), "linkerd2.config.Install")
	proto.RegisterType((*Install_Flag)(nil), "linkerd2.config.Install.Flag")
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_cc332a44e926b360) }

var fileDescriptor_cc332a44e926b360 = []byte{
	// 1094 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x86, 0x1d, 0x3b, 0xb6, 0x8f, 0xed, 0xc6, 0x66, 0xd2, 0x54, 0xcd, 0xd0, 0x2d, 0xd3, 0x50,
	0xa0, 0xe8, 0x06, 0x7b, 0x4b, 0x87, 0xb6, 0xe8, 0xd5, 0xd2, 0x5f, 0x64, 0xcd, 0xb2, 0x40, 0xc5,
	0x3a, 0x60, 0x37, 0x82, 0x2c, 0x1d, 0x2b, 0x9c, 0x29, 0xd2, 0x95, 0xa8, 0x24, 0x7d, 0x88, 0x5d,
	0x6f, 0x57, 0x7b, 0xb1, 0xbd, 0xc6, 0x1e, 0x60, 0xe0, 0x21, 0xe5, 0x38, 0x71, 0x93, 0x5d, 0x59,
	0xfc, 0xce, 0xf7, 0x9d, 0x73, 0x24, 0x7e, 0x3c, 0x34, 0x6c, 0xc6, 0x4a, 0x4e, 0x79, 0x3a, 0xb6,
	0x3f, 0xa3, 0x79, 0xae, 0xb4, 0x62, 0x1b, 0x82, 0xcb, 0x19, 0xe6, 0xc9, 0xde, 0xc8, 0xc2, 0x3b,
	0x9f, 0xa7, 0x4a, 0xa5, 0x02, 0xc7, 0x14, 0x9e, 0x94, 0xd3, 0x71, 0x52, 0xe6, 0x91, 0xe6, 0x4a,
	0x5a, 0x81, 0xff, 0x57, 0x0d, 0xd6, 0xf6, 0x85, 0x60, 0x63, 0x58, 0x4f, 0x85, 0x9a, 0x44, 0xc2,
	0xab, 0xed, 0xd6, 0x1e, 0x74, 0xf7, 0xee, 0x8c, 0xae, 0x64, 0x1a, 0xbd, 0xa1, 0x70, 0xe0, 0x68,
	0xec, 0x1b, 0x68, 0xce, 0x73, 0x75, 0xfe, 0xd1, 0xab, 0x13, 0x7f, 0x7b, 0x85, 0x7f, 0x6c, 0xa2,
	0x81, 0x25, 0xb1, 0x3d, 0x68, 0x71, 0x59, 0xe8, 0x48, 0x08, 0x6f, 0x8d, 0xf8, 0xde, 0x0a, 0xff,
	0xc0, 0xc6, 0x83, 0x8a, 0xe8, 0xff, 0x5b, 0x87, 0x75, 0x5b, 0x94, 0x7d, 0x0d, 0x43, 0x47, 0x0f,
	0x65, 0x94, 0x61, 0x31, 0x8f, 0x62, 0xa4, 0x46, 0x3b, 0xc1, 0xc0, 0x05, 0x8e, 0x2a, 0x9c, 0x7d,
	0x01, 0xdd, 0x58, 0xf2, 0x10, 0x65, 0x34, 0x11, 0x98, 0x50, 0x7f, 0xed, 0x00, 0x62, 0xc9, 0x5f,
	0x59, 0x84, 0x79, 0xd0, 0x3a, 0xc5, 0xbc, 0xe0, 0x4a, 0x52, 0x33, 0x9d, 0xa0, 0x5a, 0xb2, 0xb7,
	0x30, 0xe0, 0x09, 0x4a, 0xcd, 0xf5, 0xc7, 0x30, 0x56, 0x52, 0xe3, 0xb9, 0xf6, 0x1a, 0xd4, 0xef,
	0xee, 0x6a, 0xbf, 0x8e, 0xf8, 0xc2, 0xf2, 0x82, 0x0d, 0x7e, 0x19, 0x60, 0xef, 0x61, 0x33, 0x2a,
	0xb5, 0x0a, 0xb9, 0xfc, 0x1d, 0x63, 0xbd, 0xc8, 0xb7, 0x4e, 0xf9, 0xfc, 0x95, 0x7c, 0xfb, 0xa5,
	0x56, 0x07, 0x44, 0x75, 0x09, 0x9e, 0xd7, 0xbd, 0x5a, 0x30, 0x8c, 0xae, 0xc2, 0xec, 0x31, 0x6c,
	0xab, 0x8c, 0xeb, 0x5f, 0x71, 0x72, 0xa2, 0xd4, 0xec, 0x1d, 0x4f, 0xf0, 0xd5, 0x74, 0x8a, 0xb1,
	0x2e, 0xbc, 0x16, 0xbd, 0xea, 0x35, 0x51, 0x76, 0x1f, 0x6e, 0xc5, 0xa2, 0x2c, 0x34, 0xe6, 0x61,
	0xa2, 0xb2, 0x88, 0x4b, 0xaf, 0x4d, 0x6f, 0xdf, 0x77, 0xe8, 0x4b, 0x02, 0xfd, 0x7f, 0x5a, 0xd0,
	0xa4, 0xbd, 0x63, 0x4f, 0xa0, 0x4b, 0xbb, 0x17, 0xf2, 0x2c, 0x4a, 0xd1, 0x19, 0x63, 0x75, 0xa3,
	0x0f, 0x4c, 0x34, 0x00, 0xa2, 0xd2, 0x33, 0xfb, 0x01, 0x06, 0x4e, 0x28, 0xb9, 0x76, 0xea, 0xfa,
	0x8d, 0xea, 0x5b, 0x56, 0x2d, 0xb9, 0xb6, 0x19, 0x9e, 0x42, 0xcf, 0x7c, 0xaf, 0x5c, 0x89, 0x70,
	0xae, 0x72, 0xed, 0x4c, 0x73, 0x7b, 0xd5, 0x64, 0x2a, 0xd7, 0x41, 0xd7, 0x51, 0xcd, 0x82, 0x1d,
	0xc2, 0x16, 0x4f, 0xa5, 0xca, 0x31, 0xe4, 0x72, 0xa2, 0x4a, 0x99, 0x50, 0x82, 0xc2, 0x6b, 0xec,
	0xae, 0x3d, 0xe8, 0xee, 0xed, 0x7c, 0x3a, 0x43, 0x24, 0x53, 0x0c, 0x98, 0xd5, 0x1d, 0x58, 0x99,
	0xc1, 0x0b, 0x76, 0x04, 0xb7, 0x5d, 0x36, 0x55, 0xea, 0xe5, 0x74, 0xcd, 0xff, 0x4d, 0xb7, 0x69,
	0x85, 0x3f, 0x3b, 0x9d, 0xcd, 0xf7, 0x14, 0x7a, 0xcb, 0x6d, 0x39, 0x33, 0x5c, 0xf7, 0x5e, 0xfc,
	0xa2, 0x15, 0xf6, 0x3d, 0x40, 0x94, 0x64, 0x5c, 0x5a, 0x5d, 0xeb, 0x26, 0x5d, 0x87, 0x88, 0xa4,
	0x7a, 0x06, 0xfd, 0x4b, 0x8d, 0xd3, 0x96, 0x5f, 0x2b, 0xec, 0xa9, 0xa5, 0x66, 0xd9, 0x3e, 0xb4,
	0x73, 0x2c, 0x54, 0x99, 0xc7, 0xe8, 0x75, 0x48, 0x76, 0x7f, 0x45, 0x16, 0x38, 0x42, 0x80, 0x1f,
	0x4a, 0x9e, 0x63, 0x86, 0x52, 0x17, 0xc1, 0x42, 0xc6, 0x3e, 0x83, 0x8e, 0x35, 0x42, 0xc9, 0x13,
	0x0f, 0x76, 0x6b, 0x0f, 0xd6, 0x82, 0x36, 0x01, 0xbf, 0xf0, 0x84, 0x3d, 0x86, 0x8e, 0x50, 0x69,
	0x28, 0xf0, 0x14, 0x85, 0xd7, 0xa5, 0x02, 0x77, 0x57, 0x0a, 0x1c, 0xaa, 0xf4, 0xd0, 0x10, 0x82,
	0xb6, 0x70, 0x4f, 0xec, 0x19, 0xdc, 0x4d, 0x78, 0x61, 0x8e, 0x72, 0x88, 0xe7, 0x1a, 0x73, 0x19,
	0x89, 0x70, 0x9e, 0xab, 0x29, 0x17, 0x58, 0x78, 0x3d, 0x3a, 0x02, 0x77, 0x1c, 0xe1, 0x95, 0x8b,
	0x1f, 0xbb, 0x30, 0xfb, 0x0a, 0xfa, 0xb6, 0xa1, 0x6a, 0x00, 0xf4, 0xe9, 0x08, 0xf4, 0x08, 0x7c,
	0xef, 0xa6, 0xc0, 0x13, 0xf0, 0xae, 0xda, 0x77, 0xc1, 0xbf, 0x45, 0xfc, 0xdb, 0x97, 0xed, 0x7a,
	0x21, 0xec, 0x26, 0x38, 0x29, 0x53, 0x67, 0xf9, 0x8d, 0x9b, 0x0f, 0x0c, 0x51, 0xad, 0xdd, 0x47,
	0xb0, 0xb9, 0x24, 0x5c, 0x14, 0x1b, 0x50, 0xb1, 0xe1, 0x05, 0xb1, 0x2a, 0xf4, 0x14, 0xbc, 0x04,
	0x0b, 0xcd, 0x25, 0x8d, 0xf2, 0x30, 0x45, 0x1d, 0x4a, 0xd4, 0x67, 0x2a, 0x9f, 0x15, 0xde, 0x90,
	0x44, 0xdb, 0x4b, 0xf1, 0x37, 0xa8, 0x8f, 0x5c, 0xd4, 0x7f, 0x03, 0x4d, 0x5b, 0xf2, 0x1e, 0x80,
	0x2d, 0x66, 0x06, 0xaa, 0x9b, 0xa5, 0x1d, 0x42, 0xcc, 0x24, 0x35, 0x43, 0x74, 0x5e, 0x0a, 0x73,
	0xfa, 0x04, 0x8f, 0xed, 0x90, 0xef, 0x04, 0x60, 0xa0, 0x63, 0x42, 0xfc, 0x1d, 0x68, 0x90, 0x4b,
	0x18, 0x34, 0xc8, 0x58, 0x26, 0x43, 0x3f, 0xa0, 0x67, 0xff, 0x21, 0x74, 0x16, 0xe7, 0xc0, 0x14,
	0x32, 0x60, 0x98, 0x9b, 0x55, 0x55, 0x68, 0x5e, 0x85, 0xfd, 0xbf, 0x6b, 0xb0, 0xf5, 0x29, 0x17,
	0x99, 0x0e, 0x72, 0xfc, 0x50, 0x62, 0xa1, 0xc3, 0x78, 0x5e, 0x3a, 0x21, 0x38, 0xe8, 0xc5, 0xbc,
	0x34, 0xf3, 0xac, 0x22, 0x64, 0x98, 0xa9, 0xbc, 0xea, 0xb2, 0xef, 0xd0, 0x9f, 0x08, 0x34, 0x1e,
	0x14, 0x3c, 0xe3, 0x36, 0x8b, 0x9d, 0xf7, 0x6d, 0x02, 0x4c, 0x8e, 0x2f, 0xa1, 0x67, 0x83, 0x2e,
	0x43, 0x83, 0xe2, 0x5d, 0xc2, 0xac, 0xde, 0xbf, 0x03, 0xc3, 0x95, 0xd1, 0xfc, 0xac, 0xee, 0xd5,
	0xfc, 0x3f, 0xea, 0xb0, 0x71, 0xe5, 0x12, 0x30, 0xf9, 0x74, 0x5e, 0x16, 0xba, 0x9a, 0xb0, 0xb6,
	0xeb, 0x2e, 0x61, 0x76, 0xbe, 0xb2, 0x87, 0x30, 0xb4, 0x94, 0x48, 0xc6, 0x27, 0x2a, 0x2f, 0xc2,
	0x39, 0x66, 0xae, 0xf3, 0x0d, 0x0a, 0xec, 0x5b, 0xfc, 0x18, 0x33, 0xf6, 0x1a, 0x86, 0xbc, 0x28,
	0xca, 0x48, 0xc6, 0x18, 0x0a, 0x3e, 0x45, 0xcd, 0x33, 0x74, 0xb3, 0xf0, 0xee, 0xc8, 0xde, 0xec,
	0xa3, 0xea, 0x66, 0x1f, 0xbd, 0x74, 0x37, 0x7b, 0x30, 0xa8, 0x34, 0x87, 0x4e, 0xc2, 0xde, 0xc2,
	0x56, 0x2c, 0x54, 0x3c, 0x0b, 0x8b, 0x19, 0x9e, 0x85, 0x91, 0x10, 0xea, 0xcc, 0xc4, 0xdd, 0xdd,
	0x76, 0x43, 0x2a, 0x46, 0xb2, 0x77, 0x33, 0x3c, 0xdb, 0xaf, 0x44, 0x6c, 0x1b, 0xd6, 0x8b, 0xf8,
	0x04, 0x33, 0xf4, 0x9a, 0xd4, 0xb5, 0x5b, 0xf9, 0xbb, 0xd0, 0xae, 0x4e, 0x2b, 0xdb, 0x82, 0xa6,
	0x3d, 0xd7, 0xf6, 0x03, 0xd8, 0x85, 0xff, 0x67, 0x0d, 0x5a, 0xee, 0x9a, 0xa7, 0x5b, 0x5a, 0xf0,
	0x85, 0xd5, 0x9d, 0xc1, 0x62, 0xc1, 0x2b, 0x8f, 0x3f, 0x82, 0xe6, 0x54, 0x44, 0x69, 0xe1, 0xad,
	0xd1, 0xa8, 0xbd, 0x77, 0xdd, 0x1f, 0x86, 0xd1, 0x6b, 0x11, 0xa5, 0x81, 0xe5, 0xee, 0x7c, 0x0b,
	0x0d, 0xb3, 0x34, 0xae, 0x5c, 0xf2, 0x35, 0x3d, 0x9b, 0x9e, 0x4e, 0x23, 0x51, 0xa2, 0xab, 0x65,
	0x17, 0x3f, 0x36, 0xda, 0xb5, 0x41, 0xfd, 0xf9, 0xa3, 0xdf, 0xbe, 0x4b, 0xb9, 0x3e, 0x29, 0x27,
	0xa3, 0x58, 0x65, 0x63, 0x57, 0xa9, 0xfa, 0xdd, 0x1b, 0xbb, 0x0b, 0x46, 0x60, 0x3e, 0x4e, 0x51,
	0xba, 0xbf, 0x5c, 0x93, 0x75, 0xfa, 0x5e, 0x8f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x14, 0x95,
	0x1d, 0x5e, 0x8a, 0x09, 0x00, 0x00,
}
