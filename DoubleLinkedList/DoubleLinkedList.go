package DoubleLinkedList

import "fmt"

type Node struct {
	data       interface{}
	next, prev *Node
}

type DoubleLinkedList struct {
	head, tail *Node
	Size       uint64
}

func (d *DoubleLinkedList) AddHead(value interface{}) {
	newNode := &Node{data: value}
	if d.head == nil && d.tail == nil {
		//first node being added
		d.head = newNode
		d.tail = newNode
		newNode.next = nil
		newNode.prev = nil
		d.Size = 1
	} else {
		//list is not empty
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
		d.Size ++
	}
}

func (d *DoubleLinkedList) RemoveTail() {
	currentNode := d.tail
	d.tail = currentNode.prev
	d.tail.next = nil
	currentNode = nil
}

func (d *DoubleLinkedList) Remove(value interface{}) {
	currentNode := d.head
	for ; currentNode.data != value; currentNode = currentNode.next {
		if currentNode == nil {
			return
		}
	}

	//data found
	if currentNode.prev == nil && currentNode.next == nil {
		//only on node in list, delete and set head and tail to nil
		d.head = nil
		d.tail = nil
	} else if currentNode.prev == nil && currentNode.next != nil {
		//deleting a head node
		d.head = currentNode.next
		d.head.prev = nil
	} else if currentNode.prev != nil && currentNode.next == nil {
		//deleting a tail node
		d.tail = currentNode.prev
		d.tail.next = nil
	} else {
		//deleting a node in between head and tail
		currentNode.prev.next = currentNode.next
		currentNode.next.prev = currentNode.prev
	}

	currentNode = nil
	d.Size--
}

func (d *DoubleLinkedList) Print() {
	fmt.Println("\nL - R")
	currentNode := d.head
	for ; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("%v ", currentNode.data)
	}
	fmt.Println("\nR - L")
	currentNode = d.tail
	for ; currentNode != nil; currentNode = currentNode.prev {
		fmt.Printf("%v ", currentNode.data)
	}
}

func (d *DoubleLinkedList) Iterate() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		currentNode := d.head
		for ; currentNode != nil; currentNode = currentNode.next {
			c <- currentNode.data
		}
		close(c)
	}()
	return c
}
