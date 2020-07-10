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

// generated from https://csfieldguide.org.nz/en/interactives/rsa-key-generator/
// with format scheme PKCS #8(base64)
const samplePrivateKey = `
-----BEGIN PRIVATE KEY-----
MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAi49aL5udF81/f+KI
ip/qi5CwiKimo8SXRDTG6GrKr0MCCCUUOLGMPGqnN0/L6voWA1LKjMgprtCZ3yyn
u/rqzQIDAQABAkAFYQsa1qahajxF05drsGo74uHLAqUZntQtvtMD1knlo2JPF3mJ
9JVC/edAm6TIJEsV7x5Y2L3PX00SbhVvvp9BAiEA/zgXqRNlA3jTl8KvYIpUyp6y
djN7Lywr4pKrMoXdJHkCIQCL/KqkxJo+nzCtGR8p7gQHq8AfpwYsKIka2Q6LhmBb
9QIhANIglqpoA3T2WA/NBKPRgLpKKtjSzgsqrP8gjr9MI6TRAiBiJReOxbhOx1Vd
RwuuXg29QxFEH9oYA6N8i0nDUMcmMQIgXgv52LyczFpE03TL67yWKeIb+W5EeKgO
SOE/mhFvJns=
-----END PRIVATE KEY-----
`

// generated from https://csfieldguide.org.nz/en/interactives/rsa-key-generator/
// with format scheme PKCS #8(base64)
const samplePublicKey = `
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAIuPWi+bnRfNf3/iiIqf6ouQsIiopqPE
l0Q0xuhqyq9DAgglFDixjDxqpzdPy+r6FgNSyozIKa7Qmd8sp7v66s0CAwEAAQ==
-----END PUBLIC KEY-----
`

// generated from running onos-gui against the Dex IDP - expires 10 Jul 2030
const sampleTokenOnosGuiAndDex = `eyJhbGciOiJSUzI1NiIsImtpZCI6Ijg5NjcyODQ4NzUzOWI1MmJjZGExZGU2MjljNmZkODI2MzE5M2RiODUifQ.eyJpc3MiOiJodHRwOi8vZGV4OjMyMDAwIiwic3ViIjoiQ2lRd09HRTROamcwWWkxa1lqZzRMVFJpTnpNdE9UQmhPUzB6WTJReE5qWXhaalUwTmpnU0JXeHZZMkZzIiwiYXVkIjoib25vcy1ndWkiLCJleHAiOjE5MDk3MjU2NjMsImlhdCI6MTU5NDM2NTY2Mywibm9uY2UiOiJiM0JFVkVreVJHTi1NbmRDTUMwelFVVmZhVTF2WW1KM1VuVjNURzR4Wm1kWVMyOU5hakJRWnpKeWR6ZHgiLCJhdF9oYXNoIjoiNmhQR3ppNnR4NS16QklnTEhLa3dHQSIsImVtYWlsIjoic2VhbkBvcGVubmV0d29ya2luZy5vcmciLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InNlYW4ifQ.lLZhT58X8cFxPj1h1d-cjLMt4pt_dtCqYw4rZfLqkUNnnatOajIQv7Dg93QKJizz9HGbFSvz8rUEvkQxhGjNQjueVva97_fig6MTKP7vj2TL5L5XO0XtI39JZlGJYA3kwE1Xw43diMnDRjaU6UHmcMbjUA-aF97WDLWgiMbu-wnbfdfEj_9pT0vytFLPFVlww6EzjvNTfwneUVLFOqU0Hq9ykv6eDugKhilYWBhpWhC-hOTlJNfVn8IU3gIU2whl_YU6--4BpBJli3UKRqbrAzwnkE8-OZA6TVT3uqysuZ93_Sgm_EgXXrNp_yX8nLreTgKZtpP6e5ROK9oRH89aVg`

func Test_ParsePublicKeyFromPem(t *testing.T) {
	assert.Equal(t, 523, len(samplePrivateKey))
	assert.Equal(t, 183, len(samplePublicKey))

	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(samplePublicKey))
	assert.NilError(t, err, "unexpected error parsing public key")
	assert.Assert(t, rsaPublicKey != nil)

}

func TestJwtAuthenticator_parseToken(t *testing.T) {
	t.Skip()
	err := os.Setenv(RSAPublicKey, samplePublicKey)
	assert.NilError(t, err, "unexpected error setting env var")
	authenticator := new(JwtAuthenticator)
	token, claims, err := authenticator.parseToken(sampleTokenOnosGuiAndDex)
	assert.NilError(t, err, "unexpected error parsing token")
	assert.Assert(t, token != nil)
	assert.Assert(t, claims != nil)

}
