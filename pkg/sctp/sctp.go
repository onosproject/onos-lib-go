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

package sctp

import (
	"encoding/binary"
	"unsafe"
)

// NativeEndian ...
var NativeEndian binary.ByteOrder

//var sndRcvInfoSize uintptr

func init() {
	i := uint16(1)
	if *(*byte)(unsafe.Pointer(&i)) == 0 {
		NativeEndian = binary.BigEndian
	} else {
		NativeEndian = binary.LittleEndian
	}
	//sndRcvInfoSize = unsafe.Sizeof(SndRcvInfo{})
}
