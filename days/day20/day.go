package day20

import (
	"fmt"
	"strings"

	. "github.com/panshadow/aoc2022/utils"
)

func init() {
	RegisterTask("20/01", Task01, "20/01.twitter")
	RegisterTask("20/02", Task02, "20/01.twitter")
}

type Node struct {
	Val int
	Index int
	Prev *Node
	Next *Node
}

func (n *Node) Sibling(fwd bool) *Node {
	if fwd {
		return n.Next
	} else {
		return n.Prev
	}
}

type List struct {
	Head *Node
	Len int
}

func NewList(xs []int) *List {
	list := &List{}
	var cur *Node
	for i,x := range xs {
		Debugln("Add ", x)
		node := &Node{}
		node.Val = x
		node.Index = i
		if list.Head == nil {
			Debugln("Add head ")
			list.Head = node
			cur = list.Head
		} else {
			Debugln("Add node after ", cur.Val)
			node.Prev = cur
			cur.Next = node
			cur = node
		}
	}
	cur.Next = list.Head
	list.Head.Prev = cur
	list.Len = len(xs)
	Debugf("New List(%d): %s\n",list.Len, list)
	return list
}


func (l *List) FindIndex(index int) *Node {
	i := 0
	cur := l.Head
	for i < l.Len {
		if cur.Index == index {
			return cur
		}
		i++
		cur = cur.Next
	}

	return nil
}

func (l *List) Find(x int) *Node {
	i := 0
	cur := l.Head
	for i < l.Len {
		if cur.Val == x {
			return cur
		}
		i++
		cur = cur.Next
	}

	return nil
}

func (l *List) Cut(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
}

func (l *List) InsertBefore(src, dst *Node) {
	src.Next = dst
	src.Prev = dst.Prev
	dst.Prev.Next = src
	dst.Prev = src
}

func (l *List) InsertAfter(src, dst *Node) {
	src.Prev = dst
	src.Next = dst.Next
	dst.Next.Prev = src
	dst.Next = src
}

func (l *List) Move(node *Node) {
	distance := node.Val
	if Abs(distance) > l.Len {
		distance = distance % l.Len
	}
	if distance == 0 {
		Debugln("Don't move ", node.Val)
		return
	}
	srcNode := node
	dstNode := srcNode

	if dstNode != nil {
		Debugln("Distance is ", node.Val, distance)
		if srcNode == l.Head {
			Debugf("Move HEAD from %d to %d\n", l.Head.Val, l.Head.Next.Val)
			l.Head = l.Head.Next
		}

		srcNode.Next.Prev = srcNode.Prev
		srcNode.Prev.Next = srcNode.Next


		if distance < 0 {
			for i :=0; i<Abs(distance); i++ {
				dstNode = dstNode.Prev
			}
			srcNode.Next = dstNode
			srcNode.Prev = dstNode.Prev
			dstNode.Prev.Next = srcNode
			dstNode.Prev = srcNode
		} else {
			for i :=0; i<Abs(distance); i++ {
				dstNode = dstNode.Next
			}
			srcNode.Prev = dstNode
			srcNode.Next = dstNode.Next
			dstNode.Next.Prev = srcNode
			dstNode.Next = srcNode
		}
		Debugf("Insert %d between %d and %d\n",srcNode.Val, srcNode.Prev.Val, srcNode.Next.Val)
	}
}

func (l *List) String() string {
	out := make([]string, l.Len)
	cur := l.Head
	for i :=0; i < l.Len; i++ {
		out[i] = fmt.Sprint(cur.Val)
		cur = cur.Next
	}
	return strings.Join(out, ", ")
}

func (l *List) Check() bool {
	Debugln("Check List")
	curF := l.Head
	curB := l.Head
	for i :=0; i < l.Len; i++ {
		curF = curF.Next
		curB = curB.Prev
	}
	Debugf("curB==head %v\n", curB==l.Head)
	Debugf("curF==head: %v\n", curF==l.Head)
	Debugf("curB==curF: %v\n", curB==curF)

	return curB==l.Head && curF==l.Head
}

func (l *List) GetFrom(x int, pos int) int {
	Debug("Get from ",pos)
	pos = pos % l.Len
	Debug(pos)
	if pos < 0 {
		pos += l.Len
	}
	Debugln(pos)


	cur := l.Find(x)
	if cur != nil {
		Debugln("jump after ",cur.Val)
		for i := 0; i < pos ; i++ {
			cur = cur.Next
		}
		Debugln("Found ",cur.Val)
		return cur.Val
	} else {
		Debugln("Not found ",x)
	}


	return -1
}

func Solution(input []string, pos ...int) int {
	src := IntSlice(input)
	dst := NewList(src)
	for i,x := range src {
		node := dst.FindIndex(i)
		Debugf("Try to move %d #%d \n", x, i)
		dst.Move(node)

		if dst.Len < 20 {
			Debugf("%d > %s\n", x, dst)
		}
	}
	Debugf("> %s\n", dst)


	var result int
	for _, p := range pos {
		x := dst.GetFrom(0, p)
		Debugf("%dth of xs is %d\n",p,x)
		result += x
	}

	return result
}

func Task01(input []string) string {
	return fmt.Sprint(Solution(input, 1000, 2000, 3000))

}

func Task02(input []string) string {
	return ""
}
