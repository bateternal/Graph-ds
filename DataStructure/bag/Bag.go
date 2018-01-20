package bag

// bag data structure (multiset).
type Bag struct {
	size int
	data map[interface{}]int
}

// Creates a new empty bag.
func New() *Bag {
	return &Bag{0, make(map[interface{}]int)}
}

// Inserts an element into the bag.
func (b *Bag) Insert(val interface{}) {
	b.data[val]++
	b.size++
}

// Removes an element from the bag. If none was present, nothing is done.
func (b *Bag) Remove(val interface{}) {
	old, ok := b.data[val]
	if ok {
		if old > 1 {
			b.data[val] = old - 1
		} else {
			delete(b.data, val)
		}
		b.size--
	}
}

// Returns the total number of elemens in the bag.
func (b *Bag) Size() int {
	return b.size
}

// Counts the number of occurances of an element in the bag.
func (b *Bag) Count(val interface{}) int {
	return b.data[val]
}

// Clears the contents of a bag.
func (b *Bag) Reset() {
	*b = *New()
}
