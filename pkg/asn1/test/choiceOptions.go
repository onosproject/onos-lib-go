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
