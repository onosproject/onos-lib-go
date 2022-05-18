<!--
SPDX-FileCopyrightText: 2022-present Intel Corporation
SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

## Descriptions of the error messages
This little guide explains errors which may happen in the encoding/decoding process with Go APER library.

### Errors which can be obtained during encoding
* `bits Value is over capacity` can occur as an output of `putBitsValue()` function.
    * It happens when you're trying to encode value in lesser amount of bits than it requires, e.g., encoding 10 (1010 in
      binary) in 3 bits is impossible.
* `constraint Value is large than 65536` can occur as an output of `appendConstraintValue()` function.
    * It happens if in the value range in encoding, of typically `INTEGER`, is larger than `65536`. If this boundary is
      exceeded, then different set of rules for large numbers encoding have to be applied.
* `bitString Length is over upperbound:` can occur as an output of `appendBitString()` function.
    * It happens case of the constrained `BitString`, which doesn't have a `sizeExt` flag (= can't be extended). If the
      amount of bits in the `BitString` value is more than the upperbound, then this error is raised.
* `bitString Length is not match fix-sized:` can occur as an output of `appendBitString()` function.
    * It can happen with the fixed-size `BitString` (upperbound and lowerbound are the same), if the `BitString` value
      doesn't match pre-defined `BitString` length.
