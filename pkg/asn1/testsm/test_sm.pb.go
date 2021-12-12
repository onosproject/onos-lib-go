	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// enumerated from tes_sm.asn1:161
//{TEST-Enumerated}
type TestEnumerated int32

const (
	TestEnumerated_TEST_ENUMERATED_ENUM1 TestEnumerated = 0
	TestEnumerated_TEST_ENUMERATED_ENUM2 TestEnumerated = 1
	TestEnumerated_TEST_ENUMERATED_ENUM3 TestEnumerated = 2
	TestEnumerated_TEST_ENUMERATED_ENUM4 TestEnumerated = 3
	TestEnumerated_TEST_ENUMERATED_ENUM5 TestEnumerated = 4
	TestEnumerated_TEST_ENUMERATED_ENUM6 TestEnumerated = 5
)

// Enum value maps for TestEnumerated.
var (
	TestEnumerated_name = map[int32]string{
		0: "TEST_ENUMERATED_ENUM1",
		1: "TEST_ENUMERATED_ENUM2",
		2: "TEST_ENUMERATED_ENUM3",
		3: "TEST_ENUMERATED_ENUM4",
		4: "TEST_ENUMERATED_ENUM5",
		5: "TEST_ENUMERATED_ENUM6",
	}
	TestEnumerated_value = map[string]int32{
		"TEST_ENUMERATED_ENUM1": 0,
		"TEST_ENUMERATED_ENUM2": 1,
		"TEST_ENUMERATED_ENUM3": 2,
		"TEST_ENUMERATED_ENUM4": 3,
		"TEST_ENUMERATED_ENUM5": 4,
		"TEST_ENUMERATED_ENUM6": 5,
	}
)

func (x TestEnumerated) Enum() *TestEnumerated {
	p := new(TestEnumerated)
	*p = x
	return p
}

func (x TestEnumerated) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnumerated) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_asn1_testsm_test_sm_proto_enumTypes[0].Descriptor()
}

func (TestEnumerated) Type() protoreflect.EnumType {
	return &file_pkg_asn1_testsm_test_sm_proto_enumTypes[0]
}

func (x TestEnumerated) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestEnumerated.Descriptor instead.
func (TestEnumerated) EnumDescriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{0}
}

// enumerated from tes_sm.asn1:170
//{TEST-EnumeratedExtensible}
type TestEnumeratedExtensible int32

const (
	TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM1 TestEnumeratedExtensible = 0
	TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM2 TestEnumeratedExtensible = 1
	TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM3 TestEnumeratedExtensible = 2
	TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM4 TestEnumeratedExtensible = 3
	TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM5 TestEnumeratedExtensible = 4
	TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM6 TestEnumeratedExtensible = 5
)

// Enum value maps for TestEnumeratedExtensible.
var (
	TestEnumeratedExtensible_name = map[int32]string{
		0: "TEST_ENUMERATED_EXTENSIBLE_ENUM1",
		1: "TEST_ENUMERATED_EXTENSIBLE_ENUM2",
		2: "TEST_ENUMERATED_EXTENSIBLE_ENUM3",
		3: "TEST_ENUMERATED_EXTENSIBLE_ENUM4",
		4: "TEST_ENUMERATED_EXTENSIBLE_ENUM5",
		5: "TEST_ENUMERATED_EXTENSIBLE_ENUM6",
	}
	TestEnumeratedExtensible_value = map[string]int32{
		"TEST_ENUMERATED_EXTENSIBLE_ENUM1": 0,
		"TEST_ENUMERATED_EXTENSIBLE_ENUM2": 1,
		"TEST_ENUMERATED_EXTENSIBLE_ENUM3": 2,
		"TEST_ENUMERATED_EXTENSIBLE_ENUM4": 3,
		"TEST_ENUMERATED_EXTENSIBLE_ENUM5": 4,
		"TEST_ENUMERATED_EXTENSIBLE_ENUM6": 5,
	}
)

func (x TestEnumeratedExtensible) Enum() *TestEnumeratedExtensible {
	p := new(TestEnumeratedExtensible)
	*p = x
	return p
}

func (x TestEnumeratedExtensible) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestEnumeratedExtensible) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_asn1_testsm_test_sm_proto_enumTypes[1].Descriptor()
}

func (TestEnumeratedExtensible) Type() protoreflect.EnumType {
	return &file_pkg_asn1_testsm_test_sm_proto_enumTypes[1]
}

func (x TestEnumeratedExtensible) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestEnumeratedExtensible.Descriptor instead.
func (TestEnumeratedExtensible) EnumDescriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{1}
}

//{TEST-EnumeratedExtensible}
type TestFullyOptionalSequenceItem4 int32

const (
	TestFullyOptionalSequenceItem4_TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_ONE TestFullyOptionalSequenceItem4 = 0
	TestFullyOptionalSequenceItem4_TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_TWO TestFullyOptionalSequenceItem4 = 1
)

// Enum value maps for TestFullyOptionalSequenceItem4.
var (
	TestFullyOptionalSequenceItem4_name = map[int32]string{
		0: "TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_ONE",
		1: "TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_TWO",
	}
	TestFullyOptionalSequenceItem4_value = map[string]int32{
		"TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_ONE": 0,
		"TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_TWO": 1,
	}
)

func (x TestFullyOptionalSequenceItem4) Enum() *TestFullyOptionalSequenceItem4 {
	p := new(TestFullyOptionalSequenceItem4)
	*p = x
	return p
}

func (x TestFullyOptionalSequenceItem4) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestFullyOptionalSequenceItem4) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_asn1_testsm_test_sm_proto_enumTypes[2].Descriptor()
}

func (TestFullyOptionalSequenceItem4) Type() protoreflect.EnumType {
	return &file_pkg_asn1_testsm_test_sm_proto_enumTypes[2]
}

func (x TestFullyOptionalSequenceItem4) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestFullyOptionalSequenceItem4.Descriptor instead.
func (TestFullyOptionalSequenceItem4) EnumDescriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{2}
}

// sequence from tes_sm.asn1:15
// {TEST-UnconstrainedInt}
type TestUnconstrainedInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrUciA int32 `protobuf:"varint,1,opt,name=attr_uci_a,json=attrUciA,proto3" json:"attr_uci_a,omitempty"`
	AttrUciB int32 `protobuf:"varint,2,opt,name=attr_uci_b,json=attrUciB,proto3" json:"attr_uci_b,omitempty"`
}

func (x *TestUnconstrainedInt) Reset() {
	*x = TestUnconstrainedInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestUnconstrainedInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestUnconstrainedInt) ProtoMessage() {}

func (x *TestUnconstrainedInt) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestUnconstrainedInt.ProtoReflect.Descriptor instead.
func (*TestUnconstrainedInt) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{0}
}

func (x *TestUnconstrainedInt) GetAttrUciA() int32 {
	if x != nil {
		return x.AttrUciA
	}
	return 0
}

func (x *TestUnconstrainedInt) GetAttrUciB() int32 {
	if x != nil {
		return x.AttrUciB
	}
	return 0
}

// sequence from tes_sm.asn1:20
// {TEST-ConstrainedInt}
type TestConstrainedInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:10,valueUB:100"
	AttrCiA int32 `protobuf:"varint,1,opt,name=attr_ci_a,json=attrCiA,proto3" json:"attr_ci_a,omitempty" aper:"valueLB:10,valueUB:100"`
	// @inject_tag: aper:"valueLB:255,valueUB:65535"
	AttrCiB int32 `protobuf:"varint,2,opt,name=attr_ci_b,json=attrCiB,proto3" json:"attr_ci_b,omitempty" aper:"valueLB:255,valueUB:65535"`
	// @inject_tag: aper:"valueLB:10,valueUB:4294967295"
	AttrCiC int32 `protobuf:"varint,3,opt,name=attr_ci_c,json=attrCiC,proto3" json:"attr_ci_c,omitempty" aper:"valueLB:10,valueUB:4294967295"`
	// @inject_tag: aper:"valueUB:100"
	AttrCiD int32 `protobuf:"varint,4,opt,name=attr_ci_d,json=attrCiD,proto3" json:"attr_ci_d,omitempty" aper:"valueUB:100"`
	// @inject_tag: aper:"valueLB:10,valueUB:20"
	AttrCiE int32 `protobuf:"varint,5,opt,name=attr_ci_e,json=attrCiE,proto3" json:"attr_ci_e,omitempty" aper:"valueLB:10,valueUB:20"`
	// @inject_tag: aper:"valueLB:10,valueUB:10"
	AttrCiF int32 `protobuf:"varint,6,opt,name=attr_ci_f,json=attrCiF,proto3" json:"attr_ci_f,omitempty" aper:"valueLB:10,valueUB:10"`
	// @inject_tag: aper:"valueLB:10,valueUB:10,valueExt"
	AttrCiG int32 `protobuf:"varint,7,opt,name=attr_ci_g,json=attrCiG,proto3" json:"attr_ci_g,omitempty" aper:"valueLB:10,valueUB:10,valueExt"`
}

func (x *TestConstrainedInt) Reset() {
	*x = TestConstrainedInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConstrainedInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConstrainedInt) ProtoMessage() {}

func (x *TestConstrainedInt) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConstrainedInt.ProtoReflect.Descriptor instead.
func (*TestConstrainedInt) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{1}
}

func (x *TestConstrainedInt) GetAttrCiA() int32 {
	if x != nil {
		return x.AttrCiA
	}
	return 0
}

func (x *TestConstrainedInt) GetAttrCiB() int32 {
	if x != nil {
		return x.AttrCiB
	}
	return 0
}

func (x *TestConstrainedInt) GetAttrCiC() int32 {
	if x != nil {
		return x.AttrCiC
	}
	return 0
}

func (x *TestConstrainedInt) GetAttrCiD() int32 {
	if x != nil {
		return x.AttrCiD
	}
	return 0
}

func (x *TestConstrainedInt) GetAttrCiE() int32 {
	if x != nil {
		return x.AttrCiE
	}
	return 0
}

func (x *TestConstrainedInt) GetAttrCiF() int32 {
	if x != nil {
		return x.AttrCiF
	}
	return 0
}

func (x *TestConstrainedInt) GetAttrCiG() int32 {
	if x != nil {
		return x.AttrCiG
	}
	return 0
}

// sequence from tes_sm.asn1:29
// {TEST-UnconstrainedReal}
type TestUnconstrainedReal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrUcrA float64 `protobuf:"fixed64,1,opt,name=attr_ucr_a,json=attrUcrA,proto3" json:"attr_ucr_a,omitempty"`
	AttrUcrB float64 `protobuf:"fixed64,2,opt,name=attr_ucr_b,json=attrUcrB,proto3" json:"attr_ucr_b,omitempty"`
}

func (x *TestUnconstrainedReal) Reset() {
	*x = TestUnconstrainedReal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestUnconstrainedReal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestUnconstrainedReal) ProtoMessage() {}

func (x *TestUnconstrainedReal) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestUnconstrainedReal.ProtoReflect.Descriptor instead.
func (*TestUnconstrainedReal) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{2}
}

func (x *TestUnconstrainedReal) GetAttrUcrA() float64 {
	if x != nil {
		return x.AttrUcrA
	}
	return 0
}

