package Heap

type node struct {
	Left  *node
	Right *node
	Value int
}

type lHeap struct {
	node    *node
	content [][]*node
	count   int
}

func NewLittleHeap() *lHeap {
	heap := &lHeap{}
	heap.Init()
	return heap
}

func (boku *lHeap) Init() {
	if boku.content != nil || boku.count != 0 {
		return
	}
	boku.content = make([][]*node, 1)
	boku.content[0] = make([]*node, 0)
}

func (boku *lHeap) Push(i int) {
	el := &node{
		Value: i,
	}
	if len(boku.content[len(boku.content)-1]) < (1 << uint(len(boku.content)-1)) {
		boku.content[len(boku.content)-1] = append(boku.content[len(boku.content)-1], el)
	} else {
		l := make([]*node, 0)
		l = append(l, el)
		boku.content = append(boku.content, l)
	}
	boku.count += 1
	if boku.count > 1 {
		boku.float(el, len(boku.content[len(boku.content)-1])-1, len(boku.content)-1)
	}
}

func (boku *lHeap) Pop() int {

	if boku.count == 0 {
		return 0
	}
	result := boku.content[0][0].Value
	boku.count -= 1
	if boku.count > 0 {
		lastList := boku.content[len(boku.content)-1]
		val := lastList[len(lastList)-1].Value //最后一个元素
		boku.content[0][0].Value = val
		lastList = lastList[:len(lastList)-1]
		if len(lastList) > 0 {
			boku.content[len(boku.content)-1] = lastList
		} else {
			boku.content = boku.content[:len(boku.content)-1]
		}
		if boku.count > 1 {
			boku.sink(boku.content[0][0], 0, 0)
		}
	} else {
		boku.Init()
	}
	return result
}

func (boku *lHeap) sink(el *node, x int, y int) {
	var l *node = nil
	var r *node = nil
	if y < len(boku.content)-1 && x*2 < len(boku.content[y+1]) {
		l = boku.content[y+1][x*2]
	}
	if y < len(boku.content)-1 && x*2+1 < len(boku.content[y+1]) {
		r = boku.content[y+1][x*2+1]
	}
	if l == nil && r == nil {
		return
	}

	if r == nil {
		if el.Value > l.Value {
			l.Value += el.Value
			el.Value = l.Value - el.Value
			l.Value -= el.Value

			y += 1
			x = x * 2
			boku.sink(l, x, y)
		}
	} else {
		if r.Value > l.Value {
			if el.Value > l.Value {
				l.Value += el.Value
				el.Value = l.Value - el.Value
				l.Value -= el.Value

				y += 1
				x = x * 2
				boku.sink(l, x, y)
			}
		} else {
			if el.Value > r.Value {
				r.Value += el.Value
				el.Value = r.Value - el.Value
				r.Value -= el.Value

				y += 1
				x = x*2 + 1
				boku.sink(r, x, y)
			}
		}
	}
}

func (boku *lHeap) float(el *node, x int, y int) {
	if y == 0 {
		return
	}
	p := boku.content[y-1][x/2]
	if el.Value < p.Value {
		p.Value += el.Value
		el.Value = p.Value - el.Value
		p.Value -= el.Value
		if y-1 > 0 {
			boku.float(p, x/2, y-1)
		}
	}
}

func (boku *lHeap) Len() int {
	return boku.count
}
