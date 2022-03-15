// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package uri

// Options URI options for creating a new URI
type Options struct {
	scheme      string
	opaque      string    // encoded opaque data
	user        *Userinfo // username and password information
	host        string    // host or host:port
	path        string    // path (relative paths may omit leading slash)
	rawPath     string    // encoded path hint (see EscapedPath method)
	forceQuery  bool      // append a query ('?') even if RawQuery is empty
	rawQuery    string    // encoded query values, without '?'
	fragment    string    // fragment for references, without '#'
	rawFragment string    // encoded fragment hint (see EscapedFragment method)
}

// Option URI option
type Option func(options *Options)

// NewURI creates a new URI
func NewURI(opts ...Option) *URI {
	URIOptions := &Options{}
	for _, option := range opts {
		option(URIOptions)
	}

	return &URI{
		Scheme:      URIOptions.scheme,
		Opaque:      URIOptions.opaque,
		User:        URIOptions.user,
		Path:        URIOptions.path,
		RawPath:     URIOptions.rawPath,
		ForceQuery:  URIOptions.forceQuery,
		RawQuery:    URIOptions.rawQuery,
		Fragment:    URIOptions.fragment,
		RawFragment: URIOptions.rawFragment,
	}

}

// WithScheme sets URI scheme
func WithScheme(scheme string) func(options *Options) {
	return func(options *Options) {
		options.scheme = scheme

	}
}

// WithOpaque  sets URI opaque
func WithOpaque(opaque string) func(options *Options) {
	return func(options *Options) {
		options.opaque = opaque

	}
}

// WithUser  sets URI user information
func WithUser(user *Userinfo) func(options *Options) {
	return func(options *Options) {
		options.user = user

	}
}

// WithHost  sets URI host information
func WithHost(host string) func(options *Options) {
	return func(options *Options) {
		options.host = host

	}
}

// WithPath  sets URI path information
func WithPath(path string) func(options *Options) {
	return func(options *Options) {
		options.path = path

	}
}

// WithRawPath  sets URI raw path information
func WithRawPath(rawPath string) func(options *Options) {
	return func(options *Options) {
		options.rawPath = rawPath

	}
}

// WithForceQuery  sets URI force query information
func WithForceQuery(forceQuery bool) func(options *Options) {
	return func(options *Options) {
		options.forceQuery = forceQuery

	}
}

// WithRawQuery  sets URI raw query information
func WithRawQuery(rawQuery string) func(options *Options) {
	return func(options *Options) {
		options.rawQuery = rawQuery

	}
}

// WithFragment  sets URI fragment information
func WithFragment(fragment string) func(options *Options) {
	return func(options *Options) {
		options.fragment = fragment

	}
}

// WithRawFragment  sets URI raw fragment information
func WithRawFragment(rawFragment string) func(options *Options) {
	return func(options *Options) {
		options.rawFragment = rawFragment

	}
}
