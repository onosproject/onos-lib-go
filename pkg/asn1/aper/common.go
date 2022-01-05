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

package aper

import (
	"strconv"
	"strings"
)

// fieldParameters is the parsed representation of tag string from a structure field.
type fieldParameters struct {
	optional            bool   // true iff the type has OPTIONAL tag.
	sizeExtensible      bool   // true iff the size can be extensed.
	valueExtensible     bool   // true iff the value can be extensed.
	sizeLowerBound      *int64 // a sizeLowerBound is the minimum size of type constraint(maybe nil).
	sizeUpperBound      *int64 // a sizeUpperBound is the maximum size of type constraint(maybe nil).
	valueLowerBound     *int64 // a valueLowerBound is the minimum value of type constraint(maybe nil).
	valueUpperBound     *int64 // a valueUpperBound is the maximum value of type constraint(maybe nil).
	defaultValue        *int64 // a default value for INTEGER and ENUMERATED typed fields (maybe nil).
	openType            bool   // true iff this type is opentype.
	referenceFieldName  string // the field to get to get the corresrponding value of this type(maybe nil).
	referenceFieldValue *int64 // the field value which map to this type(maybe nil).
	choiceIndex         *uint8 // the choice index
	oneofName           string // a oneof name if present
	align               bool   // flag to explicitly align structure
	unique              bool   // true if item is used as an input for canonical ordering for CHOICE encoding
	canonicalOrder      bool   // true, if CHOICE is needed to be encoded in canonical ordering
	fromChoiceExt       bool   // true, if CHOICE item belongs to CHOICE extension
	choiceExt           bool   // true, if CHOICE can be extended with other items
}

// Given a tag string with the format specified in the package comment,
// parseFieldParameters will parse it into a fieldParameters structure,
// ignoring unknown parts of the string. TODO:PrintableString
func parseFieldParameters(str string) (params fieldParameters) {
	for _, part := range strings.Split(str, ",") {
		switch {
		case part == "choiceExt":
			params.choiceExt = true
		case part == "fromChoiceExt":
			params.fromChoiceExt = true
		case part == "align":
			params.align = true
		case part == "unique":
			params.unique = true
		case part == "canonicalOrder":
			params.canonicalOrder = true
		case part == "optional":
			params.optional = true
		case part == "sizeExt":
			params.sizeExtensible = true
		case part == "valueExt":
			params.valueExtensible = true
		case strings.HasPrefix(part, "sizeLB:"):
			i, err := strconv.ParseInt(part[7:], 10, 64)
			if err == nil {
				params.sizeLowerBound = new(int64)
				*params.sizeLowerBound = i
			}
		case strings.HasPrefix(part, "sizeUB:"):
			i, err := strconv.ParseInt(part[7:], 10, 64)
			if err == nil {
				params.sizeUpperBound = new(int64)
				*params.sizeUpperBound = i
			}
		case strings.HasPrefix(part, "valueLB:"):
			i, err := strconv.ParseInt(part[8:], 10, 64)
			if err == nil {
				params.valueLowerBound = new(int64)
				*params.valueLowerBound = i
			}
		case strings.HasPrefix(part, "valueUB:"):
			i, err := strconv.ParseInt(part[8:], 10, 64)
			if err == nil {
				params.valueUpperBound = new(int64)
				*params.valueUpperBound = i
			}
		case strings.HasPrefix(part, "default:"):
			i, err := strconv.ParseInt(part[8:], 10, 64)
			if err == nil {
				params.defaultValue = new(int64)
				*params.defaultValue = i
			}
		case part == "openType":
			params.openType = true
		case strings.HasPrefix(part, "referenceFieldName:"):
			params.referenceFieldName = part[19:]
		case strings.HasPrefix(part, "referenceFieldValue:"):
			i, err := strconv.ParseInt(part[20:], 10, 64)
			if err == nil {
				params.referenceFieldValue = new(int64)
				*params.referenceFieldValue = i
			}
		case strings.HasPrefix(part, "choiceIdx:"):
			i, err := strconv.ParseInt(part[10:], 10, 8)
			if err == nil {
				params.choiceIndex = new(uint8)
				*params.choiceIndex = uint8(i)
			}
		}
	}
	return params
}
