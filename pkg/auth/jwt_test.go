// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"github.com/dgrijalva/jwt-go"
	"gotest.tools/assert"
	"os"
	"testing"
)

// generated with "openssl genrsa -out /tmp/seans_private_key.pem"
const samplePrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEArpsPCTlPszvklnHXb97xsmeF2NdmBW54ibCwR6gfGHvDI+am
hU++F7utIZPDMB2+H8pOk79lDNNieZmIWRHqDmZMkQygY+OuTAkPMm/35dMFpAWt
jTzYM8Z1gKp/WRdoY1nbZuo3aAo1EHZYh7X6n29YfalgLCLCTo3GnieDjCi4bjMO
DNWYrdEm1qXghlaBuoNYqbhOlSvQWRPuRQsNJKQ4BaPeP41Eb3w9UvbsIjfSs6ks
LoljxbINe8noesr0Lae2nSYugVhrPdhXVS5AFz2lngvlxJCZsQydbitQiUVsJYns
weOq6k0Nn/4FUGIp9bXvEhoXR639R/KSaWhJZQIDAQABAoIBACKrbUfill27db8d
qa5v8UQAZEZTNtG7RrnoWIhR7KK66Ft3j/cGh3NE87KoGWizby32yLVzmof6bSJC
Bx3Qfc4QKAHhJPPQoKo+XkMgknOS/Bq+eeCChVd6f5hlwlWZXUPk8rizpv7EkBbN
uPRxgOspe8Ov3wXEfRqF8jszE16/3y+7cFeaR2xK+7DwvmhMOZeGGUFdqSFpzA3j
9nPdjlvob8WqLhuTMvjM83u2q3rGOcmkUEvpfqB+Jhrr+5gTkYkXfdcvS5lAFF74
q5QuXgjxn2uZhZvJpAY5sLYkqGxGeBTAIM9rcQvL4OVTgUQxeimQKHbUL+Du0m2H
of5qS90CgYEA1dlY0Lz6pHg9il6K1IW4CXm69Zk/WlFOqtyBBYt8FhdwT23wqe/W
igUAIoYe46wwk1Q1fo55o+sZHnQ30WIjdrxgsFKwBwPK5WCt0PJKQtJX95bmq+Jl
JoIYwcuWl+mwCw0i+d+lrvxEDiVi4MxXRAZAH1t5DgTeJXZD24vsRKMCgYEA0QWG
abZWZjNxqBPYaKji5LGNaapLYfBGuGNph+E87P1ezRugKmu9wp6UVfZeqzQ6qSe8
qjept28rZ4XoFr20WAglduI7tgwbaM5zPm09TIEfMgbkC+wPZoRRV/MiE93KXs8Z
CLqx9vzAc4sfhx74OMu1bqETH6BcdLiZQsaoklcCgYEAplo2GeL4Qxr6HHphGuOO
f2h/hHAa9UJMpON1Rn/0HidLia5nSXq19JXhPfoBa3BWNTWLi5B/lYDcAHG9vhbO
qZ3uxRr9redIXVjwvZrNI+AG6CYt+MXbk7IeWhrqYfA6rs4gSCqu80lwE2UH3wF2
XQdTuFDrAXnN6WxvawkU2WsCgYA5Xs7ZzRZBzvTvMSNA9rnwE1vBGOC/7Pc8PO0G
Qqg6VADlQyPfANuAw43rWkf5rcg8DZAXGFgY5QaAz5w4QbFySCogS1AgU4piZefz
xoAAs6AgVwvYyd4gQUkiXrWHxmR5SVaRssyOAinAjPsGV1XCIQeXadaZ46X8034d
efFi2QKBgF66NE55vBPexsV9hpEYenqJ7WJ3S4rYWOYz1febZbNoCsESazAIgzR1
Hme4aXuj/GXzaAQQ2VEnGhVg9+Bvj0pakS6xQu0XS8SWaRzDvwCwgn2qwcoAe8hF
WrVYLktjh2tnKm5qihBnJgj24ElFMmHn1crr1z7ccAa3HyQL2gUk
-----END RSA PRIVATE KEY-----
`

// generated from above with "openssl rsa -in /tmp/seans_private_key.pem -RSAPublicKey_out -out /tmp/seans_public_key.pem"
const samplePublicKey = `
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEArpsPCTlPszvklnHXb97xsmeF2NdmBW54ibCwR6gfGHvDI+amhU++
F7utIZPDMB2+H8pOk79lDNNieZmIWRHqDmZMkQygY+OuTAkPMm/35dMFpAWtjTzY
M8Z1gKp/WRdoY1nbZuo3aAo1EHZYh7X6n29YfalgLCLCTo3GnieDjCi4bjMODNWY
rdEm1qXghlaBuoNYqbhOlSvQWRPuRQsNJKQ4BaPeP41Eb3w9UvbsIjfSs6ksLolj
xbINe8noesr0Lae2nSYugVhrPdhXVS5AFz2lngvlxJCZsQydbitQiUVsJYnsweOq
6k0Nn/4FUGIp9bXvEhoXR639R/KSaWhJZQIDAQAB
-----END RSA PUBLIC KEY-----
`

// generated from running onos-gui against the Dex IDP
const sampleTokenOnosGuiAndDex = `eyJhbGciOiJSUzI1NiIsImtpZCI6IjQwOGFkNDcyNDg3ZjU4YTVmY2I2YzU1NjM3OTg1N2EzNmFjYzBmYTMifQ.eyJpc3MiOiJodHRwOi8vZGV4OjMyMDAwIiwic3ViIjoiQ2lRd09HRTROamcwWWkxa1lqZzRMVFJpTnpNdE9UQmhPUzB6WTJReE5qWXhaalUwTmpnU0JXeHZZMkZzIiwiYXVkIjoib25vcy1ndWkiLCJleHAiOjE1OTQzNzYzNTMsImlhdCI6MTU5NDI4OTk1Mywibm9uY2UiOiJjbE15VmkxWmJWazNWMUpyZEcxTU5uTnFRUzA0Y3pOemRUQkhWMmhPUTFoYWVVcDVRVE4wVmtsUVRHMDQiLCJhdF9oYXNoIjoiQm1xaFVKcW16N25mcEFOV2Y0M0NpZyIsImVtYWlsIjoic2VhbkBvcGVubmV0d29ya2luZy5vcmciLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InNlYW4ifQ.rA3uYnbucPUeIST-HFefoy6zOdxJFCoU7KLXlo02ZqbUBC354NzrY6IvkHCmXTIv2fjCOxzfBCANhK4hpPsL_6fNUz_judao17zwT5UxXyhxmMLsHzlIwfdx_1lKS6hdU5O-WI6KP8nv9lQsCzmyLLP-0cs19MaSg4OHXbMrg1Q6BRssZnTQ9GDiuL1yh3Io1RD5iopNgir-4lOl0lfygVFZJgua89aMWJPoJkoo-SgJGqs8fMlmk1uyqC2icD1zEhfnrdU9ApALQZ0w4GNiYJYBVk15Nb3566acee92J9Mm6r8HkKR0hEpCpDK9uwDTp6e-2x9SyMLCZdOrUgOn1A`

func Test_ParsePublicKeyFromPem(t *testing.T) {
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(samplePublicKey))
	assert.NilError(t, err, "unexpected error parsing public key")
	assert.Assert(t, rsaPublicKey != nil)

	// Trying with the private key
	rsaPublicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(samplePrivateKey))
	assert.NilError(t, err, "unexpected error parsing private key")
	assert.Assert(t, rsaPublicKey != nil)
}

func TestJwtAuthenticator_parseToken(t *testing.T) {
	err := os.Setenv(RSAPublicKey, samplePublicKey)
	assert.NilError(t, err, "unexpected error setting env var")
	authenticator := new(JwtAuthenticator)
	token, claims, err := authenticator.parseToken(sampleTokenOnosGuiAndDex)
	assert.NilError(t, err, "unexpected error parsing token")
	assert.Assert(t, token != nil)
	assert.Assert(t, claims != nil)

}
