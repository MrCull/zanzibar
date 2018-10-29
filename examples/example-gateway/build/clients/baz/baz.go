// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bazclient

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

	tchannel "github.com/uber/tchannel-go"
	zanzibar "github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz/module"
	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
)

// Client defines baz client interface.
type Client interface {
	EchoBinary(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoBinary_Args,
	) ([]byte, map[string]string, error)
	EchoBool(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoBool_Args,
	) (bool, map[string]string, error)
	EchoDouble(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoDouble_Args,
	) (float64, map[string]string, error)
	EchoEnum(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoEnum_Args,
	) (clientsBazBaz.Fruit, map[string]string, error)
	EchoI16(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI16_Args,
	) (int16, map[string]string, error)
	EchoI32(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI32_Args,
	) (int32, map[string]string, error)
	EchoI64(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI64_Args,
	) (int64, map[string]string, error)
	EchoI8(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI8_Args,
	) (int8, map[string]string, error)
	EchoString(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoString_Args,
	) (string, map[string]string, error)
	EchoStringList(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoStringList_Args,
	) ([]string, map[string]string, error)
	EchoStringMap(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoStringMap_Args,
	) (map[string]*clientsBazBase.BazResponse, map[string]string, error)
	EchoStringSet(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoStringSet_Args,
	) (map[string]struct{}, map[string]string, error)
	EchoStructList(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoStructList_Args,
	) ([]*clientsBazBase.BazResponse, map[string]string, error)

	EchoStructSet(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoStructSet_Args,
	) ([]*clientsBazBase.BazResponse, map[string]string, error)
	EchoTypedef(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoTypedef_Args,
	) (clientsBazBase.UUID, map[string]string, error)

	Call(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Call_Args,
	) (map[string]string, error)
	Compare(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Compare_Args,
	) (*clientsBazBase.BazResponse, map[string]string, error)
	GetProfile(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_GetProfile_Args,
	) (*clientsBazBaz.GetProfileResponse, map[string]string, error)
	HeaderSchema(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_HeaderSchema_Args,
	) (*clientsBazBaz.HeaderSchema, map[string]string, error)
	Ping(
		ctx context.Context,
		reqHeaders map[string]string,
	) (*clientsBazBase.BazResponse, map[string]string, error)
	DeliberateDiffNoop(
		ctx context.Context,
		reqHeaders map[string]string,
	) (map[string]string, error)
	TestUUID(
		ctx context.Context,
		reqHeaders map[string]string,
	) (map[string]string, error)
	Trans(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Trans_Args,
	) (*clientsBazBase.TransStruct, map[string]string, error)
	TransHeaders(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_TransHeaders_Args,
	) (*clientsBazBase.TransHeaders, map[string]string, error)
	TransHeadersNoReq(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_TransHeadersNoReq_Args,
	) (*clientsBazBase.TransHeaders, map[string]string, error)
	TransHeadersType(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_TransHeadersType_Args,
	) (*clientsBazBaz.TransHeaderType, map[string]string, error)
	URLTest(
		ctx context.Context,
		reqHeaders map[string]string,
	) (map[string]string, error)
}

