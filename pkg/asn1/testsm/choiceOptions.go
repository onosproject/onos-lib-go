// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsm

import "reflect"

// Choicemap - Should be generated with protoc-gen-choice from proto
var Choicemap = map[string]map[int]reflect.Type{
	"choice1": {
		1: reflect.TypeOf(Choice1_Choice1A{}),
	},
	"choice2": {
		1: reflect.TypeOf(Choice2_Choice2A{}),
		2: reflect.TypeOf(Choice2_Choice2B{}),
	},
	"choice3": {
		1: reflect.TypeOf(Choice3_Choice3A{}),
		2: reflect.TypeOf(Choice3_Choice3B{}),
		3: reflect.TypeOf(Choice3_Choice3C{}),
	},
	"choice4": {
		1: reflect.TypeOf(Choice4_Choice4A{}),
	},
	"constrained_choice1": {
		1: reflect.TypeOf(ConstrainedChoice1_ConstrainedChoice1A{}),
	},
	"constrained_choice2": {
		1: reflect.TypeOf(ConstrainedChoice2_ConstrainedChoice2A{}),
		2: reflect.TypeOf(ConstrainedChoice2_ConstrainedChoice2B{}),
	},
	"constrained_choice3": {
		1: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3A{}),
		2: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3B{}),
		3: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3C{}),
		4: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3D{}),
	},
	"constrained_choice4": {
		1: reflect.TypeOf(ConstrainedChoice4_ConstrainedChoice4A{}),
	},
	"test_nested_choice": {
		1: reflect.TypeOf(TestNestedChoice_Option1{}),
		2: reflect.TypeOf(TestNestedChoice_Option2{}),
		3: reflect.TypeOf(TestNestedChoice_Option3{}),
	},
	"mixed_choice": {
		1: reflect.TypeOf(MixedChoice_Ch1{}),
		2: reflect.TypeOf(MixedChoice_Ch2{}),
	},
	"choice_extended": {
		1: reflect.TypeOf(ChoiceExtended_ChoiceExtendedA{}),
		2: reflect.TypeOf(ChoiceExtended_ChoiceExtendedB{}),
		3: reflect.TypeOf(ChoiceExtended_ChoiceExtendedC{}),
		4: reflect.TypeOf(ChoiceExtended_ChoiceExtendedD{}),
	},
}

// CanonicalChoicemap - Should be generated with protoc-gen-choice from proto or created by hand (necessary to understand how to point to correct choice)
var CanonicalChoicemap = map[string]map[int64]reflect.Type{
	"canonical_nested_choice": {
		11: reflect.TypeOf(CanonicalNestedChoice_Ch1{}),
		21: reflect.TypeOf(CanonicalNestedChoice_Ch2{}),
		31: reflect.TypeOf(CanonicalNestedChoice_Ch3{}),
		41: reflect.TypeOf(CanonicalNestedChoice_Ch4{}),
	},
	"canonical_choice": {
		10: reflect.TypeOf(CanonicalChoice_Ch1{}),
		20: reflect.TypeOf(CanonicalChoice_Ch2{}),
		30: reflect.TypeOf(CanonicalChoice_Ch3{}),
		40: reflect.TypeOf(CanonicalChoice_Ch4{}),
		50: reflect.TypeOf(CanonicalChoice_Ch5{}),
		60: reflect.TypeOf(CanonicalChoice_Ch6{}),
	},
}
