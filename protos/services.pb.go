// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	services.proto

It has these top-level messages:
	UserLogin
	UserSignup
	RegisterID
	TeamMember
	UserInfo
	AllUsers
	ProjectOne
	ProjectResponse
	TaskModel
	AllTaskInfo
	TaskResponse
	StatusResponse
	Empty
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserLogin struct {
	Email    string `protobuf:"bytes,1,opt,name=Email,json=email" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,json=password" json:"Password,omitempty"`
}

func (m *UserLogin) Reset()                    { *m = UserLogin{} }
func (m *UserLogin) String() string            { return proto.CompactTextString(m) }
func (*UserLogin) ProtoMessage()               {}
func (*UserLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserLogin) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserSignup struct {
	Username string `protobuf:"bytes,1,opt,name=Username,json=username" json:"Username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,json=email" json:"Email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,json=password" json:"Password,omitempty"`
}

func (m *UserSignup) Reset()                    { *m = UserSignup{} }
func (m *UserSignup) String() string            { return proto.CompactTextString(m) }
func (*UserSignup) ProtoMessage()               {}
func (*UserSignup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserSignup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserSignup) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserSignup) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterID struct {
	ID string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *RegisterID) Reset()                    { *m = RegisterID{} }
func (m *RegisterID) String() string            { return proto.CompactTextString(m) }
func (*RegisterID) ProtoMessage()               {}
func (*RegisterID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type TeamMember struct {
	LeaderID string `protobuf:"bytes,1,opt,name=LeaderID,json=leaderID" json:"LeaderID,omitempty"`
	MemberID string `protobuf:"bytes,2,opt,name=MemberID,json=memberID" json:"MemberID,omitempty"`
}

func (m *TeamMember) Reset()                    { *m = TeamMember{} }
func (m *TeamMember) String() string            { return proto.CompactTextString(m) }
func (*TeamMember) ProtoMessage()               {}
func (*TeamMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TeamMember) GetLeaderID() string {
	if m != nil {
		return m.LeaderID
	}
	return ""
}

func (m *TeamMember) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

type UserInfo struct {
	ID       string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,json=username" json:"Username,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AllUsers struct {
	User *UserInfo `protobuf:"bytes,1,opt,name=User,json=user" json:"User,omitempty"`
}

func (m *AllUsers) Reset()                    { *m = AllUsers{} }
func (m *AllUsers) String() string            { return proto.CompactTextString(m) }
func (*AllUsers) ProtoMessage()               {}
func (*AllUsers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AllUsers) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type ProjectOne struct {
	ID    string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,json=title" json:"Title,omitempty"`
}

func (m *ProjectOne) Reset()                    { *m = ProjectOne{} }
func (m *ProjectOne) String() string            { return proto.CompactTextString(m) }
func (*ProjectOne) ProtoMessage()               {}
func (*ProjectOne) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ProjectOne) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ProjectOne) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type ProjectResponse struct {
	Project *ProjectOne `protobuf:"bytes,1,opt,name=Project,json=project" json:"Project,omitempty"`
}

func (m *ProjectResponse) Reset()                    { *m = ProjectResponse{} }
func (m *ProjectResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectResponse) ProtoMessage()               {}
func (*ProjectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ProjectResponse) GetProject() *ProjectOne {
	if m != nil {
		return m.Project
	}
	return nil
}

type TaskModel struct {
	ID           string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Subject      string `protobuf:"bytes,2,opt,name=Subject,json=subject" json:"Subject,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=Description,json=description" json:"Description,omitempty"`
	Status       string `protobuf:"bytes,4,opt,name=Status,json=status" json:"Status,omitempty"`
	Priority     int32  `protobuf:"varint,5,opt,name=Priority,json=priority" json:"Priority,omitempty"`
	Category     string `protobuf:"bytes,6,opt,name=Category,json=category" json:"Category,omitempty"`
	DateCreated  string `protobuf:"bytes,7,opt,name=DateCreated,json=dateCreated" json:"DateCreated,omitempty"`
	DateModified string `protobuf:"bytes,8,opt,name=DateModified,json=dateModified" json:"DateModified,omitempty"`
	StartDate    string `protobuf:"bytes,9,opt,name=StartDate,json=startDate" json:"StartDate,omitempty"`
	DueDate      string `protobuf:"bytes,10,opt,name=DueDate,json=dueDate" json:"DueDate,omitempty"`
	ProjectID    string `protobuf:"bytes,11,opt,name=ProjectID,json=projectID" json:"ProjectID,omitempty"`
	// Types that are valid to be assigned to Assignee:
	//	*TaskModel_AssigneeObj
	//	*TaskModel_AssigneeID
	Assignee isTaskModel_Assignee `protobuf_oneof:"Assignee"`
}

