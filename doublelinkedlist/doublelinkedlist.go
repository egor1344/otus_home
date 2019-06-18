package doublelinkedlist

import (
	"log"
)

type Item struct {
	value interface{}
	list  *DoubleLinkedList
	next  *Item
	prev  *Item
}

func (i *Item) GetValue() interface{} {
	if i.value != nil {
		return i.value
	}
	return nil
}

func (i *Item) SetValue(value interface{}) {
	i.value = value
}

func (i *Item) Next() *Item {
	if i.next != nil {
		return i.next
	}
	return nil
}

func (i *Item) Prev() *Item {
	if i.prev != nil {
		return i.prev
	}
	return nil
}

func (i *Item) Remove() {
	if i.next != nil && i.prev != nil {
		i.next.prev = i.prev
		i.prev.next = i.next
	}
	if i.next == nil && i.prev != nil {
		i.prev.next = nil
	}
	if i.next != nil && i.prev == nil {
		i.next.prev = nil
	}
}

type DoubleLinkedList struct {
	first *Item
	last  *Item
}

func (dll *DoubleLinkedList) Len() int {
	if dll.first == nil {
		return 0
	}
	if dll.First().next == nil {
		return 1
	}
	counter := 1
	for j := dll.First(); j.next != nil; j = j.next {
		counter++
		log.Println(counter)
	}
	return counter
}

func (dll *DoubleLinkedList) First() *Item {
	if dll.first != nil {
		return dll.first
	}
	return nil
}

func (dll *DoubleLinkedList) Past() *Item {
	if dll.last != nil {
		return dll.last
	}
	return nil
}

func (dll *DoubleLinkedList) PushFront(i Item) {
	if dll.first == nil {
		dll.first = &i
	} else {
		if dll.first.prev == nil {
			dll.first.prev = &i
			dll.first = &i
		}
	}
	if dll.last == nil {
		dll.last = &i
	}
	if dll.first.list == nil {
		dll.first.list = dll
	}

}

func (dll *DoubleLinkedList) PushBack(i Item) {
	if dll.first == nil {
		dll.first = &i
	}
	if dll.last == nil {
		dll.last = &i
	} else {
		if dll.last.next == nil {
			dll.last.next = &i
			dll.last = &i
		}
	}
	if dll.last.list == nil {
		dll.last.list = dll
	}
}
