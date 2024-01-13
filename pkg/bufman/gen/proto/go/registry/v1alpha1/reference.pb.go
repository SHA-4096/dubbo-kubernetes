// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: registry/v1alpha1/reference.proto

package registryv1alpha1

import (
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

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reference:
	//
	//	*Reference_Tag
	//	*Reference_Commit
	//	*Reference_Main
	//	*Reference_Draft
	Reference isReference_Reference `protobuf_oneof:"reference"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{0}
}

func (m *Reference) GetReference() isReference_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (x *Reference) GetTag() *RepositoryTag {
	if x, ok := x.GetReference().(*Reference_Tag); ok {
		return x.Tag
	}
	return nil
}

func (x *Reference) GetCommit() *RepositoryCommit {
	if x, ok := x.GetReference().(*Reference_Commit); ok {
		return x.Commit
	}
	return nil
}

func (x *Reference) GetMain() *RepositoryMainReference {
	if x, ok := x.GetReference().(*Reference_Main); ok {
		return x.Main
	}
	return nil
}

func (x *Reference) GetDraft() *RepositoryDraft {
	if x, ok := x.GetReference().(*Reference_Draft); ok {
		return x.Draft
	}
	return nil
}

type isReference_Reference interface {
	isReference_Reference()
}

type Reference_Tag struct {
	// The requested reference is a tag.
	Tag *RepositoryTag `protobuf:"bytes,2,opt,name=tag,proto3,oneof"`
}

type Reference_Commit struct {
	// The requested reference is a commit.
	Commit *RepositoryCommit `protobuf:"bytes,3,opt,name=commit,proto3,oneof"`
}

type Reference_Main struct {
	// The requested reference is the default reference.
	Main *RepositoryMainReference `protobuf:"bytes,5,opt,name=main,proto3,oneof"`
}

type Reference_Draft struct {
	// The requested reference is a draft commit.
	Draft *RepositoryDraft `protobuf:"bytes,6,opt,name=draft,proto3,oneof"`
}

func (*Reference_Tag) isReference_Reference() {}

func (*Reference_Commit) isReference_Reference() {}

func (*Reference_Main) isReference_Reference() {}

func (*Reference_Draft) isReference_Reference() {}

type RepositoryMainReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is always 'main'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The latest commit in this repository. If the repository has no commits,
	// this will be empty.
	Commit *RepositoryCommit `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *RepositoryMainReference) Reset() {
	*x = RepositoryMainReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_reference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryMainReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryMainReference) ProtoMessage() {}

func (x *RepositoryMainReference) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_reference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryMainReference.ProtoReflect.Descriptor instead.
func (*RepositoryMainReference) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{1}
}

func (x *RepositoryMainReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoryMainReference) GetCommit() *RepositoryCommit {
	if x != nil {
		return x.Commit
	}
	return nil
}

type RepositoryDraft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the draft
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The commit this draft points to.
	Commit *RepositoryCommit `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *RepositoryDraft) Reset() {
	*x = RepositoryDraft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_reference_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryDraft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryDraft) ProtoMessage() {}

func (x *RepositoryDraft) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_reference_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryDraft.ProtoReflect.Descriptor instead.
func (*RepositoryDraft) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{2}
}

func (x *RepositoryDraft) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepositoryDraft) GetCommit() *RepositoryCommit {
	if x != nil {
		return x.Commit
	}
	return nil
}

type GetReferenceByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the requested reference.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Owner of the repository the reference belongs to.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Name of the repository the reference belongs to.
	RepositoryName string `protobuf:"bytes,3,opt,name=repository_name,json=repositoryName,proto3" json:"repository_name,omitempty"`
}

func (x *GetReferenceByNameRequest) Reset() {
	*x = GetReferenceByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_reference_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReferenceByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReferenceByNameRequest) ProtoMessage() {}

func (x *GetReferenceByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_reference_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReferenceByNameRequest.ProtoReflect.Descriptor instead.
func (*GetReferenceByNameRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{3}
}

func (x *GetReferenceByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetReferenceByNameRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetReferenceByNameRequest) GetRepositoryName() string {
	if x != nil {
		return x.RepositoryName
	}
	return ""
}

type GetReferenceByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference *Reference `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (x *GetReferenceByNameResponse) Reset() {
	*x = GetReferenceByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_reference_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReferenceByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReferenceByNameResponse) ProtoMessage() {}

func (x *GetReferenceByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_reference_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReferenceByNameResponse.ProtoReflect.Descriptor instead.
func (*GetReferenceByNameResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{4}
}

func (x *GetReferenceByNameResponse) GetReference() *Reference {
	if x != nil {
		return x.Reference
	}
	return nil
}

type ListGitCommitMetadataForReferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String that represents the name of the reference.
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Owner of the repository the reference belongs to.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// Name of the repository the reference belongs to.
	RepositoryName string `protobuf:"bytes,3,opt,name=repository_name,json=repositoryName,proto3" json:"repository_name,omitempty"`
}

