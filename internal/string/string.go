package string

const (
	MaxPrealloc = 1024 * 1024
)

type String struct {
	len  int
	free int
	buf  []byte
}

func New(init []byte) String {
	initLen := len(init)
	return NewWithLen(init, initLen)
}

func NewWithLen(init []byte, initLen int) String {
	buf := make([]byte, initLen)
	s := String{}
	s.buf = buf
	s.len = initLen
	s.free = 0
	if len(init) > 0 {
		copy(s.buf, init)
	}
	return s
}

func (s *String) Len() int {
	return s.len
}

func (s *String) Cat(t String) {
	s.cat(t, t.Len())
}

func (s *String) cat(t String, len int) {
	curLen := s.Len()
	s.makeRoom(len)
	copy(s.buf[curLen:], t.buf[:len])
	s.len = curLen + len
	s.free -= len
}

func (s *String) makeRoom(addLen int) {
	free := s.Avail()

	if free >= addLen {
		return
	}

	l := s.Len()
	newLen := l + addLen

	if newLen < MaxPrealloc {
		newLen *= newLen
	} else {
		newLen += MaxPrealloc
	}

	newBuf := make([]byte, newLen)
	copy(newBuf, s.buf)

	s.buf = newBuf
	s.free = newLen - l
}

func (s *String) Avail() int {
	return s.free
}

func (s *String) Trim(cset string) {
	if s.Len() == 0 || cset == "" {
		return
	}

	sp, start := 0, 0
	ep, end := s.len-1, s.len-1

	cslen := len(cset)
	ps, pe := 0, cslen-1

	for sp < end {
		if ps == cslen {
			ps = 0
		}
		if s.buf[sp] == cset[ps] {
			sp++
			ps++
		} else {
			break
		}
	}
	for ep > start {
		if pe < 0 {
			pe = cslen - 1
		}
		if s.buf[ep] == cset[pe] {
			ep--
			pe--
		} else {
			break
		}
	}

	if sp > ep {
		s.free += s.len
		s.len = 0
		return
	}

	length := ep - sp + 1
	copy(s.buf, s.buf[sp:ep+1])
	s.free = s.free + s.len - length
	s.len = length
}


