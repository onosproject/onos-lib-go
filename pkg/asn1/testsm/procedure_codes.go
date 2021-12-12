// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package testsm

// Driven from e2ap_constants.proto
// TODO: Automate the generation of this file

type CanonicalChoiceID int32

const (
	CanonicalChoiceIDSampleOctetString         CanonicalChoiceID = 1
	CanonicalChoiceIDSampleConstrainedInteger  CanonicalChoiceID = 2
	CanonicalChoiceIDSampleBitString           CanonicalChoiceID = 3
	CanonicalChoiceIDTestListExtensible1       CanonicalChoiceID = 4
	CanonicalChoiceIDItem                      CanonicalChoiceID = 5
	CanonicalChoiceIDSampleNestedE2apPduChoice CanonicalChoiceID = 6
)

type CanonicalNestedChoiceID int32

const (
	CanonicalNestedChoiceIDSampleOctetString        CanonicalNestedChoiceID = 1
	CanonicalNestedChoiceIDSampleConstrainedInteger CanonicalNestedChoiceID = 2
	CanonicalNestedChoiceIDSampleBitString          CanonicalNestedChoiceID = 3
	CanonicalNestedChoiceIDTestListExtensible1      CanonicalNestedChoiceID = 4
)