func (x *TestUnconstrainedReal) GetAttrUcrB() float64 {
	if x != nil {
		return x.AttrUcrB
	}
	return 0
}

// sequence from tes_sm.asn1:34
// {TEST-ConstrainedReal}
type TestConstrainedReal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:10,valueUB:100"
	AttrCrA float64 `protobuf:"fixed64,1,opt,name=attr_cr_a,json=attrCrA,proto3" json:"attr_cr_a,omitempty" aper:"valueLB:10,valueUB:100"`
	// @inject_tag: aper:"valueLB:10"
	AttrCrB float64 `protobuf:"fixed64,2,opt,name=attr_cr_b,json=attrCrB,proto3" json:"attr_cr_b,omitempty" aper:"valueLB:10"`
	// @inject_tag: aper:"valueUB:100"
	AttrCrC float64 `protobuf:"fixed64,3,opt,name=attr_cr_c,json=attrCrC,proto3" json:"attr_cr_c,omitempty" aper:"valueUB:100"`
	// @inject_tag: aper:"valueLB:10,valueUB:20"
	AttrCrD float64 `protobuf:"fixed64,4,opt,name=attr_cr_d,json=attrCrD,proto3" json:"attr_cr_d,omitempty" aper:"valueLB:10,valueUB:20"`
	// @inject_tag: aper:"valueLB:10,valueUB:10"
	AttrCrE float64 `protobuf:"fixed64,5,opt,name=attr_cr_e,json=attrCrE,proto3" json:"attr_cr_e,omitempty" aper:"valueLB:10,valueUB:10"`
	// @inject_tag: aper:"valueLB:10,valueUB:10,valueExt"
	AttrCrF float64 `protobuf:"fixed64,6,opt,name=attr_cr_f,json=attrCrF,proto3" json:"attr_cr_f,omitempty" aper:"valueLB:10,valueUB:10,valueExt"`
}

func (x *TestConstrainedReal) Reset() {
	*x = TestConstrainedReal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConstrainedReal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConstrainedReal) ProtoMessage() {}

func (x *TestConstrainedReal) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConstrainedReal.ProtoReflect.Descriptor instead.
func (*TestConstrainedReal) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{3}
}

func (x *TestConstrainedReal) GetAttrCrA() float64 {
	if x != nil {
		return x.AttrCrA
	}
	return 0
}

func (x *TestConstrainedReal) GetAttrCrB() float64 {
	if x != nil {
		return x.AttrCrB
	}
	return 0
}

func (x *TestConstrainedReal) GetAttrCrC() float64 {
	if x != nil {
		return x.AttrCrC
	}
	return 0
}

func (x *TestConstrainedReal) GetAttrCrD() float64 {
	if x != nil {
		return x.AttrCrD
	}
	return 0
}

func (x *TestConstrainedReal) GetAttrCrE() float64 {
	if x != nil {
		return x.AttrCrE
	}
	return 0
}

func (x *TestConstrainedReal) GetAttrCrF() float64 {
	if x != nil {
		return x.AttrCrF
	}
	return 0
}

// sequence from tes_sm.asn1:43
// {TEST-BitString}
type TestBitString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrBs1 *asn1.BitString `protobuf:"bytes,1,opt,name=attr_bs1,json=attrBs1,proto3" json:"attr_bs1,omitempty"`
	// @inject_tag: aper:"sizeLB:20,sizeUB:20"
	AttrBs2 *asn1.BitString `protobuf:"bytes,2,opt,name=attr_bs2,json=attrBs2,proto3" json:"attr_bs2,omitempty" aper:"sizeLB:20,sizeUB:20"`
	// @inject_tag: aper:"sizeLB:20,sizeUB:20, valueExt"
	AttrBs3 *asn1.BitString `protobuf:"bytes,3,opt,name=attr_bs3,json=attrBs3,proto3" json:"attr_bs3,omitempty" aper:"sizeLB:20,sizeUB:20, valueExt"`
	// @inject_tag: aper:"sizeLB:0,sizeUB:18"
	AttrBs4 *asn1.BitString `protobuf:"bytes,4,opt,name=attr_bs4,json=attrBs4,proto3" json:"attr_bs4,omitempty" aper:"sizeLB:0,sizeUB:18"`
	// @inject_tag: aper:"sizeLB:22,sizeUB:32"
	AttrBs5 *asn1.BitString `protobuf:"bytes,5,opt,name=attr_bs5,json=attrBs5,proto3" json:"attr_bs5,omitempty" aper:"sizeLB:22,sizeUB:32"`
	// @inject_tag: aper:"sizeLB:28,sizeUB:32,sizeExt"
	AttrBs6 *asn1.BitString `protobuf:"bytes,6,opt,name=attr_bs6,json=attrBs6,proto3" json:"attr_bs6,omitempty" aper:"sizeLB:28,sizeUB:32,sizeExt"`
	// @inject_tag: aper:"sizeLB:22,sizeUB:36,optional"
	AttrBs7 *asn1.BitString `protobuf:"bytes,7,opt,name=attr_bs7,json=attrBs7,proto3,oneof" json:"attr_bs7,omitempty" aper:"sizeLB:22,sizeUB:36,optional"`
}

func (x *TestBitString) Reset() {
	*x = TestBitString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestBitString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestBitString) ProtoMessage() {}

func (x *TestBitString) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestBitString.ProtoReflect.Descriptor instead.
func (*TestBitString) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{4}
}

func (x *TestBitString) GetAttrBs1() *asn1.BitString {
	if x != nil {
		return x.AttrBs1
	}
	return nil
}

func (x *TestBitString) GetAttrBs2() *asn1.BitString {
	if x != nil {
		return x.AttrBs2
	}
	return nil
}

func (x *TestBitString) GetAttrBs3() *asn1.BitString {
	if x != nil {
		return x.AttrBs3
	}
	return nil
}

func (x *TestBitString) GetAttrBs4() *asn1.BitString {
	if x != nil {
		return x.AttrBs4
	}
	return nil
}

func (x *TestBitString) GetAttrBs5() *asn1.BitString {
	if x != nil {
		return x.AttrBs5
	}
	return nil
}

func (x *TestBitString) GetAttrBs6() *asn1.BitString {
	if x != nil {
		return x.AttrBs6
	}
	return nil
}

func (x *TestBitString) GetAttrBs7() *asn1.BitString {
	if x != nil {
		return x.AttrBs7
	}
	return nil
}

// sequence from tes_sm.asn1:53
// {TEST-Choices}
type TestChoices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherAttr string   `protobuf:"bytes,1,opt,name=other_attr,json=otherAttr,proto3" json:"other_attr,omitempty"`
	Choice1   *Choice1 `protobuf:"bytes,2,opt,name=choice1,proto3" json:"choice1,omitempty"`
	Choice2   *Choice2 `protobuf:"bytes,3,opt,name=choice2,proto3" json:"choice2,omitempty"`
	Choice3   *Choice3 `protobuf:"bytes,4,opt,name=choice3,proto3" json:"choice3,omitempty"`
	// @inject_tag: aper:"valueExt"
	Choice4 *Choice4 `protobuf:"bytes,5,opt,name=choice4,proto3" json:"choice4,omitempty" aper:"valueExt"`
}

func (x *TestChoices) Reset() {
	*x = TestChoices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestChoices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestChoices) ProtoMessage() {}

func (x *TestChoices) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestChoices.ProtoReflect.Descriptor instead.
func (*TestChoices) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{5}
}

func (x *TestChoices) GetOtherAttr() string {
	if x != nil {
		return x.OtherAttr
	}
	return ""
}

func (x *TestChoices) GetChoice1() *Choice1 {
	if x != nil {
		return x.Choice1
	}
	return nil
}

func (x *TestChoices) GetChoice2() *Choice2 {
	if x != nil {
		return x.Choice2
	}
	return nil
}

func (x *TestChoices) GetChoice3() *Choice3 {
	if x != nil {
		return x.Choice3
	}
	return nil
}

func (x *TestChoices) GetChoice4() *Choice4 {
	if x != nil {
		return x.Choice4
	}
	return nil
}

// sequence from tes_sm.asn1:62
// {Choice1}
type Choice1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:62
	//
	// Types that are assignable to Choice1:
	//	*Choice1_Choice1A
	Choice1 isChoice1_Choice1 `protobuf_oneof:"choice1"`
}

func (x *Choice1) Reset() {
	*x = Choice1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Choice1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choice1) ProtoMessage() {}

func (x *Choice1) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choice1.ProtoReflect.Descriptor instead.
func (*Choice1) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{6}
}

func (m *Choice1) GetChoice1() isChoice1_Choice1 {
	if m != nil {
		return m.Choice1
	}
	return nil
}

func (x *Choice1) GetChoice1A() int32 {
	if x, ok := x.GetChoice1().(*Choice1_Choice1A); ok {
		return x.Choice1A
	}
	return 0
}

type isChoice1_Choice1 interface {
	isChoice1_Choice1()
}

type Choice1_Choice1A struct {
	// @inject_tag: aper:"choiceIdx:1,valueExt"
	Choice1A int32 `protobuf:"varint,1,opt,name=choice1_a,json=choice1A,proto3,oneof" aper:"choiceIdx:1,valueExt"`
}

func (*Choice1_Choice1A) isChoice1_Choice1() {}

// sequence from tes_sm.asn1:65
// {Choice2}
type Choice2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:65
	//
	// Types that are assignable to Choice2:
	//	*Choice2_Choice2A
	//	*Choice2_Choice2B
	Choice2 isChoice2_Choice2 `protobuf_oneof:"choice2"`
}

func (x *Choice2) Reset() {
	*x = Choice2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Choice2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choice2) ProtoMessage() {}

func (x *Choice2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choice2.ProtoReflect.Descriptor instead.
func (*Choice2) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{7}
}

func (m *Choice2) GetChoice2() isChoice2_Choice2 {
	if m != nil {
		return m.Choice2
	}
	return nil
}

func (x *Choice2) GetChoice2A() int32 {
	if x, ok := x.GetChoice2().(*Choice2_Choice2A); ok {
		return x.Choice2A
	}
	return 0
}

func (x *Choice2) GetChoice2B() int32 {
	if x, ok := x.GetChoice2().(*Choice2_Choice2B); ok {
		return x.Choice2B
	}
	return 0
}

type isChoice2_Choice2 interface {
	isChoice2_Choice2()
}

type Choice2_Choice2A struct {
	// @inject_tag: aper:"choiceIdx:1,valueExt"
	Choice2A int32 `protobuf:"varint,1,opt,name=choice2_a,json=choice2A,proto3,oneof" aper:"choiceIdx:1,valueExt"`
}

type Choice2_Choice2B struct {
	// @inject_tag: aper:"choiceIdx:2,valueExt"
	Choice2B int32 `protobuf:"varint,2,opt,name=choice2_b,json=choice2B,proto3,oneof" aper:"choiceIdx:2,valueExt"`
}

func (*Choice2_Choice2A) isChoice2_Choice2() {}

func (*Choice2_Choice2B) isChoice2_Choice2() {}

