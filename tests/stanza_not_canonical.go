// Copyright 2022 The age Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bytes"

	"filippo.io/age/internal/testkit"
)

func main() {
	f := testkit.NewTestFile()
	f.VersionLine("v1")
	f.X25519(testkit.TestX25519Recipient)
	f.ArgsLine("stanza")
	f.Body(bytes.Repeat([]byte("A"), 50))
	f.TextLine(testkit.NotCanonicalBase64(f.UnreadLine()))
	f.HMAC()
	f.Payload("age")
	f.ExpectHeaderFailure()
	f.Generate()
}
