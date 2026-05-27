// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package server

// Error is a custom error type that includes an HTTP status code
type Error struct {
	Code   int
	Reason error
}

func (h *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (h *Error) Is(target error) bool { _ = "STUB: not implemented"; return false }
