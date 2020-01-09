// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: upf.api.json

/*
Package upf is a generated VPP binary API for 'upf' module.

It consists of:
	  3 types
	 22 messages
	 11 services
*/
package vpp

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "upf"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x25cc344
)

// UpfFar represents VPP binary API type 'upf_far'.
type UpfFar struct {
	ID              uint32
	ApplyAction     uint8
	DstIntf         uint8
	NetworkInstance string `struc:"[4]byte"`
	OutheadCreate   uint16
	Port            uint16
	Teid            uint32
	IP              uint32
}

func (*UpfFar) GetTypeName() string {
	return "upf_far"
}

// UpfL7Rule represents VPP binary API type 'upf_l7_rule'.
type UpfL7Rule struct {
	ID          uint32
	RegexLength uint32 `struc:"sizeof=Regex"`
	Regex       []byte
}

func (*UpfL7Rule) GetTypeName() string {
	return "upf_l7_rule"
}

// UpfPdr represents VPP binary API type 'upf_pdr'.
type UpfPdr struct {
	ID                 uint16
	Precedence         uint16
	SrcIntf            uint8
	NetworkInstance    string `struc:"[4]byte"`
	TeidFlag           uint8
	Teid               uint32
	TeidIPv4           uint32
	UeFlag             uint8
	UeIPv4             uint32
	OuterHeaderRemoval uint8
	FarID              uint32
}

func (*UpfPdr) GetTypeName() string {
	return "upf_pdr"
}

// UpfAppAddDel represents VPP binary API message 'upf_app_add_del'.
type UpfAppAddDel struct {
	Name  []byte `struc:"[64]byte"`
	Flags uint32
	IsAdd uint8
}

