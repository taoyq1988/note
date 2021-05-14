package _60_LFU

import (
	"fmt"
	"sort"
)

type Node struct {
	key   int
	value int
	seqNo int
	prev  *Node
	next  *Node
	tail  *Node
}

type LFUCache struct {
	capacity int
	data     map[int]*Node
	minSeq   int
	seqMap   map[int]*Node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		data:     make(map[int]*Node),
		seqMap:   make(map[int]*Node),
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		seqHeadNode := this.seqMap[node.seqNo]
		if seqHeadNode == node {
			if node.next == nil {
				delete(this.seqMap, node.seqNo)
			} else {
				nowTailNode := node.tail
				nextNode := node.next
				this.seqMap[node.seqNo] = nextNode
				node.next = nil
				nextNode.prev = nil
				nextNode.tail = nowTailNode
			}
		} else if seqHeadNode.tail == node {
			newTailNode := node.prev
			seqHeadNode.tail = newTailNode
			newTailNode.next = nil
			node.prev = nil
		} else {
			node.prev.next = node.next
			node.next.prev = node.prev
			node.prev = nil
			node.next = nil
		}
		node.seqNo++
		if this.seqMap[this.minSeq] == nil {
			this.minSeq = node.seqNo
		}
		node.next = nil
		if newSeqNode := this.seqMap[node.seqNo]; newSeqNode == nil {
			this.seqMap[node.seqNo] = node
			node.tail = node
		} else {
			nowTailNode := newSeqNode.tail
			newSeqNode.tail = node
			node.prev = nowTailNode
			nowTailNode.next = node
		}
		return node.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// if exist
	if node, ok := this.data[key]; ok {
		node.value = value
		seqHeadNode := this.seqMap[node.seqNo]
		if seqHeadNode == node {
			if node.next == nil {
				delete(this.seqMap, node.seqNo)
			} else {
				nowTailNode := node.tail
				nextNode := node.next
				this.seqMap[node.seqNo] = nextNode
				node.next = nil
				nextNode.prev = nil
				nextNode.tail = nowTailNode
			}
		} else if seqHeadNode.tail == node {
			newTailNode := node.prev
			seqHeadNode.tail = newTailNode
			newTailNode.next = nil
			node.prev = nil
		} else {
			node.prev.next = node.next
			node.next.prev = node.prev
			node.prev = nil
			node.next = nil
		}
		node.seqNo++
		if this.seqMap[this.minSeq] == nil {
			this.minSeq = node.seqNo
		}
		node.next = nil
		if newSeqNode := this.seqMap[node.seqNo]; newSeqNode == nil {
			this.seqMap[node.seqNo] = node
			node.tail = node
		} else {
			nowTailNode := newSeqNode.tail
			newSeqNode.tail = node
			node.prev = nowTailNode
			nowTailNode.next = node
		}
		return
	}
	// if not exist, create a new node
	node := &Node{
		key:   key,
		value: value,
		seqNo: 1,
	}
	if this.capacity == len(this.data) {
		minSeq := this.minSeq
		this.minSeq = node.seqNo
		rmNode := this.seqMap[minSeq]
		if rmNode.next == nil {
			delete(this.seqMap, minSeq)
		} else {
			nextHeadNode := rmNode.next
			nowTailNode := rmNode.tail
			nextHeadNode.prev = nil
			rmNode.next = nil
			this.seqMap[minSeq] = nextHeadNode
			nextHeadNode.tail = nowTailNode
		}
		delete(this.data, rmNode.key)
		seqOneNode := this.seqMap[node.seqNo]
		if seqOneNode == nil {
			this.seqMap[node.seqNo] = node
			node.tail = node
		} else {
			nowTailNode := seqOneNode.tail
			seqOneNode.tail = node
			nowTailNode.next = node
			node.prev = nowTailNode
		}
		this.data[key] = node
	} else {
		this.data[key] = node
		if oldSeqNode, _ := this.seqMap[node.seqNo]; oldSeqNode == nil {
			this.seqMap[node.seqNo] = node
			node.prev, node.next = nil, nil
			node.tail = node
			this.minSeq = node.seqNo
		} else {
			tailNode := oldSeqNode.tail
			tailNode.next = node
			node.prev = tailNode
			oldSeqNode.tail = node
		}
	}
}

func (this LFUCache) String() string {
	result := "====== LFUCache info start ======\n"
	result += fmt.Sprintf("the cache capacity is %d(%d/%d), now min seqNo is %d\n", this.capacity, len(this.data), this.capacity, this.minSeq)
	result += "the data is\n"
	keys1 := sortMapKey(this.data)
	for _, k := range keys1 {
		result += fmt.Sprintf("\tkey:%d, value:%d\n", k, this.data[k].value)
	}
	result += "the seq use is\n"
	keys2 := sortMapKey(this.seqMap)
	for _, k := range keys2 {
		node := this.seqMap[k]
		nodeLink := ""
		for node != nil {
			nodeLink += fmt.Sprintf("{%d=%d}->", node.key, node.value)
			node = node.next
		}
		if len(nodeLink) > 2 {
			nodeLink = nodeLink[:len(nodeLink)-2]
		}
		result += fmt.Sprintf("\tseq:%d, node link is %s\n", k, nodeLink)
	}
	result += "====== LFUCache info end ======\n"
	return result
}

func sortMapKey(m map[int]*Node) []int {
	keys := make(sort.IntSlice, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
