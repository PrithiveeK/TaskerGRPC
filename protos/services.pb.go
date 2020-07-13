// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	services.proto

It has these top-level messages:
	UserLogin
	UserSignup
	RegisterResponse
	TeamMember
	AllUsers
	ProjectResponse
	TaskRequest
	TaskInfo
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

type RegisterResponse struct {
	ID      string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
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

type AllUsers struct {
	User []*AllUsers_UserInfo `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
}

func (m *AllUsers) Reset()                    { *m = AllUsers{} }
func (m *AllUsers) String() string            { return proto.CompactTextString(m) }
func (*AllUsers) ProtoMessage()               {}
func (*AllUsers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AllUsers) GetUser() []*AllUsers_UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type AllUsers_UserInfo struct {
	ID       string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,json=username" json:"Username,omitempty"`
}

func (m *AllUsers_UserInfo) Reset()                    { *m = AllUsers_UserInfo{} }
func (m *AllUsers_UserInfo) String() string            { return proto.CompactTextString(m) }
func (*AllUsers_UserInfo) ProtoMessage()               {}
func (*AllUsers_UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *AllUsers_UserInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AllUsers_UserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ProjectResponse struct {
	Project []*ProjectResponse_ProjectOne `protobuf:"bytes,1,rep,name=Project,json=project" json:"Project,omitempty"`
}

func (m *ProjectResponse) Reset()                    { *m = ProjectResponse{} }
func (m *ProjectResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectResponse) ProtoMessage()               {}
func (*ProjectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ProjectResponse) GetProject() []*ProjectResponse_ProjectOne {
	if m != nil {
		return m.Project
	}
	return nil
}

type ProjectResponse_ProjectOne struct {
	ID    string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,json=title" json:"Title,omitempty"`
}

func (m *ProjectResponse_ProjectOne) Reset()                    { *m = ProjectResponse_ProjectOne{} }
func (m *ProjectResponse_ProjectOne) String() string            { return proto.CompactTextString(m) }
func (*ProjectResponse_ProjectOne) ProtoMessage()               {}
func (*ProjectResponse_ProjectOne) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ProjectResponse_ProjectOne) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ProjectResponse_ProjectOne) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type TaskRequest struct {
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
	AssigneeID   string `protobuf:"bytes,11,opt,name=AssigneeID,json=assigneeID" json:"AssigneeID,omitempty"`
	ProjectID    string `protobuf:"bytes,12,opt,name=ProjectID,json=projectID" json:"ProjectID,omitempty"`
}

func (m *TaskRequest) Reset()                    { *m = TaskRequest{} }
func (m *TaskRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()               {}
func (*TaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TaskRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TaskRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *TaskRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskRequest) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TaskRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *TaskRequest) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *TaskRequest) GetDateModified() string {
	if m != nil {
		return m.DateModified
	}
	return ""
}

func (m *TaskRequest) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *TaskRequest) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *TaskRequest) GetAssigneeID() string {
	if m != nil {
		return m.AssigneeID
	}
	return ""
}

func (m *TaskRequest) GetProjectID() string {
	if m != nil {
		return m.ProjectID
	}
	return ""
}

type TaskInfo struct {
	ID           string                 `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Subject      string                 `protobuf:"bytes,2,opt,name=Subject,json=subject" json:"Subject,omitempty"`
	Description  string                 `protobuf:"bytes,3,opt,name=Description,json=description" json:"Description,omitempty"`
	Status       string                 `protobuf:"bytes,4,opt,name=Status,json=status" json:"Status,omitempty"`
	Priority     int32                  `protobuf:"varint,5,opt,name=Priority,json=priority" json:"Priority,omitempty"`
	Category     string                 `protobuf:"bytes,6,opt,name=Category,json=category" json:"Category,omitempty"`
	DateCreated  string                 `protobuf:"bytes,7,opt,name=DateCreated,json=dateCreated" json:"DateCreated,omitempty"`
	DateModified string                 `protobuf:"bytes,8,opt,name=DateModified,json=dateModified" json:"DateModified,omitempty"`
	StartDate    string                 `protobuf:"bytes,9,opt,name=StartDate,json=startDate" json:"StartDate,omitempty"`
	DueDate      string                 `protobuf:"bytes,10,opt,name=DueDate,json=dueDate" json:"DueDate,omitempty"`
	ProjectID    string                 `protobuf:"bytes,11,opt,name=ProjectID,json=projectID" json:"ProjectID,omitempty"`
	Assignee     *TaskInfo_AssigneeInfo `protobuf:"bytes,12,opt,name=Assignee,json=assignee" json:"Assignee,omitempty"`
}

func (m *TaskInfo) Reset()                    { *m = TaskInfo{} }
func (m *TaskInfo) String() string            { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()               {}
func (*TaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TaskInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TaskInfo) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *TaskInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskInfo) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TaskInfo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *TaskInfo) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *TaskInfo) GetDateModified() string {
	if m != nil {
		return m.DateModified
	}
	return ""
}

func (m *TaskInfo) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *TaskInfo) GetDueDate() string {
	if m != nil {
		return m.DueDate
	}
	return ""
}

func (m *TaskInfo) GetProjectID() string {
	if m != nil {
		return m.ProjectID
	}
	return ""
}

func (m *TaskInfo) GetAssignee() *TaskInfo_AssigneeInfo {
	if m != nil {
		return m.Assignee
	}
	return nil
}

type TaskInfo_AssigneeInfo struct {
	ID       string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,json=username" json:"Username,omitempty"`
}

