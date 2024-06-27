// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package publishing

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	publishingpb "cloud.google.com/go/eventarc/publishing/apiv1/publishingpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newPublisherClientHook clientHook

// PublisherCallOptions contains the retry settings for each method of PublisherClient.
type PublisherCallOptions struct {
	PublishChannelConnectionEvents []gax.CallOption
	PublishEvents                  []gax.CallOption
}

func defaultPublisherGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("eventarcpublishing.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("eventarcpublishing.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("eventarcpublishing.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://eventarcpublishing.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPublisherCallOptions() *PublisherCallOptions {
	return &PublisherCallOptions{
		PublishChannelConnectionEvents: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		PublishEvents: []gax.CallOption{},
	}
}

func defaultPublisherRESTCallOptions() *PublisherCallOptions {
	return &PublisherCallOptions{
		PublishChannelConnectionEvents: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		PublishEvents: []gax.CallOption{},
	}
}

// internalPublisherClient is an interface that defines the methods available from Eventarc Publishing API.
type internalPublisherClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	PublishChannelConnectionEvents(context.Context, *publishingpb.PublishChannelConnectionEventsRequest, ...gax.CallOption) (*publishingpb.PublishChannelConnectionEventsResponse, error)
	PublishEvents(context.Context, *publishingpb.PublishEventsRequest, ...gax.CallOption) (*publishingpb.PublishEventsResponse, error)
}

