// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

// Package server provides the HTTP server for incoming events
package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/notifier"
)

// CreateEventProcessingHandler returns a handlerFunc that can be used to handle incoming webhooks.
// It choreographs the parsing of the incoming request, the generation of notifications, dispatching the notifications
// to the notifier, and returning a success or error response to the client.
func CreateEventProcessingHandler(parserKey string, parser event.Parser, processors []event.Processor, ntfr notifier.Notifier) http.HandlerFunc {
	_ = "STUB: not implemented" //nolint:revive
	return *new(http.HandlerFunc)
}

// Event is ignorable

func logAndSendErrorResponse(ctx context.Context, logger *slog.Logger, writer http.ResponseWriter, errorPrefix string, err error) {
	_ = "STUB: not implemented"
	return
}