func (m *TaskInfo_AssigneeInfo) Reset()                    { *m = TaskInfo_AssigneeInfo{} }
func (m *TaskInfo_AssigneeInfo) String() string            { return proto.CompactTextString(m) }
func (*TaskInfo_AssigneeInfo) ProtoMessage()               {}
func (*TaskInfo_AssigneeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *TaskInfo_AssigneeInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TaskInfo_AssigneeInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AllTaskInfo struct {
	Task []*TaskInfo `protobuf:"bytes,1,rep,name=Task,json=task" json:"Task,omitempty"`
}

func (m *AllTaskInfo) Reset()                    { *m = AllTaskInfo{} }
func (m *AllTaskInfo) String() string            { return proto.CompactTextString(m) }
func (*AllTaskInfo) ProtoMessage()               {}
func (*AllTaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AllTaskInfo) GetTask() []*TaskInfo {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskResponse struct {
	Ok      bool   `protobuf:"varint,1,opt,name=Ok,json=ok" json:"Ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,json=message" json:"Message,omitempty"`
}

func (m *TaskResponse) Reset()                    { *m = TaskResponse{} }
func (m *TaskResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()               {}
func (*TaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TaskResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *TaskResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StatusResponse struct {
	Status []*StatusResponse_StatusInfo `protobuf:"bytes,1,rep,name=Status,json=status" json:"Status,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *StatusResponse) GetStatus() []*StatusResponse_StatusInfo {
	if m != nil {
		return m.Status
	}
	return nil
}

type StatusResponse_StatusInfo struct {
	ID    string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=Count,json=count" json:"Count,omitempty"`
}

func (m *StatusResponse_StatusInfo) Reset()                    { *m = StatusResponse_StatusInfo{} }
func (m *StatusResponse_StatusInfo) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse_StatusInfo) ProtoMessage()               {}
func (*StatusResponse_StatusInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 0} }

func (m *StatusResponse_StatusInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *StatusResponse_StatusInfo) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Empty struct {
	ID string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Empty) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*UserLogin)(nil), "protos.UserLogin")
	proto.RegisterType((*UserSignup)(nil), "protos.UserSignup")
	proto.RegisterType((*RegisterResponse)(nil), "protos.RegisterResponse")
	proto.RegisterType((*TeamMember)(nil), "protos.TeamMember")
	proto.RegisterType((*AllUsers)(nil), "protos.AllUsers")
	proto.RegisterType((*AllUsers_UserInfo)(nil), "protos.AllUsers.UserInfo")
	proto.RegisterType((*ProjectResponse)(nil), "protos.ProjectResponse")
	proto.RegisterType((*ProjectResponse_ProjectOne)(nil), "protos.ProjectResponse.ProjectOne")
	proto.RegisterType((*TaskRequest)(nil), "protos.TaskRequest")
	proto.RegisterType((*TaskInfo)(nil), "protos.TaskInfo")
	proto.RegisterType((*TaskInfo_AssigneeInfo)(nil), "protos.TaskInfo.AssigneeInfo")
	proto.RegisterType((*AllTaskInfo)(nil), "protos.AllTaskInfo")
	proto.RegisterType((*TaskResponse)(nil), "protos.TaskResponse")
	proto.RegisterType((*StatusResponse)(nil), "protos.StatusResponse")
	proto.RegisterType((*StatusResponse_StatusInfo)(nil), "protos.StatusResponse.StatusInfo")
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
	Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*RegisterResponse, error)
	Signup(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllUsers, error)
	GetTeamMembers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllUsers, error)
	AddTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	UpdateTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	GetAllTasks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllTaskInfo, error)
	GetAllProjects(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProjectResponse, error)
	GetStatusCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	UpdateTaskStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type taskerAppClient struct {
	cc *grpc.ClientConn
}

