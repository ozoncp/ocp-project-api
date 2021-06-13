// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: api/ocp-project-api/ocp-project-api.proto

package ocp_project_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListProjectsRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListProjectsRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectsResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId uint64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProjectRequest) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CreateProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *CreateProjectResponse) Reset() {
	*x = CreateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectResponse) ProtoMessage() {}

func (x *CreateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectResponse.ProtoReflect.Descriptor instead.
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProjectResponse) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type MultiCreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*NewProject `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *MultiCreateProjectRequest) Reset() {
	*x = MultiCreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCreateProjectRequest) ProtoMessage() {}

func (x *MultiCreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCreateProjectRequest.ProtoReflect.Descriptor instead.
func (*MultiCreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{4}
}

func (x *MultiCreateProjectRequest) GetProjects() []*NewProject {
	if x != nil {
		return x.Projects
	}
	return nil
}

type MultiCreateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountOfCreated int64 `protobuf:"varint,1,opt,name=count_of_created,json=countOfCreated,proto3" json:"count_of_created,omitempty"`
}

func (x *MultiCreateProjectResponse) Reset() {
	*x = MultiCreateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCreateProjectResponse) ProtoMessage() {}

func (x *MultiCreateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCreateProjectResponse.ProtoReflect.Descriptor instead.
func (*MultiCreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{5}
}

func (x *MultiCreateProjectResponse) GetCountOfCreated() int64 {
	if x != nil {
		return x.CountOfCreated
	}
	return 0
}

type RemoveProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *RemoveProjectRequest) Reset() {
	*x = RemoveProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProjectRequest) ProtoMessage() {}

func (x *RemoveProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProjectRequest.ProtoReflect.Descriptor instead.
func (*RemoveProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveProjectRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type RemoveProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveProjectResponse) Reset() {
	*x = RemoveProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProjectResponse) ProtoMessage() {}

func (x *RemoveProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProjectResponse.ProtoReflect.Descriptor instead.
func (*RemoveProjectResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveProjectResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type DescribeProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DescribeProjectRequest) Reset() {
	*x = DescribeProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProjectRequest) ProtoMessage() {}

func (x *DescribeProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProjectRequest.ProtoReflect.Descriptor instead.
func (*DescribeProjectRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeProjectRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type DescribeProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *DescribeProjectResponse) Reset() {
	*x = DescribeProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeProjectResponse) ProtoMessage() {}

func (x *DescribeProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeProjectResponse.ProtoReflect.Descriptor instead.
func (*DescribeProjectResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{9}
}

func (x *DescribeProjectResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

// Project description
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseId uint64 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{10}
}

func (x *Project) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Project description for multi creation
type NewProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId uint64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewProject) Reset() {
	*x = NewProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProject) ProtoMessage() {}

func (x *NewProject) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_project_api_ocp_project_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProject.ProtoReflect.Descriptor instead.
func (*NewProject) Descriptor() ([]byte, []int) {
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP(), []int{11}
}

func (x *NewProject) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *NewProject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_ocp_project_api_ocp_project_api_proto protoreflect.FileDescriptor

var file_api_ocp_project_api_ocp_project_api_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6f, 0x63, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x4c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x22, 0x50, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x36, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x19, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x22, 0x46, 0x0a, 0x1a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f,
	0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x3e, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x40, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x17, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x4a, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x89, 0x05, 0x0a,
	0x0d, 0x4f, 0x63, 0x70, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x70, 0x69, 0x12, 0x7b,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x24,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x7b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x0f,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x27, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x71, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x80, 0x01, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2a, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x7e, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x2a, 0x16, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63,
	0x70, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x61, 0x70,
	0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_project_api_ocp_project_api_proto_rawDescOnce sync.Once
	file_api_ocp_project_api_ocp_project_api_proto_rawDescData = file_api_ocp_project_api_ocp_project_api_proto_rawDesc
)

func file_api_ocp_project_api_ocp_project_api_proto_rawDescGZIP() []byte {
	file_api_ocp_project_api_ocp_project_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_project_api_ocp_project_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_project_api_ocp_project_api_proto_rawDescData)
	})
	return file_api_ocp_project_api_ocp_project_api_proto_rawDescData
}

var file_api_ocp_project_api_ocp_project_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_ocp_project_api_ocp_project_api_proto_goTypes = []interface{}{
	(*ListProjectsRequest)(nil),        // 0: ocp.project.api.ListProjectsRequest
	(*ListProjectsResponse)(nil),       // 1: ocp.project.api.ListProjectsResponse
	(*CreateProjectRequest)(nil),       // 2: ocp.project.api.CreateProjectRequest
	(*CreateProjectResponse)(nil),      // 3: ocp.project.api.CreateProjectResponse
	(*MultiCreateProjectRequest)(nil),  // 4: ocp.project.api.MultiCreateProjectRequest
	(*MultiCreateProjectResponse)(nil), // 5: ocp.project.api.MultiCreateProjectResponse
	(*RemoveProjectRequest)(nil),       // 6: ocp.project.api.RemoveProjectRequest
	(*RemoveProjectResponse)(nil),      // 7: ocp.project.api.RemoveProjectResponse
	(*DescribeProjectRequest)(nil),     // 8: ocp.project.api.DescribeProjectRequest
	(*DescribeProjectResponse)(nil),    // 9: ocp.project.api.DescribeProjectResponse
	(*Project)(nil),                    // 10: ocp.project.api.Project
	(*NewProject)(nil),                 // 11: ocp.project.api.NewProject
}
var file_api_ocp_project_api_ocp_project_api_proto_depIdxs = []int32{
	10, // 0: ocp.project.api.ListProjectsResponse.projects:type_name -> ocp.project.api.Project
	11, // 1: ocp.project.api.MultiCreateProjectRequest.projects:type_name -> ocp.project.api.NewProject
	10, // 2: ocp.project.api.DescribeProjectResponse.project:type_name -> ocp.project.api.Project
	0,  // 3: ocp.project.api.OcpProjectApi.ListProjects:input_type -> ocp.project.api.ListProjectsRequest
	8,  // 4: ocp.project.api.OcpProjectApi.DescribeProject:input_type -> ocp.project.api.DescribeProjectRequest
	2,  // 5: ocp.project.api.OcpProjectApi.CreateProject:input_type -> ocp.project.api.CreateProjectRequest
	4,  // 6: ocp.project.api.OcpProjectApi.MultiCreateProject:input_type -> ocp.project.api.MultiCreateProjectRequest
	6,  // 7: ocp.project.api.OcpProjectApi.RemoveProject:input_type -> ocp.project.api.RemoveProjectRequest
	1,  // 8: ocp.project.api.OcpProjectApi.ListProjects:output_type -> ocp.project.api.ListProjectsResponse
	9,  // 9: ocp.project.api.OcpProjectApi.DescribeProject:output_type -> ocp.project.api.DescribeProjectResponse
	3,  // 10: ocp.project.api.OcpProjectApi.CreateProject:output_type -> ocp.project.api.CreateProjectResponse
	5,  // 11: ocp.project.api.OcpProjectApi.MultiCreateProject:output_type -> ocp.project.api.MultiCreateProjectResponse
	7,  // 12: ocp.project.api.OcpProjectApi.RemoveProject:output_type -> ocp.project.api.RemoveProjectResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_ocp_project_api_ocp_project_api_proto_init() }
func file_api_ocp_project_api_ocp_project_api_proto_init() {
	if File_api_ocp_project_api_ocp_project_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectResponse); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCreateProjectRequest); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCreateProjectResponse); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProjectRequest); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProjectResponse); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProjectRequest); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeProjectResponse); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_api_ocp_project_api_ocp_project_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProject); i {
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
			RawDescriptor: file_api_ocp_project_api_ocp_project_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_project_api_ocp_project_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_project_api_ocp_project_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_project_api_ocp_project_api_proto_msgTypes,
	}.Build()
	File_api_ocp_project_api_ocp_project_api_proto = out.File
	file_api_ocp_project_api_ocp_project_api_proto_rawDesc = nil
	file_api_ocp_project_api_ocp_project_api_proto_goTypes = nil
	file_api_ocp_project_api_ocp_project_api_proto_depIdxs = nil
}