// sequence from tes_sm.asn1:70
// {Choice3}
type Choice3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:70
	//
	// Types that are assignable to Choice3:
	//	*Choice3_Choice3A
	//	*Choice3_Choice3B
	//	*Choice3_Choice3C
	Choice3 isChoice3_Choice3 `protobuf_oneof:"choice3"`
}

func (x *Choice3) Reset() {
	*x = Choice3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Choice3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choice3) ProtoMessage() {}

func (x *Choice3) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choice3.ProtoReflect.Descriptor instead.
func (*Choice3) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{8}
}

func (m *Choice3) GetChoice3() isChoice3_Choice3 {
	if m != nil {
		return m.Choice3
	}
	return nil
}

func (x *Choice3) GetChoice3A() int32 {
	if x, ok := x.GetChoice3().(*Choice3_Choice3A); ok {
		return x.Choice3A
	}
	return 0
}

func (x *Choice3) GetChoice3B() int32 {
	if x, ok := x.GetChoice3().(*Choice3_Choice3B); ok {
		return x.Choice3B
	}
	return 0
}

func (x *Choice3) GetChoice3C() int32 {
	if x, ok := x.GetChoice3().(*Choice3_Choice3C); ok {
		return x.Choice3C
	}
	return 0
}

type isChoice3_Choice3 interface {
	isChoice3_Choice3()
}

type Choice3_Choice3A struct {
	// @inject_tag: aper:"choiceIdx:1,valueExt"
	Choice3A int32 `protobuf:"varint,1,opt,name=choice3_a,json=choice3A,proto3,oneof" aper:"choiceIdx:1,valueExt"`
}

type Choice3_Choice3B struct {
	// @inject_tag: aper:"choiceIdx:2,valueExt"
	Choice3B int32 `protobuf:"varint,2,opt,name=choice3_b,json=choice3B,proto3,oneof" aper:"choiceIdx:2,valueExt"`
}

type Choice3_Choice3C struct {
	// @inject_tag: aper:"choiceIdx:3,valueExt"
	Choice3C int32 `protobuf:"varint,3,opt,name=choice3_c,json=choice3C,proto3,oneof" aper:"choiceIdx:3,valueExt"`
}

func (*Choice3_Choice3A) isChoice3_Choice3() {}

func (*Choice3_Choice3B) isChoice3_Choice3() {}

func (*Choice3_Choice3C) isChoice3_Choice3() {}

// sequence from tes_sm.asn1:77
// {Choice4}
type Choice4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:77
	// @inject_tag: aper:"valueExt"
	//
	// Types that are assignable to Choice4:
	//	*Choice4_Choice4A
	Choice4 isChoice4_Choice4 `protobuf_oneof:"choice4" aper:"valueExt"`
}

func (x *Choice4) Reset() {
	*x = Choice4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Choice4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Choice4) ProtoMessage() {}

func (x *Choice4) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Choice4.ProtoReflect.Descriptor instead.
func (*Choice4) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{9}
}

func (m *Choice4) GetChoice4() isChoice4_Choice4 {
	if m != nil {
		return m.Choice4
	}
	return nil
}

func (x *Choice4) GetChoice4A() int32 {
	if x, ok := x.GetChoice4().(*Choice4_Choice4A); ok {
		return x.Choice4A
	}
	return 0
}

type isChoice4_Choice4 interface {
	isChoice4_Choice4()
}

type Choice4_Choice4A struct {
	// @inject_tag: aper:"choiceIdx:1,valueExt"
	Choice4A int32 `protobuf:"varint,1,opt,name=choice4_a,json=choice4A,proto3,oneof" aper:"choiceIdx:1,valueExt"`
}

func (*Choice4_Choice4A) isChoice4_Choice4() {}

// sequence from tes_sm.asn1:82
// {TEST-ConstrainedChoices}
type TestConstrainedChoices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:1,sizeUB:50,sizeExt"
	OtherCattr         string              `protobuf:"bytes,1,opt,name=other_cattr,json=otherCAttr,proto3" json:"other_cattr,omitempty" aper:"sizeLB:1,sizeUB:50,sizeExt"`
	ConstrainedChoice1 *ConstrainedChoice1 `protobuf:"bytes,2,opt,name=constrained_choice1,json=constrainedChoice1,proto3" json:"constrained_choice1,omitempty"`
	ConstrainedChoice2 *ConstrainedChoice2 `protobuf:"bytes,3,opt,name=constrained_choice2,json=constrainedChoice2,proto3" json:"constrained_choice2,omitempty"`
	ConstrainedChoice3 *ConstrainedChoice3 `protobuf:"bytes,4,opt,name=constrained_choice3,json=constrainedChoice3,proto3" json:"constrained_choice3,omitempty"`
	// @inject_tag: aper:"valueExt"
	ConstrainedChoice4 *ConstrainedChoice4 `protobuf:"bytes,5,opt,name=constrained_choice4,json=constrainedChoice4,proto3" json:"constrained_choice4,omitempty" aper:"valueExt"`
}

func (x *TestConstrainedChoices) Reset() {
	*x = TestConstrainedChoices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConstrainedChoices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConstrainedChoices) ProtoMessage() {}

func (x *TestConstrainedChoices) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConstrainedChoices.ProtoReflect.Descriptor instead.
func (*TestConstrainedChoices) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{10}
}

func (x *TestConstrainedChoices) GetOtherCattr() string {
	if x != nil {
		return x.OtherCattr
	}
	return ""
}

func (x *TestConstrainedChoices) GetConstrainedChoice1() *ConstrainedChoice1 {
	if x != nil {
		return x.ConstrainedChoice1
	}
	return nil
}

func (x *TestConstrainedChoices) GetConstrainedChoice2() *ConstrainedChoice2 {
	if x != nil {
		return x.ConstrainedChoice2
	}
	return nil
}

func (x *TestConstrainedChoices) GetConstrainedChoice3() *ConstrainedChoice3 {
	if x != nil {
		return x.ConstrainedChoice3
	}
	return nil
}

func (x *TestConstrainedChoices) GetConstrainedChoice4() *ConstrainedChoice4 {
	if x != nil {
		return x.ConstrainedChoice4
	}
	return nil
}

// sequence from tes_sm.asn1:91
// {ConstrainedChoice1}
type ConstrainedChoice1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:91
	//
	// Types that are assignable to ConstrainedChoice1:
	//	*ConstrainedChoice1_ConstrainedChoice1A
	ConstrainedChoice1 isConstrainedChoice1_ConstrainedChoice1 `protobuf_oneof:"constrained_choice1"`
}

func (x *ConstrainedChoice1) Reset() {
	*x = ConstrainedChoice1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstrainedChoice1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstrainedChoice1) ProtoMessage() {}

func (x *ConstrainedChoice1) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstrainedChoice1.ProtoReflect.Descriptor instead.
func (*ConstrainedChoice1) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{11}
}

func (m *ConstrainedChoice1) GetConstrainedChoice1() isConstrainedChoice1_ConstrainedChoice1 {
	if m != nil {
		return m.ConstrainedChoice1
	}
	return nil
}

func (x *ConstrainedChoice1) GetConstrainedChoice1A() int32 {
	if x, ok := x.GetConstrainedChoice1().(*ConstrainedChoice1_ConstrainedChoice1A); ok {
		return x.ConstrainedChoice1A
	}
	return 0
}

type isConstrainedChoice1_ConstrainedChoice1 interface {
	isConstrainedChoice1_ConstrainedChoice1()
}

type ConstrainedChoice1_ConstrainedChoice1A struct {
	// @inject_tag: aper:"choiceIdx:1,valueLB:1,valueUB:128,valueExt"
	ConstrainedChoice1A int32 `protobuf:"varint,1,opt,name=constrained_choice1_a,json=constrainedChoice1A,proto3,oneof" aper:"choiceIdx:1,valueLB:1,valueUB:128,valueExt"`
}

func (*ConstrainedChoice1_ConstrainedChoice1A) isConstrainedChoice1_ConstrainedChoice1() {}

// sequence from tes_sm.asn1:94
// {ConstrainedChoice2}
type ConstrainedChoice2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:94
	//
	// Types that are assignable to ConstrainedChoice2:
	//	*ConstrainedChoice2_ConstrainedChoice2A
	//	*ConstrainedChoice2_ConstrainedChoice2B
	ConstrainedChoice2 isConstrainedChoice2_ConstrainedChoice2 `protobuf_oneof:"constrained_choice2"`
}

func (x *ConstrainedChoice2) Reset() {
	*x = ConstrainedChoice2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstrainedChoice2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstrainedChoice2) ProtoMessage() {}

func (x *ConstrainedChoice2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstrainedChoice2.ProtoReflect.Descriptor instead.
func (*ConstrainedChoice2) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{12}
}

func (m *ConstrainedChoice2) GetConstrainedChoice2() isConstrainedChoice2_ConstrainedChoice2 {
	if m != nil {
		return m.ConstrainedChoice2
	}
	return nil
}

func (x *ConstrainedChoice2) GetConstrainedChoice2A() int32 {
	if x, ok := x.GetConstrainedChoice2().(*ConstrainedChoice2_ConstrainedChoice2A); ok {
		return x.ConstrainedChoice2A
	}
	return 0
}

func (x *ConstrainedChoice2) GetConstrainedChoice2B() int32 {
	if x, ok := x.GetConstrainedChoice2().(*ConstrainedChoice2_ConstrainedChoice2B); ok {
		return x.ConstrainedChoice2B
	}
	return 0
}

type isConstrainedChoice2_ConstrainedChoice2 interface {
	isConstrainedChoice2_ConstrainedChoice2()
}

type ConstrainedChoice2_ConstrainedChoice2A struct {
	// @inject_tag: aper:"choiceIdx:1,valueLB:0,valueUB:15,valueExt"
	ConstrainedChoice2A int32 `protobuf:"varint,1,opt,name=constrained_choice2_a,json=constrainedChoice2A,proto3,oneof" aper:"choiceIdx:1,valueLB:0,valueUB:15,valueExt"`
}

type ConstrainedChoice2_ConstrainedChoice2B struct {
	// @inject_tag: aper:"choiceIdx:2,valueLB:1,valueUB:4294967295,valueExt"
	ConstrainedChoice2B int32 `protobuf:"varint,2,opt,name=constrained_choice2_b,json=constrainedChoice2B,proto3,oneof" aper:"choiceIdx:2,valueLB:1,valueUB:4294967295,valueExt"`
}

func (*ConstrainedChoice2_ConstrainedChoice2A) isConstrainedChoice2_ConstrainedChoice2() {}

func (*ConstrainedChoice2_ConstrainedChoice2B) isConstrainedChoice2_ConstrainedChoice2() {}

// sequence from tes_sm.asn1:99
// {ConstrainedChoice3}
type ConstrainedChoice3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:99
	//
	// Types that are assignable to ConstrainedChoice3:
	//	*ConstrainedChoice3_ConstrainedChoice3A
	//	*ConstrainedChoice3_ConstrainedChoice3B
	//	*ConstrainedChoice3_ConstrainedChoice3C
	//	*ConstrainedChoice3_ConstrainedChoice3D
	ConstrainedChoice3 isConstrainedChoice3_ConstrainedChoice3 `protobuf_oneof:"constrained_choice3"`
}