func NewTaskerAppClient(cc *grpc.ClientConn) TaskerAppClient {
	return &taskerAppClient{cc}
}

func (c *taskerAppClient) Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) Signup(ctx context.Context, in *UserSignup, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/Signup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllUsers, error) {
	out := new(AllUsers)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) GetTeamMembers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllUsers, error) {
	out := new(AllUsers)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/GetTeamMembers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) AddTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/AddTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) UpdateTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/UpdateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) GetAllTasks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllTaskInfo, error) {
	out := new(AllTaskInfo)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/GetAllTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) GetAllProjects(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProjectResponse, error) {
	out := new(ProjectResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/GetAllProjects", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskerAppClient) GetStatusCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/protos.TaskerApp/GetStatusCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	Login(context.Context, *UserLogin) (*RegisterResponse, error)
	Signup(context.Context, *UserSignup) (*RegisterResponse, error)
	GetUsers(context.Context, *Empty) (*AllUsers, error)
	GetTeamMembers(context.Context, *Empty) (*AllUsers, error)
	AddTask(context.Context, *TaskRequest) (*TaskResponse, error)
	UpdateTask(context.Context, *TaskRequest) (*TaskResponse, error)
	GetAllTasks(context.Context, *Empty) (*AllTaskInfo, error)
	GetAllProjects(context.Context, *Empty) (*ProjectResponse, error)
	GetStatusCount(context.Context, *Empty) (*StatusResponse, error)
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

func _TaskerApp_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).GetUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_GetTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).GetTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/GetTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).GetTeamMembers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
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
		return srv.(TaskerAppServer).AddTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
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
		return srv.(TaskerAppServer).UpdateTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_GetAllTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).GetAllTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/GetAllTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).GetAllTasks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_GetAllProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).GetAllProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/GetAllProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).GetAllProjects(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskerApp_GetStatusCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskerAppServer).GetStatusCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TaskerApp/GetStatusCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskerAppServer).GetStatusCount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "GetUsers",
			Handler:    _TaskerApp_GetUsers_Handler,
		},
		{
			MethodName: "GetTeamMembers",
			Handler:    _TaskerApp_GetTeamMembers_Handler,
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
			MethodName: "GetAllTasks",
			Handler:    _TaskerApp_GetAllTasks_Handler,
		},
		{
			MethodName: "GetAllProjects",
			Handler:    _TaskerApp_GetAllProjects_Handler,
		},
		{
			MethodName: "GetStatusCount",
			Handler:    _TaskerApp_GetStatusCount_Handler,
		},
		{
			MethodName: "UpdateTaskStatus",
			Handler:    _TaskerApp_UpdateTaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}