// PublisherClient is a client for interacting with Eventarc Publishing API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Eventarc processes events generated by an event provider and delivers them to
// a subscriber.
//
// An event provider is a software-as-a-service (SaaS) system or
// product that can generate and deliver events through Eventarc.
//
// A third-party event provider is an event provider from outside of Google.
//
// A partner is a third-party event provider that is integrated with Eventarc.
//
// A subscriber is a GCP customer interested in receiving events.
//
// Channel is a first-class Eventarc resource that is created and managed
// by the subscriber in their GCP project. A Channel represents a subscriber’s
// intent to receive events from an event provider. A Channel is associated with
// exactly one event provider.
//
// ChannelConnection is a first-class Eventarc resource that
// is created and managed by the partner in their GCP project. A
// ChannelConnection represents a connection between a partner and a
// subscriber’s Channel. A ChannelConnection has a one-to-one mapping with a
// Channel.
//
// Publisher allows an event provider to publish events to Eventarc.
type PublisherClient struct {
	// The internal transport-dependent client.
	internalClient internalPublisherClient

	// The call options for this service.
	CallOptions *PublisherCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PublisherClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PublisherClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PublisherClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// PublishChannelConnectionEvents publish events to a ChannelConnection in a partner’s project.
func (c *PublisherClient) PublishChannelConnectionEvents(ctx context.Context, req *publishingpb.PublishChannelConnectionEventsRequest, opts ...gax.CallOption) (*publishingpb.PublishChannelConnectionEventsResponse, error) {
	return c.internalClient.PublishChannelConnectionEvents(ctx, req, opts...)
}

// PublishEvents publish events to a subscriber’s channel.
func (c *PublisherClient) PublishEvents(ctx context.Context, req *publishingpb.PublishEventsRequest, opts ...gax.CallOption) (*publishingpb.PublishEventsResponse, error) {
	return c.internalClient.PublishEvents(ctx, req, opts...)
}

// publisherGRPCClient is a client for interacting with Eventarc Publishing API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publisherGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PublisherClient
	CallOptions **PublisherCallOptions

	// The gRPC API client.
	publisherClient publishingpb.PublisherClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewPublisherClient creates a new publisher client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Eventarc processes events generated by an event provider and delivers them to
// a subscriber.
//
// An event provider is a software-as-a-service (SaaS) system or
// product that can generate and deliver events through Eventarc.
//
// A third-party event provider is an event provider from outside of Google.
//
// A partner is a third-party event provider that is integrated with Eventarc.
//
// A subscriber is a GCP customer interested in receiving events.
//
// Channel is a first-class Eventarc resource that is created and managed
// by the subscriber in their GCP project. A Channel represents a subscriber’s
// intent to receive events from an event provider. A Channel is associated with
// exactly one event provider.
//
// ChannelConnection is a first-class Eventarc resource that
// is created and managed by the partner in their GCP project. A
// ChannelConnection represents a connection between a partner and a
// subscriber’s Channel. A ChannelConnection has a one-to-one mapping with a
// Channel.
//
// Publisher allows an event provider to publish events to Eventarc.
func NewPublisherClient(ctx context.Context, opts ...option.ClientOption) (*PublisherClient, error) {
	clientOpts := defaultPublisherGRPCClientOptions()
	if newPublisherClientHook != nil {
		hookOpts, err := newPublisherClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PublisherClient{CallOptions: defaultPublisherCallOptions()}

	c := &publisherGRPCClient{
		connPool:        connPool,
		publisherClient: publishingpb.NewPublisherClient(connPool),
		CallOptions:     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *publisherGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publisherGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publisherGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type publisherRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing PublisherClient
	CallOptions **PublisherCallOptions
}

// NewPublisherRESTClient creates a new publisher rest client.
//
// Eventarc processes events generated by an event provider and delivers them to
// a subscriber.
//
// An event provider is a software-as-a-service (SaaS) system or
// product that can generate and deliver events through Eventarc.
//
// A third-party event provider is an event provider from outside of Google.
//
// A partner is a third-party event provider that is integrated with Eventarc.
//
// A subscriber is a GCP customer interested in receiving events.
//
// Channel is a first-class Eventarc resource that is created and managed
// by the subscriber in their GCP project. A Channel represents a subscriber’s
// intent to receive events from an event provider. A Channel is associated with
// exactly one event provider.
//
// ChannelConnection is a first-class Eventarc resource that
// is created and managed by the partner in their GCP project. A
// ChannelConnection represents a connection between a partner and a
// subscriber’s Channel. A ChannelConnection has a one-to-one mapping with a
// Channel.
//
// Publisher allows an event provider to publish events to Eventarc.
func NewPublisherRESTClient(ctx context.Context, opts ...option.ClientOption) (*PublisherClient, error) {
	clientOpts := append(defaultPublisherRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPublisherRESTCallOptions()
	c := &publisherRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &PublisherClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPublisherRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://eventarcpublishing.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://eventarcpublishing.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://eventarcpublishing.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://eventarcpublishing.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *publisherRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *publisherRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *publisherRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *publisherGRPCClient) PublishChannelConnectionEvents(ctx context.Context, req *publishingpb.PublishChannelConnectionEventsRequest, opts ...gax.CallOption) (*publishingpb.PublishChannelConnectionEventsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "channel_connection", url.QueryEscape(req.GetChannelConnection()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PublishChannelConnectionEvents[0:len((*c.CallOptions).PublishChannelConnectionEvents):len((*c.CallOptions).PublishChannelConnectionEvents)], opts...)
	var resp *publishingpb.PublishChannelConnectionEventsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.PublishChannelConnectionEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *publisherGRPCClient) PublishEvents(ctx context.Context, req *publishingpb.PublishEventsRequest, opts ...gax.CallOption) (*publishingpb.PublishEventsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "channel", url.QueryEscape(req.GetChannel()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PublishEvents[0:len((*c.CallOptions).PublishEvents):len((*c.CallOptions).PublishEvents)], opts...)
	var resp *publishingpb.PublishEventsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.publisherClient.PublishEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PublishChannelConnectionEvents publish events to a ChannelConnection in a partner’s project.
func (c *publisherRESTClient) PublishChannelConnectionEvents(ctx context.Context, req *publishingpb.PublishChannelConnectionEventsRequest, opts ...gax.CallOption) (*publishingpb.PublishChannelConnectionEventsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:publishEvents", req.GetChannelConnection())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "channel_connection", url.QueryEscape(req.GetChannelConnection()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).PublishChannelConnectionEvents[0:len((*c.CallOptions).PublishChannelConnectionEvents):len((*c.CallOptions).PublishChannelConnectionEvents)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &publishingpb.PublishChannelConnectionEventsResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// PublishEvents publish events to a subscriber’s channel.
func (c *publisherRESTClient) PublishEvents(ctx context.Context, req *publishingpb.PublishEventsRequest, opts ...gax.CallOption) (*publishingpb.PublishEventsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:publishEvents", req.GetChannel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "channel", url.QueryEscape(req.GetChannel()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).PublishEvents[0:len((*c.CallOptions).PublishEvents):len((*c.CallOptions).PublishEvents)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &publishingpb.PublishEventsResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
