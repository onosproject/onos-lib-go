// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type URITest struct {
	in        string
	out       *URI   // expected parse
	roundtrip string // expected result of reserializing the URL; empty means same as "in".
}

var uritests = []URITest{
	// no path
	{
		"http://www.google.com",
		&URI{
			Scheme: "http",
			Host:   "www.google.com",
		},
		"",
	},
	// path
	{
		"http://www.google.com/",
		&URI{
			Scheme: "http",
			Host:   "www.google.com",
			Path:   "/",
		},
		"",
	},
	// path with hex escaping
	{
		"http://www.google.com/file%20one%26two",
		&URI{
			Scheme:  "http",
			Host:    "www.google.com",
			Path:    "/file one&two",
			RawPath: "/file%20one%26two",
		},
		"",
	},
	// fragment with hex escaping
	{
		"http://www.google.com/#file%20one%26two",
		&URI{
			Scheme:      "http",
			Host:        "www.google.com",
			Path:        "/",
			Fragment:    "file one&two",
			RawFragment: "file%20one%26two",
		},
		"",
	},
	// user
	{
		"ftp://webmaster@www.google.com/",
		&URI{
			Scheme: "ftp",
			User:   User("webmaster"),
			Host:   "www.google.com",
			Path:   "/",
		},
		"",
	},
	// escape sequence in username
	{
		"ftp://john%20doe@www.google.com/",
		&URI{
			Scheme: "ftp",
			User:   User("john doe"),
			Host:   "www.google.com",
			Path:   "/",
		},
		"ftp://john%20doe@www.google.com/",
	},
	// empty query
	{
		"http://www.google.com/?",
		&URI{
			Scheme:     "http",
			Host:       "www.google.com",
			Path:       "/",
			ForceQuery: true,
		},
		"",
	},
	// query ending in question mark (Issue 14573)
	{
		"http://www.google.com/?foo=bar?",
		&URI{
			Scheme:   "http",
			Host:     "www.google.com",
			Path:     "/",
			RawQuery: "foo=bar?",
		},
		"",
	},
	// query
	{
		"http://www.google.com/?q=go+language",
		&URI{
			Scheme:   "http",
			Host:     "www.google.com",
			Path:     "/",
			RawQuery: "q=go+language",
		},
		"",
	},
	// query with hex escaping: NOT parsed
	{
		"http://www.google.com/?q=go%20language",
		&URI{
			Scheme:   "http",
			Host:     "www.google.com",
			Path:     "/",
			RawQuery: "q=go%20language",
		},
		"",
	},
	// %20 outside query
	{
		"http://www.google.com/a%20b?q=c+d",
		&URI{
			Scheme:   "http",
			Host:     "www.google.com",
			Path:     "/a b",
			RawQuery: "q=c+d",
		},
		"",
	},
	// path without leading /, so no parsing
	{
		"http:www.google.com/?q=go+language",
		&URI{
			Scheme:   "http",
			Opaque:   "www.google.com/",
			RawQuery: "q=go+language",
		},
		"http:www.google.com/?q=go+language",
	},
	// path without leading /, so no parsing
	{
		"http:%2f%2fwww.google.com/?q=go+language",
		&URI{
			Scheme:   "http",
			Opaque:   "%2f%2fwww.google.com/",
			RawQuery: "q=go+language",
		},
		"http:%2f%2fwww.google.com/?q=go+language",
	},
	// non-authority with path
	{
		"mailto:/webmaster@golang.org",
		&URI{
			Scheme: "mailto",
			Path:   "/webmaster@golang.org",
		},
		"mailto:///webmaster@golang.org", // unfortunate compromise
	},
	// non-authority
	{
		"mailto:webmaster@golang.org",
		&URI{
			Scheme: "mailto",
			Opaque: "webmaster@golang.org",
		},
		"",
	},
	// unescaped :// in query should not create a scheme
	{
		"/foo?query=http://bad",
		&URI{
			Path:     "/foo",
			RawQuery: "query=http://bad",
		},
		"",
	},
	// leading // without scheme should create an authority
	{
		"//foo",
		&URI{
			Host: "foo",
		},
		"",
	},
	// leading // without scheme, with userinfo, path, and query
	{
		"//user@foo/path?a=b",
		&URI{
			User:     User("user"),
			Host:     "foo",
			Path:     "/path",
			RawQuery: "a=b",
		},
		"",
	},
	// Three leading slashes isn't an authority, but doesn't return an error.
	// (We can't return an error, as this code is also used via
	// ServeHTTP -> ReadRequest -> Parse, which is arguably a
	// different URL parsing context, but currently shares the
	// same codepath)
	{
		"///threeslashes",
		&URI{
			Path: "///threeslashes",
		},
		"",
	},
	{
		"http://user:password@google.com",
		&URI{
			Scheme: "http",
			User:   UserPassword("user", "password"),
			Host:   "google.com",
		},
		"http://user:password@google.com",
	},
	// unescaped @ in username should not confuse host
	{
		"http://j@ne:password@google.com",
		&URI{
			Scheme: "http",
			User:   UserPassword("j@ne", "password"),
			Host:   "google.com",
		},
		"http://j%40ne:password@google.com",
	},
	// unescaped @ in password should not confuse host
	{
		"http://jane:p@ssword@google.com",
		&URI{
			Scheme: "http",
			User:   UserPassword("jane", "p@ssword"),
			Host:   "google.com",
		},
		"http://jane:p%40ssword@google.com",
	},
	{
		"http://j@ne:password@google.com/p@th?q=@go",
		&URI{
			Scheme:   "http",
			User:     UserPassword("j@ne", "password"),
			Host:     "google.com",
			Path:     "/p@th",
			RawQuery: "q=@go",
		},
		"http://j%40ne:password@google.com/p@th?q=@go",
	},
	{
		"http://www.google.com/?q=go+language#foo",
		&URI{
			Scheme:   "http",
			Host:     "www.google.com",
			Path:     "/",
			RawQuery: "q=go+language",
			Fragment: "foo",
		},
		"",
	},
	{
		"http://www.google.com/?q=go+language#foo&bar",
		&URI{
			Scheme:   "http",
			Host:     "www.google.com",
			Path:     "/",
			RawQuery: "q=go+language",
			Fragment: "foo&bar",
		},
		"http://www.google.com/?q=go+language#foo&bar",
	},
	{
		"http://www.google.com/?q=go+language#foo%26bar",
		&URI{
			Scheme:      "http",
			Host:        "www.google.com",
			Path:        "/",
			RawQuery:    "q=go+language",
			Fragment:    "foo&bar",
			RawFragment: "foo%26bar",
		},
		"http://www.google.com/?q=go+language#foo%26bar",
	},
	{
		"file:///home/adg/rabbits",
		&URI{
			Scheme: "file",
			Host:   "",
			Path:   "/home/adg/rabbits",
		},
		"file:///home/adg/rabbits",
	},
	// "Windows" paths are no exception to the rule.
	// See golang.org/issue/6027, especially comment #9.
	{
		"file:///C:/FooBar/Baz.txt",
		&URI{
			Scheme: "file",
			Host:   "",
			Path:   "/C:/FooBar/Baz.txt",
		},
		"file:///C:/FooBar/Baz.txt",
	},
	// case-insensitive scheme
	{
		"MaIlTo:webmaster@golang.org",
		&URI{
			Scheme: "mailto",
			Opaque: "webmaster@golang.org",
		},
		"mailto:webmaster@golang.org",
	},
	// Relative path
	{
		"a/b/c",
		&URI{
			Path: "a/b/c",
		},
		"a/b/c",
	},
	// escaped '?' in username and password
	{
		"http://%3Fam:pa%3Fsword@google.com",
		&URI{
			Scheme: "http",
			User:   UserPassword("?am", "pa?sword"),
			Host:   "google.com",
		},
		"",
	},
	// host subcomponent; IPv4 address in RFC 3986
	{
		"http://192.168.0.1/",
		&URI{
			Scheme: "http",
			Host:   "192.168.0.1",
			Path:   "/",
		},
		"",
	},
	// host and port subcomponents; IPv4 address in RFC 3986
	{
		"http://192.168.0.1:8080/",
		&URI{
			Scheme: "http",
			Host:   "192.168.0.1:8080",
			Path:   "/",
		},
		"",
	},
	// host subcomponent; IPv6 address in RFC 3986
	{
		"http://[fe80::1]/",
		&URI{
			Scheme: "http",
			Host:   "[fe80::1]",
			Path:   "/",
		},
		"",
	},
	// host and port subcomponents; IPv6 address in RFC 3986
	{
		"http://[fe80::1]:8080/",
		&URI{
			Scheme: "http",
			Host:   "[fe80::1]:8080",
			Path:   "/",
		},
		"",
	},
	// host subcomponent; IPv6 address with zone identifier in RFC 6874
	{
		"http://[fe80::1%25en0]/", // alphanum zone identifier
		&URI{
			Scheme: "http",
			Host:   "[fe80::1%en0]",
			Path:   "/",
		},
		"",
	},
	// host and port subcomponents; IPv6 address with zone identifier in RFC 6874
	{
		"http://[fe80::1%25en0]:8080/", // alphanum zone identifier
		&URI{
			Scheme: "http",
			Host:   "[fe80::1%en0]:8080",
			Path:   "/",
		},
		"",
	},
	// host subcomponent; IPv6 address with zone identifier in RFC 6874
	{
		"http://[fe80::1%25%65%6e%301-._~]/", // percent-encoded+unreserved zone identifier
		&URI{
			Scheme: "http",
			Host:   "[fe80::1%en01-._~]",
			Path:   "/",
		},
		"http://[fe80::1%25en01-._~]/",
	},
	// host and port subcomponents; IPv6 address with zone identifier in RFC 6874
	{
		"http://[fe80::1%25%65%6e%301-._~]:8080/", // percent-encoded+unreserved zone identifier
		&URI{
			Scheme: "http",
			Host:   "[fe80::1%en01-._~]:8080",
			Path:   "/",
		},
		"http://[fe80::1%25en01-._~]:8080/",
	},
	// alternate escapings of path survive round trip
	{
		"http://rest.rsc.io/foo%2fbar/baz%2Fquux?alt=media",
		&URI{
			Scheme:   "http",
			Host:     "rest.rsc.io",
			Path:     "/foo/bar/baz/quux",
			RawPath:  "/foo%2fbar/baz%2Fquux",
			RawQuery: "alt=media",
		},
		"",
	},
	// issue 12036
	{
		"mysql://a,b,c/bar",
		&URI{
			Scheme: "mysql",
			Host:   "a,b,c",
			Path:   "/bar",
		},
		"",
	},
	// worst case host, still round trips
	{
		"scheme://!$&'()*+,;=hello!:1/path",
		&URI{
			Scheme: "scheme",
			Host:   "!$&'()*+,;=hello!:1",
			Path:   "/path",
		},
		"",
	},
	// worst case path, still round trips
	{
		"http://host/!$&'()*+,;=:@[hello]",
		&URI{
			Scheme:  "http",
			Host:    "host",
			Path:    "/!$&'()*+,;=:@[hello]",
			RawPath: "/!$&'()*+,;=:@[hello]",
		},
		"",
	},
	// golang.org/issue/5684
	{
		"http://example.com/oid/[order_id]",
		&URI{
			Scheme:  "http",
			Host:    "example.com",
			Path:    "/oid/[order_id]",
			RawPath: "/oid/[order_id]",
		},
		"",
	},
	// golang.org/issue/12200 (colon with empty port)
	{
		"http://192.168.0.2:8080/foo",
		&URI{
			Scheme: "http",
			Host:   "192.168.0.2:8080",
			Path:   "/foo",
		},
		"",
	},
	{
		"http://192.168.0.2:/foo",
		&URI{
			Scheme: "http",
			Host:   "192.168.0.2:",
			Path:   "/foo",
		},
		"",
	},
	{
		// Malformed IPv6 but still accepted.
		"http://2b01:e34:ef40:7730:8e70:5aff:fefe:edac:8080/foo",
		&URI{
			Scheme: "http",
			Host:   "2b01:e34:ef40:7730:8e70:5aff:fefe:edac:8080",
			Path:   "/foo",
		},
		"",
	},
	{
		// Malformed IPv6 but still accepted.
		"http://2b01:e34:ef40:7730:8e70:5aff:fefe:edac:/foo",
		&URI{
			Scheme: "http",
			Host:   "2b01:e34:ef40:7730:8e70:5aff:fefe:edac:",
			Path:   "/foo",
		},
		"",
	},
	{
		"http://[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:8080/foo",
		&URI{
			Scheme: "http",
			Host:   "[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:8080",
			Path:   "/foo",
		},
		"",
	},
	{
		"http://[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:/foo",
		&URI{
			Scheme: "http",
			Host:   "[2b01:e34:ef40:7730:8e70:5aff:fefe:edac]:",
			Path:   "/foo",
		},
		"",
	},
	// golang.org/issue/7991 and golang.org/issue/12719 (non-ascii %-encoded in host)
	{
		"http://hello.世界.com/foo",
		&URI{
			Scheme: "http",
			Host:   "hello.世界.com",
			Path:   "/foo",
		},
		"http://hello.%E4%B8%96%E7%95%8C.com/foo",
	},
	{
		"http://hello.%e4%b8%96%e7%95%8c.com/foo",
		&URI{
			Scheme: "http",
			Host:   "hello.世界.com",
			Path:   "/foo",
		},
		"http://hello.%E4%B8%96%E7%95%8C.com/foo",
	},
	{
		"http://hello.%E4%B8%96%E7%95%8C.com/foo",
		&URI{
			Scheme: "http",
			Host:   "hello.世界.com",
			Path:   "/foo",
		},
		"",
	},
	// golang.org/issue/10433 (path beginning with //)
	{
		"http://example.com//foo",
		&URI{
			Scheme: "http",
			Host:   "example.com",
			Path:   "//foo",
		},
		"",
	},
	// test that we can reparse the host names we accept.
	{
		"myscheme://authority<\"hi\">/foo",
		&URI{
			Scheme: "myscheme",
			Host:   "authority<\"hi\">",
			Path:   "/foo",
		},
		"",
	},
	// spaces in hosts are disallowed but escaped spaces in IPv6 scope IDs are grudgingly OK.
	// This happens on Windows.
	// golang.org/issue/14002
	{
		"tcp://[2020::2020:20:2020:2020%25Windows%20Loves%20Spaces]:2020",
		&URI{
			Scheme: "tcp",
			Host:   "[2020::2020:20:2020:2020%Windows Loves Spaces]:2020",
		},
		"",
	},
	// test we can roundtrip magnet url
	// fix issue https://golang.org/issue/20054
	{
		in: "magnet:?xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn",
		out: &URI{
			Scheme:   "magnet",
			Host:     "",
			Path:     "",
			RawQuery: "xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn",
		},
		roundtrip: "magnet:?xt=urn:btih:c12fe1c06bba254a9dc9f519b335aa7c1367a88a&dn",
	},
	{
		in: "mailto:?subject=hi",
		out: &URI{
			Scheme:   "mailto",
			Host:     "",
			Path:     "",
			RawQuery: "subject=hi",
		},
		roundtrip: "mailto:?subject=hi",
	},
	{
		in: "e2:1/1234",
		out: &URI{
			Scheme: "e2",
			Opaque: "1/1234",
		},
		roundtrip: "",
	},
	{
		in: "e2:1/cafe5153f00d/3/14ea32ca",
		out: &URI{
			Scheme: "e2",
			Opaque: "1/cafe5153f00d/3/14ea32ca",
		},
		roundtrip: "",
	},
	{
		in: "uuid:f716c131-e401-4f7b-9924-16361ae7825a",
		out: &URI{
			Scheme: "uuid",
			Opaque: "f716c131-e401-4f7b-9924-16361ae7825a",
		},
	},
	{
		in: "tel:+1-816-555-1212",
		out: &URI{
			Scheme: "tel",
			Opaque: "+1-816-555-1212",
		},
	},
	{
		in: "news:comp.infosystems.www.servers.unix",
		out: &URI{
			Scheme: "news",
			Opaque: "comp.infosystems.www.servers.unix",
		},
	},
	{
		in: "urn:oasis:names:specification:docbook:dtd:xml:4.1.2",
		out: &URI{
			Scheme: "urn",
			Opaque: "oasis:names:specification:docbook:dtd:xml:4.1.2",
		},
	},
	{
		in: "ldap://[2001:db8::7]/c=GB?objectClass?one",
		out: &URI{
			Scheme:   "ldap",
			Host:     "[2001:db8::7]",
			Path:     "/c=GB",
			RawQuery: "objectClass?one",
		},
	},
}

