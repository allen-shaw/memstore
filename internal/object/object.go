package object

type Object struct {
	typ      uint
	encoding uint
	lru      uint
	ptr      interface{}
	refCount int
}
