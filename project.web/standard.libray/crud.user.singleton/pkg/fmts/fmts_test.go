package fmts

import (
	"testing"
)

//func Test_buildStr(t *testing.T) {
//	var many1String, many1Int, many1Int8, many1Int16, many1Int32, many1Int64 interface{}
//	var many1Uint, many1Uint8, many1Uint16, many1Uint32, many1Uint64, many1Uintptr interface{}
//	var many1Float32, many1Float64 interface{}
//
//	many1String = string("21")
//	many1Int = int(100)
//	many1Int8 = int8(100)
//	many1Int16 = int16(100)
//	many1Int32 = int32(100)
//	many1Int64 = int64(100)
//	many1Uint = uint(100)
//	many1Uint8 = uint8(100)
//	many1Uint16 = uint16(100)
//	many1Uint32 = uint32(100)
//	many1Uint64 = uint64(100)
//	many1Uintptr = uintptr(100)
//
//	many1Float32 = float32(100.01)
//	many1Float64 = float64(100.01)
//	type args struct {
//		str interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//		{"test_buildStr_1", args{str: many1String}, "21"},
//		{"test_buildStr_2", args{str: many1Int}, "100"},
//		{"test_buildStr_3", args{str: many1Int8}, "100"},
//		{"test_buildStr_4", args{str: many1Int16}, "100"},
//		{"test_buildStr_5", args{str: many1Int32}, "100"},
//		{"test_buildStr_6", args{str: many1Int64}, "100"},
//		{"test_buildStr_7", args{str: many1Uint}, "100"},
//		{"test_buildStr_8", args{str: many1Uint8}, "100"},
//		{"test_buildStr_9", args{str: many1Uint16}, "100"},
//		{"test_buildStr_10", args{str: many1Uint32}, "100"},
//		{"test_buildStr_11", args{str: many1Uint64}, "100"},
//		{"test_buildStr_12", args{str: many1Uintptr}, "d"}, //endereço do ponteiro
//		{"test_buildStr_13", args{str: many1Float32}, "100.010002"},
//		{"test_buildStr_14", args{str: many1Float64}, "100.010000"},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := buildStr(tt.args.str); got != tt.want {
//				t.Errorf("buildStr() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

// go test -v -run ^TestStdout$
func TestStdout(t *testing.T) {
	var manyStdout []interface{}
	manyStdout = append(manyStdout, []string{"1"})

	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test_stdout_", args{strs: manyStdout}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Stdout(tt.args.strs...)
		})
	}
}

// go test -v -run ^TestConcat$
func TestConcat(t *testing.T) {

	var many1String, many1Int, many1Int8, many1Int16, many1Int32, many1Int64 []interface{}
	var many1Uint, many1Uint8, many1Uint16, many1Uint32, many1Uint64, many1Uintptr []interface{}
	var many1Float32, many1Float64 []interface{}

	many1String = append(many1String, []string{"21", "joao", "beta"})
	many1Int = append(many1Int, []int{100, 98, 23})
	many1Int8 = append(many1Int8, []int8{100, 98, 23})
	many1Int16 = append(many1Int16, []int16{100, 98, 23})
	many1Int32 = append(many1Int32, []int32{100, 98, 23})
	many1Int64 = append(many1Int64, []int64{100, 98, 23})
	many1Uint = append(many1Uint, []uint{100, 98, 23})
	many1Uint8 = append(many1Uint8, []uint8{100, 98, 23})
	many1Uint16 = append(many1Uint16, []uint16{100, 98, 23})
	many1Uint32 = append(many1Uint32, []uint32{100, 98, 23})
	many1Uint64 = append(many1Uint64, []uint64{100, 98, 23})
	many1Uintptr = append(many1Uintptr, []uintptr{100})

	many1Float32 = append(many1Float32, []float32{100.01, 100.01})
	many1Float64 = append(many1Float64, []float64{100.01, 100.01})
	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test_concat_", args{strs: many1String}, "21joaobeta"},
		{"test_concat_", args{strs: many1Int}, "1009823"},
		{"test_concat_", args{strs: many1Int8}, "1009823"},
		{"test_concat_", args{strs: many1Int16}, "1009823"},
		{"test_concat_", args{strs: many1Int32}, "1009823"},
		{"test_concat_", args{strs: many1Int64}, "1009823"},
		{"test_concat_", args{strs: many1Uint}, "1009823"},
		{"test_concat_", args{strs: many1Uint8}, "1009823"},
		{"test_concat_", args{strs: many1Uint16}, "1009823"},
		{"test_concat_", args{strs: many1Uint32}, "1009823"},
		{"test_concat_", args{strs: many1Uint64}, "1009823"},
		//	{"test_concat_", args{strs: many1Uintptr}, "d"}, //endereço do ponteiro
		{"test_concat_", args{strs: many1Float32}, "100.010002100.010002"},
		{"test_concat_", args{strs: many1Float64}, "100.010000100.010000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.strs...); got != tt.want {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run ^TestPrintln$
func TestPrintln(t *testing.T) {
	var manyPrintln []interface{}
	manyPrintln = append(manyPrintln, []string{"1", "2"})

	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test_println_", args{strs: manyPrintln}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Println(tt.args.strs...)
		})
	}
}

// go test -v -run ^TestPrint$
func TestPrint(t *testing.T) {
	var manyPrint []interface{}
	manyPrint = append(manyPrint, []string{"1", "2"})
	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test_print_", args{strs: manyPrint}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.args.strs...)
		})
	}
}
