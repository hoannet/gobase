// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors implements functions to manipulate errors.
package errors

// New returns an error that formats as the given text.
func New(code string ,text string) BaseError {
	return &coreError{errorcode:code,errortext:text}
}

// coreError is a trivial implementation of error.
type coreError struct {
	errorcode string 
	errortext string
	
}

func (e *coreError) Error() string {
	return e.errortext
}
func (e *coreError) Code() string {
	return e.errorcode
}
