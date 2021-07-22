package helper

import (
	"unsafe"
)

func Byte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2byte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