// NewClient returns a new TChannel client for service baz.
func NewClient(deps *module.Dependencies) Client {
	serviceName := deps.Default.Config.MustGetString("clients.baz.serviceName")
	var routingKey string
	if deps.Default.Config.ContainsKey("clients.baz.routingKey") {
		routingKey = deps.Default.Config.MustGetString("clients.baz.routingKey")
	}
	sc := deps.Default.Channel.GetSubChannel(serviceName, tchannel.Isolated)

	ip := deps.Default.Config.MustGetString("clients.baz.ip")
	port := deps.Default.Config.MustGetInt("clients.baz.port")
	sc.Peers().Add(ip + ":" + strconv.Itoa(int(port)))

	var scAltName string
	if deps.Default.Config.ContainsKey("clients.baz.staging.serviceName") {
		scAltName = deps.Default.Config.MustGetString("clients.baz.staging.serviceName")
		ipAlt := deps.Default.Config.MustGetString("clients.baz.staging.ip")
		portAlt := deps.Default.Config.MustGetInt("clients.baz.staging.port")

		scAlt := deps.Default.Channel.GetSubChannel(scAltName, tchannel.Isolated)
		scAlt.Peers().Add(ipAlt + ":" + strconv.Itoa(int(portAlt)))
	} else if deps.Default.Config.ContainsKey("clients.staging.all.serviceName") {
		scAltName = deps.Default.Config.MustGetString("clients.staging.all.serviceName")
		ipAlt := deps.Default.Config.MustGetString("clients.staging.all.ip")
		portAlt := deps.Default.Config.MustGetInt("clients.staging.all.port")

		scAlt := deps.Default.Channel.GetSubChannel(scAltName, tchannel.Isolated)
		scAlt.Peers().Add(ipAlt + ":" + strconv.Itoa(int(portAlt)))
	}

	timeout := time.Millisecond * time.Duration(
		deps.Default.Config.MustGetInt("clients.baz.timeout"),
	)
	timeoutPerAttempt := time.Millisecond * time.Duration(
		deps.Default.Config.MustGetInt("clients.baz.timeoutPerAttempt"),
	)

	methodNames := map[string]string{
		"SecondService::echoBinary":        "EchoBinary",
		"SecondService::echoBool":          "EchoBool",
		"SecondService::echoDouble":        "EchoDouble",
		"SecondService::echoEnum":          "EchoEnum",
		"SecondService::echoI16":           "EchoI16",
		"SecondService::echoI32":           "EchoI32",
		"SecondService::echoI64":           "EchoI64",
		"SecondService::echoI8":            "EchoI8",
		"SecondService::echoString":        "EchoString",
		"SecondService::echoStringList":    "EchoStringList",
		"SecondService::echoStringMap":     "EchoStringMap",
		"SecondService::echoStringSet":     "EchoStringSet",
		"SecondService::echoStructList":    "EchoStructList",
		"SecondService::echoStructSet":     "EchoStructSet",
		"SecondService::echoTypedef":       "EchoTypedef",
		"SimpleService::call":              "Call",
		"SimpleService::compare":           "Compare",
		"SimpleService::getProfile":        "GetProfile",
		"SimpleService::headerSchema":      "HeaderSchema",
		"SimpleService::ping":              "Ping",
		"SimpleService::sillyNoop":         "DeliberateDiffNoop",
		"SimpleService::testUuid":          "TestUUID",
		"SimpleService::trans":             "Trans",
		"SimpleService::transHeaders":      "TransHeaders",
		"SimpleService::transHeadersNoReq": "TransHeadersNoReq",
		"SimpleService::transHeadersType":  "TransHeadersType",
		"SimpleService::urlTest":           "URLTest",
	}

	client := zanzibar.NewTChannelClientContext(
		deps.Default.Channel,
		deps.Default.Logger,
		deps.Default.ContextMetrics,
		&zanzibar.TChannelClientOption{
			ServiceName:       serviceName,
			ClientID:          "baz",
			MethodNames:       methodNames,
			Timeout:           timeout,
			TimeoutPerAttempt: timeoutPerAttempt,
			RoutingKey:        &routingKey,
			AltSubchannelName: scAltName,
		},
	)

	return &bazClient{
		client: client,
	}
}

// bazClient is the TChannel client for downstream service.
type bazClient struct {
	client *zanzibar.TChannelClient
}

