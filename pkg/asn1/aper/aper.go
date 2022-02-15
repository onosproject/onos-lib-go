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
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"reflect"
)

var log = logging.GetLogger("asn1", "aper")

// ChoiceMap - a global map of choices - specific to the Protobuf being handled
//var ChoiceMap = map[string]map[int]reflect.Type{}

// CanonicalChoiceMap - a global map of choices in canonical ordering - specific to the Protobuf being handled
//var CanonicalChoiceMap = map[string]map[int64]reflect.Type{}
//
//var canonicalOrdering = false
//var choiceCanBeExtended = false

type perBitData struct {
	bytes               []byte
	byteOffset          uint64
	bitsOffset          uint
	choiceMap           map[string]map[int]reflect.Type
	unique              int64
	canonicalOrdering   bool
	canonicalChoiceMap  map[string]map[int64]reflect.Type
	choiceCanBeExtended bool
}

func perBitLog(numBits uint64, byteOffset uint64, bitsOffset uint, value interface{}) string {
	if reflect.TypeOf(value).Kind() == reflect.Uint64 {
		return fmt.Sprintf("  [PER got %2d bits, byteOffset(after): %d, bitsOffset(after): %d, value: 0x%0x]",
			numBits, byteOffset, bitsOffset, reflect.ValueOf(value).Uint())
	}
	return fmt.Sprintf("  [PER got %2d bits, byteOffset(after): %d, bitsOffset(after): %d, value: 0x%0x]",
		numBits, byteOffset, bitsOffset, reflect.ValueOf(value).Bytes())

}

// GetBitString is to get BitString with desire size from source byte array with bit offset
func GetBitString(srcBytes []byte, bitsOffset uint, numBits uint) (dstBytes []byte, err error) {
	bitsLeft := uint(len(srcBytes))*8 - bitsOffset
	if numBits > bitsLeft {
		err = fmt.Errorf("Get bits overflow, requireBits: %d, leftBits: %d", numBits, bitsLeft)
		return
	}
	byteLen := (bitsOffset + numBits + 7) >> 3
	numBitsByteLen := (numBits + 7) >> 3
	dstBytes = make([]byte, numBitsByteLen)
	numBitsMask := byte(0xff)
	if modEight := numBits & 0x7; modEight != 0 {
		numBitsMask <<= uint8(8 - (modEight))
	}
	for i := 1; i < int(byteLen); i++ {
		dstBytes[i-1] = srcBytes[i-1]<<bitsOffset | srcBytes[i]>>(8-bitsOffset)
	}
	if byteLen == numBitsByteLen {
		dstBytes[byteLen-1] = srcBytes[byteLen-1] << bitsOffset
	}
	dstBytes[numBitsByteLen-1] &= numBitsMask
	return
}

// GetFewBits is to get Value with desire few bits from source byte with bit offset
// func GetFewBits(srcByte byte, bitsOffset uint, numBits uint) (value uint64, err error) {

// 	if numBits == 0 {
// 		value = 0
// 		return
// 	}
// 	bitsLeft := 8 - bitsOffset
// 	if bitsLeft < numBits {
// 		err = fmt.Errorf("Get bits overflow, requireBits: %d, leftBits: %d", numBits, bitsLeft)
// 		return
// 	}
// 	if bitsOffset == 0 {
// 		value = uint64(srcByte >> (8 - numBits))
// 	} else {
// 		value = uint64((srcByte << bitsOffset) >> (8 - numBits))
// 	}
// 	return
// }

// GetBitsValue is to get Value with desire bits from source byte array with bit offset
func GetBitsValue(srcBytes []byte, bitsOffset uint, numBits uint) (value uint64, err error) {
	var dstBytes []byte
	dstBytes, err = GetBitString(srcBytes, bitsOffset, numBits)
	if err != nil {
		return
	}
	for i, j := 0, numBits; j >= 8; i, j = i+1, j-8 {
		value <<= 8
		value |= uint64(uint(dstBytes[i]))
	}
	if numBitsOff := numBits & 0x7; numBitsOff != 0 {
		var mask uint = (1 << numBitsOff) - 1
		value <<= numBitsOff
		value |= uint64(uint(dstBytes[len(dstBytes)-1]>>(8-numBitsOff)) & mask)
	}
	return
}

func (pd *perBitData) bitCarry() {
	pd.byteOffset += uint64(pd.bitsOffset >> 3)
	pd.bitsOffset = pd.bitsOffset & 0x07
}

func (pd *perBitData) getBitString(numBits uint) (dstBytes []byte, err error) {

	dstBytes, err = GetBitString(pd.bytes[pd.byteOffset:], pd.bitsOffset, numBits)
	if err != nil {
		return
	}
	pd.bitsOffset += numBits

	pd.bitCarry()
	log.Debugf("%s", perBitLog(uint64(numBits), pd.byteOffset, pd.bitsOffset, dstBytes))
	return
}

