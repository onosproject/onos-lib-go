// Copyright 2020-present Open Networking Foundation.
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

package logging

import (
	"go.uber.org/zap"
	"time"
)

// Field is a structured logger field
type Field interface {
	getZapField() zap.Field
}

type zapField struct {
	field zap.Field
}

func (f *zapField) Name() string {
	return f.field.Key
}

func (f *zapField) getZapField() zap.Field {
	return f.field
}

func String(name string, value string) Field {
	return &zapField{
		field: zap.String(name, value),
	}
}

func Stringp(name string, value *string) Field {
	return &zapField{
		field: zap.Stringp(name, value),
	}
}

func Strings(name string, value []string) Field {
	return &zapField{
		field: zap.Strings(name, value),
	}
}

func Int(name string, value int) Field {
	return &zapField{
		field: zap.Int(name, value),
	}
}

func Intp(name string, value *int) Field {
	return &zapField{
		field: zap.Intp(name, value),
	}
}

func Ints(name string, value []int) Field {
	return &zapField{
		field: zap.Ints(name, value),
	}
}

func Int32(name string, value int32) Field {
	return &zapField{
		field: zap.Int32(name, value),
	}
}

func Int32p(name string, value *int32) Field {
	return &zapField{
		field: zap.Int32p(name, value),
	}
}

func Int32s(name string, value []int32) Field {
	return &zapField{
		field: zap.Int32s(name, value),
	}
}

func Int64(name string, value int64) Field {
	return &zapField{
		field: zap.Int64(name, value),
	}
}

func Int64p(name string, value *int64) Field {
	return &zapField{
		field: zap.Int64p(name, value),
	}
}

func Int64s(name string, value []int64) Field {
	return &zapField{
		field: zap.Int64s(name, value),
	}
}

func Uint(name string, value uint) Field {
	return &zapField{
		field: zap.Uint(name, value),
	}
}

func Uintp(name string, value *uint) Field {
	return &zapField{
		field: zap.Uintp(name, value),
	}
}

func Uints(name string, value []uint) Field {
	return &zapField{
		field: zap.Uints(name, value),
	}
}

func Uint32(name string, value uint32) Field {
	return &zapField{
		field: zap.Uint32(name, value),
	}
}

func Uint32p(name string, value *uint32) Field {
	return &zapField{
		field: zap.Uint32p(name, value),
	}
}

func Uint32s(name string, value []uint32) Field {
	return &zapField{
		field: zap.Uint32s(name, value),
	}
}

func Uint64(name string, value uint64) Field {
	return &zapField{
		field: zap.Uint64(name, value),
	}
}

func Uint64p(name string, value *uint64) Field {
	return &zapField{
		field: zap.Uint64p(name, value),
	}
}

func Uint64s(name string, value []uint64) Field {
	return &zapField{
		field: zap.Uint64s(name, value),
	}
}

func Float32(name string, value float32) Field {
	return &zapField{
		field: zap.Float32(name, value),
	}
}

func Float32p(name string, value *float32) Field {
	return &zapField{
		field: zap.Float32p(name, value),
	}
}

func Float32s(name string, value []float32) Field {
	return &zapField{
		field: zap.Float32s(name, value),
	}
}

func Float64(name string, value float64) Field {
	return &zapField{
		field: zap.Float64(name, value),
	}
}

func Float64p(name string, value *float64) Field {
	return &zapField{
		field: zap.Float64p(name, value),
	}
}

func Float64s(name string, value []float64) Field {
	return &zapField{
		field: zap.Float64s(name, value),
	}
}

func Complex64(name string, value complex64) Field {
	return &zapField{
		field: zap.Complex64(name, value),
	}
}

func Complex64p(name string, value *complex64) Field {
	return &zapField{
		field: zap.Complex64p(name, value),
	}
}

func Complex64s(name string, value []complex64) Field {
	return &zapField{
		field: zap.Complex64s(name, value),
	}
}

func Complex128(name string, value complex128) Field {
	return &zapField{
		field: zap.Complex128(name, value),
	}
}

func Complex128p(name string, value *complex128) Field {
	return &zapField{
		field: zap.Complex128p(name, value),
	}
}

func Complex128s(name string, value []complex128) Field {
	return &zapField{
		field: zap.Complex128s(name, value),
	}
}

func Bool(name string, value bool) Field {
	return &zapField{
		field: zap.Bool(name, value),
	}
}

func Boolp(name string, value *bool) Field {
	return &zapField{
		field: zap.Boolp(name, value),
	}
}

func Bools(name string, value []bool) Field {
	return &zapField{
		field: zap.Bools(name, value),
	}
}

func Time(name string, value time.Time) Field {
	return &zapField{
		field: zap.Time(name, value),
	}
}

func Timep(name string, value *time.Time) Field {
	return &zapField{
		field: zap.Timep(name, value),
	}
}

func Times(name string, value []time.Time) Field {
	return &zapField{
		field: zap.Times(name, value),
	}
}

func Duration(name string, value time.Duration) Field {
	return &zapField{
		field: zap.Duration(name, value),
	}
}

func Durationp(name string, value *time.Duration) Field {
	return &zapField{
		field: zap.Durationp(name, value),
	}
}

func Durations(name string, value []time.Duration) Field {
	return &zapField{
		field: zap.Durations(name, value),
	}
}

func Byte(name string, value byte) Field {
	return &zapField{
		field: zap.Binary(name, []byte{value}),
	}
}

func Bytes(name string, value []byte) Field {
	return &zapField{
		field: zap.Binary(name, value),
	}
}

func ByteString(name string, value []byte) Field {
	return &zapField{
		field: zap.ByteString(name, value),
	}
}

func ByteStrings(name string, value [][]byte) Field {
	return &zapField{
		field: zap.ByteStrings(name, value),
	}
}
