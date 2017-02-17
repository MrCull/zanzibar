package endpoints

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/clients"
	"github.com/uber/zanzibar/examples/example-gateway/endpoints/bar"
	"github.com/uber/zanzibar/examples/example-gateway/endpoints/contacts"
	"github.com/uber/zanzibar/examples/example-gateway/endpoints/googlenow"
	"github.com/uber/zanzibar/runtime"
)

type handlerFn func(
	ctx context.Context,
	inc *zanzibar.IncomingMessage,
	gateway *zanzibar.Gateway,
	clients *clients.Clients,
)

type myEndpoint struct {
	HandlerFn handlerFn
	Clients   *clients.Clients
}

func (endpoint *myEndpoint) handle(ctx context.Context, inc *zanzibar.IncomingMessage, g *zanzibar.Gateway) {
	fn := endpoint.HandlerFn
	fn(ctx, inc, g, endpoint.Clients)
}

func makeEndpoint(
	g *zanzibar.Gateway,
	endpointName string,
	handlerName string,
	handlerFn handlerFn,
) *zanzibar.Endpoint {
	myEndpoint := &myEndpoint{
		Clients:   g.Clients.(*clients.Clients),
		HandlerFn: handlerFn,
	}

	return zanzibar.NewEndpoint(
		g,
		endpointName,
		handlerName,
		myEndpoint.handle,
	)
}

// Register will register all endpoints
func Register(g *zanzibar.Gateway, router *zanzibar.Router) {
	router.Register(
		"POST", "/contacts/:userUUID/contacts",
		makeEndpoint(
			g,
			"contacts",
			"saveContacts",
			contacts.HandleSaveContactsRequest,
		),
	)

	router.Register(
		"POST", "/googlenow/add-credentials",
		makeEndpoint(
			g,
			"googlenow",
			"addCredentials",
			googlenow.HandleAddCredentialsRequest,
		),
	)

	router.Register(
		"POST", "/googlenow/check-credentials",
		makeEndpoint(
			g,
			"googlenow",
			"addCredentials",
			googlenow.HandleCheckCredentialsRequest,
		),
	)

	router.Register(
		"POST", "/bar/bar-path",
		makeEndpoint(
			g,
			"bar",
			"normal",
			bar.HandleNormalRequest,
		),
	)
	router.Register(
		"GET", "/bar/no-request-path",
		makeEndpoint(
			g,
			"bar",
			"noRequest",
			bar.HandleNoRequestRequest,
		),
	)
	router.Register(
		"GET", "/bar/missing-arg-path",
		makeEndpoint(
			g,
			"bar",
			"missingArg",
			bar.HandleMissingArgRequest,
		),
	)
	router.Register(
		"POST", "/bar/too-many-args-path",
		makeEndpoint(
			g,
			"bar",
			"tooManyArgs",
			bar.HandleTooManyArgsRequest,
		),
	)
	router.Register(
		"POST", "/bar/arg-not-struct-path",
		makeEndpoint(
			g,
			"bar",
			"argNotStrcut",
			bar.HandleArgNotStructRequest,
		),
	)
}
