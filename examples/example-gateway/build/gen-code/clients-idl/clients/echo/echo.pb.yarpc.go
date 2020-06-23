// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: clients-idl/clients/echo/echo.proto

package echo

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// EchoYARPCClient is the YARPC client-side interface for the Echo service.
type EchoYARPCClient interface {
	Echo(context.Context, *Request, ...yarpc.CallOption) (*Response, error)
}

func newEchoYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) EchoYARPCClient {
	return &_EchoYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "echo.Echo",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewEchoYARPCClient builds a new YARPC client for the Echo service.
func NewEchoYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) EchoYARPCClient {
	return newEchoYARPCClient(clientConfig, nil, options...)
}

// EchoYARPCServer is the YARPC server-side interface for the Echo service.
type EchoYARPCServer interface {
	Echo(context.Context, *Request) (*Response, error)
}

type buildEchoYARPCProceduresParams struct {
	Server      EchoYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildEchoYARPCProcedures(params buildEchoYARPCProceduresParams) []transport.Procedure {
	handler := &_EchoYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "echo.Echo",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Echo",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Echo,
							NewRequest:  newEchoServiceEchoYARPCRequest,
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

// BuildEchoYARPCProcedures prepares an implementation of the Echo service for YARPC registration.
func BuildEchoYARPCProcedures(server EchoYARPCServer) []transport.Procedure {
	return buildEchoYARPCProcedures(buildEchoYARPCProceduresParams{Server: server})
}

// FxEchoYARPCClientParams defines the input
// for NewFxEchoYARPCClient. It provides the
// paramaters to get a EchoYARPCClient in an
// Fx application.
type FxEchoYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxEchoYARPCClientResult defines the output
// of NewFxEchoYARPCClient. It provides a
// EchoYARPCClient to an Fx application.
type FxEchoYARPCClientResult struct {
	fx.Out

	Client EchoYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxEchoYARPCClient provides a EchoYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    echo.NewFxEchoYARPCClient("service-name"),
//    ...
//  )
func NewFxEchoYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxEchoYARPCClientParams) FxEchoYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxEchoYARPCClientResult{
			Client: newEchoYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxEchoYARPCProceduresParams defines the input
// for NewFxEchoYARPCProcedures. It provides the
// paramaters to get EchoYARPCServer procedures in an
// Fx application.
type FxEchoYARPCProceduresParams struct {
	fx.In

	Server      EchoYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxEchoYARPCProceduresResult defines the output
// of NewFxEchoYARPCProcedures. It provides
// EchoYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxEchoYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxEchoYARPCProcedures provides EchoYARPCServer procedures to an Fx application.
// It expects a EchoYARPCServer to be present in the container.
//
//  fx.Provide(
//    echo.NewFxEchoYARPCProcedures(),
//    ...
//  )
func NewFxEchoYARPCProcedures() interface{} {
	return func(params FxEchoYARPCProceduresParams) FxEchoYARPCProceduresResult {
		return FxEchoYARPCProceduresResult{
			Procedures: buildEchoYARPCProcedures(buildEchoYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "echo.Echo",
				FileDescriptors: yarpcFileDescriptorClosurec188a5dbce06c9a4,
			},
		}
	}
}

type _EchoYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_EchoYARPCCaller) Echo(ctx context.Context, request *Request, options ...yarpc.CallOption) (*Response, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Echo", request, newEchoServiceEchoYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*Response)
	if !ok {
		return nil, protobuf.CastError(emptyEchoServiceEchoYARPCResponse, responseMessage)
	}
	return response, err
}

type _EchoYARPCHandler struct {
	server EchoYARPCServer
}

func (h *_EchoYARPCHandler) Echo(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *Request
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Request)
		if !ok {
			return nil, protobuf.CastError(emptyEchoServiceEchoYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Echo(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newEchoServiceEchoYARPCRequest() proto.Message {
	return &Request{}
}

func newEchoServiceEchoYARPCResponse() proto.Message {
	return &Response{}
}

var (
	emptyEchoServiceEchoYARPCRequest  = &Request{}
	emptyEchoServiceEchoYARPCResponse = &Response{}
)

var yarpcFileDescriptorClosurec188a5dbce06c9a4 = [][]byte{
	// clients-idl/clients/echo/echo.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0xc9, 0x4c,
		0xcd, 0x2b, 0x29, 0xd6, 0xcd, 0x4c, 0xc9, 0xd1, 0x87, 0xb2, 0xf5, 0x53, 0x93, 0x33, 0xf2, 0xc1,
		0x84, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x0b, 0x88, 0xad, 0xa4, 0xcc, 0xc5, 0x1e, 0x94,
		0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e,
		0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x2a, 0xa9, 0x70, 0x71, 0x04, 0xa5, 0x16,
		0x17, 0xe4, 0xe7, 0x15, 0xa7, 0xe2, 0x56, 0x65, 0xa4, 0xcb, 0xc5, 0xe2, 0x9a, 0x9c, 0x91, 0x2f,
		0xa4, 0x0a, 0xa5, 0x79, 0xf5, 0xc0, 0xb6, 0x41, 0x8d, 0x97, 0xe2, 0x83, 0x71, 0x21, 0x06, 0x25,
		0xb1, 0x81, 0x9d, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xa8, 0x0e, 0xf1, 0xad, 0x00,
		0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) EchoYARPCClient {
			return NewEchoYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}