func (x *ListGitCommitMetadataForReferenceRequest) Reset() {
	*x = ListGitCommitMetadataForReferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_reference_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGitCommitMetadataForReferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGitCommitMetadataForReferenceRequest) ProtoMessage() {}

func (x *ListGitCommitMetadataForReferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_reference_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGitCommitMetadataForReferenceRequest.ProtoReflect.Descriptor instead.
func (*ListGitCommitMetadataForReferenceRequest) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{5}
}

func (x *ListGitCommitMetadataForReferenceRequest) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *ListGitCommitMetadataForReferenceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ListGitCommitMetadataForReferenceRequest) GetRepositoryName() string {
	if x != nil {
		return x.RepositoryName
	}
	return ""
}

type ListGitCommitMetadataForReferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the BSR commit the reference resolved to.
	CommitId string `protobuf:"bytes,1,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	// List of git commits and metadata associated with the resolved reference.
	GitCommitMetadatas []*GitCommitMetadata `protobuf:"bytes,2,rep,name=git_commit_metadatas,json=gitCommitMetadatas,proto3" json:"git_commit_metadatas,omitempty"`
}

func (x *ListGitCommitMetadataForReferenceResponse) Reset() {
	*x = ListGitCommitMetadataForReferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1alpha1_reference_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGitCommitMetadataForReferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGitCommitMetadataForReferenceResponse) ProtoMessage() {}

func (x *ListGitCommitMetadataForReferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1alpha1_reference_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGitCommitMetadataForReferenceResponse.ProtoReflect.Descriptor instead.
func (*ListGitCommitMetadataForReferenceResponse) Descriptor() ([]byte, []int) {
	return file_registry_v1alpha1_reference_proto_rawDescGZIP(), []int{6}
}

func (x *ListGitCommitMetadataForReferenceResponse) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *ListGitCommitMetadataForReferenceResponse) GetGitCommitMetadatas() []*GitCommitMetadata {
	if x != nil {
		return x.GitCommitMetadatas
	}
	return nil
}

var File_registry_v1alpha1_reference_proto protoreflect.FileDescriptor

var file_registry_v1alpha1_reference_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x29, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x67, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x61,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x61, 0x67, 0x48, 0x00, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x55, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62,
	0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x58, 0x0a, 0x04, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61,
	0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4d,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x52, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62,
	0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x72, 0x61, 0x66, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x64, 0x72, 0x61, 0x66, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x04, 0x10,
	0x05, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x22, 0x82, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4d,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x53, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x7a, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x62,
	0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x22, 0x6e, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x70, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x28, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb8, 0x01,
	0x0a, 0x29, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x6e, 0x0a, 0x14, 0x67, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e,
	0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x12, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x73, 0x32, 0x91, 0x03, 0x0a, 0x10, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa6, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x62, 0x75, 0x66,
	0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0xd3, 0x01, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x46, 0x6f, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x53, 0x2e, 0x62,
	0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x69, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f,
	0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x54, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0xe9, 0x02, 0x0a,
	0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2e, 0x64, 0x75, 0x62, 0x62,
	0x6f, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2f, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2,
	0x02, 0x05, 0x42, 0x44, 0x41, 0x4f, 0x52, 0xaa, 0x02, 0x29, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e,
	0x2e, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x2e, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x4f, 0x72,
	0x67, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x29, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x5c, 0x44, 0x75, 0x62,
	0x62, 0x6f, 0x5c, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x35, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e, 0x5c, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x5c, 0x41,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x5c, 0x4f, 0x72, 0x67, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x42, 0x75, 0x66, 0x6d, 0x61, 0x6e,
	0x3a, 0x3a, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x3a,
	0x3a, 0x4f, 0x72, 0x67, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registry_v1alpha1_reference_proto_rawDescOnce sync.Once
	file_registry_v1alpha1_reference_proto_rawDescData = file_registry_v1alpha1_reference_proto_rawDesc
)

func file_registry_v1alpha1_reference_proto_rawDescGZIP() []byte {
	file_registry_v1alpha1_reference_proto_rawDescOnce.Do(func() {
		file_registry_v1alpha1_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_v1alpha1_reference_proto_rawDescData)
	})
	return file_registry_v1alpha1_reference_proto_rawDescData
}

var file_registry_v1alpha1_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_registry_v1alpha1_reference_proto_goTypes = []interface{}{
	(*Reference)(nil),                                 // 0: bufman.dubbo.apache.org.registry.v1alpha1.Reference
	(*RepositoryMainReference)(nil),                   // 1: bufman.dubbo.apache.org.registry.v1alpha1.RepositoryMainReference
	(*RepositoryDraft)(nil),                           // 2: bufman.dubbo.apache.org.registry.v1alpha1.RepositoryDraft
	(*GetReferenceByNameRequest)(nil),                 // 3: bufman.dubbo.apache.org.registry.v1alpha1.GetReferenceByNameRequest
	(*GetReferenceByNameResponse)(nil),                // 4: bufman.dubbo.apache.org.registry.v1alpha1.GetReferenceByNameResponse
	(*ListGitCommitMetadataForReferenceRequest)(nil),  // 5: bufman.dubbo.apache.org.registry.v1alpha1.ListGitCommitMetadataForReferenceRequest
	(*ListGitCommitMetadataForReferenceResponse)(nil), // 6: bufman.dubbo.apache.org.registry.v1alpha1.ListGitCommitMetadataForReferenceResponse
	(*RepositoryTag)(nil),                             // 7: bufman.dubbo.apache.org.registry.v1alpha1.RepositoryTag
	(*RepositoryCommit)(nil),                          // 8: bufman.dubbo.apache.org.registry.v1alpha1.RepositoryCommit
	(*GitCommitMetadata)(nil),                         // 9: bufman.dubbo.apache.org.registry.v1alpha1.GitCommitMetadata
}
var file_registry_v1alpha1_reference_proto_depIdxs = []int32{
	7,  // 0: bufman.dubbo.apache.org.registry.v1alpha1.Reference.tag:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.RepositoryTag
	8,  // 1: bufman.dubbo.apache.org.registry.v1alpha1.Reference.commit:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.RepositoryCommit
	1,  // 2: bufman.dubbo.apache.org.registry.v1alpha1.Reference.main:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.RepositoryMainReference
	2,  // 3: bufman.dubbo.apache.org.registry.v1alpha1.Reference.draft:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.RepositoryDraft
	8,  // 4: bufman.dubbo.apache.org.registry.v1alpha1.RepositoryMainReference.commit:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.RepositoryCommit
	8,  // 5: bufman.dubbo.apache.org.registry.v1alpha1.RepositoryDraft.commit:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.RepositoryCommit
	0,  // 6: bufman.dubbo.apache.org.registry.v1alpha1.GetReferenceByNameResponse.reference:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.Reference
	9,  // 7: bufman.dubbo.apache.org.registry.v1alpha1.ListGitCommitMetadataForReferenceResponse.git_commit_metadatas:type_name -> bufman.dubbo.apache.org.registry.v1alpha1.GitCommitMetadata
	3,  // 8: bufman.dubbo.apache.org.registry.v1alpha1.ReferenceService.GetReferenceByName:input_type -> bufman.dubbo.apache.org.registry.v1alpha1.GetReferenceByNameRequest
	5,  // 9: bufman.dubbo.apache.org.registry.v1alpha1.ReferenceService.ListGitCommitMetadataForReference:input_type -> bufman.dubbo.apache.org.registry.v1alpha1.ListGitCommitMetadataForReferenceRequest
	4,  // 10: bufman.dubbo.apache.org.registry.v1alpha1.ReferenceService.GetReferenceByName:output_type -> bufman.dubbo.apache.org.registry.v1alpha1.GetReferenceByNameResponse
	6,  // 11: bufman.dubbo.apache.org.registry.v1alpha1.ReferenceService.ListGitCommitMetadataForReference:output_type -> bufman.dubbo.apache.org.registry.v1alpha1.ListGitCommitMetadataForReferenceResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_registry_v1alpha1_reference_proto_init() }
func file_registry_v1alpha1_reference_proto_init() {
	if File_registry_v1alpha1_reference_proto != nil {
		return
	}
	file_registry_v1alpha1_git_metadata_proto_init()
	file_registry_v1alpha1_repository_commit_proto_init()
	file_registry_v1alpha1_repository_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_registry_v1alpha1_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_registry_v1alpha1_reference_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryMainReference); i {
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
		file_registry_v1alpha1_reference_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryDraft); i {
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
		file_registry_v1alpha1_reference_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReferenceByNameRequest); i {
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
		file_registry_v1alpha1_reference_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReferenceByNameResponse); i {
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
		file_registry_v1alpha1_reference_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGitCommitMetadataForReferenceRequest); i {
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
		file_registry_v1alpha1_reference_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGitCommitMetadataForReferenceResponse); i {
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
	file_registry_v1alpha1_reference_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Reference_Tag)(nil),
		(*Reference_Commit)(nil),
		(*Reference_Main)(nil),
		(*Reference_Draft)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_registry_v1alpha1_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registry_v1alpha1_reference_proto_goTypes,
		DependencyIndexes: file_registry_v1alpha1_reference_proto_depIdxs,
		MessageInfos:      file_registry_v1alpha1_reference_proto_msgTypes,
	}.Build()
	File_registry_v1alpha1_reference_proto = out.File
	file_registry_v1alpha1_reference_proto_rawDesc = nil
	file_registry_v1alpha1_reference_proto_goTypes = nil
	file_registry_v1alpha1_reference_proto_depIdxs = nil
}