func TestParse(t *testing.T) {
	for _, tt := range uritests {
		u, err := Parse(tt.in)
		assert.NoError(t, err)
		assert.True(t, reflect.DeepEqual(u, tt.out))
	}
}

var stringURITests = []struct {
	uri  URI
	want string
}{
	// No leading slash on path should prepend slash on String() call
	{
		uri: URI{
			Scheme: "http",
			Host:   "www.google.com",
			Path:   "search",
		},
		want: "http://www.google.com/search",
	},
	// Relative path with first element containing ":" should be prepended with "./", golang.org/issue/17184
	{
		uri: URI{
			Path: "this:that",
		},
		want: "./this:that",
	},
	// Relative path with second element containing ":" should not be prepended with "./"
	{
		uri: URI{
			Path: "here/this:that",
		},
		want: "here/this:that",
	},
	// Non-relative path with first element containing ":" should not be prepended with "./"
	{
		uri: URI{
			Scheme: "http",
			Host:   "www.google.com",
			Path:   "this:that",
		},
		want: "http://www.google.com/this:that",
	},
}

func TestURIString(t *testing.T) {
	for _, tt := range uritests {
		u, err := Parse(tt.in)
		assert.NoError(t, err)
		expected := tt.in
		if tt.roundtrip != "" {
			expected = tt.roundtrip
		}
		s := u.String()
		assert.Equal(t, expected, s)
	}

	for _, tt := range stringURITests {
		assert.Equal(t, tt.want, tt.uri.String())
	}
}