* `OctetString Length is over upperbound` can occur as an output of `appendOctetString()` function.
    * It happens when the constrained `OctetString` (represented as `[]byte{}`) length is larger than the defined
      upperbound and there is no `sizeExt` flag (= `OctetString` can't be extended).
* `OctetString Length is not match fix-sized:` can occur as an output of `appendOctetString()` function.
    * Same as for BitString.
    * It can happen with the fixed-size OctetString (upperbound and lowerbound are the same), if the OctetString value
      doesn't match pre-defined OctetString length.
* `Error encoding REAL - value (%v) is lower than lowerbound (%v)` or `Error encoding REAL - value (%v) is higher than
  upperbound (%v)` can occur as an output of `appendReal()` function.
    * It literally tells that the constrained `REAL` value exceeded the lowerbound or upperbound.
* `Error encoding REAL - numerical argument is out of domain` is raised in `appendReal()` function, when the `0` value
  is encoded. it is currently not supported in Go APER library.
* `Error encoding REAL - checksum verification failed. Encoded %v bytes, expected %v bytes to encode` is raised in
  `appendReal()` when the actual number of encoded bytes doesn't match expected number bytes to be encoded.
* `INTEGER value is smaller than lowerbound:` or `INTEGER value is larger than upperbound:` is raised in `appendInteger()`
  when the encoded value of a constrained `INTEGER` exceeds lowerbound or upperbound.
* `SEQUENCE OF Size is larger than upperbound:` or `SEQUENCE OF Size is lower than lowerbound:` is raised in `parseSequenceOf()`
  function when the number of items in the list is larger/lower than the maximum/minimum size of the list (amount of items).
* `encoding Length %d != fix-size %d` can be raised in `parseSequenceOf()` function in case of the list which has a fixed size
  and the number of items inside the list doesn't match this condition.
* `the upper bound of CHOICE is missing` is raised in the `appendChoiceIndex()` function if the CHOICE map entry is corrupted.
    * It happens when CHOICE map doesn't have an entry for certain CHOICE option, e.g.,
      if [entry for CGI in E2SM-RC](https://github.com/onosproject/onos-e2-sm/blob/6fd4546563ed112d47a89b173abcc31982ead240/servicemodels/e2sm_rc/v1/choiceOptions/choiceOptions.go#L98-L101)
      would be an empty one.
* `unsupport value of CHOICE type is in Extensed:` is raised in the `appendChoiceIndex()` function if the CHOICE can be extended and the
  item, which is being encoded, is not defined in the CHOICE extension, e.g., if with [this definition of ENB-ID in E2SM-RC](https://github.com/onosproject/onos-e2-sm/blob/6fd4546563ed112d47a89b173abcc31982ead240/servicemodels/e2sm_rc/v1/choiceOptions/choiceOptions.go#L148-L152)
  we tried to encode the 5th item, which doesn't exist.
    * Example from the above can happen in communication of different SM versions.
    * In practice, this error often happens either if one of the CHOICE map entries is missed or if we forgot to put `fromChoiceExt` flag
      on the item from CHOICE extension.
* `unsupport value of CHOICE type:` is raised in the `appendChoiceIndex()` function if the currently encoded item (**not** from the CHOICE extension)
  is not defined in the CHOICE map entry.
* `per: Value %v has exceeded its possible upperbound and shouldn't be encoded as a Normally small non-negative whole number` is raised in
  `appendNormallySmallNonNegativeWholeNumber()` function if the range of the encoded value exceeds `32767`. If this boundary is
  exceeded, then different set of rules have to be applied.
* `CHOICE index in canonical ordering for %v was not passed, please check encoding schema` is raised in the `appendCanonicalChoiceIndex()` function.
    * It happens when the `unique` flag is not passed and the encoder doesn't know which CHOICE option it encodes.
* `Expected to have key (%v) in CanonicalChoiceMap` is raised in the `appendCanonicalChoiceIndex()` function when it can't find a CHOICE map entry
  for parsed CHOICE option.
    * Solution here is to revisit your CHOICE map and make sure that all CHOICEs from the ASN.1 definition are included.
* `UNIQUE ID (%v) doesn't correspond to it's choice option (%v), got %v` is raised in the `appendCanonicalChoiceIndex()` function when passed CHOICE
  option doesn't match the obtained from CHOICE map entry.
* `aper: cannot marshal nil value` is raised in the `makeField()` function if the obtained Golang structure is nil (applies for only mandatory items in the message).
    * Solution is to include the missing item in the message.
* `Expected %d BitString byte(s) to contain %d bits. Got %d` is raised in the `makeField()` function if the BitString value doesn't match `BitString` length.
* `Expected last %d bits of byte array to be unused, and to contain only trailing zeroes. %s` is raised in the `makeField()` function, if
  the last few bits of a `BitString` are required to be trailing zeroes, but they are not.
    * For example, `BitString` with length 18, has to store its value in 3 bytes, e.g., `[]byte{0xFF, 0xFF, 0xFF}`. The last byte should have last 6 bits zeroed
      (3*8 - 18 = 6). In that case, valid `BitString` should be `[]byte{0xFF, 0xFF, 0xC0}`.
* `nil element in SEQUENCE type` is raised in the `makeField()` function if it gets a nil item (applied for mandatory items in the message).
    * Solution is to include the missing item in the message.
* `Expected a (canonical) choice map with` is raised in the `makeField()` function if in the **canonical** CHOICE map an entry for the item is not present in the
  **canonical** CHOICE map.
* `choice Index is nil at Field` is raised in the `makeField()` function if the APER tag is not included in the Golang driven Protobuf structure.
* `Expected a choice map with %s` is raised in the `makeField()` function if the CHOICE map entry for the item is not present in the CHOICE map.
* `unsupported: Type:` is raised in the `makeField()` function if the encoding of this particular type of the structure is not supported in the Go APER library.

### Errors which can be obtained during decoding
* `Get bits overflow, requireBits: %d, leftBits: %d` is raised in the `GetBitString()` function when the requested amount of bits to decode is larger than
  what is left in the unparsed bits.
    * Possible root cause is some missing APER tags or some additional APER tags, which are not required per ASN.1 definition.
* `Align Bit is not zero in (see last octet)` is raised in the `parseAlignBits()` function if the bits, which are expected to be an alignment bits (= zero bits)
  are not zeroes.
    * That means that some information is present in these bits. That may happen if some APER tags are missing.
* `Value range is negative` is raised in the `parseConstraintValue()` function.
    * It usually happens if the upperbound is lower than the lowerbound.
* `Constraint Value is large than 65536` is raised in the `parseConstraintValue()` function when the decoding value range is greater than `65536`.
    * In case it is greater, a different set of decoding rules is applied.
* `Parsed Length Out of Constraint` is raised in the `parseLength()` function if the length of a `BitString` or an `OctetString` wasn't encoded with regard to
  the chapter 10.9.3.6 or 10.9.3.7 of the [ITU-T X.691](https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf) recommendation.
* `PER data out of range` is raised if the Go APER library thinks that it parsed more bytes than it actually did.
    * This normally doesn't happen. If it does, enable debug mode in the Go APER library and try to see where the decoding gets stuck.
* `Error while parsing encoding base of REAL, obtained` is raised in `parseReal()` function if the encoding base of the `REAL` number is not 2.
    * This issue may occur only if in the encoder is used other encoding base for `REAL` numbers than 2. In this case, encoder is not aligned with [ITU-T X.691](https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf) recommendation.
* `Error parsing scaling factor - decoded bits expected to be 0, obtained %v` is raised in `parseReal()` function.
    * It means that the encoder, which produced APER bytes, is not compliant with ITU-T X.691 recommendation.
* `Error parsing exponent - decoded bits expected to be 0, obtained %v` is raised in `parseReal()` function.
    * Again, it means that the encoder, which produced APER bytes, is not compliant with ITU-T X.691 recommendation.
* `Parse Constraint Value failed:` is raised in `parseSequenceOf()` function, and it means that the decoder failed to parse the length of the list.
    * That may happen due to invalid constraints of list, if `sizeLB` or `sizeUB` tags are missing.
* `the upper bound of CHOICE is missing` is raised in `getChoiceIndex()` function if no entries from CHOICE map were parsed.
    * It is practically the same error as was described in the **Errors which can be obtained during encoding** section.
* `sequence truncated` error is raised in `parseField()` function if all APER bytes passed at input were processed, but the decoder still expects
  bytes for some other items from the message.
    * In practice, this issue may occur if you pass only part of the bytes, but not all of them to the decoder.
* `Expected a choice map with %s` is raised in the `parseField()` function if the CHOICE map entry for the item is not present in the CHOICE map.
* `Didn't find UNIQUE flag. Please revisit ASN1 definition` is raised in `parseField()` function if no structure were defined to be a constrained
  for CHOICE encoding in canonical order.
    * This is caused by missing `unique` flag in the Go driven Protobuf structure. Solution here would be to revisit ASN.1 definition and adjust your
      Protobuf to have a `unique` tag.
* `Expected choice map %s to have index %d` is raised in `parseField()` function when the CHOICE map doesn't have an entry for parsed item.
    * Solution here is to revisit your CHOICE map and make sure that all CHOICEs from the ASN.1 definition are included.
* `Expected an index %d in a choice map with %s` is raised in the `parseField` structure.
    * The root cause and the solution are similar to the ones described in the **Errors which can be obtained during encoding** section.
* `unsupported value of CHOICE type is in Extensed was found. It's not possible to decode it without knowing it` is raised in `parseField()` function.
    * This error is raised when there is only single entry in the CHOICE map and the extension bit is found.
        * Extension bit tells the decoder that the item from the extension should be decoded, but CHOICE map doesn't have any information about items which belong to the extension.
          In this case decoder doesn't know what structure to decode and thus panic
    * This may happen in the E2* communication of different versions.
* `unsupported: %s Kind:` is raised in `parseField()` if the item type decoding is not supported in the Go APER library.
