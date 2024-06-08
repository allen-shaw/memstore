package string

import (
	"strings"
	"unsafe"
)

type String2 struct {
	buf []byte
}

func New2(init []byte) String2 {
	initLen := len(init)
	return String2{
		buf: make([]byte, initLen),
	}
}

func (s *String2) Cat(t String2) {
	s.buf = append(s.buf, t.buf...)
}

func (s *String2) Trim(cset string) {
	s.buf = s2b(strings.Trim(b2s(s.buf), cset))
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func s2b(s string) (b []byte) {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
