// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

// StackTrace stack
type StackTrace struct {
	entries []StackTraceEntry
}

// StackTraceEntry of StackTrace
type StackTraceEntry struct {
	file string
	line int
}
