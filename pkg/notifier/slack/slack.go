// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

// Package slack provides a notifier.Transport implementation for sending notifications to Slack
package slack

import (
	"context"

	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/identifier"
	"github.com/seatgeek/mailroom/pkg/notifier"
	"github.com/seatgeek/mailroom/pkg/validation"
	"github.com/slack-go/slack"
)

var ID = identifier.NewNamespaceAndKind("slack.com", identifier.KindID)

// Transport supports sending messages to Slack
type Transport struct {
	key    event.TransportKey
	client *slack.Client
}

// RichNotification is an optional interface that can be implemented by any notification supporting Slack formatting
type RichNotification interface {
	event.Notification
	// GetSlackOptions returns the slack.MsgOptions to be used when sending the notification
	GetSlackOptions() []slack.MsgOption
}

// Push sends a notification to a Slack user
// In addition to supporting common.Notification, it also supports RichNotification for more complex messages
// that might include attachments, blocks, etc.
func (s *Transport) Push(ctx context.Context, notification event.Notification) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Transport) getMessageOptions(notification event.Notification) []slack.MsgOption {
	_ = "STUB: not implemented"
	return nil
}

func (s *Transport) Key() event.TransportKey {
	_ = "STUB: not implemented"
	return *new(event.TransportKey)
}

func (s *Transport) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var (
	_ notifier.Transport   = &Transport{}
	_ validation.Validator = &Transport{}
)

// NewTransport creates a new Slack Transport
// It requires a TransportID, a Slack API token, and optionally some slack.Options
func NewTransport(key event.TransportKey, token string, opts ...slack.Option) *Transport {
	_ = "STUB: not implemented"
	return nil
}
