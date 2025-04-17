// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: tests/tests.proto

package testsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TestService_CreateTest_FullMethodName     = "/test.v1.TestService/CreateTest"
	TestService_GetTests_FullMethodName       = "/test.v1.TestService/GetTests"
	TestService_GetTest_FullMethodName        = "/test.v1.TestService/GetTest"
	TestService_SubmitAnswers_FullMethodName  = "/test.v1.TestService/SubmitAnswers"
	TestService_GetTestResults_FullMethodName = "/test.v1.TestService/GetTestResults"
	TestService_UpdateTest_FullMethodName     = "/test.v1.TestService/UpdateTest"
	TestService_DeleteTest_FullMethodName     = "/test.v1.TestService/DeleteTest"
	TestService_PublishTest_FullMethodName    = "/test.v1.TestService/PublishTest"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	CreateTest(ctx context.Context, in *CreateTestRequest, opts ...grpc.CallOption) (*CreateTestResponse, error)
	GetTests(ctx context.Context, in *GetTestsRequest, opts ...grpc.CallOption) (*GetTestsResponse, error)
	GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*GetTestResponse, error)
	SubmitAnswers(ctx context.Context, in *SubmitAnswersRequest, opts ...grpc.CallOption) (*SubmitAnswersResponse, error)
	GetTestResults(ctx context.Context, in *GetTestResultsRequest, opts ...grpc.CallOption) (*GetTestResultsResponse, error)
	UpdateTest(ctx context.Context, in *UpdateTestRequest, opts ...grpc.CallOption) (*UpdateTestResponse, error)
	DeleteTest(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*Pagination, error)
	PublishTest(ctx context.Context, in *PublishTestRequest, opts ...grpc.CallOption) (*PublishTestResponse, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) CreateTest(ctx context.Context, in *CreateTestRequest, opts ...grpc.CallOption) (*CreateTestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTestResponse)
	err := c.cc.Invoke(ctx, TestService_CreateTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) GetTests(ctx context.Context, in *GetTestsRequest, opts ...grpc.CallOption) (*GetTestsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTestsResponse)
	err := c.cc.Invoke(ctx, TestService_GetTests_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*GetTestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTestResponse)
	err := c.cc.Invoke(ctx, TestService_GetTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) SubmitAnswers(ctx context.Context, in *SubmitAnswersRequest, opts ...grpc.CallOption) (*SubmitAnswersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitAnswersResponse)
	err := c.cc.Invoke(ctx, TestService_SubmitAnswers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) GetTestResults(ctx context.Context, in *GetTestResultsRequest, opts ...grpc.CallOption) (*GetTestResultsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTestResultsResponse)
	err := c.cc.Invoke(ctx, TestService_GetTestResults_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) UpdateTest(ctx context.Context, in *UpdateTestRequest, opts ...grpc.CallOption) (*UpdateTestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTestResponse)
	err := c.cc.Invoke(ctx, TestService_UpdateTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) DeleteTest(ctx context.Context, in *DeleteTestRequest, opts ...grpc.CallOption) (*Pagination, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Pagination)
	err := c.cc.Invoke(ctx, TestService_DeleteTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PublishTest(ctx context.Context, in *PublishTestRequest, opts ...grpc.CallOption) (*PublishTestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishTestResponse)
	err := c.cc.Invoke(ctx, TestService_PublishTest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility.
type TestServiceServer interface {
	CreateTest(context.Context, *CreateTestRequest) (*CreateTestResponse, error)
	GetTests(context.Context, *GetTestsRequest) (*GetTestsResponse, error)
	GetTest(context.Context, *GetTestRequest) (*GetTestResponse, error)
	SubmitAnswers(context.Context, *SubmitAnswersRequest) (*SubmitAnswersResponse, error)
	GetTestResults(context.Context, *GetTestResultsRequest) (*GetTestResultsResponse, error)
	UpdateTest(context.Context, *UpdateTestRequest) (*UpdateTestResponse, error)
	DeleteTest(context.Context, *DeleteTestRequest) (*Pagination, error)
	PublishTest(context.Context, *PublishTestRequest) (*PublishTestResponse, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) CreateTest(context.Context, *CreateTestRequest) (*CreateTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTest not implemented")
}
func (UnimplementedTestServiceServer) GetTests(context.Context, *GetTestsRequest) (*GetTestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTests not implemented")
}
func (UnimplementedTestServiceServer) GetTest(context.Context, *GetTestRequest) (*GetTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTest not implemented")
}
func (UnimplementedTestServiceServer) SubmitAnswers(context.Context, *SubmitAnswersRequest) (*SubmitAnswersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAnswers not implemented")
}
func (UnimplementedTestServiceServer) GetTestResults(context.Context, *GetTestResultsRequest) (*GetTestResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestResults not implemented")
}
func (UnimplementedTestServiceServer) UpdateTest(context.Context, *UpdateTestRequest) (*UpdateTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTest not implemented")
}
func (UnimplementedTestServiceServer) DeleteTest(context.Context, *DeleteTestRequest) (*Pagination, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTest not implemented")
}
func (UnimplementedTestServiceServer) PublishTest(context.Context, *PublishTestRequest) (*PublishTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishTest not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}
func (UnimplementedTestServiceServer) testEmbeddedByValue()                     {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_CreateTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).CreateTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_CreateTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).CreateTest(ctx, req.(*CreateTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_GetTests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_GetTests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTests(ctx, req.(*GetTestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_GetTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_GetTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTest(ctx, req.(*GetTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_SubmitAnswers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAnswersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SubmitAnswers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_SubmitAnswers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SubmitAnswers(ctx, req.(*SubmitAnswersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_GetTestResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTestResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_GetTestResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTestResults(ctx, req.(*GetTestResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_UpdateTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UpdateTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_UpdateTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UpdateTest(ctx, req.(*UpdateTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_DeleteTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).DeleteTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_DeleteTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).DeleteTest(ctx, req.(*DeleteTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PublishTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PublishTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_PublishTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PublishTest(ctx, req.(*PublishTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTest",
			Handler:    _TestService_CreateTest_Handler,
		},
		{
			MethodName: "GetTests",
			Handler:    _TestService_GetTests_Handler,
		},
		{
			MethodName: "GetTest",
			Handler:    _TestService_GetTest_Handler,
		},
		{
			MethodName: "SubmitAnswers",
			Handler:    _TestService_SubmitAnswers_Handler,
		},
		{
			MethodName: "GetTestResults",
			Handler:    _TestService_GetTestResults_Handler,
		},
		{
			MethodName: "UpdateTest",
			Handler:    _TestService_UpdateTest_Handler,
		},
		{
			MethodName: "DeleteTest",
			Handler:    _TestService_DeleteTest_Handler,
		},
		{
			MethodName: "PublishTest",
			Handler:    _TestService_PublishTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tests/tests.proto",
}