func (m *TaskModel) Reset()                    { *m = TaskModel{} }
func (m *TaskModel) String() string            { return proto.CompactTextString(m) }
func (*TaskModel) ProtoMessage()               {}
func (*TaskModel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type isTaskModel_Assignee interface{ isTaskModel_Assignee() }

type TaskModel_AssigneeObj struct {
	AssigneeObj *UserInfo `protobuf:"bytes,12,opt,name=AssigneeObj,json=assigneeObj,oneof"`
}
type TaskModel_AssigneeID struct {
	AssigneeID string `protobuf:"bytes,13,opt,name=AssigneeID,json=assigneeID,oneof"`
}

func (*TaskModel_AssigneeObj) isTaskModel_Assignee() {}
func (*TaskModel_AssigneeID) isTaskModel_Assignee()  {}

func (m *TaskModel) GetAssignee() isTaskModel_Assignee {
	if m != nil {
		return m.Assignee
	}
	return nil
}

func (m *TaskModel) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TaskModel) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *TaskModel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskModel) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskModel) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TaskModel) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *TaskModel) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *TaskModel) GetDateModified() string {
	if m != nil {
		return m.DateModified
	}
	return ""
}

func (m *TaskModel) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *TaskModel) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *TaskModel) GetProjectID() string {
	if m != nil {
		return m.ProjectID
	}
	return ""
}

func (m *TaskModel) GetAssigneeObj() *UserInfo {
	if x, ok := m.GetAssignee().(*TaskModel_AssigneeObj); ok {
		return x.AssigneeObj
	}
	return nil
}

func (m *TaskModel) GetAssigneeID() string {
	if x, ok := m.GetAssignee().(*TaskModel_AssigneeID); ok {
		return x.AssigneeID
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskModel) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskModel_OneofMarshaler, _TaskModel_OneofUnmarshaler, _TaskModel_OneofSizer, []interface{}{
		(*TaskModel_AssigneeObj)(nil),
		(*TaskModel_AssigneeID)(nil),
	}
}

func _TaskModel_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskModel)
	// Assignee
	switch x := m.Assignee.(type) {
	case *TaskModel_AssigneeObj:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AssigneeObj); err != nil {
			return err
		}
	case *TaskModel_AssigneeID:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AssigneeID)
	case nil:
	default:
		return fmt.Errorf("TaskModel.Assignee has unexpected type %T", x)
	}
	return nil
}

func _TaskModel_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskModel)
	switch tag {
	case 12: // Assignee.AssigneeObj
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UserInfo)
		err := b.DecodeMessage(msg)
		m.Assignee = &TaskModel_AssigneeObj{msg}
		return true, err
	case 13: // Assignee.AssigneeID
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Assignee = &TaskModel_AssigneeID{x}
		return true, err
	default:
		return false, nil
	}
}

func _TaskModel_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskModel)
	// Assignee
	switch x := m.Assignee.(type) {
	case *TaskModel_AssigneeObj:
		s := proto.Size(x.AssigneeObj)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TaskModel_AssigneeID:
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AssigneeID)))
		n += len(x.AssigneeID)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AllTaskInfo struct {
	Task *TaskModel `protobuf:"bytes,1,opt,name=Task,json=task" json:"Task,omitempty"`
}

