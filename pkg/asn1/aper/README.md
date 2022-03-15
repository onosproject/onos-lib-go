<!--
SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# Free 5GC "aper" project

This library provides encoding and decoding of ASN.1 structs to APER encoding format.

This is copied from `v1.0.0` of https://github.com/free5gc/aper

This library was modified in order to correspond to the ASN.1 APER encoding rules specified
in [ITU-T X.691](https://www.itu.int/ITU-T/studygroups/com17/languages/X.691-0207.pdf) recommendation. Now this library
generates the same APER bytes as [this asn1c](https://github.com/nokia/asn1c) tool.

#### Changes:

* Library was adjusted to use Protobuf instead of Go structs
* Unified `BitString` definition, no it is handled as a complex structure:
    * Value field is represented with `[]byte`
    * Len field is represented with `uint32`
* Fixed `BitString` encoding and decoding brought by `[]byte` representation specificity
* Fixed `Integer` encoding, now it can encode negative values as well
* logrus has been replaced by own logging
* Errors are generated through our own package
* Can handle structs that have one or more private fields (lowercase)
* Fixed `CHOICE` encoding
    * Now for the correct encoding or decoding a `CHOICE` map is needed to be passed to the encoder.
        * This map can be generated
          with [`protoc-gen-choice`](https://github.com/onosproject/onos-e2-sm/protoc-gen-choice) plugin out of Protobuf
    * For `CHOICEs` with the single item inside, `CHOICE` index is not being encoded.
* Introduced correct encoding and decoding of the items from the `CHOICE` extension 
* Introduced encoding and decoding for `CHOICEs` in Canonical ordering (used in O-RAN's E2AP)
* Introduced encoding of normally small non-negative `Integers`
* Fixed race condition issue

All of the aforementioned changes were verified with unit tests.

#### Tags which have to be included in the Protobuf to ensure the correct encoding and decoding

* `choiceIdx` specifies the `CHOICE` index
* `choiceExt` specifies that the `CHOICE` structure can contain items in its extension
* `fromChoiceExt` specifies the item which belongs to the extension of `CHOICE`
* `valueExt` specifies that the ASN.1 structure like `SEQUENCE`, `ENUMERATED`, `INTEGER` can be extended
* `valueUB` specifies the upperbound of the `INTEGER`
* `valueLB` specifies the lowerbound of the `INTEGER`
* `sizeExt`  specifies that the ASN.1 structure like `OCTET STRING`, `PrintableString`, `SEQUENCE OF` can be extended
* `sizeUB` specifies the upperbound of the `OCTET STRING`, `PrintableString`, `SEQUENCE OF`
* `sizeLB` specifies the lowerbound of the `OCTET STRING`, `PrintableString`, `SEQUENCE OF`
* `optional` specifies that the item is `OPTIONAL` (i.e., not mandatory to be present in the structure) 
* `canonicalOrder` specifies that the `CHOICE` follows canonical ordering
* `unique` specifies that this item is used as an input to indicate which `CHOICE` option is encoded or decoded 
  * It is a mandatory prerequisite for Canonical `CHOICE` encoding
* `align` specifies that the Octet Alignment should be performed after this item

#### Known issues

* Encoding and decoding of `REAL` structures is not handled
* `BitString` doesn't support encoding and decoding of value extension