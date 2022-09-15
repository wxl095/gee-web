package bytesconv

import "unsafe"

func StringsToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Len int
		}{
			str,
			len(str),
		}))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
