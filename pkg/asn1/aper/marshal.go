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
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"math"
	"reflect"
)

type perRawBitData struct {
	bytes               []byte
	bitsOffset          uint
	choiceMap           map[string]map[int]reflect.Type
	unique              int64
	canonicalChoiceMap  map[string]map[int64]reflect.Type
	choiceCanBeExtended bool
}

// Assuming that UNIQUE ID is being treated as INTEGER
// By default, we don't know if UNIQUE items are present in ASN.1 definition
//var unique int64 = -1

func perRawBitLog(numBits uint64, byteLen int, bitsOffset uint, value interface{}) string {
	if reflect.TypeOf(value).Kind() == reflect.Uint64 {
		return fmt.Sprintf("  [PER put %2d bits, byteLen(after): %d, bitsOffset(after): %d, value: 0x%0x]",
			numBits, byteLen, bitsOffset, reflect.ValueOf(value).Uint())
	}
	return fmt.Sprintf("  [PER put %2d bits, byteLen(after): %d, bitsOffset(after): %d, value: 0x%0x]",
		numBits, byteLen, bitsOffset, reflect.ValueOf(value).Bytes())

}

func (pd *perRawBitData) bitCarry() {
	pd.bitsOffset = pd.bitsOffset & 0x07
}
func (pd *perRawBitData) appendAlignBits() {
	if alignBits := uint64(8-pd.bitsOffset&0x7) & 0x7; alignBits != 0 {
		log.Debugf("Aligning %d bits", alignBits)
		log.Debugf("%s", perRawBitLog(alignBits, len(pd.bytes), 0, []byte{0x00}))
	}
	pd.bitsOffset = 0
}

func (pd *perRawBitData) putBitString(bytes []byte, numBits uint) (err error) {
	bytes = bytes[:(numBits+7)>>3]
	if pd.bitsOffset == 0 {
		pd.bytes = append(pd.bytes, bytes...)
		pd.bitsOffset = (numBits & 0x7)
		log.Debugf("%s", perRawBitLog(uint64(numBits), len(pd.bytes), pd.bitsOffset, bytes))
		return
	}
	bitsLeft := 8 - pd.bitsOffset
	currentByte := len(pd.bytes) - 1
	if numBits <= bitsLeft {
		pd.bytes[currentByte] |= (bytes[0] >> pd.bitsOffset)
	} else {
		bytes = append([]byte{0x00}, bytes...)
		var shiftBytes []byte
		if shiftBytes, err = GetBitString(bytes, bitsLeft, pd.bitsOffset+numBits); err != nil {
			return
		}
		pd.bytes[currentByte] |= shiftBytes[0]
		pd.bytes = append(pd.bytes, shiftBytes[1:]...)
		bytes = bytes[1:]
	}
	pd.bitsOffset = (numBits & 0x7) + pd.bitsOffset
	pd.bitCarry()
	log.Debugf("%s", perRawBitLog(uint64(numBits), len(pd.bytes), pd.bitsOffset, bytes))
	return
}

func (pd *perRawBitData) putBitsValue(value uint64, numBits uint) (err error) {
	if numBits == 0 {
		return
	}
	Byteslen := (numBits + 7) >> 3
	tempBytes := make([]byte, Byteslen)
	bitOff := numBits & 0x7
	if bitOff == 0 {
		bitOff = 8
	}
	LeftbitOff := 8 - bitOff
	tempBytes[Byteslen-1] = byte((value << LeftbitOff) & 0xff)
	value >>= bitOff
	var i int
	for i = int(Byteslen) - 2; value > 0; i-- {
		if i < 0 {
			err = fmt.Errorf("bits Value is over capacity")
			return
		}
		tempBytes[i] = byte(value & 0xff)
		value >>= 8
	}

	return pd.putBitString(tempBytes, numBits)
}

