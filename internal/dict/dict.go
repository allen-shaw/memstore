package dict

import (
	"errors"
	"github.com/allen-shaw/memstore/internal/assert"
	"math"
)

var (
	ErrorKeyExisted = errors.New("key existed")
	ErrorExpandFail = errors.New("expand fail")
)

type Entry struct {
	key interface{}
	v   struct {
		val interface{}
		u64 uint64
		s64 int64
	}
	next *Entry
}

type Type interface {
	Hash(key interface{}) int

	KeyDup(privData, key interface{}) interface{}

	ValDup(privData, obj interface{}) interface{}

	KeyCompare(privData interface{}, key1, key2 interface{}) int

	KeyDestructor(privData, key interface{})

	ValDestructor(privData, obj interface{})
}

type dictht struct {
	table    []*Entry
	size     uint64
	sizeMask uint64
	used     uint64
}

type Dict struct {
	typ       Type
	privData  interface{}
	ht        [2]*dictht
	rehashIdx int
	iterators int
}

func (d *Dict) freeVal(entry Entry) {
	d.typ.ValDestructor(d.privData, entry.v.val)
}

func (d *Dict) setVal(entry *Entry, val interface{}) {
	entry.v.val = d.typ.ValDup(d.privData, val)
}

type Iterator struct {
	d     *Dict
	table int
	index int
	safe  int

	entry     *Entry
	nextEntry *Entry

	fingerprint uint64
}

type ScanFunc func(privData interface{}, e *Entry)

const (
	dictHtInitialSize uint64 = 4
)

func Create(typ Type, privData interface{}) *Dict {
	d := &Dict{}
	d.init(typ, privData)
	return d
}

func (d *Dict) init(typ Type, privData interface{}) error {
	d.reset(d.ht[0])
	d.reset(d.ht[1])

	d.typ = typ
	d.privData = privData
	d.rehashIdx = -1
	d.iterators = 0

	return nil
}

func (d *Dict) reset(ht *dictht) {
	ht.table = nil
	ht.size = 0
	ht.sizeMask = 0
	ht.used = 0
}

func (d *Dict) Add(key, val interface{}) error {
	entry := d.add(key, val)
	if entry != nil {
		return ErrorKeyExisted
	}

	d.setVal(entry, val)
	return nil
}

func (d *Dict) add(key, val interface{}) *Entry {
	var (
		index int
		entry *Entry
		ht    *dictht
	)

	if d.isRehashing() {
		d.rehashStep()
	}

	index = d.keyIndex(key)
	if index == -1 {
		return nil
	}

	ht = d.ht[0]
	if d.isRehashing() {
		ht = d.ht[1]
	}

	entry = &Entry{}
	entry.next = ht.table[index]
	ht.used++

	d.setKey(entry, key)
	return entry
}

func (d *Dict) isRehashing() bool {
	return d.rehashIdx != -1
}

func (d *Dict) rehashStep() {
	if d.iterators == 0 {
		d.rehash(1)
	}
}

func (d *Dict) rehash(n int) bool {
	if d.isRehashing() {
		return false
	}

	for n > 0 {
		var de, nextde *Entry

		if d.ht[0].used == 0 {
			d.ht[0].table = nil // free, TODO: 可以使用arena
			d.ht[0] = d.ht[1]
			d.reset(d.ht[1])
			d.rehashIdx = -1
			return false
		}

		assert.Assert(d.ht[0].size > uint64(d.rehashIdx))

		for d.ht[0].table[d.rehashIdx] == nil {
			d.rehashIdx++
		}

		de = d.ht[0].table[d.rehashIdx]
		for de != nil {
			nextde = de.next
			h := uint64(d.hashKey(de.key)) & d.ht[1].sizeMask
			de.next = d.ht[1].table[h]
			d.ht[1].table[h] = de

			d.ht[0].used--
			d.ht[1].used++

			de = nextde
		}

		d.ht[0].table[d.rehashIdx] = nil
		d.rehashIdx++
		n--
	}
	return true
}

func (d *Dict) hashKey(key interface{}) int {
	return d.typ.Hash(key)
}

func (d *Dict) keyIndex(key interface{}) int {
	var (
		h, idx, table uint64
		he            *Entry
	)

	if d.expandIfNeeded() != nil {
		return -1
	}

	h = uint64(d.hashKey(key))
	for table = 0; table <= 1; table++ {
		idx = h & d.ht[table].sizeMask
		he = d.ht[table].table[idx]
		for he != nil {
			if d.compareKey(key, he.key) > 0 {
				return -1
			}
			he = he.next
		}
		if d.isRehashing() {
			break
		}
	}

	return int(idx)
}

func (d *Dict) setKey(entry *Entry, key interface{}) {
	entry.key = d.typ.KeyDup(d.privData, entry.key)
}

var (
	dictCanResize        = true
	dictForceResizeRatio = uint64(5)
)

func (d *Dict) expandIfNeeded() error {
	if d.isRehashing() {
		return nil
	}

	if d.ht[0].size == 0 {
		return d.expand(dictHtInitialSize)
	}

	if d.ht[0].used >= d.ht[0].size &&
		(dictCanResize ||
			d.ht[0].used/d.ht[0].size > dictForceResizeRatio) {
		return d.expand(d.ht[0].used * 2)
	}

	return nil
}

func (d *Dict) expand(size uint64) error {
	n := &dictht{}
	realSize := d.nextPower(size)

	if d.isRehashing() || d.ht[0].used > size {
		return ErrorExpandFail
	}

	n.size = realSize
	n.sizeMask = realSize - 1
	n.table = make([]*Entry, realSize)
	n.used = 0

	if d.ht[0].table == nil {
		d.ht[0] = n
		return nil
	}

	d.ht[1] = n
	d.rehashIdx = 0
	return nil
}

func (d *Dict) nextPower(size uint64) uint64 {
	i := dictHtInitialSize
	if size >= math.MaxInt64 {
		return math.MaxInt64
	}
	for {
		if i >= size {
			return i
		}
		i *= 2
	}
}

func (d *Dict) compareKey(key1, key2 interface{}) int {
	return d.typ.KeyCompare(d.privData, key1, key2)
}
