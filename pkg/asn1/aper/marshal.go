// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	//sequenceCanBeExtended bool
}

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
		err = fmt.Errorf("constraint Value is larger than 65536")
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

func howManyBitsNeeded(value int64) (bitAmount int32) {

	if value < 0 {
		value = -value
	}

	for {
		bitAmount++
		value = value >> 1
		if value == 0 {
			break
		}
	}

	return
}

func howManyBytesNeeded(value int64) (byteAmount int) {

	if value < 0 {
		value = -value
	}

	bitAmount := howManyBitsNeeded(value)

	for {
		if bitAmount > 8 {
			bitAmount = bitAmount - 8
			byteAmount++
		} else {
			break
		}
	}
	byteAmount++

	return
}

// it looks like encoding of REAL doesn't take into account constraints at all, so we don't bother about parsing constraints (only to check if value is within bounds)
// general rules are - mantissa should be an odd number or a 0
func (pd *perRawBitData) appendReal(value float64, lb *int64, ub *int64, valueExt bool) (err error) {

	log.Debugf("Encoding REAL number %v", value)

	// checking if value is within bounds
	if lb != nil {
		lowerBound := *lb
		if value < float64(lowerBound) {
			return errors.NewInvalid("Error encoding REAL - value (%v) is lower than lowerbound (%v)", value, float64(lowerBound))
		}
	}
	if ub != nil {
		upperBound := *ub
		if value > float64(upperBound) && !valueExt {
			return errors.NewInvalid("Error encoding REAL - value (%v) is higher than upperbound (%v)", value, float64(upperBound))
		}
	}

	// treating special case
	if value == float64(0) {
		// ITU-T X.691 refers to ITU-T X.690 and none of the defines proper form to encode 0 value.
		// I assume, that it can be represented as [0x03 0x80 0x00 0x00].
		// asn1c tool by Nokia doesn't treat this case - it returns error "numerical argument out of domain", returning it here
		return errors.NewInvalid("Error encoding REAL - numerical argument is out of domain")
	}

	var mantissa int64
	var exponent int64
	var p int
	var n int
	negativeExponent := false

	// First, checking whether we encode a whole number
	if value == math.Trunc(value) {
		log.Debugf("We're encoding a whole number")
		// Valid for whole numbers: divide value on 2 until the result is odd, once result is even, stop.
		// Encode power of 2 (obtained from division) as exponent and encode result of division as a mantissa
		// If the number is even at the beginning, then encode 0 as an exponent and encode
		// number as a mantissa (don't forget to put 00 octet in the beginning, for some reason I don't understand (yet))
		mantissa = int64(value)
		// mantissa can't be negative
		if mantissa < 0 {
			mantissa = -mantissa
		}
		for {
			if mantissa%2 != 0 {
				break
			}
			exponent++
			mantissa = mantissa / 2
		}
		log.Debugf("Obtained mantissa is %v, exponent is %v", mantissa, exponent)
	} else {
		log.Debugf("We're encoding a number with a floating point")
		// For values with numbers after decimal dot (radix), representation is different
		// Steps are following: multiply initial number by two until it becomes a whole number
		// if the number is not becoming a whole one, then multiply by 2 (max. 51 times), then
		// take the resulting number and encode it the same way as a whole number (exponent is encoded as its 2's complement)

		val := value
		for i := 0; i < 52; i++ { // 52 bits is maximum size of float
			if val == math.Trunc(val) {
				break
			}
			val = val * 2
			exponent++
		}
		// get the mantissa
		mantissa = int64(math.Trunc(val))
		// mantissa can't be negative
		if mantissa < 0 {
			mantissa = -mantissa
		}
		negativeExponent = true
		log.Debugf("Obtained mantissa is %v, value in computations is %v. Exponent is %v, it is negative (%v)", mantissa, val, exponent, negativeExponent)
		log.Debugf("Computing 2's complement for exponent (%v)", exponent)
		//alternative way - works for 8 bit representation numbers
		twosComplimentExp := 256 - exponent
		log.Debugf("2's complement for exponent %v is %v", exponent, twosComplimentExp)
		exponent = twosComplimentExp
		log.Debugf("It requires %v bits to store the value", howManyBitsNeeded(exponent))
	}

	// number of bytes to carry mantissa
	p = howManyBytesNeeded(mantissa)
	log.Debugf("Exponent/2 is %v, mantissa/2 is %v, exponent is %v, p is %v, mantissa is %v", exponent%2, mantissa%2, exponent, p, mantissa)
	// ToDo - nail down correct constraints to include 0x00 before mantissa
	// current constraint at least covers all cases in the unit test.
	// I, personally, would abandon these extra 0x00 in the beginning of mantissa. It's redundant
	if (exponent%2 == 0 || mantissa%2 != 0 || exponent == 0) && p < 7 && mantissa > 32 {
		// 7 is the maximum number of bytes to carry mantissa (because of max. 51 multiplication of 2),
		// mantissa > 32 and odd - reverse engineered from Nokia's asn1c tool
		p = p + 1
	}
	// number of bytes to carry exponent
	n = howManyBytesNeeded(exponent) // should be always 1 byte

	// computing length of the bytes needed to encode a number
	byteLength := n + p + 1
	log.Debugf("Amount of bytes to encode is %v: 1 byte for header, %v bytes for exponent, %v bytes for mantissa", byteLength, n, p)

	//aligning bits first
	pd.appendAlignBits()
	// storing number of bytes for reference
	numBytesStart := len(pd.bytes)

	// putting length of the bits first
	err = pd.putBitsValue(uint64(byteLength), 8)
	if err != nil {
		return err
	}

	// composing header
	// putting 1 (mandatory)
	err = pd.putBitsValue(1, 1)
	if err != nil {
		return err
	}
	// putting sign bit
	if value >= 0 {
		// if positive number
		err = pd.putBitsValue(0, 1)
		if err != nil {
			return err
		}
	} else {
		// if negative number
		err = pd.putBitsValue(1, 1)
		if err != nil {
			return err
		}
	}
	// putting an encoding base (always 2, so 00 bits)
	err = pd.putBitsValue(0, 2)
	if err != nil {
		return err
	}
	// putting a scale factor (always set to 0)
	err = pd.putBitsValue(0, 2)
	if err != nil {
		return err
	}
	// putting an exponent (always set to be 00 bits)
	err = pd.putBitsValue(0, 2)
	if err != nil {
		return err
	}

	err = pd.putBitsValue(uint64(exponent), uint(n)*8)
	if err != nil {
		return err
	}

	err = pd.putBitsValue(uint64(mantissa), uint(p)*8)
	if err != nil {
		return err
	}

	numBytesEnd := len(pd.bytes)
	if numBytesEnd-numBytesStart != byteLength+1 { // byteLength+1 is because 1 byte in the beginning stores the length of the following bytes
		return errors.NewInvalid("Error encoding REAL - checksum verification failed. Encoded %v bytes, expected %v bytes to encode", numBytesEnd-numBytesStart, byteLength+1)
	}

	return nil
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
		log.Debugf("Encoding index of CHOICE %d with upperbound %d", rawChoice, choiceBounds)
		if choiceBounds != 1 {
			if err := pd.appendConstraintValue(int64(choiceBounds), uint64(rawChoice)); err != nil {
				return err
			}
		} else {
			log.Debugf("Choice extension contains only single item, no need to encode the Choice index")
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
			return fmt.Errorf("unsupport value of CHOICE type: %v", rawChoice)
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
// when the encoded number exceeds 127 (in decimal). If the encoded number is less than 128, then the rest the number
// is encoded in 7 bits. In original definition it should treat the boundary 64 (and if the number is less than 64,
// then it encodes the number in 6 bits). Also, no octet alignment when number is between 64 and 256 is needed.
// Nokia's distribution is treating it in theirs way. Since theirs asn1c tool is officially recommended by O-RAN,
// this library needs to be aligned with them.
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
	// ToDo - sequenceCanBeExtended may cause potential problems
	threadedBytes := &perRawBitData{[]byte(""), 0, pd.choiceMap, -1, pd.canonicalChoiceMap, false}
	if err := threadedBytes.makeField(v, params); err != nil {
		return err
	}
	// ToDo - find workaround in logging
	//log.SetLevel(log.Debug)

	// encoding the number of upcoming bytes
	return pd.appendNormallySmallNonNegativeWholeNumber(uint64(len(threadedBytes.bytes)))
}

func (pd *perRawBitData) appendOpenType(v reflect.Value, params fieldParameters) error {

	// ToDo - sequenceCanBeExtended may cause potential problems
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

	case reflect.Float64:
		if err := pd.appendReal(v.Float(), params.valueLowerBound, params.valueUpperBound, params.valueExtensible); err != nil {
			return err
		}
		return nil

	case reflect.Struct:

		structType := fieldType
		var structParams []fieldParameters
		var optionalCount uint
		fromValueExtPresent := false // this is to indicate if any items in SEQUENCE Extension are actually present
		var optionalPresents uint64
		var choiceType string
		pd.choiceCanBeExtended = false
		sequenceCanBeExtended := false
		extensionHeader := false
		// It is only possible to decode extension which is defined in the encoding schema
		// struct extensive
		if params.valueExtensible && !params.choiceExt {
			sequenceCanBeExtended = true
			log.Debugf("SEQUENCE can be extended")
			//log.Debugf("Encoding Value Extensive Bit : true")
			//if err := pd.putBitsValue(0, 1); err != nil {
			//	return err
			//}
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
				if v.Field(i).Kind() == reflect.Ptr {
					log.Debugf("Type of the structure is %v, value is %v - %v", v.Field(i).Elem().Field(1).Kind(), v.Field(i).Elem().Field(1).Int(), v.Field(i).Elem().Field(3))
					pd.unique = v.Field(i).Elem().Field(3).Int()
				} else {
					pd.unique = v.Field(i).Int()
				}
				log.Debugf("Unique of type %v was found - it is %v", reflect.ValueOf(v.Field(i)), pd.unique)
			}
			// reflect.Slice is to handle the case, when structure is a byte array (i.e., []byte)
			if tempParams.fromValueExt && (v.Field(i).Type().Kind() == reflect.Ptr || v.Field(i).Type().Kind() == reflect.Slice) && !v.Field(i).IsNil() {
				log.Debugf("%v is from SEQUENCE extension and present", structType.Field(i).Name)
				fromValueExtPresent = true
			} else if tempParams.fromValueExt && (v.Field(i).Type().Kind() == reflect.Ptr || v.Field(i).Type().Kind() == reflect.Slice) && v.Field(i).IsNil() {
				log.Debugf("%v is from SEQUENCE extension and not present", structType.Field(i).Name)
			}
			choiceType = structType.Field(i).Tag.Get("protobuf_oneof")
			if choiceType == "" {
				// for optional flag
				if tempParams.optional && !tempParams.fromValueExt { // OPTIONAL items from SEQUENCE extension are not counted in the main header
					optionalCount++
					optionalPresents <<= 1
					if !v.Field(i).IsNil() {
						optionalPresents++
					}
				} else if v.Field(i).Type().Kind() == reflect.Ptr && v.Field(i).IsNil() && !tempParams.fromValueExt {
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

		if fromValueExtPresent {
			log.Debugf("SEQUENCE extension is present. Encoding Value Extensive Bit: true")
			if err := pd.putBitsValue(1, 1); err != nil {
				return err
			}
		} else if sequenceCanBeExtended {
			log.Debugf("SEQUENCE extension is not present. Encoding Value Extensive Bit: false")
			if err := pd.putBitsValue(0, 1); err != nil {
				return err
			}
		}
		if optionalCount > 0 {
			log.Debugf("putting optional(%d), optionalPresents is %0b", optionalCount, optionalPresents)
			if err := pd.putBitsValue(optionalPresents, optionalCount); err != nil {
				return err
			}
		}

		fieldIdx = -1
		for i := 0; i < structType.NumField(); i++ {
			log.Debugf("Iteration %v", i)
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
				}
				log.Debugf("Field \"%s\" in %s is OPTIONAL and present", structType.Field(fieldIdx).Name, structType)
			}
			log.Debugf("SEQUENCE Extension presence is %v, current field name is %v", fromValueExtPresent, structType.Field(i).Name)
			log.Debugf("fromValueExt is %v", structParams[fieldIdx].fromValueExt)
			log.Debugf("ExtensionHeader encoded is %v", extensionHeader)
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

						log.Debugf("ValueExt (i.e., if current CHOICE item is extensible) is %v", structParams[fieldIdx].valueExtensible)
						log.Debugf("FromChoiceExt is %v", structParams[fieldIdx].fromChoiceExt)
						log.Debugf("Amount of values which are not in extension is %v", ieNotInExt)
						log.Debugf("Choice map length is %v", len(choices))
						log.Debugf("Choice can be extended is %v", pd.choiceCanBeExtended)
						if err := pd.appendChoiceIndex(present, structParams[fieldIdx].valueExtensible, structParams[fieldIdx].fromChoiceExt, ieNotInExt, len(choices)); err != nil {
							return err
						}
					} else {
						// ToDo - test if choice can be extended or not and put an Extensed bit
						if pd.choiceCanBeExtended {
							log.Debugf("CHOICE can be potentially extended, putting 0 bit to indicate that")
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
			} else if structParams[fieldIdx].fromValueExt && fromValueExtPresent { // making sure that the items in the extension are actually present
				// encoding items from the value extension
				log.Debugf("Current SEQUENCE can be extended is %v", sequenceCanBeExtended)
				if !extensionHeader && sequenceCanBeExtended {
					log.Debugf("Encoding SEQUENCE Extension header")
					// encoding the header of the value extension
					// obtaining total number of items
					totalItemsInExtension := structType.NumField() - i
					if totalItemsInExtension < 0 {
						log.Debugf("Something went wrong - total amount of instances in the extension is %d (negative)\n", totalItemsInExtension)
						return nil
					}
					// encoding this number
					log.Debugf("Total amount of items in the extension is %d - encoding it as a small non-negative whole number", totalItemsInExtension)
					if totalItemsInExtension <= 127 {
						// encoding number of items over 0 (i.e., totalItemsInExtension - 1)
						if err := pd.putBitsValue(uint64(totalItemsInExtension-1), 7); err != nil {
							return err
						}
					} else {
						// fallback to the old way of decoding...
						if err := pd.appendNormallySmallNonNegativeWholeNumber(uint64(totalItemsInExtension)); err != nil {
							return err
						}
					}

					// encoding indication that actual item is present in the extension
					for ext := 0; ext < totalItemsInExtension; ext++ {
						if structParams[fieldIdx+ext].fromValueExt && (v.Field(i+ext).Type().Kind() == reflect.Ptr || v.Field(i+ext).Type().Kind() == reflect.Slice) && !val.Field(i+ext).IsNil() {
							log.Debugf("%v is from SEQUENCE extension and present", structType.Field(i+ext).Name)
							if err := pd.putBitsValue(1, 1); err != nil {
								return err
							}
						} else if structParams[fieldIdx+ext].fromValueExt && (v.Field(i+ext).Type().Kind() == reflect.Ptr || v.Field(i+ext).Type().Kind() == reflect.Slice) && val.Field(i+ext).IsNil() {
							log.Debugf("%v is from SEQUENCE extension and NOT present", structType.Field(i+ext).Name)
							if err := pd.putBitsValue(0, 1); err != nil {
								return err
							}
						}
					}

					// aligning bytes
					pd.appendAlignBits()
					// indicating that extension header is encoded
					extensionHeader = true
				}
				// proceeding with encoding the items in extension in a regular way
				if !val.Field(i).IsNil() {
					log.Debugf("Encoding %v - an item from the extension", structType.Field(i).Name)
					// create threaded bytes, encode the item and encode its length
					threadedBytes := &perRawBitData{[]byte(""), 0, pd.choiceMap, -1, pd.canonicalChoiceMap, false}
					if err := threadedBytes.makeField(val.Field(i), structParams[fieldIdx]); err != nil {
						return err
					}

					// encoding the number of upcoming bytes
					if err := pd.appendNormallySmallNonNegativeWholeNumber(uint64(len(threadedBytes.bytes))); err != nil {
						return err
					}

					// if item is present in extension, encoding it, otherwise, iterating over the next items in the extension
					if err := pd.makeField(val.Field(i), structParams[fieldIdx]); err != nil {
						return err
					}
				}
			} else if !structParams[fieldIdx].fromValueExt {
				// if the value is not the CHOICE, or an item in Extension, or an OPTIONAL item,
				// then it should be mandatory present in the message - encoding it
				if err := pd.makeField(val.Field(i), structParams[fieldIdx]); err != nil {
					return err
				}
			}
			// if we are hitting the last item, set by default the extension header to false
			if fieldIdx == structType.NumField()-1 {
				extensionHeader = false
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
	// ToDo - sequenceCanBeExtended may cause potential problems
	//log.SetLevel(logging.DebugLevel)
	pd := &perRawBitData{[]byte(""), 0, choiceMap, -1, canonicalChoiceMap, false}
	err := pd.makeField(reflect.ValueOf(val), parseFieldParameters(params))
	if err != nil {
		return nil, err
	} else if len(pd.bytes) == 0 {
		pd.bytes = make([]byte, 1)
	}
	return pd.bytes, nil
}
