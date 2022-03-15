// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package test

import "reflect"

// Choicemap - Should be generated with protoc-gen-choice from proto
var Choicemap = map[string]map[int]reflect.Type{
	"Choice1": {
		1: reflect.TypeOf(TestChoices_Choice1A{}),
	},
	"Choice2": {
		1: reflect.TypeOf(TestChoices_Choice2A{}),
		2: reflect.TypeOf(TestChoices_Choice2B{}),
	},
	"Choice3": {
		1: reflect.TypeOf(TestChoices_Choice3A{}),
		2: reflect.TypeOf(TestChoices_Choice3B{}),
		3: reflect.TypeOf(TestChoices_Choice3C{}),
	},
}