func (m *AllTaskInfo) Reset()                    { *m = AllTaskInfo{} }
func (m *AllTaskInfo) String() string            { return proto.CompactTextString(m) }
func (*AllTaskInfo) ProtoMessage()               {}
func (*AllTaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AllTaskInfo) GetTask() *TaskModel {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=Ok,json=ok" json:"Ok,omitempty"`
}

func (m *TaskResponse) Reset()                    { *m = TaskResponse{} }
func (m *TaskResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()               {}
func (*TaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TaskResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type StatusResponse struct {
	Status map[string]int32 `protobuf:"bytes,1,rep,name=Status,json=status" json:"Status,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *StatusResponse) GetStatus() map[string]int32 {
	if m != nil {
		return m.Status
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*UserLogin)(nil), "protos.UserLogin")
	proto.RegisterType((*UserSignup)(nil), "protos.UserSignup")
	proto.RegisterType((*RegisterID)(nil), "protos.RegisterID")
	proto.RegisterType((*TeamMember)(nil), "protos.TeamMember")
	proto.RegisterType((*UserInfo)(nil), "protos.UserInfo")
	proto.RegisterType((*AllUsers)(nil), "protos.AllUsers")
	proto.RegisterType((*ProjectOne)(nil), "protos.ProjectOne")
	proto.RegisterType((*ProjectResponse)(nil), "protos.ProjectResponse")
	proto.RegisterType((*TaskModel)(nil), "protos.TaskModel")
	proto.RegisterType((*AllTaskInfo)(nil), "protos.AllTaskInfo")
	proto.RegisterType((*TaskResponse)(nil), "protos.TaskResponse")
	proto.RegisterType((*StatusResponse)(nil), "protos.StatusResponse")
	proto.RegisterType((*Empty)(nil), "protos.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskerApp service

type TaskerAppClient interface {
	Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*RegisterID, error)
	Signup(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*RegisterID, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (TaskerApp_GetUsersClient, error)
	GetTeamMembers(ctx context.Context, in *RegisterID, opts ...grpc.CallOption) (TaskerApp_GetTeamMembersClient, error)
	AddTask(ctx context.Context, in *TaskModel, opts ...grpc.CallOption) (*TaskResponse, error)
	UpdateTask(ctx context.Context, in *TaskModel, opts ...grpc.CallOption) (*TaskResponse, error)
	GetAllTasks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (TaskerApp_GetAllTasksClient, error)
	GetAllProjects(ctx context.Context, in *Empty, opts ...grpc.CallOption) (TaskerApp_GetAllProjectsClient, error)
	GetStatusCount(ctx context.Context, in *RegisterID, opts ...grpc.CallOption) (TaskerApp_GetStatusCountClient, error)
	UpdateTaskStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type taskerAppClient struct {
	cc *grpc.ClientConn
}

func NewTaskerAppClient(cc *grpc.ClientConn) TaskerAppClient {
	return &taskerAppClient{cc}
}

func (c *taskerAppClient) Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*RegisterID, error) {
	out := new(RegisterID)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) Signup(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*RegisterID, error) {
	out := new(RegisterID)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/Signup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (TaskerApp_GetUsersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskerApp_serviceDesc.Streams[0], c.cc, "/protos.TaskerApp/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskerAppGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskerApp_GetUsersClient interface {
	Recv() (*AllUsers, error)
	grpc.ClientStream
}

type taskerAppGetUsersClient struct {
	grpc.ClientStream
}

func (x *taskerAppGetUsersClient) Recv() (*AllUsers, error) {
	m := new(AllUsers)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskerAppClient) GetTeamMembers(ctx context.Context, in *RegisterID, opts ...grpc.CallOption) (TaskerApp_GetTeamMembersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskerApp_serviceDesc.Streams[1], c.cc, "/protos.TaskerApp/GetTeamMembers", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskerAppGetTeamMembersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskerApp_GetTeamMembersClient interface {
	Recv() (*AllUsers, error)
	grpc.ClientStream
}

type taskerAppGetTeamMembersClient struct {
	grpc.ClientStream
}

func (x *taskerAppGetTeamMembersClient) Recv() (*AllUsers, error) {
	m := new(AllUsers)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskerAppClient) AddTask(ctx context.Context, in *TaskModel, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/AddTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) UpdateTask(ctx context.Context, in *TaskModel, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/UpdateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) GetAllTasks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (TaskerApp_GetAllTasksClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskerApp_serviceDesc.Streams[2], c.cc, "/protos.TaskerApp/GetAllTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskerAppGetAllTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskerApp_GetAllTasksClient interface {
	Recv() (*AllTaskInfo, error)
	grpc.ClientStream
}

type taskerAppGetAllTasksClient struct {
	grpc.ClientStream
}

func (x *taskerAppGetAllTasksClient) Recv() (*AllTaskInfo, error) {
	m := new(AllTaskInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskerAppClient) GetAllProjects(ctx context.Context, in *Empty, opts ...grpc.CallOption) (TaskerApp_GetAllProjectsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskerApp_serviceDesc.Streams[3], c.cc, "/protos.TaskerApp/GetAllProjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskerAppGetAllProjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskerApp_GetAllProjectsClient interface {
	Recv() (*ProjectResponse, error)
	grpc.ClientStream
}

type taskerAppGetAllProjectsClient struct {
	grpc.ClientStream
}

func (x *taskerAppGetAllProjectsClient) Recv() (*ProjectResponse, error) {
	m := new(ProjectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskerAppClient) GetStatusCount(ctx context.Context, in *RegisterID, opts ...grpc.CallOption) (TaskerApp_GetStatusCountClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskerApp_serviceDesc.Streams[4], c.cc, "/protos.TaskerApp/GetStatusCount", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskerAppGetStatusCountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskerApp_GetStatusCountClient interface {
	Recv() (*StatusResponse, error)
	grpc.ClientStream
}

type taskerAppGetStatusCountClient struct {
	grpc.ClientStream
}

func (x *taskerAppGetStatusCountClient) Recv() (*StatusResponse, error) {
	m := new(StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskerAppClient) UpdateTaskStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/UpdateTaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskerApp service

type TaskerAppServer interface {
	Login(context.Context, *UserLogin) (*RegisterID, error)
	Signup(context.Context, *UserSignup) (*RegisterID, error)
	GetUsers(*Empty, TaskerApp_GetUsersServer) error
	GetTeamMembers(*RegisterID, TaskerApp_GetTeamMembersServer) error
	AddTask(context.Context, *TaskModel) (*TaskResponse, error)
	UpdateTask(context.Context, *TaskModel) (*TaskResponse, error)
	GetAllTasks(*Empty, TaskerApp_GetAllTasksServer) error
	GetAllProjects(*Empty, TaskerApp_GetAllProjectsServer) error
	GetStatusCount(*RegisterID, TaskerApp_GetStatusCountServer) error
	UpdateTaskStatus(context.Context, *Empty) (*Empty, error)
}

func RegisterTaskerAppServer(s *grpc.Server, srv TaskerAppServer) {
	s.RegisterService(&_TaskerApp_serviceDesc, srv)
}

func _TaskerApp_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).Login(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).Signup(ctx, req.(*UserSignup))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskerAppServer).GetUsers(m, &taskerAppGetUsersServer{stream})
}

type TaskerApp_GetUsersServer interface {
	Send(*AllUsers) error
	grpc.ServerStream
}

type taskerAppGetUsersServer struct {
	grpc.ServerStream
}

func (x *taskerAppGetUsersServer) Send(m *AllUsers) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskerApp_GetTeamMembers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskerAppServer).GetTeamMembers(m, &taskerAppGetTeamMembersServer{stream})
}

type TaskerApp_GetTeamMembersServer interface {
	Send(*AllUsers) error
	grpc.ServerStream
}

type taskerAppGetTeamMembersServer struct {
	grpc.ServerStream
}

func (x *taskerAppGetTeamMembersServer) Send(m *AllUsers) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskerApp_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).AddTask(ctx, req.(*TaskModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).UpdateTask(ctx, req.(*TaskModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_GetAllTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskerAppServer).GetAllTasks(m, &taskerAppGetAllTasksServer{stream})
}

type TaskerApp_GetAllTasksServer interface {
	Send(*AllTaskInfo) error
	grpc.ServerStream
}

type taskerAppGetAllTasksServer struct {
	grpc.ServerStream
}

func (x *taskerAppGetAllTasksServer) Send(m *AllTaskInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskerApp_GetAllProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskerAppServer).GetAllProjects(m, &taskerAppGetAllProjectsServer{stream})
}

type TaskerApp_GetAllProjectsServer interface {
	Send(*ProjectResponse) error
	grpc.ServerStream
}

type taskerAppGetAllProjectsServer struct {
	grpc.ServerStream
}

func (x *taskerAppGetAllProjectsServer) Send(m *ProjectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskerApp_GetStatusCount_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskerAppServer).GetStatusCount(m, &taskerAppGetStatusCountServer{stream})
}

type TaskerApp_GetStatusCountServer interface {
	Send(*StatusResponse) error
	grpc.ServerStream
}

type taskerAppGetStatusCountServer struct {
	grpc.ServerStream
}

func (x *taskerAppGetStatusCountServer) Send(m *StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskerApp_UpdateTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).UpdateTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/UpdateTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).UpdateTaskStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskerApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.TaskerApp",
	HandlerType: (*TaskerAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _TaskerApp_Login_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _TaskerApp_Signup_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _TaskerApp_AddTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskerApp_UpdateTask_Handler,
		},
		{
			MethodName: "UpdateTaskStatus",
			Handler:    _TaskerApp_UpdateTaskStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _TaskerApp_GetUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTeamMembers",
			Handler:       _TaskerApp_GetTeamMembers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllTasks",
			Handler:       _TaskerApp_GetAllTasks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllProjects",
			Handler:       _TaskerApp_GetAllProjects_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStatusCount",
			Handler:       _TaskerApp_GetStatusCount_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}

func init() { proto.RegisterFile("services.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x6b, 0xe3, 0x56,
	0x10, 0x8d, 0x64, 0xcb, 0x96, 0x47, 0x8e, 0x9b, 0xdc, 0x86, 0x54, 0x98, 0x50, 0x8c, 0x68, 0x21,
	0x0f, 0x6d, 0x30, 0x4e, 0x5b, 0xd2, 0xc0, 0xee, 0xe2, 0x8d, 0x42, 0x62, 0x48, 0x48, 0x90, 0x93,
	0x97, 0x7d, 0x93, 0xad, 0x89, 0x51, 0x2c, 0x4b, 0x42, 0xf7, 0x3a, 0x8b, 0x7f, 0xc1, 0xfe, 0x91,
	0x85, 0xfd, 0x9b, 0xcb, 0xfd, 0x92, 0x3f, 0xf7, 0x61, 0x9f, 0xac, 0x33, 0x33, 0x67, 0x8e, 0xe6,
	0x9e, 0xd1, 0x35, 0xb4, 0x28, 0x16, 0x6f, 0xf1, 0x18, 0xe9, 0x59, 0x5e, 0x64, 0x2c, 0x23, 0x35,
	0xf1, 0x43, 0xbd, 0x77, 0xd0, 0x78, 0xa6, 0x58, 0xdc, 0x65, 0x93, 0x38, 0x25, 0x47, 0x60, 0x5d,
	0xcf, 0xc2, 0x38, 0x71, 0x8d, 0x8e, 0x71, 0xda, 0x08, 0x2c, 0xe4, 0x80, 0xb4, 0xc1, 0x7e, 0x0c,
	0x29, 0xfd, 0x9c, 0x15, 0x91, 0x6b, 0x8a, 0x84, 0x9d, 0x2b, 0xec, 0x7d, 0x02, 0xe0, 0xf4, 0x61,
	0x3c, 0x49, 0xe7, 0x39, 0xaf, 0xe4, 0x28, 0x0d, 0x67, 0xa8, 0x5a, 0xd8, 0x73, 0x85, 0x97, 0xbd,
	0xcd, 0x1f, 0xf5, 0xae, 0x6c, 0xf4, 0x3e, 0x01, 0x08, 0x70, 0x12, 0x53, 0x86, 0xc5, 0xc0, 0x27,
	0x2d, 0x30, 0x07, 0xbe, 0xea, 0x6a, 0xc6, 0xbe, 0xe7, 0x03, 0x3c, 0x61, 0x38, 0xbb, 0xc7, 0xd9,
	0x08, 0x0b, 0xde, 0xe7, 0x0e, 0xc3, 0x88, 0x57, 0x6a, 0xe5, 0x44, 0x61, 0x9e, 0x93, 0x55, 0x03,
	0x5f, 0xbf, 0xff, 0x4c, 0x61, 0xef, 0x3f, 0xf9, 0xc6, 0x83, 0xf4, 0x25, 0xdb, 0x54, 0x58, 0x9b,
	0xc6, 0x5c, 0x9f, 0xc6, 0xeb, 0x82, 0xdd, 0x4f, 0x12, 0x9e, 0xa6, 0xe4, 0x0f, 0xa8, 0xf2, 0x07,
	0xc1, 0x74, 0x7a, 0x07, 0xf2, 0x80, 0xe9, 0x99, 0xee, 0x1b, 0x54, 0x39, 0xcb, 0xeb, 0x01, 0x3c,
	0x16, 0xd9, 0x2b, 0x8e, 0xd9, 0x43, 0x8a, 0x5b, 0x5a, 0x47, 0x60, 0x3d, 0xc5, 0x2c, 0xd1, 0x42,
	0x16, 0xe3, 0xc0, 0xfb, 0x00, 0xbf, 0x28, 0x4e, 0x80, 0x34, 0xcf, 0x52, 0x8a, 0xe4, 0x2f, 0xa8,
	0xab, 0x90, 0xd2, 0x23, 0x5a, 0x6f, 0xd9, 0x3d, 0xa8, 0xe7, 0xf2, 0xd9, 0xfb, 0x56, 0x81, 0xc6,
	0x53, 0x48, 0xa7, 0xf7, 0x59, 0x84, 0xc9, 0x96, 0xa8, 0x0b, 0xf5, 0xe1, 0x7c, 0x24, 0x7a, 0x49,
	0xd9, 0x3a, 0x95, 0x90, 0x74, 0xc0, 0xf1, 0x91, 0x8e, 0x8b, 0x38, 0x67, 0x71, 0x96, 0x2a, 0x67,
	0x9c, 0x68, 0x19, 0x22, 0xc7, 0x50, 0x1b, 0xb2, 0x90, 0xcd, 0xa9, 0x5b, 0x15, 0xc9, 0x1a, 0x15,
	0x48, 0x18, 0x5a, 0xc4, 0x59, 0x11, 0xb3, 0x85, 0x6b, 0x75, 0x8c, 0x53, 0x2b, 0xb0, 0x73, 0x85,
	0x79, 0xee, 0x2a, 0x64, 0x38, 0xc9, 0x8a, 0x85, 0x5b, 0x93, 0x07, 0x3a, 0x56, 0x58, 0x28, 0x86,
	0x0c, 0xaf, 0x0a, 0x0c, 0x19, 0x46, 0x6e, 0x5d, 0x29, 0x2e, 0x43, 0xc4, 0x83, 0x26, 0xaf, 0xb8,
	0xcf, 0xa2, 0xf8, 0x25, 0xc6, 0xc8, 0xb5, 0x45, 0x49, 0x33, 0x5a, 0x89, 0x91, 0x13, 0x68, 0x0c,
	0x59, 0x58, 0x30, 0x5e, 0xe8, 0x36, 0x44, 0x41, 0x83, 0xea, 0x00, 0x9f, 0xd7, 0x9f, 0xa3, 0xc8,
	0x81, 0x9c, 0x37, 0x92, 0x90, 0xf3, 0xd4, 0xf1, 0x0d, 0x7c, 0xd7, 0x91, 0xbc, 0x5c, 0x07, 0xc8,
	0x3f, 0xe0, 0xf4, 0x29, 0x8d, 0x27, 0x29, 0xe2, 0xc3, 0xe8, 0xd5, 0x6d, 0xee, 0xf6, 0xf9, 0x76,
	0x2f, 0x70, 0xc2, 0x65, 0x19, 0xe9, 0x00, 0x68, 0xd6, 0xc0, 0x77, 0xf7, 0x79, 0xd3, 0xdb, 0xbd,
	0x00, 0xc2, 0x32, 0xf6, 0x11, 0xc0, 0xd6, 0x15, 0x1e, 0xd7, 0x48, 0x12, 0xee, 0x95, 0xd8, 0xc5,
	0x3f, 0xa1, 0xca, 0x9f, 0x95, 0xc7, 0x87, 0x5a, 0xab, 0xf4, 0x32, 0xa8, 0xb2, 0x90, 0x4e, 0xbd,
	0xdf, 0xa1, 0xc9, 0x43, 0xe5, 0x76, 0xb4, 0xc0, 0x7c, 0x90, 0x24, 0x3b, 0x30, 0xb3, 0xa9, 0xf7,
	0xc5, 0x80, 0x96, 0xb4, 0xa9, 0x2c, 0xb9, 0x2c, 0x8d, 0x33, 0x3a, 0x95, 0x53, 0xa7, 0xe7, 0xe9,
	0xde, 0xeb, 0x75, 0x0a, 0x5e, 0xa7, 0xac, 0x58, 0x68, 0x73, 0xdb, 0xff, 0x83, 0xb3, 0x12, 0x26,
	0x07, 0x50, 0x99, 0xe2, 0x42, 0x2d, 0x14, 0x7f, 0xe4, 0x6b, 0xfc, 0x16, 0x26, 0x73, 0xb9, 0xc6,
	0x56, 0x20, 0xc1, 0xa5, 0x79, 0x61, 0x78, 0x75, 0xfe, 0xf9, 0xe7, 0x6c, 0xd1, 0xfb, 0x5a, 0x95,
	0x2b, 0x89, 0x45, 0x3f, 0xcf, 0xc9, 0x19, 0x58, 0xf2, 0xea, 0x39, 0x5c, 0x3d, 0x4e, 0x11, 0x6a,
	0x97, 0x9b, 0xbd, 0x72, 0x0b, 0x74, 0xa1, 0xa6, 0xee, 0x1a, 0xb2, 0x4a, 0x90, 0xb1, 0x9d, 0x8c,
	0xbf, 0xc1, 0xbe, 0x41, 0x26, 0xbf, 0xd4, 0x7d, 0x9d, 0x17, 0xaf, 0xd2, 0x2e, 0x2d, 0xd4, 0x9f,
	0x72, 0xd7, 0x20, 0x17, 0xd0, 0xba, 0x41, 0xb6, 0xbc, 0x59, 0x28, 0xd9, 0xd1, 0x74, 0x27, 0xb3,
	0x07, 0xf5, 0x7e, 0x14, 0xf1, 0xd1, 0xc8, 0xb6, 0x5f, 0xed, 0xa3, 0xd5, 0x50, 0x69, 0xc6, 0xbf,
	0x00, 0xcf, 0x39, 0xdf, 0xe0, 0x9f, 0xa3, 0x9d, 0x83, 0x73, 0x83, 0x4c, 0xed, 0xcb, 0xd6, 0x58,
	0xbf, 0xae, 0xbc, 0x9c, 0x5e, 0xa8, 0xae, 0x41, 0x2e, 0xc5, 0x64, 0xfd, 0x24, 0x51, 0x9b, 0xbe,
	0xc5, 0xfb, 0x6d, 0xe3, 0x26, 0xd1, 0x72, 0x5d, 0x83, 0xbc, 0x17, 0x5c, 0xe9, 0xfd, 0x55, 0x36,
	0x4f, 0xd9, 0xce, 0x53, 0x39, 0xde, 0xbd, 0x4a, 0x5d, 0x83, 0x74, 0xe1, 0x60, 0x39, 0xa7, 0xcc,
	0x6e, 0xaa, 0xaf, 0xc3, 0x91, 0xfc, 0x7f, 0x3a, 0xff, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xa5,
	0x59, 0x38, 0xb8, 0x06, 0x00, 0x00,
}
