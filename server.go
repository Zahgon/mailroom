// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package mailroom

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/notifier"
	"github.com/seatgeek/mailroom/pkg/notifier/preference"
	"github.com/seatgeek/mailroom/pkg/user"
)

// Server is the heart of the mailroom application
// It listens for incoming webhooks, parses them, generates notifications, and dispatches them to users.
type Server struct {
	listenAddr         string
	parsers            map[string]event.Parser
	processors         []event.Processor
	notifier           notifier.Notifier
	transports         []notifier.Transport
	defaultPreferences preference.Provider
	userStore          user.Store
	router             *mux.Router
}

type Opt func(s *Server)

// New returns a new server
func New(opts ...Opt) *Server { _ = "STUB: not implemented"; return nil }

// WithListenAddr sets the IP and port the server listens on, in the form "host:port"
func WithListenAddr(addr string) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithParser adds an event.Parser to the server with the given key.
// The key is used as the API endpoint for the server.
func WithParser(key string, parser event.Parser) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithParserAndGenerator is a convenience function that adds an event.Parser and its corresponding processor (which generates notifications) in a single call.
func WithParserAndGenerator(key string, parser event.Parser, generator event.Processor) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithProcessors adds event.Processor instances to the server in the order given.
func WithProcessors(processors ...event.Processor) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithTransports adds notifier.Transport instances to the server
func WithTransports(transports ...notifier.Transport) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithUserStore sets the user.Store for the server
func WithUserStore(us user.Store) Opt { _ = "STUB: not implemented"; return *new(Opt) }

// WithDefaultPreferences sets the default preferences for the server
func WithDefaultPreferences(prefs preference.Provider) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

// WithRouter sets the mux.Router used for the server
func WithRouter(router *mux.Router) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func (s *Server) validate(ctx context.Context) error {
	_ = "STUB: not implemented" //nolint:revive // high cognitive complexity okay here
	return nil
}

// Run starts the server in a Goroutine and blocks until the server is shut down.
// If the given context is canceled, the server will attempt to shut down gracefully.
func (s *Server) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) serveHttp(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Mount all parsers

// Expose routes for managing user preferences

// Run the server in a Goroutine

// Wait for the context to be canceled

//nolint:contextcheck

// Or wait for the server to exit on its own (with some error)

func transportKeys(transports []notifier.Transport) []event.TransportKey {
	_ = "STUB: not implemented"
	return nil
}