func (x *ConstrainedChoice3) Reset() {
	*x = ConstrainedChoice3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstrainedChoice3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstrainedChoice3) ProtoMessage() {}

func (x *ConstrainedChoice3) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstrainedChoice3.ProtoReflect.Descriptor instead.
func (*ConstrainedChoice3) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{13}
}

func (m *ConstrainedChoice3) GetConstrainedChoice3() isConstrainedChoice3_ConstrainedChoice3 {
	if m != nil {
		return m.ConstrainedChoice3
	}
	return nil
}

func (x *ConstrainedChoice3) GetConstrainedChoice3A() int32 {
	if x, ok := x.GetConstrainedChoice3().(*ConstrainedChoice3_ConstrainedChoice3A); ok {
		return x.ConstrainedChoice3A
	}
	return 0
}

func (x *ConstrainedChoice3) GetConstrainedChoice3B() int32 {
	if x, ok := x.GetConstrainedChoice3().(*ConstrainedChoice3_ConstrainedChoice3B); ok {
		return x.ConstrainedChoice3B
	}
	return 0
}

func (x *ConstrainedChoice3) GetConstrainedChoice3C() int32 {
	if x, ok := x.GetConstrainedChoice3().(*ConstrainedChoice3_ConstrainedChoice3C); ok {
		return x.ConstrainedChoice3C
	}
	return 0
}

func (x *ConstrainedChoice3) GetConstrainedChoice3D() int32 {
	if x, ok := x.GetConstrainedChoice3().(*ConstrainedChoice3_ConstrainedChoice3D); ok {
		return x.ConstrainedChoice3D
	}
	return 0
}

type isConstrainedChoice3_ConstrainedChoice3 interface {
	isConstrainedChoice3_ConstrainedChoice3()
}

type ConstrainedChoice3_ConstrainedChoice3A struct {
	// @inject_tag: aper:"choiceIdx:1,valueExt"
	ConstrainedChoice3A int32 `protobuf:"varint,1,opt,name=constrained_choice3_a,json=constrainedChoice3A,proto3,oneof" aper:"choiceIdx:1,valueExt"`
}

type ConstrainedChoice3_ConstrainedChoice3B struct {
	// @inject_tag: aper:"choiceIdx:2,valueLB:0,valueUB:15,valueExt"
	ConstrainedChoice3B int32 `protobuf:"varint,2,opt,name=constrained_choice3_b,json=constrainedChoice3B,proto3,oneof" aper:"choiceIdx:2,valueLB:0,valueUB:15,valueExt"`
}

type ConstrainedChoice3_ConstrainedChoice3C struct {
	// @inject_tag: aper:"choiceIdx:3,valueLB:1,valueUB:4294967295,valueExt"
	ConstrainedChoice3C int32 `protobuf:"varint,3,opt,name=constrained_choice3_c,json=constrainedChoice3C,proto3,oneof" aper:"choiceIdx:3,valueLB:1,valueUB:4294967295,valueExt"`
}

type ConstrainedChoice3_ConstrainedChoice3D struct {
	// @inject_tag: aper:"choiceIdx:4,valueLB:1,valueUB:1,valueExt"
	ConstrainedChoice3D int32 `protobuf:"varint,4,opt,name=constrained_choice3_d,json=constrainedChoice3D,proto3,oneof" aper:"choiceIdx:4,valueLB:1,valueUB:1,valueExt"`
}

func (*ConstrainedChoice3_ConstrainedChoice3A) isConstrainedChoice3_ConstrainedChoice3() {}

func (*ConstrainedChoice3_ConstrainedChoice3B) isConstrainedChoice3_ConstrainedChoice3() {}

func (*ConstrainedChoice3_ConstrainedChoice3C) isConstrainedChoice3_ConstrainedChoice3() {}

func (*ConstrainedChoice3_ConstrainedChoice3D) isConstrainedChoice3_ConstrainedChoice3() {}

// sequence from tes_sm.asn1:107
// {ConstrainedChoice4}
type ConstrainedChoice4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:107
	// @inject_tag: aper:"valueExt"
	//
	// Types that are assignable to ConstrainedChoice4:
	//	*ConstrainedChoice4_ConstrainedChoice4A
	ConstrainedChoice4 isConstrainedChoice4_ConstrainedChoice4 `protobuf_oneof:"constrained_choice4" aper:"valueExt"`
}

func (x *ConstrainedChoice4) Reset() {
	*x = ConstrainedChoice4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstrainedChoice4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstrainedChoice4) ProtoMessage() {}

func (x *ConstrainedChoice4) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstrainedChoice4.ProtoReflect.Descriptor instead.
func (*ConstrainedChoice4) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{14}
}

func (m *ConstrainedChoice4) GetConstrainedChoice4() isConstrainedChoice4_ConstrainedChoice4 {
	if m != nil {
		return m.ConstrainedChoice4
	}
	return nil
}

func (x *ConstrainedChoice4) GetConstrainedChoice4A() int32 {
	if x, ok := x.GetConstrainedChoice4().(*ConstrainedChoice4_ConstrainedChoice4A); ok {
		return x.ConstrainedChoice4A
	}
	return 0
}

type isConstrainedChoice4_ConstrainedChoice4 interface {
	isConstrainedChoice4_ConstrainedChoice4()
}

type ConstrainedChoice4_ConstrainedChoice4A struct {
	// @inject_tag: aper:"choiceIdx:1,valueLB:1,valueUB:128,valueExt"
	ConstrainedChoice4A int32 `protobuf:"varint,1,opt,name=constrained_choice4_a,json=constrainedChoice4A,proto3,oneof" aper:"choiceIdx:1,valueLB:1,valueUB:128,valueExt"`
}

func (*ConstrainedChoice4_ConstrainedChoice4A) isConstrainedChoice4_ConstrainedChoice4() {}

// sequence from tes_sm.asn1:107
// {TEST-NestedChoice}
type TestNestedChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// choice from tes_sm.asn1:107
	// @inject_tag: aper:"valueExt"
	//
	// Types that are assignable to TestNestedChoice:
	//	*TestNestedChoice_Option1
	//	*TestNestedChoice_Option2
	//	*TestNestedChoice_Option3
	TestNestedChoice isTestNestedChoice_TestNestedChoice `protobuf_oneof:"test_nested_choice" aper:"valueExt"`
}

func (x *TestNestedChoice) Reset() {
	*x = TestNestedChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestNestedChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestNestedChoice) ProtoMessage() {}

func (x *TestNestedChoice) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestNestedChoice.ProtoReflect.Descriptor instead.
func (*TestNestedChoice) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{15}
}

func (m *TestNestedChoice) GetTestNestedChoice() isTestNestedChoice_TestNestedChoice {
	if m != nil {
		return m.TestNestedChoice
	}
	return nil
}

func (x *TestNestedChoice) GetOption1() *Choice3 {
	if x, ok := x.GetTestNestedChoice().(*TestNestedChoice_Option1); ok {
		return x.Option1
	}
	return nil
}

func (x *TestNestedChoice) GetOption2() *ConstrainedChoice3 {
	if x, ok := x.GetTestNestedChoice().(*TestNestedChoice_Option2); ok {
		return x.Option2
	}
	return nil
}

func (x *TestNestedChoice) GetOption3() *ConstrainedChoice4 {
	if x, ok := x.GetTestNestedChoice().(*TestNestedChoice_Option3); ok {
		return x.Option3
	}
	return nil
}

type isTestNestedChoice_TestNestedChoice interface {
	isTestNestedChoice_TestNestedChoice()
}

type TestNestedChoice_Option1 struct {
	Option1 *Choice3 `protobuf:"bytes,1,opt,name=option1,proto3,oneof"`
}

type TestNestedChoice_Option2 struct {
	Option2 *ConstrainedChoice3 `protobuf:"bytes,2,opt,name=option2,proto3,oneof"`
}

type TestNestedChoice_Option3 struct {
	// @inject_tag: aper:"valueExt"
	Option3 *ConstrainedChoice4 `protobuf:"bytes,3,opt,name=option3,proto3,oneof" aper:"valueExt"`
}

func (*TestNestedChoice_Option1) isTestNestedChoice_TestNestedChoice() {}

func (*TestNestedChoice_Option2) isTestNestedChoice_TestNestedChoice() {}

func (*TestNestedChoice_Option3) isTestNestedChoice_TestNestedChoice() {}

// sequence from tes_sm.asn1:112
// {TEST-OctetString}
type TestOctetString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrOs1 []byte `protobuf:"bytes,1,opt,name=attr_os1,json=attrOs1,proto3" json:"attr_os1,omitempty"`
	// @inject_tag: aper:"sizeLB:2,sizeUB:2"
	AttrOs2 []byte `protobuf:"bytes,2,opt,name=attr_os2,json=attrOs2,proto3" json:"attr_os2,omitempty" aper:"sizeLB:2,sizeUB:2"`
	// @inject_tag: aper:"sizeLB:2,sizeUB:2,sizeExt"
	AttrOs3 []byte `protobuf:"bytes,3,opt,name=attr_os3,json=attrOs3,proto3" json:"attr_os3,omitempty" aper:"sizeLB:2,sizeUB:2,sizeExt"`
	// @inject_tag: aper:"sizeLB:0,sizeUB:3"
	AttrOs4 []byte `protobuf:"bytes,4,opt,name=attr_os4,json=attrOs4,proto3" json:"attr_os4,omitempty" aper:"sizeLB:0,sizeUB:3"`
	// @inject_tag: aper:"sizeLB:2,sizeUB:3"
	AttrOs5 []byte `protobuf:"bytes,5,opt,name=attr_os5,json=attrOs5,proto3" json:"attr_os5,omitempty" aper:"sizeLB:2,sizeUB:3"`
	// @inject_tag: aper:"sizeLB:1,sizeUB:3,sizeExt"
	AttrOs6 []byte `protobuf:"bytes,6,opt,name=attr_os6,json=attrOs6,proto3" json:"attr_os6,omitempty" aper:"sizeLB:1,sizeUB:3,sizeExt"`
	// @inject_tag: aper:"optional,sizeLB:2,sizeUB:6"
	AttrOs7 []byte `protobuf:"bytes,7,opt,name=attr_os7,json=attrOs7,proto3,oneof" json:"attr_os7,omitempty" aper:"optional,sizeLB:2,sizeUB:6"`
}

func (x *TestOctetString) Reset() {
	*x = TestOctetString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestOctetString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestOctetString) ProtoMessage() {}

func (x *TestOctetString) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestOctetString.ProtoReflect.Descriptor instead.
func (*TestOctetString) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{16}
}

func (x *TestOctetString) GetAttrOs1() []byte {
	if x != nil {
		return x.AttrOs1
	}
	return nil
}

func (x *TestOctetString) GetAttrOs2() []byte {
	if x != nil {
		return x.AttrOs2
	}
	return nil
}

func (x *TestOctetString) GetAttrOs3() []byte {
	if x != nil {
		return x.AttrOs3
	}
	return nil
}

func (x *TestOctetString) GetAttrOs4() []byte {
	if x != nil {
		return x.AttrOs4
	}
	return nil
}

