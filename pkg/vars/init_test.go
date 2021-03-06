// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package vars

import (
	"fmt"
	"strconv"
	"strings"
)

// most of test vars should match with go/src/strconv/atof_test.go
// since parsing Value from Collection should pass undes same conditions.

type stringTest struct {
	key     string
	in      string
	want    string
	wantErr error
}

var stringTests = []stringTest{
	{"GOARCH", "amd64", "amd64", nil},
	{"GOHOSTARCH", "amd", "amd", nil},
	{"GOHOSTOS", "linux", "linux", nil},
	{"GOOS", "linux", "linux", nil},
	{"GOPATH", "/go-workspace", "/go-workspace", nil},
	{"GOROOT", "/usr/lib/golang", "/usr/lib/golang", nil},
	{"GOTOOLDIR", "/usr/lib/golang/pkg/tool/linux_amd64", "/usr/lib/golang/pkg/tool/linux_amd64", nil},
	{"GCCGO", "gccgo", "gccgo", nil},
	{"CC", "gcc", "gcc", nil},
	{"GOGCCFLAGS", "-fPIC -m64 -pthread -fmessage-length=0", "-fPIC -m64 -pthread -fmessage-length=0", nil},
	{"CXX", "g++", "g++", nil},
	{"PKG_CONFIG", "pkg-config", "pkg-config", nil},
	{"CGO_ENABLED", "1", "1", nil},
	{"CGO_CFLAGS", "-g -O2", "-g -O2", nil},
	{"CGO_CPPFLAGS", "", "", nil},
	{"CGO_CXXFLAGS", "-g -O2", "-g -O2", nil},
	{"CGO_FFLAGS", "-g -O2", "-g -O2", nil},
	{"CGO_LDFLAGS", "-g -O2", "-g -O2", nil},
}

