package feb

type Iterator interface {
	hasNext() bool
	next() int
}

type PeekingIterator struct {
	iter   Iterator
	peeked int
	nextE  int
}

func Constructor(iter Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:   iter,
		peeked: iter.next(),
		nextE:  iter.next(),
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.iter.hasNext() || this.peeked != this.nextE
}

func (this *PeekingIterator) next() int {
	prev := this.peek()
	this.peeked = this.nextE
	this.nextE = this.iter.next()

	return prev
}

func (this *PeekingIterator) peek() int {
	return this.peeked
}
