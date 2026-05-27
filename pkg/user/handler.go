// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package user

import (
	"context"
	"net/http"

	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/notifier/preference"
)

// PreferencesHandler exposes an HTTP API for managing user preferences
type PreferencesHandler struct {
	userStore  Store
	parsers    map[string]event.Parser
	transports []event.TransportKey
	defaults   preference.Provider
}

// NewPreferencesHandler creates a new PreferencesHandler for managing user preferences
func NewPreferencesHandler(userStore Store, parsers map[string]event.Parser, transports []event.TransportKey, defaults preference.Provider) *PreferencesHandler {
	_ = "STUB: not implemented"
	return nil
}

type preferencesBody struct {
	Preferences preference.Map `json:"preferences"`
}

// GetPreferences returns the preferences for a given user
func (ph *PreferencesHandler) GetPreferences(writer http.ResponseWriter, request *http.Request) {
	_ = "STUB: not implemented"
	return
}

// UpdatePreferences updates the preferences for a given user
func (ph *PreferencesHandler) UpdatePreferences(writer http.ResponseWriter, request *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Builds a current mapping of user preferences based on what is stored in the
// user store and the parsers and transports that are registered with the server.
//
// Only event types and transports that are currently active in the server will
// be included in the preference map. User is opted in to any preference that is
// not stored.
func (ph *PreferencesHandler) buildCurrentUserPreferences(ctx context.Context, p preference.Provider) preference.Map {
	_ = "STUB: not implemented"
	return *new(preference.Map)
}

// No preference or fallback was available, so show it as enabled by default.

// fakeNotificationFor creates a fake notification for the given event type.
// This is needed because preferences are based on notifications and their context,
// so we need to simulate such a notification to check preferences against it.
// Most preferences are usually based on just the event type, and will usually return
// nil to fall back to some default later in the chain - this is why we can get away
// with using a fake notification here.
//
// But if you're reading this and thinking "this dirty hack doesn't work for my needs"
// then please open an issue to explain your use case so we can improve this!
func fakeNotificationFor(eventType event.Type) event.Notification {
	_ = "STUB: not implemented"
	return *new(event.Notification)
}

type transport struct {
	Key event.TransportKey `json:"key"`
}

type source struct {
	Key        string                 `json:"key"`
	EventTypes []event.TypeDescriptor `json:"event_types"`
}

type availableOptions struct {
	Sources    []source    `json:"sources"`
	Transports []transport `json:"transports"`
}

// ListOptions returns the available sources and transports for setting preferences
func (ph *PreferencesHandler) ListOptions(writer http.ResponseWriter, request *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Sort sources by key (asc)

func writeJson(ctx context.Context, writer http.ResponseWriter, value any) {
	_ = "STUB: not implemented"
	return
}