type RequestURITest struct {
	uri *URI
	out string
}

var requritests = []RequestURITest{
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("")),
		out: "/",
	},
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("/a b")),
		out: "/a%20b",
	},
	// golang.org/issue/4860 variant 1
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithOpaque("/%2F/%2F/")),
		out: "/%2F/%2F/",
	},
	// golang.org/issue/4860 variant 2
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithOpaque("//other.example.com/%2F/%2F/")),
		out: "http://other.example.com/%2F/%2F/",
	},
	// better fix for issue 4860
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("/////"),
			WithRawPath("/%2F/%2F/")),
		out: "/%2F/%2F/",
	},
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("/////"), WithRawPath("/WRONG/")),
		out: "/////",
	},
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("/a b"), WithRawQuery("q=go+language")),
		out: "/a%20b?q=go+language",
	},
	{
		uri: NewURI(WithScheme("http"), WithHost("example.com"),
			WithPath("/a b"), WithRawPath("/a b"), WithRawQuery("q=go+language")),
		out: "/a%20b?q=go+language",
	},
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("/a?b"), WithRawPath("/a?b"),
			WithRawQuery("q=go+language")),
		out: "/a%3Fb?q=go+language",
	},
	{
		uri: NewURI(WithScheme("myschema"),
			WithOpaque("opaque")),
		out: "opaque",
	},
	{
		uri: NewURI(WithScheme("myschema"),
			WithOpaque("opaque"),
			WithRawQuery("q=go+language")),
		out: "opaque?q=go+language",
	},
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("//foo")),
		out: "//foo",
	},
	{
		uri: NewURI(WithScheme("http"),
			WithHost("example.com"),
			WithPath("/foo"),
			WithForceQuery(true)),
		out: "/foo?",
	},
}

func TestRequestURI(t *testing.T) {
	for _, tt := range requritests {
		s := tt.uri.RequestURI()
		assert.Equal(t, tt.out, s)
	}
}