func (x *TestOctetString) GetAttrOs5() []byte {
	if x != nil {
		return x.AttrOs5
	}
	return nil
}

func (x *TestOctetString) GetAttrOs6() []byte {
	if x != nil {
		return x.AttrOs6
	}
	return nil
}

func (x *TestOctetString) GetAttrOs7() []byte {
	if x != nil {
		return x.AttrOs7
	}
	return nil
}

// sequence from tes_sm.asn1:122
// {TEST-PrintableString}
type TestPrintableString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrPs1 string `protobuf:"bytes,1,opt,name=attr_ps1,json=attrPs1,proto3" json:"attr_ps1,omitempty"`
	// @inject_tag: aper:"sizeLB:2,sizeUB:2"
	AttrPs2 string `protobuf:"bytes,2,opt,name=attr_ps2,json=attrPs2,proto3" json:"attr_ps2,omitempty" aper:"sizeLB:2,sizeUB:2"`
	// @inject_tag: aper:"sizeLB:2,sizeUB:2,sizeExt"
	AttrPs3 string `protobuf:"bytes,3,opt,name=attr_ps3,json=attrPs3,proto3" json:"attr_ps3,omitempty" aper:"sizeLB:2,sizeUB:2,sizeExt"`
	// @inject_tag: aper:"sizeLB:0,sizeUB:3"
	AttrPs4 string `protobuf:"bytes,4,opt,name=attr_ps4,json=attrPs4,proto3" json:"attr_ps4,omitempty" aper:"sizeLB:0,sizeUB:3"`
	// @inject_tag: aper:"sizeLB:2,sizeUB:3"
	AttrPs5 string `protobuf:"bytes,5,opt,name=attr_ps5,json=attrPs5,proto3" json:"attr_ps5,omitempty" aper:"sizeLB:2,sizeUB:3"`
	// @inject_tag: aper:"sizeLB:1,sizeUB:3,sizeExt"
	AttrPs6 string `protobuf:"bytes,6,opt,name=attr_ps6,json=attrPs6,proto3" json:"attr_ps6,omitempty" aper:"sizeLB:1,sizeUB:3,sizeExt"`
	// @inject_tag: aper:"optional,sizeLB:2,sizeUB:6"
	AttrPs7 *string `protobuf:"bytes,7,opt,name=attr_ps7,json=attrPs7,proto3,oneof" json:"attr_ps7,omitempty" aper:"optional,sizeLB:2,sizeUB:6"`
}

func (x *TestPrintableString) Reset() {
	*x = TestPrintableString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPrintableString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPrintableString) ProtoMessage() {}

func (x *TestPrintableString) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPrintableString.ProtoReflect.Descriptor instead.
func (*TestPrintableString) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{17}
}

func (x *TestPrintableString) GetAttrPs1() string {
	if x != nil {
		return x.AttrPs1
	}
	return ""
}

func (x *TestPrintableString) GetAttrPs2() string {
	if x != nil {
		return x.AttrPs2
	}
	return ""
}

func (x *TestPrintableString) GetAttrPs3() string {
	if x != nil {
		return x.AttrPs3
	}
	return ""
}

func (x *TestPrintableString) GetAttrPs4() string {
	if x != nil {
		return x.AttrPs4
	}
	return ""
}

func (x *TestPrintableString) GetAttrPs5() string {
	if x != nil {
		return x.AttrPs5
	}
	return ""
}

func (x *TestPrintableString) GetAttrPs6() string {
	if x != nil {
		return x.AttrPs6
	}
	return ""
}

func (x *TestPrintableString) GetAttrPs7() string {
	if x != nil && x.AttrPs7 != nil {
		return *x.AttrPs7
	}
	return ""
}

// sequence from tes_sm.asn1:133
// {TEST-List1}
type TestList1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:0,sizeUB:12,valueExt"
	Value []*Item `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" aper:"sizeLB:0,sizeUB:12,valueExt"`
}

func (x *TestList1) Reset() {
	*x = TestList1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestList1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestList1) ProtoMessage() {}

func (x *TestList1) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestList1.ProtoReflect.Descriptor instead.
func (*TestList1) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{18}
}

func (x *TestList1) GetValue() []*Item {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from tes_sm.asn1:134
// {Item}
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"optional"
	Item1 *int32 `protobuf:"varint,1,opt,name=item1,proto3,oneof" json:"item1,omitempty" aper:"optional"`
	// @inject_tag: aper:"sizeLB:3,sizeUB:7"
	Item2 *asn1.BitString `protobuf:"bytes,2,opt,name=item2,proto3" json:"item2,omitempty" aper:"sizeLB:3,sizeUB:7"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{19}
}

func (x *Item) GetItem1() int32 {
	if x != nil && x.Item1 != nil {
		return *x.Item1
	}
	return 0
}

func (x *Item) GetItem2() *asn1.BitString {
	if x != nil {
		return x.Item2
	}
	return nil
}

// sequence from tes_sm.asn1:140
// {TEST-List2}
type TestList2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:0,sizeUB:12,valueExt"
	Value []*ItemExtensible `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" aper:"sizeLB:0,sizeUB:12,valueExt"`
}

func (x *TestList2) Reset() {
	*x = TestList2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestList2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestList2) ProtoMessage() {}

func (x *TestList2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestList2.ProtoReflect.Descriptor instead.
func (*TestList2) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{20}
}

func (x *TestList2) GetValue() []*ItemExtensible {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from tes_sm.asn1:141
// {ItemExtensible}
type ItemExtensible struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item1 int32 `protobuf:"varint,1,opt,name=item1,proto3" json:"item1,omitempty"`
	// @inject_tag: aper:"optional,sizeLB:3,sizeUB:7"
	Item2 []byte `protobuf:"bytes,2,opt,name=item2,proto3,oneof" json:"item2,omitempty" aper:"optional,sizeLB:3,sizeUB:7"`
}

func (x *ItemExtensible) Reset() {
	*x = ItemExtensible{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemExtensible) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemExtensible) ProtoMessage() {}

func (x *ItemExtensible) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemExtensible.ProtoReflect.Descriptor instead.
func (*ItemExtensible) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{21}
}

func (x *ItemExtensible) GetItem1() int32 {
	if x != nil {
		return x.Item1
	}
	return 0
}

func (x *ItemExtensible) GetItem2() []byte {
	if x != nil {
		return x.Item2
	}
	return nil
}

// sequence from tes_sm.asn1:147
// {TEST-FullyOptionalSequence}
type TestFullyOptionalSequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"optional"
	Item1 *int32 `protobuf:"varint,1,opt,name=item1,proto3,oneof" json:"item1,omitempty" aper:"optional"`
	// @inject_tag: aper:"optional,sizeLB:3,sizeUB:7"
	Item2 []byte `protobuf:"bytes,2,opt,name=item2,proto3,oneof" json:"item2,omitempty" aper:"optional,sizeLB:3,sizeUB:7"`
	// @inject_tag: aper:"optional"
	Item3 *bool `protobuf:"varint,3,opt,name=item3,proto3,oneof" json:"item3,omitempty" aper:"optional"`
	// @inject_tag: aper:"optional,valueLB:0,valueUB:1,valueExt"
	Item4 *TestFullyOptionalSequenceItem4 `protobuf:"varint,4,opt,name=item4,proto3,enum=aper.test.v1.TestFullyOptionalSequenceItem4,oneof" json:"item4,omitempty" aper:"optional,valueLB:0,valueUB:1,valueExt"`
	// @inject_tag: aper:"optional,valueLB:0,valueUB:0"
	Item5 *int32 `protobuf:"varint,5,opt,name=item5,proto3,oneof" json:"item5,omitempty" aper:"optional,valueLB:0,valueUB:0"`
}

func (x *TestFullyOptionalSequence) Reset() {
	*x = TestFullyOptionalSequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestFullyOptionalSequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestFullyOptionalSequence) ProtoMessage() {}

func (x *TestFullyOptionalSequence) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestFullyOptionalSequence.ProtoReflect.Descriptor instead.
func (*TestFullyOptionalSequence) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{22}
}

func (x *TestFullyOptionalSequence) GetItem1() int32 {
	if x != nil && x.Item1 != nil {
		return *x.Item1
	}
	return 0
}

func (x *TestFullyOptionalSequence) GetItem2() []byte {
	if x != nil {
		return x.Item2
	}
	return nil
}

func (x *TestFullyOptionalSequence) GetItem3() bool {
	if x != nil && x.Item3 != nil {
		return *x.Item3
	}
	return false
}

func (x *TestFullyOptionalSequence) GetItem4() TestFullyOptionalSequenceItem4 {
	if x != nil && x.Item4 != nil {
		return *x.Item4
	}
	return TestFullyOptionalSequenceItem4_TEST_FULLY_OPTIONAL_SEQUENCE_ITEM4_ONE
}

func (x *TestFullyOptionalSequence) GetItem5() int32 {
	if x != nil && x.Item5 != nil {
		return *x.Item5
	}
	return 0
}

// sequence from tes_sm.asn1:156
// {TEST-ListExtensible1}
type TestListExtensible1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:0,sizeUB:4,sizeExt"
	Value []*Item `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" aper:"sizeLB:0,sizeUB:4,sizeExt"`
}

func (x *TestListExtensible1) Reset() {
	*x = TestListExtensible1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestListExtensible1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestListExtensible1) ProtoMessage() {}

func (x *TestListExtensible1) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestListExtensible1.ProtoReflect.Descriptor instead.
func (*TestListExtensible1) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{23}
}

func (x *TestListExtensible1) GetValue() []*Item {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from tes_sm.asn1:157
// {TEST-ListExtensible2}
type TestListExtensible2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:0,sizeUB:4,sizeExt,valueExt"
	Value []*ItemExtensible `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" aper:"sizeLB:0,sizeUB:4,sizeExt,valueExt"`
}

func (x *TestListExtensible2) Reset() {
	*x = TestListExtensible2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestListExtensible2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestListExtensible2) ProtoMessage() {}

func (x *TestListExtensible2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestListExtensible2.ProtoReflect.Descriptor instead.
func (*TestListExtensible2) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{24}
}

func (x *TestListExtensible2) GetValue() []*ItemExtensible {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from tes_sm.asn1:158
// {TEST-ListExtensible3}
type TestListExtensible3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:0,sizeUB:4,sizeExt,valueExt"
	Value []*TestFullyOptionalSequence `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" aper:"sizeLB:0,sizeUB:4,sizeExt,valueExt"`
}

func (x *TestListExtensible3) Reset() {
	*x = TestListExtensible3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestListExtensible3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestListExtensible3) ProtoMessage() {}

func (x *TestListExtensible3) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestListExtensible3.ProtoReflect.Descriptor instead.
func (*TestListExtensible3) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{25}
}

func (x *TestListExtensible3) GetValue() []*TestFullyOptionalSequence {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from tes_sm.asn1:160
// {TEST-List3}
type TestList3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:0,sizeUB:12,valueExt"
	Value []*TestFullyOptionalSequence `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" aper:"sizeLB:0,sizeUB:12,valueExt"`
}

func (x *TestList3) Reset() {
	*x = TestList3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestList3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestList3) ProtoMessage() {}

