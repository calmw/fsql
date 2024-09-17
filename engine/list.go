package engine

const (
	BNODE_FREE_LIST  = 3
	FREE_LIST_HEADER = 4 + 8 + 8
	FREE_LIST_CAP    = (BTREE_PAGE_SIZE - FREE_LIST_HEADER) / 8
)

//type FreeList struct {
//	head uint64
//	// 用于管理磁盘上页的回调
//	get func(uint64) BNode  // 引用指针
//	new func(BNode) uint64  // 分配一个新的页
//	use func(uint64, BNode) // 释放一个页
//}
//
//func (fl *FreeList) Get(topn int) uint64 {
//	assert(0 <= topn && topn < fl.Total())
//	node := fl.get(fl.head)
//	for flnSize(node) <= topn {
//		topn -= flnSize(node)
//		next := flnNext(node)
//		assert(next != 0)
//		node = fl.get(next)
//	}
//	return flnPtr(node, flnSize(node)-topn-1)
//}
//
//// remove `popn` pointers and add some new pointers
//func (fl *FreeList) Update(popn int, freed []uint64) {
//	assert(popn <= fl.Total())
//	if popn == 0 && len(freed) == 0 {
//		return // nothing to do
//	}
//	// prepare to construct the new list
//	total := fl.Total()
//	reuse := []uint64{}
//	for fl.head != 0 && len(reuse)*FREE_LIST_CAP < len(freed) {
//		node := fl.get(fl.head)
//		freed = append(freed, fl.head) // recyle the node itself
//		if popn >= flnSize(node) {
//			// phase 1
//			// remove all pointers in this node
//			popn -= flnSize(node)
//		} else {
//			// phase 2:
//			// remove some pointers
//			remain := flnSize(node) - popn
//			popn = 0
//			// reuse pointers from the free list itself
//			for remain > 0 && len(reuse)*FREE_LIST_CAP < len(freed)+remain {
//				remain--
//				reuse = append(reuse, flnPtr(node, remain))
//			}
//			// move the node into the `freed` list
//			for i := 0; i < remain; i++ {
//				freed = append(freed, flnPtr(node, i))
//			}
//		}
//		// discard the node and move to the next node
//		total -= flnSize(node)
//		fl.head = flnNext(node)
//	}
//	assert(len(reuse)*FREE_LIST_CAP >= len(freed) || fl.head == 0)
//	// phase 3: prepend new nodes
//	flPush(fl, freed, reuse)
//	// done
//	flnSetTotal(fl.get(fl.head), uint64(total+len(freed)))
//}
//
//func flPush(fl *FreeList, freed []uint64, reuse []uint64) {
//	for len(freed) > 0 {
//		new := BNode{make([]byte, BTREE_PAGE_SIZE)}
//		// construct a new node
//		size := len(freed)
//		if size > FREE_LIST_CAP {
//			size = FREE_LIST_CAP
//		}
//		flnSetHeader(new, uint16(size), fl.head)
//		for i, ptr := range freed[:size] {
//			flnSetPtr(new, i, ptr)
//		}
//		freed = freed[size:]
//		if len(reuse) > 0 {
//			// reuse a pointer from the list
//			fl.head, reuse = reuse[0], reuse[1:]
//			fl.use(fl.head, new)
//		} else {
//			// or append a page to house the new node
//			fl.head = fl.new(new)
//		}
//	}
//	assert(len(reuse) == 0)
//}
