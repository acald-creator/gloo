// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/static/static.proto

package static

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Static upstreams are used to route request to services listening at fixed IP/Host & Port pairs.
// Static upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
// for additional service-specific configuration.
// Unlike upstreams created by service discovery, Static Upstreams must be created manually by users
type UpstreamSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of addresses and ports
	// at least one must be specified
	Hosts []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Attempt to use outbound TLS
	// If not explicitly set, Gloo will automatically set this to true for port 443
	UseTls *wrappers.BoolValue `protobuf:"bytes,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *options.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// When set, automatically set the sni address to use to the addr field.
	// If both this and host.sni_addr are set, host.sni_addr has priority.
	// defaults to "true".
	AutoSniRewrite *wrappers.BoolValue `protobuf:"bytes,6,opt,name=auto_sni_rewrite,json=autoSniRewrite,proto3" json:"auto_sni_rewrite,omitempty"`
}

func (x *UpstreamSpec) Reset() {
	*x = UpstreamSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec) ProtoMessage() {}

func (x *UpstreamSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSpec.ProtoReflect.Descriptor instead.
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSpec) GetHosts() []*Host {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *UpstreamSpec) GetUseTls() *wrappers.BoolValue {
	if x != nil {
		return x.UseTls
	}
	return nil
}

func (x *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if x != nil {
		return x.ServiceSpec
	}
	return nil
}

func (x *UpstreamSpec) GetAutoSniRewrite() *wrappers.BoolValue {
	if x != nil {
		return x.AutoSniRewrite
	}
	return nil
}

// Represents a single instance of an upstream
type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address (hostname or IP)
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Port the instance is listening on
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Address to use for SNI if using ssl.
	SniAddr string `protobuf:"bytes,4,opt,name=sni_addr,json=sniAddr,proto3" json:"sni_addr,omitempty"`
	// The optional load balancing weight of the upstream host; at least 1.
	// Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	LoadBalancingWeight *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	// (Enterprise Only): Host specific health checking configuration.
	HealthCheckConfig *Host_HealthCheckConfig `protobuf:"bytes,3,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescGZIP(), []int{1}
}

func (x *Host) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Host) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Host) GetSniAddr() string {
	if x != nil {
		return x.SniAddr
	}
	return ""
}

func (x *Host) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if x != nil {
		return x.LoadBalancingWeight
	}
	return nil
}

func (x *Host) GetHealthCheckConfig() *Host_HealthCheckConfig {
	if x != nil {
		return x.HealthCheckConfig
	}
	return nil
}

type Host_HealthCheckConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Enterprise Only): Path to use when health checking this specific host.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// (Enterprise Only): Method to use when health checking this specific host.
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *Host_HealthCheckConfig) Reset() {
	*x = Host_HealthCheckConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host_HealthCheckConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host_HealthCheckConfig) ProtoMessage() {}

func (x *Host_HealthCheckConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host_HealthCheckConfig.ProtoReflect.Descriptor instead.
func (*Host_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Host_HealthCheckConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Host_HealthCheckConfig) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02,
	0x0a, 0x0c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x37,
	0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f, 0x74,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x75, 0x73, 0x65, 0x54, 0x6c, 0x73, 0x12, 0x44, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x44, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x6e, 0x69, 0x5f, 0x72,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x53, 0x6e,
	0x69, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0xca, 0x02, 0x0a, 0x04, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6e, 0x69,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6e, 0x69,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x59, 0x0a, 0x15, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x01, 0x52, 0x13, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x63, 0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x11, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x3f, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x4d, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_goTypes = []interface{}{
	(*UpstreamSpec)(nil),           // 0: static.options.gloo.solo.io.UpstreamSpec
	(*Host)(nil),                   // 1: static.options.gloo.solo.io.Host
	(*Host_HealthCheckConfig)(nil), // 2: static.options.gloo.solo.io.Host.HealthCheckConfig
	(*wrappers.BoolValue)(nil),     // 3: google.protobuf.BoolValue
	(*options.ServiceSpec)(nil),    // 4: options.gloo.solo.io.ServiceSpec
	(*wrappers.UInt32Value)(nil),   // 5: google.protobuf.UInt32Value
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_depIdxs = []int32{
	1, // 0: static.options.gloo.solo.io.UpstreamSpec.hosts:type_name -> static.options.gloo.solo.io.Host
	3, // 1: static.options.gloo.solo.io.UpstreamSpec.use_tls:type_name -> google.protobuf.BoolValue
	4, // 2: static.options.gloo.solo.io.UpstreamSpec.service_spec:type_name -> options.gloo.solo.io.ServiceSpec
	3, // 3: static.options.gloo.solo.io.UpstreamSpec.auto_sni_rewrite:type_name -> google.protobuf.BoolValue
	5, // 4: static.options.gloo.solo.io.Host.load_balancing_weight:type_name -> google.protobuf.UInt32Value
	2, // 5: static.options.gloo.solo.io.Host.health_check_config:type_name -> static.options.gloo.solo.io.Host.HealthCheckConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamSpec); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host_HealthCheckConfig); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_static_static_proto_depIdxs = nil
}