func (x *TestList3) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestList3.ProtoReflect.Descriptor instead.
func (*TestList3) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{26}
}

func (x *TestList3) GetValue() []*TestFullyOptionalSequence {
	if x != nil {
		return x.Value
	}
	return nil
}

// sequence from test_sm.asn1:193
// {TEST-TopLevelPDU}
type TestTopLevelPdu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opt1 *TestUnconstrainedInt `protobuf:"bytes,1,opt,name=opt1,proto3" json:"opt1,omitempty"`
	// @inject_tag: aper:"optional"
	Opt2 *TestConstrainedReal `protobuf:"bytes,2,opt,name=opt2,proto3,oneof" json:"opt2,omitempty" aper:"optional"`
	// @inject_tag: aper:"valueExt"
	Opt3 *TestNestedChoice `protobuf:"bytes,3,opt,name=opt3,proto3" json:"opt3,omitempty" aper:"valueExt"`
	Opt4 *TestBitString    `protobuf:"bytes,4,opt,name=opt4,proto3" json:"opt4,omitempty"`
	// @inject_tag: aper:"optional"
	Opt5 *TestOctetString `protobuf:"bytes,5,opt,name=opt5,proto3,oneof" json:"opt5,omitempty" aper:"optional"`
	// @inject_tag: aper:""
	Opt6 *TestListExtensible3 `protobuf:"bytes,6,opt,name=opt6,proto3" json:"opt6,omitempty"`
	// @inject_tag: aper:"valueExt"
	Opt7 TestEnumeratedExtensible `protobuf:"varint,7,opt,name=opt7,proto3,enum=aper.test.v1.TestEnumeratedExtensible" json:"opt7,omitempty" aper:"valueExt"`
}

func (x *TestTopLevelPdu) Reset() {
	*x = TestTopLevelPdu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestTopLevelPdu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestTopLevelPdu) ProtoMessage() {}

func (x *TestTopLevelPdu) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestTopLevelPdu.ProtoReflect.Descriptor instead.
func (*TestTopLevelPdu) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{27}
}

func (x *TestTopLevelPdu) GetOpt1() *TestUnconstrainedInt {
	if x != nil {
		return x.Opt1
	}
	return nil
}

func (x *TestTopLevelPdu) GetOpt2() *TestConstrainedReal {
	if x != nil {
		return x.Opt2
	}
	return nil
}

func (x *TestTopLevelPdu) GetOpt3() *TestNestedChoice {
	if x != nil {
		return x.Opt3
	}
	return nil
}

func (x *TestTopLevelPdu) GetOpt4() *TestBitString {
	if x != nil {
		return x.Opt4
	}
	return nil
}

func (x *TestTopLevelPdu) GetOpt5() *TestOctetString {
	if x != nil {
		return x.Opt5
	}
	return nil
}

func (x *TestTopLevelPdu) GetOpt6() *TestListExtensible3 {
	if x != nil {
		return x.Opt6
	}
	return nil
}

func (x *TestTopLevelPdu) GetOpt7() TestEnumeratedExtensible {
	if x != nil {
		return x.Opt7
	}
	return TestEnumeratedExtensible_TEST_ENUMERATED_EXTENSIBLE_ENUM1
}

type SampleE2ApPduChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: aper:"valueLB:0,valueUB:65535,unique"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" aper:"valueLB:0,valueUB:65535,unique"`
	//@inject_tag: aper:"valueLB:0,valueUB:2"
	Criticality int32 `protobuf:"varint,2,opt,name=criticality,proto3" json:"criticality,omitempty" aper:"valueLB:0,valueUB:2"`
	//@inject_tag: aper:"canonicalOrder"
	Ch *CanonicalChoice `protobuf:"bytes,3,opt,name=ch,proto3" json:"ch,omitempty" aper:"canonicalOrder"`
}

func (x *SampleE2ApPduChoice) Reset() {
	*x = SampleE2ApPduChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[28]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleE2ApPduChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleE2ApPduChoice) ProtoMessage() {}

func (x *SampleE2ApPduChoice) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[28]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleE2ApPduChoice.ProtoReflect.Descriptor instead.
func (*SampleE2ApPduChoice) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{28}
}

func (x *SampleE2ApPduChoice) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SampleE2ApPduChoice) GetCriticality() int32 {
	if x != nil {
		return x.Criticality
	}
	return 0
}

func (x *SampleE2ApPduChoice) GetCh() *CanonicalChoice {
	if x != nil {
		return x.Ch
	}
	return nil
}

type CanonicalChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CanonicalChoice:
	//	*CanonicalChoice_Ch1
	//	*CanonicalChoice_Ch2
	//	*CanonicalChoice_Ch3
	//	*CanonicalChoice_Ch4
	//	*CanonicalChoice_Ch5
	CanonicalChoice isCanonicalChoice_CanonicalChoice `protobuf_oneof:"canonical_choice"`
}

func (x *CanonicalChoice) Reset() {
	*x = CanonicalChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[29]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanonicalChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanonicalChoice) ProtoMessage() {}

func (x *CanonicalChoice) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[29]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanonicalChoice.ProtoReflect.Descriptor instead.
func (*CanonicalChoice) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{29}
}

func (m *CanonicalChoice) GetCanonicalChoice() isCanonicalChoice_CanonicalChoice {
	if m != nil {
		return m.CanonicalChoice
	}
	return nil
}

func (x *CanonicalChoice) GetCh1() *SampleOctetString {
	if x, ok := x.GetCanonicalChoice().(*CanonicalChoice_Ch1); ok {
		return x.Ch1
	}
	return nil
}

func (x *CanonicalChoice) GetCh2() *SampleConstrainedInteger {
	if x, ok := x.GetCanonicalChoice().(*CanonicalChoice_Ch2); ok {
		return x.Ch2
	}
	return nil
}

func (x *CanonicalChoice) GetCh3() *SampleBitString {
	if x, ok := x.GetCanonicalChoice().(*CanonicalChoice_Ch3); ok {
		return x.Ch3
	}
	return nil
}

func (x *CanonicalChoice) GetCh4() *TestListExtensible1 {
	if x, ok := x.GetCanonicalChoice().(*CanonicalChoice_Ch4); ok {
		return x.Ch4
	}
	return nil
}

func (x *CanonicalChoice) GetCh5() *Item {
	if x, ok := x.GetCanonicalChoice().(*CanonicalChoice_Ch5); ok {
		return x.Ch5
	}
	return nil
}

type isCanonicalChoice_CanonicalChoice interface {
	isCanonicalChoice_CanonicalChoice()
}

type CanonicalChoice_Ch1 struct {
	//@inject_tag: aper:"choiceIdx:1"
	Ch1 *SampleOctetString `protobuf:"bytes,1,opt,name=ch1,proto3,oneof" aper:"choiceIdx:1"`
}

type CanonicalChoice_Ch2 struct {
	//@inject_tag: aper:"choiceIdx:2"
	Ch2 *SampleConstrainedInteger `protobuf:"bytes,2,opt,name=ch2,proto3,oneof" aper:"choiceIdx:2"`
}

type CanonicalChoice_Ch3 struct {
	//@inject_tag: aper:"choiceIdx:3"
	Ch3 *SampleBitString `protobuf:"bytes,3,opt,name=ch3,proto3,oneof" aper:"choiceIdx:3"`
}

type CanonicalChoice_Ch4 struct {
	//@inject_tag: aper:"choiceIdx:4"
	Ch4 *TestListExtensible1 `protobuf:"bytes,4,opt,name=ch4,proto3,oneof" aper:"choiceIdx:4"`
}

type CanonicalChoice_Ch5 struct {
	//@inject_tag: aper:"choiceIdx:5,valueExt"
	Ch5 *Item `protobuf:"bytes,5,opt,name=ch5,proto3,oneof" aper:"choiceIdx:5,valueExt"`
}

func (*CanonicalChoice_Ch1) isCanonicalChoice_CanonicalChoice() {}

func (*CanonicalChoice_Ch2) isCanonicalChoice_CanonicalChoice() {}

func (*CanonicalChoice_Ch3) isCanonicalChoice_CanonicalChoice() {}

func (*CanonicalChoice_Ch4) isCanonicalChoice_CanonicalChoice() {}

func (*CanonicalChoice_Ch5) isCanonicalChoice_CanonicalChoice() {}

type SampleOctetString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SampleOctetString) Reset() {
	*x = SampleOctetString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[30]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleOctetString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleOctetString) ProtoMessage() {}

func (x *SampleOctetString) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[30]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleOctetString.ProtoReflect.Descriptor instead.
func (*SampleOctetString) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{30}
}

func (x *SampleOctetString) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SampleConstrainedInteger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"valueLB:0,valueUB:255,valueExt"
	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty" aper:"valueLB:0,valueUB:255,valueExt"`
}

func (x *SampleConstrainedInteger) Reset() {
	*x = SampleConstrainedInteger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[31]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleConstrainedInteger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleConstrainedInteger) ProtoMessage() {}

func (x *SampleConstrainedInteger) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[31]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleConstrainedInteger.ProtoReflect.Descriptor instead.
func (*SampleConstrainedInteger) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{31}
}

func (x *SampleConstrainedInteger) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SampleBitString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: aper:"sizeLB:24,sizeUB:50,sizeExt"
	Value *asn1.BitString `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" aper:"sizeLB:24,sizeUB:50,sizeExt"`
}

func (x *SampleBitString) Reset() {
	*x = SampleBitString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[32]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleBitString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleBitString) ProtoMessage() {}

func (x *SampleBitString) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_asn1_testsm_test_sm_proto_msgTypes[32]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleBitString.ProtoReflect.Descriptor instead.
func (*SampleBitString) Descriptor() ([]byte, []int) {
	return file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP(), []int{32}
}

