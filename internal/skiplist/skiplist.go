package skiplist

import (
	"github.com/allen-shaw/memstore/internal/object"
	"math/rand/v2"
)

type skiplistLevel struct {
	forward *Node
	span    uint
}

type Node struct {
	obj   *object.Object
	score float64

	backward *Node

	level []skiplistLevel
}

type Skiplist struct {
	header, tail *Node
	length       uint
	level        int
}

const (
	skiplistMaxLevel = 32
	skiplistP        = 0.25
)

func Create() *Skiplist {
	zsl := &Skiplist{}
	zsl.level = 1
	zsl.length = 0

	zsl.header = CreateNode(skiplistMaxLevel, 0, nil)
	for j := 0; j < skiplistMaxLevel; j++ {
		zsl.header.level[j].forward = nil
		zsl.header.level[j].span = 0
	}
	zsl.header.backward = nil
	zsl.tail = nil
	return zsl
}

func CreateNode(level int, score float64, obj *object.Object) *Node {
	zn := &Node{}
	zn.level = make([]skiplistLevel, level)
	zn.score = score
	zn.obj = obj
	return zn
}

func (zsl *Skiplist) Insert(score float64, obj *object.Object) *Node {
	var (
		x        *Node
		update   = make([]*Node, skiplistMaxLevel)
		i, level int
		rank     = make([]uint, skiplistMaxLevel)
	)

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		if i == zsl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score || (x.level[i].forward.score == score &&
				compareStringObjects(x.level[i].forward.obj, obj) < 0)) {

			rank[i] += x.level[i].span
			x = x.level[i].forward
		}

		update[i] = x
	}

	level = zsl.randomLevel()
	if level > zsl.level {
		for i = zsl.level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.header
			update[i].level[i].span = zsl.length
		}
		zsl.level = level
	}

	x = CreateNode(level, score, obj)
	for i = 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i = level; i < zsl.level; i++ {
		update[i].level[i].span++
	}

	x.backward = update[0]
	if update[0] == zsl.header {
		x.backward = nil
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		zsl.tail = x
	}

	zsl.length++
	return x
}

func (zsl *Skiplist) randomLevel() int {
	level := 1
	for float64(rand.Int()&0xFFFF) < (skiplistP * 0xFFFF) {
		level += 1
	}

	if level < skiplistMaxLevel {
		return level
	}
	return skiplistMaxLevel
}

func compareStringObjects(obj1, obj2 *object.Object) int {
	panic("no implement")
}