func (*UpfAppAddDel) GetMessageName() string {
	return "upf_app_add_del"
}
func (*UpfAppAddDel) GetCrcString() string {
	return "a4149885"
}
func (*UpfAppAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfAppAddDelReply represents VPP binary API message 'upf_app_add_del_reply'.
type UpfAppAddDelReply struct {
	Retval int32
}

func (*UpfAppAddDelReply) GetMessageName() string {
	return "upf_app_add_del_reply"
}
func (*UpfAppAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfAppAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfAppFlowTimeoutSet represents VPP binary API message 'upf_app_flow_timeout_set'.
type UpfAppFlowTimeoutSet struct {
	Type         uint8
	DefaultValue uint16
}

func (*UpfAppFlowTimeoutSet) GetMessageName() string {
	return "upf_app_flow_timeout_set"
}
func (*UpfAppFlowTimeoutSet) GetCrcString() string {
	return "2dac8f63"
}
func (*UpfAppFlowTimeoutSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfAppFlowTimeoutSetReply represents VPP binary API message 'upf_app_flow_timeout_set_reply'.
type UpfAppFlowTimeoutSetReply struct {
	Retval int32
}

func (*UpfAppFlowTimeoutSetReply) GetMessageName() string {
	return "upf_app_flow_timeout_set_reply"
}
func (*UpfAppFlowTimeoutSetReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfAppFlowTimeoutSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfAppIPRuleAddDel represents VPP binary API message 'upf_app_ip_rule_add_del'.
type UpfAppIPRuleAddDel struct {
	App            []byte `struc:"[64]byte"`
	ID             uint32
	IsIPv6         uint8
	SrcIPAddr      []byte `struc:"[16]byte"`
	SrcIPPrefixLen uint8
	DstIPAddr      []byte `struc:"[16]byte"`
	DstIPPrefixLen uint8
	IsAdd          uint8
}

func (*UpfAppIPRuleAddDel) GetMessageName() string {
	return "upf_app_ip_rule_add_del"
}
func (*UpfAppIPRuleAddDel) GetCrcString() string {
	return "48eea4f1"
}
func (*UpfAppIPRuleAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfAppIPRuleAddDelReply represents VPP binary API message 'upf_app_ip_rule_add_del_reply'.
type UpfAppIPRuleAddDelReply struct {
	Retval int32
}

func (*UpfAppIPRuleAddDelReply) GetMessageName() string {
	return "upf_app_ip_rule_add_del_reply"
}
func (*UpfAppIPRuleAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfAppIPRuleAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfAppL7RuleAddDel represents VPP binary API message 'upf_app_l7_rule_add_del'.
type UpfAppL7RuleAddDel struct {
	App   []byte `struc:"[64]byte"`
	ID    uint32
	Regex []byte `struc:"[1024]byte"`
	IsAdd uint8
}

func (*UpfAppL7RuleAddDel) GetMessageName() string {
	return "upf_app_l7_rule_add_del"
}
func (*UpfAppL7RuleAddDel) GetCrcString() string {
	return "35e936b3"
}
func (*UpfAppL7RuleAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfAppL7RuleAddDelReply represents VPP binary API message 'upf_app_l7_rule_add_del_reply'.
type UpfAppL7RuleAddDelReply struct {
	Retval int32
}

func (*UpfAppL7RuleAddDelReply) GetMessageName() string {
	return "upf_app_l7_rule_add_del_reply"
}
func (*UpfAppL7RuleAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfAppL7RuleAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfApplicationL7RuleDetails represents VPP binary API message 'upf_application_l7_rule_details'.
type UpfApplicationL7RuleDetails struct {
	ID    uint32
	Regex []byte `struc:"[1024]byte"`
}

func (*UpfApplicationL7RuleDetails) GetMessageName() string {
	return "upf_application_l7_rule_details"
}
func (*UpfApplicationL7RuleDetails) GetCrcString() string {
	return "5df24e1f"
}
func (*UpfApplicationL7RuleDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfApplicationL7RuleDump represents VPP binary API message 'upf_application_l7_rule_dump'.
type UpfApplicationL7RuleDump struct {
	App []byte `struc:"[64]byte"`
}

func (*UpfApplicationL7RuleDump) GetMessageName() string {
	return "upf_application_l7_rule_dump"
}
func (*UpfApplicationL7RuleDump) GetCrcString() string {
	return "0b99fe11"
}
func (*UpfApplicationL7RuleDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfApplicationsDetails represents VPP binary API message 'upf_applications_details'.
type UpfApplicationsDetails struct {
	Name  []byte `struc:"[64]byte"`
	Flags uint32
}

func (*UpfApplicationsDetails) GetMessageName() string {
	return "upf_applications_details"
}
func (*UpfApplicationsDetails) GetCrcString() string {
	return "72cd4b5d"
}
func (*UpfApplicationsDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfApplicationsDump represents VPP binary API message 'upf_applications_dump'.
type UpfApplicationsDump struct{}

func (*UpfApplicationsDump) GetMessageName() string {
	return "upf_applications_dump"
}
func (*UpfApplicationsDump) GetCrcString() string {
	return "51077d14"
}
func (*UpfApplicationsDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfEnableDisable represents VPP binary API message 'upf_enable_disable'.
type UpfEnableDisable struct {
	EnableDisable uint8
	SwIfIndex     uint32
}

func (*UpfEnableDisable) GetMessageName() string {
	return "upf_enable_disable"
}
func (*UpfEnableDisable) GetCrcString() string {
	return "57298519"
}
func (*UpfEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfEnableDisableReply represents VPP binary API message 'upf_enable_disable_reply'.
type UpfEnableDisableReply struct {
	Retval int32
}

func (*UpfEnableDisableReply) GetMessageName() string {
	return "upf_enable_disable_reply"
}
func (*UpfEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfPfcpSessionAdd represents VPP binary API message 'upf_pfcp_session_add'.
type UpfPfcpSessionAdd struct {
	SessionID uint64
	Fields    uint32
	Pdrs      []UpfPdr `struc:"[3]UpfPdr"`
	Fars      []UpfFar `struc:"[3]UpfFar"`
}

func (*UpfPfcpSessionAdd) GetMessageName() string {
	return "upf_pfcp_session_add"
}
func (*UpfPfcpSessionAdd) GetCrcString() string {
	return "615aa179"
}
func (*UpfPfcpSessionAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfPfcpSessionAddReply represents VPP binary API message 'upf_pfcp_session_add_reply'.
type UpfPfcpSessionAddReply struct {
	Retval int32
}

func (*UpfPfcpSessionAddReply) GetMessageName() string {
	return "upf_pfcp_session_add_reply"
}
func (*UpfPfcpSessionAddReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfPfcpSessionAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfPfcpSessionDel represents VPP binary API message 'upf_pfcp_session_del'.
type UpfPfcpSessionDel struct {
	SessionID uint64
}

func (*UpfPfcpSessionDel) GetMessageName() string {
	return "upf_pfcp_session_del"
}
func (*UpfPfcpSessionDel) GetCrcString() string {
	return "350f1aaa"
}
func (*UpfPfcpSessionDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfPfcpSessionDelReply represents VPP binary API message 'upf_pfcp_session_del_reply'.
type UpfPfcpSessionDelReply struct {
	Retval int32
}

func (*UpfPfcpSessionDelReply) GetMessageName() string {
	return "upf_pfcp_session_del_reply"
}
func (*UpfPfcpSessionDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfPfcpSessionDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfPfcpSessionUpdate represents VPP binary API message 'upf_pfcp_session_update'.
type UpfPfcpSessionUpdate struct {
	SessionID uint64
	Fars      []UpfFar `struc:"[3]UpfFar"`
}

func (*UpfPfcpSessionUpdate) GetMessageName() string {
	return "upf_pfcp_session_update"
}
func (*UpfPfcpSessionUpdate) GetCrcString() string {
	return "f18229e9"
}
func (*UpfPfcpSessionUpdate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfPfcpSessionUpdateReply represents VPP binary API message 'upf_pfcp_session_update_reply'.
type UpfPfcpSessionUpdateReply struct {
	Retval int32
}

func (*UpfPfcpSessionUpdateReply) GetMessageName() string {
	return "upf_pfcp_session_update_reply"
}
func (*UpfPfcpSessionUpdateReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfPfcpSessionUpdateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// UpfUpdateApp represents VPP binary API message 'upf_update_app'.
type UpfUpdateApp struct {
	App         []byte `struc:"[64]byte"`
	L7RuleCount uint32 `struc:"sizeof=L7Rules"`
	L7Rules     []UpfL7Rule
}

func (*UpfUpdateApp) GetMessageName() string {
	return "upf_update_app"
}
func (*UpfUpdateApp) GetCrcString() string {
	return "50f53737"
}
func (*UpfUpdateApp) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// UpfUpdateAppReply represents VPP binary API message 'upf_update_app_reply'.
type UpfUpdateAppReply struct {
	Retval int32
}

func (*UpfUpdateAppReply) GetMessageName() string {
	return "upf_update_app_reply"
}
func (*UpfUpdateAppReply) GetCrcString() string {
	return "e8d4e804"
}
func (*UpfUpdateAppReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*UpfAppAddDel)(nil), "upf.UpfAppAddDel")
	api.RegisterMessage((*UpfAppAddDelReply)(nil), "upf.UpfAppAddDelReply")
	api.RegisterMessage((*UpfAppFlowTimeoutSet)(nil), "upf.UpfAppFlowTimeoutSet")
	api.RegisterMessage((*UpfAppFlowTimeoutSetReply)(nil), "upf.UpfAppFlowTimeoutSetReply")
	api.RegisterMessage((*UpfAppIPRuleAddDel)(nil), "upf.UpfAppIPRuleAddDel")
	api.RegisterMessage((*UpfAppIPRuleAddDelReply)(nil), "upf.UpfAppIPRuleAddDelReply")
	api.RegisterMessage((*UpfAppL7RuleAddDel)(nil), "upf.UpfAppL7RuleAddDel")
	api.RegisterMessage((*UpfAppL7RuleAddDelReply)(nil), "upf.UpfAppL7RuleAddDelReply")
	api.RegisterMessage((*UpfApplicationL7RuleDetails)(nil), "upf.UpfApplicationL7RuleDetails")
	api.RegisterMessage((*UpfApplicationL7RuleDump)(nil), "upf.UpfApplicationL7RuleDump")
	api.RegisterMessage((*UpfApplicationsDetails)(nil), "upf.UpfApplicationsDetails")
	api.RegisterMessage((*UpfApplicationsDump)(nil), "upf.UpfApplicationsDump")
	api.RegisterMessage((*UpfEnableDisable)(nil), "upf.UpfEnableDisable")
	api.RegisterMessage((*UpfEnableDisableReply)(nil), "upf.UpfEnableDisableReply")
	api.RegisterMessage((*UpfPfcpSessionAdd)(nil), "upf.UpfPfcpSessionAdd")
	api.RegisterMessage((*UpfPfcpSessionAddReply)(nil), "upf.UpfPfcpSessionAddReply")
	api.RegisterMessage((*UpfPfcpSessionDel)(nil), "upf.UpfPfcpSessionDel")
	api.RegisterMessage((*UpfPfcpSessionDelReply)(nil), "upf.UpfPfcpSessionDelReply")
	api.RegisterMessage((*UpfPfcpSessionUpdate)(nil), "upf.UpfPfcpSessionUpdate")
	api.RegisterMessage((*UpfPfcpSessionUpdateReply)(nil), "upf.UpfPfcpSessionUpdateReply")
	api.RegisterMessage((*UpfUpdateApp)(nil), "upf.UpfUpdateApp")
	api.RegisterMessage((*UpfUpdateAppReply)(nil), "upf.UpfUpdateAppReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*UpfAppAddDel)(nil),
		(*UpfAppAddDelReply)(nil),
		(*UpfAppFlowTimeoutSet)(nil),
		(*UpfAppFlowTimeoutSetReply)(nil),
		(*UpfAppIPRuleAddDel)(nil),
		(*UpfAppIPRuleAddDelReply)(nil),
		(*UpfAppL7RuleAddDel)(nil),
		(*UpfAppL7RuleAddDelReply)(nil),
		(*UpfApplicationL7RuleDetails)(nil),
		(*UpfApplicationL7RuleDump)(nil),
		(*UpfApplicationsDetails)(nil),
		(*UpfApplicationsDump)(nil),
		(*UpfEnableDisable)(nil),
		(*UpfEnableDisableReply)(nil),
		(*UpfPfcpSessionAdd)(nil),
		(*UpfPfcpSessionAddReply)(nil),
		(*UpfPfcpSessionDel)(nil),
		(*UpfPfcpSessionDelReply)(nil),
		(*UpfPfcpSessionUpdate)(nil),
		(*UpfPfcpSessionUpdateReply)(nil),
		(*UpfUpdateApp)(nil),
		(*UpfUpdateAppReply)(nil),
	}
}

// RPCService represents RPC service API for upf module.
type RPCService interface {
	DumpUpfApplicationL7Rule(ctx context.Context, in *UpfApplicationL7RuleDump) (RPCService_DumpUpfApplicationL7RuleClient, error)
	DumpUpfApplications(ctx context.Context, in *UpfApplicationsDump) (RPCService_DumpUpfApplicationsClient, error)
	UpfAppAddDel(ctx context.Context, in *UpfAppAddDel) (*UpfAppAddDelReply, error)
	UpfAppFlowTimeoutSet(ctx context.Context, in *UpfAppFlowTimeoutSet) (*UpfAppFlowTimeoutSetReply, error)
	UpfAppIPRuleAddDel(ctx context.Context, in *UpfAppIPRuleAddDel) (*UpfAppIPRuleAddDelReply, error)
	UpfAppL7RuleAddDel(ctx context.Context, in *UpfAppL7RuleAddDel) (*UpfAppL7RuleAddDelReply, error)
	UpfEnableDisable(ctx context.Context, in *UpfEnableDisable) (*UpfEnableDisableReply, error)
	UpfPfcpSessionAdd(ctx context.Context, in *UpfPfcpSessionAdd) (*UpfPfcpSessionAddReply, error)
	UpfPfcpSessionDel(ctx context.Context, in *UpfPfcpSessionDel) (*UpfPfcpSessionDelReply, error)
	UpfPfcpSessionUpdate(ctx context.Context, in *UpfPfcpSessionUpdate) (*UpfPfcpSessionUpdateReply, error)
	UpfUpdateApp(ctx context.Context, in *UpfUpdateApp) (*UpfUpdateAppReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpUpfApplicationL7Rule(ctx context.Context, in *UpfApplicationL7RuleDump) (RPCService_DumpUpfApplicationL7RuleClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpUpfApplicationL7RuleClient{stream}
	return x, nil
}

type RPCService_DumpUpfApplicationL7RuleClient interface {
	Recv() (*UpfApplicationL7RuleDetails, error)
}

type serviceClient_DumpUpfApplicationL7RuleClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpUpfApplicationL7RuleClient) Recv() (*UpfApplicationL7RuleDetails, error) {
	m := new(UpfApplicationL7RuleDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpUpfApplications(ctx context.Context, in *UpfApplicationsDump) (RPCService_DumpUpfApplicationsClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpUpfApplicationsClient{stream}
	return x, nil
}

type RPCService_DumpUpfApplicationsClient interface {
	Recv() (*UpfApplicationsDetails, error)
}

type serviceClient_DumpUpfApplicationsClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpUpfApplicationsClient) Recv() (*UpfApplicationsDetails, error) {
	m := new(UpfApplicationsDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) UpfAppAddDel(ctx context.Context, in *UpfAppAddDel) (*UpfAppAddDelReply, error) {
	out := new(UpfAppAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfAppFlowTimeoutSet(ctx context.Context, in *UpfAppFlowTimeoutSet) (*UpfAppFlowTimeoutSetReply, error) {
	out := new(UpfAppFlowTimeoutSetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfAppIPRuleAddDel(ctx context.Context, in *UpfAppIPRuleAddDel) (*UpfAppIPRuleAddDelReply, error) {
	out := new(UpfAppIPRuleAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfAppL7RuleAddDel(ctx context.Context, in *UpfAppL7RuleAddDel) (*UpfAppL7RuleAddDelReply, error) {
	out := new(UpfAppL7RuleAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfEnableDisable(ctx context.Context, in *UpfEnableDisable) (*UpfEnableDisableReply, error) {
	out := new(UpfEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfPfcpSessionAdd(ctx context.Context, in *UpfPfcpSessionAdd) (*UpfPfcpSessionAddReply, error) {
	out := new(UpfPfcpSessionAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfPfcpSessionDel(ctx context.Context, in *UpfPfcpSessionDel) (*UpfPfcpSessionDelReply, error) {
	out := new(UpfPfcpSessionDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfPfcpSessionUpdate(ctx context.Context, in *UpfPfcpSessionUpdate) (*UpfPfcpSessionUpdateReply, error) {
	out := new(UpfPfcpSessionUpdateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpfUpdateApp(ctx context.Context, in *UpfUpdateApp) (*UpfUpdateAppReply, error) {
	out := new(UpfUpdateAppReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
