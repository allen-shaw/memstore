package intset

type IntSet struct {
	encoding uint32
	length   uint32
	contents []byte
}
