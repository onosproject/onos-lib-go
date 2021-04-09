// Copyright 2021-present Open Networking Foundation.
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

package hex

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// Asn1BytesToByte This function was created for debugging purposes
// It takes as an input string of asn1 bytes (in HEXadecimal format)
// and converts them to the array of bytes ([]byte) which could be passed
// to the input of *ASN1toProto function
// Input data could be obtained in a following way:
// t.Logf("E2SM-KPM-ActionDefinition (Format3) asn1Bytes are \n%x", asn1Bytes)
func Asn1BytesToByte(str string) ([]byte, error) {

	return hex.DecodeString(str)
}

// DumpToByte This function was created for debugging purposes.
// It takes as an input output of hex.Dump() for asn1 bytes
// and converts them to the array of bytes ([]byte)
// which could be passed to the input of *ASN1toProto function
// Input data could be obtained in a following way:
// t.Logf("E2SM-KPM-ActionDefinition (Format3) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
func DumpToByte(str string) ([]byte, error) {

	r, err := regexp.Compile("([\t\n\f\r ][0-9a-f]{2})")
	if err != nil {
		return nil, err
	}
	out := r.FindAllString(str, -1)

	res := ""
	escapeElement := 16
	for i, slice := range out {
		postprcs := strings.ReplaceAll(slice, " ", "")
		postprcss := strings.ReplaceAll(postprcs, "  ", "")
		if i != escapeElement {
			res = res + fmt.Sprintf("%v", strings.ReplaceAll(postprcss, "\n", ""))
		} else {
			escapeElement = escapeElement + 17
		}
	}

	b, err := hex.DecodeString(res)
	if err != nil {
		return nil, err
	}

	return b, nil
}
