package main

import "fmt"

type node struct {
	data string
	next *node
	prev *node
}

type linkedList struct {
	head         *node
	tail         *node
	length       int
	maxCacheSize int
}

func (l *linkedList) fillDefaults() {
	if l.maxCacheSize == 0 {
		l.maxCacheSize = 3
	}
}

func (l *linkedList) prepend(n *node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
		n.prev = nil
		n.next = nil
	} else if l.length == 1 {
		l.head.next = nil
	}

	l.head.prev = n
	n.next = l.head
	n.prev = nil
	l.head = n
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length > 0 {
		fmt.Printf("%s ", toPrint.data)
		toPrint = toPrint.next
		l.length -= 1
	}
	fmt.Printf("\n")
}

func (l *linkedList) removeTail() {
	l.tail = l.tail.prev
	l.tail.next = nil
	l.length--
}

func main() {
	cache := linkedList{}
	hash := map[string]*node{}

	cache.fillDefaults()

	clientRequest := func(req string) *node {
		var currentNode *node

		if hash[req] != nil {
			currentNode = hash[req]

			if currentNode == cache.head {
				return currentNode
			}

			if currentNode == cache.tail {
				cache.tail = currentNode.prev
			}

			currentNode.prev.next = currentNode.next

			if currentNode.next != nil {
				currentNode.next.prev = currentNode.prev
			}

			cache.length--
			cache.prepend(currentNode)
		} else {
			if cache.length == cache.maxCacheSize {
				lastNode := cache.tail
				cache.removeTail()
				delete(hash, lastNode.data)
			}

			newNode := &node{data: req, next: nil, prev: nil}
			cache.prepend(newNode)
			hash[req] = newNode
			currentNode = newNode
		}
		return currentNode
	}

	clientRequest("chocolate")
	clientRequest("vanilla")
	clientRequest("strawberry")
	cache.printListData() // strawberry vanilla chocolate

	clientRequest("chocolate")
	cache.printListData() // chocolate strawberry vanilla

	clientRequest("pound")
	cache.printListData() // pound chocolate strawberry

	clientRequest("chocolate")
	cache.printListData() // chocolate pound strawberry

	clientRequest("marble")
	cache.printListData() // marble chocolate pound
}