func (pd *perBitData) getBitsValue(numBits uint) (value uint64, err error) {
	value, err = GetBitsValue(pd.bytes[pd.byteOffset:], pd.bitsOffset, numBits)
	if err != nil {
		return
	}
	pd.bitsOffset += numBits
	pd.bitCarry()
	log.Debugf("%s", perBitLog(uint64(numBits), pd.byteOffset, pd.bitsOffset, value))
	return
}

func (pd *perBitData) parseAlignBits() error {

	if (pd.bitsOffset & 0x7) > 0 {
		alignBits := 8 - ((pd.bitsOffset) & 0x7)
		log.Debugf("Aligning %d bits", alignBits)
		if val, err := pd.getBitsValue(alignBits); err != nil {
			return err
		} else if val != 0 {
			return fmt.Errorf("Align Bit is not zero in %v", hex.Dump(pd.bytes[pd.byteOffset:pd.byteOffset+1]))
		}
	} else if pd.bitsOffset != 0 {
		pd.bitCarry()
	}
	return nil
}

func (pd *perBitData) parseConstraintValue(valueRange int64) (value uint64, err error) {
	log.Debugf("Getting Constraint Value with range %d", valueRange)

	var bytes uint
	if valueRange <= 255 {
		if valueRange < 0 {
			err = fmt.Errorf("Value range is negative")
			return
		}
		var i uint
		// 1 ~ 8 bits
		for i = 1; i <= 8; i++ {
			upper := 1 << i
			if int64(upper) >= valueRange {
				break
			}
		}
		value, err = pd.getBitsValue(i)
		return
	} else if valueRange == 256 {
		bytes = 1
	} else if valueRange <= 65536 {
		bytes = 2
	} else {
		err = fmt.Errorf("Constraint Value is large than 65536")
		return
	}
	if err = pd.parseAlignBits(); err != nil {
		return
	}
	value, err = pd.getBitsValue(bytes * 8)
	return value, err
}

func (pd *perBitData) parseLength(sizeRange int64, repeat *bool) (value uint64, err error) {
	*repeat = false
	if sizeRange <= 65536 && sizeRange > 0 {
		return pd.parseConstraintValue(sizeRange)
	}

	if err = pd.parseAlignBits(); err != nil {
		return
	}
	firstByte, err := pd.getBitsValue(8)
	if err != nil {
		return
	}
	if (firstByte & 128) == 0 { // #10.9.3.6
		value = firstByte & 0x7F
		return
	} else if (firstByte & 64) == 0 { // #10.9.3.7
		var secondByte uint64
		if secondByte, err = pd.getBitsValue(8); err != nil {
			return
		}
		value = ((firstByte & 63) << 8) | secondByte
		return
	}
	firstByte &= 63
	if firstByte < 1 || firstByte > 4 {
		err = fmt.Errorf("Parse Length Out of Constraint")
		return
	}
	*repeat = true
	value = 16384 * firstByte
	return value, err
}

