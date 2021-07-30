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
}
