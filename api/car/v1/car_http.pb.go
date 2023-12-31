// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.3
// source: car/v1/car.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCarAuthToken = "/api.car.v1.Car/AuthToken"
const OperationCarGetUser = "/api.car.v1.Car/GetUser"
const OperationCarGetWechatContacts = "/api.car.v1.Car/GetWechatContacts"
const OperationCarHealthCheck = "/api.car.v1.Car/HealthCheck"
const OperationCarListUser = "/api.car.v1.Car/ListUser"

type CarHTTPServer interface {
	AuthToken(context.Context, *AuthTokenRequest) (*AuthTokenReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	GetWechatContacts(context.Context, *GetWechatContactsRequest) (*GetWechatContactsReply, error)
	HealthCheck(context.Context, *structpb.Value) (*HealthReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
}

func RegisterCarHTTPServer(s *http.Server, srv CarHTTPServer) {
	r := s.Route("/")
	r.GET("/listuser", _Car_ListUser0_HTTP_Handler(srv))
	r.GET("/getuser/{id}", _Car_GetUser0_HTTP_Handler(srv))
	r.POST("/oauth/token", _Car_AuthToken0_HTTP_Handler(srv))
	r.GET("/wechat/contacts", _Car_GetWechatContacts0_HTTP_Handler(srv))
	r.GET("/health", _Car_HealthCheck0_HTTP_Handler(srv))
}

func _Car_ListUser0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarListUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUser(ctx, req.(*ListUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserReply)
		return ctx.Result(200, reply)
	}
}

func _Car_GetUser0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _Car_AuthToken0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthTokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarAuthToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthToken(ctx, req.(*AuthTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Car_GetWechatContacts0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetWechatContactsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarGetWechatContacts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWechatContacts(ctx, req.(*GetWechatContactsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetWechatContactsReply)
		return ctx.Result(200, reply)
	}
}

func _Car_HealthCheck0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in structpb.Value
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarHealthCheck)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HealthCheck(ctx, req.(*structpb.Value))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HealthReply)
		return ctx.Result(200, reply)
	}
}

type CarHTTPClient interface {
	AuthToken(ctx context.Context, req *AuthTokenRequest, opts ...http.CallOption) (rsp *AuthTokenReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	GetWechatContacts(ctx context.Context, req *GetWechatContactsRequest, opts ...http.CallOption) (rsp *GetWechatContactsReply, err error)
	HealthCheck(ctx context.Context, req *structpb.Value, opts ...http.CallOption) (rsp *HealthReply, err error)
	ListUser(ctx context.Context, req *ListUserRequest, opts ...http.CallOption) (rsp *ListUserReply, err error)
}

type CarHTTPClientImpl struct {
	cc *http.Client
}

func NewCarHTTPClient(client *http.Client) CarHTTPClient {
	return &CarHTTPClientImpl{client}
}

func (c *CarHTTPClientImpl) AuthToken(ctx context.Context, in *AuthTokenRequest, opts ...http.CallOption) (*AuthTokenReply, error) {
	var out AuthTokenReply
	pattern := "/oauth/token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCarAuthToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CarHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/getuser/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCarGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CarHTTPClientImpl) GetWechatContacts(ctx context.Context, in *GetWechatContactsRequest, opts ...http.CallOption) (*GetWechatContactsReply, error) {
	var out GetWechatContactsReply
	pattern := "/wechat/contacts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCarGetWechatContacts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CarHTTPClientImpl) HealthCheck(ctx context.Context, in *structpb.Value, opts ...http.CallOption) (*HealthReply, error) {
	var out HealthReply
	pattern := "/health"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCarHealthCheck))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CarHTTPClientImpl) ListUser(ctx context.Context, in *ListUserRequest, opts ...http.CallOption) (*ListUserReply, error) {
	var out ListUserReply
	pattern := "/listuser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCarListUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
