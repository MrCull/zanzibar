// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: clients/mirror/mirror.proto

package mirror

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// MirrorYARPCClient is the YARPC client-side interface for the Mirror service.
type MirrorYARPCClient interface {
	Mirror(context.Context, *Request, ...yarpc.CallOption) (*Response, error)
}

func newMirrorYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) MirrorYARPCClient {
	return &_MirrorYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "mirror.Mirror",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewMirrorYARPCClient builds a new YARPC client for the Mirror service.
func NewMirrorYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) MirrorYARPCClient {
	return newMirrorYARPCClient(clientConfig, nil, options...)
}

// MirrorYARPCServer is the YARPC server-side interface for the Mirror service.
type MirrorYARPCServer interface {
	Mirror(context.Context, *Request) (*Response, error)
}

type buildMirrorYARPCProceduresParams struct {
	Server      MirrorYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildMirrorYARPCProcedures(params buildMirrorYARPCProceduresParams) []transport.Procedure {
	handler := &_MirrorYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "mirror.Mirror",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Mirror",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Mirror,
							NewRequest:  newMirrorServiceMirrorYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildMirrorYARPCProcedures prepares an implementation of the Mirror service for YARPC registration.
func BuildMirrorYARPCProcedures(server MirrorYARPCServer) []transport.Procedure {
	return buildMirrorYARPCProcedures(buildMirrorYARPCProceduresParams{Server: server})
}

// FxMirrorYARPCClientParams defines the input
// for NewFxMirrorYARPCClient. It provides the
// paramaters to get a MirrorYARPCClient in an
// Fx application.
type FxMirrorYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxMirrorYARPCClientResult defines the output
// of NewFxMirrorYARPCClient. It provides a
// MirrorYARPCClient to an Fx application.
type FxMirrorYARPCClientResult struct {
	fx.Out

	Client MirrorYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxMirrorYARPCClient provides a MirrorYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    mirror.NewFxMirrorYARPCClient("service-name"),
//    ...
//  )
func NewFxMirrorYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxMirrorYARPCClientParams) FxMirrorYARPCClientResult {
		return FxMirrorYARPCClientResult{
			Client: newMirrorYARPCClient(params.Provider.ClientConfig(name), params.AnyResolver, options...),
		}
	}
}

// FxMirrorYARPCProceduresParams defines the input
// for NewFxMirrorYARPCProcedures. It provides the
// paramaters to get MirrorYARPCServer procedures in an
// Fx application.
type FxMirrorYARPCProceduresParams struct {
	fx.In

	Server      MirrorYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxMirrorYARPCProceduresResult defines the output
// of NewFxMirrorYARPCProcedures. It provides
// MirrorYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxMirrorYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxMirrorYARPCProcedures provides MirrorYARPCServer procedures to an Fx application.
// It expects a MirrorYARPCServer to be present in the container.
//
//  fx.Provide(
//    mirror.NewFxMirrorYARPCProcedures(),
//    ...
//  )
func NewFxMirrorYARPCProcedures() interface{} {
	return func(params FxMirrorYARPCProceduresParams) FxMirrorYARPCProceduresResult {
		return FxMirrorYARPCProceduresResult{
			Procedures: buildMirrorYARPCProcedures(buildMirrorYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "mirror.Mirror",
				FileDescriptors: yarpcFileDescriptorClosure735477c7170897b9,
			},
		}
	}
}

type _MirrorYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_MirrorYARPCCaller) Mirror(ctx context.Context, request *Request, options ...yarpc.CallOption) (*Response, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Mirror", request, newMirrorServiceMirrorYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*Response)
	if !ok {
		return nil, protobuf.CastError(emptyMirrorServiceMirrorYARPCResponse, responseMessage)
	}
	return response, err
}

type _MirrorYARPCHandler struct {
	server MirrorYARPCServer
}

func (h *_MirrorYARPCHandler) Mirror(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *Request
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Request)
		if !ok {
			return nil, protobuf.CastError(emptyMirrorServiceMirrorYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Mirror(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newMirrorServiceMirrorYARPCRequest() proto.Message {
	return &Request{}
}

func newMirrorServiceMirrorYARPCResponse() proto.Message {
	return &Response{}
}

var (
	emptyMirrorServiceMirrorYARPCRequest  = &Request{}
	emptyMirrorServiceMirrorYARPCResponse = &Response{}
)

var yarpcFileDescriptorClosure735477c7170897b9 = [][]byte{
	// clients/mirror/mirror.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0xc9, 0x4c,
		0xcd, 0x2b, 0x29, 0xd6, 0xcf, 0xcd, 0x2c, 0x2a, 0xca, 0x2f, 0x82, 0x52, 0x7a, 0x05, 0x45, 0xf9,
		0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92, 0x32, 0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
		0x89, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
		0x06, 0x67, 0x10, 0x8c, 0xab, 0xa4, 0xc2, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
		0x8a, 0x5b, 0x95, 0x91, 0x29, 0x17, 0x9b, 0x2f, 0xd8, 0x50, 0x21, 0x6d, 0x38, 0x8b, 0x5f, 0x0f,
		0x6a, 0x2b, 0xd4, 0x12, 0x29, 0x01, 0x84, 0x00, 0xc4, 0xc0, 0x24, 0x36, 0xb0, 0x83, 0x8c, 0x01,
		0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x62, 0x7b, 0xeb, 0xaf, 0x00, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) MirrorYARPCClient {
			return NewMirrorYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