func (pd *perRawBitData) appendConstraintValue(valueRange int64, value uint64) (err error) {
	log.Debugf("Putting Constraint Value %d with range %d", value, valueRange)

	var bytes uint
	if valueRange <= 255 {
		if valueRange < 0 {
			err = fmt.Errorf("value range is negative: %v", valueRange)
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
		err = pd.putBitsValue(value, i)
		return
	} else if valueRange == 256 {
		bytes = 1
	} else if valueRange <= 65536 {
		bytes = 2
	} else {
		err = fmt.Errorf("constraint Value is large than 65536")
		return
	}
	pd.appendAlignBits()
	err = pd.putBitsValue(value, bytes*8)
	return
}

func (pd *perRawBitData) appendLength(sizeRange int64, value uint64) (err error) {
	if sizeRange <= 65536 && sizeRange > 0 {
		return pd.appendConstraintValue(sizeRange, value)
	}
	pd.appendAlignBits()
	log.Debugf("Putting Length of Value : %d", value)
	if value <= 127 {
		err = pd.putBitsValue(value, 8)
		return
	} else if value <= 16383 {
		value |= 0x8000
		err = pd.putBitsValue(value, 16)
		return
	}

	value = (value >> 14) | 0xc0
	err = pd.putBitsValue(value, 8)
	return
}

func (pd *perRawBitData) appendBitString(bytes []byte, bitsLength uint64, extensive bool,
	lowerBoundPtr *int64, upperBoundPtr *int64) (err error) {
	var lb, ub, sizeRange int64 = 0, -1, -1
	if lowerBoundPtr != nil {
		lb = *lowerBoundPtr
		if upperBoundPtr != nil {
			ub = *upperBoundPtr
			if bitsLength <= uint64(ub) {
				sizeRange = ub - lb + 1
			} else if !extensive {
				err = fmt.Errorf("bitString Length is over upperbound: obtained bytes %v of length %v, UB is %v", bytes, bitsLength, ub)
				return
			}
			if extensive {
				log.Debugf("Putting size Extension Value")
				if sizeRange == -1 {
					if errTmp := pd.putBitsValue(1, 1); errTmp != nil {
						log.Errorf("putBitsValue(1, 1) error: %v", errTmp)
					}
					lb = 0
				} else {
					if errTmp := pd.putBitsValue(0, 1); errTmp != nil {
						log.Errorf("putBitsValue(0, 1) error: %v", errTmp)
					}
				}
			}

		}
	}

	if ub > 65535 {
		sizeRange = -1
	}
	sizes := (bitsLength + 7) >> 3
	shift := (8 - bitsLength&0x7)
	if shift != 8 {
		bytes[sizes-1] &= (0xff << shift)
	}

	if sizeRange == 1 {
		if bitsLength != uint64(ub) {
			err = fmt.Errorf("bitString Length(%d) is not match fix-sized : %d", bitsLength, ub)
		}
		log.Debugf("Encoding BIT STRING size %d", ub)
		if sizes > 2 {
			pd.appendAlignBits()
			pd.bytes = append(pd.bytes, bytes...)
			pd.bitsOffset = uint(ub & 0x7)
			log.Debugf("%s", perRawBitLog(bitsLength, len(pd.bytes), pd.bitsOffset, bytes))
		} else {
			err = pd.putBitString(bytes, uint(bitsLength))
		}
		log.Debugf("Encoded BIT STRING (length = %d): 0x%0x", bitsLength, bytes)
		return
	}
	rawLength := bitsLength - uint64(lb)

	var byteOffset, partOfRawLength uint64
	for {
		if rawLength > 65536 {
			partOfRawLength = 65536
		} else if rawLength >= 16384 {
			partOfRawLength = rawLength & 0xc000
		} else {
			partOfRawLength = rawLength
		}
		if err = pd.appendLength(sizeRange, partOfRawLength); err != nil {
			return
		}
		partOfRawLength += uint64(lb)
		sizes := (partOfRawLength + 7) >> 3
		log.Debugf("Encoding BIT STRING size %d", partOfRawLength)
		if partOfRawLength == 0 {
			return
		}
		pd.appendAlignBits()
		pd.bytes = append(pd.bytes, bytes[byteOffset:byteOffset+sizes]...)
		log.Debugf("%s", perRawBitLog(partOfRawLength, len(pd.bytes), pd.bitsOffset, bytes))
		log.Debugf("Encoded BIT STRING (length = %d): 0x%0x", partOfRawLength,
			bytes[byteOffset:byteOffset+sizes])
		rawLength -= (partOfRawLength - uint64(lb))
		if rawLength > 0 {
			byteOffset += sizes
		} else {
			pd.bitsOffset += uint(partOfRawLength & 0x7)
			// pd.appendAlignBits()
			break
		}
	}
	return err

}

func (pd *perRawBitData) appendOctetString(bytes []byte, extensive bool, lowerBoundPtr *int64,
	upperBoundPtr *int64) error {
	byteLen := uint64(len(bytes))
	var lb, ub, sizeRange int64 = 0, -1, -1
	if lowerBoundPtr != nil {
		lb = *lowerBoundPtr
		if upperBoundPtr != nil {
			ub = *upperBoundPtr
			if byteLen <= uint64(ub) {
				sizeRange = ub - lb + 1
			} else if !extensive {
				err := fmt.Errorf("OctetString Length is over upperbound: obtained bytes %v of length %v, UB is %v", bytes, byteLen, ub)
				return err
			}
			if extensive {
				log.Debugf("Putting size Extension Value")
				if sizeRange == -1 {
					if errTmp := pd.putBitsValue(1, 1); errTmp != nil {
						log.Debugf("putBitsValue(1, 1) err: %v", errTmp)
					}
					lb = 0
				} else {
					if errTmp := pd.putBitsValue(0, 1); errTmp != nil {
						log.Debugf("putBitsValue(0, 1) err: %v", errTmp)
					}
				}
			}

		}
	}

	if ub > 65535 {
		sizeRange = -1
	}

	if sizeRange == 1 {
		if byteLen != uint64(ub) {
			err := fmt.Errorf("OctetString Length(%d) is not match fix-sized : %d", byteLen, ub)
			return err
		}
		log.Debugf("Encoding OCTET STRING size %d", ub)
		if byteLen > 2 {
			pd.appendAlignBits()
			pd.bytes = append(pd.bytes, bytes...)
			log.Debugf("%s", perRawBitLog(byteLen*8, len(pd.bytes), 0, bytes))
		} else {
			err := pd.putBitString(bytes, uint(byteLen*8))
			return err
		}
		log.Debugf("Encoded OCTET STRING (length = %d): 0x%0x", byteLen, bytes)
		return nil
	}
	rawLength := byteLen - uint64(lb)

	var byteOffset, partOfRawLength uint64
	for {
		if rawLength > 65536 {
			partOfRawLength = 65536
		} else if rawLength >= 16384 {
			partOfRawLength = rawLength & 0xc000
		} else {
			partOfRawLength = rawLength
		}
		if err := pd.appendLength(sizeRange, partOfRawLength); err != nil {
			return err
		}
		partOfRawLength += uint64(lb)
		log.Debugf("Encoding OCTET STRING size %d", partOfRawLength)
		if partOfRawLength == 0 {
			return nil
		}
		pd.appendAlignBits()
		pd.bytes = append(pd.bytes, bytes[byteOffset:byteOffset+partOfRawLength]...)
		log.Debugf("%s", perRawBitLog(partOfRawLength*8, len(pd.bytes), pd.bitsOffset, bytes))
		log.Debugf("Encoded OCTET STRING (length = %d): 0x%0x", partOfRawLength,
			bytes[byteOffset:byteOffset+partOfRawLength])
		rawLength -= (partOfRawLength - uint64(lb))
		if rawLength > 0 {
			byteOffset += partOfRawLength
		} else {
			// pd.appendAlignBits()
			break
		}
	}
	return nil

}

func (pd *perRawBitData) appendBool(value bool) (err error) {
	log.Debugf("Encoding BOOLEAN Value %t", value)
	if value {
		err = pd.putBitsValue(1, 1)
		log.Debugf("Encoded BOOLEAN Value : 0x1")
	} else {
		err = pd.putBitsValue(0, 1)
		log.Debugf("Encoded BOOLEAN Value : 0x0")
	}
	return
}

func (pd *perRawBitData) appendInteger(value int64, extensive bool, lowerBoundPtr *int64, upperBoundPtr *int64) error {
	var lb, valueRange int64 = 0, 0
	if lowerBoundPtr != nil {
		lb = *lowerBoundPtr
		if value < lb {
			return fmt.Errorf("INTEGER value is smaller than lowerbound: obtained %v, LB is %v", value, lb)
		}
		if upperBoundPtr != nil {
			ub := *upperBoundPtr
			if value <= ub {
				valueRange = ub - lb + 1
			} else if !extensive {
				return fmt.Errorf("INTEGER value is larger than upperbound: obtained %v, UB is %v", value, ub)
			}
			if extensive {
				log.Debugf("Putting value Extension bit")
				if valueRange == 0 {
					log.Debugf("Encoding INTEGER with Unconstraint Value")
					valueRange = -1
					if errTmp := pd.putBitsValue(1, 1); errTmp != nil {
						fmt.Printf("pd.putBitsValue(1, 1) error: %v", errTmp)
					}
				} else {
					log.Debugf("Encoding INTEGER with Value Range(%d..%d)", lb, ub)
					if errTmp := pd.putBitsValue(0, 1); errTmp != nil {
						fmt.Printf("pd.putBitsValue(0, 1) error: %v", errTmp)
					}
				}
			}

		} else {
			log.Debugf("Encoding INTEGER with Semi-Constraint Range(%d..)", lb)
		}
	} else {
		log.Debugf("Encoding INTEGER with Unconstraint Value")
		valueRange = -1
	}

	unsignedValue := uint64(value)
	var rawLength uint
	if valueRange == 1 {
		log.Debugf("Value of INTEGER is fixed")

		return nil
	}
	if value < 0 {
		y := value >> 63
		valueXor := value ^ y
		unsignedValue = uint64((valueXor - y))
	}
	if valueRange <= 0 {
		unsignedValue >>= 7
	} else if valueRange <= 65536 {
		return pd.appendConstraintValue(valueRange, uint64(value-lb))
	} else {
		unsignedValue >>= 8
	}
	for rawLength = 1; rawLength <= 127; rawLength++ {
		if unsignedValue == 0 {
			break
		}
		unsignedValue >>= 8
	}

	// putting length
	if valueRange <= 0 {
		// semi-constraint or unconstraint
		pd.appendAlignBits()
		pd.bytes = append(pd.bytes, byte(rawLength))
		log.Debugf("Encoding INTEGER Length %d in one byte", rawLength)

		log.Debugf("%s", perRawBitLog(8, len(pd.bytes), pd.bitsOffset, uint64(rawLength)))
	} else {
		// valueRange > 65536
		var byteLen uint
		unsignedValueRange := uint64(valueRange - 1)
		for byteLen = 1; byteLen <= 127; byteLen++ {
			unsignedValueRange >>= 8
			if unsignedValueRange <= 1 {
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

		// New implementation of algorithm
		absoluteDistanceToLB := value - lb
		for rawLength = 1; rawLength <= 127; rawLength++ {
			if absoluteDistanceToLB == 0 {
				break
			}
			absoluteDistanceToLB >>= 8
		}
		if value-lb == 0 {
			rawLength = 1
		} else {
			rawLength--
		}
		log.Debugf("Encoding INTEGER Length %d-1 in %d bits", rawLength, i)
		if err := pd.putBitsValue(uint64(rawLength-1), i); err != nil {
			return err
		}
	}

	log.Debugf("Encoding INTEGER %d with %d bytes", value, rawLength)

	rawLength *= 8
	pd.appendAlignBits()

	if valueRange < 0 {
		mask := int64(1<<rawLength - 1)
		return pd.putBitsValue(uint64(value&mask), rawLength)
	}
	value -= lb
	return pd.putBitsValue(uint64(value), rawLength)
}

// append ENUMERATED type but do not implement extensive value and different value with index
func (pd *perRawBitData) appendEnumerated(value uint64, extensive bool, lowerBoundPtr *int64,
	upperBoundPtr *int64) error {
	if lowerBoundPtr == nil || upperBoundPtr == nil {
		return fmt.Errorf("ENUMERATED value constraint is error - make sure that at least LB or UB tag is passed")
	}
	lb, ub := *lowerBoundPtr, *upperBoundPtr
	if signedValue := int64(value); signedValue > ub {
		if extensive {
			return fmt.Errorf("Unsupport the extensive value of ENUMERATED")
		}
		return fmt.Errorf("ENUMERATED value is larger than upperbound: obtained %v, UB is %v", value, ub)
	} else if signedValue < lb {
		return fmt.Errorf("ENUMERATED value is smaller than lowerbound: obtained %v, LB is %v", value, lb)
	}
	if extensive {
		if err := pd.putBitsValue(0, 1); err != nil {
			return err
		}
	}

	valueRange := ub - lb + 1
	log.Debugf("Encoding ENUMERATED Value : %d with Value Range(%d..%d)", value, lb, ub)
	if valueRange > 1 {
		return pd.appendConstraintValue(valueRange, value)
	}
	return nil

}

func (pd *perRawBitData) parseSequenceOf(v reflect.Value, params fieldParameters) error {
	var lb, ub, sizeRange int64 = 0, -1, -1
	numElements := int64(v.Len())
	if params.sizeLowerBound != nil && *params.sizeLowerBound < 65536 {
		lb = *params.sizeLowerBound
	}
	if params.sizeUpperBound != nil && *params.sizeUpperBound < 65536 {
		ub = *params.sizeUpperBound
		if params.sizeExtensible {
			if numElements > ub {
				if err := pd.putBitsValue(1, 1); err != nil {
					return err
				}
			} else {
				if err := pd.putBitsValue(0, 1); err != nil {
					return err
				}
				sizeRange = ub - lb + 1
			}
		} else if numElements > ub {
			return fmt.Errorf("SEQUENCE OF Size is larger than upperbound: %v, size is %v, UB is %v", v.Type(), numElements, ub)
		} else {
			sizeRange = ub - lb + 1
		}
	} else {
		sizeRange = -1
	}

	if numElements < lb {
		return fmt.Errorf("SEQUENCE OF Size is lower than lowerbound: %v, size is %v, LB is %v", v.Type(), numElements, lb)
	} else if sizeRange == 1 {
		log.Debugf("Encoding Length of \"SEQUENCE OF\"  with fix-size %d", ub)
		if numElements != ub {
			return fmt.Errorf("encoding Length %d != fix-size %d", numElements, ub)
		}
	} else if sizeRange > 0 {
		log.Debugf("Encoding Length(%d) of \"SEQUENCE OF\"  with Size Range(%d..%d)", numElements, lb, ub)
		if err := pd.appendConstraintValue(sizeRange, uint64(numElements-lb)); err != nil {
			return err
		}
	} else {
		log.Debugf("Encoding Length(%d) of \"SEQUENCE OF\" with Semi-Constraint Range(%d..)", numElements, lb)
		pd.appendAlignBits()
		pd.bytes = append(pd.bytes, byte(numElements&0xff))
		log.Debugf("%s", perRawBitLog(8, len(pd.bytes), pd.bitsOffset, uint64(numElements)))
	}
	log.Debugf("Encoding  \"SEQUENCE OF\" struct %s with len(%d)", v.Type().Elem().Name(), numElements)
	params.sizeExtensible = false
	params.sizeUpperBound = nil
	params.sizeLowerBound = nil
	for i := 0; i < v.Len(); i++ {
		if err := pd.makeField(v.Index(i), params); err != nil {
			return err
		}
	}
	return nil
}

func (pd *perRawBitData) appendChoiceIndex(present int, extensive bool, fromChoiceExtension bool, numItemsNotInExtension int, choiceMapLen int) error {
	log.Debugf("Current present is %v", present)
	if fromChoiceExtension {
		// putting an extensive bit first
		if err := pd.putBitsValue(1, 1); err != nil {
			return err
		}
		rawChoice := present - 1 - numItemsNotInExtension
		choiceBounds := choiceMapLen - numItemsNotInExtension
		if choiceBounds < 1 {
			return fmt.Errorf("the upper bound of CHOICE is missing")
		} else if extensive && rawChoice > choiceBounds {
			return fmt.Errorf("unsupport value of CHOICE type is in Extensed: %v", rawChoice)
		}
		log.Debugf("Encoding Present index of CHOICE  %d - 1", present)
		if choiceBounds != 1 {
			if err := pd.appendConstraintValue(int64(choiceBounds), uint64(rawChoice)); err != nil {
				return err
			}
		}
	} else {
		if pd.choiceCanBeExtended {
			if err := pd.putBitsValue(0, 1); err != nil {
				return err
			}
		}
		rawChoice := present - 1
		choiceBounds := numItemsNotInExtension
		log.Debugf("The upperbound of choice is %v", choiceBounds)
		if choiceBounds < 1 {
			return fmt.Errorf("the upper bound of CHOICE is missing")
		} else if extensive && rawChoice > choiceBounds {
			return fmt.Errorf("unsupport value of CHOICE type is in Extensed: %v", rawChoice)
		}
		log.Debugf("Encoding Present index of CHOICE  %d - 1", present)
		if choiceBounds != 1 {
			if err := pd.appendConstraintValue(int64(choiceBounds), uint64(rawChoice)); err != nil {
				return err
			}
		}
	}
	return nil
}

// appendNormallySmallNonNegativeWholeNumber function does not fully correspond to its original definition
// provided in chapter 20.4 of Olivier DuBuisson book "ASN.1. Communication between Heterogeneous systems".
// Instead, this function was aligned to correspond to the needs of E2AP APER encoding handled by asn1c tool,
// which is provided by Nokia (https://github.com/nokia/asn1c). In particular, it adds 1 in the header only
// when the encoded number exceeds 127 (in decimal). If number is less than 128, then the rest the number
// is encoded in 7 bits. In original definition it should treat the boundary 64 (and if the number is less than 64,
// then it encodes the number in 6 bits). Also, no octet alignment when number is between 64 and 256 is needed.
// Nokia's distribution is treating it in theirs way. Since theirs asn1c tool is officially recommended by O-RAN,
// we need to be aligned with them.
func (pd *perRawBitData) appendNormallySmallNonNegativeWholeNumber(value uint64) error {

	if value > 32767 {
		return fmt.Errorf("aper: Value %v has exceeded its possible upperbound and shouldn't be encoded as "+
			"Normally small non-negative whole number. If this issue is related to the E2AP then it is a PANIC!! T_T", value)
	}
	if value > 127 {
		if err := pd.putBitsValue(1, 1); err != nil {
			return err
		}
		if value < 256 {
			pd.appendAlignBits()
			return pd.putBitsValue(value, 8)
		}
		return pd.putBitsValue(value, 15)
	}
	if err := pd.putBitsValue(0, 1); err != nil {
		return err
	}
	return pd.putBitsValue(value, 7)
}

// Canonical CHOICE index is literally number of bytes which are following after current byte. Could be re-used as a checksum.
// In fact, we don't even need to know about the structure in a CHOICE option, but it would be good to check it (especially for the decoding).
func (pd *perRawBitData) appendCanonicalChoiceIndex(canonicalChoiceMap map[int64]reflect.Type, v reflect.Value, params fieldParameters) error {

	if pd.unique == -1 {
		return fmt.Errorf("CHOICE index in canonical ordering for %v was not passed, please check encoding schema", v.Type())
	}
	log.Debugf("UNIQUE index is %v", pd.unique)

	// Verifying that this CHOICE option exists in a CanonicalChoiceMap
	val, ok := canonicalChoiceMap[pd.unique]
	if !ok {
		return errors.NewInvalid("Expected to have key (%v) in CanonicalChoiceMap\n%v", pd.unique, canonicalChoiceMap)
	}

	//Now comparing obtained CHOICE option with actually passed CHOICE option
	if val.Name() != v.Elem().Type().Name() {
		return errors.NewInvalid("UNIQUE ID (%v) doesn't correspond to it's choice option (%v), got %v", pd.unique, canonicalChoiceMap[pd.unique].Name(), v.Elem().Type().Name())
	}

	// Verification of correct CHOICE option was done, now setting unique variable back to -1 and waiting for the other CHOICE to come
	pd.unique = -1

	// aligning bits first - necessary to encode in full byte
	pd.appendAlignBits()

	// ToDo - find workaround in logging
	//log.SetLevel(log.Info)
	threadedBytes := &perRawBitData{[]byte(""), 0, pd.choiceMap, -1, pd.canonicalChoiceMap, false}
	if err := threadedBytes.makeField(v, params); err != nil {
		return err
	}
	// ToDo - find workaround in logging
	//log.SetLevel(log.Debug)

	// encoding the number of upcoming bytes
	if err := pd.appendNormallySmallNonNegativeWholeNumber(uint64(len(threadedBytes.bytes))); err != nil {
		return err
	}

	return nil
}

func (pd *perRawBitData) appendOpenType(v reflect.Value, params fieldParameters) error {

	pdOpenType := &perRawBitData{[]byte(""), 0, pd.choiceMap, -1, pd.canonicalChoiceMap, false}
	log.Debugf("Encoding OpenType %s to temp RawData", v.Type().String())
	if err := pdOpenType.makeField(v, params); err != nil {
		return err
	}
	openTypeBytes := pdOpenType.bytes
	rawLength := uint64(len(pdOpenType.bytes))
	log.Debugf("Encoding OpenType %s RawData : 0x%0x(%d bytes)", v.Type().String(), pdOpenType.bytes,
		rawLength)

	var byteOffset, partOfRawLength uint64
	for {
		if rawLength > 65536 {
			partOfRawLength = 65536
		} else if rawLength >= 16384 {
			partOfRawLength = rawLength & 0xc000
		} else {
			partOfRawLength = rawLength
		}
		if err := pd.appendLength(-1, partOfRawLength); err != nil {
			return err
		}
		log.Debugf("Encoding Part of OpenType RawData size %d", partOfRawLength)
		if partOfRawLength == 0 {
			return nil
		}
		pd.appendAlignBits()
		pd.bytes = append(pd.bytes, openTypeBytes[byteOffset:byteOffset+partOfRawLength]...)
		log.Debugf("%s", perRawBitLog(partOfRawLength*8, len(pd.bytes), pd.bitsOffset, openTypeBytes))
		log.Debugf("Encoded OpenType RawData (length = %d): 0x%0x", partOfRawLength,
			openTypeBytes[byteOffset:byteOffset+partOfRawLength])
		rawLength -= partOfRawLength
		if rawLength > 0 {
			byteOffset += partOfRawLength
		} else {
			pd.appendAlignBits()
			break
		}
	}

	log.Debugf("Encoded OpenType %s", v.Type().String())
	return nil
}

func (pd *perRawBitData) makeField(v reflect.Value, params fieldParameters) error {

	//log.Debugf("Current bit offset is %v", pd.bitsOffset)
	//if pd.bitsOffset != 0 {
	//	log.Debugf("Shifting bits to the right explicitly..")
	//	err := pd.shiftLastBit()
	//	if err != nil {
	//		return err
	//	}
	//	log.Debugf("Bits offset is %v, explicitly aligning bits", pd.bitsOffset)
	//	pd.appendAlignBits()
	//}
	//log.Debugf("Current APER bytes with bits offset %v are \n%v", pd.bitsOffset, hex.Dump(pd.bytes))

	log.Debugf("Encoding %s %s", v.Type().String(), v.Kind().String())
	if !v.IsValid() {
		return fmt.Errorf("aper: cannot marshal nil value")
	}
	// If the field is an interface{} then recurse into it.
	if v.Kind() == reflect.Interface && v.Type().NumMethod() == 0 {
		return pd.makeField(v.Elem(), params)
	}
	if v.Kind() == reflect.Ptr {
		return pd.makeField(v.Elem(), params)
	}
	fieldType := v.Type()

	// We deal with the structures defined in this package first.
	switch fieldType {
	case BitStringType:
		bytes := v.Field(3).Bytes()
		length := v.Field(4).Uint()
		expected := int(math.Ceil(float64(length) / 8))
		unused := 8 - int(math.Mod(float64(length), 8))
		if unused == 8 {
			unused = 0
		}
		unusedMask := (1 << unused) - 1
		log.Debugf("Handling BitString with %v. Len %d", bytes, length)
		if len(bytes) != expected {
			return errors.NewInvalid("Expected %d BitString byte(s) to contain %d bits. Got %d",
				expected, length, len(bytes))
		} else if len(bytes) > 0 && bytes[len(bytes)-1]&byte(unusedMask) != 0 {
			return errors.NewInvalid("Expected last %d bits of byte array to be unused, and to contain only trailing zeroes. %s",
				unused, hex.Dump(bytes))
		}
		err := pd.appendBitString(v.Field(3).Bytes(), v.Field(4).Uint(), params.sizeExtensible, params.sizeLowerBound,
			params.sizeUpperBound)
		return err
	case reflect.TypeOf([]uint8{}):
		err := pd.appendOctetString(v.Bytes(), params.sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		return err
	default:
		log.Debugf("not a built in type %v", fieldType)
	}
	switch val := v; val.Kind() {
	case reflect.Bool:
		err := pd.appendBool(v.Bool())
		return err
	case reflect.Int, reflect.Int32, reflect.Int64:
		if err := pd.appendInteger(v.Int(), params.valueExtensible, params.valueLowerBound, params.valueUpperBound); err != nil {
			return err
		}
		if params.unique && v.Int() > 0 {
			pd.unique = v.Int()
		}
		if params.align {
			pd.appendAlignBits()
		}
		return nil

	case reflect.Struct:

		structType := fieldType
		var structParams []fieldParameters
		var optionalCount uint
		var optionalPresents uint64
		var choiceType string
		pd.choiceCanBeExtended = false
		// ToDo - currently this is an incorrect treatment of possible extensions with structs.
		// It is only possible to decode extension which is defined in the encoding schema
		// struct extensive TODO: support extensed type
		if params.valueExtensible && !params.choiceExt {
			log.Debugf("Encoding Value Extensive Bit : true")
			if err := pd.putBitsValue(0, 1); err != nil {
				return err
			}
		}
		if params.choiceExt {
			pd.choiceCanBeExtended = true
			log.Debugf("CHOICE can be extended")
		}
		//sequenceType = structType.NumField() <= 0 || structType.Field(0).Name != "Present"
		// pass tag for optional
		fieldIdx := -1
		for i := 0; i < structType.NumField(); i++ {
			fieldIdx++
			if structType.Field(i).PkgPath != "" {
				log.Debugf("struct %s ignoring unexported field : %s", structType.Name(), structType.Field(i).Name)
				continue
			}
			log.Debugf("Handling %s", structType.Field(i).Name)
			tempParams := parseFieldParameters(structType.Field(i).Tag.Get("aper"))
			if tempParams.unique {
				pd.unique = v.Field(i).Int()
				log.Debugf("Unique of type %v was found - it is %v", reflect.ValueOf(v.Field(i)), pd.unique)
			}
			choiceType = structType.Field(i).Tag.Get("protobuf_oneof")
			if choiceType == "" {
				// for optional flag
				if tempParams.optional {
					optionalCount++
					optionalPresents <<= 1
					if !v.Field(i).IsNil() {
						optionalPresents++
					}
				} else if v.Field(i).Type().Kind() == reflect.Ptr && v.Field(i).IsNil() {
					return fmt.Errorf("nil element in SEQUENCE type %v", v.Field(i).Type())
				}
			} else {
				if v.Field(i).Interface() == nil {
					continue
				}
				concreteType := reflect.TypeOf(v.Field(i).Interface()).Elem()
				tempParams = parseFieldParameters(concreteType.Field(0).Tag.Get("aper"))
				tempParams.oneofName = choiceType
			}

			structParams = append(structParams, tempParams)
		}
		if optionalCount > 0 {
			log.Debugf("putting optional(%d), optionalPresents is %0b", optionalCount, optionalPresents)
			if err := pd.putBitsValue(optionalPresents, optionalCount); err != nil {
				return err
			}
		}

		//// CHOICE or OpenType
		//if choiceType != "" { // TODO: remove hard coding
		//	present := int(*structParams[0].choiceIndex)
		//	ub := structParams[0].valueUpperBound
		//	if err := pd.appendChoiceIndex(present, structParams[0].valueExtensible, ub); err != nil {
		//		return err
		//	}
		//
		//	if err := pd.makeField(val.Field(fieldIdx), structParams[0]); err != nil {
		//		return err
		//	}
		//	return nil
		//}

		fieldIdx = -1
		for i := 0; i < structType.NumField(); i++ {
			if structType.Field(i).PkgPath != "" {
				log.Debugf("struct %s ignoring unexported field : %s", structType.Name(), structType.Field(i).Name)
				continue
			}
			fieldIdx++
			// optional
			if len(structParams) <= fieldIdx {
				continue
			}
			if structParams[fieldIdx].optional && optionalCount > 0 {
				optionalCount--
				if optionalPresents&(1<<optionalCount) == 0 {
					log.Debugf("Field \"%s\" in %s is OPTIONAL and not present", structType.Field(fieldIdx).Name, structType)
					continue
				} else {
					log.Debugf("Field \"%s\" in %s is OPTIONAL and present", structType.Field(fieldIdx).Name, structType)
				}
			}
			// for open type reference
			//if structParams[fieldIdx].openType {
			//	fieldName := structParams[fieldIdx].referenceFieldName
			//	var index int
			//	for index = 0; index < fieldIdx; index++ {
			//		if structType.Field(index).Name == fieldName {
			//			break
			//		}
			//	}
			//	if index == fieldIdx {
			//		return fmt.Errorf("open type is not reference to the other field in the struct")
			//	}
			//	structParams[fieldIdx].referenceFieldValue = new(int64)
			//	value, err := getReferenceFieldValue(val.Field(index))
			//	if err != nil {
			//		return err
			//	}
			//	*structParams[fieldIdx].referenceFieldValue = value
			//}
			//if params.unique {
			//	choiceID = v.Field(i)
			//}
			if structParams[fieldIdx].oneofName != "" {
				if params.canonicalOrder {
					tempParams := structParams[fieldIdx]
					tempParams.valueExtensible = false
					canonicalChoices, ok := pd.canonicalChoiceMap[choiceType]
					if !ok {
						return errors.NewInvalid("Expected a (canonical) choice map with %s", choiceType)
					}
					if err := pd.appendCanonicalChoiceIndex(canonicalChoices, reflect.ValueOf(v.Field(i).Interface()), tempParams); err != nil {
						return err
					}
				} else {
					if structParams[fieldIdx].choiceIndex == nil {
						return fmt.Errorf("choice Index is nil at Field %v, Index %v.\n Make sure all aper tags are injected in your proto", v.Field(i).Type(), fieldIdx)
					}
					present := int(*structParams[fieldIdx].choiceIndex)
					choices, ok := pd.choiceMap[choiceType]
					//When there is only one item in the choice, you don't need to encode choice index
					if !ok {
						return errors.NewInvalid("Expected a choice map with %s", choiceType)
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
								return errors.NewInvalid("Expected an index %d in a choice map with %s", j, choiceType)
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

						log.Debugf("ValueExt is %v", structParams[fieldIdx].valueExtensible)
						log.Debugf("FromChoiceExt is %v", structParams[fieldIdx].fromChoiceExt)
						log.Debugf("Amount of values which are not in extension is %v", ieNotInExt)
						log.Debugf("Choice can be extended is %v", pd.choiceCanBeExtended)
						if err := pd.appendChoiceIndex(present, structParams[fieldIdx].valueExtensible, structParams[fieldIdx].fromChoiceExt, ieNotInExt, len(choices)); err != nil {
							return err
						}
					} else {
						// ToDo - test if choice can be extensed or not and put and Extensed bit
						if pd.choiceCanBeExtended {
							log.Debugf("CHOICE can be potentially extensed, putting 0 bit to indicate that")
							if err := pd.putBitsValue(0, 1); err != nil {
								return err
							}
						}
					}
				}
				tempParams := structParams[fieldIdx]
				tempParams.valueExtensible = false
				// Here the CHOICE field is being encoded
				if err := pd.makeField(reflect.ValueOf(v.Field(i).Interface()), tempParams); err != nil {
					return err
				}
			} else {
				if err := pd.makeField(val.Field(i), structParams[fieldIdx]); err != nil {
					return err
				}
			}
		}
		return nil
	case reflect.Slice:
		err := pd.parseSequenceOf(v, params)
		return err
	case reflect.String:
		printableString := v.String()
		log.Debugf("Encoding PrintableString : \"%s\" using Octet String decoding method", printableString)
		err := pd.appendOctetString([]byte(printableString), params.sizeExtensible, params.sizeLowerBound,
			params.sizeUpperBound)
		return err
	//case reflect.Array:
	//	log.Debugf("Ignoring array: \"%s\"", v.String())
	//	return nil
	default:
		log.Debugf("Unhandled: \"%s\"", v.String())
	}
	return fmt.Errorf("unsupported: Type:%s Kind:%s", v.Type().String(), v.Kind().String())
}

// Marshal returns the ASN.1 encoding of val.
func Marshal(val interface{}, choiceMap map[string]map[int]reflect.Type, canonicalChoiceMap map[string]map[int64]reflect.Type) ([]byte, error) {
	return MarshalWithParams(val, "", choiceMap, canonicalChoiceMap)
}

// MarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func MarshalWithParams(val interface{}, params string, choiceMap map[string]map[int]reflect.Type, canonicalChoiceMap map[string]map[int64]reflect.Type) ([]byte, error) {
	pd := &perRawBitData{[]byte(""), 0, choiceMap, -1, canonicalChoiceMap, false}
	err := pd.makeField(reflect.ValueOf(val), parseFieldParameters(params))
	if err != nil {
		return nil, err
	} else if len(pd.bytes) == 0 {
		pd.bytes = make([]byte, 1)
	}
	return pd.bytes, nil
}
