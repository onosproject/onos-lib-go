// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

// Driven from e2ap_constants.proto
// TODO: Automate the generation of this file

// CanonicalChoiceID - should be generated with protoc-gen-cgo.
// It is a necessary prerequisite for canonical ordering in CHOICE encoding
type CanonicalChoiceID int32

const (
	CanonicalChoiceIDSampleOctetString         CanonicalChoiceID = 10
	CanonicalChoiceIDSampleConstrainedInteger  CanonicalChoiceID = 20
	CanonicalChoiceIDSampleBitString           CanonicalChoiceID = 30
	CanonicalChoiceIDTestListExtensible1       CanonicalChoiceID = 40
	CanonicalChoiceIDItem                      CanonicalChoiceID = 50
	CanonicalChoiceIDSampleNestedE2apPduChoice CanonicalChoiceID = 60
)

// CanonicalNestedChoiceID - should be generated with protoc-gen-cgo.
// It is a necessary prerequisite for canonical ordering in CHOICE encoding
type CanonicalNestedChoiceID int32

const (
	CanonicalNestedChoiceIDSampleOctetString        CanonicalNestedChoiceID = 11
	CanonicalNestedChoiceIDSampleConstrainedInteger CanonicalNestedChoiceID = 21
	CanonicalNestedChoiceIDSampleBitString          CanonicalNestedChoiceID = 31
	CanonicalNestedChoiceIDTestListExtensible1      CanonicalNestedChoiceID = 41
)
