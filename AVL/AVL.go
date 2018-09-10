package AVL

type node struct {
	parent *node
	height int
	Left   *node
	Right  *node
	root   *Tree
	Value  int
}
/*
AVL
 */

type Tree struct {
	node  *node
	count int
}

const INIT_VALUE = -1

func NewAVL() *Tree {
	t := new(Tree)
	t.node = &node{
		root:  t,
		Value: INIT_VALUE,
	}
	return t
}

func (this *Tree) Put(i int) bool {
	if this.node.put(i) {
		this.count += 1
		return true
	}
	return false
}

func (this *Tree) Find(i int) *node {
	return this.node.find(i)
}

func (this *Tree) Count() int {
	return this.count
}

func (this *Tree) Del(i int) bool {
	if this.node.find(i).del() {
		this.count -= 1
		return true
	}
	return false
}

func (this *Tree) Tree() *node {
	return this.node
}

func (this *Tree) List() []int {
	l := make([]int, 0)
	if this.node != nil {
		l = list(this.node, l)
		return l
	}
	return nil
}

/***************************************************************************/
/*
遍历树
 */
func list(n *node, l []int) []int {
	l = append(l, n.Value)
	if n.Left != nil {
		l = list(n.Left, l)
	}
	if n.Right != nil {
		l = list(n.Right, l)
	}
	return l
}

/*
查找结点
 */
func (this *node) find(i int) *node {
	if this == nil {
		return nil
	}
	if this.Value > i {
		return this.Left.find(i)
	} else if this.Value < i {
		return this.Right.find(i)
	} else if this.Value == i {
		return this
	} else {
		return nil
	}
}

/*
放置结点
 */
func (this *node) put(value int) bool {
	if this.Value == INIT_VALUE {
		this.Value = value
		this.height = 1
		this.heightBubble()
		return true
	}
	if this.Value > value {
		if this.Left == nil {
			this.Left = &node{
				parent: this,
				Value:  INIT_VALUE,
			}
		}
		return this.Left.put(value)
	} else if this.Value < value {
		if this.Right == nil {
			this.Right = &node{
				parent: this,
				Value:  INIT_VALUE,
			}
		}
		return this.Right.put(value)
	} else {
		return false
	}
}

/*
删除结点
 */
func (this *node) del() bool {
	if this == nil {
		return false
	}
	switch {
	case this.Left == nil && this.Right == nil:
		this.parentChange(nil)
		this.heightBubble()
		return true
	case this.Left != nil && this.Right == nil:
		this.Left.height -= 1
		this.parentChange(this.Left)
		this.heightBubble()
		return true
	case this.Left == nil && this.Right != nil:
		this.Right.height -= 1
		this.parentChange(this.Right)
		this.heightBubble()
		return true
	case this.Left != nil && this.Right != nil:
		r := this.Right.findHeir()
		this.Value = r.Value
		r.del()
		return true
	}
	return false
}

/*
翻树
 */
func (this *node) rotate() {
	var k1 *node
	var k2 *node
	var k3 *node
	var double = func(k1 *node, k2 *node, k3 *node) {
		if k2.Right != nil {
			k2.Right.parent = k3
			k3.Left = k2.Right
		} else {
			k3.Left = nil
		}
		if k2.Left != nil {
			k2.Left.parent = k1
			k1.Right = k2.Left
		} else {
			k1.Right = nil
		}
		k2.Left = k1
		k2.Right = k3
		k1.parent = k2
		k3.parent = k2
	}

	switch {
	case this.Left.Height() > this.Right.Height() && this.Left.Left.Height() >= this.Left.Right.Height(): //左左
		k1 := this
		k2 := this.Left
		this.parentChange(k2)
		if k2.Right != nil {
			k2.Right.parent = k1
			k1.Left = k2.Right
		} else {
			k1.Left = nil
		}
		k1.parent = k2
		k2.Right = k1
		k1.height -= 1
		k2.heightBubble()
	case this.Right.Height() > this.Left.Height() && this.Right.Right.Height() >= this.Right.Left.Height():
		k1 := this
		k2 := this.Right
		this.parentChange(k2)
		if k2.Left != nil {
			k2.Left.parent = k1
			k1.Right = k2.Left
		} else {
			k1.Right = nil
		}
		k1.parent = k2
		k2.Left = k1
		k1.height -= 1
		k2.heightBubble()
	case this.Right.Height() > this.Left.Height() && this.Right.Left.Height() > this.Right.Right.Height(): //右左
		k1 = this
		k2 = this.Right.Left
		k3 = this.Right
		double(k1, k2, k3)
		k1.height -= 2
		k3.height -= 1
		k2.height += 1
		this.parentChange(k3)
	case this.Left.Height() > this.Right.Height() && this.Left.Right.Height() > this.Left.Left.Height(): //左右
		k3 = this
		k2 = this.Left.Right
		k1 = this.Left
		double(k1, k2, k3)
		k1.height -= 1
		k3.height -= 2
		k2.height += 1
		this.parentChange(k3)
	}
}

/*
冒泡式更新结点高度
 */
func (this *node) heightBubble() {
	var max = func(a int, b int) (int, int) {
		if a > b {
			return a, a - b
		}
		return b, b - a
	}
	_parent := this.parent
	if _parent == nil {
		return
	}
	nh, d := max(_parent.Left.Height(), _parent.Right.Height())

	if d >= 2 {
		_parent.rotate()
	} else if nh+1 == _parent.Height() {
		return
	} else {
		_parent.height = nh + 1
		_parent.heightBubble()
	}
}

func (this *node) Height() int {
	if this == nil {
		return 0
	}
	return this.height
}

func (this *node) parentChange(s *node) {
	if this.root != nil {
		if s == nil { //刪除
			this.root.node = s
		} else {
			this.root.node = s
			s.root = this.root
			this.root = nil
		}
	} else {
		if this.parent.Left == this {
			this.parent.Left = s
		} else {
			this.parent.Right = s
		}
		if s != nil {
			s.parent = this.parent
		}
	}
}

/*
查找最小叶子
 */
func (this *node) findHeir() *node {
	if this.Left == nil {
		return this
	} else {
		return this.Left.findHeir()
	}
	return nil
}
