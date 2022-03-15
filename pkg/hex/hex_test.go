// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package hex

import (
	"encoding/hex"
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestServicemodel_Asn1BytesToByte(t *testing.T) {
	indHdrFmt1Hex, err := Asn1BytesToByte("1f21222324187478740000034f4e4640736f6d6554797065066f6e660c3734370" +
		"0d4bc08803039201a85")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", indHdrFmt1Hex)

	indMsgFmt1Hex, err := Asn1BytesToByte("0e4030380000036f6e66011500004020747269616c013ffba021222340400" +
		"10203000a7c0f000f0001724000fa00000400007a0001c7000314000000400300023039440480001a8500")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", indMsgFmt1Hex)

	actDefFmt3Hex, err := Asn1BytesToByte("00010c4000036f6e6600000040747269616c000048210200c90115203038")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", actDefFmt3Hex)

	evntTrigDefHex, err := Asn1BytesToByte("00010c4000036f6e6600000040747269616c000048210200c90115203038")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", evntTrigDefHex)

	ranFuncDescHex, err := Asn1BytesToByte("74046f6e660000056f69643132330700736f6d654465736372697074696f6e" +
		"01150000430021222300d4bc08803039201a8500000000034f4e4600212223000000200000010b01006f6e66010f00010b01006f6e66010f0" +
		"00041a04f70656e4e6574776f726b696e67000017012f0118")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", ranFuncDescHex)

	ctrlHdrHex, err := Asn1BytesToByte("3412f410abd4bc000101")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", ctrlHdrHex)

	ctrlMsgHex, err := Asn1BytesToByte("0000010100504349000114")
	assert.NilError(t, err)
	fmt.Printf("Output of Asn1BytesToByte is \n%x\n", ctrlMsgHex)

	ctrlOutHex, err := Asn1BytesToByte("20000000001400010a")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%x\n", ctrlOutHex)
}

func TestServicemodel_HexDumpToByte(t *testing.T) {
	indHdrFmt1Hex, err := DumpToByte("00000000  1f 21 22 23 24 18 74 78  74 00 00 03 4f 4e 46 40  " +
		"|.!\"#$.txt...ONF@|\n        00000010  73 6f 6d 65 54 79 70 65  06 6f 6e 66 0c 37 34 37  |someType.onf.747|\n" +
		"        00000020  00 d4 bc 08 80 30 39 20  1a 85                    |.....09 ..|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(indHdrFmt1Hex))

	indMsgFmt1Hex, err := DumpToByte("00000000  0e 40 30 38 00 00 03 6f  6e 66 01 15 00 00 40 20  |.@08...onf....@ |" +
		"	00000010  74 72 69 61 6c 01 3f fb  a0 21 22 23 40 40 01 02  |trial.?..!\"#@@..|" +
		"	00000020  03 00 0a 7c 0f 00 0f 00  01 72 40 00 fa 00 00 04  |...|.....r@.....|" +
		"	00000030  00 00 7a 00 01 c7 00 03  14 00 00 00 40 03 00 02  |..z.........@...|" +
		"	00000040  30 39 44 04 80 00 1a 85  00                       |09D......|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(indMsgFmt1Hex))

	actDefFmt3Hex, err := DumpToByte("00000000  00 01 0c 40 00 03 6f 6e  66 00 00 00 40 74 72 69  " +
		"|...@..onf...@tri|\n        00000010  61 6c 00 00 48 21 02 00  c9 01 15 20 30 38        |al..H!..... 08|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(actDefFmt3Hex))

	evntTrigDefHex, err := DumpToByte("00000000  00 01 0c                                          |...|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(evntTrigDefHex))

	ranFuncDescHex, err := DumpToByte("00000000  74 04 6f 6e 66 00 00 05  6f 69 64 31 32 33 07 00  |t.onf...oid123..|" +
		"	00000010  73 6f 6d 65 44 65 73 63  72 69 70 74 69 6f 6e 01  |someDescription.|" +
		"	00000020  15 00 00 43 00 21 22 23  00 d4 bc 08 80 30 39 20  |...C.!\"#.....09 |" +
		"	00000030  1a 85 00 00 00 00 03 4f  4e 46 00 21 22 23 00 00  |.......ONF.!\"#..|" +
		"	00000040  00 20 00 00 01 0b 01 00  6f 6e 66 01 0f 00 01 0b  |. ......onf.....|" +
		"	00000050  01 00 6f 6e 66 01 0f 00  00 41 a0 4f 70 65 6e 4e  |..onf....A.OpenN|" +
		"	00000060  65 74 77 6f 72 6b 69 6e  67 00 00 17 01 2f 01 18  |etworking..../..|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(ranFuncDescHex))

	ctrlHdrHex, err := DumpToByte("00000000  34 12 f4 10 ab d4 bc 00  01 01                    |4.........|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(ctrlHdrHex))

	ctrlMsgHex, err := DumpToByte("00000000  00 00 01 01 00 50 43 49  00 01 14                 |.....PCI...|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(ctrlMsgHex))

	ctrlOutHex, err := DumpToByte("00000000  20 00 00 00 00 14 00 01  0a                       | ........|")
	assert.NilError(t, err)
	fmt.Printf("Output of HexDumpToByte is \n%v\n", hex.Dump(ctrlOutHex))
}
