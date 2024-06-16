package list

type ListNode[T any] struct {
	prev  *ListNode[T]
	next  *ListNode[T]
	value T
}

type Direction int

const (
	StartHead Direction = 0 // 从表头向尾部迭代
	StartTail Direction = 1 // 从表尾向表头迭代
)

type ListIter[T any] struct {
	next      *ListNode[T]
	direction Direction // 迭代方向
}

type List[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]

	// 节点复制函数
	// 节点释放函数
	// 节点比较函数

	len int
}

func (l *List[T]) Length() int {
	return l.len
}

func (l *List[T]) First() *ListNode[T] {
	return l.head
}

func (l *List[T]) Last() *ListNode[T] {
	return l.tail
}

func Create[T any]() List[T] {

}

func (l *List[T]) AddNodeHead(value T) {

}

func (l *List[T]) AddNodeTail(value T) {

}