func init() { proto.RegisterFile("services.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcd, 0x4e, 0xeb, 0x56,
	0x10, 0x56, 0x7e, 0x9c, 0x38, 0xe3, 0x34, 0x4d, 0x0f, 0x08, 0xdc, 0xa8, 0xad, 0x52, 0xab, 0x0b,
	0xa4, 0xaa, 0x51, 0x09, 0x88, 0x9f, 0x8a, 0x2e, 0x22, 0x5c, 0xa1, 0x48, 0x20, 0x90, 0x03, 0x9b,
	0xee, 0x4c, 0x3c, 0x44, 0x6e, 0x1c, 0xdb, 0xf8, 0x9c, 0xb4, 0x62, 0xd5, 0xc5, 0x7d, 0x91, 0xfb,
	0x1a, 0xf7, 0x81, 0xee, 0x7b, 0x5c, 0x9d, 0x3f, 0x3b, 0x71, 0xc2, 0xd5, 0x65, 0x7f, 0x37, 0xe0,
	0x6f, 0x7e, 0xce, 0xcc, 0x7c, 0x33, 0x67, 0x4e, 0xa0, 0x43, 0x31, 0xfb, 0x37, 0x9c, 0x22, 0x1d,
	0xa4, 0x59, 0xc2, 0x12, 0xd2, 0x10, 0xff, 0xa8, 0xf3, 0x27, 0xb4, 0x1e, 0x28, 0x66, 0xd7, 0xc9,
	0x2c, 0x8c, 0xc9, 0x2e, 0x18, 0x7f, 0x2d, 0xfc, 0x30, 0xb2, 0x2b, 0xfd, 0xca, 0x41, 0xcb, 0x33,
	0x90, 0x03, 0xd2, 0x03, 0xf3, 0xce, 0xa7, 0xf4, 0xbf, 0x24, 0x0b, 0xec, 0xaa, 0x50, 0x98, 0xa9,
	0xc2, 0xce, 0xdf, 0x00, 0xdc, 0x7d, 0x12, 0xce, 0xe2, 0x65, 0xca, 0x2d, 0x39, 0x8a, 0xfd, 0x05,
	0xaa, 0x23, 0xcc, 0xa5, 0xc2, 0xc5, 0xd9, 0xd5, 0xd7, 0xce, 0xae, 0x95, 0xce, 0xbe, 0x80, 0xae,
	0x87, 0xb3, 0x90, 0x32, 0xcc, 0x3c, 0xa4, 0x69, 0x12, 0x53, 0x24, 0x1d, 0xa8, 0x8e, 0x5d, 0x75,
	0x76, 0x35, 0x74, 0x89, 0x0d, 0xcd, 0x05, 0x52, 0xea, 0xcf, 0x50, 0x9d, 0xab, 0xa1, 0xe3, 0x02,
	0xdc, 0xa3, 0xbf, 0xb8, 0xc1, 0xc5, 0x23, 0x66, 0x3c, 0xce, 0x35, 0xfa, 0x01, 0x66, 0xb9, 0xb7,
	0x19, 0x29, 0xcc, 0x75, 0xd2, 0x6a, 0xec, 0xea, 0xfa, 0x16, 0x0a, 0x3b, 0xcf, 0x60, 0x8e, 0xa2,
	0x88, 0x17, 0x45, 0xc9, 0x6f, 0x50, 0xe7, 0xd5, 0xd8, 0x95, 0x7e, 0xed, 0xc0, 0x1a, 0x7e, 0x2f,
	0x89, 0xa4, 0x03, 0xad, 0x1f, 0xf0, 0xbf, 0xe3, 0xf8, 0x29, 0xf1, 0x84, 0x59, 0xef, 0x44, 0x92,
	0xc1, 0x25, 0x1b, 0x69, 0xaf, 0x12, 0x55, 0x5d, 0x27, 0xca, 0x79, 0x57, 0x81, 0x6f, 0xef, 0xb2,
	0xe4, 0x1f, 0x9c, 0xb2, 0xbc, 0xec, 0x0b, 0x68, 0x2a, 0x91, 0x8a, 0xee, 0xe8, 0xe8, 0x25, 0x4b,
	0x8d, 0x6f, 0x63, 0xf4, 0x9a, 0xa9, 0xfc, 0xee, 0x0d, 0x01, 0x0a, 0xf1, 0x46, 0x2e, 0xbb, 0x60,
	0xdc, 0x87, 0x2c, 0xd2, 0x89, 0x18, 0x8c, 0x03, 0xe7, 0x63, 0x15, 0xac, 0x7b, 0x9f, 0xce, 0x3d,
	0x7c, 0x5e, 0x22, 0x65, 0xdb, 0x88, 0x9f, 0x2c, 0x1f, 0x45, 0x46, 0x8a, 0x78, 0x2a, 0x21, 0xe9,
	0x83, 0xe5, 0x22, 0x9d, 0x66, 0x61, 0xca, 0xc2, 0x24, 0x56, 0x5d, 0xb5, 0x82, 0x42, 0x44, 0xf6,
	0xa0, 0x31, 0x61, 0x3e, 0x5b, 0x52, 0xbb, 0x2e, 0x94, 0x0d, 0x2a, 0x90, 0x18, 0x86, 0x2c, 0x4c,
	0xb2, 0x90, 0xbd, 0xd8, 0x46, 0xbf, 0x72, 0x60, 0x78, 0x66, 0xaa, 0x30, 0xd7, 0x5d, 0xfa, 0x0c,
	0x67, 0x49, 0xf6, 0x62, 0x37, 0x24, 0x63, 0x53, 0x85, 0x45, 0x44, 0x9f, 0xe1, 0x65, 0x86, 0x3e,
	0xc3, 0xc0, 0x6e, 0xaa, 0x88, 0x85, 0x88, 0x38, 0xd0, 0xe6, 0x16, 0x37, 0x49, 0x10, 0x3e, 0x85,
	0x18, 0xd8, 0xa6, 0x30, 0x69, 0x07, 0x2b, 0x32, 0xf2, 0x03, 0xb4, 0x26, 0xcc, 0xcf, 0x18, 0x37,
	0xb4, 0x5b, 0xc2, 0xa0, 0x45, 0xb5, 0x80, 0xd7, 0xeb, 0x2e, 0x51, 0xe8, 0x40, 0xd6, 0x1b, 0x48,
	0x48, 0x7e, 0x02, 0x18, 0x51, 0x1a, 0xce, 0x62, 0xc4, 0xb1, 0x6b, 0x5b, 0x42, 0x09, 0x7e, 0x2e,
	0xe1, 0xe7, 0x2a, 0xf6, 0xc7, 0xae, 0xdd, 0x96, 0xe7, 0xa6, 0x5a, 0xe0, 0x7c, 0xa8, 0x81, 0xc9,
	0x79, 0xde, 0x3a, 0x26, 0x5f, 0x49, 0xd6, 0x24, 0xaf, 0x91, 0x68, 0x95, 0x48, 0x24, 0xe7, 0x60,
	0xea, 0x16, 0x08, 0x86, 0xad, 0xe1, 0x8f, 0xfa, 0x7e, 0x68, 0x6e, 0x07, 0x79, 0x8f, 0xf8, 0x0d,
	0x35, 0x75, 0x7f, 0x7a, 0x7f, 0x40, 0x7b, 0x55, 0xf3, 0xa6, 0x9b, 0x7a, 0x04, 0xd6, 0x28, 0x8a,
	0xf2, 0xee, 0xfd, 0x02, 0x75, 0xfe, 0xad, 0x6e, 0x68, 0xb7, 0x9c, 0x81, 0x57, 0x67, 0x3e, 0x9d,
	0x3b, 0x67, 0xd0, 0x96, 0xf7, 0xaa, 0xd8, 0x68, 0xb7, 0x73, 0x11, 0xd0, 0xf4, 0xaa, 0xc9, 0x9c,
	0x73, 0x70, 0xb3, 0x7d, 0xa3, 0xfd, 0x0f, 0x1d, 0xd9, 0xd1, 0xdc, 0xf7, 0x3c, 0xef, 0xb1, 0x8c,
	0xf9, 0xb3, 0x8e, 0xb9, 0x6e, 0xa7, 0xa0, 0x48, 0x42, 0x8d, 0x01, 0xdf, 0x09, 0x85, 0x74, 0xdb,
	0x4e, 0xb8, 0x4c, 0x96, 0xb1, 0x1c, 0xbb, 0x9a, 0x67, 0x4c, 0x39, 0x70, 0xf6, 0xf9, 0x0a, 0x4f,
	0xd9, 0x4b, 0xd9, 0x7c, 0xf8, 0xbe, 0x0e, 0x2d, 0x5e, 0x14, 0x66, 0xa3, 0x34, 0x25, 0xc7, 0x60,
	0xc8, 0xe7, 0xe4, 0x3b, 0x9d, 0x4e, 0xfe, 0xc2, 0xf4, 0x6c, 0x2d, 0xda, 0xd8, 0xec, 0x27, 0xd0,
	0x50, 0xaf, 0x08, 0x59, 0x75, 0x93, 0xb2, 0xcf, 0xf8, 0xfd, 0x0a, 0xe6, 0x15, 0x32, 0xb9, 0xa1,
	0xbf, 0xd1, 0x56, 0x22, 0xcd, 0x5e, 0xb7, 0xbc, 0xa2, 0xc9, 0x21, 0x74, 0xae, 0x90, 0x15, 0xef,
	0xc2, 0x17, 0xb8, 0x1c, 0x43, 0x73, 0x14, 0x04, 0xbc, 0x3a, 0xb2, 0xb3, 0xda, 0x52, 0xb5, 0x18,
	0x7b, 0xbb, 0xeb, 0x42, 0x95, 0xd5, 0x29, 0xc0, 0x43, 0xca, 0x27, 0xff, 0xad, 0x8e, 0x87, 0x60,
	0x5d, 0x21, 0x53, 0x63, 0xb5, 0x91, 0xde, 0xce, 0x4a, 0x7a, 0xf9, 0xdc, 0x9d, 0x89, 0xa2, 0x46,
	0x51, 0xa4, 0x6e, 0xc8, 0x86, 0xd7, 0xfe, 0x2b, 0x8f, 0x05, 0x39, 0x15, 0x9e, 0x72, 0x0e, 0x44,
	0xbf, 0xcb, 0x9e, 0x7b, 0xdb, 0x07, 0x8a, 0xfc, 0x0e, 0xdd, 0xa2, 0x3c, 0xa9, 0x2b, 0xbb, 0xae,
	0xc3, 0x47, 0xf9, 0x7b, 0xe3, 0xe8, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x00, 0xca, 0x48,
	0x88, 0x08, 0x00, 0x00,
}