package doublelinkedlist

import (
	"fmt"
	"reflect"
)

type ItemInterface interface {
	GetValue() interface{}
	SetValue(interface{})
	Next() *Item
	Prev() *Item
	Remove() bool
}

type Item struct {
	list  *DoubleLinkedList
	value reflect.Value
}

func (i *Item) GetValue() interface{} {
	return i.value.Interface()
}

func (i *Item) SetValue(value interface{}) {
	i.value = reflect.ValueOf(value)
}

func (i *Item) Next() *Item {
	return i
}

func (i *Item) Prev() *Item {
	return i
}

func (i *Item) Remove() bool {
	return false
}

type DoubleLinkedListInterface interface {
	Len() int
	First() int
	Past() int
	PushFront(v interface{}) // добавить значение в начало
	PushBack(v interface{})  // добавить значение в конец
}

type DoubleLinkedList struct {
}

func (dll *DoubleLinkedList) Len() int {
	return 0
}

func (dll *DoubleLinkedList) First() *Item {
	return &Item{}
}

func (dll *DoubleLinkedList) Past() *Item {
	return &Item{}
}

func (dll *DoubleLinkedList) PushFront(v interface{}) {
	fmt.Println("PushFront")
}

func (dll *DoubleLinkedList) PushBack(v interface{}) {
	fmt.Println("PushBack")
}
