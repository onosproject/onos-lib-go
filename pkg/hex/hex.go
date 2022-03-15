// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