func (x *SampleBitString) GetValue() *asn1.BitString {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_pkg_asn1_testsm_test_sm_proto protoreflect.FileDescriptor

var file_pkg_asn1_testsm_test_sm_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x73, 0x6e, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x61,
	0x73, 0x6e, 0x31, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x52, 0x0a, 0x14, 0x54, 0x65, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x5f, 0x75, 0x63, 0x69, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61,
	0x74, 0x74, 0x72, 0x55, 0x63, 0x69, 0x41, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x5f,
	0x75, 0x63, 0x69, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x74, 0x74,
	0x72, 0x55, 0x63, 0x69, 0x42, 0x22, 0xd8, 0x01, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x09,
	0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x69, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x69, 0x41, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x63, 0x69, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x74, 0x74,
	0x72, 0x43, 0x69, 0x42, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x69, 0x5f,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x69, 0x43,
	0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x69, 0x5f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x69, 0x44, 0x12, 0x1a, 0x0a, 0x09,
	0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x69, 0x5f, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x69, 0x45, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x63, 0x69, 0x5f, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x74, 0x74,
	0x72, 0x43, 0x69, 0x46, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x69, 0x5f,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x69, 0x47,
	0x22, 0x53, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x5f, 0x75, 0x63, 0x72, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61,
	0x74, 0x74, 0x72, 0x55, 0x63, 0x72, 0x41, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x5f,
	0x75, 0x63, 0x72, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x74, 0x74,
	0x72, 0x55, 0x63, 0x72, 0x42, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x61, 0x6c, 0x12, 0x1a, 0x0a,
	0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x72, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x72, 0x41, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74,
	0x72, 0x5f, 0x63, 0x72, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x74,
	0x74, 0x72, 0x43, 0x72, 0x42, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x72,
	0x5f, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x72,
	0x43, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x72, 0x5f, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x72, 0x44, 0x12, 0x1a, 0x0a,
	0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x63, 0x72, 0x5f, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x43, 0x72, 0x45, 0x12, 0x1a, 0x0a, 0x09, 0x61, 0x74, 0x74,
	0x72, 0x5f, 0x63, 0x72, 0x5f, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x74,
	0x74, 0x72, 0x43, 0x72, 0x46, 0x22, 0xea, 0x02, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x42, 0x69,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f,
	0x62, 0x73, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61,
	0x74, 0x74, 0x72, 0x42, 0x73, 0x31, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x62,
	0x73, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x74,
	0x74, 0x72, 0x42, 0x73, 0x32, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x62, 0x73,
	0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x74, 0x74,
	0x72, 0x42, 0x73, 0x33, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x62, 0x73, 0x34,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72,
	0x42, 0x73, 0x34, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x62, 0x73, 0x35, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x42,
	0x73, 0x35, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x62, 0x73, 0x36, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x42, 0x73,
	0x36, 0x12, 0x32, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x62, 0x73, 0x37, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x42,
	0x73, 0x37, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x62,
	0x73, 0x37, 0x22, 0xf0, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x41, 0x74, 0x74,
	0x72, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x31, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x32, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x52, 0x07, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x33, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x52, 0x07, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x34, 0x22, 0x33, 0x0a, 0x07, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31,
	0x12, 0x1d, 0x0a, 0x09, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x5f, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x41, 0x42,
	0x09, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x22, 0x52, 0x0a, 0x07, 0x43, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x32, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32,
	0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x32, 0x41, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x5f,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x32, 0x42, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x22, 0x71,
	0x0a, 0x07, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x33, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08,
	0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x41, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x33, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x42, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x33, 0x5f, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x33, 0x43, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x33, 0x22, 0x33, 0x0a, 0x07, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x12, 0x1d, 0x0a, 0x09,
	0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x08, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x41, 0x42, 0x09, 0x0a, 0x07, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x22, 0x85, 0x03, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x43, 0x41, 0x74,
	0x74, 0x72, 0x12, 0x51, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x31, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x31, 0x12, 0x51, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x32, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x12, 0x51, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x12, 0x51, 0x0a, 0x13, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x22, 0x61,
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x31, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x5f, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x31, 0x41, 0x42, 0x15, 0x0a, 0x13, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x31, 0x22, 0x97, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x5f,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x41, 0x12, 0x34,
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x32, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x32, 0x42, 0x42, 0x15, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x32, 0x22, 0x83, 0x02, 0x0a, 0x12,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x33, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65,
	0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x41, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x5f,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x42, 0x12, 0x34,
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x33, 0x5f, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x33, 0x43, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x5f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x44, 0x42, 0x15, 0x0a, 0x13, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x33, 0x22, 0x61, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x5f, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x41, 0x42, 0x15, 0x0a,
	0x13, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x34, 0x22, 0xd7, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x33, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x31, 0x12, 0x3c, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x33, 0x48,
	0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x12, 0x3c, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70,
	0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x34, 0x48, 0x00, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x42, 0x14, 0x0a, 0x12, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x22, 0xe0,
	0x01, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x63, 0x74, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6f, 0x73, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x4f, 0x73, 0x31, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6f, 0x73, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x61, 0x74, 0x74, 0x72, 0x4f, 0x73, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x6f, 0x73, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72,
	0x4f, 0x73, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6f, 0x73, 0x34, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x4f, 0x73, 0x34, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6f, 0x73, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x4f, 0x73, 0x35, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74,
	0x72, 0x5f, 0x6f, 0x73, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x74, 0x74,
	0x72, 0x4f, 0x73, 0x36, 0x12, 0x1e, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6f, 0x73, 0x37,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x4f, 0x73,
	0x37, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6f, 0x73,
	0x37, 0x22, 0xe4, 0x01, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74,
	0x72, 0x5f, 0x70, 0x73, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x74,
	0x72, 0x50, 0x73, 0x31, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x70, 0x73, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x50, 0x73, 0x32, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x70, 0x73, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x50, 0x73, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74,
	0x74, 0x72, 0x5f, 0x70, 0x73, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74,
	0x74, 0x72, 0x50, 0x73, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x70, 0x73,
	0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x50, 0x73, 0x35,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x70, 0x73, 0x36, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x74, 0x72, 0x50, 0x73, 0x36, 0x12, 0x1e, 0x0a, 0x08, 0x61,
	0x74, 0x74, 0x72, 0x5f, 0x70, 0x73, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x61, 0x74, 0x74, 0x72, 0x50, 0x73, 0x37, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x5f, 0x70, 0x73, 0x37, 0x22, 0x35, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x31, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x55, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x31, 0x88,
	0x01, 0x01, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x74, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x32, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x31, 0x22, 0x3f, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4b, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x31, 0x12,
	0x19, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x32, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x32, 0x22, 0x82, 0x02, 0x0a, 0x19, 0x54, 0x65, 0x73, 0x74, 0x46, 0x75, 0x6c,
	0x6c, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x31, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x32, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x33,
	0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x34, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x34,
	0x48, 0x03, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x34, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x35, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x31, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x32, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x33, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x34, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x35, 0x22, 0x3f, 0x0a, 0x13, 0x54, 0x65, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x31,
	0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x13, 0x54, 0x65,
	0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x32, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x54, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x33, 0x12, 0x3d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70,
	0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4a, 0x0a, 0x09, 0x54,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x33, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa7, 0x03, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74,
	0x54, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x64, 0x75, 0x12, 0x36, 0x0a, 0x04, 0x6f,
	0x70, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x65, 0x72,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x55, 0x6e, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x52, 0x04, 0x6f,
	0x70, 0x74, 0x31, 0x12, 0x3a, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x52, 0x65, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x32, 0x88, 0x01, 0x01, 0x12,
	0x32, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x04, 0x6f,
	0x70, 0x74, 0x33, 0x12, 0x2f, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x42, 0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04,
	0x6f, 0x70, 0x74, 0x34, 0x12, 0x36, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x35, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x63, 0x74, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x48, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x35, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x04,
	0x6f, 0x70, 0x74, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x33, 0x52, 0x04, 0x6f,
	0x70, 0x74, 0x36, 0x12, 0x3a, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x37, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6f, 0x70, 0x74, 0x32, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6f, 0x70, 0x74,
	0x35, 0x22, 0x76, 0x0a, 0x13, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x32, 0x61, 0x70, 0x50,
	0x64, 0x75, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x02, 0x63, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x43,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x02, 0x63, 0x68, 0x22, 0xa8, 0x02, 0x0a, 0x0f, 0x43, 0x61,
	0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a,
	0x03, 0x63, 0x68, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x4f, 0x63, 0x74, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x03, 0x63,
	0x68, 0x31, 0x12, 0x3a, 0x0a, 0x03, 0x63, 0x68, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x03, 0x63, 0x68, 0x32, 0x12, 0x31,
	0x0a, 0x03, 0x63, 0x68, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70,
	0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x42, 0x69, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x03, 0x63, 0x68,
	0x33, 0x12, 0x35, 0x0a, 0x03, 0x63, 0x68, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x31, 0x48, 0x00, 0x52, 0x03, 0x63, 0x68, 0x34, 0x12, 0x26, 0x0a, 0x03, 0x63, 0x68, 0x35, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x03, 0x63, 0x68, 0x35,
	0x42, 0x12, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x63,
	0x74, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x30, 0x0a, 0x18, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x3b, 0x0a, 0x0f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x69, 0x74, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x73, 0x6e, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0xb2,
	0x01, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52,
	0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x31, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x45, 0x4e, 0x55, 0x4d, 0x32, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x53, 0x54, 0x5f,
	0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x33,
	0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45,
	0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x34, 0x10, 0x03, 0x12, 0x19, 0x0a,
	0x15, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x35, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x53, 0x54,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x4e, 0x55, 0x4d,
	0x36, 0x10, 0x05, 0x2a, 0xfe, 0x01, 0x0a, 0x18, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x12, 0x24, 0x0a, 0x20, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41,
	0x54, 0x45, 0x44, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x45,
	0x4e, 0x55, 0x4d, 0x31, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45,
	0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53,
	0x49, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x32, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f,
	0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x33,
	0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45,
	0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x42, 0x4c, 0x45,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x34, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x45, 0x53, 0x54,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x4e, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x35, 0x10, 0x04, 0x12, 0x24,
	0x0a, 0x20, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x45, 0x52, 0x41, 0x54, 0x45,
	0x44, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x55,
	0x4d, 0x36, 0x10, 0x05, 0x2a, 0x78, 0x0a, 0x1e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x34, 0x12, 0x2a, 0x0a, 0x26, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x46,
	0x55, 0x4c, 0x4c, 0x59, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45,
	0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x34, 0x5f, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x2a, 0x0a, 0x26, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x59,
	0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e,
	0x43, 0x45, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x34, 0x5f, 0x54, 0x57, 0x4f, 0x10, 0x01, 0x42, 0x11,
	0x5a, 0x0f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x73, 0x6e, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_asn1_testsm_test_sm_proto_rawDescOnce sync.Once
	file_pkg_asn1_testsm_test_sm_proto_rawDescData = file_pkg_asn1_testsm_test_sm_proto_rawDesc
)

func file_pkg_asn1_testsm_test_sm_proto_rawDescGZIP() []byte {
	file_pkg_asn1_testsm_test_sm_proto_rawDescOnce.Do(func() {
		file_pkg_asn1_testsm_test_sm_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_asn1_testsm_test_sm_proto_rawDescData)
	})
	return file_pkg_asn1_testsm_test_sm_proto_rawDescData
}

