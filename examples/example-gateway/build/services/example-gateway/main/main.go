// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
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

package main

import (
	"flag"
	"runtime"
	"strings"

	"go.uber.org/zap"

	"github.com/uber/zanzibar/config"
	"github.com/uber/zanzibar/runtime"

	service "github.com/uber/zanzibar/examples/example-gateway/build/services/example-gateway"
	module "github.com/uber/zanzibar/examples/example-gateway/build/services/example-gateway/module"
)

var configFiles = flag.String(
	"config",
	"",
	"an ordered, semi-colon separated list of configuration files to use",
)

func getDirName() string {
	_, file, _, _ := runtime.Caller(0)
	return zanzibar.GetDirnameFromRuntimeCaller(file)
}

func getConfig() *zanzibar.StaticConfig {
	var files []string

	if configFiles == nil {
		files = []string{}
	} else {
		files = strings.Split(*configFiles, ";")
	}

	return config.NewRuntimeConfigOrDie(files, nil)
}

func createGateway() (*zanzibar.Gateway, error) {
	config := getConfig()

	gateway, _, err := service.CreateGateway(config, nil)
	if err != nil {
		return nil, err
	}

	return gateway, nil
}

func registerEndpoints(g *zanzibar.Gateway, deps *module.Dependencies) error {
	err0 := deps.Endpoint.Bar.Register(g)
	if err0 != nil {
		return err0
	}
	err1 := deps.Endpoint.Baz.Register(g)
	if err1 != nil {
		return err1
	}
	err2 := deps.Endpoint.BazTChannel.Register(g)
	if err2 != nil {
		return err2
	}
	err3 := deps.Endpoint.Contacts.Register(g)
	if err3 != nil {
		return err3
	}
	err4 := deps.Endpoint.Googlenow.Register(g)
	if err4 != nil {
		return err4
	}
	return nil
}

func logAndWait(server *zanzibar.Gateway) {
	server.Logger.Info("Started ExampleGateway",
		zap.String("realHTTPAddr", server.RealHTTPAddr),
		zap.String("realTChannelAddr", server.RealTChannelAddr),
		zap.Any("config", server.InspectOrDie()),
	)

	// TODO: handle sigterm gracefully
	server.Wait()
	// TODO: emit metrics about startup.
	// TODO: setup and configure tracing/jeager.
}

func main() {
	flag.Parse()
	server, err := createGateway()
	if err != nil {
		panic(err)
	}

	err = server.Bootstrap()
	if err != nil {
		panic(err)
	}

	logAndWait(server)
}
