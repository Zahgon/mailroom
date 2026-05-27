// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package notification

import (
	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/identifier"
	slack2 "github.com/seatgeek/mailroom/pkg/notifier/slack"
	"github.com/slack-go/slack"
)

type builderOpts struct {
	context             event.Context
	recipients          identifier.Set
	fallbackMessage     string
	messagePerTransport map[event.TransportKey]string
	slackOpts           []slack.MsgOption
}

// Builder provides a fluent interface for constructing rich notification objects
type Builder struct {
	opts builderOpts
}

// NewBuilder creates a new fluent Builder instance
func NewBuilder(context event.Context) *Builder { _ = "STUB: not implemented"; return nil }

// WithRecipient sets the recipient of the notification
// It's like WithRecipientIdentifiers, but it accepts a single identifier set
func (b *Builder) WithRecipient(identifiers identifier.Set) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithRecipientIdentifiers sets the recipient of the notification
// It's like WithRecipient but it accepts multiple identifiers as variadic arguments
func (b *Builder) WithRecipientIdentifiers(identifiers ...identifier.Identifier) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithDefaultMessage sets the default message to be used if no message is provided for a specific transport
func (b *Builder) WithDefaultMessage(message string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithMessageForTransport sets a specific message to be used for a specific transport
func (b *Builder) WithMessageForTransport(transportKey event.TransportKey, message string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// WithSlackOptions sets the Slack options (like attachments, blocks, etc.) to be used when sending the notification
func (b *Builder) WithSlackOptions(opts ...slack.MsgOption) *Builder {
	_ = "STUB: not implemented"
	return nil
}

// Build constructs the rich notification object from the previously set options
func (b *Builder) Build() slack2.RichNotification {
	_ = "STUB: not implemented"
	return *new(slack2.RichNotification)
}

var _ slack2.RichNotification = &builderOpts{}

func (b *builderOpts) Context() event.Context {
	_ = "STUB: not implemented"
	return *new(event.Context)
}

func (b *builderOpts) Recipient() identifier.Set {
	_ = "STUB: not implemented"
	return *new(identifier.Set)
}

func (b *builderOpts) Render(key event.TransportKey) string { _ = "STUB: not implemented"; return "" }

func (b *builderOpts) GetSlackOptions() []slack.MsgOption { _ = "STUB: not implemented"; return nil }

func (b *builderOpts) WithRecipient(recipient identifier.Set) event.Notification {
	_ = "STUB: not implemented"
	return *new(event.Notification)
}

func (b *builderOpts) Copy() event.Notification {
	_ = "STUB: not implemented"
	return *new(event.Notification)
}