func genStringTestBytes() []byte {
	var out []byte
	for _, data := range stringTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type atobTest struct {
	key     string
	in      string
	want    bool
	wantErr error
}

var atobTests = []atobTest{
	{"ATOB_1", "", false, strconv.ErrSyntax},
	{"ATOB_2", "asdf", false, strconv.ErrSyntax},
	{"ATOB_3", "0", false, nil},
	{"ATOB_4", "f", false, nil},
	{"ATOB_5", "F", false, nil},
	{"ATOB_6", "FALSE", false, nil},
	{"ATOB_7", "false", false, nil},
	{"ATOB_8", "False", false, nil},
	{"ATOB_9", "1", true, nil},
	{"ATOB_10", "t", true, nil},
	{"ATOB_11", "T", true, nil},
	{"ATOB_12", "TRUE", true, nil},
	{"ATOB_13", "true", true, nil},
	{"ATOB_14", "True", true, nil},
}

func genAtobTestBytes() []byte {
	var out []byte
	for _, data := range atobTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type atofTest struct {
	key     string
	in      string
	want    string
	wantErr error
}

var atofTests = []atofTest{
	{"FLOAT_1", "", "0", strconv.ErrSyntax},
	{"FLOAT_2", "1", "1", nil},
	{"FLOAT_3", "+1", "1", nil},
	{"FLOAT_4", "1x", "0", strconv.ErrSyntax},
	{"FLOAT_5", "1.1.", "0", strconv.ErrSyntax},
	{"FLOAT_6", "1e23", "1e+23", nil},
	{"FLOAT_7", "1E23", "1e+23", nil},
	{"FLOAT_8", "100000000000000000000000", "1e+23", nil},
	{"FLOAT_9", "1e-100", "1e-100", nil},
	{"FLOAT_10", "123456700", "1.234567e+08", nil},
	{"FLOAT_11", "99999999999999974834176", "9.999999999999997e+22", nil},
	{"FLOAT_12", "100000000000000000000001", "1.0000000000000001e+23", nil},
	{"FLOAT_13", "100000000000000008388608", "1.0000000000000001e+23", nil},
	{"FLOAT_14", "100000000000000016777215", "1.0000000000000001e+23", nil},
	{"FLOAT_15", "100000000000000016777216", "1.0000000000000003e+23", nil},
	{"FLOAT_16", "-1", "-1", nil},
	{"FLOAT_17", "-0.1", "-0.1", nil},
	{"FLOAT_18", "-0", "-0", nil},
	{"FLOAT_19", "1e-20", "1e-20", nil},
	{"FLOAT_20", "625e-3", "0.625", nil},

	// zeros
	{"FLOAT_21", "0", "0", nil},
	{"FLOAT_22", "0e0", "0", nil},
	{"FLOAT_24", "-0e0", "-0", nil},
	{"FLOAT_25", "+0e0", "0", nil},
	{"FLOAT_26", "0e-0", "0", nil},
	{"FLOAT_27", "-0e-0", "-0", nil},
	{"FLOAT_28", "+0e-0", "0", nil},
	{"FLOAT_29", "0e+0", "0", nil},
	{"FLOAT_30", "-0e+0", "-0", nil},
	{"FLOAT_31", "+0e+0", "0", nil},
	{"FLOAT_32", "0e+01234567890123456789", "0", nil},
	{"FLOAT_33", "0.00e-01234567890123456789", "0", nil},
	{"FLOAT_34", "-0e+01234567890123456789", "-0", nil},
	{"FLOAT_35", "-0.00e-01234567890123456789", "-0", nil},
	{"FLOAT_36", "0e291", "0", nil}, // issue 15364
	{"FLOAT_37", "0e292", "0", nil}, // issue 15364
	{"FLOAT_38", "0e347", "0", nil}, // issue 15364
	{"FLOAT_39", "0e348", "0", nil}, // issue 15364
	{"FLOAT_40", "-0e291", "-0", nil},
	{"FLOAT_41", "-0e292", "-0", nil},
	{"FLOAT_42", "-0e347", "-0", nil},
	{"FLOAT_43", "-0e348", "-0", nil},

	// NaNs
	{"FLOAT_44", "nan", "NaN", nil},
	{"FLOAT_45", "NaN", "NaN", nil},
	{"FLOAT_46", "NAN", "NaN", nil},

	// Infs
	{"FLOAT_47", "inf", "+Inf", nil},
	{"FLOAT_48", "-Inf", "-Inf", nil},
	{"FLOAT_49", "+INF", "+Inf", nil},
	{"FLOAT_50", "-Infinity", "-Inf", nil},
	{"FLOAT_51", "+INFINITY", "+Inf", nil},
	{"FLOAT_52", "Infinity", "+Inf", nil},

	// largest float64
	{"FLOAT_53", "1.7976931348623157e308", "1.7976931348623157e+308", nil},
	{"FLOAT_54", "-1.7976931348623157e308", "-1.7976931348623157e+308", nil},
	// next float64 - too large
	{"FLOAT_55", "1.7976931348623159e308", "+Inf", strconv.ErrRange},
	{"FLOAT_56", "-1.7976931348623159e308", "-Inf", strconv.ErrRange},
	// the border is ...158079
	// borderline - okay
	{"FLOAT_57", "1.7976931348623158e308", "1.7976931348623157e+308", nil},
	{"FLOAT_58", "-1.7976931348623158e308", "-1.7976931348623157e+308", nil},
	// borderline - too large
	{"FLOAT_59", "1.797693134862315808e308", "+Inf", strconv.ErrRange},
	{"FLOAT_60", "-1.797693134862315808e308", "-Inf", strconv.ErrRange},

	// a little too large
	{"FLOAT_61", "1e308", "1e+308", nil},
	{"FLOAT_62", "2e308", "+Inf", strconv.ErrRange},
	{"FLOAT_63", "1e309", "+Inf", strconv.ErrRange},

	// way too large
	{"FLOAT_64", "1e310", "+Inf", strconv.ErrRange},
	{"FLOAT_65", "-1e310", "-Inf", strconv.ErrRange},
	{"FLOAT_66", "1e400", "+Inf", strconv.ErrRange},
	{"FLOAT_67", "-1e400", "-Inf", strconv.ErrRange},
	{"FLOAT_68", "1e400000", "+Inf", strconv.ErrRange},
	{"FLOAT_69", "-1e400000", "-Inf", strconv.ErrRange},

	// denormalized
	{"FLOAT_70", "1e-305", "1e-305", nil},
	{"FLOAT_71", "1e-306", "1e-306", nil},
	{"FLOAT_72", "1e-307", "1e-307", nil},
	{"FLOAT_73", "1e-308", "1e-308", nil},
	{"FLOAT_74", "1e-309", "1e-309", nil},
	{"FLOAT_75", "1e-310", "1e-310", nil},
	{"FLOAT_76", "1e-322", "1e-322", nil},
	// smallest denormal
	{"FLOAT_77", "5e-324", "5e-324", nil},
	{"FLOAT_78", "4e-324", "5e-324", nil},
	{"FLOAT_79", "3e-324", "5e-324", nil},
	// too small
	{"FLOAT_80", "2e-324", "0", nil},
	// way too small
	{"FLOAT_81", "1e-350", "0", nil},
	{"FLOAT_82", "1e-400000", "0", nil},

	// try to overflow exponent
	{"FLOAT_83", "1e-4294967296", "0", nil},
	{"FLOAT_84", "1e+4294967296", "+Inf", strconv.ErrRange},
	{"FLOAT_85", "1e-18446744073709551616", "0", nil},
	{"FLOAT_86", "1e+18446744073709551616", "+Inf", strconv.ErrRange},

	// Parse errors
	{"FLOAT_87", "1e", "0", strconv.ErrSyntax},
	{"FLOAT_88", "1e-", "0", strconv.ErrSyntax},
	{"FLOAT_89", ".e-1", "0", strconv.ErrSyntax},
	{"FLOAT_90", "1\x00.2", "0", strconv.ErrSyntax},

	// http://www.exploringbinary.com/java-hangs-when-converting-2-2250738585072012e-308/
	{"FLOAT_91", "2.2250738585072012e-308", "2.2250738585072014e-308", nil},
	// http://www.exploringbinary.com/php-hangs-on-numeric-value-2-2250738585072011e-308/
	{"FLOAT_92", "2.2250738585072011e-308", "2.225073858507201e-308", nil},

	// A very large number (initially wrongly parsed by the fast algorithm).
	{"FLOAT_93", "4.630813248087435e+307", "4.630813248087435e+307", nil},

	// A different kind of very large number.
	{"FLOAT_94", "22.222222222222222", "22.22222222222222", nil},
	{"FLOAT_95", "2." + strings.Repeat("2", 4000) + "e+1", "22.22222222222222", nil},

	// Exactly halfway between 1 and math.Nextafter(1, 2).
	// Round to even (down).
	{"FLOAT_96", "1.00000000000000011102230246251565404236316680908203125", "1", nil},
	// Slightly lower; still round down.
	{"FLOAT_97", "1.00000000000000011102230246251565404236316680908203124", "1", nil},
	// Slightly higher; round up.
	{"FLOAT_98", "1.00000000000000011102230246251565404236316680908203126", "1.0000000000000002", nil},
	// Slightly higher, but you have to read all the way to the end.
	{"FLOAT_99", "1.00000000000000011102230246251565404236316680908203125" + strings.Repeat("0", 10000) + "1", "1.0000000000000002", nil},
}

func genAtofTestBytes() []byte {
	var out []byte
	for _, data := range atofTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

var atof32Tests = []atofTest{
	// Exactly halfway between 1 and the next float32.
	// Round to even (down).
	{"FLOAT_1", "1.000000059604644775390625", "1", nil},
	// Slightly lower.
	{"FLOAT_2", "1.000000059604644775390624", "1", nil},
	// Slightly higher.
	{"FLOAT_3", "1.000000059604644775390626", "1.0000001", nil},
	// Slightly higher, but you have to read all the way to the end.
	{"FLOAT_4", "1.000000059604644775390625" + strings.Repeat("0", 10000) + "1", "1.0000001", nil},

	// largest float32: (1<<128) * (1 - 2^-24)
	{"FLOAT_5", "340282346638528859811704183484516925440", "3.4028235e+38", nil},
	{"FLOAT_6", "-340282346638528859811704183484516925440", "-3.4028235e+38", nil},
	// next float32 - too large
	{"FLOAT_7", "3.4028236e38", "+Inf", strconv.ErrRange},
	{"FLOAT_8", "-3.4028236e38", "-Inf", strconv.ErrRange},
	// the border is 3.40282356779...e+38
	// borderline - okay
	{"FLOAT_9", "3.402823567e38", "3.4028235e+38", nil},
	{"FLOAT_10", "-3.402823567e38", "-3.4028235e+38", nil},
	// borderline - too large
	{"FLOAT_11", "3.4028235678e38", "+Inf", strconv.ErrRange},
	{"FLOAT_12", "-3.4028235678e38", "-Inf", strconv.ErrRange},

	// Denormals: less than 2^-126
	{"FLOAT_13", "1e-38", "1e-38", nil},
	{"FLOAT_14", "1e-39", "1e-39", nil},
	{"FLOAT_15", "1e-40", "1e-40", nil},
	{"FLOAT_16", "1e-41", "1e-41", nil},
	{"FLOAT_17", "1e-42", "1e-42", nil},
	{"FLOAT_18", "1e-43", "1e-43", nil},
	{"FLOAT_20", "1e-44", "1e-44", nil},
	{"FLOAT_21", "6e-45", "6e-45", nil}, // 4p-149 = 5.6e-45
	{"FLOAT_22", "5e-45", "6e-45", nil},
	// Smallest denormal
	{"FLOAT_23", "1e-45", "1e-45", nil}, // 1p-149 = 1.4e-45
	{"FLOAT_24", "2e-45", "1e-45", nil},

	// 2^92 = 8388608p+69 = 4951760157141521099596496896 (4.9517602e27)
	// is an exact power of two that needs 8 decimal digits to be correctly
	// parsed back.
	// The float32 before is 16777215p+68 = 4.95175986e+27
	// The halfway is 4.951760009. A bad algorithm that thinks the previous
	// float32 is 8388607p+69 will shorten incorrectly to 4.95176e+27.
	{"FLOAT_25", "4951760157141521099596496896", "4.9517602e+27", nil},
}

func genAtof32TestBytes() []byte {
	var out []byte
	for _, data := range atof32Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type atoui64Test struct {
	key     string
	in      string
	want    uint64
	wantErr error
}

var atoui64Tests = []atoui64Test{
	{"UINT64_1", "", 0, strconv.ErrSyntax},
	{"UINT64_2", "0", 0, nil},
	{"UINT64_3", "1", 1, nil},
	{"UINT64_4", "12345", 12345, nil},
	{"UINT64_5", "012345", 12345, nil},
	{"UINT64_6", "12345x", 0, strconv.ErrSyntax},
	{"UINT64_7", "98765432100", 98765432100, nil},
	{"UINT64_8", "18446744073709551615", 1<<64 - 1, nil},
	{"UINT64_9", "18446744073709551616", 1<<64 - 1, strconv.ErrRange},
	{"UINT64_10", "18446744073709551620", 1<<64 - 1, strconv.ErrRange},
}

func genAtoui64TestBytes() []byte {
	var out []byte
	for _, data := range atoui64Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

var btoui64Tests = []atoui64Test{
	{"UINT64_1", "", 0, strconv.ErrSyntax},
	{"UINT64_2", "0", 0, nil},
	{"UINT64_3", "0x", 0, strconv.ErrSyntax},
	{"UINT64_4", "0X", 0, strconv.ErrSyntax},
	{"UINT64_5", "1", 1, nil},
	{"UINT64_6", "12345", 12345, nil},
	{"UINT64_7", "012345", 012345, nil},
	{"UINT64_8", "0x12345", 0x12345, nil},
	{"UINT64_9", "0X12345", 0x12345, nil},
	{"UINT64_10", "12345x", 0, strconv.ErrSyntax},
	{"UINT64_11", "0xabcdefg123", 0, strconv.ErrSyntax},
	{"UINT64_12", "123456789abc", 0, strconv.ErrSyntax},
	{"UINT64_13", "98765432100", 98765432100, nil},
	{"UINT64_14", "18446744073709551615", 1<<64 - 1, nil},
	{"UINT64_15", "18446744073709551616", 1<<64 - 1, strconv.ErrRange},
	{"UINT64_16", "18446744073709551620", 1<<64 - 1, strconv.ErrRange},
	{"UINT64_17", "0xFFFFFFFFFFFFFFFF", 1<<64 - 1, nil},
	{"UINT64_18", "0x10000000000000000", 1<<64 - 1, strconv.ErrRange},
	{"UINT64_19", "01777777777777777777777", 1<<64 - 1, nil},
	{"UINT64_20", "01777777777777777777778", 0, strconv.ErrSyntax},
	{"UINT64_21", "02000000000000000000000", 1<<64 - 1, strconv.ErrRange},
	{"UINT64_22", "0200000000000000000000", 1 << 61, nil},
}

func genBtoui64TestBytes() []byte {
	var out []byte
	for _, data := range btoui64Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type atoi64Test struct {
	key     string
	in      string
	want    int64
	wantErr error
}

var atoi64Tests = []atoi64Test{
	{"INT64_1", "", 0, strconv.ErrSyntax},
	{"INT64_2", "0", 0, nil},
	{"INT64_3", "-0", 0, nil},
	{"INT64_4", "1", 1, nil},
	{"INT64_5", "-1", -1, nil},
	{"INT64_6", "12345", 12345, nil},
	{"INT64_7", "-12345", -12345, nil},
	{"INT64_8", "012345", 12345, nil},
	{"INT64_9", "-012345", -12345, nil},
	{"INT64_10", "98765432100", 98765432100, nil},
	{"INT64_11", "-98765432100", -98765432100, nil},
	{"INT64_12", "9223372036854775807", 1<<63 - 1, nil},
	{"INT64_13", "-9223372036854775807", -(1<<63 - 1), nil},
	{"INT64_14", "9223372036854775808", 1<<63 - 1, strconv.ErrRange},
	{"INT64_15", "-9223372036854775808", -1 << 63, nil},
	{"INT64_16", "9223372036854775809", 1<<63 - 1, strconv.ErrRange},
	{"INT64_17", "-9223372036854775809", -1 << 63, strconv.ErrRange},
}

func genAtoi64TestBytes() []byte {
	var out []byte
	for _, data := range atoi64Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type btoi64Test struct {
	key     string
	in      string
	base    int
	want    int64
	wantErr error
}

var btoi64Tests = []btoi64Test{
	{"INT64_1", "", 0, 0, strconv.ErrSyntax},
	{"INT64_2", "0", 0, 0, nil},
	{"INT64_3", "-0", 0, 0, nil},
	{"INT64_4", "1", 0, 1, nil},
	{"INT64_5", "-1", 0, -1, nil},
	{"INT64_6", "12345", 0, 12345, nil},
	{"INT64_7", "-12345", 0, -12345, nil},
	{"INT64_8", "012345", 0, 012345, nil},
	{"INT64_9", "-012345", 0, -012345, nil},
	{"INT64_10", "0x12345", 0, 0x12345, nil},
	{"INT64_11", "-0X12345", 0, -0x12345, nil},
	{"INT64_12", "12345x", 0, 0, strconv.ErrSyntax},
	{"INT64_13", "-12345x", 0, 0, strconv.ErrSyntax},
	{"INT64_14", "98765432100", 0, 98765432100, nil},
	{"INT64_15", "-98765432100", 0, -98765432100, nil},
	{"INT64_16", "9223372036854775807", 0, 1<<63 - 1, nil},
	{"INT64_17", "-9223372036854775807", 0, -(1<<63 - 1), nil},
	{"INT64_18", "9223372036854775808", 0, 1<<63 - 1, strconv.ErrRange},
	{"INT64_19", "-9223372036854775808", 0, -1 << 63, nil},
	{"INT64_20", "9223372036854775809", 0, 1<<63 - 1, strconv.ErrRange},
	{"INT64_21", "-9223372036854775809", 0, -1 << 63, strconv.ErrRange},

	// other bases
	{"INT64_22", "g", 17, 16, nil},
	{"INT64_24", "10", 25, 25, nil},
	{"INT64_25", "holycow", 35, (((((17*35+24)*35+21)*35+34)*35+12)*35+24)*35 + 32, nil},
	{"INT64_26", "holycow", 36, (((((17*36+24)*36+21)*36+34)*36+12)*36+24)*36 + 32, nil},

	// base 2
	{"INT64_27", "0", 2, 0, nil},
	{"INT64_28", "-1", 2, -1, nil},
	{"INT64_29", "1010", 2, 10, nil},
	{"INT64_30", "1000000000000000", 2, 1 << 15, nil},
	{"INT64_31", "111111111111111111111111111111111111111111111111111111111111111", 2, 1<<63 - 1, nil},
	{"INT64_32", "1000000000000000000000000000000000000000000000000000000000000000", 2, 1<<63 - 1, strconv.ErrRange},
	{"INT64_33", "-1000000000000000000000000000000000000000000000000000000000000000", 2, -1 << 63, nil},
	{"INT64_34", "-1000000000000000000000000000000000000000000000000000000000000001", 2, -1 << 63, strconv.ErrRange},

	// base 8
	{"INT64_35", "-10", 8, -8, nil},
	{"INT64_36", "57635436545", 8, 057635436545, nil},
	{"INT64_37", "100000000", 8, 1 << 24, nil},

	// base 16
	{"INT64_38", "10", 16, 16, nil},
	{"INT64_39", "-123456789abcdef", 16, -0x123456789abcdef, nil},
	{"INT64_40", "7fffffffffffffff", 16, 1<<63 - 1, nil},
}

func genBtoi64TestBytes() []byte {
	var out []byte
	for _, data := range btoi64Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type atoui32Test struct {
	key     string
	in      string
	want    uint32
	wantErr error
}

var atoui32Tests = []atoui32Test{
	{"UINT32_1", "", 0, strconv.ErrSyntax},
	{"UINT32_2", "0", 0, nil},
	{"UINT32_3", "1", 1, nil},
	{"UINT32_4", "12345", 12345, nil},
	{"UINT32_5", "012345", 12345, nil},
	{"UINT32_6", "12345x", 0, strconv.ErrSyntax},
	{"UINT32_7", "987654321", 987654321, nil},
	{"UINT32_8", "4294967295", 1<<32 - 1, nil},
	{"UINT32_9", "4294967296", 1<<32 - 1, strconv.ErrRange},
}

func genAtoui32TestBytes() []byte {
	var out []byte
	for _, data := range atoui32Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type atoi32Test struct {
	key     string
	in      string
	want    int32
	wantErr error
}

var atoi32tests = []atoi32Test{
	{"INT32_1", "", 0, strconv.ErrSyntax},
	{"INT32_2", "0", 0, nil},
	{"INT32_3", "-0", 0, nil},
	{"INT32_4", "1", 1, nil},
	{"INT32_5", "-1", -1, nil},
	{"INT32_6", "12345", 12345, nil},
	{"INT32_7", "-12345", -12345, nil},
	{"INT32_8", "012345", 12345, nil},
	{"INT32_9", "-012345", -12345, nil},
	{"INT32_10", "12345x", 0, strconv.ErrSyntax},
	{"INT32_11", "-12345x", 0, strconv.ErrSyntax},
	{"INT32_12", "987654321", 987654321, nil},
	{"INT32_13", "-987654321", -987654321, nil},
	{"INT32_14", "2147483647", 1<<31 - 1, nil},
	{"INT32_15", "-2147483647", -(1<<31 - 1), nil},
	{"INT32_16", "2147483648", 1<<31 - 1, strconv.ErrRange},
	{"INT32_17", "-2147483648", -1 << 31, nil},
	{"INT32_18", "2147483649", 1<<31 - 1, strconv.ErrRange},
	{"INT32_19", "-2147483649", -1 << 31, strconv.ErrRange},
}

func genAtoi32TestBytes() []byte {
	var out []byte
	for _, data := range atoui32Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type complex64Test struct {
	key     string
	in      string
	want    complex64
	wantErr error
}

var complex64Tests = []complex64Test{
	// Exactly halfway between 1 and the next float32.
	// Round to even (down).
	{"COMPLEX64_1", "1.000000059604644775390625 1.000000059604644775390624", complex64(complex(1.000000059604644775390625, 1.000000059604644775390624)), nil},
	// Slightly lower.
	{"COMPLEX64_2", "1", complex64(0), strconv.ErrSyntax},
	// Slightly higher.
	{"COMPLEX64_3", "1.000000059604644775390626 2", complex64(complex(1.0000001, 2)), nil},
	{"COMPLEX64_4", "1x -0", complex64(0), strconv.ErrSyntax},
	{"COMPLEX64_5", "-0 1x", complex64(0), strconv.ErrSyntax},
}

func genComplex64TestBytes() []byte {
	var out []byte
	for _, data := range complex64Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

type complex128Test struct {
	key     string
	in      string
	want    complex128
	wantErr error
}

var complex128Tests = []complex128Test{
	{"COMPLEX128_1", " 1", complex128(0), strconv.ErrSyntax},
	{"COMPLEX128_2", "+1 -1", complex128(complex(1, -1)), nil},
	{"COMPLEX128_3", "1x -0", complex128(0), strconv.ErrSyntax},
	{"COMPLEX128_3", "-0 1x", complex128(0), strconv.ErrSyntax},
	{"COMPLEX128_4", "1.1. 0", complex128(0), strconv.ErrSyntax},
	{"COMPLEX128_5", "1e23 1E23", complex128(complex(1e+23, 1e+23)), nil},
	{"COMPLEX128_6", "100000000000000000000000 1e-100", complex128(complex(1e+23, 1e-100)), nil},
	{"COMPLEX128_7", "123456700 1e-100", complex128(complex(1.234567e+08, 1e-100)), nil},
	{"COMPLEX128_8", "99999999999999974834176 100000000000000000000001", complex128(complex(9.999999999999997e+22, 1.0000000000000001e+23)), nil},
	{"COMPLEX128_9", "100000000000000008388608 100000000000000016777215", complex128(complex(1.0000000000000001e+23, 1.0000000000000001e+23)), nil},
	{"COMPLEX128_10", "1e-20 625e-3", complex128(complex(1e-20, 0.625)), nil},
}

func genComplex128TestBytes() []byte {
	var out []byte
	for _, data := range complex128Tests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}
