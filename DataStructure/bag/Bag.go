package bag

// bag data structure (multiset).
type Bag struct {
	size int
	data []int
}

// Creates a new empty bag.
func New() *Bag {
	return &Bag{0, make([]int,0)}
}

// Inserts an element into the bag.
func (b *Bag) Insert(val int) {
	b.data = append(b.data,val)
}

func (b *Bag) Get() []int{
	return b.data
}

// Removes an element from the bag. If none was present, nothing is done.
func (b *Bag) Remove(i int) {
	for k:=0;k<len(b.data);k++{
		if b.data[k]==i{
			b.data[k] = b.data[b.Size()-1]
			b.data = b.data[:b.Size()-1]
			b.size--
			break
		}
	}
}


// Returns the total number of elemens in the bag.
func (b *Bag) Size() int {
	return len(b.data)
}


// Clears the contents of a bag.
func (b *Bag) Reset() {
	*b = *New()
}
