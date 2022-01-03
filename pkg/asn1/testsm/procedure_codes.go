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