var file_pkg_asn1_testsm_test_sm_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pkg_asn1_testsm_test_sm_proto_msgTypes = make([]protoimpl.MessageInfo, 33)
var file_pkg_asn1_testsm_test_sm_proto_goTypes = []interface{}{
	(TestEnumerated)(0),                 // 0: aper.test.v1.TestEnumerated
	(TestEnumeratedExtensible)(0),       // 1: aper.test.v1.TestEnumeratedExtensible
	(TestFullyOptionalSequenceItem4)(0), // 2: aper.test.v1.TestFullyOptionalSequenceItem4
	(*TestUnconstrainedInt)(nil),        // 3: aper.test.v1.TestUnconstrainedInt
	(*TestConstrainedInt)(nil),          // 4: aper.test.v1.TestConstrainedInt
	(*TestUnconstrainedReal)(nil),       // 5: aper.test.v1.TestUnconstrainedReal
	(*TestConstrainedReal)(nil),         // 6: aper.test.v1.TestConstrainedReal
	(*TestBitString)(nil),               // 7: aper.test.v1.TestBitString
	(*TestChoices)(nil),                 // 8: aper.test.v1.TestChoices
	(*Choice1)(nil),                     // 9: aper.test.v1.Choice1
	(*Choice2)(nil),                     // 10: aper.test.v1.Choice2
	(*Choice3)(nil),                     // 11: aper.test.v1.Choice3
	(*Choice4)(nil),                     // 12: aper.test.v1.Choice4
	(*TestConstrainedChoices)(nil),      // 13: aper.test.v1.TestConstrainedChoices
	(*ConstrainedChoice1)(nil),          // 14: aper.test.v1.ConstrainedChoice1
	(*ConstrainedChoice2)(nil),          // 15: aper.test.v1.ConstrainedChoice2
	(*ConstrainedChoice3)(nil),          // 16: aper.test.v1.ConstrainedChoice3
	(*ConstrainedChoice4)(nil),          // 17: aper.test.v1.ConstrainedChoice4
	(*TestNestedChoice)(nil),            // 18: aper.test.v1.TestNestedChoice
	(*TestOctetString)(nil),             // 19: aper.test.v1.TestOctetString
	(*TestPrintableString)(nil),         // 20: aper.test.v1.TestPrintableString
	(*TestList1)(nil),                   // 21: aper.test.v1.TestList1
	(*Item)(nil),                        // 22: aper.test.v1.Item
	(*TestList2)(nil),                   // 23: aper.test.v1.TestList2
	(*ItemExtensible)(nil),              // 24: aper.test.v1.ItemExtensible
	(*TestFullyOptionalSequence)(nil),   // 25: aper.test.v1.TestFullyOptionalSequence
	(*TestListExtensible1)(nil),         // 26: aper.test.v1.TestListExtensible1
	(*TestListExtensible2)(nil),         // 27: aper.test.v1.TestListExtensible2
	(*TestListExtensible3)(nil),         // 28: aper.test.v1.TestListExtensible3
	(*TestList3)(nil),                   // 29: aper.test.v1.TestList3
	(*TestTopLevelPdu)(nil),             // 30: aper.test.v1.TestTopLevelPdu
	(*SampleE2ApPduChoice)(nil),         // 31: aper.test.v1.sampleE2apPduChoice
	(*CanonicalChoice)(nil),             // 32: aper.test.v1.CanonicalChoice
	(*SampleOctetString)(nil),           // 33: aper.test.v1.SampleOctetString
	(*SampleConstrainedInteger)(nil),    // 34: aper.test.v1.SampleConstrainedInteger
	(*SampleBitString)(nil),             // 35: aper.test.v1.SampleBitString
	(*asn1.BitString)(nil),              // 36: asn1.v1.BitString
}
var file_pkg_asn1_testsm_test_sm_proto_depIdxs = []int32{
	36, // 0: aper.test.v1.TestBitString.attr_bs1:type_name -> asn1.v1.BitString
	36, // 1: aper.test.v1.TestBitString.attr_bs2:type_name -> asn1.v1.BitString
	36, // 2: aper.test.v1.TestBitString.attr_bs3:type_name -> asn1.v1.BitString
	36, // 3: aper.test.v1.TestBitString.attr_bs4:type_name -> asn1.v1.BitString
	36, // 4: aper.test.v1.TestBitString.attr_bs5:type_name -> asn1.v1.BitString
	36, // 5: aper.test.v1.TestBitString.attr_bs6:type_name -> asn1.v1.BitString
	36, // 6: aper.test.v1.TestBitString.attr_bs7:type_name -> asn1.v1.BitString
	9,  // 7: aper.test.v1.TestChoices.choice1:type_name -> aper.test.v1.Choice1
	10, // 8: aper.test.v1.TestChoices.choice2:type_name -> aper.test.v1.Choice2
	11, // 9: aper.test.v1.TestChoices.choice3:type_name -> aper.test.v1.Choice3
	12, // 10: aper.test.v1.TestChoices.choice4:type_name -> aper.test.v1.Choice4
	14, // 11: aper.test.v1.TestConstrainedChoices.constrained_choice1:type_name -> aper.test.v1.ConstrainedChoice1
	15, // 12: aper.test.v1.TestConstrainedChoices.constrained_choice2:type_name -> aper.test.v1.ConstrainedChoice2
	16, // 13: aper.test.v1.TestConstrainedChoices.constrained_choice3:type_name -> aper.test.v1.ConstrainedChoice3
	17, // 14: aper.test.v1.TestConstrainedChoices.constrained_choice4:type_name -> aper.test.v1.ConstrainedChoice4
	11, // 15: aper.test.v1.TestNestedChoice.option1:type_name -> aper.test.v1.Choice3
	16, // 16: aper.test.v1.TestNestedChoice.option2:type_name -> aper.test.v1.ConstrainedChoice3
	17, // 17: aper.test.v1.TestNestedChoice.option3:type_name -> aper.test.v1.ConstrainedChoice4
	22, // 18: aper.test.v1.TestList1.value:type_name -> aper.test.v1.Item
	36, // 19: aper.test.v1.Item.item2:type_name -> asn1.v1.BitString
	24, // 20: aper.test.v1.TestList2.value:type_name -> aper.test.v1.ItemExtensible
	2,  // 21: aper.test.v1.TestFullyOptionalSequence.item4:type_name -> aper.test.v1.TestFullyOptionalSequenceItem4
	22, // 22: aper.test.v1.TestListExtensible1.value:type_name -> aper.test.v1.Item
	24, // 23: aper.test.v1.TestListExtensible2.value:type_name -> aper.test.v1.ItemExtensible
	25, // 24: aper.test.v1.TestListExtensible3.value:type_name -> aper.test.v1.TestFullyOptionalSequence
	25, // 25: aper.test.v1.TestList3.value:type_name -> aper.test.v1.TestFullyOptionalSequence
	3,  // 26: aper.test.v1.TestTopLevelPdu.opt1:type_name -> aper.test.v1.TestUnconstrainedInt
	6,  // 27: aper.test.v1.TestTopLevelPdu.opt2:type_name -> aper.test.v1.TestConstrainedReal
	18, // 28: aper.test.v1.TestTopLevelPdu.opt3:type_name -> aper.test.v1.TestNestedChoice
	7,  // 29: aper.test.v1.TestTopLevelPdu.opt4:type_name -> aper.test.v1.TestBitString
	19, // 30: aper.test.v1.TestTopLevelPdu.opt5:type_name -> aper.test.v1.TestOctetString
	28, // 31: aper.test.v1.TestTopLevelPdu.opt6:type_name -> aper.test.v1.TestListExtensible3
	1,  // 32: aper.test.v1.TestTopLevelPdu.opt7:type_name -> aper.test.v1.TestEnumeratedExtensible
	32, // 33: aper.test.v1.sampleE2apPduChoice.ch:type_name -> aper.test.v1.CanonicalChoice
	33, // 34: aper.test.v1.CanonicalChoice.ch1:type_name -> aper.test.v1.SampleOctetString
	34, // 35: aper.test.v1.CanonicalChoice.ch2:type_name -> aper.test.v1.SampleConstrainedInteger
	35, // 36: aper.test.v1.CanonicalChoice.ch3:type_name -> aper.test.v1.SampleBitString
	26, // 37: aper.test.v1.CanonicalChoice.ch4:type_name -> aper.test.v1.TestListExtensible1
	22, // 38: aper.test.v1.CanonicalChoice.ch5:type_name -> aper.test.v1.Item
	36, // 39: aper.test.v1.SampleBitString.value:type_name -> asn1.v1.BitString
	40, // [40:40] is the sub-list for method output_type
	40, // [40:40] is the sub-list for method input_type
	40, // [40:40] is the sub-list for extension type_name
	40, // [40:40] is the sub-list for extension extendee
	0,  // [0:40] is the sub-list for field type_name
}

func init() { file_pkg_asn1_testsm_test_sm_proto_init() }
func file_pkg_asn1_testsm_test_sm_proto_init() {
	if File_pkg_asn1_testsm_test_sm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestUnconstrainedInt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConstrainedInt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestUnconstrainedReal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConstrainedReal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestBitString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestChoices); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Choice1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Choice2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Choice3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Choice4); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConstrainedChoices); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstrainedChoice1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstrainedChoice2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstrainedChoice3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstrainedChoice4); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestNestedChoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestOctetString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPrintableString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestList1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestList2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemExtensible); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestFullyOptionalSequence); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestListExtensible1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestListExtensible2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestListExtensible3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestList3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[27].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestTopLevelPdu); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[28].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleE2ApPduChoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[29].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanonicalChoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[30].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleOctetString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[31].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleConstrainedInteger); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_asn1_testsm_test_sm_proto_msgTypes[32].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleBitString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Choice1_Choice1A)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*Choice2_Choice2A)(nil),
		(*Choice2_Choice2B)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*Choice3_Choice3A)(nil),
		(*Choice3_Choice3B)(nil),
		(*Choice3_Choice3C)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Choice4_Choice4A)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*ConstrainedChoice1_ConstrainedChoice1A)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[12].OneofWrappers = []interface{}{
		(*ConstrainedChoice2_ConstrainedChoice2A)(nil),
		(*ConstrainedChoice2_ConstrainedChoice2B)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[13].OneofWrappers = []interface{}{
		(*ConstrainedChoice3_ConstrainedChoice3A)(nil),
		(*ConstrainedChoice3_ConstrainedChoice3B)(nil),
		(*ConstrainedChoice3_ConstrainedChoice3C)(nil),
		(*ConstrainedChoice3_ConstrainedChoice3D)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[14].OneofWrappers = []interface{}{
		(*ConstrainedChoice4_ConstrainedChoice4A)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[15].OneofWrappers = []interface{}{
		(*TestNestedChoice_Option1)(nil),
		(*TestNestedChoice_Option2)(nil),
		(*TestNestedChoice_Option3)(nil),
	}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[16].OneofWrappers = []interface{}{}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[17].OneofWrappers = []interface{}{}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[19].OneofWrappers = []interface{}{}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[21].OneofWrappers = []interface{}{}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[22].OneofWrappers = []interface{}{}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[27].OneofWrappers = []interface{}{}
	file_pkg_asn1_testsm_test_sm_proto_msgTypes[29].OneofWrappers = []interface{}{
		(*CanonicalChoice_Ch1)(nil),
		(*CanonicalChoice_Ch2)(nil),
		(*CanonicalChoice_Ch3)(nil),
		(*CanonicalChoice_Ch4)(nil),
		(*CanonicalChoice_Ch5)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_asn1_testsm_test_sm_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   33,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_asn1_testsm_test_sm_proto_goTypes,
		DependencyIndexes: file_pkg_asn1_testsm_test_sm_proto_depIdxs,
		EnumInfos:         file_pkg_asn1_testsm_test_sm_proto_enumTypes,
		MessageInfos:      file_pkg_asn1_testsm_test_sm_proto_msgTypes,
	}.Build()
	File_pkg_asn1_testsm_test_sm_proto = out.File
	file_pkg_asn1_testsm_test_sm_proto_rawDesc = nil
	file_pkg_asn1_testsm_test_sm_proto_goTypes = nil
	file_pkg_asn1_testsm_test_sm_proto_depIdxs = nil
}