// EchoBinary is a client RPC call for method "SecondService::echoBinary"
func (c *bazClient) EchoBinary(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoBinary_Args,
) ([]byte, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoBinary_Result
	var resp []byte

	logger := c.client.Loggers["SecondService::echoBinary"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoBinary", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoBinary")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoBinary_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoBool is a client RPC call for method "SecondService::echoBool"
func (c *bazClient) EchoBool(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoBool_Args,
) (bool, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoBool_Result
	var resp bool

	logger := c.client.Loggers["SecondService::echoBool"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoBool", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoBool")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoBool_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoDouble is a client RPC call for method "SecondService::echoDouble"
func (c *bazClient) EchoDouble(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoDouble_Args,
) (float64, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoDouble_Result
	var resp float64

	logger := c.client.Loggers["SecondService::echoDouble"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoDouble", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoDouble")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoDouble_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoEnum is a client RPC call for method "SecondService::echoEnum"
func (c *bazClient) EchoEnum(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoEnum_Args,
) (clientsBazBaz.Fruit, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoEnum_Result
	var resp clientsBazBaz.Fruit

	logger := c.client.Loggers["SecondService::echoEnum"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoEnum", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoEnum")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoEnum_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoI16 is a client RPC call for method "SecondService::echoI16"
func (c *bazClient) EchoI16(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI16_Args,
) (int16, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI16_Result
	var resp int16

	logger := c.client.Loggers["SecondService::echoI16"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoI16", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI16")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI16_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoI32 is a client RPC call for method "SecondService::echoI32"
func (c *bazClient) EchoI32(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI32_Args,
) (int32, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI32_Result
	var resp int32

	logger := c.client.Loggers["SecondService::echoI32"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoI32", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI32")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI32_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoI64 is a client RPC call for method "SecondService::echoI64"
func (c *bazClient) EchoI64(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI64_Args,
) (int64, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI64_Result
	var resp int64

	logger := c.client.Loggers["SecondService::echoI64"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoI64", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI64")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI64_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoI8 is a client RPC call for method "SecondService::echoI8"
func (c *bazClient) EchoI8(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI8_Args,
) (int8, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI8_Result
	var resp int8

	logger := c.client.Loggers["SecondService::echoI8"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoI8", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI8")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI8_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoString is a client RPC call for method "SecondService::echoString"
func (c *bazClient) EchoString(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoString_Args,
) (string, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoString_Result
	var resp string

	logger := c.client.Loggers["SecondService::echoString"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoString", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoString")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoString_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoStringList is a client RPC call for method "SecondService::echoStringList"
func (c *bazClient) EchoStringList(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStringList_Args,
) ([]string, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoStringList_Result
	var resp []string

	logger := c.client.Loggers["SecondService::echoStringList"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoStringList", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoStringList")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoStringList_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoStringMap is a client RPC call for method "SecondService::echoStringMap"
func (c *bazClient) EchoStringMap(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStringMap_Args,
) (map[string]*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoStringMap_Result
	var resp map[string]*clientsBazBase.BazResponse

	logger := c.client.Loggers["SecondService::echoStringMap"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoStringMap", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoStringMap")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoStringMap_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoStringSet is a client RPC call for method "SecondService::echoStringSet"
func (c *bazClient) EchoStringSet(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStringSet_Args,
) (map[string]struct{}, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoStringSet_Result
	var resp map[string]struct{}

	logger := c.client.Loggers["SecondService::echoStringSet"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoStringSet", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoStringSet")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoStringSet_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoStructList is a client RPC call for method "SecondService::echoStructList"
func (c *bazClient) EchoStructList(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStructList_Args,
) ([]*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoStructList_Result
	var resp []*clientsBazBase.BazResponse

	logger := c.client.Loggers["SecondService::echoStructList"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoStructList", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoStructList")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoStructList_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoStructSet is a client RPC call for method "SecondService::echoStructSet"
func (c *bazClient) EchoStructSet(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoStructSet_Args,
) ([]*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoStructSet_Result
	var resp []*clientsBazBase.BazResponse

	logger := c.client.Loggers["SecondService::echoStructSet"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoStructSet", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoStructSet")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoStructSet_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// EchoTypedef is a client RPC call for method "SecondService::echoTypedef"
func (c *bazClient) EchoTypedef(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoTypedef_Args,
) (clientsBazBase.UUID, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoTypedef_Result
	var resp clientsBazBase.UUID

	logger := c.client.Loggers["SecondService::echoTypedef"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SecondService", "echoTypedef", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoTypedef")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoTypedef_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// Call is a client RPC call for method "SimpleService::call"
func (c *bazClient) Call(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Call_Args,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_Call_Result

	logger := c.client.Loggers["SimpleService::call"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "call", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for Call")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return nil, err
	}

	return respHeaders, err
}

// Compare is a client RPC call for method "SimpleService::compare"
func (c *bazClient) Compare(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Compare_Args,
) (*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SimpleService_Compare_Result
	var resp *clientsBazBase.BazResponse

	logger := c.client.Loggers["SimpleService::compare"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "compare", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.OtherAuthErr != nil:
			err = result.OtherAuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for Compare")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_Compare_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// GetProfile is a client RPC call for method "SimpleService::getProfile"
func (c *bazClient) GetProfile(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_GetProfile_Args,
) (*clientsBazBaz.GetProfileResponse, map[string]string, error) {
	var result clientsBazBaz.SimpleService_GetProfile_Result
	var resp *clientsBazBaz.GetProfileResponse

	logger := c.client.Loggers["SimpleService::getProfile"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "getProfile", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for GetProfile")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_GetProfile_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// HeaderSchema is a client RPC call for method "SimpleService::headerSchema"
func (c *bazClient) HeaderSchema(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_HeaderSchema_Args,
) (*clientsBazBaz.HeaderSchema, map[string]string, error) {
	var result clientsBazBaz.SimpleService_HeaderSchema_Result
	var resp *clientsBazBaz.HeaderSchema

	logger := c.client.Loggers["SimpleService::headerSchema"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "headerSchema", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.OtherAuthErr != nil:
			err = result.OtherAuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for HeaderSchema")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_HeaderSchema_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// Ping is a client RPC call for method "SimpleService::ping"
func (c *bazClient) Ping(
	ctx context.Context,
	reqHeaders map[string]string,
) (*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SimpleService_Ping_Result
	var resp *clientsBazBase.BazResponse

	logger := c.client.Loggers["SimpleService::ping"]

	args := &clientsBazBaz.SimpleService_Ping_Args{}
	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "ping", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for Ping")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_Ping_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// DeliberateDiffNoop is a client RPC call for method "SimpleService::sillyNoop"
func (c *bazClient) DeliberateDiffNoop(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_SillyNoop_Result

	logger := c.client.Loggers["SimpleService::sillyNoop"]

	args := &clientsBazBaz.SimpleService_SillyNoop_Args{}
	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "sillyNoop", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.ServerErr != nil:
			err = result.ServerErr
		default:
			err = errors.New("bazClient received no result or unknown exception for SillyNoop")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return nil, err
	}

	return respHeaders, err
}

// TestUUID is a client RPC call for method "SimpleService::testUuid"
func (c *bazClient) TestUUID(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_TestUuid_Result

	logger := c.client.Loggers["SimpleService::testUuid"]

	args := &clientsBazBaz.SimpleService_TestUuid_Args{}
	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "testUuid", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for TestUuid")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return nil, err
	}

	return respHeaders, err
}

// Trans is a client RPC call for method "SimpleService::trans"
func (c *bazClient) Trans(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Trans_Args,
) (*clientsBazBase.TransStruct, map[string]string, error) {
	var result clientsBazBaz.SimpleService_Trans_Result
	var resp *clientsBazBase.TransStruct

	logger := c.client.Loggers["SimpleService::trans"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "trans", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.OtherAuthErr != nil:
			err = result.OtherAuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for Trans")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_Trans_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// TransHeaders is a client RPC call for method "SimpleService::transHeaders"
func (c *bazClient) TransHeaders(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_TransHeaders_Args,
) (*clientsBazBase.TransHeaders, map[string]string, error) {
	var result clientsBazBaz.SimpleService_TransHeaders_Result
	var resp *clientsBazBase.TransHeaders

	logger := c.client.Loggers["SimpleService::transHeaders"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "transHeaders", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.OtherAuthErr != nil:
			err = result.OtherAuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for TransHeaders")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_TransHeaders_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// TransHeadersNoReq is a client RPC call for method "SimpleService::transHeadersNoReq"
func (c *bazClient) TransHeadersNoReq(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_TransHeadersNoReq_Args,
) (*clientsBazBase.TransHeaders, map[string]string, error) {
	var result clientsBazBaz.SimpleService_TransHeadersNoReq_Result
	var resp *clientsBazBase.TransHeaders

	logger := c.client.Loggers["SimpleService::transHeadersNoReq"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "transHeadersNoReq", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for TransHeadersNoReq")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_TransHeadersNoReq_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// TransHeadersType is a client RPC call for method "SimpleService::transHeadersType"
func (c *bazClient) TransHeadersType(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_TransHeadersType_Args,
) (*clientsBazBaz.TransHeaderType, map[string]string, error) {
	var result clientsBazBaz.SimpleService_TransHeadersType_Result
	var resp *clientsBazBaz.TransHeaderType

	logger := c.client.Loggers["SimpleService::transHeadersType"]

	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "transHeadersType", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.OtherAuthErr != nil:
			err = result.OtherAuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for TransHeadersType")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_TransHeadersType_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}

// URLTest is a client RPC call for method "SimpleService::urlTest"
func (c *bazClient) URLTest(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_UrlTest_Result

	logger := c.client.Loggers["SimpleService::urlTest"]

	args := &clientsBazBaz.SimpleService_UrlTest_Args{}
	caller := c.client.Call
	if strings.EqualFold(reqHeaders["X-Zanzibar-Use-Staging"], "true") {
		caller = c.client.CallThruAltChannel
	}
	success, respHeaders, err := caller(
		ctx, "SimpleService", "urlTest", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for UrlTest")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return nil, err
	}

	return respHeaders, err
}