func (pd *perBitData) parseBitString(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (*asn1.BitString, error) {
	var lb, ub, sizeRange int64 = 0, -1, -1
	if !extensed {
		if lowerBoundPtr != nil {
			lb = *lowerBoundPtr
		}
		if upperBoundPtr != nil {
			ub = *upperBoundPtr
			sizeRange = ub - lb + 1
		}
	}
	if ub > 65535 {
		sizeRange = -1
	}
	// initailization
	bitString := asn1.BitString{Value: make([]byte, 0), Len: 0}
	// lowerbound == upperbound
	if sizeRange == 1 {
		sizes := uint64(ub+7) >> 3
		bitString.Len = uint32(uint64(ub))
		log.Debugf("Decoding BIT STRING size %d", ub)
		log.Debugf("Decoding BIT STRING size %d", bitString.Len)
		if sizes > 2 {
			if err := pd.parseAlignBits(); err != nil {
				return nil, err
			}
			if (pd.byteOffset + sizes) > uint64(len(pd.bytes)) {
				err := fmt.Errorf("PER data out of range")
				return nil, err
			}
			if _, err := bitString.UpdateValue(pd.bytes[pd.byteOffset : pd.byteOffset+sizes]); err != nil {
				return nil, err
			}
			// Truncating last trailing bits -- Length of a BitString should be already set!
			if _, err := bitString.TruncateValue(); err != nil {
				return nil, err
			}
			//bitString.Value = append(bitString.Value, pd.bytes[pd.byteOffset:pd.byteOffset+sizes]...)
			pd.byteOffset += sizes
			pd.bitsOffset = uint(ub & 0x7)
			if pd.bitsOffset > 0 {
				pd.byteOffset--
			}
			log.Debugf("%s", perBitLog(uint64(ub), pd.byteOffset, pd.bitsOffset, bitString.Value))
		} else {
			bytes, err := pd.getBitString(uint(ub))
			if err != nil {
				log.Errorf("PD GetBitString error: %+v", err)
				return nil, err
			}
			if _, err = bitString.UpdateValue(bytes); err != nil {
				return nil, err
			}
		}
		log.Debugf("Decoded BIT STRING (length = %d): %0.8b", ub, bitString.Value)
		return &bitString, nil

	}
	repeat := false
	for {
		var rawLength uint64
		length, err := pd.parseLength(sizeRange, &repeat)
		if err != nil {
			return nil, err
		}
		rawLength = length
		rawLength += uint64(lb)
		log.Debugf("Decoding BIT STRING size %d", rawLength)
		if rawLength == 0 {
			return nil, nil
		}
		sizes := (rawLength + 7) >> 3
		if err := pd.parseAlignBits(); err != nil {
			return nil, err
		}

		if (pd.byteOffset + sizes) > uint64(len(pd.bytes)) {
			return nil, errors.NewInvalid("PER data out of range")
		}
		//bitString.Value = append(bitString.Value, pd.bytes[pd.byteOffset:pd.byteOffset+sizes]...)
		bitString.Len += uint32(rawLength)
		// we need to get length before we want to decode with UpdateValue
		if _, err := bitString.UpdateValue(pd.bytes[pd.byteOffset : pd.byteOffset+sizes]); err != nil {
			return nil, err
		}
		// Truncating last trailing bits -- Length of a BitString should be already set!
		if _, err := bitString.TruncateValue(); err != nil {
			return nil, err
		}
		pd.byteOffset += sizes
		pd.bitsOffset = uint(rawLength & 0x7)
		if pd.bitsOffset != 0 {
			pd.byteOffset--
		}
		log.Debugf("%s", perBitLog(rawLength, pd.byteOffset, pd.bitsOffset, bitString.Value))
		log.Debugf("Decoded BIT STRING (length = %d): %0.8b", rawLength, bitString.Value)

		if !repeat {
			// if err = pd.parseAlignBits(); err != nil {
			// 	return
			// }
			break
		}
	}
	return &bitString, nil
}
func (pd *perBitData) parseOctetString(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (
	[]byte, error) {
	var lb, ub, sizeRange int64 = 0, -1, -1
	if !extensed {
		if lowerBoundPtr != nil {
			lb = *lowerBoundPtr
		}
		if upperBoundPtr != nil {
			ub = *upperBoundPtr
			sizeRange = ub - lb + 1
		}
	}
	if ub > 65535 {
		sizeRange = -1
	}
	// initailization
	octetString := []byte("")
	// lowerbound == upperbound
	if sizeRange == 1 {
		log.Debugf("Decoding OCTET STRING size %d", ub)
		if ub > 2 {
			unsignedUB := uint64(ub)
			if err := pd.parseAlignBits(); err != nil {
				return octetString, err
			}
			if (int64(pd.byteOffset) + ub) > int64(len(pd.bytes)) {
				err := fmt.Errorf("per data out of range")
				return octetString, err
			}
			octetString = pd.bytes[pd.byteOffset : pd.byteOffset+unsignedUB]
			pd.byteOffset += uint64(ub)
			log.Debugf("%s", perBitLog(8*unsignedUB, pd.byteOffset, pd.bitsOffset, octetString))
		} else {
			octet, err := pd.getBitString(uint(ub * 8))
			if err != nil {
				return octetString, err
			}
			octetString = octet
		}
		log.Debugf("Decoded OCTET STRING (length = %d): 0x%0x", ub, octetString)
		return octetString, nil

	}
	repeat := false
	for {
		var rawLength uint64
		length, err := pd.parseLength(sizeRange, &repeat)
		if err != nil {
			return octetString, err
		}
		rawLength = length
		rawLength += uint64(lb)
		log.Debugf("Decoding OCTET STRING size %d", rawLength)
		if rawLength == 0 {
			return octetString, nil
		} else if err := pd.parseAlignBits(); err != nil {
			return octetString, err
		}
		if (rawLength + pd.byteOffset) > uint64(len(pd.bytes)) {
			err := fmt.Errorf("per data out of range ")
			return octetString, err
		}
		octetString = append(octetString, pd.bytes[pd.byteOffset:pd.byteOffset+rawLength]...)
		pd.byteOffset += rawLength
		log.Debugf("%s", perBitLog(8*rawLength, pd.byteOffset, pd.bitsOffset, octetString))
		log.Debugf("Decoded OCTET STRING (length = %d): 0x%0x", rawLength, octetString)
		if !repeat {
			// if err = pd.parseAlignBits(); err != nil {
			// 	return
			// }
			break
		}
	}
	return octetString, nil
}

func (pd *perBitData) parseBool() (value bool, err error) {
	log.Debugf("Decoding BOOLEAN Value")
	bit, err1 := pd.getBitsValue(1)
	if err1 != nil {
		err = err1
		return
	}
	if bit == 1 {
		value = true
	} else {
		value = false
	}
	log.Debugf("Decoded BOOLEAN Value %v", value)
	return
}

func (pd *perBitData) parseInteger(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (int64, error) {
	var lb, ub, valueRange int64 = 0, -1, 0
	if !extensed {
		if lowerBoundPtr == nil {
			log.Debugf("Decoding INTEGER with Unconstraint Value")
			valueRange = -1
		} else {
			lb = *lowerBoundPtr
			if upperBoundPtr != nil {
				ub = *upperBoundPtr
				valueRange = ub - lb + 1
				log.Debugf("Decoding INTEGER with Value Range(%d..%d)", lb, ub)
			} else {
				log.Debugf("Decoding INTEGER with Semi-Constraint Range(%d..)", lb)
			}
		}
	} else {
		valueRange = -1
		log.Debugf("Decoding INTEGER with Extensive Value")
	}
	var rawLength uint
	if valueRange == 1 {
		return ub, nil
	} else if valueRange <= 0 {
		// semi-constraint or unconstraint
		if err := pd.parseAlignBits(); err != nil {
			return int64(0), err
		}
		if pd.byteOffset >= uint64(len(pd.bytes)) {
			return int64(0), fmt.Errorf("per data out of range")
		}
		rawLength = uint(pd.bytes[pd.byteOffset])
		pd.byteOffset++
		log.Debugf("%s", perBitLog(8, pd.byteOffset, pd.bitsOffset, uint64(rawLength)))
	} else if valueRange <= 65536 {
		rawValue, err := pd.parseConstraintValue(valueRange)
		if err != nil {
			return int64(0), err
		}
		return int64(rawValue) + lb, nil
	} else {
		// valueRange > 65536
		var byteLen uint
		unsignedValueRange := uint64(valueRange - 1)
		for byteLen = 1; byteLen <= 127; byteLen++ {
			unsignedValueRange >>= 8
			if unsignedValueRange == 0 {
				break
			}
		}
		var i, upper uint
		// 1 ~ 8 bits
		for i = 1; i <= 8; i++ {
			upper = 1 << i
			if upper >= byteLen {
				break
			}
		}
		tempLength, err := pd.getBitsValue(i)
		if err != nil {
			return int64(0), err
		}
		rawLength = uint(tempLength)
		rawLength++
		if err := pd.parseAlignBits(); err != nil {
			return int64(0), err
		}
	}
	log.Debugf("Decoding INTEGER Length with %d bytes", rawLength)

	if rawValue, err := pd.getBitsValue(rawLength * 8); err != nil {
		return int64(0), err
	} else if valueRange < 0 {
		signedBitMask := uint64(1 << (rawLength*8 - 1))
		valueMask := signedBitMask - 1
		// negative
		if rawValue&signedBitMask > 0 {
			return int64((^rawValue)&valueMask+1) * -1, nil
		}
		return int64(rawValue) + lb, nil
	} else {
		return int64(rawValue) + lb, nil
	}
}

// parse ENUMERATED type but do not implement extensive value and different value with index
//func (pd *perBitData) parseEnumerated(extensed bool, lowerBoundPtr *int64, upperBoundPtr *int64) (value uint64,
//	err error) {
//	if extensed {
//		err = fmt.Errorf("Unsupport the extensive value of ENUMERATED ")
//		return
//	}
//	if lowerBoundPtr == nil || upperBoundPtr == nil {
//		err = fmt.Errorf("ENUMERATED value constraint is error ")
//		return
//	}
//	lb, ub := *lowerBoundPtr, *upperBoundPtr
//	valueRange := ub - lb + 1
//	log.Debugf("Decoding ENUMERATED with Value Range(%d..%d)", lb, ub)
//	if valueRange > 1 {
//		value, err = pd.parseConstraintValue(valueRange)
//	}
//	log.Debugf("Decoded ENUMERATED Value : %d", value)
//	return
//
//}

func (pd *perBitData) parseSequenceOf(sizeExtensed bool, params fieldParameters, sliceType reflect.Type) (
	reflect.Value, error) {
	var sliceContent reflect.Value
	var lb int64
	var sizeRange int64
	if params.sizeLowerBound != nil && *params.sizeLowerBound < 65536 {
		lb = *params.sizeLowerBound
	}
	if !sizeExtensed && params.sizeUpperBound != nil && *params.sizeUpperBound < 65536 {
		ub := *params.sizeUpperBound
		sizeRange = ub - lb + 1
		log.Debugf("Decoding Length of \"SEQUENCE OF\"  with Size Range(%d..%d)", lb, ub)
	} else {
		sizeRange = -1
		log.Debugf("Decoding Length of \"SEQUENCE OF\" with Semi-Constraint Range(%d..)", lb)
	}

	var numElements uint64
	if sizeRange > 1 {
		if numElementsTmp, err := pd.parseConstraintValue(sizeRange); err != nil {
			log.Errorf("Parse Constraint Value failed: %+v", err)
		} else {
			numElements = numElementsTmp
		}
		numElements += uint64(lb)
	} else if sizeRange == 1 {
		numElements += uint64(lb)
	} else {
		if err := pd.parseAlignBits(); err != nil {
			return sliceContent, err
		}
		if pd.byteOffset >= uint64(len(pd.bytes)) {
			err := fmt.Errorf("per data out of range")
			return sliceContent, err
		}
		numElements = uint64(pd.bytes[pd.byteOffset])
		pd.byteOffset++
		log.Debugf("%s", perBitLog(8, pd.byteOffset, pd.bitsOffset, numElements))
	}
	log.Debugf("Decoding  \"SEQUENCE OF\" struct %s with len(%d)", sliceType.Elem().Name(), numElements)
	params.sizeExtensible = false
	params.sizeUpperBound = nil
	params.sizeLowerBound = nil
	intNumElements := int(numElements)
	sliceContent = reflect.MakeSlice(sliceType, intNumElements, intNumElements)
	for i := 0; i < intNumElements; i++ {
		err := parseField(sliceContent.Index(i), pd, params)
		if err != nil {
			return sliceContent, err
		}
	}
	return sliceContent, nil
}

func (pd *perBitData) getChoiceIndex(extensed bool, fromChoiceExtension bool, numItemsNotInExtension int, choiceMapLen int) (present int, err error) {

	if pd.choiceCanBeExtended {
		// This flag has already served for its purpose. Setting it back to its initial value
		pd.choiceCanBeExtended = false

		isExtended := false
		if bitsValue, err1 := pd.getBitsValue(1); err1 != nil {
			err = err1
		} else if bitsValue != 0 {
			isExtended = true
		}

		if isExtended {
			log.Debugf("Choice is extended. Parsing items from extension")
			upperBound := choiceMapLen - numItemsNotInExtension
			if upperBound == 1 {
				present = upperBound + 1
			} else {
				if upperBound < 1 {
					err = fmt.Errorf("the upper bound of CHOICE is missing")
				} else if rawChoice, err1 := pd.parseConstraintValue(int64(upperBound)); err1 != nil {
					err = err1
				} else {
					present = int(rawChoice) + 1 + numItemsNotInExtension
				}
			}
		} else {
			log.Debugf("Choice is not extended. Parsing main items")
			if numItemsNotInExtension == 1 {
				present = 1
			} else {
				upperBound := numItemsNotInExtension
				if upperBound < 1 {
					err = fmt.Errorf("the upper bound of CHOICE is missing")
				} else if rawChoice, err1 := pd.parseConstraintValue(int64(numItemsNotInExtension)); err1 != nil {
					err = err1
				} else {
					present = int(rawChoice) + 1
				}
			}
		}
	} else {
		upperBound := choiceMapLen
		if upperBound < 1 {
			err = fmt.Errorf("the upper bound of CHOICE is missing")
		} else if rawChoice, err1 := pd.parseConstraintValue(int64(upperBound)); err1 != nil {
			err = err1
		} else {
			log.Debugf("Decoded Present index of CHOICE is %d + 1", rawChoice)
			present = int(rawChoice) + 1
		}
	}

	return present, err
}

func (pd *perBitData) getCanonicalChoiceIndex() error {

	err := pd.parseAlignBits()
	if err != nil {
		return err
	}
	log.Debugf("Parsing %v bytes", len(pd.bytes[pd.byteOffset:]))

	ext, err := pd.getBitsValue(1)
	if err != nil {
		return err
	}

	if ext == 0 {
		numBytes, err := pd.getBitsValue(7)
		if err != nil {
			return err
		}
		//ToDo - valid only when Canonical CHOICE is the last part of the message
		//if numBytes != uint64(len(pd.bytes[pd.byteOffset:])) {
		//	return errors.NewInvalid("Checksum didn't pass. Expecting %v bytes, but have %v bytes to decode", numBytes, len(pd.bytes[pd.byteOffset:]))
		//}
		log.Debugf("Decoding %v bytes", numBytes)
	} else if ext == 1 {
		numBytes, err := pd.getBitsValue(15)
		if err != nil {
			return err
		}
		//ToDo - valid only when Canonical CHOICE is the last part of the message
		//if numBytes != uint64(len(pd.bytes[pd.byteOffset:])) {
		//	return errors.NewInvalid("Checksum didn't pass. Expecting %v bytes, but have %v bytes to decode", numBytes, len(pd.bytes[pd.byteOffset:]))
		//}
		log.Debugf("Decoding %v bytes", numBytes)
	}

	return nil
}

//func getReferenceFieldValue(v reflect.Value) (value int64, err error) {
//	fieldType := v.Type()
//	switch v.Kind() {
//	case reflect.Int, reflect.Int32, reflect.Int64:
//		value = v.Int()
//	case reflect.Struct:
//		if fieldType.Field(0).Name == "Present" {
//			present := int(v.Field(0).Int())
//			if present == 0 {
//				err = fmt.Errorf("ReferenceField Value present is 0(present's field number)")
//			} else if present >= fieldType.NumField() {
//				err = fmt.Errorf("'Present' is bigger than number of struct field")
//			} else {
//				value, err = getReferenceFieldValue(v.Field(present))
//			}
//		} else {
//			value, err = getReferenceFieldValue(v.Field(0))
//		}
//	default:
//		err = fmt.Errorf("OpenType reference only support INTEGER")
//	}
//	return
//}

// parseField is the main parsing function. Given a byte slice and an offset
// into the array, it will try to parse a suitable ASN.1 value out and store it
// in the given Value. TODO : ObjectIdenfier, handle extension Field
func parseField(v reflect.Value, pd *perBitData, params fieldParameters) error {
	log.Debugf("Decoding %s", v.Type().Name())
	fieldType := v.Type()

	// If we have run out of data return error.
	if pd.byteOffset == uint64(len(pd.bytes)) {
		return fmt.Errorf("sequence truncated")
	}
	if v.Kind() == reflect.Ptr {
		ptr := reflect.New(fieldType.Elem())
		v.Set(ptr)
		return parseField(v.Elem(), pd, params)
	}
	sizeExtensible := false
	valueExtensible := false
	if params.sizeExtensible {
		if bitsValue, err1 := pd.getBitsValue(1); err1 != nil {
			return err1
		} else if bitsValue != 0 {
			sizeExtensible = true
		}
		log.Debugf("Decoded Size Extensive Bit : %t", sizeExtensible)
	}
	if params.valueExtensible && v.Kind() != reflect.Slice && !params.choiceExt {
		if bitsValue, err1 := pd.getBitsValue(1); err1 != nil {
			return err1
		} else if bitsValue != 0 {
			valueExtensible = true
		}
		log.Debugf("Decoded Value Extensive Bit : %t", valueExtensible)
	}
	if params.choiceExt && v.Kind() != reflect.Slice {
		// We have to make this variable global. In the decoding we§re parsing parent structure first
		// and then drilling down to its child.  Once we've drilled down, we don't see previous (local) flag anymore.
		pd.choiceCanBeExtended = true
		log.Debugf("CHOICE can be extended")
	}

	// Setting explicitly flag for canonical ordering here due to the specificity of passing flags to choices
	// (canonicalOrder flag is being set one level above than for regular choice)
	if params.canonicalOrder {
		pd.canonicalOrdering = true
		log.Debugf("Setting canonicalOrdering flag to true. Next CHOICE is expected to be in canonical ordering")
	}

	// We deal with the structures defined in this package first.
	switch fieldType {
	case BitStringType:
		bitString, err1 := pd.parseBitString(sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		if err1 != nil {
			return err1
		}
		v.Field(3).Set(reflect.ValueOf(bitString.Value))
		v.Field(4).Set(reflect.ValueOf(bitString.Len))
		return nil
	case reflect.TypeOf([]uint8{}):
		octetString, err := pd.parseOctetString(sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(octetString))
		return nil
	default:
		log.Debugf("not a built in field type %v", fieldType)
	}
	switch val := v; val.Kind() {
	case reflect.Bool:
		parsedBool, err := pd.parseBool()
		if err != nil {
			return err
		}
		val.SetBool(parsedBool)
		return nil
	case reflect.Int, reflect.Int32, reflect.Int64:
		parsedInt, err := pd.parseInteger(valueExtensible, params.valueLowerBound, params.valueUpperBound)
		if err != nil {
			return err
		}
		val.SetInt(parsedInt)
		log.Debugf("Decoded INTEGER Value: %d", parsedInt)
		if params.unique {
			pd.unique = parsedInt
			log.Debugf("UNIQUE flag was found, it is %v", pd.unique)
		}
		return nil
	case reflect.Struct:
		structType := fieldType
		var structParams []fieldParameters
		var optionalCount uint
		var optionalPresents uint64

		// pass tag for optional
		fieldIdx := -1
		for i := 0; i < structType.NumField(); i++ {
			if structType.Field(i).PkgPath != "" {
				log.Debugf("struct %s ignoring unexported field : %s", structType.Name(), structType.Field(i).Name)
				continue
			}
			fieldIdx++
			tempParams := parseFieldParameters(structType.Field(i).Tag.Get("aper"))
			tempParams.oneofName = structType.Field(i).Tag.Get("protobuf_oneof")
			// for optional flag
			if tempParams.optional {
				optionalCount++
			}
			structParams = append(structParams, tempParams)
		}

		if optionalCount > 0 {
			optionalPresentsTmp, err := pd.getBitsValue(optionalCount)
			if err != nil {
				return err
			}
			optionalPresents = optionalPresentsTmp
			log.Debugf("optionalPresents is %0b", optionalPresents)
		}

		fieldIdx = -1
		for i := 0; i < structType.NumField(); i++ {
			if structType.Field(i).PkgPath != "" {
				log.Debugf("struct %s ignoring unexported field : %s", structType.Name(), structType.Field(i).Name)
				continue
			}
			fieldIdx++
			if structParams[fieldIdx].optional && optionalCount > 0 {
				optionalCount--
				if optionalPresents&(1<<optionalCount) == 0 {
					log.Debugf("Field \"%s\" in %s is OPTIONAL and not present", structType.Field(i).Name, structType)
					continue
				} else {
					log.Debugf("Field \"%s\" in %s is OPTIONAL and present", structType.Field(i).Name, structType)
				}
			}

			if err := parseField(val.Field(i), pd, structParams[fieldIdx]); err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		sliceType := fieldType
		newSlice, err := pd.parseSequenceOf(sizeExtensible, params, sliceType)
		if err != nil {
			return err
		}
		val.Set(newSlice)
		return nil
	case reflect.String:
		log.Debugf("Decoding PrintableString using Octet String decoding method")

		octetString, err := pd.parseOctetString(sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		if err != nil {
			return err
		}
		printableString := string(octetString)
		val.SetString(printableString)
		log.Debugf("Decoded PrintableString : \"%s\"", printableString)
		return nil
	case reflect.Interface:
		var choiceIdx int //:= 1
		var err error
		var choiceStruct reflect.Value
		if pd.canonicalOrdering {
			canonicalChoices, ok := pd.canonicalChoiceMap[params.oneofName]
			if !ok {
				return errors.NewInvalid("Expected a choice map with %s", params.oneofName)
			}
			// Parsing number of bytes which are following and verifying checksum
			err := pd.getCanonicalChoiceIndex()
			if err != nil {
				return err
			}

			if pd.unique == -1 {
				return errors.NewInvalid("Didn't find UNIQUE flag. Please revisit ASN1 definition")
			}

			choiceType, ok := canonicalChoices[pd.unique]
			if !ok {
				return errors.NewInvalid("Expected choice map %s to have index %d", params.oneofName, pd.unique)
			}
			choiceStruct = reflect.New(choiceType)
			if v.CanSet() {
				v.Set(choiceStruct)
			}

			log.Debugf("type is %s", choiceType.String())

			// Setting unique to -1 in order to reset CHOICE value (since we've already extracted it)
			pd.unique = -1
			// Setting this flag back to false in order to indicate the next CHOICE in canonical ordering
			pd.canonicalOrdering = false
		} else {
			choices, ok := pd.choiceMap[params.oneofName]
			if !ok {
				return errors.NewInvalid("Expected a choice map with %s", params.oneofName)
			}
			if len(choices) > 1 {
				var ieNotInExt = 0
				// Initiating flag which will indicate whether any extension item is presented in CHOICE
				var flag = false
				// Getting number of items in choice extension, if it exists
				// Assuming that items in CHOICE are sorted: firstly coming main items, then coming items from extension
				for j := 1; j <= len(choices); j++ {
					ie, ok := choices[j]
					if !ok {
						return errors.NewInvalid("Expected an index %d in a choice map with %s", j, params.oneofName)
					}
					itemParams := parseFieldParameters(ie.Field(0).Tag.Get("aper"))
					if itemParams.fromChoiceExt {
						flag = true
						ieNotInExt = j - 1
						break
					}
				}
				// When no items in extensions are found
				if !flag {
					ieNotInExt = len(choices)
				}

				log.Debugf("ValueExt is %v", params.valueExtensible)
				log.Debugf("FromChoiceExt is %v", params.fromChoiceExt)
				log.Debugf("Amount of values which are not in extension is %v", ieNotInExt)
				log.Debugf("Choice can be extended is %v", pd.choiceCanBeExtended)

				choiceIdx, err = pd.getChoiceIndex(params.valueExtensible, params.fromChoiceExt, ieNotInExt, len(choices))
				if err != nil {
					return err
				}

				log.Debugf("Handling interface %s for 'oneof' %s %d/%d", v.Type().String(), params.oneofName, choiceIdx, len(choices))
				choiceType, ok := choices[choiceIdx]
				if !ok {
					return errors.NewInvalid("Expected choice map %s to have index %d", params.oneofName, choiceIdx)
				}
				choiceStruct = reflect.New(choiceType)
				if v.CanSet() {
					v.Set(choiceStruct)
				}

				log.Debugf("type is %s", choiceType.String())
			} else {
				// treating the case when there is only single option
				// Firstly checking extension bit, if it is defined in the encoding schema
				if pd.choiceCanBeExtended {
					if bitsValue, err1 := pd.getBitsValue(1); err1 != nil {
						return err1
					} else if bitsValue != 0 {
						return errors.NewInvalid("unsupported value of CHOICE type is in Extensed was found. It's not possible to decode it without knowing it")
					}
					choiceType, ok := choices[1]
					if !ok {
						return errors.NewInvalid("Expected choice map %s to have index %d", params.oneofName, 1)
					}
					choiceStruct = reflect.New(choiceType)
					if v.CanSet() {
						v.Set(choiceStruct)
					}

					log.Debugf("type is %s", choiceType.String())

				} else {
					choiceType, ok := choices[1]
					if !ok {
						return errors.NewInvalid("Expected choice map %s to have index %d", params.oneofName, 1)
					}
					choiceStruct = reflect.New(choiceType)
					if v.CanSet() {
						v.Set(choiceStruct)
					}

					log.Debugf("type is %s", choiceType.String())
				}
			}
		}

		if err = parseField(choiceStruct.Elem(), pd, params); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unsupported: %s Kind: %s", v.Type().String(), v.Kind().String())
}

// Unmarshal parses the APER-encoded ASN.1 data structure b
// and uses the reflect package to fill in an arbitrary value pointed at by value.
// Because Unmarshal uses the reflect package, the structs
// being written to must use upper case field names.
//
// An ASN.1 INTEGER can be written to an int, int32, int64,
// If the encoded value does not fit in the Go type,
// Unmarshal returns a parse error.
//
// An ASN.1 BIT STRING can be written to a BitString.
//
// An ASN.1 OCTET STRING can be written to a []byte.
//
// An ASN.1 OBJECT IDENTIFIER can be written to an
// ObjectIdentifier.
//
// An ASN.1 ENUMERATED can be written to an Enumerated.
//
// Any of the above ASN.1 values can be written to an interface{}.
// The value stored in the interface has the corresponding Go type.
// For integers, that type is int64.
//
// An ASN.1 SEQUENCE OF x can be written
// to a slice if an x can be written to the slice's element type.
//
// An ASN.1 SEQUENCE can be written to a struct
// if each of the elements in the sequence can be
// written to the corresponding element in the struct.
//
// The following tags on struct fields have special meaning to Unmarshal:
//
//	optional        	OPTIONAL tag in SEQUENCE
//	sizeExt             specifies that size  is extensible
//	valueExt            specifies that value is extensible
//	sizeLB		        set the minimum value of size constraint
//	sizeUB              set the maximum value of value constraint
//	valueLB		        set the minimum value of size constraint
//	valueUB             set the maximum value of value constraint
//	default             sets the default value
//	openType            specifies the open Type
//  referenceFieldName	the string of the reference field for this type (only if openType used)
//  referenceFieldValue	the corresponding value of the reference field for this type (only if openType used)
//
// Other ASN.1 types are not supported; if it encounters them,
// Unmarshal returns a parse error.
func Unmarshal(b []byte, value interface{}, choiceMap map[string]map[int]reflect.Type, canonicalChoiceMap map[string]map[int64]reflect.Type) error {
	return UnmarshalWithParams(b, value, "", choiceMap, canonicalChoiceMap)
}

// UnmarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func UnmarshalWithParams(b []byte, value interface{}, params string, choiceMap map[string]map[int]reflect.Type, canonicalChoiceMap map[string]map[int64]reflect.Type) error {
	v := reflect.ValueOf(value).Elem()
	pd := &perBitData{b, 0, 0, choiceMap, -1, false, canonicalChoiceMap, false}
	err := parseField(v, pd, parseFieldParameters(params))
	if err != nil {
		return fmt.Errorf("Decoding failed with error %v\n%v", err, hex.Dump(b))
	}
	return nil
}
