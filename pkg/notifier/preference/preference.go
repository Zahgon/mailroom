// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package preference

import (
	"context"

	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/validation"
)

// Provider provides a mechanism for determining whether a user wants to receive some notification via some transport.
type Provider interface {
	// Wants returns whether the user wants to receive the given event via the given transport.
	// It returns nil if there is no explicit preference.
	Wants(context.Context, event.Notification, event.TransportKey) *bool
}

// Func is a function type that implements the Provider interface.
type Func func(context.Context, event.Notification, event.TransportKey) *bool

func (f Func) Wants(ctx context.Context, notification event.Notification, transport event.TransportKey) *bool {
	_ = "STUB: not implemented"
	return nil
}

// Chain is a sequence of Provider instances that will be checked in order until one returns a non-nil value.
type Chain []Provider

var _ validation.Validator = (*Chain)(nil)

func (c Chain) Wants(ctx context.Context, notification event.Notification, transport event.TransportKey) *bool {
	_ = "STUB: not implemented"
	return nil
}

// If no preferences matched, return nil to indicate no explicit preference.

func (c Chain) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Map defines preferences by event type and transport.
// For example, a user may want to receive PR review request notifications via Slack but not email.
type Map map[event.Type]map[event.TransportKey]bool

func (p Map) Wants(_ context.Context, notification event.Notification, transport event.TransportKey) *bool {
	_ = "STUB: not implemented"
	return nil
}

// No preference set for this event

// No preference set for this transport

// Default returns a Provider implementation that always returns the given boolean value.
func Default(wants bool) Func { _ = "STUB: not implemented"; return *new(Func) }
